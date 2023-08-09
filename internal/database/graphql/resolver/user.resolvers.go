package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/database/ent"
)

// Self is the resolver for the self field.
func (r *queryResolver) Self(ctx context.Context) (*ent.User, error) {
	user := chix.IdentFromContext[ent.User](ctx)
	if user == nil {
		return nil, fmt.Errorf("not authenticated")
	}

	return user, nil
}
