// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"
	"database/sql"
	"net/url"
	"sync"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/apex/log"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	_ "github.com/lrstanley/liam.sh/internal/database/ent/runtime" // required by ent.
	"github.com/lrstanley/liam.sh/internal/models"
	sqlite "modernc.org/sqlite"
)

var defaultConnectionValues = url.Values{
	"cache": {"shared"},
	"_pragma": {
		"foreign_keys(1)",
		"journal_mode(WAL)",
		"temp_store(2)",
		"synchronous(1)",
		"mmap_size(30000000000)",
	},
	"_busy_timeout": {"15"},
}

var sqliteInit sync.Once

// Open new postgres connection.
func Open(ctx context.Context, config models.ConfigDatabase) *ent.Client {
	sqliteInit.Do(func() {
		sql.Register("sqlite3", &sqlite.Driver{})
	})

	uri, err := url.Parse(config.URL)
	if err != nil {
		log.FromContext(ctx).WithError(err).Fatal("failed to parse database connection")
		return nil
	}

	values := uri.Query()
	for k, v := range defaultConnectionValues {
		if _, ok := values[k]; !ok {
			values[k] = v
		}
	}
	uri.RawQuery = values.Encode()

	db, err := ent.Open("sqlite3", uri.String())
	if err != nil {
		log.FromContext(ctx).WithError(err).Fatal("failed to open database connection")
		return nil
	}

	log.FromContext(ctx).Info("connected to database")
	return db
}

func Migrate(ctx context.Context, db *ent.Client) {
	log.FromContext(ctx).Info("initiating database schema migration")
	if db == nil {
		panic("database client is nil")
	}

	if err := db.Schema.Create(
		ctx,
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		schema.WithGlobalUniqueID(true),
		schema.WithForeignKeys(true),
	); err != nil {
		log.FromContext(ctx).WithError(err).Fatal("failed to create schema")
	}
	log.FromContext(ctx).Info("database schema migration complete")
}
