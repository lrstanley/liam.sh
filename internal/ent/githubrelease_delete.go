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
	"github.com/lrstanley/liam.sh/internal/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// GithubReleaseDelete is the builder for deleting a GithubRelease entity.
type GithubReleaseDelete struct {
	config
	hooks    []Hook
	mutation *GithubReleaseMutation
}

// Where appends a list predicates to the GithubReleaseDelete builder.
func (grd *GithubReleaseDelete) Where(ps ...predicate.GithubRelease) *GithubReleaseDelete {
	grd.mutation.Where(ps...)
	return grd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (grd *GithubReleaseDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(grd.hooks) == 0 {
		affected, err = grd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubReleaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			grd.mutation = mutation
			affected, err = grd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(grd.hooks) - 1; i >= 0; i-- {
			if grd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = grd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, grd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (grd *GithubReleaseDelete) ExecX(ctx context.Context) int {
	n, err := grd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (grd *GithubReleaseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: githubrelease.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githubrelease.FieldID,
			},
		},
	}
	if ps := grd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, grd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// GithubReleaseDeleteOne is the builder for deleting a single GithubRelease entity.
type GithubReleaseDeleteOne struct {
	grd *GithubReleaseDelete
}

// Exec executes the deletion query.
func (grdo *GithubReleaseDeleteOne) Exec(ctx context.Context) error {
	n, err := grdo.grd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{githubrelease.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (grdo *GithubReleaseDeleteOne) ExecX(ctx context.Context) {
	grdo.grd.ExecX(ctx)
}
