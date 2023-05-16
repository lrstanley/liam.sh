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
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/database/ent/label"
	"github.com/lrstanley/liam.sh/internal/database/ent/post"
	"github.com/lrstanley/liam.sh/internal/database/ent/predicate"
)

// LabelUpdate is the builder for updating Label entities.
type LabelUpdate struct {
	config
	hooks    []Hook
	mutation *LabelMutation
}

// Where appends a list predicates to the LabelUpdate builder.
func (lu *LabelUpdate) Where(ps ...predicate.Label) *LabelUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUpdateTime sets the "update_time" field.
func (lu *LabelUpdate) SetUpdateTime(t time.Time) *LabelUpdate {
	lu.mutation.SetUpdateTime(t)
	return lu
}

// SetName sets the "name" field.
func (lu *LabelUpdate) SetName(s string) *LabelUpdate {
	lu.mutation.SetName(s)
	return lu
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (lu *LabelUpdate) AddPostIDs(ids ...int) *LabelUpdate {
	lu.mutation.AddPostIDs(ids...)
	return lu
}

// AddPosts adds the "posts" edges to the Post entity.
func (lu *LabelUpdate) AddPosts(p ...*Post) *LabelUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lu.AddPostIDs(ids...)
}

// AddGithubRepositoryIDs adds the "github_repositories" edge to the GithubRepository entity by IDs.
func (lu *LabelUpdate) AddGithubRepositoryIDs(ids ...int) *LabelUpdate {
	lu.mutation.AddGithubRepositoryIDs(ids...)
	return lu
}

// AddGithubRepositories adds the "github_repositories" edges to the GithubRepository entity.
func (lu *LabelUpdate) AddGithubRepositories(g ...*GithubRepository) *LabelUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return lu.AddGithubRepositoryIDs(ids...)
}

// Mutation returns the LabelMutation object of the builder.
func (lu *LabelUpdate) Mutation() *LabelMutation {
	return lu.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (lu *LabelUpdate) ClearPosts() *LabelUpdate {
	lu.mutation.ClearPosts()
	return lu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (lu *LabelUpdate) RemovePostIDs(ids ...int) *LabelUpdate {
	lu.mutation.RemovePostIDs(ids...)
	return lu
}

// RemovePosts removes "posts" edges to Post entities.
func (lu *LabelUpdate) RemovePosts(p ...*Post) *LabelUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lu.RemovePostIDs(ids...)
}

// ClearGithubRepositories clears all "github_repositories" edges to the GithubRepository entity.
func (lu *LabelUpdate) ClearGithubRepositories() *LabelUpdate {
	lu.mutation.ClearGithubRepositories()
	return lu
}

// RemoveGithubRepositoryIDs removes the "github_repositories" edge to GithubRepository entities by IDs.
func (lu *LabelUpdate) RemoveGithubRepositoryIDs(ids ...int) *LabelUpdate {
	lu.mutation.RemoveGithubRepositoryIDs(ids...)
	return lu
}

// RemoveGithubRepositories removes "github_repositories" edges to GithubRepository entities.
func (lu *LabelUpdate) RemoveGithubRepositories(g ...*GithubRepository) *LabelUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return lu.RemoveGithubRepositoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LabelUpdate) Save(ctx context.Context) (int, error) {
	if err := lu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LabelUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LabelUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LabelUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LabelUpdate) defaults() error {
	if _, ok := lu.mutation.UpdateTime(); !ok {
		if label.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized label.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := label.UpdateDefaultUpdateTime()
		lu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lu *LabelUpdate) check() error {
	if v, ok := lu.mutation.Name(); ok {
		if err := label.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Label.name": %w`, err)}
		}
	}
	return nil
}

