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
	"github.com/google/go-github/v50/github"
	"github.com/lrstanley/liam.sh/internal/ent/githubgist"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// GithubGistUpdate is the builder for updating GithubGist entities.
type GithubGistUpdate struct {
	config
	hooks    []Hook
	mutation *GithubGistMutation
}

// Where appends a list predicates to the GithubGistUpdate builder.
func (ggu *GithubGistUpdate) Where(ps ...predicate.GithubGist) *GithubGistUpdate {
	ggu.mutation.Where(ps...)
	return ggu
}

// SetGistID sets the "gist_id" field.
func (ggu *GithubGistUpdate) SetGistID(s string) *GithubGistUpdate {
	ggu.mutation.SetGistID(s)
	return ggu
}

// SetHTMLURL sets the "html_url" field.
func (ggu *GithubGistUpdate) SetHTMLURL(s string) *GithubGistUpdate {
	ggu.mutation.SetHTMLURL(s)
	return ggu
}

// SetPublic sets the "public" field.
func (ggu *GithubGistUpdate) SetPublic(b bool) *GithubGistUpdate {
	ggu.mutation.SetPublic(b)
	return ggu
}

// SetCreatedAt sets the "created_at" field.
func (ggu *GithubGistUpdate) SetCreatedAt(t time.Time) *GithubGistUpdate {
	ggu.mutation.SetCreatedAt(t)
	return ggu
}

// SetUpdatedAt sets the "updated_at" field.
func (ggu *GithubGistUpdate) SetUpdatedAt(t time.Time) *GithubGistUpdate {
	ggu.mutation.SetUpdatedAt(t)
	return ggu
}

// SetDescription sets the "description" field.
func (ggu *GithubGistUpdate) SetDescription(s string) *GithubGistUpdate {
	ggu.mutation.SetDescription(s)
	return ggu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ggu *GithubGistUpdate) SetNillableDescription(s *string) *GithubGistUpdate {
	if s != nil {
		ggu.SetDescription(*s)
	}
	return ggu
}

// ClearDescription clears the value of the "description" field.
func (ggu *GithubGistUpdate) ClearDescription() *GithubGistUpdate {
	ggu.mutation.ClearDescription()
	return ggu
}

// SetOwner sets the "owner" field.
func (ggu *GithubGistUpdate) SetOwner(gi *github.User) *GithubGistUpdate {
	ggu.mutation.SetOwner(gi)
	return ggu
}

// SetName sets the "name" field.
func (ggu *GithubGistUpdate) SetName(s string) *GithubGistUpdate {
	ggu.mutation.SetName(s)
	return ggu
}

// SetType sets the "type" field.
func (ggu *GithubGistUpdate) SetType(s string) *GithubGistUpdate {
	ggu.mutation.SetType(s)
	return ggu
}

