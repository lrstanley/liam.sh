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
	"github.com/lrstanley/liam.sh/internal/database/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrelease"
)

// GithubAssetCreate is the builder for creating a GithubAsset entity.
type GithubAssetCreate struct {
	config
	mutation *GithubAssetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAssetID sets the "asset_id" field.
func (gac *GithubAssetCreate) SetAssetID(i int64) *GithubAssetCreate {
	gac.mutation.SetAssetID(i)
	return gac
}

// SetBrowserDownloadURL sets the "browser_download_url" field.
func (gac *GithubAssetCreate) SetBrowserDownloadURL(s string) *GithubAssetCreate {
	gac.mutation.SetBrowserDownloadURL(s)
	return gac
}

// SetName sets the "name" field.
func (gac *GithubAssetCreate) SetName(s string) *GithubAssetCreate {
	gac.mutation.SetName(s)
	return gac
}

// SetLabel sets the "label" field.
func (gac *GithubAssetCreate) SetLabel(s string) *GithubAssetCreate {
	gac.mutation.SetLabel(s)
	return gac
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (gac *GithubAssetCreate) SetNillableLabel(s *string) *GithubAssetCreate {
	if s != nil {
		gac.SetLabel(*s)
	}
	return gac
}

// SetState sets the "state" field.
func (gac *GithubAssetCreate) SetState(s string) *GithubAssetCreate {
	gac.mutation.SetState(s)
	return gac
}

// SetNillableState sets the "state" field if the given value is not nil.
func (gac *GithubAssetCreate) SetNillableState(s *string) *GithubAssetCreate {
	if s != nil {
		gac.SetState(*s)
	}
	return gac
}

// SetContentType sets the "content_type" field.
func (gac *GithubAssetCreate) SetContentType(s string) *GithubAssetCreate {
	gac.mutation.SetContentType(s)
	return gac
}

// SetSize sets the "size" field.
func (gac *GithubAssetCreate) SetSize(i int64) *GithubAssetCreate {
	gac.mutation.SetSize(i)
	return gac
}

// SetDownloadCount sets the "download_count" field.
func (gac *GithubAssetCreate) SetDownloadCount(i int64) *GithubAssetCreate {
	gac.mutation.SetDownloadCount(i)
	return gac
}

// SetCreatedAt sets the "created_at" field.
func (gac *GithubAssetCreate) SetCreatedAt(t time.Time) *GithubAssetCreate {
	gac.mutation.SetCreatedAt(t)
	return gac
}

// SetUpdatedAt sets the "updated_at" field.
func (gac *GithubAssetCreate) SetUpdatedAt(t time.Time) *GithubAssetCreate {
	gac.mutation.SetUpdatedAt(t)
	return gac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gac *GithubAssetCreate) SetNillableUpdatedAt(t *time.Time) *GithubAssetCreate {
	if t != nil {
		gac.SetUpdatedAt(*t)
	}
	return gac
}

// SetUploader sets the "uploader" field.
func (gac *GithubAssetCreate) SetUploader(gi *github.User) *GithubAssetCreate {
	gac.mutation.SetUploader(gi)
	return gac
}

// SetReleaseID sets the "release" edge to the GithubRelease entity by ID.
func (gac *GithubAssetCreate) SetReleaseID(id int) *GithubAssetCreate {
	gac.mutation.SetReleaseID(id)
	return gac
}

// SetRelease sets the "release" edge to the GithubRelease entity.
func (gac *GithubAssetCreate) SetRelease(g *GithubRelease) *GithubAssetCreate {
	return gac.SetReleaseID(g.ID)
}

// Mutation returns the GithubAssetMutation object of the builder.
func (gac *GithubAssetCreate) Mutation() *GithubAssetMutation {
	return gac.mutation
}

// Save creates the GithubAsset in the database.
func (gac *GithubAssetCreate) Save(ctx context.Context) (*GithubAsset, error) {
	return withHooks(ctx, gac.sqlSave, gac.mutation, gac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gac *GithubAssetCreate) SaveX(ctx context.Context) *GithubAsset {
	v, err := gac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gac *GithubAssetCreate) Exec(ctx context.Context) error {
	_, err := gac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gac *GithubAssetCreate) ExecX(ctx context.Context) {
	if err := gac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gac *GithubAssetCreate) check() error {
	if _, ok := gac.mutation.AssetID(); !ok {
		return &ValidationError{Name: "asset_id", err: errors.New(`ent: missing required field "GithubAsset.asset_id"`)}
	}
	if v, ok := gac.mutation.AssetID(); ok {
		if err := githubasset.AssetIDValidator(v); err != nil {
			return &ValidationError{Name: "asset_id", err: fmt.Errorf(`ent: validator failed for field "GithubAsset.asset_id": %w`, err)}
		}
	}
	if _, ok := gac.mutation.BrowserDownloadURL(); !ok {
		return &ValidationError{Name: "browser_download_url", err: errors.New(`ent: missing required field "GithubAsset.browser_download_url"`)}
	}
	if v, ok := gac.mutation.BrowserDownloadURL(); ok {
		if err := githubasset.BrowserDownloadURLValidator(v); err != nil {
			return &ValidationError{Name: "browser_download_url", err: fmt.Errorf(`ent: validator failed for field "GithubAsset.browser_download_url": %w`, err)}
		}
	}
	if _, ok := gac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "GithubAsset.name"`)}
	}
	if v, ok := gac.mutation.Name(); ok {
		if err := githubasset.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "GithubAsset.name": %w`, err)}
		}
	}
	if _, ok := gac.mutation.ContentType(); !ok {
		return &ValidationError{Name: "content_type", err: errors.New(`ent: missing required field "GithubAsset.content_type"`)}
	}
	if _, ok := gac.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "GithubAsset.size"`)}
	}
	if _, ok := gac.mutation.DownloadCount(); !ok {
		return &ValidationError{Name: "download_count", err: errors.New(`ent: missing required field "GithubAsset.download_count"`)}
	}
	if _, ok := gac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GithubAsset.created_at"`)}
	}
	if _, ok := gac.mutation.Uploader(); !ok {
		return &ValidationError{Name: "uploader", err: errors.New(`ent: missing required field "GithubAsset.uploader"`)}
	}
	if _, ok := gac.mutation.ReleaseID(); !ok {
		return &ValidationError{Name: "release", err: errors.New(`ent: missing required edge "GithubAsset.release"`)}
	}
	return nil
}

