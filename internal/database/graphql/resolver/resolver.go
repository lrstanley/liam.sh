// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/lrstanley/clix"
	"github.com/lrstanley/liam.sh/internal/database/graphql/gqlhandler"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the GQL resolver root.
type Resolver struct {
	cli *clix.CLI[models.Flags]
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, cli *clix.CLI[models.Flags]) graphql.ExecutableSchema {
	return gqlhandler.NewExecutableSchema(gqlhandler.Config{
		Resolvers: &Resolver{
			cli: cli,
		},
	})
}
