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
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/database/ent/label"
	"github.com/lrstanley/liam.sh/internal/database/ent/predicate"
)

// GithubRepositoryUpdate is the builder for updating GithubRepository entities.
type GithubRepositoryUpdate struct {
	config
	hooks    []Hook
	mutation *GithubRepositoryMutation
}

// Where appends a list predicates to the GithubRepositoryUpdate builder.
func (gru *GithubRepositoryUpdate) Where(ps ...predicate.GithubRepository) *GithubRepositoryUpdate {
	gru.mutation.Where(ps...)
	return gru
}

// SetRepoID sets the "repo_id" field.
func (gru *GithubRepositoryUpdate) SetRepoID(i int64) *GithubRepositoryUpdate {
	gru.mutation.ResetRepoID()
	gru.mutation.SetRepoID(i)
	return gru
}

// AddRepoID adds i to the "repo_id" field.
func (gru *GithubRepositoryUpdate) AddRepoID(i int64) *GithubRepositoryUpdate {
	gru.mutation.AddRepoID(i)
	return gru
}

// SetName sets the "name" field.
func (gru *GithubRepositoryUpdate) SetName(s string) *GithubRepositoryUpdate {
	gru.mutation.SetName(s)
	return gru
}

// SetFullName sets the "full_name" field.
func (gru *GithubRepositoryUpdate) SetFullName(s string) *GithubRepositoryUpdate {
	gru.mutation.SetFullName(s)
	return gru
}

// SetOwnerLogin sets the "owner_login" field.
func (gru *GithubRepositoryUpdate) SetOwnerLogin(s string) *GithubRepositoryUpdate {
	gru.mutation.SetOwnerLogin(s)
	return gru
}

// SetOwner sets the "owner" field.
func (gru *GithubRepositoryUpdate) SetOwner(gi *github.User) *GithubRepositoryUpdate {
	gru.mutation.SetOwner(gi)
	return gru
}