func (gac *GithubAssetCreate) sqlSave(ctx context.Context) (*GithubAsset, error) {
	if err := gac.check(); err != nil {
		return nil, err
	}
	_node, _spec := gac.createSpec()
	if err := sqlgraph.CreateNode(ctx, gac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gac.mutation.id = &_node.ID
	gac.mutation.done = true
	return _node, nil
}

func (gac *GithubAssetCreate) createSpec() (*GithubAsset, *sqlgraph.CreateSpec) {
	var (
		_node = &GithubAsset{config: gac.config}
		_spec = sqlgraph.NewCreateSpec(githubasset.Table, sqlgraph.NewFieldSpec(githubasset.FieldID, field.TypeInt))
	)
	_spec.OnConflict = gac.conflict
	if value, ok := gac.mutation.AssetID(); ok {
		_spec.SetField(githubasset.FieldAssetID, field.TypeInt64, value)
		_node.AssetID = value
	}
	if value, ok := gac.mutation.BrowserDownloadURL(); ok {
		_spec.SetField(githubasset.FieldBrowserDownloadURL, field.TypeString, value)
		_node.BrowserDownloadURL = value
	}
	if value, ok := gac.mutation.Name(); ok {
		_spec.SetField(githubasset.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gac.mutation.Label(); ok {
		_spec.SetField(githubasset.FieldLabel, field.TypeString, value)
		_node.Label = value
	}
	if value, ok := gac.mutation.State(); ok {
		_spec.SetField(githubasset.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := gac.mutation.ContentType(); ok {
		_spec.SetField(githubasset.FieldContentType, field.TypeString, value)
		_node.ContentType = value
	}
	if value, ok := gac.mutation.Size(); ok {
		_spec.SetField(githubasset.FieldSize, field.TypeInt64, value)
		_node.Size = value
	}
	if value, ok := gac.mutation.DownloadCount(); ok {
		_spec.SetField(githubasset.FieldDownloadCount, field.TypeInt64, value)
		_node.DownloadCount = value
	}
	if value, ok := gac.mutation.CreatedAt(); ok {
		_spec.SetField(githubasset.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gac.mutation.UpdatedAt(); ok {
		_spec.SetField(githubasset.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gac.mutation.Uploader(); ok {
		_spec.SetField(githubasset.FieldUploader, field.TypeJSON, value)
		_node.Uploader = value
	}
	if nodes := gac.mutation.ReleaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   githubasset.ReleaseTable,
			Columns: []string{githubasset.ReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrelease.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.github_release_assets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GithubAsset.Create().
//		SetAssetID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GithubAssetUpsert) {
//			SetAssetID(v+v).
//		}).
//		Exec(ctx)
func (gac *GithubAssetCreate) OnConflict(opts ...sql.ConflictOption) *GithubAssetUpsertOne {
	gac.conflict = opts
	return &GithubAssetUpsertOne{
		create: gac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GithubAsset.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gac *GithubAssetCreate) OnConflictColumns(columns ...string) *GithubAssetUpsertOne {
	gac.conflict = append(gac.conflict, sql.ConflictColumns(columns...))
	return &GithubAssetUpsertOne{
		create: gac,
	}
}

type (
	// GithubAssetUpsertOne is the builder for "upsert"-ing
	//  one GithubAsset node.
	GithubAssetUpsertOne struct {
		create *GithubAssetCreate
	}

	// GithubAssetUpsert is the "OnConflict" setter.
	GithubAssetUpsert struct {
		*sql.UpdateSet
	}
)

// SetAssetID sets the "asset_id" field.
func (u *GithubAssetUpsert) SetAssetID(v int64) *GithubAssetUpsert {
	u.Set(githubasset.FieldAssetID, v)
	return u
}

// UpdateAssetID sets the "asset_id" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateAssetID() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldAssetID)
	return u
}

// AddAssetID adds v to the "asset_id" field.
func (u *GithubAssetUpsert) AddAssetID(v int64) *GithubAssetUpsert {
	u.Add(githubasset.FieldAssetID, v)
	return u
}

// SetBrowserDownloadURL sets the "browser_download_url" field.
func (u *GithubAssetUpsert) SetBrowserDownloadURL(v string) *GithubAssetUpsert {
	u.Set(githubasset.FieldBrowserDownloadURL, v)
	return u
}

// UpdateBrowserDownloadURL sets the "browser_download_url" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateBrowserDownloadURL() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldBrowserDownloadURL)
	return u
}

// SetName sets the "name" field.
func (u *GithubAssetUpsert) SetName(v string) *GithubAssetUpsert {
	u.Set(githubasset.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateName() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldName)
	return u
}

// SetLabel sets the "label" field.
func (u *GithubAssetUpsert) SetLabel(v string) *GithubAssetUpsert {
	u.Set(githubasset.FieldLabel, v)
	return u
}

// UpdateLabel sets the "label" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateLabel() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldLabel)
	return u
}

// ClearLabel clears the value of the "label" field.
func (u *GithubAssetUpsert) ClearLabel() *GithubAssetUpsert {
	u.SetNull(githubasset.FieldLabel)
	return u
}

// SetState sets the "state" field.
func (u *GithubAssetUpsert) SetState(v string) *GithubAssetUpsert {
	u.Set(githubasset.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateState() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldState)
	return u
}

// ClearState clears the value of the "state" field.
func (u *GithubAssetUpsert) ClearState() *GithubAssetUpsert {
	u.SetNull(githubasset.FieldState)
	return u
}

// SetContentType sets the "content_type" field.
func (u *GithubAssetUpsert) SetContentType(v string) *GithubAssetUpsert {
	u.Set(githubasset.FieldContentType, v)
	return u
}

// UpdateContentType sets the "content_type" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateContentType() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldContentType)
	return u
}

// SetSize sets the "size" field.
func (u *GithubAssetUpsert) SetSize(v int64) *GithubAssetUpsert {
	u.Set(githubasset.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateSize() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *GithubAssetUpsert) AddSize(v int64) *GithubAssetUpsert {
	u.Add(githubasset.FieldSize, v)
	return u
}

// SetDownloadCount sets the "download_count" field.
func (u *GithubAssetUpsert) SetDownloadCount(v int64) *GithubAssetUpsert {
	u.Set(githubasset.FieldDownloadCount, v)
	return u
}

// UpdateDownloadCount sets the "download_count" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateDownloadCount() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldDownloadCount)
	return u
}

// AddDownloadCount adds v to the "download_count" field.
func (u *GithubAssetUpsert) AddDownloadCount(v int64) *GithubAssetUpsert {
	u.Add(githubasset.FieldDownloadCount, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubAssetUpsert) SetCreatedAt(v time.Time) *GithubAssetUpsert {
	u.Set(githubasset.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateCreatedAt() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GithubAssetUpsert) SetUpdatedAt(v time.Time) *GithubAssetUpsert {
	u.Set(githubasset.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateUpdatedAt() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *GithubAssetUpsert) ClearUpdatedAt() *GithubAssetUpsert {
	u.SetNull(githubasset.FieldUpdatedAt)
	return u
}

// SetUploader sets the "uploader" field.
func (u *GithubAssetUpsert) SetUploader(v *github.User) *GithubAssetUpsert {
	u.Set(githubasset.FieldUploader, v)
	return u
}

// UpdateUploader sets the "uploader" field to the value that was provided on create.
func (u *GithubAssetUpsert) UpdateUploader() *GithubAssetUpsert {
	u.SetExcluded(githubasset.FieldUploader)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.GithubAsset.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GithubAssetUpsertOne) UpdateNewValues() *GithubAssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GithubAsset.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GithubAssetUpsertOne) Ignore() *GithubAssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GithubAssetUpsertOne) DoNothing() *GithubAssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GithubAssetCreate.OnConflict
// documentation for more info.
func (u *GithubAssetUpsertOne) Update(set func(*GithubAssetUpsert)) *GithubAssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GithubAssetUpsert{UpdateSet: update})
	}))
	return u
}

// SetAssetID sets the "asset_id" field.
func (u *GithubAssetUpsertOne) SetAssetID(v int64) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetAssetID(v)
	})
}

// AddAssetID adds v to the "asset_id" field.
func (u *GithubAssetUpsertOne) AddAssetID(v int64) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.AddAssetID(v)
	})
}

