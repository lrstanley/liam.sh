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
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/liam.sh/internal/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// GithubAssetQuery is the builder for querying GithubAsset entities.
type GithubAssetQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.GithubAsset
	withRelease *GithubReleaseQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*GithubAsset) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GithubAssetQuery builder.
func (gaq *GithubAssetQuery) Where(ps ...predicate.GithubAsset) *GithubAssetQuery {
	gaq.predicates = append(gaq.predicates, ps...)
	return gaq
}

// Limit the number of records to be returned by this query.
func (gaq *GithubAssetQuery) Limit(limit int) *GithubAssetQuery {
	gaq.ctx.Limit = &limit
	return gaq
}

// Offset to start from.
func (gaq *GithubAssetQuery) Offset(offset int) *GithubAssetQuery {
	gaq.ctx.Offset = &offset
	return gaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gaq *GithubAssetQuery) Unique(unique bool) *GithubAssetQuery {
	gaq.ctx.Unique = &unique
	return gaq
}

// Order specifies how the records should be ordered.
func (gaq *GithubAssetQuery) Order(o ...OrderFunc) *GithubAssetQuery {
	gaq.order = append(gaq.order, o...)
	return gaq
}

// QueryRelease chains the current query on the "release" edge.
func (gaq *GithubAssetQuery) QueryRelease() *GithubReleaseQuery {
	query := (&GithubReleaseClient{config: gaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(githubasset.Table, githubasset.FieldID, selector),
			sqlgraph.To(githubrelease.Table, githubrelease.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, githubasset.ReleaseTable, githubasset.ReleaseColumn),
		)
		fromU = sqlgraph.SetNeighbors(gaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GithubAsset entity from the query.
// Returns a *NotFoundError when no GithubAsset was found.
func (gaq *GithubAssetQuery) First(ctx context.Context) (*GithubAsset, error) {
	nodes, err := gaq.Limit(1).All(setContextOp(ctx, gaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{githubasset.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gaq *GithubAssetQuery) FirstX(ctx context.Context) *GithubAsset {
	node, err := gaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GithubAsset ID from the query.
// Returns a *NotFoundError when no GithubAsset ID was found.
func (gaq *GithubAssetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gaq.Limit(1).IDs(setContextOp(ctx, gaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{githubasset.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gaq *GithubAssetQuery) FirstIDX(ctx context.Context) int {
	id, err := gaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GithubAsset entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GithubAsset entity is found.
// Returns a *NotFoundError when no GithubAsset entities are found.
func (gaq *GithubAssetQuery) Only(ctx context.Context) (*GithubAsset, error) {
	nodes, err := gaq.Limit(2).All(setContextOp(ctx, gaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{githubasset.Label}
	default:
		return nil, &NotSingularError{githubasset.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gaq *GithubAssetQuery) OnlyX(ctx context.Context) *GithubAsset {
	node, err := gaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GithubAsset ID in the query.
// Returns a *NotSingularError when more than one GithubAsset ID is found.
// Returns a *NotFoundError when no entities are found.
func (gaq *GithubAssetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gaq.Limit(2).IDs(setContextOp(ctx, gaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{githubasset.Label}
	default:
		err = &NotSingularError{githubasset.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gaq *GithubAssetQuery) OnlyIDX(ctx context.Context) int {
	id, err := gaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GithubAssets.
func (gaq *GithubAssetQuery) All(ctx context.Context) ([]*GithubAsset, error) {
	ctx = setContextOp(ctx, gaq.ctx, "All")
	if err := gaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GithubAsset, *GithubAssetQuery]()
	return withInterceptors[[]*GithubAsset](ctx, gaq, qr, gaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gaq *GithubAssetQuery) AllX(ctx context.Context) []*GithubAsset {
	nodes, err := gaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GithubAsset IDs.
func (gaq *GithubAssetQuery) IDs(ctx context.Context) (ids []int, err error) {
	if gaq.ctx.Unique == nil && gaq.path != nil {
		gaq.Unique(true)
	}
	ctx = setContextOp(ctx, gaq.ctx, "IDs")
	if err = gaq.Select(githubasset.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gaq *GithubAssetQuery) IDsX(ctx context.Context) []int {
	ids, err := gaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gaq *GithubAssetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gaq.ctx, "Count")
	if err := gaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gaq, querierCount[*GithubAssetQuery](), gaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gaq *GithubAssetQuery) CountX(ctx context.Context) int {
	count, err := gaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gaq *GithubAssetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gaq.ctx, "Exist")
	switch _, err := gaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gaq *GithubAssetQuery) ExistX(ctx context.Context) bool {
	exist, err := gaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GithubAssetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gaq *GithubAssetQuery) Clone() *GithubAssetQuery {
	if gaq == nil {
		return nil
	}
	return &GithubAssetQuery{
		config:      gaq.config,
		ctx:         gaq.ctx.Clone(),
		order:       append([]OrderFunc{}, gaq.order...),
		inters:      append([]Interceptor{}, gaq.inters...),
		predicates:  append([]predicate.GithubAsset{}, gaq.predicates...),
		withRelease: gaq.withRelease.Clone(),
		// clone intermediate query.
		sql:  gaq.sql.Clone(),
		path: gaq.path,
	}
}

// WithRelease tells the query-builder to eager-load the nodes that are connected to
// the "release" edge. The optional arguments are used to configure the query builder of the edge.
func (gaq *GithubAssetQuery) WithRelease(opts ...func(*GithubReleaseQuery)) *GithubAssetQuery {
	query := (&GithubReleaseClient{config: gaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gaq.withRelease = query
	return gaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AssetID int64 `json:"asset_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GithubAsset.Query().
//		GroupBy(githubasset.FieldAssetID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gaq *GithubAssetQuery) GroupBy(field string, fields ...string) *GithubAssetGroupBy {
	gaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GithubAssetGroupBy{build: gaq}
	grbuild.flds = &gaq.ctx.Fields
	grbuild.label = githubasset.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AssetID int64 `json:"asset_id,omitempty"`
//	}
//
//	client.GithubAsset.Query().
//		Select(githubasset.FieldAssetID).
//		Scan(ctx, &v)
func (gaq *GithubAssetQuery) Select(fields ...string) *GithubAssetSelect {
	gaq.ctx.Fields = append(gaq.ctx.Fields, fields...)
	sbuild := &GithubAssetSelect{GithubAssetQuery: gaq}
	sbuild.label = githubasset.Label
	sbuild.flds, sbuild.scan = &gaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GithubAssetSelect configured with the given aggregations.
func (gaq *GithubAssetQuery) Aggregate(fns ...AggregateFunc) *GithubAssetSelect {
	return gaq.Select().Aggregate(fns...)
}

func (gaq *GithubAssetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gaq); err != nil {
				return err
			}
		}
	}
	for _, f := range gaq.ctx.Fields {
		if !githubasset.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gaq.path != nil {
		prev, err := gaq.path(ctx)
		if err != nil {
			return err
		}
		gaq.sql = prev
	}
	if githubasset.Policy == nil {
		return errors.New("ent: uninitialized githubasset.Policy (forgotten import ent/runtime?)")
	}
	if err := githubasset.Policy.EvalQuery(ctx, gaq); err != nil {
		return err
	}
	return nil
}

func (gaq *GithubAssetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GithubAsset, error) {
	var (
		nodes       = []*GithubAsset{}
		withFKs     = gaq.withFKs
		_spec       = gaq.querySpec()
		loadedTypes = [1]bool{
			gaq.withRelease != nil,
		}
	)
	if gaq.withRelease != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, githubasset.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GithubAsset).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GithubAsset{config: gaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(gaq.modifiers) > 0 {
		_spec.Modifiers = gaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gaq.withRelease; query != nil {
		if err := gaq.loadRelease(ctx, query, nodes, nil,
			func(n *GithubAsset, e *GithubRelease) { n.Edges.Release = e }); err != nil {
			return nil, err
		}
	}
	for i := range gaq.loadTotal {
		if err := gaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gaq *GithubAssetQuery) loadRelease(ctx context.Context, query *GithubReleaseQuery, nodes []*GithubAsset, init func(*GithubAsset), assign func(*GithubAsset, *GithubRelease)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GithubAsset)
	for i := range nodes {
		if nodes[i].github_release_assets == nil {
			continue
		}
		fk := *nodes[i].github_release_assets
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(githubrelease.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "github_release_assets" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gaq *GithubAssetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gaq.querySpec()
	if len(gaq.modifiers) > 0 {
		_spec.Modifiers = gaq.modifiers
	}
	_spec.Node.Columns = gaq.ctx.Fields
	if len(gaq.ctx.Fields) > 0 {
		_spec.Unique = gaq.ctx.Unique != nil && *gaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gaq.driver, _spec)
}

func (gaq *GithubAssetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(githubasset.Table, githubasset.Columns, sqlgraph.NewFieldSpec(githubasset.FieldID, field.TypeInt))
	_spec.From = gaq.sql
	if unique := gaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gaq.path != nil {
		_spec.Unique = true
	}
	if fields := gaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githubasset.FieldID)
		for i := range fields {
			if fields[i] != githubasset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gaq *GithubAssetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gaq.driver.Dialect())
	t1 := builder.Table(githubasset.Table)
	columns := gaq.ctx.Fields
	if len(columns) == 0 {
		columns = githubasset.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gaq.sql != nil {
		selector = gaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gaq.ctx.Unique != nil && *gaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gaq.predicates {
		p(selector)
	}
	for _, p := range gaq.order {
		p(selector)
	}
	if offset := gaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GithubAssetGroupBy is the group-by builder for GithubAsset entities.
type GithubAssetGroupBy struct {
	selector
	build *GithubAssetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gagb *GithubAssetGroupBy) Aggregate(fns ...AggregateFunc) *GithubAssetGroupBy {
	gagb.fns = append(gagb.fns, fns...)
	return gagb
}

// Scan applies the selector query and scans the result into the given value.
func (gagb *GithubAssetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gagb.build.ctx, "GroupBy")
	if err := gagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GithubAssetQuery, *GithubAssetGroupBy](ctx, gagb.build, gagb, gagb.build.inters, v)
}

func (gagb *GithubAssetGroupBy) sqlScan(ctx context.Context, root *GithubAssetQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gagb.fns))
	for _, fn := range gagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gagb.flds)+len(gagb.fns))
		for _, f := range *gagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GithubAssetSelect is the builder for selecting fields of GithubAsset entities.
type GithubAssetSelect struct {
	*GithubAssetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gas *GithubAssetSelect) Aggregate(fns ...AggregateFunc) *GithubAssetSelect {
	gas.fns = append(gas.fns, fns...)
	return gas
}

// Scan applies the selector query and scans the result into the given value.
func (gas *GithubAssetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gas.ctx, "Select")
	if err := gas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GithubAssetQuery, *GithubAssetSelect](ctx, gas.GithubAssetQuery, gas, gas.inters, v)
}

func (gas *GithubAssetSelect) sqlScan(ctx context.Context, root *GithubAssetQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gas.fns))
	for _, fn := range gas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
