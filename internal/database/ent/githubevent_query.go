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
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubevent"
	"github.com/lrstanley/liam.sh/internal/database/ent/predicate"
)

// GithubEventQuery is the builder for querying GithubEvent entities.
type GithubEventQuery struct {
	config
	ctx        *QueryContext
	order      []githubevent.OrderOption
	inters     []Interceptor
	predicates []predicate.GithubEvent
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GithubEventQuery builder.
func (geq *GithubEventQuery) Where(ps ...predicate.GithubEvent) *GithubEventQuery {
	geq.predicates = append(geq.predicates, ps...)
	return geq
}

// Limit the number of records to be returned by this query.
func (geq *GithubEventQuery) Limit(limit int) *GithubEventQuery {
	geq.ctx.Limit = &limit
	return geq
}

// Offset to start from.
func (geq *GithubEventQuery) Offset(offset int) *GithubEventQuery {
	geq.ctx.Offset = &offset
	return geq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (geq *GithubEventQuery) Unique(unique bool) *GithubEventQuery {
	geq.ctx.Unique = &unique
	return geq
}

// Order specifies how the records should be ordered.
func (geq *GithubEventQuery) Order(o ...githubevent.OrderOption) *GithubEventQuery {
	geq.order = append(geq.order, o...)
	return geq
}

// First returns the first GithubEvent entity from the query.
// Returns a *NotFoundError when no GithubEvent was found.
func (geq *GithubEventQuery) First(ctx context.Context) (*GithubEvent, error) {
	nodes, err := geq.Limit(1).All(setContextOp(ctx, geq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{githubevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (geq *GithubEventQuery) FirstX(ctx context.Context) *GithubEvent {
	node, err := geq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GithubEvent ID from the query.
// Returns a *NotFoundError when no GithubEvent ID was found.
func (geq *GithubEventQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = geq.Limit(1).IDs(setContextOp(ctx, geq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{githubevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (geq *GithubEventQuery) FirstIDX(ctx context.Context) int {
	id, err := geq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GithubEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GithubEvent entity is found.
// Returns a *NotFoundError when no GithubEvent entities are found.
func (geq *GithubEventQuery) Only(ctx context.Context) (*GithubEvent, error) {
	nodes, err := geq.Limit(2).All(setContextOp(ctx, geq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{githubevent.Label}
	default:
		return nil, &NotSingularError{githubevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (geq *GithubEventQuery) OnlyX(ctx context.Context) *GithubEvent {
	node, err := geq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GithubEvent ID in the query.
// Returns a *NotSingularError when more than one GithubEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (geq *GithubEventQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = geq.Limit(2).IDs(setContextOp(ctx, geq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{githubevent.Label}
	default:
		err = &NotSingularError{githubevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (geq *GithubEventQuery) OnlyIDX(ctx context.Context) int {
	id, err := geq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GithubEvents.
func (geq *GithubEventQuery) All(ctx context.Context) ([]*GithubEvent, error) {
	ctx = setContextOp(ctx, geq.ctx, ent.OpQueryAll)
	if err := geq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GithubEvent, *GithubEventQuery]()
	return withInterceptors[[]*GithubEvent](ctx, geq, qr, geq.inters)
}

// AllX is like All, but panics if an error occurs.
func (geq *GithubEventQuery) AllX(ctx context.Context) []*GithubEvent {
	nodes, err := geq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GithubEvent IDs.
func (geq *GithubEventQuery) IDs(ctx context.Context) (ids []int, err error) {
	if geq.ctx.Unique == nil && geq.path != nil {
		geq.Unique(true)
	}
	ctx = setContextOp(ctx, geq.ctx, ent.OpQueryIDs)
	if err = geq.Select(githubevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (geq *GithubEventQuery) IDsX(ctx context.Context) []int {
	ids, err := geq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (geq *GithubEventQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, geq.ctx, ent.OpQueryCount)
	if err := geq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, geq, querierCount[*GithubEventQuery](), geq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (geq *GithubEventQuery) CountX(ctx context.Context) int {
	count, err := geq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (geq *GithubEventQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, geq.ctx, ent.OpQueryExist)
	switch _, err := geq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (geq *GithubEventQuery) ExistX(ctx context.Context) bool {
	exist, err := geq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GithubEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (geq *GithubEventQuery) Clone() *GithubEventQuery {
	if geq == nil {
		return nil
	}
	return &GithubEventQuery{
		config:     geq.config,
		ctx:        geq.ctx.Clone(),
		order:      append([]githubevent.OrderOption{}, geq.order...),
		inters:     append([]Interceptor{}, geq.inters...),
		predicates: append([]predicate.GithubEvent{}, geq.predicates...),
		// clone intermediate query.
		sql:  geq.sql.Clone(),
		path: geq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EventID string `json:"event_id"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GithubEvent.Query().
//		GroupBy(githubevent.FieldEventID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (geq *GithubEventQuery) GroupBy(field string, fields ...string) *GithubEventGroupBy {
	geq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GithubEventGroupBy{build: geq}
	grbuild.flds = &geq.ctx.Fields
	grbuild.label = githubevent.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EventID string `json:"event_id"`
//	}
//
//	client.GithubEvent.Query().
//		Select(githubevent.FieldEventID).
//		Scan(ctx, &v)
func (geq *GithubEventQuery) Select(fields ...string) *GithubEventSelect {
	geq.ctx.Fields = append(geq.ctx.Fields, fields...)
	sbuild := &GithubEventSelect{GithubEventQuery: geq}
	sbuild.label = githubevent.Label
	sbuild.flds, sbuild.scan = &geq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GithubEventSelect configured with the given aggregations.
func (geq *GithubEventQuery) Aggregate(fns ...AggregateFunc) *GithubEventSelect {
	return geq.Select().Aggregate(fns...)
}

func (geq *GithubEventQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range geq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, geq); err != nil {
				return err
			}
		}
	}
	for _, f := range geq.ctx.Fields {
		if !githubevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if geq.path != nil {
		prev, err := geq.path(ctx)
		if err != nil {
			return err
		}
		geq.sql = prev
	}
	if githubevent.Policy == nil {
		return errors.New("ent: uninitialized githubevent.Policy (forgotten import ent/runtime?)")
	}
	if err := githubevent.Policy.EvalQuery(ctx, geq); err != nil {
		return err
	}
	return nil
}

func (geq *GithubEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GithubEvent, error) {
	var (
		nodes = []*GithubEvent{}
		_spec = geq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GithubEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GithubEvent{config: geq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, geq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (geq *GithubEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := geq.querySpec()
	_spec.Node.Columns = geq.ctx.Fields
	if len(geq.ctx.Fields) > 0 {
		_spec.Unique = geq.ctx.Unique != nil && *geq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, geq.driver, _spec)
}

func (geq *GithubEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(githubevent.Table, githubevent.Columns, sqlgraph.NewFieldSpec(githubevent.FieldID, field.TypeInt))
	_spec.From = geq.sql
	if unique := geq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if geq.path != nil {
		_spec.Unique = true
	}
	if fields := geq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githubevent.FieldID)
		for i := range fields {
			if fields[i] != githubevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := geq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := geq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := geq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := geq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (geq *GithubEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(geq.driver.Dialect())
	t1 := builder.Table(githubevent.Table)
	columns := geq.ctx.Fields
	if len(columns) == 0 {
		columns = githubevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if geq.sql != nil {
		selector = geq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if geq.ctx.Unique != nil && *geq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range geq.predicates {
		p(selector)
	}
	for _, p := range geq.order {
		p(selector)
	}
	if offset := geq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := geq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GithubEventGroupBy is the group-by builder for GithubEvent entities.
type GithubEventGroupBy struct {
	selector
	build *GithubEventQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gegb *GithubEventGroupBy) Aggregate(fns ...AggregateFunc) *GithubEventGroupBy {
	gegb.fns = append(gegb.fns, fns...)
	return gegb
}

// Scan applies the selector query and scans the result into the given value.
func (gegb *GithubEventGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gegb.build.ctx, ent.OpQueryGroupBy)
	if err := gegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GithubEventQuery, *GithubEventGroupBy](ctx, gegb.build, gegb, gegb.build.inters, v)
}

func (gegb *GithubEventGroupBy) sqlScan(ctx context.Context, root *GithubEventQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gegb.fns))
	for _, fn := range gegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gegb.flds)+len(gegb.fns))
		for _, f := range *gegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GithubEventSelect is the builder for selecting fields of GithubEvent entities.
type GithubEventSelect struct {
	*GithubEventQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ges *GithubEventSelect) Aggregate(fns ...AggregateFunc) *GithubEventSelect {
	ges.fns = append(ges.fns, fns...)
	return ges
}

// Scan applies the selector query and scans the result into the given value.
func (ges *GithubEventSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ges.ctx, ent.OpQuerySelect)
	if err := ges.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GithubEventQuery, *GithubEventSelect](ctx, ges.GithubEventQuery, ges, ges.inters, v)
}

func (ges *GithubEventSelect) sqlScan(ctx context.Context, root *GithubEventQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ges.fns))
	for _, fn := range ges.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ges.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ges.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
