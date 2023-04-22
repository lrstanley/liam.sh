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
	"github.com/google/go-github/v52/github"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubgist"
)

// GithubGistCreate is the builder for creating a GithubGist entity.
type GithubGistCreate struct {
	config
	mutation *GithubGistMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetGistID sets the "gist_id" field.
func (ggc *GithubGistCreate) SetGistID(s string) *GithubGistCreate {
	ggc.mutation.SetGistID(s)
	return ggc
}

// SetHTMLURL sets the "html_url" field.
func (ggc *GithubGistCreate) SetHTMLURL(s string) *GithubGistCreate {
	ggc.mutation.SetHTMLURL(s)
	return ggc
}

// SetPublic sets the "public" field.
func (ggc *GithubGistCreate) SetPublic(b bool) *GithubGistCreate {
	ggc.mutation.SetPublic(b)
	return ggc
}

// SetCreatedAt sets the "created_at" field.
func (ggc *GithubGistCreate) SetCreatedAt(t time.Time) *GithubGistCreate {
	ggc.mutation.SetCreatedAt(t)
	return ggc
}

// SetUpdatedAt sets the "updated_at" field.
func (ggc *GithubGistCreate) SetUpdatedAt(t time.Time) *GithubGistCreate {
	ggc.mutation.SetUpdatedAt(t)
	return ggc
}

// SetDescription sets the "description" field.
func (ggc *GithubGistCreate) SetDescription(s string) *GithubGistCreate {
	ggc.mutation.SetDescription(s)
	return ggc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ggc *GithubGistCreate) SetNillableDescription(s *string) *GithubGistCreate {
	if s != nil {
		ggc.SetDescription(*s)
	}
	return ggc
}

// SetOwner sets the "owner" field.
func (ggc *GithubGistCreate) SetOwner(gi *github.User) *GithubGistCreate {
	ggc.mutation.SetOwner(gi)
	return ggc
}

// SetName sets the "name" field.
func (ggc *GithubGistCreate) SetName(s string) *GithubGistCreate {
	ggc.mutation.SetName(s)
	return ggc
}

// SetType sets the "type" field.
func (ggc *GithubGistCreate) SetType(s string) *GithubGistCreate {
	ggc.mutation.SetType(s)
	return ggc
}

// SetLanguage sets the "language" field.
func (ggc *GithubGistCreate) SetLanguage(s string) *GithubGistCreate {
	ggc.mutation.SetLanguage(s)
	return ggc
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (ggc *GithubGistCreate) SetNillableLanguage(s *string) *GithubGistCreate {
	if s != nil {
		ggc.SetLanguage(*s)
	}
	return ggc
}

// SetSize sets the "size" field.
func (ggc *GithubGistCreate) SetSize(i int64) *GithubGistCreate {
	ggc.mutation.SetSize(i)
	return ggc
}

// SetRawURL sets the "raw_url" field.
func (ggc *GithubGistCreate) SetRawURL(s string) *GithubGistCreate {
	ggc.mutation.SetRawURL(s)
	return ggc
}

// SetContent sets the "content" field.
func (ggc *GithubGistCreate) SetContent(s string) *GithubGistCreate {
	ggc.mutation.SetContent(s)
	return ggc
}

// Mutation returns the GithubGistMutation object of the builder.
func (ggc *GithubGistCreate) Mutation() *GithubGistMutation {
	return ggc.mutation
}

// Save creates the GithubGist in the database.
func (ggc *GithubGistCreate) Save(ctx context.Context) (*GithubGist, error) {
	return withHooks[*GithubGist, GithubGistMutation](ctx, ggc.sqlSave, ggc.mutation, ggc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ggc *GithubGistCreate) SaveX(ctx context.Context) *GithubGist {
	v, err := ggc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ggc *GithubGistCreate) Exec(ctx context.Context) error {
	_, err := ggc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ggc *GithubGistCreate) ExecX(ctx context.Context) {
	if err := ggc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ggc *GithubGistCreate) check() error {
	if _, ok := ggc.mutation.GistID(); !ok {
		return &ValidationError{Name: "gist_id", err: errors.New(`ent: missing required field "GithubGist.gist_id"`)}
	}
	if _, ok := ggc.mutation.HTMLURL(); !ok {
		return &ValidationError{Name: "html_url", err: errors.New(`ent: missing required field "GithubGist.html_url"`)}
	}
	if v, ok := ggc.mutation.HTMLURL(); ok {
		if err := githubgist.HTMLURLValidator(v); err != nil {
			return &ValidationError{Name: "html_url", err: fmt.Errorf(`ent: validator failed for field "GithubGist.html_url": %w`, err)}
		}
	}
	if _, ok := ggc.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`ent: missing required field "GithubGist.public"`)}
	}
	if _, ok := ggc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GithubGist.created_at"`)}
	}
	if _, ok := ggc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "GithubGist.updated_at"`)}
	}
	if _, ok := ggc.mutation.Owner(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required field "GithubGist.owner"`)}
	}
	if _, ok := ggc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "GithubGist.name"`)}
	}
	if _, ok := ggc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "GithubGist.type"`)}
	}
	if _, ok := ggc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "GithubGist.size"`)}
	}
	if _, ok := ggc.mutation.RawURL(); !ok {
		return &ValidationError{Name: "raw_url", err: errors.New(`ent: missing required field "GithubGist.raw_url"`)}
	}
	if _, ok := ggc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "GithubGist.content"`)}
	}
	return nil
}

