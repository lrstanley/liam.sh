package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/lrstanley/clix"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/graphql/gqlhandler"
	"github.com/lrstanley/liam.sh/internal/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the GQL resolver root.
type Resolver struct {
	client *ent.Client
	cli    *clix.CLI[models.Flags]
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, cli *clix.CLI[models.Flags]) graphql.ExecutableSchema {
	return gqlhandler.NewExecutableSchema(gqlhandler.Config{
		Resolvers: &Resolver{
			client: client,
			cli:    cli,
		},
	})
}
