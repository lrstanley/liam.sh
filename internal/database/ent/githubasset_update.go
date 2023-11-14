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
	"github.com/lrstanley/liam.sh/internal/database/ent/predicate"
)

// GithubAssetUpdate is the builder for updating GithubAsset entities.
type GithubAssetUpdate struct {
	config
	hooks    []Hook
	mutation *GithubAssetMutation
}

// Where appends a list predicates to the GithubAssetUpdate builder.
func (gau *GithubAssetUpdate) Where(ps ...predicate.GithubAsset) *GithubAssetUpdate {
	gau.mutation.Where(ps...)
	return gau
}

// SetAssetID sets the "asset_id" field.
func (gau *GithubAssetUpdate) SetAssetID(i int64) *GithubAssetUpdate {
	gau.mutation.ResetAssetID()
	gau.mutation.SetAssetID(i)
	return gau
}

// SetNillableAssetID sets the "asset_id" field if the given value is not nil.
func (gau *GithubAssetUpdate) SetNillableAssetID(i *int64) *GithubAssetUpdate {
	if i != nil {
		gau.SetAssetID(*i)
	}
	return gau
}

// AddAssetID adds i to the "asset_id" field.
func (gau *GithubAssetUpdate) AddAssetID(i int64) *GithubAssetUpdate {
	gau.mutation.AddAssetID(i)
	return gau
}

// SetBrowserDownloadURL sets the "browser_download_url" field.
func (gau *GithubAssetUpdate) SetBrowserDownloadURL(s string) *GithubAssetUpdate {
	gau.mutation.SetBrowserDownloadURL(s)
	return gau
}

// SetNillableBrowserDownloadURL sets the "browser_download_url" field if the given value is not nil.
func (gau *GithubAssetUpdate) SetNillableBrowserDownloadURL(s *string) *GithubAssetUpdate {
	if s != nil {
		gau.SetBrowserDownloadURL(*s)
	}
	return gau
}

// SetName sets the "name" field.
func (gau *GithubAssetUpdate) SetName(s string) *GithubAssetUpdate {
	gau.mutation.SetName(s)
	return gau
}

// SetNillableName sets the "name" field if the given value is not nil.
func (gau *GithubAssetUpdate) SetNillableName(s *string) *GithubAssetUpdate {
	if s != nil {
		gau.SetName(*s)
	}
	return gau
}

