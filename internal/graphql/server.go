// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package graphql

import (
	"bytes"
	"embed"
	"fmt"
	"path/filepath"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/lrstanley/clix"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/graphql/resolver"
	"github.com/lrstanley/liam.sh/internal/models"
)

//go:embed schema/*.gql
var rawSchema embed.FS

var Schema string

func init() {
	files, err := rawSchema.ReadDir("schema")
	if err != nil {
		panic(err)
	}

	buf := bytes.Buffer{}

	for i, file := range files {
		if file.IsDir() {
			continue
		}

		path := filepath.Join("schema", file.Name())

		f, err := rawSchema.Open(path)
		if err != nil {
			panic(err)
		}

		if i != 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(fmt.Sprintf("# source: %s\n", path))
		_, _ = buf.ReadFrom(f)
	}

	Schema = buf.String()
}

func New(db *ent.Client, cli *clix.CLI[models.Flags]) *handler.Server {
	srv := handler.NewDefaultServer(resolver.NewSchema(db, cli))
	srv.AroundOperations(requestLogger)
	srv.SetRecoverFunc(recoverFunc)
	srv.AroundOperations(injectClient(db))
	srv.SetErrorPresenter(errorPresenter)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	if !cli.Debug {
		srv.SetQueryCache(lru.New(512))
		srv.Use(extension.FixedComplexityLimit(100))
	}

	srv.Use(extension.Introspection{})
	srv.Use(entgql.Transactioner{
		TxOpener: db,
	})

	return srv
}