// SetPublic sets the "public" field.
func (gru *GithubRepositoryUpdate) SetPublic(b bool) *GithubRepositoryUpdate {
	gru.mutation.SetPublic(b)
	return gru
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (gru *GithubRepositoryUpdate) SetNillablePublic(b *bool) *GithubRepositoryUpdate {
	if b != nil {
		gru.SetPublic(*b)
	}
	return gru
}

// SetHTMLURL sets the "html_url" field.
func (gru *GithubRepositoryUpdate) SetHTMLURL(s string) *GithubRepositoryUpdate {
	gru.mutation.SetHTMLURL(s)
	return gru
}

// SetDescription sets the "description" field.
func (gru *GithubRepositoryUpdate) SetDescription(s string) *GithubRepositoryUpdate {
	gru.mutation.SetDescription(s)
	return gru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gru *GithubRepositoryUpdate) SetNillableDescription(s *string) *GithubRepositoryUpdate {
	if s != nil {
		gru.SetDescription(*s)
	}
	return gru
}

// ClearDescription clears the value of the "description" field.
func (gru *GithubRepositoryUpdate) ClearDescription() *GithubRepositoryUpdate {
	gru.mutation.ClearDescription()
	return gru
}

// SetFork sets the "fork" field.
func (gru *GithubRepositoryUpdate) SetFork(b bool) *GithubRepositoryUpdate {
	gru.mutation.SetFork(b)
	return gru
}

// SetNillableFork sets the "fork" field if the given value is not nil.
func (gru *GithubRepositoryUpdate) SetNillableFork(b *bool) *GithubRepositoryUpdate {
	if b != nil {
		gru.SetFork(*b)
	}
	return gru
}

// SetHomepage sets the "homepage" field.
func (gru *GithubRepositoryUpdate) SetHomepage(s string) *GithubRepositoryUpdate {
	gru.mutation.SetHomepage(s)
	return gru
}

// SetNillableHomepage sets the "homepage" field if the given value is not nil.
func (gru *GithubRepositoryUpdate) SetNillableHomepage(s *string) *GithubRepositoryUpdate {
	if s != nil {
		gru.SetHomepage(*s)
	}
	return gru
}

// ClearHomepage clears the value of the "homepage" field.
func (gru *GithubRepositoryUpdate) ClearHomepage() *GithubRepositoryUpdate {
	gru.mutation.ClearHomepage()
	return gru
}

// SetStarCount sets the "star_count" field.
func (gru *GithubRepositoryUpdate) SetStarCount(i int) *GithubRepositoryUpdate {
	gru.mutation.ResetStarCount()
	gru.mutation.SetStarCount(i)
	return gru
}

// SetNillableStarCount sets the "star_count" field if the given value is not nil.
func (gru *GithubRepositoryUpdate) SetNillableStarCount(i *int) *GithubRepositoryUpdate {
	if i != nil {
		gru.SetStarCount(*i)
	}
	return gru
}

// AddStarCount adds i to the "star_count" field.
func (gru *GithubRepositoryUpdate) AddStarCount(i int) *GithubRepositoryUpdate {
	gru.mutation.AddStarCount(i)
	return gru
}

// SetDefaultBranch sets the "default_branch" field.
func (gru *GithubRepositoryUpdate) SetDefaultBranch(s string) *GithubRepositoryUpdate {
	gru.mutation.SetDefaultBranch(s)
	return gru
}

// SetIsTemplate sets the "is_template" field.
func (gru *GithubRepositoryUpdate) SetIsTemplate(b bool) *GithubRepositoryUpdate {
	gru.mutation.SetIsTemplate(b)
	return gru
}

// SetNillableIsTemplate sets the "is_template" field if the given value is not nil.
func (gru *GithubRepositoryUpdate) SetNillableIsTemplate(b *bool) *GithubRepositoryUpdate {
	if b != nil {
		gru.SetIsTemplate(*b)
	}
	return gru
}

// SetHasIssues sets the "has_issues" field.
func (gru *GithubRepositoryUpdate) SetHasIssues(b bool) *GithubRepositoryUpdate {
	gru.mutation.SetHasIssues(b)
	return gru
}

// SetNillableHasIssues sets the "has_issues" field if the given value is not nil.
func (gru *GithubRepositoryUpdate) SetNillableHasIssues(b *bool) *GithubRepositoryUpdate {
	if b != nil {
		gru.SetHasIssues(*b)
	}
	return gru
}

// SetArchived sets the "archived" field.
func (gru *GithubRepositoryUpdate) SetArchived(b bool) *GithubRepositoryUpdate {
	gru.mutation.SetArchived(b)
	return gru
}

// SetNillableArchived sets the "archived" field if the given value is not nil.
func (gru *GithubRepositoryUpdate) SetNillableArchived(b *bool) *GithubRepositoryUpdate {
	if b != nil {
		gru.SetArchived(*b)
	}
	return gru
}

// SetPushedAt sets the "pushed_at" field.
func (gru *GithubRepositoryUpdate) SetPushedAt(t time.Time) *GithubRepositoryUpdate {
	gru.mutation.SetPushedAt(t)
	return gru
}

// SetNillablePushedAt sets the "pushed_at" field if the given value is not nil.
func (gru *GithubRepositoryUpdate) SetNillablePushedAt(t *time.Time) *GithubRepositoryUpdate {
	if t != nil {
		gru.SetPushedAt(*t)
	}
	return gru
}

// ClearPushedAt clears the value of the "pushed_at" field.
func (gru *GithubRepositoryUpdate) ClearPushedAt() *GithubRepositoryUpdate {
	gru.mutation.ClearPushedAt()
	return gru
}

// SetCreatedAt sets the "created_at" field.
func (gru *GithubRepositoryUpdate) SetCreatedAt(t time.Time) *GithubRepositoryUpdate {
	gru.mutation.SetCreatedAt(t)
	return gru
}

// SetUpdatedAt sets the "updated_at" field.
func (gru *GithubRepositoryUpdate) SetUpdatedAt(t time.Time) *GithubRepositoryUpdate {
	gru.mutation.SetUpdatedAt(t)
	return gru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gru *GithubRepositoryUpdate) SetNillableUpdatedAt(t *time.Time) *GithubRepositoryUpdate {
	if t != nil {
		gru.SetUpdatedAt(*t)
	}
	return gru
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gru *GithubRepositoryUpdate) ClearUpdatedAt() *GithubRepositoryUpdate {
	gru.mutation.ClearUpdatedAt()
	return gru
}

// SetLicense sets the "license" field.
func (gru *GithubRepositoryUpdate) SetLicense(gi *github.License) *GithubRepositoryUpdate {
	gru.mutation.SetLicense(gi)
	return gru
}

// ClearLicense clears the value of the "license" field.
func (gru *GithubRepositoryUpdate) ClearLicense() *GithubRepositoryUpdate {
	gru.mutation.ClearLicense()
	return gru
}

// AddLabelIDs adds the "labels" edge to the Label entity by IDs.
func (gru *GithubRepositoryUpdate) AddLabelIDs(ids ...int) *GithubRepositoryUpdate {
	gru.mutation.AddLabelIDs(ids...)
	return gru
}

// AddLabels adds the "labels" edges to the Label entity.
func (gru *GithubRepositoryUpdate) AddLabels(l ...*Label) *GithubRepositoryUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return gru.AddLabelIDs(ids...)
}

