// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"
	"log/slog"
	"net/url"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	_ "github.com/lrstanley/liam.sh/internal/database/ent/runtime" // required by ent.
	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/lrstanley/x/entx"
)

//go:generate go run -mod=readonly entc.go

// Open new database connection. See [entx.Driver] for supported DSN schemes.
func Open(ctx context.Context, config models.ConfigDatabase) *ent.Client {
	dsn, err := url.Parse(config.URL)
	if err != nil {
		panic(err)
	}

	return ent.NewClient(ent.Driver(entx.MustDriver(ctx, slog.Default(), dsn)))
}

func Migrate(ctx context.Context, db *ent.Client) {
	slog.InfoContext(ctx, "initiating database schema migration")
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
		slog.ErrorContext(ctx, "failed to create schema", "error", err)
		panic(err)
	}
	slog.InfoContext(ctx, "database schema migration complete")
}
