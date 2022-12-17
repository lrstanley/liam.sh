package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/database/graphql/gqlhandler"
	"github.com/lrstanley/liam.sh/internal/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return ent.FromContext(ctx).Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return ent.FromContext(ctx).Noders(ctx, ids)
}

// GithubAssets is the resolver for the githubAssets field.
func (r *queryResolver) GithubAssets(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubAssetOrder, where *ent.GithubAssetWhereInput) (*ent.GithubAssetConnection, error) {
	return ent.FromContext(ctx).GithubAsset.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubAssetOrder(orderBy),
		ent.WithGithubAssetFilter(where.Filter),
	)
}

// GithubEvents is the resolver for the githubEvents field.
func (r *queryResolver) GithubEvents(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubEventOrder, where *ent.GithubEventWhereInput) (*ent.GithubEventConnection, error) {
	return ent.FromContext(ctx).GithubEvent.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubEventOrder(orderBy),
		ent.WithGithubEventFilter(where.Filter),
	)
}

// GithubGists is the resolver for the githubGists field.
func (r *queryResolver) GithubGists(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubGistOrder, where *ent.GithubGistWhereInput) (*ent.GithubGistConnection, error) {
	return ent.FromContext(ctx).GithubGist.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubGistOrder(orderBy),
		ent.WithGithubGistFilter(where.Filter),
	)
}

// GithubReleases is the resolver for the githubReleases field.
func (r *queryResolver) GithubReleases(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubReleaseOrder, where *ent.GithubReleaseWhereInput) (*ent.GithubReleaseConnection, error) {
	return ent.FromContext(ctx).GithubRelease.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubReleaseOrder(orderBy),
		ent.WithGithubReleaseFilter(where.Filter),
	)
}

// GithubRepositories is the resolver for the githubRepositories field.
func (r *queryResolver) GithubRepositories(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubRepositoryOrder, where *ent.GithubRepositoryWhereInput) (*ent.GithubRepositoryConnection, error) {
	return ent.FromContext(ctx).GithubRepository.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubRepositoryOrder(orderBy),
		ent.WithGithubRepositoryFilter(where.Filter),
	)
}

// Labels is the resolver for the labels field.
func (r *queryResolver) Labels(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.LabelOrder, where *ent.LabelWhereInput) (*ent.LabelConnection, error) {
	return ent.FromContext(ctx).Label.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithLabelOrder(orderBy),
		ent.WithLabelFilter(where.Filter),
	)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	conn, err := ent.FromContext(ctx).Post.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithPostOrder(orderBy),
		ent.WithPostFilter(where.Filter),
	)

	if err == nil && conn.TotalCount == 1 {
		post := conn.Edges[0].Node

		if post.ContentHTML == "" {
			return conn, nil
		}

		go database.PostViewCounter(ctx, post)
	}

	return conn, err
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return ent.FromContext(ctx).User.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithUserOrder(orderBy),
		ent.WithUserFilter(where.Filter),
	)
}

// Query returns gqlhandler.QueryResolver implementation.
func (r *Resolver) Query() gqlhandler.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
