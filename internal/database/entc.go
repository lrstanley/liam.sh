//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/lrstanley/entrest"
)

const header = `// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.`

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rest, err := entrest.NewExtension(&entrest.Config{
		SpecFromPath:          "../cmd/httpserver/base-openapi.json",
		MaxItemsPerPage:       1000,
		Handler:               entrest.HandlerChi,
		StrictMutate:          true,
		GlobalRequestHeaders:  entrest.RequestIDHeader,
		GlobalResponseHeaders: entrest.RateLimitHeaders,
	})
	checkError(err)

	err = entc.Generate(
		"./database/schema",
		&gen.Config{
			Target:  "./database/ent/",
			Schema:  "github.com/lrstanley/liam.sh/internal/database/schema",
			Package: "github.com/lrstanley/liam.sh/internal/database/ent",
			Header:  header,
			Features: []gen.Feature{
				gen.FeaturePrivacy,
				gen.FeatureEntQL,
				gen.FeatureSnapshot,
				gen.FeatureUpsert,
				gen.FeatureIntercept,
			},
		},
		entc.Extensions(rest),
	)
	checkError(err)
}
