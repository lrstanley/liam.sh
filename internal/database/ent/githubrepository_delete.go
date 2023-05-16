// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/database/ent/predicate"
)

// GithubRepositoryDelete is the builder for deleting a GithubRepository entity.
type GithubRepositoryDelete struct {
	config
	hooks    []Hook
	mutation *GithubRepositoryMutation
}

// Where appends a list predicates to the GithubRepositoryDelete builder.
func (grd *GithubRepositoryDelete) Where(ps ...predicate.GithubRepository) *GithubRepositoryDelete {
	grd.mutation.Where(ps...)
	return grd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (grd *GithubRepositoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, grd.sqlExec, grd.mutation, grd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (grd *GithubRepositoryDelete) ExecX(ctx context.Context) int {
	n, err := grd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (grd *GithubRepositoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(githubrepository.Table, sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt))
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
	grd.mutation.done = true
	return affected, err
}

// GithubRepositoryDeleteOne is the builder for deleting a single GithubRepository entity.
type GithubRepositoryDeleteOne struct {
	grd *GithubRepositoryDelete
}

// Where appends a list predicates to the GithubRepositoryDelete builder.
func (grdo *GithubRepositoryDeleteOne) Where(ps ...predicate.GithubRepository) *GithubRepositoryDeleteOne {
	grdo.grd.mutation.Where(ps...)
	return grdo
}

// Exec executes the deletion query.
func (grdo *GithubRepositoryDeleteOne) Exec(ctx context.Context) error {
	n, err := grdo.grd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{githubrepository.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (grdo *GithubRepositoryDeleteOne) ExecX(ctx context.Context) {
	if err := grdo.Exec(ctx); err != nil {
		panic(err)
	}
}
