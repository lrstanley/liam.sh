// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/go-github/v44/github"
	"github.com/lrstanley/liam.sh/internal/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// GithubReleaseUpdate is the builder for updating GithubRelease entities.
type GithubReleaseUpdate struct {
	config
	hooks    []Hook
	mutation *GithubReleaseMutation
}

// Where appends a list predicates to the GithubReleaseUpdate builder.
func (gru *GithubReleaseUpdate) Where(ps ...predicate.GithubRelease) *GithubReleaseUpdate {
	gru.mutation.Where(ps...)
	return gru
}

// SetReleaseID sets the "release_id" field.
func (gru *GithubReleaseUpdate) SetReleaseID(i int64) *GithubReleaseUpdate {
	gru.mutation.ResetReleaseID()
	gru.mutation.SetReleaseID(i)
	return gru
}

// AddReleaseID adds i to the "release_id" field.
func (gru *GithubReleaseUpdate) AddReleaseID(i int64) *GithubReleaseUpdate {
	gru.mutation.AddReleaseID(i)
	return gru
}

// SetHTMLURL sets the "html_url" field.
func (gru *GithubReleaseUpdate) SetHTMLURL(s string) *GithubReleaseUpdate {
	gru.mutation.SetHTMLURL(s)
	return gru
}

// SetTagName sets the "tag_name" field.
func (gru *GithubReleaseUpdate) SetTagName(s string) *GithubReleaseUpdate {
	gru.mutation.SetTagName(s)
	return gru
}

// SetTargetCommitish sets the "target_commitish" field.
func (gru *GithubReleaseUpdate) SetTargetCommitish(s string) *GithubReleaseUpdate {
	gru.mutation.SetTargetCommitish(s)
	return gru
}

// SetName sets the "name" field.
func (gru *GithubReleaseUpdate) SetName(s string) *GithubReleaseUpdate {
	gru.mutation.SetName(s)
	return gru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (gru *GithubReleaseUpdate) SetNillableName(s *string) *GithubReleaseUpdate {
	if s != nil {
		gru.SetName(*s)
	}
	return gru
}

// ClearName clears the value of the "name" field.
func (gru *GithubReleaseUpdate) ClearName() *GithubReleaseUpdate {
	gru.mutation.ClearName()
	return gru
}

// SetDraft sets the "draft" field.
func (gru *GithubReleaseUpdate) SetDraft(b bool) *GithubReleaseUpdate {
	gru.mutation.SetDraft(b)
	return gru
}

// SetPrerelease sets the "prerelease" field.
func (gru *GithubReleaseUpdate) SetPrerelease(b bool) *GithubReleaseUpdate {
	gru.mutation.SetPrerelease(b)
	return gru
}

// SetCreatedAt sets the "created_at" field.
func (gru *GithubReleaseUpdate) SetCreatedAt(t time.Time) *GithubReleaseUpdate {
	gru.mutation.SetCreatedAt(t)
	return gru
}

// SetPublishedAt sets the "published_at" field.
func (gru *GithubReleaseUpdate) SetPublishedAt(t time.Time) *GithubReleaseUpdate {
	gru.mutation.SetPublishedAt(t)
	return gru
}

// SetAuthor sets the "author" field.
func (gru *GithubReleaseUpdate) SetAuthor(gi *github.User) *GithubReleaseUpdate {
	gru.mutation.SetAuthor(gi)
	return gru
}

// SetRepositoryID sets the "repository" edge to the GithubRepository entity by ID.
func (gru *GithubReleaseUpdate) SetRepositoryID(id int) *GithubReleaseUpdate {
	gru.mutation.SetRepositoryID(id)
	return gru
}

// SetRepository sets the "repository" edge to the GithubRepository entity.
func (gru *GithubReleaseUpdate) SetRepository(g *GithubRepository) *GithubReleaseUpdate {
	return gru.SetRepositoryID(g.ID)
}

// AddAssetIDs adds the "assets" edge to the GithubAsset entity by IDs.
func (gru *GithubReleaseUpdate) AddAssetIDs(ids ...int) *GithubReleaseUpdate {
	gru.mutation.AddAssetIDs(ids...)
	return gru
}

// AddAssets adds the "assets" edges to the GithubAsset entity.
func (gru *GithubReleaseUpdate) AddAssets(g ...*GithubAsset) *GithubReleaseUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gru.AddAssetIDs(ids...)
}

// Mutation returns the GithubReleaseMutation object of the builder.
func (gru *GithubReleaseUpdate) Mutation() *GithubReleaseMutation {
	return gru.mutation
}

// ClearRepository clears the "repository" edge to the GithubRepository entity.
func (gru *GithubReleaseUpdate) ClearRepository() *GithubReleaseUpdate {
	gru.mutation.ClearRepository()
	return gru
}

