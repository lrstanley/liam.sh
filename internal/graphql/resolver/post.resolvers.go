package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/apex/log"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/ent/post"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput) (*ent.Post, error) {
	input.AuthorID = chix.IDFromContext[int](ctx)
	return ent.FromContext(ctx).Post.Create().SetInput(input).Save(ctx)
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input ent.UpdatePostInput) (*ent.Post, error) {
	return ent.FromContext(ctx).Post.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id int) (int, error) {
	return id, ent.FromContext(ctx).Post.DeleteOneID(id).Exec(ctx)
}

// RegeneratePosts is the resolver for the regeneratePosts field.
func (r *mutationResolver) RegeneratePosts(ctx context.Context) (bool, error) {
	p := []struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	}{}

	tx, err := ent.FromContext(ctx).Tx(ctx)
	if err != nil {
		return false, err
	}

	if err = tx.Post.Query().Select(post.FieldID, post.FieldContent).Scan(ctx, &p); err != nil {
		_ = tx.Rollback()
		return false, err
	}

	for _, v := range p {
		if v.Content == "" {
			continue
		}

		if err = tx.Post.UpdateOneID(v.ID).SetContent(v.Content).Exec(ctx); err != nil {
			_ = tx.Rollback()
			return false, err
		}
		log.FromContext(ctx).WithField("id", v.ID).Debug("regenerated post")
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return false, err
	}
	return true, nil
}