// UpdateAssetID sets the "asset_id" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateAssetID() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateAssetID()
	})
}

// SetBrowserDownloadURL sets the "browser_download_url" field.
func (u *GithubAssetUpsertOne) SetBrowserDownloadURL(v string) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetBrowserDownloadURL(v)
	})
}

// UpdateBrowserDownloadURL sets the "browser_download_url" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateBrowserDownloadURL() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateBrowserDownloadURL()
	})
}

// SetName sets the "name" field.
func (u *GithubAssetUpsertOne) SetName(v string) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateName() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateName()
	})
}

// SetLabel sets the "label" field.
func (u *GithubAssetUpsertOne) SetLabel(v string) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetLabel(v)
	})
}

// UpdateLabel sets the "label" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateLabel() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateLabel()
	})
}

// ClearLabel clears the value of the "label" field.
func (u *GithubAssetUpsertOne) ClearLabel() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.ClearLabel()
	})
}

// SetState sets the "state" field.
func (u *GithubAssetUpsertOne) SetState(v string) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateState() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateState()
	})
}

// ClearState clears the value of the "state" field.
func (u *GithubAssetUpsertOne) ClearState() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.ClearState()
	})
}

// SetContentType sets the "content_type" field.
func (u *GithubAssetUpsertOne) SetContentType(v string) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetContentType(v)
	})
}