// ClearAssets clears all "assets" edges to the GithubAsset entity.
func (gru *GithubReleaseUpdate) ClearAssets() *GithubReleaseUpdate {
	gru.mutation.ClearAssets()
	return gru
}

// RemoveAssetIDs removes the "assets" edge to GithubAsset entities by IDs.
func (gru *GithubReleaseUpdate) RemoveAssetIDs(ids ...int) *GithubReleaseUpdate {
	gru.mutation.RemoveAssetIDs(ids...)
	return gru
}

// RemoveAssets removes "assets" edges to GithubAsset entities.
func (gru *GithubReleaseUpdate) RemoveAssets(g ...*GithubAsset) *GithubReleaseUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gru.RemoveAssetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gru *GithubReleaseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gru.hooks) == 0 {
		if err = gru.check(); err != nil {
			return 0, err
		}
		affected, err = gru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubReleaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gru.check(); err != nil {
				return 0, err
			}
			gru.mutation = mutation
			affected, err = gru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gru.hooks) - 1; i >= 0; i-- {
			if gru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gru *GithubReleaseUpdate) SaveX(ctx context.Context) int {
	affected, err := gru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gru *GithubReleaseUpdate) Exec(ctx context.Context) error {
	_, err := gru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gru *GithubReleaseUpdate) ExecX(ctx context.Context) {
	if err := gru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gru *GithubReleaseUpdate) check() error {
	if v, ok := gru.mutation.ReleaseID(); ok {
		if err := githubrelease.ReleaseIDValidator(v); err != nil {
			return &ValidationError{Name: "release_id", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.release_id": %w`, err)}
		}
	}
	if v, ok := gru.mutation.HTMLURL(); ok {
		if err := githubrelease.HTMLURLValidator(v); err != nil {
			return &ValidationError{Name: "html_url", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.html_url": %w`, err)}
		}
	}
	if v, ok := gru.mutation.TagName(); ok {
		if err := githubrelease.TagNameValidator(v); err != nil {
			return &ValidationError{Name: "tag_name", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.tag_name": %w`, err)}
		}
	}
	if v, ok := gru.mutation.TargetCommitish(); ok {
		if err := githubrelease.TargetCommitishValidator(v); err != nil {
			return &ValidationError{Name: "target_commitish", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.target_commitish": %w`, err)}
		}
	}
	if _, ok := gru.mutation.RepositoryID(); gru.mutation.RepositoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GithubRelease.repository"`)
	}
	return nil
}

func (gru *GithubReleaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   githubrelease.Table,
			Columns: githubrelease.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githubrelease.FieldID,
			},
		},
	}
	if ps := gru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gru.mutation.ReleaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: githubrelease.FieldReleaseID,
		})
	}
	if value, ok := gru.mutation.AddedReleaseID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: githubrelease.FieldReleaseID,
		})
	}
	if value, ok := gru.mutation.HTMLURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: githubrelease.FieldHTMLURL,
		})
	}
	if value, ok := gru.mutation.TagName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: githubrelease.FieldTagName,
		})
	}
	if value, ok := gru.mutation.TargetCommitish(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: githubrelease.FieldTargetCommitish,
		})
	}
	if value, ok := gru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: githubrelease.FieldName,
		})
	}
	if gru.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: githubrelease.FieldName,
		})
	}
	if value, ok := gru.mutation.Draft(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: githubrelease.FieldDraft,
		})
	}
	if value, ok := gru.mutation.Prerelease(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: githubrelease.FieldPrerelease,
		})
	}
	if value, ok := gru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: githubrelease.FieldCreatedAt,
		})
	}
	if value, ok := gru.mutation.PublishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: githubrelease.FieldPublishedAt,
		})
	}
	if value, ok := gru.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: githubrelease.FieldAuthor,
		})
	}
	if gru.mutation.RepositoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   githubrelease.RepositoryTable,
			Columns: []string{githubrelease.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubrepository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   githubrelease.RepositoryTable,
			Columns: []string{githubrelease.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubrepository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gru.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrelease.AssetsTable,
			Columns: []string{githubrelease.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubasset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.RemovedAssetsIDs(); len(nodes) > 0 && !gru.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrelease.AssetsTable,
			Columns: []string{githubrelease.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubasset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrelease.AssetsTable,
			Columns: []string{githubrelease.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubasset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubrelease.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GithubReleaseUpdateOne is the builder for updating a single GithubRelease entity.
type GithubReleaseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GithubReleaseMutation
}

// SetReleaseID sets the "release_id" field.
func (gruo *GithubReleaseUpdateOne) SetReleaseID(i int64) *GithubReleaseUpdateOne {
	gruo.mutation.ResetReleaseID()
	gruo.mutation.SetReleaseID(i)
	return gruo
}

// AddReleaseID adds i to the "release_id" field.
func (gruo *GithubReleaseUpdateOne) AddReleaseID(i int64) *GithubReleaseUpdateOne {
	gruo.mutation.AddReleaseID(i)
	return gruo
}

// SetHTMLURL sets the "html_url" field.
func (gruo *GithubReleaseUpdateOne) SetHTMLURL(s string) *GithubReleaseUpdateOne {
	gruo.mutation.SetHTMLURL(s)
	return gruo
}

// SetTagName sets the "tag_name" field.
func (gruo *GithubReleaseUpdateOne) SetTagName(s string) *GithubReleaseUpdateOne {
	gruo.mutation.SetTagName(s)
	return gruo
}

// SetTargetCommitish sets the "target_commitish" field.
func (gruo *GithubReleaseUpdateOne) SetTargetCommitish(s string) *GithubReleaseUpdateOne {
	gruo.mutation.SetTargetCommitish(s)
	return gruo
}

// SetName sets the "name" field.
func (gruo *GithubReleaseUpdateOne) SetName(s string) *GithubReleaseUpdateOne {
	gruo.mutation.SetName(s)
	return gruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (gruo *GithubReleaseUpdateOne) SetNillableName(s *string) *GithubReleaseUpdateOne {
	if s != nil {
		gruo.SetName(*s)
	}
	return gruo
}

// ClearName clears the value of the "name" field.
func (gruo *GithubReleaseUpdateOne) ClearName() *GithubReleaseUpdateOne {
	gruo.mutation.ClearName()
	return gruo
}

// SetDraft sets the "draft" field.
func (gruo *GithubReleaseUpdateOne) SetDraft(b bool) *GithubReleaseUpdateOne {
	gruo.mutation.SetDraft(b)
	return gruo
}

// SetPrerelease sets the "prerelease" field.
func (gruo *GithubReleaseUpdateOne) SetPrerelease(b bool) *GithubReleaseUpdateOne {
	gruo.mutation.SetPrerelease(b)
	return gruo
}

// SetCreatedAt sets the "created_at" field.
func (gruo *GithubReleaseUpdateOne) SetCreatedAt(t time.Time) *GithubReleaseUpdateOne {
	gruo.mutation.SetCreatedAt(t)
	return gruo
}

// SetPublishedAt sets the "published_at" field.
func (gruo *GithubReleaseUpdateOne) SetPublishedAt(t time.Time) *GithubReleaseUpdateOne {
	gruo.mutation.SetPublishedAt(t)
	return gruo
}

// SetAuthor sets the "author" field.
func (gruo *GithubReleaseUpdateOne) SetAuthor(gi *github.User) *GithubReleaseUpdateOne {
	gruo.mutation.SetAuthor(gi)
	return gruo
}

// SetRepositoryID sets the "repository" edge to the GithubRepository entity by ID.
func (gruo *GithubReleaseUpdateOne) SetRepositoryID(id int) *GithubReleaseUpdateOne {
	gruo.mutation.SetRepositoryID(id)
	return gruo
}

// SetRepository sets the "repository" edge to the GithubRepository entity.
func (gruo *GithubReleaseUpdateOne) SetRepository(g *GithubRepository) *GithubReleaseUpdateOne {
	return gruo.SetRepositoryID(g.ID)
}

// AddAssetIDs adds the "assets" edge to the GithubAsset entity by IDs.
func (gruo *GithubReleaseUpdateOne) AddAssetIDs(ids ...int) *GithubReleaseUpdateOne {
	gruo.mutation.AddAssetIDs(ids...)
	return gruo
}

// AddAssets adds the "assets" edges to the GithubAsset entity.
func (gruo *GithubReleaseUpdateOne) AddAssets(g ...*GithubAsset) *GithubReleaseUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gruo.AddAssetIDs(ids...)
}

// Mutation returns the GithubReleaseMutation object of the builder.
func (gruo *GithubReleaseUpdateOne) Mutation() *GithubReleaseMutation {
	return gruo.mutation
}

// ClearRepository clears the "repository" edge to the GithubRepository entity.
func (gruo *GithubReleaseUpdateOne) ClearRepository() *GithubReleaseUpdateOne {
	gruo.mutation.ClearRepository()
	return gruo
}

// ClearAssets clears all "assets" edges to the GithubAsset entity.
func (gruo *GithubReleaseUpdateOne) ClearAssets() *GithubReleaseUpdateOne {
	gruo.mutation.ClearAssets()
	return gruo
}

// RemoveAssetIDs removes the "assets" edge to GithubAsset entities by IDs.
func (gruo *GithubReleaseUpdateOne) RemoveAssetIDs(ids ...int) *GithubReleaseUpdateOne {
	gruo.mutation.RemoveAssetIDs(ids...)
	return gruo
}

// RemoveAssets removes "assets" edges to GithubAsset entities.
func (gruo *GithubReleaseUpdateOne) RemoveAssets(g ...*GithubAsset) *GithubReleaseUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gruo.RemoveAssetIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gruo *GithubReleaseUpdateOne) Select(field string, fields ...string) *GithubReleaseUpdateOne {
	gruo.fields = append([]string{field}, fields...)
	return gruo
}

// Save executes the query and returns the updated GithubRelease entity.
func (gruo *GithubReleaseUpdateOne) Save(ctx context.Context) (*GithubRelease, error) {
	var (
		err  error
		node *GithubRelease
	)
	if len(gruo.hooks) == 0 {
		if err = gruo.check(); err != nil {
			return nil, err
		}
		node, err = gruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubReleaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gruo.check(); err != nil {
				return nil, err
			}
			gruo.mutation = mutation
			node, err = gruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gruo.hooks) - 1; i >= 0; i-- {
			if gruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GithubRelease)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GithubReleaseMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gruo *GithubReleaseUpdateOne) SaveX(ctx context.Context) *GithubRelease {
	node, err := gruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gruo *GithubReleaseUpdateOne) Exec(ctx context.Context) error {
	_, err := gruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gruo *GithubReleaseUpdateOne) ExecX(ctx context.Context) {
	if err := gruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gruo *GithubReleaseUpdateOne) check() error {
	if v, ok := gruo.mutation.ReleaseID(); ok {
		if err := githubrelease.ReleaseIDValidator(v); err != nil {
			return &ValidationError{Name: "release_id", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.release_id": %w`, err)}
		}
	}
	if v, ok := gruo.mutation.HTMLURL(); ok {
		if err := githubrelease.HTMLURLValidator(v); err != nil {
			return &ValidationError{Name: "html_url", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.html_url": %w`, err)}
		}
	}
	if v, ok := gruo.mutation.TagName(); ok {
		if err := githubrelease.TagNameValidator(v); err != nil {
			return &ValidationError{Name: "tag_name", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.tag_name": %w`, err)}
		}
	}
	if v, ok := gruo.mutation.TargetCommitish(); ok {
		if err := githubrelease.TargetCommitishValidator(v); err != nil {
			return &ValidationError{Name: "target_commitish", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.target_commitish": %w`, err)}
		}
	}
	if _, ok := gruo.mutation.RepositoryID(); gruo.mutation.RepositoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GithubRelease.repository"`)
	}
	return nil
}

func (gruo *GithubReleaseUpdateOne) sqlSave(ctx context.Context) (_node *GithubRelease, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   githubrelease.Table,
			Columns: githubrelease.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githubrelease.FieldID,
			},
		},
	}
	id, ok := gruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GithubRelease.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githubrelease.FieldID)
		for _, f := range fields {
			if !githubrelease.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != githubrelease.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gruo.mutation.ReleaseID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: githubrelease.FieldReleaseID,
		})
	}
	if value, ok := gruo.mutation.AddedReleaseID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: githubrelease.FieldReleaseID,
		})
	}
	if value, ok := gruo.mutation.HTMLURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: githubrelease.FieldHTMLURL,
		})
	}
	if value, ok := gruo.mutation.TagName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: githubrelease.FieldTagName,
		})
	}
	if value, ok := gruo.mutation.TargetCommitish(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: githubrelease.FieldTargetCommitish,
		})
	}
	if value, ok := gruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: githubrelease.FieldName,
		})
	}
	if gruo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: githubrelease.FieldName,
		})
	}
	if value, ok := gruo.mutation.Draft(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: githubrelease.FieldDraft,
		})
	}
	if value, ok := gruo.mutation.Prerelease(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: githubrelease.FieldPrerelease,
		})
	}
	if value, ok := gruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: githubrelease.FieldCreatedAt,
		})
	}
	if value, ok := gruo.mutation.PublishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: githubrelease.FieldPublishedAt,
		})
	}
	if value, ok := gruo.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: githubrelease.FieldAuthor,
		})
	}
	if gruo.mutation.RepositoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   githubrelease.RepositoryTable,
			Columns: []string{githubrelease.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubrepository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   githubrelease.RepositoryTable,
			Columns: []string{githubrelease.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubrepository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gruo.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrelease.AssetsTable,
			Columns: []string{githubrelease.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubasset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.RemovedAssetsIDs(); len(nodes) > 0 && !gruo.mutation.AssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrelease.AssetsTable,
			Columns: []string{githubrelease.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubasset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrelease.AssetsTable,
			Columns: []string{githubrelease.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubasset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GithubRelease{config: gruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubrelease.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
