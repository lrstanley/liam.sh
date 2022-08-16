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
	"github.com/lrstanley/liam.sh/internal/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/ent/label"
	"github.com/lrstanley/liam.sh/internal/ent/post"
)

// LabelCreate is the builder for creating a Label entity.
type LabelCreate struct {
	config
	mutation *LabelMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (lc *LabelCreate) SetCreateTime(t time.Time) *LabelCreate {
	lc.mutation.SetCreateTime(t)
	return lc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (lc *LabelCreate) SetNillableCreateTime(t *time.Time) *LabelCreate {
	if t != nil {
		lc.SetCreateTime(*t)
	}
	return lc
}

// SetUpdateTime sets the "update_time" field.
func (lc *LabelCreate) SetUpdateTime(t time.Time) *LabelCreate {
	lc.mutation.SetUpdateTime(t)
	return lc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (lc *LabelCreate) SetNillableUpdateTime(t *time.Time) *LabelCreate {
	if t != nil {
		lc.SetUpdateTime(*t)
	}
	return lc
}

// SetName sets the "name" field.
func (lc *LabelCreate) SetName(s string) *LabelCreate {
	lc.mutation.SetName(s)
	return lc
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (lc *LabelCreate) AddPostIDs(ids ...int) *LabelCreate {
	lc.mutation.AddPostIDs(ids...)
	return lc
}

// AddPosts adds the "posts" edges to the Post entity.
func (lc *LabelCreate) AddPosts(p ...*Post) *LabelCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return lc.AddPostIDs(ids...)
}

// AddGithubRepositoryIDs adds the "github_repositories" edge to the GithubRepository entity by IDs.
func (lc *LabelCreate) AddGithubRepositoryIDs(ids ...int) *LabelCreate {
	lc.mutation.AddGithubRepositoryIDs(ids...)
	return lc
}

// AddGithubRepositories adds the "github_repositories" edges to the GithubRepository entity.
func (lc *LabelCreate) AddGithubRepositories(g ...*GithubRepository) *LabelCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return lc.AddGithubRepositoryIDs(ids...)
}

// Mutation returns the LabelMutation object of the builder.
func (lc *LabelCreate) Mutation() *LabelMutation {
	return lc.mutation
}

// Save creates the Label in the database.
func (lc *LabelCreate) Save(ctx context.Context) (*Label, error) {
	var (
		err  error
		node *Label
	)
	if err := lc.defaults(); err != nil {
		return nil, err
	}
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LabelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Label)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LabelMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LabelCreate) SaveX(ctx context.Context) *Label {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LabelCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LabelCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LabelCreate) defaults() error {
	if _, ok := lc.mutation.CreateTime(); !ok {
		if label.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized label.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := label.DefaultCreateTime()
		lc.mutation.SetCreateTime(v)
	}
	if _, ok := lc.mutation.UpdateTime(); !ok {
		if label.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized label.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := label.DefaultUpdateTime()
		lc.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lc *LabelCreate) check() error {
	if _, ok := lc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Label.create_time"`)}
	}
	if _, ok := lc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Label.update_time"`)}
	}
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Label.name"`)}
	}
	if v, ok := lc.mutation.Name(); ok {
		if err := label.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Label.name": %w`, err)}
		}
	}
	return nil
}

func (lc *LabelCreate) sqlSave(ctx context.Context) (*Label, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lc *LabelCreate) createSpec() (*Label, *sqlgraph.CreateSpec) {
	var (
		_node = &Label{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: label.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: label.FieldID,
			},
		}
	)
	_spec.OnConflict = lc.conflict
	if value, ok := lc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: label.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := lc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: label.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := lc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: label.FieldName,
		})
		_node.Name = value
	}
	if nodes := lc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.PostsTable,
			Columns: label.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.GithubRepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   label.GithubRepositoriesTable,
			Columns: label.GithubRepositoriesPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Label.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LabelUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (lc *LabelCreate) OnConflict(opts ...sql.ConflictOption) *LabelUpsertOne {
	lc.conflict = opts
	return &LabelUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Label.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lc *LabelCreate) OnConflictColumns(columns ...string) *LabelUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &LabelUpsertOne{
		create: lc,
	}
}