// UpdateContentType sets the "content_type" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateContentType() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateContentType()
	})
}

// SetSize sets the "size" field.
func (u *GithubAssetUpsertOne) SetSize(v int64) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *GithubAssetUpsertOne) AddSize(v int64) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateSize() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateSize()
	})
}

// SetDownloadCount sets the "download_count" field.
func (u *GithubAssetUpsertOne) SetDownloadCount(v int64) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetDownloadCount(v)
	})
}

// AddDownloadCount adds v to the "download_count" field.
func (u *GithubAssetUpsertOne) AddDownloadCount(v int64) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.AddDownloadCount(v)
	})
}

// UpdateDownloadCount sets the "download_count" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateDownloadCount() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateDownloadCount()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubAssetUpsertOne) SetCreatedAt(v time.Time) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateCreatedAt() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GithubAssetUpsertOne) SetUpdatedAt(v time.Time) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateUpdatedAt() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *GithubAssetUpsertOne) ClearUpdatedAt() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetUploader sets the "uploader" field.
func (u *GithubAssetUpsertOne) SetUploader(v *github.User) *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetUploader(v)
	})
}

// UpdateUploader sets the "uploader" field to the value that was provided on create.
func (u *GithubAssetUpsertOne) UpdateUploader() *GithubAssetUpsertOne {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateUploader()
	})
}

