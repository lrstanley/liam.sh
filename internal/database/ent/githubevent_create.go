// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
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
	"github.com/lrstanley/liam.sh/internal/database/ent/githubevent"
)

// GithubEventCreate is the builder for creating a GithubEvent entity.
type GithubEventCreate struct {
	config
	mutation *GithubEventMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEventID sets the "event_id" field.
func (gec *GithubEventCreate) SetEventID(s string) *GithubEventCreate {
	gec.mutation.SetEventID(s)
	return gec
}

// SetEventType sets the "event_type" field.
func (gec *GithubEventCreate) SetEventType(s string) *GithubEventCreate {
	gec.mutation.SetEventType(s)
	return gec
}

// SetCreatedAt sets the "created_at" field.
func (gec *GithubEventCreate) SetCreatedAt(t time.Time) *GithubEventCreate {
	gec.mutation.SetCreatedAt(t)
	return gec
}

// SetPublic sets the "public" field.
func (gec *GithubEventCreate) SetPublic(b bool) *GithubEventCreate {
	gec.mutation.SetPublic(b)
	return gec
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (gec *GithubEventCreate) SetNillablePublic(b *bool) *GithubEventCreate {
	if b != nil {
		gec.SetPublic(*b)
	}
	return gec
}

// SetActorID sets the "actor_id" field.
func (gec *GithubEventCreate) SetActorID(i int64) *GithubEventCreate {
	gec.mutation.SetActorID(i)
	return gec
}

// SetActor sets the "actor" field.
func (gec *GithubEventCreate) SetActor(gi *github.User) *GithubEventCreate {
	gec.mutation.SetActor(gi)
	return gec
}

// SetRepoID sets the "repo_id" field.
func (gec *GithubEventCreate) SetRepoID(i int64) *GithubEventCreate {
	gec.mutation.SetRepoID(i)
	return gec
}

// SetRepo sets the "repo" field.
func (gec *GithubEventCreate) SetRepo(gi *github.Repository) *GithubEventCreate {
	gec.mutation.SetRepo(gi)
	return gec
}

// SetPayload sets the "payload" field.
func (gec *GithubEventCreate) SetPayload(m map[string]interface{}) *GithubEventCreate {
	gec.mutation.SetPayload(m)
	return gec
}

// Mutation returns the GithubEventMutation object of the builder.
func (gec *GithubEventCreate) Mutation() *GithubEventMutation {
	return gec.mutation
}

// Save creates the GithubEvent in the database.
func (gec *GithubEventCreate) Save(ctx context.Context) (*GithubEvent, error) {
	if err := gec.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gec.sqlSave, gec.mutation, gec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gec *GithubEventCreate) SaveX(ctx context.Context) *GithubEvent {
	v, err := gec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gec *GithubEventCreate) Exec(ctx context.Context) error {
	_, err := gec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gec *GithubEventCreate) ExecX(ctx context.Context) {
	if err := gec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gec *GithubEventCreate) defaults() error {
	if _, ok := gec.mutation.Public(); !ok {
		v := githubevent.DefaultPublic
		gec.mutation.SetPublic(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gec *GithubEventCreate) check() error {
	if _, ok := gec.mutation.EventID(); !ok {
		return &ValidationError{Name: "event_id", err: errors.New(`ent: missing required field "GithubEvent.event_id"`)}
	}
	if _, ok := gec.mutation.EventType(); !ok {
		return &ValidationError{Name: "event_type", err: errors.New(`ent: missing required field "GithubEvent.event_type"`)}
	}
	if v, ok := gec.mutation.EventType(); ok {
		if err := githubevent.EventTypeValidator(v); err != nil {
			return &ValidationError{Name: "event_type", err: fmt.Errorf(`ent: validator failed for field "GithubEvent.event_type": %w`, err)}
		}
	}
	if _, ok := gec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GithubEvent.created_at"`)}
	}
	if _, ok := gec.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`ent: missing required field "GithubEvent.public"`)}
	}
	if _, ok := gec.mutation.ActorID(); !ok {
		return &ValidationError{Name: "actor_id", err: errors.New(`ent: missing required field "GithubEvent.actor_id"`)}
	}
	if _, ok := gec.mutation.Actor(); !ok {
		return &ValidationError{Name: "actor", err: errors.New(`ent: missing required field "GithubEvent.actor"`)}
	}
	if _, ok := gec.mutation.RepoID(); !ok {
		return &ValidationError{Name: "repo_id", err: errors.New(`ent: missing required field "GithubEvent.repo_id"`)}
	}
	if v, ok := gec.mutation.RepoID(); ok {
		if err := githubevent.RepoIDValidator(v); err != nil {
			return &ValidationError{Name: "repo_id", err: fmt.Errorf(`ent: validator failed for field "GithubEvent.repo_id": %w`, err)}
		}
	}
	if _, ok := gec.mutation.Repo(); !ok {
		return &ValidationError{Name: "repo", err: errors.New(`ent: missing required field "GithubEvent.repo"`)}
	}
	if _, ok := gec.mutation.Payload(); !ok {
		return &ValidationError{Name: "payload", err: errors.New(`ent: missing required field "GithubEvent.payload"`)}
	}
	return nil
}

