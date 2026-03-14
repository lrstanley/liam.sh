// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"
	"iter"
	"log/slog"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	_ "github.com/lrstanley/liam.sh/internal/database/ent/runtime" // required by ent.
	"github.com/lrstanley/liam.sh/internal/models"
)

//go:generate go run -mod=readonly entc.go

// Open new postgres connection.
func Open(ctx context.Context, config models.ConfigDatabase) *ent.Client {
	poolConfig, err := pgxpool.ParseConfig(config.URL)
	if err != nil {
		panic(err)
	}

	poolConfig.PingTimeout = 5 * time.Second

	db, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		panic(err)
	}

	return ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, stdlib.OpenDBFromPool(db))))
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

type PagableQuery[P any, T any] interface {
	Limit(int) P
	Offset(int) P
	All(context.Context) ([]*T, error)
}

func Chunked[P PagableQuery[P, T], T any](
	ctx context.Context,
	size int,
	query PagableQuery[P, T],
) iter.Seq2[*T, error] {
	if size <= 1 {
		size = 250
	}

	return func(yield func(*T, error) bool) {
		var offset int
		var results []*T
		var err error

		for {
			results, err = query.Limit(size).Offset(offset).All(ctx)
			if err != nil {
				yield(nil, err)
				return
			}

			for _, result := range results {
				if !yield(result, nil) {
					return
				}
			}

			if len(results) < size {
				return
			}

			offset += size
		}
	}
}
