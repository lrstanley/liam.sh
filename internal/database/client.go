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

	atlas "ariga.io/atlas/sql/migrate"
	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/apex/log"
	pgx "github.com/jackc/pgx/v5/stdlib"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/migrate"
	_ "github.com/lrstanley/liam.sh/internal/database/ent/runtime"
	"github.com/lrstanley/liam.sh/internal/models"
)

func init() {
	// TODO: https://github.com/ariga/atlas/issues/1415
	sql.Register("postgres", pgx.GetDefaultDriver())
}

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
			tctx, cancel := context.WithTimeout(ctx, 5*time.Second)

			err = db.PingContext(tctx)
			cancel()
		}

		if err != nil {
			logger.WithError(err).WithField("attempt", attempt).Error("failed to open database connection")

			if attempt > 2 {
				logger.Fatal("failed to open database connection after 3 attempts")
			}
			time.Sleep(time.Second * 5)
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
			entcache.TTL(120*time.Second),
			entcache.Levels(entcache.NewLRU(256)),
		),
	))
}

func GenerateMigrations(ctx context.Context, logger log.Interface, config models.ConfigDatabase, migrationPath, name string) {
	logger.Info("generating schema migrations")

	uri, err := ParseURL(config)
	if err != nil {
		logger.WithError(err).Fatal("failed to parse database url")
	}

	log.Info(uri)

	dir, err := atlas.NewLocalDir(migrationPath)
	if err != nil {
		logger.WithError(err).Fatal("failed to initialize migrations directory")
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithGlobalUniqueID(true),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	}

	err = migrate.NamedDiff(ctx, uri, name, opts...)
	if err != nil {
		logger.WithError(err).Fatal("failed to generate migration")
	}

	logger.Info("schema migration generated")
}
