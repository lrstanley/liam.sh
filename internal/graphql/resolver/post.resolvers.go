package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput) (*ent.Post, error) {
	input.AuthorID = chix.IDFromContext[int](ctx)
	return r.client.Post.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input ent.UpdatePostInput) (*ent.Post, error) {
	return r.client.Post.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *mutationResolver) DeletePost(ctx context.Context, id int) (int, error) {
	return id, r.client.Post.DeleteOneID(id).Exec(ctx)
}
