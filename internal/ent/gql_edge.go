// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (gr *GithubRepository) Labels(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *LabelOrder, where *LabelWhereInput,
	opts ...LabelPaginateOption,
) (*LabelConnection, error) {
	totalCount := gr.Edges.totalCount[0]
	if nodes, err := gr.Edges.LabelsOrErr(); err == nil {
		conn := &LabelConnection{Edges: []*LabelEdge{}}
		if totalCount != nil {
			conn.TotalCount = *totalCount
		}
		opts = append(opts, WithLabelOrder(orderBy))
		pager, err := newLabelPager(opts)
		if err != nil {
			return nil, err
		}
		conn.build(nodes, pager, first, last)
		return conn, nil
	}
	query := gr.QueryLabels()
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newLabelPager(opts)
	if err != nil {
		return nil, err
	}
	if query, err = pager.applyFilter(query); err != nil {
		return nil, err
	}
	conn := &LabelConnection{Edges: []*LabelEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
			if totalCount != nil {
				conn.TotalCount = *totalCount
			} else if conn.TotalCount, err = query.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := query.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	query = pager.applyCursors(query, after, before)
	query = pager.applyOrder(query, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		query.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := query.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := query.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}
	conn.build(nodes, pager, first, last)
	return conn, nil
}

func (l *Label) Posts(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PostOrder, where *PostWhereInput,
	opts ...PostPaginateOption,
) (*PostConnection, error) {
	totalCount := l.Edges.totalCount[0]
	if nodes, err := l.Edges.PostsOrErr(); err == nil {
		conn := &PostConnection{Edges: []*PostEdge{}}
		if totalCount != nil {
			conn.TotalCount = *totalCount
		}
		opts = append(opts, WithPostOrder(orderBy))
		pager, err := newPostPager(opts)
		if err != nil {
			return nil, err
		}
		conn.build(nodes, pager, first, last)
		return conn, nil
	}
	query := l.QueryPosts()
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newPostPager(opts)
	if err != nil {
		return nil, err
	}
	if query, err = pager.applyFilter(query); err != nil {
		return nil, err
	}
	conn := &PostConnection{Edges: []*PostEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
			if totalCount != nil {
				conn.TotalCount = *totalCount
			} else if conn.TotalCount, err = query.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := query.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	query = pager.applyCursors(query, after, before)
	query = pager.applyOrder(query, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		query.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := query.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := query.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}
	conn.build(nodes, pager, first, last)
	return conn, nil
}

func (l *Label) GithubRepositories(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *GithubRepositoryOrder, where *GithubRepositoryWhereInput,
	opts ...GithubRepositoryPaginateOption,
) (*GithubRepositoryConnection, error) {
	totalCount := l.Edges.totalCount[1]
	if nodes, err := l.Edges.GithubRepositoriesOrErr(); err == nil {
		conn := &GithubRepositoryConnection{Edges: []*GithubRepositoryEdge{}}
		if totalCount != nil {
			conn.TotalCount = *totalCount
		}
		opts = append(opts, WithGithubRepositoryOrder(orderBy))
		pager, err := newGithubRepositoryPager(opts)
		if err != nil {
			return nil, err
		}
		conn.build(nodes, pager, first, last)
		return conn, nil
	}
	query := l.QueryGithubRepositories()
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newGithubRepositoryPager(opts)
	if err != nil {
		return nil, err
	}
	if query, err = pager.applyFilter(query); err != nil {
		return nil, err
	}
	conn := &GithubRepositoryConnection{Edges: []*GithubRepositoryEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
			if totalCount != nil {
				conn.TotalCount = *totalCount
			} else if conn.TotalCount, err = query.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := query.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	query = pager.applyCursors(query, after, before)
	query = pager.applyOrder(query, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		query.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := query.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := query.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}
	conn.build(nodes, pager, first, last)
	return conn, nil
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
	opts ...LabelPaginateOption,
) (*LabelConnection, error) {
	totalCount := po.Edges.totalCount[1]
	if nodes, err := po.Edges.LabelsOrErr(); err == nil {
		conn := &LabelConnection{Edges: []*LabelEdge{}}
		if totalCount != nil {
			conn.TotalCount = *totalCount
		}
		opts = append(opts, WithLabelOrder(orderBy))
		pager, err := newLabelPager(opts)
		if err != nil {
			return nil, err
		}
		conn.build(nodes, pager, first, last)
		return conn, nil
	}
	query := po.QueryLabels()
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newLabelPager(opts)
	if err != nil {
		return nil, err
	}
	if query, err = pager.applyFilter(query); err != nil {
		return nil, err
	}
	conn := &LabelConnection{Edges: []*LabelEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
			if totalCount != nil {
				conn.TotalCount = *totalCount
			} else if conn.TotalCount, err = query.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := query.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	query = pager.applyCursors(query, after, before)
	query = pager.applyOrder(query, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		query.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := query.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := query.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}
	conn.build(nodes, pager, first, last)
	return conn, nil
}

func (u *User) Posts(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PostOrder, where *PostWhereInput,
	opts ...PostPaginateOption,
) (*PostConnection, error) {
	totalCount := u.Edges.totalCount[0]
	if nodes, err := u.Edges.PostsOrErr(); err == nil {
		conn := &PostConnection{Edges: []*PostEdge{}}
		if totalCount != nil {
			conn.TotalCount = *totalCount
		}
		opts = append(opts, WithPostOrder(orderBy))
		pager, err := newPostPager(opts)
		if err != nil {
			return nil, err
		}
		conn.build(nodes, pager, first, last)
		return conn, nil
	}
	query := u.QueryPosts()
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newPostPager(opts)
	if err != nil {
		return nil, err
	}
	if query, err = pager.applyFilter(query); err != nil {
		return nil, err
	}
	conn := &PostConnection{Edges: []*PostEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
			if totalCount != nil {
				conn.TotalCount = *totalCount
			} else if conn.TotalCount, err = query.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := query.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	query = pager.applyCursors(query, after, before)
	query = pager.applyOrder(query, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		query.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := query.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := query.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}
	conn.build(nodes, pager, first, last)
	return conn, nil
}