func (ggc *GithubGistCreate) sqlSave(ctx context.Context) (*GithubGist, error) {
	if err := ggc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ggc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ggc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ggc.mutation.id = &_node.ID
	ggc.mutation.done = true
	return _node, nil
}

func (ggc *GithubGistCreate) createSpec() (*GithubGist, *sqlgraph.CreateSpec) {
	var (
		_node = &GithubGist{config: ggc.config}
		_spec = sqlgraph.NewCreateSpec(githubgist.Table, sqlgraph.NewFieldSpec(githubgist.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ggc.conflict
	if value, ok := ggc.mutation.GistID(); ok {
		_spec.SetField(githubgist.FieldGistID, field.TypeString, value)
		_node.GistID = value
	}
	if value, ok := ggc.mutation.HTMLURL(); ok {
		_spec.SetField(githubgist.FieldHTMLURL, field.TypeString, value)
		_node.HTMLURL = value
	}
	if value, ok := ggc.mutation.Public(); ok {
		_spec.SetField(githubgist.FieldPublic, field.TypeBool, value)
		_node.Public = value
	}
	if value, ok := ggc.mutation.CreatedAt(); ok {
		_spec.SetField(githubgist.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ggc.mutation.UpdatedAt(); ok {
		_spec.SetField(githubgist.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ggc.mutation.Description(); ok {
		_spec.SetField(githubgist.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ggc.mutation.Owner(); ok {
		_spec.SetField(githubgist.FieldOwner, field.TypeJSON, value)
		_node.Owner = value
	}
	if value, ok := ggc.mutation.Name(); ok {
		_spec.SetField(githubgist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ggc.mutation.GetType(); ok {
		_spec.SetField(githubgist.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := ggc.mutation.Language(); ok {
		_spec.SetField(githubgist.FieldLanguage, field.TypeString, value)
		_node.Language = value
	}
	if value, ok := ggc.mutation.Size(); ok {
		_spec.SetField(githubgist.FieldSize, field.TypeInt64, value)
		_node.Size = value
	}
	if value, ok := ggc.mutation.RawURL(); ok {
		_spec.SetField(githubgist.FieldRawURL, field.TypeString, value)
		_node.RawURL = value
	}
	if value, ok := ggc.mutation.Content(); ok {
		_spec.SetField(githubgist.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GithubGist.Create().
//		SetGistID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GithubGistUpsert) {
//			SetGistID(v+v).
//		}).
//		Exec(ctx)
func (ggc *GithubGistCreate) OnConflict(opts ...sql.ConflictOption) *GithubGistUpsertOne {
	ggc.conflict = opts
	return &GithubGistUpsertOne{
		create: ggc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GithubGist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ggc *GithubGistCreate) OnConflictColumns(columns ...string) *GithubGistUpsertOne {
	ggc.conflict = append(ggc.conflict, sql.ConflictColumns(columns...))
	return &GithubGistUpsertOne{
		create: ggc,
	}
}

type (
	// GithubGistUpsertOne is the builder for "upsert"-ing
	//  one GithubGist node.
	GithubGistUpsertOne struct {
		create *GithubGistCreate
	}

	// GithubGistUpsert is the "OnConflict" setter.
	GithubGistUpsert struct {
		*sql.UpdateSet
	}
)

// SetGistID sets the "gist_id" field.
func (u *GithubGistUpsert) SetGistID(v string) *GithubGistUpsert {
	u.Set(githubgist.FieldGistID, v)
	return u
}

// UpdateGistID sets the "gist_id" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateGistID() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldGistID)
	return u
}

// SetHTMLURL sets the "html_url" field.
func (u *GithubGistUpsert) SetHTMLURL(v string) *GithubGistUpsert {
	u.Set(githubgist.FieldHTMLURL, v)
	return u
}

// UpdateHTMLURL sets the "html_url" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateHTMLURL() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldHTMLURL)
	return u
}

// SetPublic sets the "public" field.
func (u *GithubGistUpsert) SetPublic(v bool) *GithubGistUpsert {
	u.Set(githubgist.FieldPublic, v)
	return u
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdatePublic() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldPublic)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubGistUpsert) SetCreatedAt(v time.Time) *GithubGistUpsert {
	u.Set(githubgist.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateCreatedAt() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GithubGistUpsert) SetUpdatedAt(v time.Time) *GithubGistUpsert {
	u.Set(githubgist.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateUpdatedAt() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldUpdatedAt)
	return u
}

// SetDescription sets the "description" field.
func (u *GithubGistUpsert) SetDescription(v string) *GithubGistUpsert {
	u.Set(githubgist.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateDescription() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *GithubGistUpsert) ClearDescription() *GithubGistUpsert {
	u.SetNull(githubgist.FieldDescription)
	return u
}

// SetOwner sets the "owner" field.
func (u *GithubGistUpsert) SetOwner(v *github.User) *GithubGistUpsert {
	u.Set(githubgist.FieldOwner, v)
	return u
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateOwner() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldOwner)
	return u
}

// SetName sets the "name" field.
func (u *GithubGistUpsert) SetName(v string) *GithubGistUpsert {
	u.Set(githubgist.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateName() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldName)
	return u
}

// SetType sets the "type" field.
func (u *GithubGistUpsert) SetType(v string) *GithubGistUpsert {
	u.Set(githubgist.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateType() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldType)
	return u
}

// SetLanguage sets the "language" field.
func (u *GithubGistUpsert) SetLanguage(v string) *GithubGistUpsert {
	u.Set(githubgist.FieldLanguage, v)
	return u
}

// UpdateLanguage sets the "language" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateLanguage() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldLanguage)
	return u
}

// ClearLanguage clears the value of the "language" field.
func (u *GithubGistUpsert) ClearLanguage() *GithubGistUpsert {
	u.SetNull(githubgist.FieldLanguage)
	return u
}

// SetSize sets the "size" field.
func (u *GithubGistUpsert) SetSize(v int64) *GithubGistUpsert {
	u.Set(githubgist.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateSize() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *GithubGistUpsert) AddSize(v int64) *GithubGistUpsert {
	u.Add(githubgist.FieldSize, v)
	return u
}

// SetRawURL sets the "raw_url" field.
func (u *GithubGistUpsert) SetRawURL(v string) *GithubGistUpsert {
	u.Set(githubgist.FieldRawURL, v)
	return u
}

// UpdateRawURL sets the "raw_url" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateRawURL() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldRawURL)
	return u
}

// SetContent sets the "content" field.
func (u *GithubGistUpsert) SetContent(v string) *GithubGistUpsert {
	u.Set(githubgist.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *GithubGistUpsert) UpdateContent() *GithubGistUpsert {
	u.SetExcluded(githubgist.FieldContent)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.GithubGist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GithubGistUpsertOne) UpdateNewValues() *GithubGistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GithubGist.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GithubGistUpsertOne) Ignore() *GithubGistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GithubGistUpsertOne) DoNothing() *GithubGistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GithubGistCreate.OnConflict
// documentation for more info.
func (u *GithubGistUpsertOne) Update(set func(*GithubGistUpsert)) *GithubGistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GithubGistUpsert{UpdateSet: update})
	}))
	return u
}

// SetGistID sets the "gist_id" field.
func (u *GithubGistUpsertOne) SetGistID(v string) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetGistID(v)
	})
}

// UpdateGistID sets the "gist_id" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateGistID() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateGistID()
	})
}

// SetHTMLURL sets the "html_url" field.
func (u *GithubGistUpsertOne) SetHTMLURL(v string) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetHTMLURL(v)
	})
}