func (gec *GithubEventCreate) sqlSave(ctx context.Context) (*GithubEvent, error) {
	if err := gec.check(); err != nil {
		return nil, err
	}
	_node, _spec := gec.createSpec()
	if err := sqlgraph.CreateNode(ctx, gec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gec.mutation.id = &_node.ID
	gec.mutation.done = true
	return _node, nil
}

func (gec *GithubEventCreate) createSpec() (*GithubEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &GithubEvent{config: gec.config}
		_spec = sqlgraph.NewCreateSpec(githubevent.Table, sqlgraph.NewFieldSpec(githubevent.FieldID, field.TypeInt))
	)
	_spec.OnConflict = gec.conflict
	if value, ok := gec.mutation.EventID(); ok {
		_spec.SetField(githubevent.FieldEventID, field.TypeString, value)
		_node.EventID = value
	}
	if value, ok := gec.mutation.EventType(); ok {
		_spec.SetField(githubevent.FieldEventType, field.TypeString, value)
		_node.EventType = value
	}
	if value, ok := gec.mutation.CreatedAt(); ok {
		_spec.SetField(githubevent.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gec.mutation.Public(); ok {
		_spec.SetField(githubevent.FieldPublic, field.TypeBool, value)
		_node.Public = value
	}
	if value, ok := gec.mutation.ActorID(); ok {
		_spec.SetField(githubevent.FieldActorID, field.TypeInt64, value)
		_node.ActorID = value
	}
	if value, ok := gec.mutation.Actor(); ok {
		_spec.SetField(githubevent.FieldActor, field.TypeJSON, value)
		_node.Actor = value
	}
	if value, ok := gec.mutation.RepoID(); ok {
		_spec.SetField(githubevent.FieldRepoID, field.TypeInt64, value)
		_node.RepoID = value
	}
	if value, ok := gec.mutation.Repo(); ok {
		_spec.SetField(githubevent.FieldRepo, field.TypeJSON, value)
		_node.Repo = value
	}
	if value, ok := gec.mutation.Payload(); ok {
		_spec.SetField(githubevent.FieldPayload, field.TypeJSON, value)
		_node.Payload = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GithubEvent.Create().
//		SetEventID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GithubEventUpsert) {
//			SetEventID(v+v).
//		}).
//		Exec(ctx)
func (gec *GithubEventCreate) OnConflict(opts ...sql.ConflictOption) *GithubEventUpsertOne {
	gec.conflict = opts
	return &GithubEventUpsertOne{
		create: gec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GithubEvent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gec *GithubEventCreate) OnConflictColumns(columns ...string) *GithubEventUpsertOne {
	gec.conflict = append(gec.conflict, sql.ConflictColumns(columns...))
	return &GithubEventUpsertOne{
		create: gec,
	}
}

type (
	// GithubEventUpsertOne is the builder for "upsert"-ing
	//  one GithubEvent node.
	GithubEventUpsertOne struct {
		create *GithubEventCreate
	}

	// GithubEventUpsert is the "OnConflict" setter.
	GithubEventUpsert struct {
		*sql.UpdateSet
	}
)

// SetEventID sets the "event_id" field.
func (u *GithubEventUpsert) SetEventID(v string) *GithubEventUpsert {
	u.Set(githubevent.FieldEventID, v)
	return u
}

// UpdateEventID sets the "event_id" field to the value that was provided on create.
func (u *GithubEventUpsert) UpdateEventID() *GithubEventUpsert {
	u.SetExcluded(githubevent.FieldEventID)
	return u
}

// SetEventType sets the "event_type" field.
func (u *GithubEventUpsert) SetEventType(v string) *GithubEventUpsert {
	u.Set(githubevent.FieldEventType, v)
	return u
}

// UpdateEventType sets the "event_type" field to the value that was provided on create.
func (u *GithubEventUpsert) UpdateEventType() *GithubEventUpsert {
	u.SetExcluded(githubevent.FieldEventType)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubEventUpsert) SetCreatedAt(v time.Time) *GithubEventUpsert {
	u.Set(githubevent.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubEventUpsert) UpdateCreatedAt() *GithubEventUpsert {
	u.SetExcluded(githubevent.FieldCreatedAt)
	return u
}

// SetPublic sets the "public" field.
func (u *GithubEventUpsert) SetPublic(v bool) *GithubEventUpsert {
	u.Set(githubevent.FieldPublic, v)
	return u
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GithubEventUpsert) UpdatePublic() *GithubEventUpsert {
	u.SetExcluded(githubevent.FieldPublic)
	return u
}

// SetActorID sets the "actor_id" field.
func (u *GithubEventUpsert) SetActorID(v int64) *GithubEventUpsert {
	u.Set(githubevent.FieldActorID, v)
	return u
}

// UpdateActorID sets the "actor_id" field to the value that was provided on create.
func (u *GithubEventUpsert) UpdateActorID() *GithubEventUpsert {
	u.SetExcluded(githubevent.FieldActorID)
	return u
}

// AddActorID adds v to the "actor_id" field.
func (u *GithubEventUpsert) AddActorID(v int64) *GithubEventUpsert {
	u.Add(githubevent.FieldActorID, v)
	return u
}

// SetActor sets the "actor" field.
func (u *GithubEventUpsert) SetActor(v *github.User) *GithubEventUpsert {
	u.Set(githubevent.FieldActor, v)
	return u
}

// UpdateActor sets the "actor" field to the value that was provided on create.
func (u *GithubEventUpsert) UpdateActor() *GithubEventUpsert {
	u.SetExcluded(githubevent.FieldActor)
	return u
}

// SetRepoID sets the "repo_id" field.
func (u *GithubEventUpsert) SetRepoID(v int64) *GithubEventUpsert {
	u.Set(githubevent.FieldRepoID, v)
	return u
}

// UpdateRepoID sets the "repo_id" field to the value that was provided on create.
func (u *GithubEventUpsert) UpdateRepoID() *GithubEventUpsert {
	u.SetExcluded(githubevent.FieldRepoID)
	return u
}

// AddRepoID adds v to the "repo_id" field.
func (u *GithubEventUpsert) AddRepoID(v int64) *GithubEventUpsert {
	u.Add(githubevent.FieldRepoID, v)
	return u
}

// SetRepo sets the "repo" field.
func (u *GithubEventUpsert) SetRepo(v *github.Repository) *GithubEventUpsert {
	u.Set(githubevent.FieldRepo, v)
	return u
}

// UpdateRepo sets the "repo" field to the value that was provided on create.
func (u *GithubEventUpsert) UpdateRepo() *GithubEventUpsert {
	u.SetExcluded(githubevent.FieldRepo)
	return u
}

// SetPayload sets the "payload" field.
func (u *GithubEventUpsert) SetPayload(v map[string]interface{}) *GithubEventUpsert {
	u.Set(githubevent.FieldPayload, v)
	return u
}

// UpdatePayload sets the "payload" field to the value that was provided on create.
func (u *GithubEventUpsert) UpdatePayload() *GithubEventUpsert {
	u.SetExcluded(githubevent.FieldPayload)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.GithubEvent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GithubEventUpsertOne) UpdateNewValues() *GithubEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GithubEvent.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GithubEventUpsertOne) Ignore() *GithubEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GithubEventUpsertOne) DoNothing() *GithubEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GithubEventCreate.OnConflict
// documentation for more info.
func (u *GithubEventUpsertOne) Update(set func(*GithubEventUpsert)) *GithubEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GithubEventUpsert{UpdateSet: update})
	}))
	return u
}

