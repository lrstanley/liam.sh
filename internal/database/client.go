// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/apex/log"
	_ "github.com/jackc/pgx/v5/stdlib" // load postgres driver.
	"github.com/lrstanley/liam.sh/internal/database/ent"
	_ "github.com/lrstanley/liam.sh/internal/database/ent/runtime" // required by ent.
	"github.com/lrstanley/liam.sh/internal/models"
)

const (
	dbCacheTTL        = 120 * time.Second
	dbCacheMaxEntries = 256
)

var Ping func(context.Context) error

func ParseURL(config models.ConfigDatabase) (string, error) {
	uri, err := url.Parse(config.URL)
	if err != nil {
		return "", fmt.Errorf("failed to parse database url: %w", err)
	}

	var username, password string
	if uri.User != nil {
		username = uri.User.Username()
		password, _ = uri.User.Password()
	}

	if config.Database != "" {
		uri.Path = config.Database
	}

	if config.Username != "" {
		username = config.Username
	}

	if config.Password != "" {
		password = config.Password
	}

	if username != "" {
		uri.User = url.UserPassword(username, password)
	}

	return uri.String(), nil
}

// Open new postgres connection.
func Open(ctx context.Context, logger log.Interface, config models.ConfigDatabase) *ent.Client {
	var db *sql.DB

	uri, err := ParseURL(config)
	if err != nil {
		logger.WithError(err).Fatal("failed to parse database url")
	}

	var attempt int
	for {
		attempt++
		db, err = sql.Open("pgx", uri)

		if err == nil {
			tctx, cancel := context.WithTimeout(ctx, 5*time.Second) // nolint:gomnd

			err = db.PingContext(tctx)
			cancel()
		}

		if err != nil {
			logger.WithError(err).WithField("attempt", attempt).Error("failed to open database connection")

			if attempt > 2 { // nolint:gomnd
				logger.Fatal("failed to open database connection after 3 attempts")
			}
			time.Sleep(time.Second * 5) // nolint:gomnd
			continue
		}
		break
	}

	logger.Info("connected to database")
	Ping = db.PingContext

	// Create an ent.Driver from db.
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(
		entcache.NewDriver(
			driver,
			entcache.TTL(dbCacheTTL),
			entcache.Levels(entcache.NewLRU(dbCacheMaxEntries)),
		),
	))
}

func Migrate(ctx context.Context, logger log.Interface) {
	logger.Info("initiating database schema migration")
	db := ent.FromContext(ctx)
	if db == nil {
		panic("database client is nil")
	}

	if err := db.Schema.Create(
		entcache.Skip(ctx),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		schema.WithGlobalUniqueID(true),
	); err != nil {
		logger.WithError(err).Fatal("failed to create schema")
	}

	logger.Info("database schema migration complete")
}