// UpdateHTMLURL sets the "html_url" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateHTMLURL() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateHTMLURL()
	})
}

// SetPublic sets the "public" field.
func (u *GithubGistUpsertOne) SetPublic(v bool) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdatePublic() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdatePublic()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubGistUpsertOne) SetCreatedAt(v time.Time) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateCreatedAt() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GithubGistUpsertOne) SetUpdatedAt(v time.Time) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateUpdatedAt() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDescription sets the "description" field.
func (u *GithubGistUpsertOne) SetDescription(v string) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateDescription() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *GithubGistUpsertOne) ClearDescription() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.ClearDescription()
	})
}

// SetOwner sets the "owner" field.
func (u *GithubGistUpsertOne) SetOwner(v *github.User) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetOwner(v)
	})
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateOwner() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateOwner()
	})
}

// SetName sets the "name" field.
func (u *GithubGistUpsertOne) SetName(v string) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateName() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *GithubGistUpsertOne) SetType(v string) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateType() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateType()
	})
}

// SetLanguage sets the "language" field.
func (u *GithubGistUpsertOne) SetLanguage(v string) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetLanguage(v)
	})
}

// UpdateLanguage sets the "language" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateLanguage() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateLanguage()
	})
}

// ClearLanguage clears the value of the "language" field.
func (u *GithubGistUpsertOne) ClearLanguage() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.ClearLanguage()
	})
}

