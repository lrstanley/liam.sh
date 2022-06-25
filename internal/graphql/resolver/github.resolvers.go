package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/go-github/v44/github"
	"github.com/lrstanley/liam.sh/internal/gh"
)

func (r *queryResolver) GithubUser(ctx context.Context) (*github.User, error) {
	user := gh.User.Load()
	if user == nil {
		return nil, fmt.Errorf("information not available yet")
	}

	return user, nil
}