type (
	// LabelUpsertOne is the builder for "upsert"-ing
	//  one Label node.
	LabelUpsertOne struct {
		create *LabelCreate
	}

	// LabelUpsert is the "OnConflict" setter.
	LabelUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *LabelUpsert) SetCreateTime(v time.Time) *LabelUpsert {
	u.Set(label.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *LabelUpsert) UpdateCreateTime() *LabelUpsert {
	u.SetExcluded(label.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *LabelUpsert) SetUpdateTime(v time.Time) *LabelUpsert {
	u.Set(label.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *LabelUpsert) UpdateUpdateTime() *LabelUpsert {
	u.SetExcluded(label.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *LabelUpsert) SetName(v string) *LabelUpsert {
	u.Set(label.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *LabelUpsert) UpdateName() *LabelUpsert {
	u.SetExcluded(label.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Label.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *LabelUpsertOne) UpdateNewValues() *LabelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(label.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Label.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *LabelUpsertOne) Ignore() *LabelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LabelUpsertOne) DoNothing() *LabelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LabelCreate.OnConflict
// documentation for more info.
func (u *LabelUpsertOne) Update(set func(*LabelUpsert)) *LabelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LabelUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *LabelUpsertOne) SetCreateTime(v time.Time) *LabelUpsertOne {
	return u.Update(func(s *LabelUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *LabelUpsertOne) UpdateCreateTime() *LabelUpsertOne {
	return u.Update(func(s *LabelUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *LabelUpsertOne) SetUpdateTime(v time.Time) *LabelUpsertOne {
	return u.Update(func(s *LabelUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *LabelUpsertOne) UpdateUpdateTime() *LabelUpsertOne {
	return u.Update(func(s *LabelUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *LabelUpsertOne) SetName(v string) *LabelUpsertOne {
	return u.Update(func(s *LabelUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *LabelUpsertOne) UpdateName() *LabelUpsertOne {
	return u.Update(func(s *LabelUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *LabelUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LabelCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LabelUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LabelUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LabelUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LabelCreateBulk is the builder for creating many Label entities in bulk.
type LabelCreateBulk struct {
	config
	builders []*LabelCreate
	conflict []sql.ConflictOption
}

// Save creates the Label entities in the database.
func (lcb *LabelCreateBulk) Save(ctx context.Context) ([]*Label, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Label, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LabelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LabelCreateBulk) SaveX(ctx context.Context) []*Label {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LabelCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LabelCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Label.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LabelUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (lcb *LabelCreateBulk) OnConflict(opts ...sql.ConflictOption) *LabelUpsertBulk {
	lcb.conflict = opts
	return &LabelUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Label.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (lcb *LabelCreateBulk) OnConflictColumns(columns ...string) *LabelUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &LabelUpsertBulk{
		create: lcb,
	}
}

// LabelUpsertBulk is the builder for "upsert"-ing
// a bulk of Label nodes.
type LabelUpsertBulk struct {
	create *LabelCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Label.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *LabelUpsertBulk) UpdateNewValues() *LabelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(label.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Label.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *LabelUpsertBulk) Ignore() *LabelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LabelUpsertBulk) DoNothing() *LabelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LabelCreateBulk.OnConflict
// documentation for more info.
func (u *LabelUpsertBulk) Update(set func(*LabelUpsert)) *LabelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LabelUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *LabelUpsertBulk) SetCreateTime(v time.Time) *LabelUpsertBulk {
	return u.Update(func(s *LabelUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *LabelUpsertBulk) UpdateCreateTime() *LabelUpsertBulk {
	return u.Update(func(s *LabelUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *LabelUpsertBulk) SetUpdateTime(v time.Time) *LabelUpsertBulk {
	return u.Update(func(s *LabelUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *LabelUpsertBulk) UpdateUpdateTime() *LabelUpsertBulk {
	return u.Update(func(s *LabelUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *LabelUpsertBulk) SetName(v string) *LabelUpsertBulk {
	return u.Update(func(s *LabelUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *LabelUpsertBulk) UpdateName() *LabelUpsertBulk {
	return u.Update(func(s *LabelUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *LabelUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LabelCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LabelCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LabelUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
