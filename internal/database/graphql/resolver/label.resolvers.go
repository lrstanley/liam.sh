package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"

	"github.com/lrstanley/liam.sh/internal/ent"
)

// CreateLabel is the resolver for the createLabel field.
func (r *mutationResolver) CreateLabel(ctx context.Context, input ent.CreateLabelInput) (*ent.Label, error) {
	return ent.FromContext(ctx).Label.Create().SetInput(input).Save(ctx)
}

// UpdateLabel is the resolver for the updateLabel field.
func (r *mutationResolver) UpdateLabel(ctx context.Context, id int, input ent.UpdateLabelInput) (*ent.Label, error) {
	return ent.FromContext(ctx).Label.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteLabel is the resolver for the deleteLabel field.
func (r *mutationResolver) DeleteLabel(ctx context.Context, id int) (int, error) {
	return id, ent.FromContext(ctx).Label.DeleteOneID(id).Exec(ctx)
}
