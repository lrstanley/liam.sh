package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent"
)

// Self is the resolver for the self field.
func (r *queryResolver) Self(ctx context.Context) (*ent.User, error) {
	user := chix.IdentFromContext[ent.User](ctx)
	if user == nil {
		return nil, fmt.Errorf("not authenticated")
	}

	return user, nil
}