// SetEventID sets the "event_id" field.
func (u *GithubEventUpsertOne) SetEventID(v string) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetEventID(v)
	})
}

// UpdateEventID sets the "event_id" field to the value that was provided on create.
func (u *GithubEventUpsertOne) UpdateEventID() *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateEventID()
	})
}

// SetEventType sets the "event_type" field.
func (u *GithubEventUpsertOne) SetEventType(v string) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetEventType(v)
	})
}

// UpdateEventType sets the "event_type" field to the value that was provided on create.
func (u *GithubEventUpsertOne) UpdateEventType() *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateEventType()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubEventUpsertOne) SetCreatedAt(v time.Time) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubEventUpsertOne) UpdateCreatedAt() *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetPublic sets the "public" field.
func (u *GithubEventUpsertOne) SetPublic(v bool) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GithubEventUpsertOne) UpdatePublic() *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdatePublic()
	})
}

// SetActorID sets the "actor_id" field.
func (u *GithubEventUpsertOne) SetActorID(v int64) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetActorID(v)
	})
}

// AddActorID adds v to the "actor_id" field.
func (u *GithubEventUpsertOne) AddActorID(v int64) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.AddActorID(v)
	})
}

// UpdateActorID sets the "actor_id" field to the value that was provided on create.
func (u *GithubEventUpsertOne) UpdateActorID() *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateActorID()
	})
}