// SetLabel sets the "label" field.
func (gau *GithubAssetUpdate) SetLabel(s string) *GithubAssetUpdate {
	gau.mutation.SetLabel(s)
	return gau
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (gau *GithubAssetUpdate) SetNillableLabel(s *string) *GithubAssetUpdate {
	if s != nil {
		gau.SetLabel(*s)
	}
	return gau
}

// ClearLabel clears the value of the "label" field.
func (gau *GithubAssetUpdate) ClearLabel() *GithubAssetUpdate {
	gau.mutation.ClearLabel()
	return gau
}

// SetState sets the "state" field.
func (gau *GithubAssetUpdate) SetState(s string) *GithubAssetUpdate {
	gau.mutation.SetState(s)
	return gau
}

// SetNillableState sets the "state" field if the given value is not nil.
func (gau *GithubAssetUpdate) SetNillableState(s *string) *GithubAssetUpdate {
	if s != nil {
		gau.SetState(*s)
	}
	return gau
}

// ClearState clears the value of the "state" field.
func (gau *GithubAssetUpdate) ClearState() *GithubAssetUpdate {
	gau.mutation.ClearState()
	return gau
}

// SetContentType sets the "content_type" field.
func (gau *GithubAssetUpdate) SetContentType(s string) *GithubAssetUpdate {
	gau.mutation.SetContentType(s)
	return gau
}

// SetNillableContentType sets the "content_type" field if the given value is not nil.
func (gau *GithubAssetUpdate) SetNillableContentType(s *string) *GithubAssetUpdate {
	if s != nil {
		gau.SetContentType(*s)
	}
	return gau
}

// SetSize sets the "size" field.
func (gau *GithubAssetUpdate) SetSize(i int64) *GithubAssetUpdate {
	gau.mutation.ResetSize()
	gau.mutation.SetSize(i)
	return gau
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (gau *GithubAssetUpdate) SetNillableSize(i *int64) *GithubAssetUpdate {
	if i != nil {
		gau.SetSize(*i)
	}
	return gau
}

// AddSize adds i to the "size" field.
func (gau *GithubAssetUpdate) AddSize(i int64) *GithubAssetUpdate {
	gau.mutation.AddSize(i)
	return gau
}

// SetDownloadCount sets the "download_count" field.
func (gau *GithubAssetUpdate) SetDownloadCount(i int64) *GithubAssetUpdate {
	gau.mutation.ResetDownloadCount()
	gau.mutation.SetDownloadCount(i)
	return gau
}

// SetNillableDownloadCount sets the "download_count" field if the given value is not nil.
func (gau *GithubAssetUpdate) SetNillableDownloadCount(i *int64) *GithubAssetUpdate {
	if i != nil {
		gau.SetDownloadCount(*i)
	}
	return gau
}

// AddDownloadCount adds i to the "download_count" field.
func (gau *GithubAssetUpdate) AddDownloadCount(i int64) *GithubAssetUpdate {
	gau.mutation.AddDownloadCount(i)
	return gau
}

// SetCreatedAt sets the "created_at" field.
func (gau *GithubAssetUpdate) SetCreatedAt(t time.Time) *GithubAssetUpdate {
	gau.mutation.SetCreatedAt(t)
	return gau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gau *GithubAssetUpdate) SetNillableCreatedAt(t *time.Time) *GithubAssetUpdate {
	if t != nil {
		gau.SetCreatedAt(*t)
	}
	return gau
}

// SetUpdatedAt sets the "updated_at" field.
func (gau *GithubAssetUpdate) SetUpdatedAt(t time.Time) *GithubAssetUpdate {
	gau.mutation.SetUpdatedAt(t)
	return gau
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gau *GithubAssetUpdate) SetNillableUpdatedAt(t *time.Time) *GithubAssetUpdate {
	if t != nil {
		gau.SetUpdatedAt(*t)
	}
	return gau
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gau *GithubAssetUpdate) ClearUpdatedAt() *GithubAssetUpdate {
	gau.mutation.ClearUpdatedAt()
	return gau
}

// SetUploader sets the "uploader" field.
func (gau *GithubAssetUpdate) SetUploader(gi *github.User) *GithubAssetUpdate {
	gau.mutation.SetUploader(gi)
	return gau
}

// SetReleaseID sets the "release" edge to the GithubRelease entity by ID.
func (gau *GithubAssetUpdate) SetReleaseID(id int) *GithubAssetUpdate {
	gau.mutation.SetReleaseID(id)
	return gau
}

// SetRelease sets the "release" edge to the GithubRelease entity.
func (gau *GithubAssetUpdate) SetRelease(g *GithubRelease) *GithubAssetUpdate {
	return gau.SetReleaseID(g.ID)
}

// Mutation returns the GithubAssetMutation object of the builder.
func (gau *GithubAssetUpdate) Mutation() *GithubAssetMutation {
	return gau.mutation
}

// ClearRelease clears the "release" edge to the GithubRelease entity.
func (gau *GithubAssetUpdate) ClearRelease() *GithubAssetUpdate {
	gau.mutation.ClearRelease()
	return gau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gau *GithubAssetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gau.sqlSave, gau.mutation, gau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gau *GithubAssetUpdate) SaveX(ctx context.Context) int {
	affected, err := gau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gau *GithubAssetUpdate) Exec(ctx context.Context) error {
	_, err := gau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gau *GithubAssetUpdate) ExecX(ctx context.Context) {
	if err := gau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gau *GithubAssetUpdate) check() error {
	if v, ok := gau.mutation.AssetID(); ok {
		if err := githubasset.AssetIDValidator(v); err != nil {
			return &ValidationError{Name: "asset_id", err: fmt.Errorf(`ent: validator failed for field "GithubAsset.asset_id": %w`, err)}
		}
	}
	if v, ok := gau.mutation.BrowserDownloadURL(); ok {
		if err := githubasset.BrowserDownloadURLValidator(v); err != nil {
			return &ValidationError{Name: "browser_download_url", err: fmt.Errorf(`ent: validator failed for field "GithubAsset.browser_download_url": %w`, err)}
		}
	}
	if v, ok := gau.mutation.Name(); ok {
		if err := githubasset.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "GithubAsset.name": %w`, err)}
		}
	}
	if _, ok := gau.mutation.ReleaseID(); gau.mutation.ReleaseCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GithubAsset.release"`)
	}
	return nil
}

func (gau *GithubAssetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(githubasset.Table, githubasset.Columns, sqlgraph.NewFieldSpec(githubasset.FieldID, field.TypeInt))
	if ps := gau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gau.mutation.AssetID(); ok {
		_spec.SetField(githubasset.FieldAssetID, field.TypeInt64, value)
	}
	if value, ok := gau.mutation.AddedAssetID(); ok {
		_spec.AddField(githubasset.FieldAssetID, field.TypeInt64, value)
	}
	if value, ok := gau.mutation.BrowserDownloadURL(); ok {
		_spec.SetField(githubasset.FieldBrowserDownloadURL, field.TypeString, value)
	}
	if value, ok := gau.mutation.Name(); ok {
		_spec.SetField(githubasset.FieldName, field.TypeString, value)
	}
	if value, ok := gau.mutation.Label(); ok {
		_spec.SetField(githubasset.FieldLabel, field.TypeString, value)
	}
	if gau.mutation.LabelCleared() {
		_spec.ClearField(githubasset.FieldLabel, field.TypeString)
	}
	if value, ok := gau.mutation.State(); ok {
		_spec.SetField(githubasset.FieldState, field.TypeString, value)
	}
	if gau.mutation.StateCleared() {
		_spec.ClearField(githubasset.FieldState, field.TypeString)
	}
	if value, ok := gau.mutation.ContentType(); ok {
		_spec.SetField(githubasset.FieldContentType, field.TypeString, value)
	}
	if value, ok := gau.mutation.Size(); ok {
		_spec.SetField(githubasset.FieldSize, field.TypeInt64, value)
	}
	if value, ok := gau.mutation.AddedSize(); ok {
		_spec.AddField(githubasset.FieldSize, field.TypeInt64, value)
	}
	if value, ok := gau.mutation.DownloadCount(); ok {
		_spec.SetField(githubasset.FieldDownloadCount, field.TypeInt64, value)
	}
	if value, ok := gau.mutation.AddedDownloadCount(); ok {
		_spec.AddField(githubasset.FieldDownloadCount, field.TypeInt64, value)
	}
	if value, ok := gau.mutation.CreatedAt(); ok {
		_spec.SetField(githubasset.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := gau.mutation.UpdatedAt(); ok {
		_spec.SetField(githubasset.FieldUpdatedAt, field.TypeTime, value)
	}
	if gau.mutation.UpdatedAtCleared() {
		_spec.ClearField(githubasset.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := gau.mutation.Uploader(); ok {
		_spec.SetField(githubasset.FieldUploader, field.TypeJSON, value)
	}
	if gau.mutation.ReleaseCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gau.mutation.ReleaseIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubasset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gau.mutation.done = true
	return n, nil
}

// GithubAssetUpdateOne is the builder for updating a single GithubAsset entity.
type GithubAssetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GithubAssetMutation
}

// SetAssetID sets the "asset_id" field.
func (gauo *GithubAssetUpdateOne) SetAssetID(i int64) *GithubAssetUpdateOne {
	gauo.mutation.ResetAssetID()
	gauo.mutation.SetAssetID(i)
	return gauo
}

// SetNillableAssetID sets the "asset_id" field if the given value is not nil.
func (gauo *GithubAssetUpdateOne) SetNillableAssetID(i *int64) *GithubAssetUpdateOne {
	if i != nil {
		gauo.SetAssetID(*i)
	}
	return gauo
}

// AddAssetID adds i to the "asset_id" field.
func (gauo *GithubAssetUpdateOne) AddAssetID(i int64) *GithubAssetUpdateOne {
	gauo.mutation.AddAssetID(i)
	return gauo
}

// SetBrowserDownloadURL sets the "browser_download_url" field.
func (gauo *GithubAssetUpdateOne) SetBrowserDownloadURL(s string) *GithubAssetUpdateOne {
	gauo.mutation.SetBrowserDownloadURL(s)
	return gauo
}

// SetNillableBrowserDownloadURL sets the "browser_download_url" field if the given value is not nil.
func (gauo *GithubAssetUpdateOne) SetNillableBrowserDownloadURL(s *string) *GithubAssetUpdateOne {
	if s != nil {
		gauo.SetBrowserDownloadURL(*s)
	}
	return gauo
}

// SetName sets the "name" field.
func (gauo *GithubAssetUpdateOne) SetName(s string) *GithubAssetUpdateOne {
	gauo.mutation.SetName(s)
	return gauo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (gauo *GithubAssetUpdateOne) SetNillableName(s *string) *GithubAssetUpdateOne {
	if s != nil {
		gauo.SetName(*s)
	}
	return gauo
}

// SetLabel sets the "label" field.
func (gauo *GithubAssetUpdateOne) SetLabel(s string) *GithubAssetUpdateOne {
	gauo.mutation.SetLabel(s)
	return gauo
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (gauo *GithubAssetUpdateOne) SetNillableLabel(s *string) *GithubAssetUpdateOne {
	if s != nil {
		gauo.SetLabel(*s)
	}
	return gauo
}

// ClearLabel clears the value of the "label" field.
func (gauo *GithubAssetUpdateOne) ClearLabel() *GithubAssetUpdateOne {
	gauo.mutation.ClearLabel()
	return gauo
}

// SetState sets the "state" field.
func (gauo *GithubAssetUpdateOne) SetState(s string) *GithubAssetUpdateOne {
	gauo.mutation.SetState(s)
	return gauo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (gauo *GithubAssetUpdateOne) SetNillableState(s *string) *GithubAssetUpdateOne {
	if s != nil {
		gauo.SetState(*s)
	}
	return gauo
}

// ClearState clears the value of the "state" field.
func (gauo *GithubAssetUpdateOne) ClearState() *GithubAssetUpdateOne {
	gauo.mutation.ClearState()
	return gauo
}

// SetContentType sets the "content_type" field.
func (gauo *GithubAssetUpdateOne) SetContentType(s string) *GithubAssetUpdateOne {
	gauo.mutation.SetContentType(s)
	return gauo
}

// SetNillableContentType sets the "content_type" field if the given value is not nil.
func (gauo *GithubAssetUpdateOne) SetNillableContentType(s *string) *GithubAssetUpdateOne {
	if s != nil {
		gauo.SetContentType(*s)
	}
	return gauo
}

// SetSize sets the "size" field.
func (gauo *GithubAssetUpdateOne) SetSize(i int64) *GithubAssetUpdateOne {
	gauo.mutation.ResetSize()
	gauo.mutation.SetSize(i)
	return gauo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (gauo *GithubAssetUpdateOne) SetNillableSize(i *int64) *GithubAssetUpdateOne {
	if i != nil {
		gauo.SetSize(*i)
	}
	return gauo
}

// AddSize adds i to the "size" field.
func (gauo *GithubAssetUpdateOne) AddSize(i int64) *GithubAssetUpdateOne {
	gauo.mutation.AddSize(i)
	return gauo
}

// SetDownloadCount sets the "download_count" field.
func (gauo *GithubAssetUpdateOne) SetDownloadCount(i int64) *GithubAssetUpdateOne {
	gauo.mutation.ResetDownloadCount()
	gauo.mutation.SetDownloadCount(i)
	return gauo
}

// SetNillableDownloadCount sets the "download_count" field if the given value is not nil.
func (gauo *GithubAssetUpdateOne) SetNillableDownloadCount(i *int64) *GithubAssetUpdateOne {
	if i != nil {
		gauo.SetDownloadCount(*i)
	}
	return gauo
}

// AddDownloadCount adds i to the "download_count" field.
func (gauo *GithubAssetUpdateOne) AddDownloadCount(i int64) *GithubAssetUpdateOne {
	gauo.mutation.AddDownloadCount(i)
	return gauo
}

// SetCreatedAt sets the "created_at" field.
func (gauo *GithubAssetUpdateOne) SetCreatedAt(t time.Time) *GithubAssetUpdateOne {
	gauo.mutation.SetCreatedAt(t)
	return gauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gauo *GithubAssetUpdateOne) SetNillableCreatedAt(t *time.Time) *GithubAssetUpdateOne {
	if t != nil {
		gauo.SetCreatedAt(*t)
	}
	return gauo
}

// SetUpdatedAt sets the "updated_at" field.
func (gauo *GithubAssetUpdateOne) SetUpdatedAt(t time.Time) *GithubAssetUpdateOne {
	gauo.mutation.SetUpdatedAt(t)
	return gauo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gauo *GithubAssetUpdateOne) SetNillableUpdatedAt(t *time.Time) *GithubAssetUpdateOne {
	if t != nil {
		gauo.SetUpdatedAt(*t)
	}
	return gauo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gauo *GithubAssetUpdateOne) ClearUpdatedAt() *GithubAssetUpdateOne {
	gauo.mutation.ClearUpdatedAt()
	return gauo
}

// SetUploader sets the "uploader" field.
func (gauo *GithubAssetUpdateOne) SetUploader(gi *github.User) *GithubAssetUpdateOne {
	gauo.mutation.SetUploader(gi)
	return gauo
}

// SetReleaseID sets the "release" edge to the GithubRelease entity by ID.
func (gauo *GithubAssetUpdateOne) SetReleaseID(id int) *GithubAssetUpdateOne {
	gauo.mutation.SetReleaseID(id)
	return gauo
}

// SetRelease sets the "release" edge to the GithubRelease entity.
func (gauo *GithubAssetUpdateOne) SetRelease(g *GithubRelease) *GithubAssetUpdateOne {
	return gauo.SetReleaseID(g.ID)
}

// Mutation returns the GithubAssetMutation object of the builder.
func (gauo *GithubAssetUpdateOne) Mutation() *GithubAssetMutation {
	return gauo.mutation
}

// ClearRelease clears the "release" edge to the GithubRelease entity.
func (gauo *GithubAssetUpdateOne) ClearRelease() *GithubAssetUpdateOne {
	gauo.mutation.ClearRelease()
	return gauo
}

// Where appends a list predicates to the GithubAssetUpdate builder.
func (gauo *GithubAssetUpdateOne) Where(ps ...predicate.GithubAsset) *GithubAssetUpdateOne {
	gauo.mutation.Where(ps...)
	return gauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gauo *GithubAssetUpdateOne) Select(field string, fields ...string) *GithubAssetUpdateOne {
	gauo.fields = append([]string{field}, fields...)
	return gauo
}

// Save executes the query and returns the updated GithubAsset entity.
func (gauo *GithubAssetUpdateOne) Save(ctx context.Context) (*GithubAsset, error) {
	return withHooks(ctx, gauo.sqlSave, gauo.mutation, gauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gauo *GithubAssetUpdateOne) SaveX(ctx context.Context) *GithubAsset {
	node, err := gauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gauo *GithubAssetUpdateOne) Exec(ctx context.Context) error {
	_, err := gauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gauo *GithubAssetUpdateOne) ExecX(ctx context.Context) {
	if err := gauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gauo *GithubAssetUpdateOne) check() error {
	if v, ok := gauo.mutation.AssetID(); ok {
		if err := githubasset.AssetIDValidator(v); err != nil {
			return &ValidationError{Name: "asset_id", err: fmt.Errorf(`ent: validator failed for field "GithubAsset.asset_id": %w`, err)}
		}
	}
	if v, ok := gauo.mutation.BrowserDownloadURL(); ok {
		if err := githubasset.BrowserDownloadURLValidator(v); err != nil {
			return &ValidationError{Name: "browser_download_url", err: fmt.Errorf(`ent: validator failed for field "GithubAsset.browser_download_url": %w`, err)}
		}
	}
	if v, ok := gauo.mutation.Name(); ok {
		if err := githubasset.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "GithubAsset.name": %w`, err)}
		}
	}
	if _, ok := gauo.mutation.ReleaseID(); gauo.mutation.ReleaseCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GithubAsset.release"`)
	}
	return nil
}

func (gauo *GithubAssetUpdateOne) sqlSave(ctx context.Context) (_node *GithubAsset, err error) {
	if err := gauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(githubasset.Table, githubasset.Columns, sqlgraph.NewFieldSpec(githubasset.FieldID, field.TypeInt))
	id, ok := gauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GithubAsset.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githubasset.FieldID)
		for _, f := range fields {
			if !githubasset.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != githubasset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gauo.mutation.AssetID(); ok {
		_spec.SetField(githubasset.FieldAssetID, field.TypeInt64, value)
	}
	if value, ok := gauo.mutation.AddedAssetID(); ok {
		_spec.AddField(githubasset.FieldAssetID, field.TypeInt64, value)
	}
	if value, ok := gauo.mutation.BrowserDownloadURL(); ok {
		_spec.SetField(githubasset.FieldBrowserDownloadURL, field.TypeString, value)
	}
	if value, ok := gauo.mutation.Name(); ok {
		_spec.SetField(githubasset.FieldName, field.TypeString, value)
	}
	if value, ok := gauo.mutation.Label(); ok {
		_spec.SetField(githubasset.FieldLabel, field.TypeString, value)
	}
	if gauo.mutation.LabelCleared() {
		_spec.ClearField(githubasset.FieldLabel, field.TypeString)
	}
	if value, ok := gauo.mutation.State(); ok {
		_spec.SetField(githubasset.FieldState, field.TypeString, value)
	}
	if gauo.mutation.StateCleared() {
		_spec.ClearField(githubasset.FieldState, field.TypeString)
	}
	if value, ok := gauo.mutation.ContentType(); ok {
		_spec.SetField(githubasset.FieldContentType, field.TypeString, value)
	}
	if value, ok := gauo.mutation.Size(); ok {
		_spec.SetField(githubasset.FieldSize, field.TypeInt64, value)
	}
	if value, ok := gauo.mutation.AddedSize(); ok {
		_spec.AddField(githubasset.FieldSize, field.TypeInt64, value)
	}
	if value, ok := gauo.mutation.DownloadCount(); ok {
		_spec.SetField(githubasset.FieldDownloadCount, field.TypeInt64, value)
	}
	if value, ok := gauo.mutation.AddedDownloadCount(); ok {
		_spec.AddField(githubasset.FieldDownloadCount, field.TypeInt64, value)
	}
	if value, ok := gauo.mutation.CreatedAt(); ok {
		_spec.SetField(githubasset.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := gauo.mutation.UpdatedAt(); ok {
		_spec.SetField(githubasset.FieldUpdatedAt, field.TypeTime, value)
	}
	if gauo.mutation.UpdatedAtCleared() {
		_spec.ClearField(githubasset.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := gauo.mutation.Uploader(); ok {
		_spec.SetField(githubasset.FieldUploader, field.TypeJSON, value)
	}
	if gauo.mutation.ReleaseCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gauo.mutation.ReleaseIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GithubAsset{config: gauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubasset.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gauo.mutation.done = true
	return _node, nil
}
