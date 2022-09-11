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
	"github.com/lrstanley/liam.sh/internal/ent/post"
	"github.com/lrstanley/liam.sh/internal/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (uc *UserCreate) SetCreateTime(t time.Time) *UserCreate {
	uc.mutation.SetCreateTime(t)
	return uc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreateTime(*t)
	}
	return uc
}

// SetUpdateTime sets the "update_time" field.
func (uc *UserCreate) SetUpdateTime(t time.Time) *UserCreate {
	uc.mutation.SetUpdateTime(t)
	return uc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdateTime(*t)
	}
	return uc
}

// SetUserID sets the "user_id" field.
func (uc *UserCreate) SetUserID(i int) *UserCreate {
	uc.mutation.SetUserID(i)
	return uc
}

// SetLogin sets the "login" field.
func (uc *UserCreate) SetLogin(s string) *UserCreate {
	uc.mutation.SetLogin(s)
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetAvatarURL sets the "avatar_url" field.
func (uc *UserCreate) SetAvatarURL(s string) *UserCreate {
	uc.mutation.SetAvatarURL(s)
	return uc
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatarURL(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatarURL(*s)
	}
	return uc
}

// SetHTMLURL sets the "html_url" field.
func (uc *UserCreate) SetHTMLURL(s string) *UserCreate {
	uc.mutation.SetHTMLURL(s)
	return uc
}

// SetNillableHTMLURL sets the "html_url" field if the given value is not nil.
func (uc *UserCreate) SetNillableHTMLURL(s *string) *UserCreate {
	if s != nil {
		uc.SetHTMLURL(*s)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetLocation sets the "location" field.
func (uc *UserCreate) SetLocation(s string) *UserCreate {
	uc.mutation.SetLocation(s)
	return uc
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (uc *UserCreate) SetNillableLocation(s *string) *UserCreate {
	if s != nil {
		uc.SetLocation(*s)
	}
	return uc
}

// SetBio sets the "bio" field.
func (uc *UserCreate) SetBio(s string) *UserCreate {
	uc.mutation.SetBio(s)
	return uc
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (uc *UserCreate) SetNillableBio(s *string) *UserCreate {
	if s != nil {
		uc.SetBio(*s)
	}
	return uc
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uc *UserCreate) AddPostIDs(ids ...int) *UserCreate {
	uc.mutation.AddPostIDs(ids...)
	return uc
}

// AddPosts adds the "posts" edges to the Post entity.
func (uc *UserCreate) AddPosts(p ...*Post) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPostIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if err := uc.defaults(); err != nil {
		return nil, err
	}
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() error {
	if _, ok := uc.mutation.CreateTime(); !ok {
		if user.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := user.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		if user.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized user.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := user.DefaultUpdateTime()
		uc.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "User.create_time"`)}
	}
	if _, ok := uc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "User.update_time"`)}
	}
	if _, ok := uc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "User.user_id"`)}
	}
	if v, ok := uc.mutation.UserID(); ok {
		if err := user.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "User.user_id": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Login(); !ok {
		return &ValidationError{Name: "login", err: errors.New(`ent: missing required field "User.login"`)}
	}
	if v, ok := uc.mutation.Login(); ok {
		if err := user.LoginValidator(v); err != nil {
			return &ValidationError{Name: "login", err: fmt.Errorf(`ent: validator failed for field "User.login": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uc.mutation.AvatarURL(); ok {
		if err := user.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf(`ent: validator failed for field "User.avatar_url": %w`, err)}
		}
	}
	if v, ok := uc.mutation.HTMLURL(); ok {
		if err := user.HTMLURLValidator(v); err != nil {
			return &ValidationError{Name: "html_url", err: fmt.Errorf(`ent: validator failed for field "User.html_url": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Location(); ok {
		if err := user.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "User.location": %w`, err)}
		}
	}
	if v, ok := uc.mutation.Bio(); ok {
		if err := user.BioValidator(v); err != nil {
			return &ValidationError{Name: "bio", err: fmt.Errorf(`ent: validator failed for field "User.bio": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	_spec.OnConflict = uc.conflict
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := uc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := uc.mutation.Login(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLogin,
		})
		_node.Login = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
		_node.Name = value
	}
	if value, ok := uc.mutation.AvatarURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
		_node.AvatarURL = value
	}
	if value, ok := uc.mutation.HTMLURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldHTMLURL,
		})
		_node.HTMLURL = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := uc.mutation.Location(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLocation,
		})
		_node.Location = value
	}
	if value, ok := uc.mutation.Bio(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldBio,
		})
		_node.Bio = value
	}
	if nodes := uc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsert) SetUpdateTime(v time.Time) *UserUpsert {
	u.Set(user.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdateTime() *UserUpsert {
	u.SetExcluded(user.FieldUpdateTime)
	return u
}

// SetLogin sets the "login" field.
func (u *UserUpsert) SetLogin(v string) *UserUpsert {
	u.Set(user.FieldLogin, v)
	return u
}

// UpdateLogin sets the "login" field to the value that was provided on create.
func (u *UserUpsert) UpdateLogin() *UserUpsert {
	u.SetExcluded(user.FieldLogin)
	return u
}

// SetName sets the "name" field.
func (u *UserUpsert) SetName(v string) *UserUpsert {
	u.Set(user.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsert) UpdateName() *UserUpsert {
	u.SetExcluded(user.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *UserUpsert) ClearName() *UserUpsert {
	u.SetNull(user.FieldName)
	return u
}

// SetAvatarURL sets the "avatar_url" field.
func (u *UserUpsert) SetAvatarURL(v string) *UserUpsert {
	u.Set(user.FieldAvatarURL, v)
	return u
}

// UpdateAvatarURL sets the "avatar_url" field to the value that was provided on create.
func (u *UserUpsert) UpdateAvatarURL() *UserUpsert {
	u.SetExcluded(user.FieldAvatarURL)
	return u
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (u *UserUpsert) ClearAvatarURL() *UserUpsert {
	u.SetNull(user.FieldAvatarURL)
	return u
}

// SetHTMLURL sets the "html_url" field.
func (u *UserUpsert) SetHTMLURL(v string) *UserUpsert {
	u.Set(user.FieldHTMLURL, v)
	return u
}

// UpdateHTMLURL sets the "html_url" field to the value that was provided on create.
func (u *UserUpsert) UpdateHTMLURL() *UserUpsert {
	u.SetExcluded(user.FieldHTMLURL)
	return u
}

// ClearHTMLURL clears the value of the "html_url" field.
func (u *UserUpsert) ClearHTMLURL() *UserUpsert {
	u.SetNull(user.FieldHTMLURL)
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsert) SetEmail(v string) *UserUpsert {
	u.Set(user.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsert) UpdateEmail() *UserUpsert {
	u.SetExcluded(user.FieldEmail)
	return u
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsert) ClearEmail() *UserUpsert {
	u.SetNull(user.FieldEmail)
	return u
}

// SetLocation sets the "location" field.
func (u *UserUpsert) SetLocation(v string) *UserUpsert {
	u.Set(user.FieldLocation, v)
	return u
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *UserUpsert) UpdateLocation() *UserUpsert {
	u.SetExcluded(user.FieldLocation)
	return u
}

// ClearLocation clears the value of the "location" field.
func (u *UserUpsert) ClearLocation() *UserUpsert {
	u.SetNull(user.FieldLocation)
	return u
}

// SetBio sets the "bio" field.
func (u *UserUpsert) SetBio(v string) *UserUpsert {
	u.Set(user.FieldBio, v)
	return u
}

// UpdateBio sets the "bio" field to the value that was provided on create.
func (u *UserUpsert) UpdateBio() *UserUpsert {
	u.SetExcluded(user.FieldBio)
	return u
}

// ClearBio clears the value of the "bio" field.
func (u *UserUpsert) ClearBio() *UserUpsert {
	u.SetNull(user.FieldBio)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(user.FieldCreateTime)
		}
		if _, exists := u.create.mutation.UserID(); exists {
			s.SetIgnore(user.FieldUserID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsertOne) SetUpdateTime(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdateTime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetLogin sets the "login" field.
func (u *UserUpsertOne) SetLogin(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetLogin(v)
	})
}

// UpdateLogin sets the "login" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateLogin() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLogin()
	})
}

// SetName sets the "name" field.
func (u *UserUpsertOne) SetName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *UserUpsertOne) ClearName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearName()
	})
}

// SetAvatarURL sets the "avatar_url" field.
func (u *UserUpsertOne) SetAvatarURL(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetAvatarURL(v)
	})
}

// UpdateAvatarURL sets the "avatar_url" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateAvatarURL() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAvatarURL()
	})
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (u *UserUpsertOne) ClearAvatarURL() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearAvatarURL()
	})
}

// SetHTMLURL sets the "html_url" field.
func (u *UserUpsertOne) SetHTMLURL(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetHTMLURL(v)
	})
}

// UpdateHTMLURL sets the "html_url" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateHTMLURL() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateHTMLURL()
	})
}

// ClearHTMLURL clears the value of the "html_url" field.
func (u *UserUpsertOne) ClearHTMLURL() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearHTMLURL()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertOne) SetEmail(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsertOne) ClearEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearEmail()
	})
}

// SetLocation sets the "location" field.
func (u *UserUpsertOne) SetLocation(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetLocation(v)
	})
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateLocation() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLocation()
	})
}

// ClearLocation clears the value of the "location" field.
func (u *UserUpsertOne) ClearLocation() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearLocation()
	})
}

// SetBio sets the "bio" field.
func (u *UserUpsertOne) SetBio(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetBio(v)
	})
}

// UpdateBio sets the "bio" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateBio() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBio()
	})
}

// ClearBio clears the value of the "bio" field.
func (u *UserUpsertOne) ClearBio() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearBio()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(user.FieldCreateTime)
			}
			if _, exists := b.mutation.UserID(); exists {
				s.SetIgnore(user.FieldUserID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *UserUpsertBulk) SetUpdateTime(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdateTime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetLogin sets the "login" field.
func (u *UserUpsertBulk) SetLogin(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetLogin(v)
	})
}

// UpdateLogin sets the "login" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateLogin() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLogin()
	})
}

// SetName sets the "name" field.
func (u *UserUpsertBulk) SetName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *UserUpsertBulk) ClearName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearName()
	})
}

// SetAvatarURL sets the "avatar_url" field.
func (u *UserUpsertBulk) SetAvatarURL(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetAvatarURL(v)
	})
}

// UpdateAvatarURL sets the "avatar_url" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateAvatarURL() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAvatarURL()
	})
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (u *UserUpsertBulk) ClearAvatarURL() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearAvatarURL()
	})
}

// SetHTMLURL sets the "html_url" field.
func (u *UserUpsertBulk) SetHTMLURL(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetHTMLURL(v)
	})
}

// UpdateHTMLURL sets the "html_url" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateHTMLURL() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateHTMLURL()
	})
}

// ClearHTMLURL clears the value of the "html_url" field.
func (u *UserUpsertBulk) ClearHTMLURL() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearHTMLURL()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertBulk) SetEmail(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsertBulk) ClearEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearEmail()
	})
}

// SetLocation sets the "location" field.
func (u *UserUpsertBulk) SetLocation(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetLocation(v)
	})
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateLocation() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLocation()
	})
}

// ClearLocation clears the value of the "location" field.
func (u *UserUpsertBulk) ClearLocation() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearLocation()
	})
}

// SetBio sets the "bio" field.
func (u *UserUpsertBulk) SetBio(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetBio(v)
	})
}

// UpdateBio sets the "bio" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateBio() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateBio()
	})
}

// ClearBio clears the value of the "bio" field.
func (u *UserUpsertBulk) ClearBio() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearBio()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