// SetLanguage sets the "language" field.
func (ggu *GithubGistUpdate) SetLanguage(s string) *GithubGistUpdate {
	ggu.mutation.SetLanguage(s)
	return ggu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (ggu *GithubGistUpdate) SetNillableLanguage(s *string) *GithubGistUpdate {
	if s != nil {
		ggu.SetLanguage(*s)
	}
	return ggu
}

// ClearLanguage clears the value of the "language" field.
func (ggu *GithubGistUpdate) ClearLanguage() *GithubGistUpdate {
	ggu.mutation.ClearLanguage()
	return ggu
}

// SetSize sets the "size" field.
func (ggu *GithubGistUpdate) SetSize(i int64) *GithubGistUpdate {
	ggu.mutation.ResetSize()
	ggu.mutation.SetSize(i)
	return ggu
}

// AddSize adds i to the "size" field.
func (ggu *GithubGistUpdate) AddSize(i int64) *GithubGistUpdate {
	ggu.mutation.AddSize(i)
	return ggu
}

// SetRawURL sets the "raw_url" field.
func (ggu *GithubGistUpdate) SetRawURL(s string) *GithubGistUpdate {
	ggu.mutation.SetRawURL(s)
	return ggu
}

// SetContent sets the "content" field.
func (ggu *GithubGistUpdate) SetContent(s string) *GithubGistUpdate {
	ggu.mutation.SetContent(s)
	return ggu
}

// Mutation returns the GithubGistMutation object of the builder.
func (ggu *GithubGistUpdate) Mutation() *GithubGistMutation {
	return ggu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ggu *GithubGistUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GithubGistMutation](ctx, ggu.sqlSave, ggu.mutation, ggu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ggu *GithubGistUpdate) SaveX(ctx context.Context) int {
	affected, err := ggu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ggu *GithubGistUpdate) Exec(ctx context.Context) error {
	_, err := ggu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ggu *GithubGistUpdate) ExecX(ctx context.Context) {
	if err := ggu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ggu *GithubGistUpdate) check() error {
	if v, ok := ggu.mutation.HTMLURL(); ok {
		if err := githubgist.HTMLURLValidator(v); err != nil {
			return &ValidationError{Name: "html_url", err: fmt.Errorf(`ent: validator failed for field "GithubGist.html_url": %w`, err)}
		}
	}
	return nil
}

func (ggu *GithubGistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ggu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(githubgist.Table, githubgist.Columns, sqlgraph.NewFieldSpec(githubgist.FieldID, field.TypeInt))
	if ps := ggu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ggu.mutation.GistID(); ok {
		_spec.SetField(githubgist.FieldGistID, field.TypeString, value)
	}
	if value, ok := ggu.mutation.HTMLURL(); ok {
		_spec.SetField(githubgist.FieldHTMLURL, field.TypeString, value)
	}
	if value, ok := ggu.mutation.Public(); ok {
		_spec.SetField(githubgist.FieldPublic, field.TypeBool, value)
	}
	if value, ok := ggu.mutation.CreatedAt(); ok {
		_spec.SetField(githubgist.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ggu.mutation.UpdatedAt(); ok {
		_spec.SetField(githubgist.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ggu.mutation.Description(); ok {
		_spec.SetField(githubgist.FieldDescription, field.TypeString, value)
	}
	if ggu.mutation.DescriptionCleared() {
		_spec.ClearField(githubgist.FieldDescription, field.TypeString)
	}
	if value, ok := ggu.mutation.Owner(); ok {
		_spec.SetField(githubgist.FieldOwner, field.TypeJSON, value)
	}
	if value, ok := ggu.mutation.Name(); ok {
		_spec.SetField(githubgist.FieldName, field.TypeString, value)
	}
	if value, ok := ggu.mutation.GetType(); ok {
		_spec.SetField(githubgist.FieldType, field.TypeString, value)
	}
	if value, ok := ggu.mutation.Language(); ok {
		_spec.SetField(githubgist.FieldLanguage, field.TypeString, value)
	}
	if ggu.mutation.LanguageCleared() {
		_spec.ClearField(githubgist.FieldLanguage, field.TypeString)
	}
	if value, ok := ggu.mutation.Size(); ok {
		_spec.SetField(githubgist.FieldSize, field.TypeInt64, value)
	}
	if value, ok := ggu.mutation.AddedSize(); ok {
		_spec.AddField(githubgist.FieldSize, field.TypeInt64, value)
	}
	if value, ok := ggu.mutation.RawURL(); ok {
		_spec.SetField(githubgist.FieldRawURL, field.TypeString, value)
	}
	if value, ok := ggu.mutation.Content(); ok {
		_spec.SetField(githubgist.FieldContent, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ggu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubgist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ggu.mutation.done = true
	return n, nil
}

// GithubGistUpdateOne is the builder for updating a single GithubGist entity.
type GithubGistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GithubGistMutation
}

// SetGistID sets the "gist_id" field.
func (gguo *GithubGistUpdateOne) SetGistID(s string) *GithubGistUpdateOne {
	gguo.mutation.SetGistID(s)
	return gguo
}

// SetHTMLURL sets the "html_url" field.
func (gguo *GithubGistUpdateOne) SetHTMLURL(s string) *GithubGistUpdateOne {
	gguo.mutation.SetHTMLURL(s)
	return gguo
}

// SetPublic sets the "public" field.
func (gguo *GithubGistUpdateOne) SetPublic(b bool) *GithubGistUpdateOne {
	gguo.mutation.SetPublic(b)
	return gguo
}

// SetCreatedAt sets the "created_at" field.
func (gguo *GithubGistUpdateOne) SetCreatedAt(t time.Time) *GithubGistUpdateOne {
	gguo.mutation.SetCreatedAt(t)
	return gguo
}

// SetUpdatedAt sets the "updated_at" field.
func (gguo *GithubGistUpdateOne) SetUpdatedAt(t time.Time) *GithubGistUpdateOne {
	gguo.mutation.SetUpdatedAt(t)
	return gguo
}

// SetDescription sets the "description" field.
func (gguo *GithubGistUpdateOne) SetDescription(s string) *GithubGistUpdateOne {
	gguo.mutation.SetDescription(s)
	return gguo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gguo *GithubGistUpdateOne) SetNillableDescription(s *string) *GithubGistUpdateOne {
	if s != nil {
		gguo.SetDescription(*s)
	}
	return gguo
}

// ClearDescription clears the value of the "description" field.
func (gguo *GithubGistUpdateOne) ClearDescription() *GithubGistUpdateOne {
	gguo.mutation.ClearDescription()
	return gguo
}

// SetOwner sets the "owner" field.
func (gguo *GithubGistUpdateOne) SetOwner(gi *github.User) *GithubGistUpdateOne {
	gguo.mutation.SetOwner(gi)
	return gguo
}

// SetName sets the "name" field.
func (gguo *GithubGistUpdateOne) SetName(s string) *GithubGistUpdateOne {
	gguo.mutation.SetName(s)
	return gguo
}

// SetType sets the "type" field.
func (gguo *GithubGistUpdateOne) SetType(s string) *GithubGistUpdateOne {
	gguo.mutation.SetType(s)
	return gguo
}

// SetLanguage sets the "language" field.
func (gguo *GithubGistUpdateOne) SetLanguage(s string) *GithubGistUpdateOne {
	gguo.mutation.SetLanguage(s)
	return gguo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (gguo *GithubGistUpdateOne) SetNillableLanguage(s *string) *GithubGistUpdateOne {
	if s != nil {
		gguo.SetLanguage(*s)
	}
	return gguo
}

// ClearLanguage clears the value of the "language" field.
func (gguo *GithubGistUpdateOne) ClearLanguage() *GithubGistUpdateOne {
	gguo.mutation.ClearLanguage()
	return gguo
}

// SetSize sets the "size" field.
func (gguo *GithubGistUpdateOne) SetSize(i int64) *GithubGistUpdateOne {
	gguo.mutation.ResetSize()
	gguo.mutation.SetSize(i)
	return gguo
}

// AddSize adds i to the "size" field.
func (gguo *GithubGistUpdateOne) AddSize(i int64) *GithubGistUpdateOne {
	gguo.mutation.AddSize(i)
	return gguo
}

// SetRawURL sets the "raw_url" field.
func (gguo *GithubGistUpdateOne) SetRawURL(s string) *GithubGistUpdateOne {
	gguo.mutation.SetRawURL(s)
	return gguo
}

// SetContent sets the "content" field.
func (gguo *GithubGistUpdateOne) SetContent(s string) *GithubGistUpdateOne {
	gguo.mutation.SetContent(s)
	return gguo
}

// Mutation returns the GithubGistMutation object of the builder.
func (gguo *GithubGistUpdateOne) Mutation() *GithubGistMutation {
	return gguo.mutation
}

// Where appends a list predicates to the GithubGistUpdate builder.
func (gguo *GithubGistUpdateOne) Where(ps ...predicate.GithubGist) *GithubGistUpdateOne {
	gguo.mutation.Where(ps...)
	return gguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gguo *GithubGistUpdateOne) Select(field string, fields ...string) *GithubGistUpdateOne {
	gguo.fields = append([]string{field}, fields...)
	return gguo
}

// Save executes the query and returns the updated GithubGist entity.
func (gguo *GithubGistUpdateOne) Save(ctx context.Context) (*GithubGist, error) {
	return withHooks[*GithubGist, GithubGistMutation](ctx, gguo.sqlSave, gguo.mutation, gguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gguo *GithubGistUpdateOne) SaveX(ctx context.Context) *GithubGist {
	node, err := gguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gguo *GithubGistUpdateOne) Exec(ctx context.Context) error {
	_, err := gguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gguo *GithubGistUpdateOne) ExecX(ctx context.Context) {
	if err := gguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gguo *GithubGistUpdateOne) check() error {
	if v, ok := gguo.mutation.HTMLURL(); ok {
		if err := githubgist.HTMLURLValidator(v); err != nil {
			return &ValidationError{Name: "html_url", err: fmt.Errorf(`ent: validator failed for field "GithubGist.html_url": %w`, err)}
		}
	}
	return nil
}

func (gguo *GithubGistUpdateOne) sqlSave(ctx context.Context) (_node *GithubGist, err error) {
	if err := gguo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(githubgist.Table, githubgist.Columns, sqlgraph.NewFieldSpec(githubgist.FieldID, field.TypeInt))
	id, ok := gguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GithubGist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githubgist.FieldID)
		for _, f := range fields {
			if !githubgist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != githubgist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gguo.mutation.GistID(); ok {
		_spec.SetField(githubgist.FieldGistID, field.TypeString, value)
	}
	if value, ok := gguo.mutation.HTMLURL(); ok {
		_spec.SetField(githubgist.FieldHTMLURL, field.TypeString, value)
	}
	if value, ok := gguo.mutation.Public(); ok {
		_spec.SetField(githubgist.FieldPublic, field.TypeBool, value)
	}
	if value, ok := gguo.mutation.CreatedAt(); ok {
		_spec.SetField(githubgist.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := gguo.mutation.UpdatedAt(); ok {
		_spec.SetField(githubgist.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gguo.mutation.Description(); ok {
		_spec.SetField(githubgist.FieldDescription, field.TypeString, value)
	}
	if gguo.mutation.DescriptionCleared() {
		_spec.ClearField(githubgist.FieldDescription, field.TypeString)
	}
	if value, ok := gguo.mutation.Owner(); ok {
		_spec.SetField(githubgist.FieldOwner, field.TypeJSON, value)
	}
	if value, ok := gguo.mutation.Name(); ok {
		_spec.SetField(githubgist.FieldName, field.TypeString, value)
	}
	if value, ok := gguo.mutation.GetType(); ok {
		_spec.SetField(githubgist.FieldType, field.TypeString, value)
	}
	if value, ok := gguo.mutation.Language(); ok {
		_spec.SetField(githubgist.FieldLanguage, field.TypeString, value)
	}
	if gguo.mutation.LanguageCleared() {
		_spec.ClearField(githubgist.FieldLanguage, field.TypeString)
	}
	if value, ok := gguo.mutation.Size(); ok {
		_spec.SetField(githubgist.FieldSize, field.TypeInt64, value)
	}
	if value, ok := gguo.mutation.AddedSize(); ok {
		_spec.AddField(githubgist.FieldSize, field.TypeInt64, value)
	}
	if value, ok := gguo.mutation.RawURL(); ok {
		_spec.SetField(githubgist.FieldRawURL, field.TypeString, value)
	}
	if value, ok := gguo.mutation.Content(); ok {
		_spec.SetField(githubgist.FieldContent, field.TypeString, value)
	}
	_node = &GithubGist{config: gguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubgist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gguo.mutation.done = true
	return _node, nil
}
