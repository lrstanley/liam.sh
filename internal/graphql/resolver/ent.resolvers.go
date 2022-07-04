package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lrstanley/liam.sh/internal/database"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/graphql/gqlhandler"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Githubassets is the resolver for the githubassets field.
func (r *queryResolver) Githubassets(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubAssetOrder, where *ent.GithubAssetWhereInput) (*ent.GithubAssetConnection, error) {
	return r.client.GithubAsset.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubAssetOrder(orderBy),
		ent.WithGithubAssetFilter(where.Filter),
	)
}

// Githubevents is the resolver for the githubevents field.
func (r *queryResolver) Githubevents(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubEventOrder, where *ent.GithubEventWhereInput) (*ent.GithubEventConnection, error) {
	return r.client.GithubEvent.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubEventOrder(orderBy),
		ent.WithGithubEventFilter(where.Filter),
	)
}

// Githubgists is the resolver for the githubgists field.
func (r *queryResolver) Githubgists(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubGistOrder, where *ent.GithubGistWhereInput) (*ent.GithubGistConnection, error) {
	return r.client.GithubGist.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubGistOrder(orderBy),
		ent.WithGithubGistFilter(where.Filter),
	)
}

// Githubreleases is the resolver for the githubreleases field.
func (r *queryResolver) Githubreleases(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubReleaseOrder, where *ent.GithubReleaseWhereInput) (*ent.GithubReleaseConnection, error) {
	return r.client.GithubRelease.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubReleaseOrder(orderBy),
		ent.WithGithubReleaseFilter(where.Filter),
	)
}

// Githubrepositories is the resolver for the githubrepositories field.
func (r *queryResolver) Githubrepositories(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubRepositoryOrder, where *ent.GithubRepositoryWhereInput) (*ent.GithubRepositoryConnection, error) {
	return r.client.GithubRepository.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubRepositoryOrder(orderBy),
		ent.WithGithubRepositoryFilter(where.Filter),
	)
}

// Labels is the resolver for the labels field.
func (r *queryResolver) Labels(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.LabelOrder, where *ent.LabelWhereInput) (*ent.LabelConnection, error) {
	return r.client.Label.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithLabelOrder(orderBy),
		ent.WithLabelFilter(where.Filter),
	)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	conn, err := r.client.Post.Query().Paginate(
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
	return r.client.User.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithUserOrder(orderBy),
		ent.WithUserFilter(where.Filter),
	)
}

// Query returns gqlhandler.QueryResolver implementation.
func (r *Resolver) Query() gqlhandler.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