func (lu *LabelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(label.Table, label.Columns, sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UpdateTime(); ok {
		_spec.SetField(label.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(label.FieldName, field.TypeString, value)
	}
	if lu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.PostsTable,
			Columns: label.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !lu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.PostsTable,
			Columns: label.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.PostsTable,
			Columns: label.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.GithubRepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.GithubRepositoriesTable,
			Columns: label.GithubRepositoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedGithubRepositoriesIDs(); len(nodes) > 0 && !lu.mutation.GithubRepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.GithubRepositoriesTable,
			Columns: label.GithubRepositoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.GithubRepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.GithubRepositoriesTable,
			Columns: label.GithubRepositoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{label.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LabelUpdateOne is the builder for updating a single Label entity.
type LabelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LabelMutation
}

// SetUpdateTime sets the "update_time" field.
func (luo *LabelUpdateOne) SetUpdateTime(t time.Time) *LabelUpdateOne {
	luo.mutation.SetUpdateTime(t)
	return luo
}

// SetName sets the "name" field.
func (luo *LabelUpdateOne) SetName(s string) *LabelUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (luo *LabelUpdateOne) AddPostIDs(ids ...int) *LabelUpdateOne {
	luo.mutation.AddPostIDs(ids...)
	return luo
}

// AddPosts adds the "posts" edges to the Post entity.
func (luo *LabelUpdateOne) AddPosts(p ...*Post) *LabelUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return luo.AddPostIDs(ids...)
}

// AddGithubRepositoryIDs adds the "github_repositories" edge to the GithubRepository entity by IDs.
func (luo *LabelUpdateOne) AddGithubRepositoryIDs(ids ...int) *LabelUpdateOne {
	luo.mutation.AddGithubRepositoryIDs(ids...)
	return luo
}

// AddGithubRepositories adds the "github_repositories" edges to the GithubRepository entity.
func (luo *LabelUpdateOne) AddGithubRepositories(g ...*GithubRepository) *LabelUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return luo.AddGithubRepositoryIDs(ids...)
}

// Mutation returns the LabelMutation object of the builder.
func (luo *LabelUpdateOne) Mutation() *LabelMutation {
	return luo.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (luo *LabelUpdateOne) ClearPosts() *LabelUpdateOne {
	luo.mutation.ClearPosts()
	return luo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (luo *LabelUpdateOne) RemovePostIDs(ids ...int) *LabelUpdateOne {
	luo.mutation.RemovePostIDs(ids...)
	return luo
}

// RemovePosts removes "posts" edges to Post entities.
func (luo *LabelUpdateOne) RemovePosts(p ...*Post) *LabelUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return luo.RemovePostIDs(ids...)
}

// ClearGithubRepositories clears all "github_repositories" edges to the GithubRepository entity.
func (luo *LabelUpdateOne) ClearGithubRepositories() *LabelUpdateOne {
	luo.mutation.ClearGithubRepositories()
	return luo
}

// RemoveGithubRepositoryIDs removes the "github_repositories" edge to GithubRepository entities by IDs.
func (luo *LabelUpdateOne) RemoveGithubRepositoryIDs(ids ...int) *LabelUpdateOne {
	luo.mutation.RemoveGithubRepositoryIDs(ids...)
	return luo
}

// RemoveGithubRepositories removes "github_repositories" edges to GithubRepository entities.
func (luo *LabelUpdateOne) RemoveGithubRepositories(g ...*GithubRepository) *LabelUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return luo.RemoveGithubRepositoryIDs(ids...)
}

// Where appends a list predicates to the LabelUpdate builder.
func (luo *LabelUpdateOne) Where(ps ...predicate.Label) *LabelUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LabelUpdateOne) Select(field string, fields ...string) *LabelUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Label entity.
func (luo *LabelUpdateOne) Save(ctx context.Context) (*Label, error) {
	if err := luo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LabelUpdateOne) SaveX(ctx context.Context) *Label {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LabelUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LabelUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LabelUpdateOne) defaults() error {
	if _, ok := luo.mutation.UpdateTime(); !ok {
		if label.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized label.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := label.UpdateDefaultUpdateTime()
		luo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (luo *LabelUpdateOne) check() error {
	if v, ok := luo.mutation.Name(); ok {
		if err := label.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Label.name": %w`, err)}
		}
	}
	return nil
}

func (luo *LabelUpdateOne) sqlSave(ctx context.Context) (_node *Label, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(label.Table, label.Columns, sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Label.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, label.FieldID)
		for _, f := range fields {
			if !label.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != label.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.UpdateTime(); ok {
		_spec.SetField(label.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(label.FieldName, field.TypeString, value)
	}
	if luo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.PostsTable,
			Columns: label.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !luo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.PostsTable,
			Columns: label.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.PostsTable,
			Columns: label.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.GithubRepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.GithubRepositoriesTable,
			Columns: label.GithubRepositoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedGithubRepositoriesIDs(); len(nodes) > 0 && !luo.mutation.GithubRepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.GithubRepositoriesTable,
			Columns: label.GithubRepositoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.GithubRepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.GithubRepositoriesTable,
			Columns: label.GithubRepositoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Label{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{label.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
