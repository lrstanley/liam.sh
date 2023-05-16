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
	"github.com/lrstanley/liam.sh/internal/database/ent/label"
	"github.com/lrstanley/liam.sh/internal/database/ent/post"
	"github.com/lrstanley/liam.sh/internal/database/ent/user"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (pc *PostCreate) SetCreateTime(t time.Time) *PostCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreateTime(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PostCreate) SetUpdateTime(t time.Time) *PostCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PostCreate) SetNillableUpdateTime(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetSlug sets the "slug" field.
func (pc *PostCreate) SetSlug(s string) *PostCreate {
	pc.mutation.SetSlug(s)
	return pc
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetContent sets the "content" field.
func (pc *PostCreate) SetContent(s string) *PostCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetContentHTML sets the "content_html" field.
func (pc *PostCreate) SetContentHTML(s string) *PostCreate {
	pc.mutation.SetContentHTML(s)
	return pc
}

// SetSummary sets the "summary" field.
func (pc *PostCreate) SetSummary(s string) *PostCreate {
	pc.mutation.SetSummary(s)
	return pc
}

// SetPublishedAt sets the "published_at" field.
func (pc *PostCreate) SetPublishedAt(t time.Time) *PostCreate {
	pc.mutation.SetPublishedAt(t)
	return pc
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (pc *PostCreate) SetNillablePublishedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetPublishedAt(*t)
	}
	return pc
}

// SetViewCount sets the "view_count" field.
func (pc *PostCreate) SetViewCount(i int) *PostCreate {
	pc.mutation.SetViewCount(i)
	return pc
}

// SetNillableViewCount sets the "view_count" field if the given value is not nil.
func (pc *PostCreate) SetNillableViewCount(i *int) *PostCreate {
	if i != nil {
		pc.SetViewCount(*i)
	}
	return pc
}

// SetPublic sets the "public" field.
func (pc *PostCreate) SetPublic(b bool) *PostCreate {
	pc.mutation.SetPublic(b)
	return pc
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (pc *PostCreate) SetNillablePublic(b *bool) *PostCreate {
	if b != nil {
		pc.SetPublic(*b)
	}
	return pc
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (pc *PostCreate) SetAuthorID(id int) *PostCreate {
	pc.mutation.SetAuthorID(id)
	return pc
}

// SetAuthor sets the "author" edge to the User entity.
func (pc *PostCreate) SetAuthor(u *User) *PostCreate {
	return pc.SetAuthorID(u.ID)
}

// AddLabelIDs adds the "labels" edge to the Label entity by IDs.
func (pc *PostCreate) AddLabelIDs(ids ...int) *PostCreate {
	pc.mutation.AddLabelIDs(ids...)
	return pc
}

// AddLabels adds the "labels" edges to the Label entity.
func (pc *PostCreate) AddLabels(l ...*Label) *PostCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pc.AddLabelIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		if post.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized post.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := post.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		if post.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized post.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := post.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.PublishedAt(); !ok {
		if post.DefaultPublishedAt == nil {
			return fmt.Errorf("ent: uninitialized post.DefaultPublishedAt (forgotten import ent/runtime?)")
		}
		v := post.DefaultPublishedAt()
		pc.mutation.SetPublishedAt(v)
	}
	if _, ok := pc.mutation.ViewCount(); !ok {
		v := post.DefaultViewCount
		pc.mutation.SetViewCount(v)
	}
	if _, ok := pc.mutation.Public(); !ok {
		v := post.DefaultPublic
		pc.mutation.SetPublic(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Post.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Post.update_time"`)}
	}
	if _, ok := pc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Post.slug"`)}
	}
	if v, ok := pc.mutation.Slug(); ok {
		if err := post.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Post.slug": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Post.title"`)}
	}
	if v, ok := pc.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Post.content"`)}
	}
	if v, ok := pc.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Post.content": %w`, err)}
		}
	}
	if _, ok := pc.mutation.ContentHTML(); !ok {
		return &ValidationError{Name: "content_html", err: errors.New(`ent: missing required field "Post.content_html"`)}
	}
	if v, ok := pc.mutation.ContentHTML(); ok {
		if err := post.ContentHTMLValidator(v); err != nil {
			return &ValidationError{Name: "content_html", err: fmt.Errorf(`ent: validator failed for field "Post.content_html": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Summary(); !ok {
		return &ValidationError{Name: "summary", err: errors.New(`ent: missing required field "Post.summary"`)}
	}
	if v, ok := pc.mutation.Summary(); ok {
		if err := post.SummaryValidator(v); err != nil {
			return &ValidationError{Name: "summary", err: fmt.Errorf(`ent: validator failed for field "Post.summary": %w`, err)}
		}
	}
	if _, ok := pc.mutation.PublishedAt(); !ok {
		return &ValidationError{Name: "published_at", err: errors.New(`ent: missing required field "Post.published_at"`)}
	}
	if _, ok := pc.mutation.ViewCount(); !ok {
		return &ValidationError{Name: "view_count", err: errors.New(`ent: missing required field "Post.view_count"`)}
	}
	if v, ok := pc.mutation.ViewCount(); ok {
		if err := post.ViewCountValidator(v); err != nil {
			return &ValidationError{Name: "view_count", err: fmt.Errorf(`ent: validator failed for field "Post.view_count": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`ent: missing required field "Post.public"`)}
	}
	if _, ok := pc.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required edge "Post.author"`)}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(post.Table, sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(post.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(post.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Slug(); ok {
		_spec.SetField(post.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := pc.mutation.ContentHTML(); ok {
		_spec.SetField(post.FieldContentHTML, field.TypeString, value)
		_node.ContentHTML = value
	}
	if value, ok := pc.mutation.Summary(); ok {
		_spec.SetField(post.FieldSummary, field.TypeString, value)
		_node.Summary = value
	}
	if value, ok := pc.mutation.PublishedAt(); ok {
		_spec.SetField(post.FieldPublishedAt, field.TypeTime, value)
		_node.PublishedAt = value
	}
	if value, ok := pc.mutation.ViewCount(); ok {
		_spec.SetField(post.FieldViewCount, field.TypeInt, value)
		_node.ViewCount = value
	}
	if value, ok := pc.mutation.Public(); ok {
		_spec.SetField(post.FieldPublic, field.TypeBool, value)
		_node.Public = value
	}
	if nodes := pc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_posts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.LabelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.LabelsTable,
			Columns: post.LabelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt),
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
//	client.Post.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PostUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pc *PostCreate) OnConflict(opts ...sql.ConflictOption) *PostUpsertOne {
	pc.conflict = opts
	return &PostUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PostCreate) OnConflictColumns(columns ...string) *PostUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PostUpsertOne{
		create: pc,
	}
}

type (
	// PostUpsertOne is the builder for "upsert"-ing
	//  one Post node.
	PostUpsertOne struct {
		create *PostCreate
	}

	// PostUpsert is the "OnConflict" setter.
	PostUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *PostUpsert) SetUpdateTime(v time.Time) *PostUpsert {
	u.Set(post.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PostUpsert) UpdateUpdateTime() *PostUpsert {
	u.SetExcluded(post.FieldUpdateTime)
	return u
}

// SetSlug sets the "slug" field.
func (u *PostUpsert) SetSlug(v string) *PostUpsert {
	u.Set(post.FieldSlug, v)
	return u
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *PostUpsert) UpdateSlug() *PostUpsert {
	u.SetExcluded(post.FieldSlug)
	return u
}

// SetTitle sets the "title" field.
func (u *PostUpsert) SetTitle(v string) *PostUpsert {
	u.Set(post.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsert) UpdateTitle() *PostUpsert {
	u.SetExcluded(post.FieldTitle)
	return u
}

// SetContent sets the "content" field.
func (u *PostUpsert) SetContent(v string) *PostUpsert {
	u.Set(post.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *PostUpsert) UpdateContent() *PostUpsert {
	u.SetExcluded(post.FieldContent)
	return u
}

// SetContentHTML sets the "content_html" field.
func (u *PostUpsert) SetContentHTML(v string) *PostUpsert {
	u.Set(post.FieldContentHTML, v)
	return u
}

// UpdateContentHTML sets the "content_html" field to the value that was provided on create.
func (u *PostUpsert) UpdateContentHTML() *PostUpsert {
	u.SetExcluded(post.FieldContentHTML)
	return u
}

// SetSummary sets the "summary" field.
func (u *PostUpsert) SetSummary(v string) *PostUpsert {
	u.Set(post.FieldSummary, v)
	return u
}

// UpdateSummary sets the "summary" field to the value that was provided on create.
func (u *PostUpsert) UpdateSummary() *PostUpsert {
	u.SetExcluded(post.FieldSummary)
	return u
}

// SetPublishedAt sets the "published_at" field.
func (u *PostUpsert) SetPublishedAt(v time.Time) *PostUpsert {
	u.Set(post.FieldPublishedAt, v)
	return u
}

// UpdatePublishedAt sets the "published_at" field to the value that was provided on create.
func (u *PostUpsert) UpdatePublishedAt() *PostUpsert {
	u.SetExcluded(post.FieldPublishedAt)
	return u
}

// SetViewCount sets the "view_count" field.
func (u *PostUpsert) SetViewCount(v int) *PostUpsert {
	u.Set(post.FieldViewCount, v)
	return u
}

// UpdateViewCount sets the "view_count" field to the value that was provided on create.
func (u *PostUpsert) UpdateViewCount() *PostUpsert {
	u.SetExcluded(post.FieldViewCount)
	return u
}

// AddViewCount adds v to the "view_count" field.
func (u *PostUpsert) AddViewCount(v int) *PostUpsert {
	u.Add(post.FieldViewCount, v)
	return u
}

// SetPublic sets the "public" field.
func (u *PostUpsert) SetPublic(v bool) *PostUpsert {
	u.Set(post.FieldPublic, v)
	return u
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *PostUpsert) UpdatePublic() *PostUpsert {
	u.SetExcluded(post.FieldPublic)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PostUpsertOne) UpdateNewValues() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(post.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PostUpsertOne) Ignore() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PostUpsertOne) DoNothing() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PostCreate.OnConflict
// documentation for more info.
func (u *PostUpsertOne) Update(set func(*PostUpsert)) *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PostUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *PostUpsertOne) SetUpdateTime(v time.Time) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateUpdateTime() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetSlug sets the "slug" field.
func (u *PostUpsertOne) SetSlug(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetSlug(v)
	})
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateSlug() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateSlug()
	})
}

// SetTitle sets the "title" field.
func (u *PostUpsertOne) SetTitle(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateTitle() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *PostUpsertOne) SetContent(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateContent() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateContent()
	})
}

// SetContentHTML sets the "content_html" field.
func (u *PostUpsertOne) SetContentHTML(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetContentHTML(v)
	})
}

// UpdateContentHTML sets the "content_html" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateContentHTML() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateContentHTML()
	})
}

// SetSummary sets the "summary" field.
func (u *PostUpsertOne) SetSummary(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetSummary(v)
	})
}

// UpdateSummary sets the "summary" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateSummary() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateSummary()
	})
}

// SetPublishedAt sets the "published_at" field.
func (u *PostUpsertOne) SetPublishedAt(v time.Time) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetPublishedAt(v)
	})
}

// UpdatePublishedAt sets the "published_at" field to the value that was provided on create.
func (u *PostUpsertOne) UpdatePublishedAt() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdatePublishedAt()
	})
}

// SetViewCount sets the "view_count" field.
func (u *PostUpsertOne) SetViewCount(v int) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetViewCount(v)
	})
}

// AddViewCount adds v to the "view_count" field.
func (u *PostUpsertOne) AddViewCount(v int) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.AddViewCount(v)
	})
}

// UpdateViewCount sets the "view_count" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateViewCount() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateViewCount()
	})
}

// SetPublic sets the "public" field.
func (u *PostUpsertOne) SetPublic(v bool) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *PostUpsertOne) UpdatePublic() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdatePublic()
	})
}

// Exec executes the query.
func (u *PostUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PostCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PostUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PostUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PostUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	builders []*PostCreate
	conflict []sql.ConflictOption
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Post.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PostUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pcb *PostCreateBulk) OnConflict(opts ...sql.ConflictOption) *PostUpsertBulk {
	pcb.conflict = opts
	return &PostUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PostCreateBulk) OnConflictColumns(columns ...string) *PostUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PostUpsertBulk{
		create: pcb,
	}
}

// PostUpsertBulk is the builder for "upsert"-ing
// a bulk of Post nodes.
type PostUpsertBulk struct {
	create *PostCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PostUpsertBulk) UpdateNewValues() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(post.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PostUpsertBulk) Ignore() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PostUpsertBulk) DoNothing() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PostCreateBulk.OnConflict
// documentation for more info.
func (u *PostUpsertBulk) Update(set func(*PostUpsert)) *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PostUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *PostUpsertBulk) SetUpdateTime(v time.Time) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateUpdateTime() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetSlug sets the "slug" field.
func (u *PostUpsertBulk) SetSlug(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetSlug(v)
	})
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateSlug() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateSlug()
	})
}

// SetTitle sets the "title" field.
func (u *PostUpsertBulk) SetTitle(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateTitle() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *PostUpsertBulk) SetContent(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateContent() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateContent()
	})
}

// SetContentHTML sets the "content_html" field.
func (u *PostUpsertBulk) SetContentHTML(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetContentHTML(v)
	})
}

// UpdateContentHTML sets the "content_html" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateContentHTML() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateContentHTML()
	})
}

// SetSummary sets the "summary" field.
func (u *PostUpsertBulk) SetSummary(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetSummary(v)
	})
}

// UpdateSummary sets the "summary" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateSummary() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateSummary()
	})
}

// SetPublishedAt sets the "published_at" field.
func (u *PostUpsertBulk) SetPublishedAt(v time.Time) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetPublishedAt(v)
	})
}

// UpdatePublishedAt sets the "published_at" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdatePublishedAt() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdatePublishedAt()
	})
}

// SetViewCount sets the "view_count" field.
func (u *PostUpsertBulk) SetViewCount(v int) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetViewCount(v)
	})
}

// AddViewCount adds v to the "view_count" field.
func (u *PostUpsertBulk) AddViewCount(v int) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.AddViewCount(v)
	})
}

// UpdateViewCount sets the "view_count" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateViewCount() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateViewCount()
	})
}

// SetPublic sets the "public" field.
func (u *PostUpsertBulk) SetPublic(v bool) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdatePublic() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdatePublic()
	})
}

// Exec executes the query.
func (u *PostUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PostCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PostCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PostUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
