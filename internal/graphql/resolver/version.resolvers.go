package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lrstanley/clix"
)

// Version is the resolver for the version field.
func (r *queryResolver) Version(ctx context.Context) (*clix.NonSensitiveVersion, error) {
	return r.cli.GetVersionInfo().NonSensitive(), nil
}
