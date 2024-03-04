// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (ga *GithubAsset) Release(ctx context.Context) (*GithubRelease, error) {
	result, err := ga.Edges.ReleaseOrErr()
	if IsNotLoaded(err) {
		result, err = ga.QueryRelease().Only(ctx)
	}
	return result, err
}

func (gr *GithubRelease) Repository(ctx context.Context) (*GithubRepository, error) {
	result, err := gr.Edges.RepositoryOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QueryRepository().Only(ctx)
	}
	return result, err
}

func (gr *GithubRelease) Assets(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *GithubAssetOrder, where *GithubAssetWhereInput,
) (*GithubAssetConnection, error) {
	opts := []GithubAssetPaginateOption{
		WithGithubAssetOrder(orderBy),
		WithGithubAssetFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := gr.Edges.totalCount[1][alias]
	if nodes, err := gr.NamedAssets(alias); err == nil || hasTotalCount {
		pager, err := newGithubAssetPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &GithubAssetConnection{Edges: []*GithubAssetEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return gr.QueryAssets().Paginate(ctx, after, first, before, last, opts...)
}

func (gr *GithubRepository) Labels(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *LabelOrder, where *LabelWhereInput,
) (*LabelConnection, error) {
	opts := []LabelPaginateOption{
		WithLabelOrder(orderBy),
		WithLabelFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := gr.Edges.totalCount[0][alias]
	if nodes, err := gr.NamedLabels(alias); err == nil || hasTotalCount {
		pager, err := newLabelPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &LabelConnection{Edges: []*LabelEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return gr.QueryLabels().Paginate(ctx, after, first, before, last, opts...)
}

func (gr *GithubRepository) Releases(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *GithubReleaseOrder, where *GithubReleaseWhereInput,
) (*GithubReleaseConnection, error) {
	opts := []GithubReleasePaginateOption{
		WithGithubReleaseOrder(orderBy),
		WithGithubReleaseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := gr.Edges.totalCount[1][alias]
	if nodes, err := gr.NamedReleases(alias); err == nil || hasTotalCount {
		pager, err := newGithubReleasePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &GithubReleaseConnection{Edges: []*GithubReleaseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return gr.QueryReleases().Paginate(ctx, after, first, before, last, opts...)
}

func (l *Label) Posts(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PostOrder, where *PostWhereInput,
) (*PostConnection, error) {
	opts := []PostPaginateOption{
		WithPostOrder(orderBy),
		WithPostFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := l.Edges.totalCount[0][alias]
	if nodes, err := l.NamedPosts(alias); err == nil || hasTotalCount {
		pager, err := newPostPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &PostConnection{Edges: []*PostEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return l.QueryPosts().Paginate(ctx, after, first, before, last, opts...)
}

func (l *Label) GithubRepositories(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *GithubRepositoryOrder, where *GithubRepositoryWhereInput,
) (*GithubRepositoryConnection, error) {
	opts := []GithubRepositoryPaginateOption{
		WithGithubRepositoryOrder(orderBy),
		WithGithubRepositoryFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := l.Edges.totalCount[1][alias]
	if nodes, err := l.NamedGithubRepositories(alias); err == nil || hasTotalCount {
		pager, err := newGithubRepositoryPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &GithubRepositoryConnection{Edges: []*GithubRepositoryEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return l.QueryGithubRepositories().Paginate(ctx, after, first, before, last, opts...)
}

func (po *Post) Author(ctx context.Context) (*User, error) {
	result, err := po.Edges.AuthorOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryAuthor().Only(ctx)
	}
	return result, err
}

func (po *Post) Labels(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *LabelOrder, where *LabelWhereInput,
) (*LabelConnection, error) {
	opts := []LabelPaginateOption{
		WithLabelOrder(orderBy),
		WithLabelFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := po.Edges.totalCount[1][alias]
	if nodes, err := po.NamedLabels(alias); err == nil || hasTotalCount {
		pager, err := newLabelPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &LabelConnection{Edges: []*LabelEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return po.QueryLabels().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) Posts(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PostOrder, where *PostWhereInput,
) (*PostConnection, error) {
	opts := []PostPaginateOption{
		WithPostOrder(orderBy),
		WithPostFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[0][alias]
	if nodes, err := u.NamedPosts(alias); err == nil || hasTotalCount {
		pager, err := newPostPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &PostConnection{Edges: []*PostEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryPosts().Paginate(ctx, after, first, before, last, opts...)
}
