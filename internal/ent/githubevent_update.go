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
	"github.com/google/go-github/v48/github"
	"github.com/lrstanley/liam.sh/internal/ent/githubevent"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// GithubEventUpdate is the builder for updating GithubEvent entities.
type GithubEventUpdate struct {
	config
	hooks    []Hook
	mutation *GithubEventMutation
}

// Where appends a list predicates to the GithubEventUpdate builder.
func (geu *GithubEventUpdate) Where(ps ...predicate.GithubEvent) *GithubEventUpdate {
	geu.mutation.Where(ps...)
	return geu
}

// SetEventID sets the "event_id" field.
func (geu *GithubEventUpdate) SetEventID(s string) *GithubEventUpdate {
	geu.mutation.SetEventID(s)
	return geu
}

// SetEventType sets the "event_type" field.
func (geu *GithubEventUpdate) SetEventType(s string) *GithubEventUpdate {
	geu.mutation.SetEventType(s)
	return geu
}

// SetCreatedAt sets the "created_at" field.
func (geu *GithubEventUpdate) SetCreatedAt(t time.Time) *GithubEventUpdate {
	geu.mutation.SetCreatedAt(t)
	return geu
}

// SetPublic sets the "public" field.
func (geu *GithubEventUpdate) SetPublic(b bool) *GithubEventUpdate {
	geu.mutation.SetPublic(b)
	return geu
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (geu *GithubEventUpdate) SetNillablePublic(b *bool) *GithubEventUpdate {
	if b != nil {
		geu.SetPublic(*b)
	}
	return geu
}

// SetActorID sets the "actor_id" field.
func (geu *GithubEventUpdate) SetActorID(i int64) *GithubEventUpdate {
	geu.mutation.ResetActorID()
	geu.mutation.SetActorID(i)
	return geu
}

// AddActorID adds i to the "actor_id" field.
func (geu *GithubEventUpdate) AddActorID(i int64) *GithubEventUpdate {
	geu.mutation.AddActorID(i)
	return geu
}

// SetActor sets the "actor" field.
func (geu *GithubEventUpdate) SetActor(gi *github.User) *GithubEventUpdate {
	geu.mutation.SetActor(gi)
	return geu
}

// SetRepoID sets the "repo_id" field.
func (geu *GithubEventUpdate) SetRepoID(i int64) *GithubEventUpdate {
	geu.mutation.ResetRepoID()
	geu.mutation.SetRepoID(i)
	return geu
}

// AddRepoID adds i to the "repo_id" field.
func (geu *GithubEventUpdate) AddRepoID(i int64) *GithubEventUpdate {
	geu.mutation.AddRepoID(i)
	return geu
}

// SetRepo sets the "repo" field.
func (geu *GithubEventUpdate) SetRepo(gi *github.Repository) *GithubEventUpdate {
	geu.mutation.SetRepo(gi)
	return geu
}

// SetPayload sets the "payload" field.
func (geu *GithubEventUpdate) SetPayload(m map[string]interface{}) *GithubEventUpdate {
	geu.mutation.SetPayload(m)
	return geu
}

// Mutation returns the GithubEventMutation object of the builder.
func (geu *GithubEventUpdate) Mutation() *GithubEventMutation {
	return geu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (geu *GithubEventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(geu.hooks) == 0 {
		if err = geu.check(); err != nil {
			return 0, err
		}
		affected, err = geu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = geu.check(); err != nil {
				return 0, err
			}
			geu.mutation = mutation
			affected, err = geu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(geu.hooks) - 1; i >= 0; i-- {
			if geu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = geu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, geu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (geu *GithubEventUpdate) SaveX(ctx context.Context) int {
	affected, err := geu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (geu *GithubEventUpdate) Exec(ctx context.Context) error {
	_, err := geu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (geu *GithubEventUpdate) ExecX(ctx context.Context) {
	if err := geu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (geu *GithubEventUpdate) check() error {
	if v, ok := geu.mutation.EventType(); ok {
		if err := githubevent.EventTypeValidator(v); err != nil {
			return &ValidationError{Name: "event_type", err: fmt.Errorf(`ent: validator failed for field "GithubEvent.event_type": %w`, err)}
		}
	}
	if v, ok := geu.mutation.RepoID(); ok {
		if err := githubevent.RepoIDValidator(v); err != nil {
			return &ValidationError{Name: "repo_id", err: fmt.Errorf(`ent: validator failed for field "GithubEvent.repo_id": %w`, err)}
		}
	}
	return nil
}

func (geu *GithubEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   githubevent.Table,
			Columns: githubevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githubevent.FieldID,
			},
		},
	}
	if ps := geu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := geu.mutation.EventID(); ok {
		_spec.SetField(githubevent.FieldEventID, field.TypeString, value)
	}
	if value, ok := geu.mutation.EventType(); ok {
		_spec.SetField(githubevent.FieldEventType, field.TypeString, value)
	}
	if value, ok := geu.mutation.CreatedAt(); ok {
		_spec.SetField(githubevent.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := geu.mutation.Public(); ok {
		_spec.SetField(githubevent.FieldPublic, field.TypeBool, value)
	}
	if value, ok := geu.mutation.ActorID(); ok {
		_spec.SetField(githubevent.FieldActorID, field.TypeInt64, value)
	}
	if value, ok := geu.mutation.AddedActorID(); ok {
		_spec.AddField(githubevent.FieldActorID, field.TypeInt64, value)
	}
	if value, ok := geu.mutation.Actor(); ok {
		_spec.SetField(githubevent.FieldActor, field.TypeJSON, value)
	}
	if value, ok := geu.mutation.RepoID(); ok {
		_spec.SetField(githubevent.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := geu.mutation.AddedRepoID(); ok {
		_spec.AddField(githubevent.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := geu.mutation.Repo(); ok {
		_spec.SetField(githubevent.FieldRepo, field.TypeJSON, value)
	}
	if value, ok := geu.mutation.Payload(); ok {
		_spec.SetField(githubevent.FieldPayload, field.TypeJSON, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, geu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GithubEventUpdateOne is the builder for updating a single GithubEvent entity.
type GithubEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GithubEventMutation
}

// SetEventID sets the "event_id" field.
func (geuo *GithubEventUpdateOne) SetEventID(s string) *GithubEventUpdateOne {
	geuo.mutation.SetEventID(s)
	return geuo
}

// SetEventType sets the "event_type" field.
func (geuo *GithubEventUpdateOne) SetEventType(s string) *GithubEventUpdateOne {
	geuo.mutation.SetEventType(s)
	return geuo
}

// SetCreatedAt sets the "created_at" field.
func (geuo *GithubEventUpdateOne) SetCreatedAt(t time.Time) *GithubEventUpdateOne {
	geuo.mutation.SetCreatedAt(t)
	return geuo
}

// SetPublic sets the "public" field.
func (geuo *GithubEventUpdateOne) SetPublic(b bool) *GithubEventUpdateOne {
	geuo.mutation.SetPublic(b)
	return geuo
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (geuo *GithubEventUpdateOne) SetNillablePublic(b *bool) *GithubEventUpdateOne {
	if b != nil {
		geuo.SetPublic(*b)
	}
	return geuo
}

// SetActorID sets the "actor_id" field.
func (geuo *GithubEventUpdateOne) SetActorID(i int64) *GithubEventUpdateOne {
	geuo.mutation.ResetActorID()
	geuo.mutation.SetActorID(i)
	return geuo
}

// AddActorID adds i to the "actor_id" field.
func (geuo *GithubEventUpdateOne) AddActorID(i int64) *GithubEventUpdateOne {
	geuo.mutation.AddActorID(i)
	return geuo
}

// SetActor sets the "actor" field.
func (geuo *GithubEventUpdateOne) SetActor(gi *github.User) *GithubEventUpdateOne {
	geuo.mutation.SetActor(gi)
	return geuo
}

// SetRepoID sets the "repo_id" field.
func (geuo *GithubEventUpdateOne) SetRepoID(i int64) *GithubEventUpdateOne {
	geuo.mutation.ResetRepoID()
	geuo.mutation.SetRepoID(i)
	return geuo
}

// AddRepoID adds i to the "repo_id" field.
func (geuo *GithubEventUpdateOne) AddRepoID(i int64) *GithubEventUpdateOne {
	geuo.mutation.AddRepoID(i)
	return geuo
}

// SetRepo sets the "repo" field.
func (geuo *GithubEventUpdateOne) SetRepo(gi *github.Repository) *GithubEventUpdateOne {
	geuo.mutation.SetRepo(gi)
	return geuo
}

// SetPayload sets the "payload" field.
func (geuo *GithubEventUpdateOne) SetPayload(m map[string]interface{}) *GithubEventUpdateOne {
	geuo.mutation.SetPayload(m)
	return geuo
}

// Mutation returns the GithubEventMutation object of the builder.
func (geuo *GithubEventUpdateOne) Mutation() *GithubEventMutation {
	return geuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (geuo *GithubEventUpdateOne) Select(field string, fields ...string) *GithubEventUpdateOne {
	geuo.fields = append([]string{field}, fields...)
	return geuo
}

// Save executes the query and returns the updated GithubEvent entity.
func (geuo *GithubEventUpdateOne) Save(ctx context.Context) (*GithubEvent, error) {
	var (
		err  error
		node *GithubEvent
	)
	if len(geuo.hooks) == 0 {
		if err = geuo.check(); err != nil {
			return nil, err
		}
		node, err = geuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = geuo.check(); err != nil {
				return nil, err
			}
			geuo.mutation = mutation
			node, err = geuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(geuo.hooks) - 1; i >= 0; i-- {
			if geuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = geuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, geuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GithubEvent)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GithubEventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (geuo *GithubEventUpdateOne) SaveX(ctx context.Context) *GithubEvent {
	node, err := geuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (geuo *GithubEventUpdateOne) Exec(ctx context.Context) error {
	_, err := geuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (geuo *GithubEventUpdateOne) ExecX(ctx context.Context) {
	if err := geuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (geuo *GithubEventUpdateOne) check() error {
	if v, ok := geuo.mutation.EventType(); ok {
		if err := githubevent.EventTypeValidator(v); err != nil {
			return &ValidationError{Name: "event_type", err: fmt.Errorf(`ent: validator failed for field "GithubEvent.event_type": %w`, err)}
		}
	}
	if v, ok := geuo.mutation.RepoID(); ok {
		if err := githubevent.RepoIDValidator(v); err != nil {
			return &ValidationError{Name: "repo_id", err: fmt.Errorf(`ent: validator failed for field "GithubEvent.repo_id": %w`, err)}
		}
	}
	return nil
}

func (geuo *GithubEventUpdateOne) sqlSave(ctx context.Context) (_node *GithubEvent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   githubevent.Table,
			Columns: githubevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githubevent.FieldID,
			},
		},
	}
	id, ok := geuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GithubEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := geuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githubevent.FieldID)
		for _, f := range fields {
			if !githubevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != githubevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := geuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := geuo.mutation.EventID(); ok {
		_spec.SetField(githubevent.FieldEventID, field.TypeString, value)
	}
	if value, ok := geuo.mutation.EventType(); ok {
		_spec.SetField(githubevent.FieldEventType, field.TypeString, value)
	}
	if value, ok := geuo.mutation.CreatedAt(); ok {
		_spec.SetField(githubevent.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := geuo.mutation.Public(); ok {
		_spec.SetField(githubevent.FieldPublic, field.TypeBool, value)
	}
	if value, ok := geuo.mutation.ActorID(); ok {
		_spec.SetField(githubevent.FieldActorID, field.TypeInt64, value)
	}
	if value, ok := geuo.mutation.AddedActorID(); ok {
		_spec.AddField(githubevent.FieldActorID, field.TypeInt64, value)
	}
	if value, ok := geuo.mutation.Actor(); ok {
		_spec.SetField(githubevent.FieldActor, field.TypeJSON, value)
	}
	if value, ok := geuo.mutation.RepoID(); ok {
		_spec.SetField(githubevent.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := geuo.mutation.AddedRepoID(); ok {
		_spec.AddField(githubevent.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := geuo.mutation.Repo(); ok {
		_spec.SetField(githubevent.FieldRepo, field.TypeJSON, value)
	}
	if value, ok := geuo.mutation.Payload(); ok {
		_spec.SetField(githubevent.FieldPayload, field.TypeJSON, value)
	}
	_node = &GithubEvent{config: geuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, geuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
