package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/graphql/gqlhandler"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Labels(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.LabelOrder, where *ent.LabelWhereInput) (*ent.LabelConnection, error) {
	return r.client.Label.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithLabelOrder(orderBy),
			ent.WithLabelFilter(where.Filter),
		)
}

func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	return r.client.Post.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithPostOrder(orderBy),
			ent.WithPostFilter(where.Filter),
		)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(
			ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
}

// Query returns gqlhandler.QueryResolver implementation.
func (r *Resolver) Query() gqlhandler.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
