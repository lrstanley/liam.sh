package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lrstanley/liam.sh/internal/ent"
)

func (r *mutationResolver) CreateLabel(ctx context.Context, input ent.CreateLabelInput) (*ent.Label, error) {
	return r.client.Label.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateLabel(ctx context.Context, id int, input ent.UpdateLabelInput) (*ent.Label, error) {
	return r.client.Label.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeleteLabel(ctx context.Context, id int) (int, error) {
	return id, r.client.Label.DeleteOneID(id).Exec(ctx)
}
