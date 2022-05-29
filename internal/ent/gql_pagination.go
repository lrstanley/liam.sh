// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/lrstanley/liam.sh/internal/ent/githubevent"
	"github.com/lrstanley/liam.sh/internal/ent/label"
	"github.com/lrstanley/liam.sh/internal/ent/post"
	"github.com/lrstanley/liam.sh/internal/ent/user"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Name == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// GithubEventEdge is the edge representation of GithubEvent.
type GithubEventEdge struct {
	Node   *GithubEvent `json:"node"`
	Cursor Cursor       `json:"cursor"`
}

// GithubEventConnection is the connection containing edges to GithubEvent.
type GithubEventConnection struct {
	Edges      []*GithubEventEdge `json:"edges"`
	PageInfo   PageInfo           `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

func (c *GithubEventConnection) build(nodes []*GithubEvent, pager *githubeventPager, first, last *int) {
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *GithubEvent
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *GithubEvent {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *GithubEvent {
			return nodes[i]
		}
	}
	c.Edges = make([]*GithubEventEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &GithubEventEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// GithubEventPaginateOption enables pagination customization.
type GithubEventPaginateOption func(*githubeventPager) error

// WithGithubEventOrder configures pagination ordering.
func WithGithubEventOrder(order *GithubEventOrder) GithubEventPaginateOption {
	if order == nil {
		order = DefaultGithubEventOrder
	}
	o := *order
	return func(pager *githubeventPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultGithubEventOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithGithubEventFilter configures pagination filter.
func WithGithubEventFilter(filter func(*GithubEventQuery) (*GithubEventQuery, error)) GithubEventPaginateOption {
	return func(pager *githubeventPager) error {
		if filter == nil {
			return errors.New("GithubEventQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type githubeventPager struct {
	order  *GithubEventOrder
	filter func(*GithubEventQuery) (*GithubEventQuery, error)
}

func newGithubEventPager(opts []GithubEventPaginateOption) (*githubeventPager, error) {
	pager := &githubeventPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultGithubEventOrder
	}
	return pager, nil
}

func (p *githubeventPager) applyFilter(query *GithubEventQuery) (*GithubEventQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *githubeventPager) toCursor(ge *GithubEvent) Cursor {
	return p.order.Field.toCursor(ge)
}

func (p *githubeventPager) applyCursors(query *GithubEventQuery, after, before *Cursor) *GithubEventQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultGithubEventOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *githubeventPager) applyOrder(query *GithubEventQuery, reverse bool) *GithubEventQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultGithubEventOrder.Field {
		query = query.Order(direction.orderFunc(DefaultGithubEventOrder.Field.field))
	}
	return query
}

func (p *githubeventPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultGithubEventOrder.Field {
			b.Comma().Ident(DefaultGithubEventOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to GithubEvent.
func (ge *GithubEventQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...GithubEventPaginateOption,
) (*GithubEventConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newGithubEventPager(opts)
	if err != nil {
		return nil, err
	}
	if ge, err = pager.applyFilter(ge); err != nil {
		return nil, err
	}
	conn := &GithubEventConnection{Edges: []*GithubEventEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
			if conn.TotalCount, err = ge.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := ge.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	ge = pager.applyCursors(ge, after, before)
	ge = pager.applyOrder(ge, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		ge.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := ge.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := ge.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}
	conn.build(nodes, pager, first, last)
	return conn, nil
}

var (
	// GithubEventOrderFieldEventID orders GithubEvent by event_id.
	GithubEventOrderFieldEventID = &GithubEventOrderField{
		field: githubevent.FieldEventID,
		toCursor: func(ge *GithubEvent) Cursor {
			return Cursor{
				ID:    ge.ID,
				Value: ge.EventID,
			}
		},
	}
	// GithubEventOrderFieldEventType orders GithubEvent by event_type.
	GithubEventOrderFieldEventType = &GithubEventOrderField{
		field: githubevent.FieldEventType,
		toCursor: func(ge *GithubEvent) Cursor {
			return Cursor{
				ID:    ge.ID,
				Value: ge.EventType,
			}
		},
	}
	// GithubEventOrderFieldCreatedAt orders GithubEvent by created_at.
	GithubEventOrderFieldCreatedAt = &GithubEventOrderField{
		field: githubevent.FieldCreatedAt,
		toCursor: func(ge *GithubEvent) Cursor {
			return Cursor{
				ID:    ge.ID,
				Value: ge.CreatedAt,
			}
		},
	}
	// GithubEventOrderFieldActorID orders GithubEvent by actor_id.
	GithubEventOrderFieldActorID = &GithubEventOrderField{
		field: githubevent.FieldActorID,
		toCursor: func(ge *GithubEvent) Cursor {
			return Cursor{
				ID:    ge.ID,
				Value: ge.ActorID,
			}
		},
	}
	// GithubEventOrderFieldRepoID orders GithubEvent by repo_id.
	GithubEventOrderFieldRepoID = &GithubEventOrderField{
		field: githubevent.FieldRepoID,
		toCursor: func(ge *GithubEvent) Cursor {
			return Cursor{
				ID:    ge.ID,
				Value: ge.RepoID,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f GithubEventOrderField) String() string {
	var str string
	switch f.field {
	case githubevent.FieldEventID:
		str = "EVENT_ID"
	case githubevent.FieldEventType:
		str = "EVENT_TYPE"
	case githubevent.FieldCreatedAt:
		str = "CREATED_AT"
	case githubevent.FieldActorID:
		str = "ACTOR_ID"
	case githubevent.FieldRepoID:
		str = "REPO_ID"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f GithubEventOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *GithubEventOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("GithubEventOrderField %T must be a string", v)
	}
	switch str {
	case "EVENT_ID":
		*f = *GithubEventOrderFieldEventID
	case "EVENT_TYPE":
		*f = *GithubEventOrderFieldEventType
	case "CREATED_AT":
		*f = *GithubEventOrderFieldCreatedAt
	case "ACTOR_ID":
		*f = *GithubEventOrderFieldActorID
	case "REPO_ID":
		*f = *GithubEventOrderFieldRepoID
	default:
		return fmt.Errorf("%s is not a valid GithubEventOrderField", str)
	}
	return nil
}

// GithubEventOrderField defines the ordering field of GithubEvent.
type GithubEventOrderField struct {
	field    string
	toCursor func(*GithubEvent) Cursor
}

// GithubEventOrder defines the ordering of GithubEvent.
type GithubEventOrder struct {
	Direction OrderDirection         `json:"direction"`
	Field     *GithubEventOrderField `json:"field"`
}

// DefaultGithubEventOrder is the default ordering of GithubEvent.
var DefaultGithubEventOrder = &GithubEventOrder{
	Direction: OrderDirectionAsc,
	Field: &GithubEventOrderField{
		field: githubevent.FieldID,
		toCursor: func(ge *GithubEvent) Cursor {
			return Cursor{ID: ge.ID}
		},
	},
}

// ToEdge converts GithubEvent into GithubEventEdge.
func (ge *GithubEvent) ToEdge(order *GithubEventOrder) *GithubEventEdge {
	if order == nil {
		order = DefaultGithubEventOrder
	}
	return &GithubEventEdge{
		Node:   ge,
		Cursor: order.Field.toCursor(ge),
	}
}

// LabelEdge is the edge representation of Label.
type LabelEdge struct {
	Node   *Label `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// LabelConnection is the connection containing edges to Label.
type LabelConnection struct {
	Edges      []*LabelEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *LabelConnection) build(nodes []*Label, pager *labelPager, first, last *int) {
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Label
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Label {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Label {
			return nodes[i]
		}
	}
	c.Edges = make([]*LabelEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &LabelEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// LabelPaginateOption enables pagination customization.
type LabelPaginateOption func(*labelPager) error

// WithLabelOrder configures pagination ordering.
func WithLabelOrder(order *LabelOrder) LabelPaginateOption {
	if order == nil {
		order = DefaultLabelOrder
	}
	o := *order
	return func(pager *labelPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultLabelOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithLabelFilter configures pagination filter.
func WithLabelFilter(filter func(*LabelQuery) (*LabelQuery, error)) LabelPaginateOption {
	return func(pager *labelPager) error {
		if filter == nil {
			return errors.New("LabelQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type labelPager struct {
	order  *LabelOrder
	filter func(*LabelQuery) (*LabelQuery, error)
}

func newLabelPager(opts []LabelPaginateOption) (*labelPager, error) {
	pager := &labelPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultLabelOrder
	}
	return pager, nil
}

func (p *labelPager) applyFilter(query *LabelQuery) (*LabelQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *labelPager) toCursor(l *Label) Cursor {
	return p.order.Field.toCursor(l)
}

func (p *labelPager) applyCursors(query *LabelQuery, after, before *Cursor) *LabelQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultLabelOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *labelPager) applyOrder(query *LabelQuery, reverse bool) *LabelQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultLabelOrder.Field {
		query = query.Order(direction.orderFunc(DefaultLabelOrder.Field.field))
	}
	return query
}

func (p *labelPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultLabelOrder.Field {
			b.Comma().Ident(DefaultLabelOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Label.
func (l *LabelQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...LabelPaginateOption,
) (*LabelConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newLabelPager(opts)
	if err != nil {
		return nil, err
	}
	if l, err = pager.applyFilter(l); err != nil {
		return nil, err
	}
	conn := &LabelConnection{Edges: []*LabelEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
			if conn.TotalCount, err = l.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := l.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	l = pager.applyCursors(l, after, before)
	l = pager.applyOrder(l, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		l.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := l.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := l.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}
	conn.build(nodes, pager, first, last)
	return conn, nil
}

var (
	// LabelOrderFieldName orders Label by name.
	LabelOrderFieldName = &LabelOrderField{
		field: label.FieldName,
		toCursor: func(l *Label) Cursor {
			return Cursor{
				ID:    l.ID,
				Value: l.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f LabelOrderField) String() string {
	var str string
	switch f.field {
	case label.FieldName:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f LabelOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *LabelOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("LabelOrderField %T must be a string", v)
	}
	switch str {
	case "NAME":
		*f = *LabelOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid LabelOrderField", str)
	}
	return nil
}

// LabelOrderField defines the ordering field of Label.
type LabelOrderField struct {
	field    string
	toCursor func(*Label) Cursor
}

// LabelOrder defines the ordering of Label.
type LabelOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *LabelOrderField `json:"field"`
}

// DefaultLabelOrder is the default ordering of Label.
var DefaultLabelOrder = &LabelOrder{
	Direction: OrderDirectionAsc,
	Field: &LabelOrderField{
		field: label.FieldID,
		toCursor: func(l *Label) Cursor {
			return Cursor{ID: l.ID}
		},
	},
}

// ToEdge converts Label into LabelEdge.
func (l *Label) ToEdge(order *LabelOrder) *LabelEdge {
	if order == nil {
		order = DefaultLabelOrder
	}
	return &LabelEdge{
		Node:   l,
		Cursor: order.Field.toCursor(l),
	}
}

// PostEdge is the edge representation of Post.
type PostEdge struct {
	Node   *Post  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// PostConnection is the connection containing edges to Post.
type PostConnection struct {
	Edges      []*PostEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *PostConnection) build(nodes []*Post, pager *postPager, first, last *int) {
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Post
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Post {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Post {
			return nodes[i]
		}
	}
	c.Edges = make([]*PostEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &PostEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// PostPaginateOption enables pagination customization.
type PostPaginateOption func(*postPager) error

// WithPostOrder configures pagination ordering.
func WithPostOrder(order *PostOrder) PostPaginateOption {
	if order == nil {
		order = DefaultPostOrder
	}
	o := *order
	return func(pager *postPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultPostOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithPostFilter configures pagination filter.
func WithPostFilter(filter func(*PostQuery) (*PostQuery, error)) PostPaginateOption {
	return func(pager *postPager) error {
		if filter == nil {
			return errors.New("PostQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type postPager struct {
	order  *PostOrder
	filter func(*PostQuery) (*PostQuery, error)
}

func newPostPager(opts []PostPaginateOption) (*postPager, error) {
	pager := &postPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultPostOrder
	}
	return pager, nil
}

func (p *postPager) applyFilter(query *PostQuery) (*PostQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *postPager) toCursor(po *Post) Cursor {
	return p.order.Field.toCursor(po)
}

func (p *postPager) applyCursors(query *PostQuery, after, before *Cursor) *PostQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultPostOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *postPager) applyOrder(query *PostQuery, reverse bool) *PostQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultPostOrder.Field {
		query = query.Order(direction.orderFunc(DefaultPostOrder.Field.field))
	}
	return query
}

func (p *postPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultPostOrder.Field {
			b.Comma().Ident(DefaultPostOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Post.
func (po *PostQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...PostPaginateOption,
) (*PostConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newPostPager(opts)
	if err != nil {
		return nil, err
	}
	if po, err = pager.applyFilter(po); err != nil {
		return nil, err
	}
	conn := &PostConnection{Edges: []*PostEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
			if conn.TotalCount, err = po.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := po.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	po = pager.applyCursors(po, after, before)
	po = pager.applyOrder(po, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		po.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := po.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := po.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}
	conn.build(nodes, pager, first, last)
	return conn, nil
}

var (
	// PostOrderFieldSlug orders Post by slug.
	PostOrderFieldSlug = &PostOrderField{
		field: post.FieldSlug,
		toCursor: func(po *Post) Cursor {
			return Cursor{
				ID:    po.ID,
				Value: po.Slug,
			}
		},
	}
	// PostOrderFieldTitle orders Post by title.
	PostOrderFieldTitle = &PostOrderField{
		field: post.FieldTitle,
		toCursor: func(po *Post) Cursor {
			return Cursor{
				ID:    po.ID,
				Value: po.Title,
			}
		},
	}
	// PostOrderFieldPublishedAt orders Post by published_at.
	PostOrderFieldPublishedAt = &PostOrderField{
		field: post.FieldPublishedAt,
		toCursor: func(po *Post) Cursor {
			return Cursor{
				ID:    po.ID,
				Value: po.PublishedAt,
			}
		},
	}
	// PostOrderFieldViewCount orders Post by view_count.
	PostOrderFieldViewCount = &PostOrderField{
		field: post.FieldViewCount,
		toCursor: func(po *Post) Cursor {
			return Cursor{
				ID:    po.ID,
				Value: po.ViewCount,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f PostOrderField) String() string {
	var str string
	switch f.field {
	case post.FieldSlug:
		str = "SLUG"
	case post.FieldTitle:
		str = "TITLE"
	case post.FieldPublishedAt:
		str = "DATE"
	case post.FieldViewCount:
		str = "VIEW_COUNT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f PostOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *PostOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("PostOrderField %T must be a string", v)
	}
	switch str {
	case "SLUG":
		*f = *PostOrderFieldSlug
	case "TITLE":
		*f = *PostOrderFieldTitle
	case "DATE":
		*f = *PostOrderFieldPublishedAt
	case "VIEW_COUNT":
		*f = *PostOrderFieldViewCount
	default:
		return fmt.Errorf("%s is not a valid PostOrderField", str)
	}
	return nil
}

// PostOrderField defines the ordering field of Post.
type PostOrderField struct {
	field    string
	toCursor func(*Post) Cursor
}

// PostOrder defines the ordering of Post.
type PostOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *PostOrderField `json:"field"`
}

// DefaultPostOrder is the default ordering of Post.
var DefaultPostOrder = &PostOrder{
	Direction: OrderDirectionAsc,
	Field: &PostOrderField{
		field: post.FieldID,
		toCursor: func(po *Post) Cursor {
			return Cursor{ID: po.ID}
		},
	},
}

// ToEdge converts Post into PostEdge.
func (po *Post) ToEdge(order *PostOrder) *PostEdge {
	if order == nil {
		order = DefaultPostOrder
	}
	return &PostEdge{
		Node:   po,
		Cursor: order.Field.toCursor(po),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, first, last *int) {
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) *UserQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultUserOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *userPager) applyOrder(query *UserQuery, reverse bool) *UserQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}
	return query
}

func (p *userPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
			if conn.TotalCount, err = u.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := u.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	u = pager.applyCursors(u, after, before)
	u = pager.applyOrder(u, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := u.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}
	conn.build(nodes, pager, first, last)
	return conn, nil
}

var (
	// UserOrderFieldLogin orders User by login.
	UserOrderFieldLogin = &UserOrderField{
		field: user.FieldLogin,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Login,
			}
		},
	}
	// UserOrderFieldName orders User by name.
	UserOrderFieldName = &UserOrderField{
		field: user.FieldName,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Name,
			}
		},
	}
	// UserOrderFieldEmail orders User by email.
	UserOrderFieldEmail = &UserOrderField{
		field: user.FieldEmail,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Email,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserOrderField) String() string {
	var str string
	switch f.field {
	case user.FieldLogin:
		str = "LOGIN"
	case user.FieldName:
		str = "NAME"
	case user.FieldEmail:
		str = "EMAIL"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserOrderField %T must be a string", v)
	}
	switch str {
	case "LOGIN":
		*f = *UserOrderFieldLogin
	case "NAME":
		*f = *UserOrderFieldName
	case "EMAIL":
		*f = *UserOrderFieldEmail
	default:
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
