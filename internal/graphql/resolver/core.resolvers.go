package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lrstanley/liam.sh/internal/graphql/gqlhandler"
)

// Ping is the resolver for the ping field.
func (r *mutationResolver) Ping(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns gqlhandler.MutationResolver implementation.
func (r *Resolver) Mutation() gqlhandler.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