// Exec executes the query.
func (u *GithubAssetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GithubAssetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GithubAssetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GithubAssetUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GithubAssetUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GithubAssetCreateBulk is the builder for creating many GithubAsset entities in bulk.
type GithubAssetCreateBulk struct {
	config
	err      error
	builders []*GithubAssetCreate
	conflict []sql.ConflictOption
}

// Save creates the GithubAsset entities in the database.
func (gacb *GithubAssetCreateBulk) Save(ctx context.Context) ([]*GithubAsset, error) {
	if gacb.err != nil {
		return nil, gacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gacb.builders))
	nodes := make([]*GithubAsset, len(gacb.builders))
	mutators := make([]Mutator, len(gacb.builders))
	for i := range gacb.builders {
		func(i int, root context.Context) {
			builder := gacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GithubAssetMutation)
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
					_, err = mutators[i+1].Mutate(root, gacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gacb *GithubAssetCreateBulk) SaveX(ctx context.Context) []*GithubAsset {
	v, err := gacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gacb *GithubAssetCreateBulk) Exec(ctx context.Context) error {
	_, err := gacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gacb *GithubAssetCreateBulk) ExecX(ctx context.Context) {
	if err := gacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GithubAsset.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GithubAssetUpsert) {
//			SetAssetID(v+v).
//		}).
//		Exec(ctx)
func (gacb *GithubAssetCreateBulk) OnConflict(opts ...sql.ConflictOption) *GithubAssetUpsertBulk {
	gacb.conflict = opts
	return &GithubAssetUpsertBulk{
		create: gacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GithubAsset.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gacb *GithubAssetCreateBulk) OnConflictColumns(columns ...string) *GithubAssetUpsertBulk {
	gacb.conflict = append(gacb.conflict, sql.ConflictColumns(columns...))
	return &GithubAssetUpsertBulk{
		create: gacb,
	}
}

// GithubAssetUpsertBulk is the builder for "upsert"-ing
// a bulk of GithubAsset nodes.
type GithubAssetUpsertBulk struct {
	create *GithubAssetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GithubAsset.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GithubAssetUpsertBulk) UpdateNewValues() *GithubAssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GithubAsset.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GithubAssetUpsertBulk) Ignore() *GithubAssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GithubAssetUpsertBulk) DoNothing() *GithubAssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GithubAssetCreateBulk.OnConflict
// documentation for more info.
func (u *GithubAssetUpsertBulk) Update(set func(*GithubAssetUpsert)) *GithubAssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GithubAssetUpsert{UpdateSet: update})
	}))
	return u
}

