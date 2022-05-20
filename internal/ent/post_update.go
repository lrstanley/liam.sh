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
	"github.com/lrstanley/liam.sh/internal/ent/label"
	"github.com/lrstanley/liam.sh/internal/ent/post"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
	"github.com/lrstanley/liam.sh/internal/ent/user"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdateTime sets the "update_time" field.
func (pu *PostUpdate) SetUpdateTime(t time.Time) *PostUpdate {
	pu.mutation.SetUpdateTime(t)
	return pu
}

// SetSlug sets the "slug" field.
func (pu *PostUpdate) SetSlug(s string) *PostUpdate {
	pu.mutation.SetSlug(s)
	return pu
}

// SetTitle sets the "title" field.
func (pu *PostUpdate) SetTitle(s string) *PostUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetContent sets the "content" field.
func (pu *PostUpdate) SetContent(s string) *PostUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetContentHTML sets the "content_html" field.
func (pu *PostUpdate) SetContentHTML(s string) *PostUpdate {
	pu.mutation.SetContentHTML(s)
	return pu
}

// SetPublishedAt sets the "published_at" field.
func (pu *PostUpdate) SetPublishedAt(t time.Time) *PostUpdate {
	pu.mutation.SetPublishedAt(t)
	return pu
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (pu *PostUpdate) SetNillablePublishedAt(t *time.Time) *PostUpdate {
	if t != nil {
		pu.SetPublishedAt(*t)
	}
	return pu
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (pu *PostUpdate) SetAuthorID(id int) *PostUpdate {
	pu.mutation.SetAuthorID(id)
	return pu
}

// SetAuthor sets the "author" edge to the User entity.
func (pu *PostUpdate) SetAuthor(u *User) *PostUpdate {
	return pu.SetAuthorID(u.ID)
}

// AddLabelIDs adds the "labels" edge to the Label entity by IDs.
func (pu *PostUpdate) AddLabelIDs(ids ...int) *PostUpdate {
	pu.mutation.AddLabelIDs(ids...)
	return pu
}

// AddLabels adds the "labels" edges to the Label entity.
func (pu *PostUpdate) AddLabels(l ...*Label) *PostUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pu.AddLabelIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (pu *PostUpdate) ClearAuthor() *PostUpdate {
	pu.mutation.ClearAuthor()
	return pu
}

// ClearLabels clears all "labels" edges to the Label entity.
func (pu *PostUpdate) ClearLabels() *PostUpdate {
	pu.mutation.ClearLabels()
	return pu
}

// RemoveLabelIDs removes the "labels" edge to Label entities by IDs.
func (pu *PostUpdate) RemoveLabelIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveLabelIDs(ids...)
	return pu
}

// RemoveLabels removes "labels" edges to Label entities.
func (pu *PostUpdate) RemoveLabels(l ...*Label) *PostUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pu.RemoveLabelIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PostUpdate) defaults() error {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		if post.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized post.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := post.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pu *PostUpdate) check() error {
	if v, ok := pu.mutation.Slug(); ok {
		if err := post.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Post.slug": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Post.content": %w`, err)}
		}
	}
	if v, ok := pu.mutation.ContentHTML(); ok {
		if err := post.ContentHTMLValidator(v); err != nil {
			return &ValidationError{Name: "content_html", err: fmt.Errorf(`ent: validator failed for field "Post.content_html": %w`, err)}
		}
	}
	if _, ok := pu.mutation.AuthorID(); pu.mutation.AuthorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Post.author"`)
	}
	return nil
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: post.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldUpdateTime,
		})
	}
	if value, ok := pu.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldSlug,
		})
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldTitle,
		})
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldContent,
		})
	}
	if value, ok := pu.mutation.ContentHTML(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldContentHTML,
		})
	}
	if value, ok := pu.mutation.PublishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldPublishedAt,
		})
	}
	if pu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LabelsTable,
			Columns: post.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: label.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedLabelsIDs(); len(nodes) > 0 && !pu.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LabelsTable,
			Columns: post.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: label.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.LabelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LabelsTable,
			Columns: post.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: label.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetUpdateTime sets the "update_time" field.
