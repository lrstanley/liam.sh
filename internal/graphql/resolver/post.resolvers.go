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

func (r *mutationResolver) RegeneratePosts(ctx context.Context) (bool, error) {
	p := []struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	}{}

	tx, err := r.client.Tx(ctx)
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