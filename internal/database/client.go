// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"
	"database/sql"
	"iter"
	"net/url"
	"sync"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/apex/log"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	_ "github.com/lrstanley/liam.sh/internal/database/ent/runtime" // required by ent.
	"github.com/lrstanley/liam.sh/internal/models"
	sqlite "modernc.org/sqlite"
)

//go:generate go run -mod=readonly entc.go

var defaultConnectionValues = url.Values{
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

	db, err := sql.Open("sqlite3", uri.String())
	if err != nil {
		log.FromContext(ctx).WithError(err).Fatal("failed to open database connection")
		return nil
	}
	db.SetMaxOpenConns(1)

	return ent.NewClient(ent.Driver(entsql.OpenDB(dialect.SQLite, db)))
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