func (puo *PostUpdateOne) SetUpdateTime(t time.Time) *PostUpdateOne {
	puo.mutation.SetUpdateTime(t)
	return puo
}

// SetSlug sets the "slug" field.
func (puo *PostUpdateOne) SetSlug(s string) *PostUpdateOne {
	puo.mutation.SetSlug(s)
	return puo
}

// SetTitle sets the "title" field.
func (puo *PostUpdateOne) SetTitle(s string) *PostUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetContent sets the "content" field.
func (puo *PostUpdateOne) SetContent(s string) *PostUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetContentHTML sets the "content_html" field.
func (puo *PostUpdateOne) SetContentHTML(s string) *PostUpdateOne {
	puo.mutation.SetContentHTML(s)
	return puo
}

// SetPublishedAt sets the "published_at" field.
func (puo *PostUpdateOne) SetPublishedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetPublishedAt(t)
	return puo
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillablePublishedAt(t *time.Time) *PostUpdateOne {
	if t != nil {
		puo.SetPublishedAt(*t)
	}
	return puo
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (puo *PostUpdateOne) SetAuthorID(id int) *PostUpdateOne {
	puo.mutation.SetAuthorID(id)
	return puo
}

// SetAuthor sets the "author" edge to the User entity.
func (puo *PostUpdateOne) SetAuthor(u *User) *PostUpdateOne {
	return puo.SetAuthorID(u.ID)
}

// AddLabelIDs adds the "labels" edge to the Label entity by IDs.
func (puo *PostUpdateOne) AddLabelIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddLabelIDs(ids...)
	return puo
}

// AddLabels adds the "labels" edges to the Label entity.
func (puo *PostUpdateOne) AddLabels(l ...*Label) *PostUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return puo.AddLabelIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (puo *PostUpdateOne) ClearAuthor() *PostUpdateOne {
	puo.mutation.ClearAuthor()
	return puo
}

// ClearLabels clears all "labels" edges to the Label entity.
func (puo *PostUpdateOne) ClearLabels() *PostUpdateOne {
	puo.mutation.ClearLabels()
	return puo
}

// RemoveLabelIDs removes the "labels" edge to Label entities by IDs.
func (puo *PostUpdateOne) RemoveLabelIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveLabelIDs(ids...)
	return puo
}

// RemoveLabels removes "labels" edges to Label entities.
func (puo *PostUpdateOne) RemoveLabels(l ...*Label) *PostUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return puo.RemoveLabelIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	var (
		err  error
		node *Post
	)
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Post)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PostMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PostUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		if post.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized post.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := post.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (puo *PostUpdateOne) check() error {
	if v, ok := puo.mutation.Slug(); ok {
		if err := post.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Post.slug": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Post.content": %w`, err)}
		}
	}
	if v, ok := puo.mutation.ContentHTML(); ok {
		if err := post.ContentHTMLValidator(v); err != nil {
			return &ValidationError{Name: "content_html", err: fmt.Errorf(`ent: validator failed for field "Post.content_html": %w`, err)}
		}
	}
	if _, ok := puo.mutation.AuthorID(); puo.mutation.AuthorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Post.author"`)
	}
	return nil
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: post.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldUpdateTime,
		})
	}
	if value, ok := puo.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldSlug,
		})
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldTitle,
		})
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldContent,
		})
	}
	if value, ok := puo.mutation.ContentHTML(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldContentHTML,
		})
	}
	if value, ok := puo.mutation.PublishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: post.FieldPublishedAt,
		})
	}
	if puo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LabelsTable,
			Columns: post.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: label.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedLabelsIDs(); len(nodes) > 0 && !puo.mutation.LabelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LabelsTable,
			Columns: post.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: label.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.LabelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LabelsTable,
			Columns: post.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: label.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