// SetAssetID sets the "asset_id" field.
func (u *GithubAssetUpsertBulk) SetAssetID(v int64) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetAssetID(v)
	})
}

// AddAssetID adds v to the "asset_id" field.
func (u *GithubAssetUpsertBulk) AddAssetID(v int64) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.AddAssetID(v)
	})
}

// UpdateAssetID sets the "asset_id" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateAssetID() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateAssetID()
	})
}

// SetBrowserDownloadURL sets the "browser_download_url" field.
func (u *GithubAssetUpsertBulk) SetBrowserDownloadURL(v string) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetBrowserDownloadURL(v)
	})
}

// UpdateBrowserDownloadURL sets the "browser_download_url" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateBrowserDownloadURL() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateBrowserDownloadURL()
	})
}

// SetName sets the "name" field.
func (u *GithubAssetUpsertBulk) SetName(v string) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateName() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateName()
	})
}

// SetLabel sets the "label" field.
func (u *GithubAssetUpsertBulk) SetLabel(v string) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetLabel(v)
	})
}

// UpdateLabel sets the "label" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateLabel() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateLabel()
	})
}

// ClearLabel clears the value of the "label" field.
func (u *GithubAssetUpsertBulk) ClearLabel() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.ClearLabel()
	})
}

// SetState sets the "state" field.
func (u *GithubAssetUpsertBulk) SetState(v string) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateState() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateState()
	})
}

// ClearState clears the value of the "state" field.
func (u *GithubAssetUpsertBulk) ClearState() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.ClearState()
	})
}

// SetContentType sets the "content_type" field.
func (u *GithubAssetUpsertBulk) SetContentType(v string) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetContentType(v)
	})
}

// UpdateContentType sets the "content_type" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateContentType() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateContentType()
	})
}

// SetSize sets the "size" field.
func (u *GithubAssetUpsertBulk) SetSize(v int64) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *GithubAssetUpsertBulk) AddSize(v int64) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateSize() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateSize()
	})
}

// SetDownloadCount sets the "download_count" field.
func (u *GithubAssetUpsertBulk) SetDownloadCount(v int64) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetDownloadCount(v)
	})
}

// AddDownloadCount adds v to the "download_count" field.
func (u *GithubAssetUpsertBulk) AddDownloadCount(v int64) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.AddDownloadCount(v)
	})
}

// UpdateDownloadCount sets the "download_count" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateDownloadCount() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateDownloadCount()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubAssetUpsertBulk) SetCreatedAt(v time.Time) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateCreatedAt() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GithubAssetUpsertBulk) SetUpdatedAt(v time.Time) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateUpdatedAt() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *GithubAssetUpsertBulk) ClearUpdatedAt() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetUploader sets the "uploader" field.
func (u *GithubAssetUpsertBulk) SetUploader(v *github.User) *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.SetUploader(v)
	})
}

// UpdateUploader sets the "uploader" field to the value that was provided on create.
func (u *GithubAssetUpsertBulk) UpdateUploader() *GithubAssetUpsertBulk {
	return u.Update(func(s *GithubAssetUpsert) {
		s.UpdateUploader()
	})
}

// Exec executes the query.
func (u *GithubAssetUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GithubAssetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GithubAssetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GithubAssetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
