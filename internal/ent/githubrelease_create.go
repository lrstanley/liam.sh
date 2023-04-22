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
	"github.com/lrstanley/liam.sh/internal/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/ent/githubrepository"
)

// GithubReleaseCreate is the builder for creating a GithubRelease entity.
type GithubReleaseCreate struct {
	config
	mutation *GithubReleaseMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetReleaseID sets the "release_id" field.
func (grc *GithubReleaseCreate) SetReleaseID(i int64) *GithubReleaseCreate {
	grc.mutation.SetReleaseID(i)
	return grc
}

// SetHTMLURL sets the "html_url" field.
func (grc *GithubReleaseCreate) SetHTMLURL(s string) *GithubReleaseCreate {
	grc.mutation.SetHTMLURL(s)
	return grc
}

// SetTagName sets the "tag_name" field.
func (grc *GithubReleaseCreate) SetTagName(s string) *GithubReleaseCreate {
	grc.mutation.SetTagName(s)
	return grc
}

// SetTargetCommitish sets the "target_commitish" field.
func (grc *GithubReleaseCreate) SetTargetCommitish(s string) *GithubReleaseCreate {
	grc.mutation.SetTargetCommitish(s)
	return grc
}

// SetName sets the "name" field.
func (grc *GithubReleaseCreate) SetName(s string) *GithubReleaseCreate {
	grc.mutation.SetName(s)
	return grc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (grc *GithubReleaseCreate) SetNillableName(s *string) *GithubReleaseCreate {
	if s != nil {
		grc.SetName(*s)
	}
	return grc
}

// SetDraft sets the "draft" field.
func (grc *GithubReleaseCreate) SetDraft(b bool) *GithubReleaseCreate {
	grc.mutation.SetDraft(b)
	return grc
}

// SetPrerelease sets the "prerelease" field.
func (grc *GithubReleaseCreate) SetPrerelease(b bool) *GithubReleaseCreate {
	grc.mutation.SetPrerelease(b)
	return grc
}

// SetCreatedAt sets the "created_at" field.
func (grc *GithubReleaseCreate) SetCreatedAt(t time.Time) *GithubReleaseCreate {
	grc.mutation.SetCreatedAt(t)
	return grc
}

// SetPublishedAt sets the "published_at" field.
func (grc *GithubReleaseCreate) SetPublishedAt(t time.Time) *GithubReleaseCreate {
	grc.mutation.SetPublishedAt(t)
	return grc
}

// SetAuthor sets the "author" field.
func (grc *GithubReleaseCreate) SetAuthor(gi *github.User) *GithubReleaseCreate {
	grc.mutation.SetAuthor(gi)
	return grc
}

// SetRepositoryID sets the "repository" edge to the GithubRepository entity by ID.
func (grc *GithubReleaseCreate) SetRepositoryID(id int) *GithubReleaseCreate {
	grc.mutation.SetRepositoryID(id)
	return grc
}

// SetRepository sets the "repository" edge to the GithubRepository entity.
func (grc *GithubReleaseCreate) SetRepository(g *GithubRepository) *GithubReleaseCreate {
	return grc.SetRepositoryID(g.ID)
}

// AddAssetIDs adds the "assets" edge to the GithubAsset entity by IDs.
func (grc *GithubReleaseCreate) AddAssetIDs(ids ...int) *GithubReleaseCreate {
	grc.mutation.AddAssetIDs(ids...)
	return grc
}

// AddAssets adds the "assets" edges to the GithubAsset entity.
func (grc *GithubReleaseCreate) AddAssets(g ...*GithubAsset) *GithubReleaseCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return grc.AddAssetIDs(ids...)
}

// Mutation returns the GithubReleaseMutation object of the builder.
func (grc *GithubReleaseCreate) Mutation() *GithubReleaseMutation {
	return grc.mutation
}

// Save creates the GithubRelease in the database.
func (grc *GithubReleaseCreate) Save(ctx context.Context) (*GithubRelease, error) {
	return withHooks[*GithubRelease, GithubReleaseMutation](ctx, grc.sqlSave, grc.mutation, grc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (grc *GithubReleaseCreate) SaveX(ctx context.Context) *GithubRelease {
	v, err := grc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (grc *GithubReleaseCreate) Exec(ctx context.Context) error {
	_, err := grc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (grc *GithubReleaseCreate) ExecX(ctx context.Context) {
	if err := grc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (grc *GithubReleaseCreate) check() error {
	if _, ok := grc.mutation.ReleaseID(); !ok {
		return &ValidationError{Name: "release_id", err: errors.New(`ent: missing required field "GithubRelease.release_id"`)}
	}
	if v, ok := grc.mutation.ReleaseID(); ok {
		if err := githubrelease.ReleaseIDValidator(v); err != nil {
			return &ValidationError{Name: "release_id", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.release_id": %w`, err)}
		}
	}
	if _, ok := grc.mutation.HTMLURL(); !ok {
		return &ValidationError{Name: "html_url", err: errors.New(`ent: missing required field "GithubRelease.html_url"`)}
	}
	if v, ok := grc.mutation.HTMLURL(); ok {
		if err := githubrelease.HTMLURLValidator(v); err != nil {
			return &ValidationError{Name: "html_url", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.html_url": %w`, err)}
		}
	}
	if _, ok := grc.mutation.TagName(); !ok {
		return &ValidationError{Name: "tag_name", err: errors.New(`ent: missing required field "GithubRelease.tag_name"`)}
	}
	if v, ok := grc.mutation.TagName(); ok {
		if err := githubrelease.TagNameValidator(v); err != nil {
			return &ValidationError{Name: "tag_name", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.tag_name": %w`, err)}
		}
	}
	if _, ok := grc.mutation.TargetCommitish(); !ok {
		return &ValidationError{Name: "target_commitish", err: errors.New(`ent: missing required field "GithubRelease.target_commitish"`)}
	}
	if v, ok := grc.mutation.TargetCommitish(); ok {
		if err := githubrelease.TargetCommitishValidator(v); err != nil {
			return &ValidationError{Name: "target_commitish", err: fmt.Errorf(`ent: validator failed for field "GithubRelease.target_commitish": %w`, err)}
		}
	}
	if _, ok := grc.mutation.Draft(); !ok {
		return &ValidationError{Name: "draft", err: errors.New(`ent: missing required field "GithubRelease.draft"`)}
	}
	if _, ok := grc.mutation.Prerelease(); !ok {
		return &ValidationError{Name: "prerelease", err: errors.New(`ent: missing required field "GithubRelease.prerelease"`)}
	}
	if _, ok := grc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GithubRelease.created_at"`)}
	}
	if _, ok := grc.mutation.PublishedAt(); !ok {
		return &ValidationError{Name: "published_at", err: errors.New(`ent: missing required field "GithubRelease.published_at"`)}
	}
	if _, ok := grc.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required field "GithubRelease.author"`)}
	}
	if _, ok := grc.mutation.RepositoryID(); !ok {
		return &ValidationError{Name: "repository", err: errors.New(`ent: missing required edge "GithubRelease.repository"`)}
	}
	return nil
}

func (grc *GithubReleaseCreate) sqlSave(ctx context.Context) (*GithubRelease, error) {
	if err := grc.check(); err != nil {
		return nil, err
	}
	_node, _spec := grc.createSpec()
	if err := sqlgraph.CreateNode(ctx, grc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	grc.mutation.id = &_node.ID
	grc.mutation.done = true
	return _node, nil
}

func (grc *GithubReleaseCreate) createSpec() (*GithubRelease, *sqlgraph.CreateSpec) {
	var (
		_node = &GithubRelease{config: grc.config}
		_spec = sqlgraph.NewCreateSpec(githubrelease.Table, sqlgraph.NewFieldSpec(githubrelease.FieldID, field.TypeInt))
	)
	_spec.OnConflict = grc.conflict
	if value, ok := grc.mutation.ReleaseID(); ok {
		_spec.SetField(githubrelease.FieldReleaseID, field.TypeInt64, value)
		_node.ReleaseID = value
	}
	if value, ok := grc.mutation.HTMLURL(); ok {
		_spec.SetField(githubrelease.FieldHTMLURL, field.TypeString, value)
		_node.HTMLURL = value
	}
	if value, ok := grc.mutation.TagName(); ok {
		_spec.SetField(githubrelease.FieldTagName, field.TypeString, value)
		_node.TagName = value
	}
	if value, ok := grc.mutation.TargetCommitish(); ok {
		_spec.SetField(githubrelease.FieldTargetCommitish, field.TypeString, value)
		_node.TargetCommitish = value
	}
	if value, ok := grc.mutation.Name(); ok {
		_spec.SetField(githubrelease.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := grc.mutation.Draft(); ok {
		_spec.SetField(githubrelease.FieldDraft, field.TypeBool, value)
		_node.Draft = value
	}
	if value, ok := grc.mutation.Prerelease(); ok {
		_spec.SetField(githubrelease.FieldPrerelease, field.TypeBool, value)
		_node.Prerelease = value
	}
	if value, ok := grc.mutation.CreatedAt(); ok {
		_spec.SetField(githubrelease.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := grc.mutation.PublishedAt(); ok {
		_spec.SetField(githubrelease.FieldPublishedAt, field.TypeTime, value)
		_node.PublishedAt = value
	}
	if value, ok := grc.mutation.Author(); ok {
		_spec.SetField(githubrelease.FieldAuthor, field.TypeJSON, value)
		_node.Author = value
	}
	if nodes := grc.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   githubrelease.RepositoryTable,
			Columns: []string{githubrelease.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.github_repository_releases = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := grc.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrelease.AssetsTable,
			Columns: []string{githubrelease.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubasset.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GithubRelease.Create().
//		SetReleaseID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GithubReleaseUpsert) {
//			SetReleaseID(v+v).
//		}).
//		Exec(ctx)
func (grc *GithubReleaseCreate) OnConflict(opts ...sql.ConflictOption) *GithubReleaseUpsertOne {
	grc.conflict = opts
	return &GithubReleaseUpsertOne{
		create: grc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GithubRelease.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (grc *GithubReleaseCreate) OnConflictColumns(columns ...string) *GithubReleaseUpsertOne {
	grc.conflict = append(grc.conflict, sql.ConflictColumns(columns...))
	return &GithubReleaseUpsertOne{
		create: grc,
	}
}

type (
	// GithubReleaseUpsertOne is the builder for "upsert"-ing
	//  one GithubRelease node.
	GithubReleaseUpsertOne struct {
		create *GithubReleaseCreate
	}

	// GithubReleaseUpsert is the "OnConflict" setter.
	GithubReleaseUpsert struct {
		*sql.UpdateSet
	}
)

// SetReleaseID sets the "release_id" field.
func (u *GithubReleaseUpsert) SetReleaseID(v int64) *GithubReleaseUpsert {
	u.Set(githubrelease.FieldReleaseID, v)
	return u
}

// UpdateReleaseID sets the "release_id" field to the value that was provided on create.
func (u *GithubReleaseUpsert) UpdateReleaseID() *GithubReleaseUpsert {
	u.SetExcluded(githubrelease.FieldReleaseID)
	return u
}

// AddReleaseID adds v to the "release_id" field.
func (u *GithubReleaseUpsert) AddReleaseID(v int64) *GithubReleaseUpsert {
	u.Add(githubrelease.FieldReleaseID, v)
	return u
}

// SetHTMLURL sets the "html_url" field.
func (u *GithubReleaseUpsert) SetHTMLURL(v string) *GithubReleaseUpsert {
	u.Set(githubrelease.FieldHTMLURL, v)
	return u
}

// UpdateHTMLURL sets the "html_url" field to the value that was provided on create.
func (u *GithubReleaseUpsert) UpdateHTMLURL() *GithubReleaseUpsert {
	u.SetExcluded(githubrelease.FieldHTMLURL)
	return u
}

// SetTagName sets the "tag_name" field.
func (u *GithubReleaseUpsert) SetTagName(v string) *GithubReleaseUpsert {
	u.Set(githubrelease.FieldTagName, v)
	return u
}

// UpdateTagName sets the "tag_name" field to the value that was provided on create.
func (u *GithubReleaseUpsert) UpdateTagName() *GithubReleaseUpsert {
	u.SetExcluded(githubrelease.FieldTagName)
	return u
}

// SetTargetCommitish sets the "target_commitish" field.
func (u *GithubReleaseUpsert) SetTargetCommitish(v string) *GithubReleaseUpsert {
	u.Set(githubrelease.FieldTargetCommitish, v)
	return u
}

// UpdateTargetCommitish sets the "target_commitish" field to the value that was provided on create.
func (u *GithubReleaseUpsert) UpdateTargetCommitish() *GithubReleaseUpsert {
	u.SetExcluded(githubrelease.FieldTargetCommitish)
	return u
}

// SetName sets the "name" field.
func (u *GithubReleaseUpsert) SetName(v string) *GithubReleaseUpsert {
	u.Set(githubrelease.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GithubReleaseUpsert) UpdateName() *GithubReleaseUpsert {
	u.SetExcluded(githubrelease.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *GithubReleaseUpsert) ClearName() *GithubReleaseUpsert {
	u.SetNull(githubrelease.FieldName)
	return u
}

// SetDraft sets the "draft" field.
func (u *GithubReleaseUpsert) SetDraft(v bool) *GithubReleaseUpsert {
	u.Set(githubrelease.FieldDraft, v)
	return u
}

// UpdateDraft sets the "draft" field to the value that was provided on create.
func (u *GithubReleaseUpsert) UpdateDraft() *GithubReleaseUpsert {
	u.SetExcluded(githubrelease.FieldDraft)
	return u
}

// SetPrerelease sets the "prerelease" field.
func (u *GithubReleaseUpsert) SetPrerelease(v bool) *GithubReleaseUpsert {
	u.Set(githubrelease.FieldPrerelease, v)
	return u
}

// UpdatePrerelease sets the "prerelease" field to the value that was provided on create.
func (u *GithubReleaseUpsert) UpdatePrerelease() *GithubReleaseUpsert {
	u.SetExcluded(githubrelease.FieldPrerelease)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubReleaseUpsert) SetCreatedAt(v time.Time) *GithubReleaseUpsert {
	u.Set(githubrelease.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubReleaseUpsert) UpdateCreatedAt() *GithubReleaseUpsert {
	u.SetExcluded(githubrelease.FieldCreatedAt)
	return u
}

// SetPublishedAt sets the "published_at" field.
func (u *GithubReleaseUpsert) SetPublishedAt(v time.Time) *GithubReleaseUpsert {
	u.Set(githubrelease.FieldPublishedAt, v)
	return u
}

// UpdatePublishedAt sets the "published_at" field to the value that was provided on create.
func (u *GithubReleaseUpsert) UpdatePublishedAt() *GithubReleaseUpsert {
	u.SetExcluded(githubrelease.FieldPublishedAt)
	return u
}

// SetAuthor sets the "author" field.
func (u *GithubReleaseUpsert) SetAuthor(v *github.User) *GithubReleaseUpsert {
	u.Set(githubrelease.FieldAuthor, v)
	return u
}

// UpdateAuthor sets the "author" field to the value that was provided on create.
func (u *GithubReleaseUpsert) UpdateAuthor() *GithubReleaseUpsert {
	u.SetExcluded(githubrelease.FieldAuthor)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.GithubRelease.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GithubReleaseUpsertOne) UpdateNewValues() *GithubReleaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GithubRelease.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GithubReleaseUpsertOne) Ignore() *GithubReleaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GithubReleaseUpsertOne) DoNothing() *GithubReleaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GithubReleaseCreate.OnConflict
// documentation for more info.
func (u *GithubReleaseUpsertOne) Update(set func(*GithubReleaseUpsert)) *GithubReleaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GithubReleaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetReleaseID sets the "release_id" field.
func (u *GithubReleaseUpsertOne) SetReleaseID(v int64) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetReleaseID(v)
	})
}

// AddReleaseID adds v to the "release_id" field.
func (u *GithubReleaseUpsertOne) AddReleaseID(v int64) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.AddReleaseID(v)
	})
}

// UpdateReleaseID sets the "release_id" field to the value that was provided on create.
func (u *GithubReleaseUpsertOne) UpdateReleaseID() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateReleaseID()
	})
}

// SetHTMLURL sets the "html_url" field.
func (u *GithubReleaseUpsertOne) SetHTMLURL(v string) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetHTMLURL(v)
	})
}

// UpdateHTMLURL sets the "html_url" field to the value that was provided on create.
func (u *GithubReleaseUpsertOne) UpdateHTMLURL() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateHTMLURL()
	})
}

// SetTagName sets the "tag_name" field.
func (u *GithubReleaseUpsertOne) SetTagName(v string) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetTagName(v)
	})
}

// UpdateTagName sets the "tag_name" field to the value that was provided on create.
func (u *GithubReleaseUpsertOne) UpdateTagName() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateTagName()
	})
}

// SetTargetCommitish sets the "target_commitish" field.
func (u *GithubReleaseUpsertOne) SetTargetCommitish(v string) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetTargetCommitish(v)
	})
}

// UpdateTargetCommitish sets the "target_commitish" field to the value that was provided on create.
func (u *GithubReleaseUpsertOne) UpdateTargetCommitish() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateTargetCommitish()
	})
}

// SetName sets the "name" field.
func (u *GithubReleaseUpsertOne) SetName(v string) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GithubReleaseUpsertOne) UpdateName() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *GithubReleaseUpsertOne) ClearName() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.ClearName()
	})
}

// SetDraft sets the "draft" field.
func (u *GithubReleaseUpsertOne) SetDraft(v bool) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetDraft(v)
	})
}

// UpdateDraft sets the "draft" field to the value that was provided on create.
func (u *GithubReleaseUpsertOne) UpdateDraft() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateDraft()
	})
}

// SetPrerelease sets the "prerelease" field.
func (u *GithubReleaseUpsertOne) SetPrerelease(v bool) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetPrerelease(v)
	})
}

// UpdatePrerelease sets the "prerelease" field to the value that was provided on create.
func (u *GithubReleaseUpsertOne) UpdatePrerelease() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdatePrerelease()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubReleaseUpsertOne) SetCreatedAt(v time.Time) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubReleaseUpsertOne) UpdateCreatedAt() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetPublishedAt sets the "published_at" field.
func (u *GithubReleaseUpsertOne) SetPublishedAt(v time.Time) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetPublishedAt(v)
	})
}

// UpdatePublishedAt sets the "published_at" field to the value that was provided on create.
func (u *GithubReleaseUpsertOne) UpdatePublishedAt() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdatePublishedAt()
	})
}

// SetAuthor sets the "author" field.
func (u *GithubReleaseUpsertOne) SetAuthor(v *github.User) *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetAuthor(v)
	})
}

// UpdateAuthor sets the "author" field to the value that was provided on create.
func (u *GithubReleaseUpsertOne) UpdateAuthor() *GithubReleaseUpsertOne {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateAuthor()
	})
}

// Exec executes the query.
func (u *GithubReleaseUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GithubReleaseCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GithubReleaseUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GithubReleaseUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GithubReleaseUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GithubReleaseCreateBulk is the builder for creating many GithubRelease entities in bulk.
type GithubReleaseCreateBulk struct {
	config
	builders []*GithubReleaseCreate
	conflict []sql.ConflictOption
}

// Save creates the GithubRelease entities in the database.
func (grcb *GithubReleaseCreateBulk) Save(ctx context.Context) ([]*GithubRelease, error) {
	specs := make([]*sqlgraph.CreateSpec, len(grcb.builders))
	nodes := make([]*GithubRelease, len(grcb.builders))
	mutators := make([]Mutator, len(grcb.builders))
	for i := range grcb.builders {
		func(i int, root context.Context) {
			builder := grcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GithubReleaseMutation)
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
					_, err = mutators[i+1].Mutate(root, grcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = grcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, grcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, grcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (grcb *GithubReleaseCreateBulk) SaveX(ctx context.Context) []*GithubRelease {
	v, err := grcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (grcb *GithubReleaseCreateBulk) Exec(ctx context.Context) error {
	_, err := grcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (grcb *GithubReleaseCreateBulk) ExecX(ctx context.Context) {
	if err := grcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GithubRelease.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GithubReleaseUpsert) {
//			SetReleaseID(v+v).
//		}).
//		Exec(ctx)
func (grcb *GithubReleaseCreateBulk) OnConflict(opts ...sql.ConflictOption) *GithubReleaseUpsertBulk {
	grcb.conflict = opts
	return &GithubReleaseUpsertBulk{
		create: grcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GithubRelease.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (grcb *GithubReleaseCreateBulk) OnConflictColumns(columns ...string) *GithubReleaseUpsertBulk {
	grcb.conflict = append(grcb.conflict, sql.ConflictColumns(columns...))
	return &GithubReleaseUpsertBulk{
		create: grcb,
	}
}

// GithubReleaseUpsertBulk is the builder for "upsert"-ing
// a bulk of GithubRelease nodes.
type GithubReleaseUpsertBulk struct {
	create *GithubReleaseCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GithubRelease.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GithubReleaseUpsertBulk) UpdateNewValues() *GithubReleaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GithubRelease.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GithubReleaseUpsertBulk) Ignore() *GithubReleaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GithubReleaseUpsertBulk) DoNothing() *GithubReleaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GithubReleaseCreateBulk.OnConflict
// documentation for more info.
func (u *GithubReleaseUpsertBulk) Update(set func(*GithubReleaseUpsert)) *GithubReleaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GithubReleaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetReleaseID sets the "release_id" field.
func (u *GithubReleaseUpsertBulk) SetReleaseID(v int64) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetReleaseID(v)
	})
}

// AddReleaseID adds v to the "release_id" field.
func (u *GithubReleaseUpsertBulk) AddReleaseID(v int64) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.AddReleaseID(v)
	})
}

// UpdateReleaseID sets the "release_id" field to the value that was provided on create.
func (u *GithubReleaseUpsertBulk) UpdateReleaseID() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateReleaseID()
	})
}

// SetHTMLURL sets the "html_url" field.
func (u *GithubReleaseUpsertBulk) SetHTMLURL(v string) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetHTMLURL(v)
	})
}

// UpdateHTMLURL sets the "html_url" field to the value that was provided on create.
func (u *GithubReleaseUpsertBulk) UpdateHTMLURL() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateHTMLURL()
	})
}

// SetTagName sets the "tag_name" field.
func (u *GithubReleaseUpsertBulk) SetTagName(v string) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetTagName(v)
	})
}

// UpdateTagName sets the "tag_name" field to the value that was provided on create.
func (u *GithubReleaseUpsertBulk) UpdateTagName() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateTagName()
	})
}

// SetTargetCommitish sets the "target_commitish" field.
func (u *GithubReleaseUpsertBulk) SetTargetCommitish(v string) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetTargetCommitish(v)
	})
}

// UpdateTargetCommitish sets the "target_commitish" field to the value that was provided on create.
func (u *GithubReleaseUpsertBulk) UpdateTargetCommitish() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateTargetCommitish()
	})
}

// SetName sets the "name" field.
func (u *GithubReleaseUpsertBulk) SetName(v string) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GithubReleaseUpsertBulk) UpdateName() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *GithubReleaseUpsertBulk) ClearName() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.ClearName()
	})
}

// SetDraft sets the "draft" field.
func (u *GithubReleaseUpsertBulk) SetDraft(v bool) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetDraft(v)
	})
}

// UpdateDraft sets the "draft" field to the value that was provided on create.
func (u *GithubReleaseUpsertBulk) UpdateDraft() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateDraft()
	})
}

// SetPrerelease sets the "prerelease" field.
func (u *GithubReleaseUpsertBulk) SetPrerelease(v bool) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetPrerelease(v)
	})
}

// UpdatePrerelease sets the "prerelease" field to the value that was provided on create.
func (u *GithubReleaseUpsertBulk) UpdatePrerelease() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdatePrerelease()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubReleaseUpsertBulk) SetCreatedAt(v time.Time) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubReleaseUpsertBulk) UpdateCreatedAt() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetPublishedAt sets the "published_at" field.
func (u *GithubReleaseUpsertBulk) SetPublishedAt(v time.Time) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetPublishedAt(v)
	})
}

// UpdatePublishedAt sets the "published_at" field to the value that was provided on create.
func (u *GithubReleaseUpsertBulk) UpdatePublishedAt() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdatePublishedAt()
	})
}

// SetAuthor sets the "author" field.
func (u *GithubReleaseUpsertBulk) SetAuthor(v *github.User) *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.SetAuthor(v)
	})
}

// UpdateAuthor sets the "author" field to the value that was provided on create.
func (u *GithubReleaseUpsertBulk) UpdateAuthor() *GithubReleaseUpsertBulk {
	return u.Update(func(s *GithubReleaseUpsert) {
		s.UpdateAuthor()
	})
}

// Exec executes the query.
func (u *GithubReleaseUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GithubReleaseCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GithubReleaseCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GithubReleaseUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