// SetSize sets the "size" field.
func (u *GithubGistUpsertOne) SetSize(v int64) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *GithubGistUpsertOne) AddSize(v int64) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateSize() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateSize()
	})
}

// SetRawURL sets the "raw_url" field.
func (u *GithubGistUpsertOne) SetRawURL(v string) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetRawURL(v)
	})
}

// UpdateRawURL sets the "raw_url" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateRawURL() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateRawURL()
	})
}

// SetContent sets the "content" field.
func (u *GithubGistUpsertOne) SetContent(v string) *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *GithubGistUpsertOne) UpdateContent() *GithubGistUpsertOne {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *GithubGistUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GithubGistCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GithubGistUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GithubGistUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GithubGistUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GithubGistCreateBulk is the builder for creating many GithubGist entities in bulk.
type GithubGistCreateBulk struct {
	config
	builders []*GithubGistCreate
	conflict []sql.ConflictOption
}

// Save creates the GithubGist entities in the database.
func (ggcb *GithubGistCreateBulk) Save(ctx context.Context) ([]*GithubGist, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ggcb.builders))
	nodes := make([]*GithubGist, len(ggcb.builders))
	mutators := make([]Mutator, len(ggcb.builders))
	for i := range ggcb.builders {
		func(i int, root context.Context) {
			builder := ggcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GithubGistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ggcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ggcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ggcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ggcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ggcb *GithubGistCreateBulk) SaveX(ctx context.Context) []*GithubGist {
	v, err := ggcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ggcb *GithubGistCreateBulk) Exec(ctx context.Context) error {
	_, err := ggcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ggcb *GithubGistCreateBulk) ExecX(ctx context.Context) {
	if err := ggcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GithubGist.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GithubGistUpsert) {
//			SetGistID(v+v).
//		}).
//		Exec(ctx)
func (ggcb *GithubGistCreateBulk) OnConflict(opts ...sql.ConflictOption) *GithubGistUpsertBulk {
	ggcb.conflict = opts
	return &GithubGistUpsertBulk{
		create: ggcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GithubGist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ggcb *GithubGistCreateBulk) OnConflictColumns(columns ...string) *GithubGistUpsertBulk {
	ggcb.conflict = append(ggcb.conflict, sql.ConflictColumns(columns...))
	return &GithubGistUpsertBulk{
		create: ggcb,
	}
}

// GithubGistUpsertBulk is the builder for "upsert"-ing
// a bulk of GithubGist nodes.
type GithubGistUpsertBulk struct {
	create *GithubGistCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GithubGist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GithubGistUpsertBulk) UpdateNewValues() *GithubGistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GithubGist.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GithubGistUpsertBulk) Ignore() *GithubGistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GithubGistUpsertBulk) DoNothing() *GithubGistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GithubGistCreateBulk.OnConflict
// documentation for more info.
func (u *GithubGistUpsertBulk) Update(set func(*GithubGistUpsert)) *GithubGistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GithubGistUpsert{UpdateSet: update})
	}))
	return u
}