// SetActor sets the "actor" field.
func (u *GithubEventUpsertOne) SetActor(v *github.User) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetActor(v)
	})
}

// UpdateActor sets the "actor" field to the value that was provided on create.
func (u *GithubEventUpsertOne) UpdateActor() *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateActor()
	})
}

// SetRepoID sets the "repo_id" field.
func (u *GithubEventUpsertOne) SetRepoID(v int64) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetRepoID(v)
	})
}

// AddRepoID adds v to the "repo_id" field.
func (u *GithubEventUpsertOne) AddRepoID(v int64) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.AddRepoID(v)
	})
}

// UpdateRepoID sets the "repo_id" field to the value that was provided on create.
func (u *GithubEventUpsertOne) UpdateRepoID() *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateRepoID()
	})
}

// SetRepo sets the "repo" field.
func (u *GithubEventUpsertOne) SetRepo(v *github.Repository) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetRepo(v)
	})
}

// UpdateRepo sets the "repo" field to the value that was provided on create.
func (u *GithubEventUpsertOne) UpdateRepo() *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateRepo()
	})
}

// SetPayload sets the "payload" field.
func (u *GithubEventUpsertOne) SetPayload(v map[string]interface{}) *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetPayload(v)
	})
}

// UpdatePayload sets the "payload" field to the value that was provided on create.
func (u *GithubEventUpsertOne) UpdatePayload() *GithubEventUpsertOne {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdatePayload()
	})
}

// Exec executes the query.
func (u *GithubEventUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GithubEventCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GithubEventUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GithubEventUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GithubEventUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GithubEventCreateBulk is the builder for creating many GithubEvent entities in bulk.
type GithubEventCreateBulk struct {
	config
	err      error
	builders []*GithubEventCreate
	conflict []sql.ConflictOption
}

// Save creates the GithubEvent entities in the database.
func (gecb *GithubEventCreateBulk) Save(ctx context.Context) ([]*GithubEvent, error) {
	if gecb.err != nil {
		return nil, gecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gecb.builders))
	nodes := make([]*GithubEvent, len(gecb.builders))
	mutators := make([]Mutator, len(gecb.builders))
	for i := range gecb.builders {
		func(i int, root context.Context) {
			builder := gecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GithubEventMutation)
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
					_, err = mutators[i+1].Mutate(root, gecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gecb *GithubEventCreateBulk) SaveX(ctx context.Context) []*GithubEvent {
	v, err := gecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gecb *GithubEventCreateBulk) Exec(ctx context.Context) error {
	_, err := gecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gecb *GithubEventCreateBulk) ExecX(ctx context.Context) {
	if err := gecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GithubEvent.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GithubEventUpsert) {
//			SetEventID(v+v).
//		}).
//		Exec(ctx)
func (gecb *GithubEventCreateBulk) OnConflict(opts ...sql.ConflictOption) *GithubEventUpsertBulk {
	gecb.conflict = opts
	return &GithubEventUpsertBulk{
		create: gecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GithubEvent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gecb *GithubEventCreateBulk) OnConflictColumns(columns ...string) *GithubEventUpsertBulk {
	gecb.conflict = append(gecb.conflict, sql.ConflictColumns(columns...))
	return &GithubEventUpsertBulk{
		create: gecb,
	}
}

// GithubEventUpsertBulk is the builder for "upsert"-ing
// a bulk of GithubEvent nodes.
type GithubEventUpsertBulk struct {
	create *GithubEventCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GithubEvent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GithubEventUpsertBulk) UpdateNewValues() *GithubEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GithubEvent.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GithubEventUpsertBulk) Ignore() *GithubEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GithubEventUpsertBulk) DoNothing() *GithubEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GithubEventCreateBulk.OnConflict
// documentation for more info.
func (u *GithubEventUpsertBulk) Update(set func(*GithubEventUpsert)) *GithubEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GithubEventUpsert{UpdateSet: update})
	}))
	return u
}

