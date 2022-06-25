// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/liam.sh/internal/ent/githubgist"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// GithubGistDelete is the builder for deleting a GithubGist entity.
type GithubGistDelete struct {
	config
	hooks    []Hook
	mutation *GithubGistMutation
}

// Where appends a list predicates to the GithubGistDelete builder.
func (ggd *GithubGistDelete) Where(ps ...predicate.GithubGist) *GithubGistDelete {
	ggd.mutation.Where(ps...)
	return ggd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ggd *GithubGistDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ggd.hooks) == 0 {
		affected, err = ggd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubGistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ggd.mutation = mutation
			affected, err = ggd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ggd.hooks) - 1; i >= 0; i-- {
			if ggd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ggd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ggd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ggd *GithubGistDelete) ExecX(ctx context.Context) int {
	n, err := ggd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ggd *GithubGistDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: githubgist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githubgist.FieldID,
			},
		},
	}
	if ps := ggd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ggd.driver, _spec)
}

// GithubGistDeleteOne is the builder for deleting a single GithubGist entity.
type GithubGistDeleteOne struct {
	ggd *GithubGistDelete
}

// Exec executes the deletion query.
func (ggdo *GithubGistDeleteOne) Exec(ctx context.Context) error {
	n, err := ggdo.ggd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{githubgist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ggdo *GithubGistDeleteOne) ExecX(ctx context.Context) {
	ggdo.ggd.ExecX(ctx)
}