// AddReleaseIDs adds the "releases" edge to the GithubRelease entity by IDs.
func (gru *GithubRepositoryUpdate) AddReleaseIDs(ids ...int) *GithubRepositoryUpdate {
	gru.mutation.AddReleaseIDs(ids...)
	return gru
}

// AddReleases adds the "releases" edges to the GithubRelease entity.
func (gru *GithubRepositoryUpdate) AddReleases(g ...*GithubRelease) *GithubRepositoryUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gru.AddReleaseIDs(ids...)
}

// Mutation returns the GithubRepositoryMutation object of the builder.
func (gru *GithubRepositoryUpdate) Mutation() *GithubRepositoryMutation {
	return gru.mutation
}

// ClearLabels clears all "labels" edges to the Label entity.
func (gru *GithubRepositoryUpdate) ClearLabels() *GithubRepositoryUpdate {
	gru.mutation.ClearLabels()
	return gru
}

// RemoveLabelIDs removes the "labels" edge to Label entities by IDs.
func (gru *GithubRepositoryUpdate) RemoveLabelIDs(ids ...int) *GithubRepositoryUpdate {
	gru.mutation.RemoveLabelIDs(ids...)
	return gru
}

// RemoveLabels removes "labels" edges to Label entities.
func (gru *GithubRepositoryUpdate) RemoveLabels(l ...*Label) *GithubRepositoryUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return gru.RemoveLabelIDs(ids...)
}

// ClearReleases clears all "releases" edges to the GithubRelease entity.
func (gru *GithubRepositoryUpdate) ClearReleases() *GithubRepositoryUpdate {
	gru.mutation.ClearReleases()
	return gru
}

// RemoveReleaseIDs removes the "releases" edge to GithubRelease entities by IDs.
func (gru *GithubRepositoryUpdate) RemoveReleaseIDs(ids ...int) *GithubRepositoryUpdate {
	gru.mutation.RemoveReleaseIDs(ids...)
	return gru
}