// SetEventID sets the "event_id" field.
func (u *GithubEventUpsertBulk) SetEventID(v string) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetEventID(v)
	})
}

// UpdateEventID sets the "event_id" field to the value that was provided on create.
func (u *GithubEventUpsertBulk) UpdateEventID() *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateEventID()
	})
}

// SetEventType sets the "event_type" field.
func (u *GithubEventUpsertBulk) SetEventType(v string) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetEventType(v)
	})
}

// UpdateEventType sets the "event_type" field to the value that was provided on create.
func (u *GithubEventUpsertBulk) UpdateEventType() *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateEventType()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *GithubEventUpsertBulk) SetCreatedAt(v time.Time) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GithubEventUpsertBulk) UpdateCreatedAt() *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetPublic sets the "public" field.
func (u *GithubEventUpsertBulk) SetPublic(v bool) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GithubEventUpsertBulk) UpdatePublic() *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdatePublic()
	})
}

// SetActorID sets the "actor_id" field.
func (u *GithubEventUpsertBulk) SetActorID(v int64) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetActorID(v)
	})
}

// AddActorID adds v to the "actor_id" field.
func (u *GithubEventUpsertBulk) AddActorID(v int64) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.AddActorID(v)
	})
}

// UpdateActorID sets the "actor_id" field to the value that was provided on create.
func (u *GithubEventUpsertBulk) UpdateActorID() *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateActorID()
	})
}

// SetActor sets the "actor" field.
func (u *GithubEventUpsertBulk) SetActor(v *github.User) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetActor(v)
	})
}

// UpdateActor sets the "actor" field to the value that was provided on create.
func (u *GithubEventUpsertBulk) UpdateActor() *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateActor()
	})
}

// SetRepoID sets the "repo_id" field.
func (u *GithubEventUpsertBulk) SetRepoID(v int64) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetRepoID(v)
	})
}

// AddRepoID adds v to the "repo_id" field.
func (u *GithubEventUpsertBulk) AddRepoID(v int64) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.AddRepoID(v)
	})
}

// UpdateRepoID sets the "repo_id" field to the value that was provided on create.
func (u *GithubEventUpsertBulk) UpdateRepoID() *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateRepoID()
	})
}

// SetRepo sets the "repo" field.
func (u *GithubEventUpsertBulk) SetRepo(v *github.Repository) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetRepo(v)
	})
}

// UpdateRepo sets the "repo" field to the value that was provided on create.
func (u *GithubEventUpsertBulk) UpdateRepo() *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdateRepo()
	})
}

// SetPayload sets the "payload" field.
func (u *GithubEventUpsertBulk) SetPayload(v map[string]interface{}) *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.SetPayload(v)
	})
}

// UpdatePayload sets the "payload" field to the value that was provided on create.
func (u *GithubEventUpsertBulk) UpdatePayload() *GithubEventUpsertBulk {
	return u.Update(func(s *GithubEventUpsert) {
		s.UpdatePayload()
	})
}

// Exec executes the query.
func (u *GithubEventUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GithubEventCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GithubEventCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GithubEventUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