// SetGistID sets the "gist_id" field.
func (u *GithubGistUpsertBulk) SetGistID(v string) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetGistID(v)
	})
}

// UpdateGistID sets the "gist_id" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateGistID() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateGistID()
	})
}

// SetHTMLURL sets the "html_url" field.
func (u *GithubGistUpsertBulk) SetHTMLURL(v string) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetHTMLURL(v)
	})
}

// UpdateHTMLURL sets the "html_url" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateHTMLURL() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateHTMLURL()
	})
}

// SetPublic sets the "public" field.
func (u *GithubGistUpsertBulk) SetPublic(v bool) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdatePublic() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdatePublic()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubGistUpsertBulk) SetCreatedAt(v time.Time) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateCreatedAt() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GithubGistUpsertBulk) SetUpdatedAt(v time.Time) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateUpdatedAt() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDescription sets the "description" field.
func (u *GithubGistUpsertBulk) SetDescription(v string) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateDescription() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *GithubGistUpsertBulk) ClearDescription() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.ClearDescription()
	})
}

// SetOwner sets the "owner" field.
func (u *GithubGistUpsertBulk) SetOwner(v *github.User) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetOwner(v)
	})
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateOwner() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateOwner()
	})
}

// SetName sets the "name" field.
func (u *GithubGistUpsertBulk) SetName(v string) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateName() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *GithubGistUpsertBulk) SetType(v string) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateType() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateType()
	})
}

// SetLanguage sets the "language" field.
func (u *GithubGistUpsertBulk) SetLanguage(v string) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetLanguage(v)
	})
}

// UpdateLanguage sets the "language" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateLanguage() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateLanguage()
	})
}

// ClearLanguage clears the value of the "language" field.
func (u *GithubGistUpsertBulk) ClearLanguage() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.ClearLanguage()
	})
}

// SetSize sets the "size" field.
func (u *GithubGistUpsertBulk) SetSize(v int64) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *GithubGistUpsertBulk) AddSize(v int64) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateSize() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateSize()
	})
}

// SetRawURL sets the "raw_url" field.
func (u *GithubGistUpsertBulk) SetRawURL(v string) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetRawURL(v)
	})
}

// UpdateRawURL sets the "raw_url" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateRawURL() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateRawURL()
	})
}

// SetContent sets the "content" field.
func (u *GithubGistUpsertBulk) SetContent(v string) *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *GithubGistUpsertBulk) UpdateContent() *GithubGistUpsertBulk {
	return u.Update(func(s *GithubGistUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *GithubGistUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GithubGistCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GithubGistCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GithubGistUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