// RemoveReleases removes "releases" edges to GithubRelease entities.
func (gru *GithubRepositoryUpdate) RemoveReleases(g ...*GithubRelease) *GithubRepositoryUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gru.RemoveReleaseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gru *GithubRepositoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gru.sqlSave, gru.mutation, gru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gru *GithubRepositoryUpdate) SaveX(ctx context.Context) int {
	affected, err := gru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gru *GithubRepositoryUpdate) Exec(ctx context.Context) error {
	_, err := gru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gru *GithubRepositoryUpdate) ExecX(ctx context.Context) {
	if err := gru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gru *GithubRepositoryUpdate) check() error {
	if v, ok := gru.mutation.RepoID(); ok {
		if err := githubrepository.RepoIDValidator(v); err != nil {
			return &ValidationError{Name: "repo_id", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.repo_id": %w`, err)}
		}
	}
	if v, ok := gru.mutation.Name(); ok {
		if err := githubrepository.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.name": %w`, err)}
		}
	}
	if v, ok := gru.mutation.FullName(); ok {
		if err := githubrepository.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.full_name": %w`, err)}
		}
	}
	if v, ok := gru.mutation.OwnerLogin(); ok {
		if err := githubrepository.OwnerLoginValidator(v); err != nil {
			return &ValidationError{Name: "owner_login", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.owner_login": %w`, err)}
		}
	}
	if v, ok := gru.mutation.HTMLURL(); ok {
		if err := githubrepository.HTMLURLValidator(v); err != nil {
			return &ValidationError{Name: "html_url", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.html_url": %w`, err)}
		}
	}
	if v, ok := gru.mutation.StarCount(); ok {
		if err := githubrepository.StarCountValidator(v); err != nil {
			return &ValidationError{Name: "star_count", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.star_count": %w`, err)}
		}
	}
	return nil
}

func (gru *GithubRepositoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(githubrepository.Table, githubrepository.Columns, sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt))
	if ps := gru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gru.mutation.RepoID(); ok {
		_spec.SetField(githubrepository.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := gru.mutation.AddedRepoID(); ok {
		_spec.AddField(githubrepository.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := gru.mutation.Name(); ok {
		_spec.SetField(githubrepository.FieldName, field.TypeString, value)
	}
	if value, ok := gru.mutation.FullName(); ok {
		_spec.SetField(githubrepository.FieldFullName, field.TypeString, value)
	}
	if value, ok := gru.mutation.OwnerLogin(); ok {
		_spec.SetField(githubrepository.FieldOwnerLogin, field.TypeString, value)
	}
	if value, ok := gru.mutation.Owner(); ok {
		_spec.SetField(githubrepository.FieldOwner, field.TypeJSON, value)
	}
	if value, ok := gru.mutation.Public(); ok {
		_spec.SetField(githubrepository.FieldPublic, field.TypeBool, value)
	}
	if value, ok := gru.mutation.HTMLURL(); ok {
		_spec.SetField(githubrepository.FieldHTMLURL, field.TypeString, value)
	}
	if value, ok := gru.mutation.Description(); ok {
		_spec.SetField(githubrepository.FieldDescription, field.TypeString, value)
	}
	if gru.mutation.DescriptionCleared() {
		_spec.ClearField(githubrepository.FieldDescription, field.TypeString)
	}
	if value, ok := gru.mutation.Fork(); ok {
		_spec.SetField(githubrepository.FieldFork, field.TypeBool, value)
	}
	if value, ok := gru.mutation.Homepage(); ok {
		_spec.SetField(githubrepository.FieldHomepage, field.TypeString, value)
	}
	if gru.mutation.HomepageCleared() {
		_spec.ClearField(githubrepository.FieldHomepage, field.TypeString)
	}
	if value, ok := gru.mutation.StarCount(); ok {
		_spec.SetField(githubrepository.FieldStarCount, field.TypeInt, value)
	}
	if value, ok := gru.mutation.AddedStarCount(); ok {
		_spec.AddField(githubrepository.FieldStarCount, field.TypeInt, value)
	}
	if value, ok := gru.mutation.DefaultBranch(); ok {
		_spec.SetField(githubrepository.FieldDefaultBranch, field.TypeString, value)
	}
	if value, ok := gru.mutation.IsTemplate(); ok {
		_spec.SetField(githubrepository.FieldIsTemplate, field.TypeBool, value)
	}
	if value, ok := gru.mutation.HasIssues(); ok {
		_spec.SetField(githubrepository.FieldHasIssues, field.TypeBool, value)
	}
	if value, ok := gru.mutation.Archived(); ok {
		_spec.SetField(githubrepository.FieldArchived, field.TypeBool, value)
	}
	if value, ok := gru.mutation.PushedAt(); ok {
		_spec.SetField(githubrepository.FieldPushedAt, field.TypeTime, value)
	}
	if gru.mutation.PushedAtCleared() {
		_spec.ClearField(githubrepository.FieldPushedAt, field.TypeTime)
	}
	if value, ok := gru.mutation.CreatedAt(); ok {
		_spec.SetField(githubrepository.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := gru.mutation.UpdatedAt(); ok {
		_spec.SetField(githubrepository.FieldUpdatedAt, field.TypeTime, value)
	}
	if gru.mutation.UpdatedAtCleared() {
		_spec.ClearField(githubrepository.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := gru.mutation.License(); ok {
		_spec.SetField(githubrepository.FieldLicense, field.TypeJSON, value)
	}
	if gru.mutation.LicenseCleared() {
		_spec.ClearField(githubrepository.FieldLicense, field.TypeJSON)
	}
	if gru.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   githubrepository.LabelsTable,
			Columns: githubrepository.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.RemovedLabelsIDs(); len(nodes) > 0 && !gru.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   githubrepository.LabelsTable,
			Columns: githubrepository.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.LabelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   githubrepository.LabelsTable,
			Columns: githubrepository.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gru.mutation.ReleasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrepository.ReleasesTable,
			Columns: []string{githubrepository.ReleasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrelease.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.RemovedReleasesIDs(); len(nodes) > 0 && !gru.mutation.ReleasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrepository.ReleasesTable,
			Columns: []string{githubrepository.ReleasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrelease.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gru.mutation.ReleasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrepository.ReleasesTable,
			Columns: []string{githubrepository.ReleasesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, gru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubrepository.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gru.mutation.done = true
	return n, nil
}

// GithubRepositoryUpdateOne is the builder for updating a single GithubRepository entity.
type GithubRepositoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GithubRepositoryMutation
}

// SetRepoID sets the "repo_id" field.
func (gruo *GithubRepositoryUpdateOne) SetRepoID(i int64) *GithubRepositoryUpdateOne {
	gruo.mutation.ResetRepoID()
	gruo.mutation.SetRepoID(i)
	return gruo
}

// AddRepoID adds i to the "repo_id" field.
func (gruo *GithubRepositoryUpdateOne) AddRepoID(i int64) *GithubRepositoryUpdateOne {
	gruo.mutation.AddRepoID(i)
	return gruo
}

// SetName sets the "name" field.
func (gruo *GithubRepositoryUpdateOne) SetName(s string) *GithubRepositoryUpdateOne {
	gruo.mutation.SetName(s)
	return gruo
}

// SetFullName sets the "full_name" field.
func (gruo *GithubRepositoryUpdateOne) SetFullName(s string) *GithubRepositoryUpdateOne {
	gruo.mutation.SetFullName(s)
	return gruo
}

// SetOwnerLogin sets the "owner_login" field.
func (gruo *GithubRepositoryUpdateOne) SetOwnerLogin(s string) *GithubRepositoryUpdateOne {
	gruo.mutation.SetOwnerLogin(s)
	return gruo
}

// SetOwner sets the "owner" field.
func (gruo *GithubRepositoryUpdateOne) SetOwner(gi *github.User) *GithubRepositoryUpdateOne {
	gruo.mutation.SetOwner(gi)
	return gruo
}

// SetPublic sets the "public" field.
func (gruo *GithubRepositoryUpdateOne) SetPublic(b bool) *GithubRepositoryUpdateOne {
	gruo.mutation.SetPublic(b)
	return gruo
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (gruo *GithubRepositoryUpdateOne) SetNillablePublic(b *bool) *GithubRepositoryUpdateOne {
	if b != nil {
		gruo.SetPublic(*b)
	}
	return gruo
}

// SetHTMLURL sets the "html_url" field.
func (gruo *GithubRepositoryUpdateOne) SetHTMLURL(s string) *GithubRepositoryUpdateOne {
	gruo.mutation.SetHTMLURL(s)
	return gruo
}

// SetDescription sets the "description" field.
func (gruo *GithubRepositoryUpdateOne) SetDescription(s string) *GithubRepositoryUpdateOne {
	gruo.mutation.SetDescription(s)
	return gruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gruo *GithubRepositoryUpdateOne) SetNillableDescription(s *string) *GithubRepositoryUpdateOne {
	if s != nil {
		gruo.SetDescription(*s)
	}
	return gruo
}

// ClearDescription clears the value of the "description" field.
func (gruo *GithubRepositoryUpdateOne) ClearDescription() *GithubRepositoryUpdateOne {
	gruo.mutation.ClearDescription()
	return gruo
}

// SetFork sets the "fork" field.
func (gruo *GithubRepositoryUpdateOne) SetFork(b bool) *GithubRepositoryUpdateOne {
	gruo.mutation.SetFork(b)
	return gruo
}

// SetNillableFork sets the "fork" field if the given value is not nil.
func (gruo *GithubRepositoryUpdateOne) SetNillableFork(b *bool) *GithubRepositoryUpdateOne {
	if b != nil {
		gruo.SetFork(*b)
	}
	return gruo
}

// SetHomepage sets the "homepage" field.
func (gruo *GithubRepositoryUpdateOne) SetHomepage(s string) *GithubRepositoryUpdateOne {
	gruo.mutation.SetHomepage(s)
	return gruo
}

// SetNillableHomepage sets the "homepage" field if the given value is not nil.
func (gruo *GithubRepositoryUpdateOne) SetNillableHomepage(s *string) *GithubRepositoryUpdateOne {
	if s != nil {
		gruo.SetHomepage(*s)
	}
	return gruo
}

// ClearHomepage clears the value of the "homepage" field.
func (gruo *GithubRepositoryUpdateOne) ClearHomepage() *GithubRepositoryUpdateOne {
	gruo.mutation.ClearHomepage()
	return gruo
}

// SetStarCount sets the "star_count" field.
func (gruo *GithubRepositoryUpdateOne) SetStarCount(i int) *GithubRepositoryUpdateOne {
	gruo.mutation.ResetStarCount()
	gruo.mutation.SetStarCount(i)
	return gruo
}

// SetNillableStarCount sets the "star_count" field if the given value is not nil.
func (gruo *GithubRepositoryUpdateOne) SetNillableStarCount(i *int) *GithubRepositoryUpdateOne {
	if i != nil {
		gruo.SetStarCount(*i)
	}
	return gruo
}

// AddStarCount adds i to the "star_count" field.
func (gruo *GithubRepositoryUpdateOne) AddStarCount(i int) *GithubRepositoryUpdateOne {
	gruo.mutation.AddStarCount(i)
	return gruo
}

// SetDefaultBranch sets the "default_branch" field.
func (gruo *GithubRepositoryUpdateOne) SetDefaultBranch(s string) *GithubRepositoryUpdateOne {
	gruo.mutation.SetDefaultBranch(s)
	return gruo
}

// SetIsTemplate sets the "is_template" field.
func (gruo *GithubRepositoryUpdateOne) SetIsTemplate(b bool) *GithubRepositoryUpdateOne {
	gruo.mutation.SetIsTemplate(b)
	return gruo
}

// SetNillableIsTemplate sets the "is_template" field if the given value is not nil.
func (gruo *GithubRepositoryUpdateOne) SetNillableIsTemplate(b *bool) *GithubRepositoryUpdateOne {
	if b != nil {
		gruo.SetIsTemplate(*b)
	}
	return gruo
}

// SetHasIssues sets the "has_issues" field.
func (gruo *GithubRepositoryUpdateOne) SetHasIssues(b bool) *GithubRepositoryUpdateOne {
	gruo.mutation.SetHasIssues(b)
	return gruo
}

// SetNillableHasIssues sets the "has_issues" field if the given value is not nil.
func (gruo *GithubRepositoryUpdateOne) SetNillableHasIssues(b *bool) *GithubRepositoryUpdateOne {
	if b != nil {
		gruo.SetHasIssues(*b)
	}
	return gruo
}

// SetArchived sets the "archived" field.
func (gruo *GithubRepositoryUpdateOne) SetArchived(b bool) *GithubRepositoryUpdateOne {
	gruo.mutation.SetArchived(b)
	return gruo
}

// SetNillableArchived sets the "archived" field if the given value is not nil.
func (gruo *GithubRepositoryUpdateOne) SetNillableArchived(b *bool) *GithubRepositoryUpdateOne {
	if b != nil {
		gruo.SetArchived(*b)
	}
	return gruo
}

// SetPushedAt sets the "pushed_at" field.
func (gruo *GithubRepositoryUpdateOne) SetPushedAt(t time.Time) *GithubRepositoryUpdateOne {
	gruo.mutation.SetPushedAt(t)
	return gruo
}

// SetNillablePushedAt sets the "pushed_at" field if the given value is not nil.
func (gruo *GithubRepositoryUpdateOne) SetNillablePushedAt(t *time.Time) *GithubRepositoryUpdateOne {
	if t != nil {
		gruo.SetPushedAt(*t)
	}
	return gruo
}

// ClearPushedAt clears the value of the "pushed_at" field.
func (gruo *GithubRepositoryUpdateOne) ClearPushedAt() *GithubRepositoryUpdateOne {
	gruo.mutation.ClearPushedAt()
	return gruo
}

// SetCreatedAt sets the "created_at" field.
func (gruo *GithubRepositoryUpdateOne) SetCreatedAt(t time.Time) *GithubRepositoryUpdateOne {
	gruo.mutation.SetCreatedAt(t)
	return gruo
}

// SetUpdatedAt sets the "updated_at" field.
func (gruo *GithubRepositoryUpdateOne) SetUpdatedAt(t time.Time) *GithubRepositoryUpdateOne {
	gruo.mutation.SetUpdatedAt(t)
	return gruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gruo *GithubRepositoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *GithubRepositoryUpdateOne {
	if t != nil {
		gruo.SetUpdatedAt(*t)
	}
	return gruo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (gruo *GithubRepositoryUpdateOne) ClearUpdatedAt() *GithubRepositoryUpdateOne {
	gruo.mutation.ClearUpdatedAt()
	return gruo
}

// SetLicense sets the "license" field.
func (gruo *GithubRepositoryUpdateOne) SetLicense(gi *github.License) *GithubRepositoryUpdateOne {
	gruo.mutation.SetLicense(gi)
	return gruo
}

// ClearLicense clears the value of the "license" field.
func (gruo *GithubRepositoryUpdateOne) ClearLicense() *GithubRepositoryUpdateOne {
	gruo.mutation.ClearLicense()
	return gruo
}

// AddLabelIDs adds the "labels" edge to the Label entity by IDs.
func (gruo *GithubRepositoryUpdateOne) AddLabelIDs(ids ...int) *GithubRepositoryUpdateOne {
	gruo.mutation.AddLabelIDs(ids...)
	return gruo
}

// AddLabels adds the "labels" edges to the Label entity.
func (gruo *GithubRepositoryUpdateOne) AddLabels(l ...*Label) *GithubRepositoryUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return gruo.AddLabelIDs(ids...)
}

// AddReleaseIDs adds the "releases" edge to the GithubRelease entity by IDs.
func (gruo *GithubRepositoryUpdateOne) AddReleaseIDs(ids ...int) *GithubRepositoryUpdateOne {
	gruo.mutation.AddReleaseIDs(ids...)
	return gruo
}

// AddReleases adds the "releases" edges to the GithubRelease entity.
func (gruo *GithubRepositoryUpdateOne) AddReleases(g ...*GithubRelease) *GithubRepositoryUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gruo.AddReleaseIDs(ids...)
}

// Mutation returns the GithubRepositoryMutation object of the builder.
func (gruo *GithubRepositoryUpdateOne) Mutation() *GithubRepositoryMutation {
	return gruo.mutation
}

// ClearLabels clears all "labels" edges to the Label entity.
func (gruo *GithubRepositoryUpdateOne) ClearLabels() *GithubRepositoryUpdateOne {
	gruo.mutation.ClearLabels()
	return gruo
}

// RemoveLabelIDs removes the "labels" edge to Label entities by IDs.
func (gruo *GithubRepositoryUpdateOne) RemoveLabelIDs(ids ...int) *GithubRepositoryUpdateOne {
	gruo.mutation.RemoveLabelIDs(ids...)
	return gruo
}

// RemoveLabels removes "labels" edges to Label entities.
func (gruo *GithubRepositoryUpdateOne) RemoveLabels(l ...*Label) *GithubRepositoryUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return gruo.RemoveLabelIDs(ids...)
}

// ClearReleases clears all "releases" edges to the GithubRelease entity.
func (gruo *GithubRepositoryUpdateOne) ClearReleases() *GithubRepositoryUpdateOne {
	gruo.mutation.ClearReleases()
	return gruo
}

// RemoveReleaseIDs removes the "releases" edge to GithubRelease entities by IDs.
func (gruo *GithubRepositoryUpdateOne) RemoveReleaseIDs(ids ...int) *GithubRepositoryUpdateOne {
	gruo.mutation.RemoveReleaseIDs(ids...)
	return gruo
}

// RemoveReleases removes "releases" edges to GithubRelease entities.
func (gruo *GithubRepositoryUpdateOne) RemoveReleases(g ...*GithubRelease) *GithubRepositoryUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gruo.RemoveReleaseIDs(ids...)
}

// Where appends a list predicates to the GithubRepositoryUpdate builder.
func (gruo *GithubRepositoryUpdateOne) Where(ps ...predicate.GithubRepository) *GithubRepositoryUpdateOne {
	gruo.mutation.Where(ps...)
	return gruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gruo *GithubRepositoryUpdateOne) Select(field string, fields ...string) *GithubRepositoryUpdateOne {
	gruo.fields = append([]string{field}, fields...)
	return gruo
}

// Save executes the query and returns the updated GithubRepository entity.
func (gruo *GithubRepositoryUpdateOne) Save(ctx context.Context) (*GithubRepository, error) {
	return withHooks(ctx, gruo.sqlSave, gruo.mutation, gruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gruo *GithubRepositoryUpdateOne) SaveX(ctx context.Context) *GithubRepository {
	node, err := gruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gruo *GithubRepositoryUpdateOne) Exec(ctx context.Context) error {
	_, err := gruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gruo *GithubRepositoryUpdateOne) ExecX(ctx context.Context) {
	if err := gruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gruo *GithubRepositoryUpdateOne) check() error {
	if v, ok := gruo.mutation.RepoID(); ok {
		if err := githubrepository.RepoIDValidator(v); err != nil {
			return &ValidationError{Name: "repo_id", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.repo_id": %w`, err)}
		}
	}
	if v, ok := gruo.mutation.Name(); ok {
		if err := githubrepository.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.name": %w`, err)}
		}
	}
	if v, ok := gruo.mutation.FullName(); ok {
		if err := githubrepository.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.full_name": %w`, err)}
		}
	}
	if v, ok := gruo.mutation.OwnerLogin(); ok {
		if err := githubrepository.OwnerLoginValidator(v); err != nil {
			return &ValidationError{Name: "owner_login", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.owner_login": %w`, err)}
		}
	}
	if v, ok := gruo.mutation.HTMLURL(); ok {
		if err := githubrepository.HTMLURLValidator(v); err != nil {
			return &ValidationError{Name: "html_url", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.html_url": %w`, err)}
		}
	}
	if v, ok := gruo.mutation.StarCount(); ok {
		if err := githubrepository.StarCountValidator(v); err != nil {
			return &ValidationError{Name: "star_count", err: fmt.Errorf(`ent: validator failed for field "GithubRepository.star_count": %w`, err)}
		}
	}
	return nil
}

func (gruo *GithubRepositoryUpdateOne) sqlSave(ctx context.Context) (_node *GithubRepository, err error) {
	if err := gruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(githubrepository.Table, githubrepository.Columns, sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt))
	id, ok := gruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GithubRepository.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githubrepository.FieldID)
		for _, f := range fields {
			if !githubrepository.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != githubrepository.FieldID {
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
	if value, ok := gruo.mutation.RepoID(); ok {
		_spec.SetField(githubrepository.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := gruo.mutation.AddedRepoID(); ok {
		_spec.AddField(githubrepository.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := gruo.mutation.Name(); ok {
		_spec.SetField(githubrepository.FieldName, field.TypeString, value)
	}
	if value, ok := gruo.mutation.FullName(); ok {
		_spec.SetField(githubrepository.FieldFullName, field.TypeString, value)
	}
	if value, ok := gruo.mutation.OwnerLogin(); ok {
		_spec.SetField(githubrepository.FieldOwnerLogin, field.TypeString, value)
	}
	if value, ok := gruo.mutation.Owner(); ok {
		_spec.SetField(githubrepository.FieldOwner, field.TypeJSON, value)
	}
	if value, ok := gruo.mutation.Public(); ok {
		_spec.SetField(githubrepository.FieldPublic, field.TypeBool, value)
	}
	if value, ok := gruo.mutation.HTMLURL(); ok {
		_spec.SetField(githubrepository.FieldHTMLURL, field.TypeString, value)
	}
	if value, ok := gruo.mutation.Description(); ok {
		_spec.SetField(githubrepository.FieldDescription, field.TypeString, value)
	}
	if gruo.mutation.DescriptionCleared() {
		_spec.ClearField(githubrepository.FieldDescription, field.TypeString)
	}
	if value, ok := gruo.mutation.Fork(); ok {
		_spec.SetField(githubrepository.FieldFork, field.TypeBool, value)
	}
	if value, ok := gruo.mutation.Homepage(); ok {
		_spec.SetField(githubrepository.FieldHomepage, field.TypeString, value)
	}
	if gruo.mutation.HomepageCleared() {
		_spec.ClearField(githubrepository.FieldHomepage, field.TypeString)
	}
	if value, ok := gruo.mutation.StarCount(); ok {
		_spec.SetField(githubrepository.FieldStarCount, field.TypeInt, value)
	}
	if value, ok := gruo.mutation.AddedStarCount(); ok {
		_spec.AddField(githubrepository.FieldStarCount, field.TypeInt, value)
	}
	if value, ok := gruo.mutation.DefaultBranch(); ok {
		_spec.SetField(githubrepository.FieldDefaultBranch, field.TypeString, value)
	}
	if value, ok := gruo.mutation.IsTemplate(); ok {
		_spec.SetField(githubrepository.FieldIsTemplate, field.TypeBool, value)
	}
	if value, ok := gruo.mutation.HasIssues(); ok {
		_spec.SetField(githubrepository.FieldHasIssues, field.TypeBool, value)
	}
	if value, ok := gruo.mutation.Archived(); ok {
		_spec.SetField(githubrepository.FieldArchived, field.TypeBool, value)
	}
	if value, ok := gruo.mutation.PushedAt(); ok {
		_spec.SetField(githubrepository.FieldPushedAt, field.TypeTime, value)
	}
	if gruo.mutation.PushedAtCleared() {
		_spec.ClearField(githubrepository.FieldPushedAt, field.TypeTime)
	}
	if value, ok := gruo.mutation.CreatedAt(); ok {
		_spec.SetField(githubrepository.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := gruo.mutation.UpdatedAt(); ok {
		_spec.SetField(githubrepository.FieldUpdatedAt, field.TypeTime, value)
	}
	if gruo.mutation.UpdatedAtCleared() {
		_spec.ClearField(githubrepository.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := gruo.mutation.License(); ok {
		_spec.SetField(githubrepository.FieldLicense, field.TypeJSON, value)
	}
	if gruo.mutation.LicenseCleared() {
		_spec.ClearField(githubrepository.FieldLicense, field.TypeJSON)
	}
	if gruo.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   githubrepository.LabelsTable,
			Columns: githubrepository.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.RemovedLabelsIDs(); len(nodes) > 0 && !gruo.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   githubrepository.LabelsTable,
			Columns: githubrepository.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.LabelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   githubrepository.LabelsTable,
			Columns: githubrepository.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gruo.mutation.ReleasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrepository.ReleasesTable,
			Columns: []string{githubrepository.ReleasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrelease.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.RemovedReleasesIDs(); len(nodes) > 0 && !gruo.mutation.ReleasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrepository.ReleasesTable,
			Columns: []string{githubrepository.ReleasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(githubrelease.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gruo.mutation.ReleasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githubrepository.ReleasesTable,
			Columns: []string{githubrepository.ReleasesColumn},
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
	_node = &GithubRepository{config: gruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubrepository.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gruo.mutation.done = true
	return _node, nil
}
