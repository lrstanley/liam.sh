// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"
	"database/sql"
	"sync"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/apex/log"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	_ "github.com/lrstanley/liam.sh/internal/database/ent/runtime" // required by ent.
	"github.com/lrstanley/liam.sh/internal/models"
	sqlite "modernc.org/sqlite"
)

var sqliteInit sync.Once

// Open new postgres connection.
func Open(ctx context.Context, config models.ConfigDatabase) *ent.Client {
	sqliteInit.Do(func() {
		sql.Register("sqlite3", &sqlite.Driver{})
	})

	db, err := ent.Open("sqlite3", config.URL)
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
