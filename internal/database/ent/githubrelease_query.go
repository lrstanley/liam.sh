// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/database/ent/predicate"
)

// GithubReleaseQuery is the builder for querying GithubRelease entities.
type GithubReleaseQuery struct {
	config
	ctx             *QueryContext
	order           []githubrelease.OrderOption
	inters          []Interceptor
	predicates      []predicate.GithubRelease
	withRepository  *GithubRepositoryQuery
	withAssets      *GithubAssetQuery
	withFKs         bool
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*GithubRelease) error
	withNamedAssets map[string]*GithubAssetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GithubReleaseQuery builder.
func (grq *GithubReleaseQuery) Where(ps ...predicate.GithubRelease) *GithubReleaseQuery {
	grq.predicates = append(grq.predicates, ps...)
	return grq
}

// Limit the number of records to be returned by this query.
func (grq *GithubReleaseQuery) Limit(limit int) *GithubReleaseQuery {
	grq.ctx.Limit = &limit
	return grq
}

// Offset to start from.
func (grq *GithubReleaseQuery) Offset(offset int) *GithubReleaseQuery {
	grq.ctx.Offset = &offset
	return grq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (grq *GithubReleaseQuery) Unique(unique bool) *GithubReleaseQuery {
	grq.ctx.Unique = &unique
	return grq
}

// Order specifies how the records should be ordered.
func (grq *GithubReleaseQuery) Order(o ...githubrelease.OrderOption) *GithubReleaseQuery {
	grq.order = append(grq.order, o...)
	return grq
}

// QueryRepository chains the current query on the "repository" edge.
func (grq *GithubReleaseQuery) QueryRepository() *GithubRepositoryQuery {
	query := (&GithubRepositoryClient{config: grq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := grq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(githubrelease.Table, githubrelease.FieldID, selector),
			sqlgraph.To(githubrepository.Table, githubrepository.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, githubrelease.RepositoryTable, githubrelease.RepositoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(grq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssets chains the current query on the "assets" edge.
func (grq *GithubReleaseQuery) QueryAssets() *GithubAssetQuery {
	query := (&GithubAssetClient{config: grq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := grq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(githubrelease.Table, githubrelease.FieldID, selector),
			sqlgraph.To(githubasset.Table, githubasset.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, githubrelease.AssetsTable, githubrelease.AssetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(grq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GithubRelease entity from the query.
// Returns a *NotFoundError when no GithubRelease was found.
func (grq *GithubReleaseQuery) First(ctx context.Context) (*GithubRelease, error) {
	nodes, err := grq.Limit(1).All(setContextOp(ctx, grq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{githubrelease.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (grq *GithubReleaseQuery) FirstX(ctx context.Context) *GithubRelease {
	node, err := grq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GithubRelease ID from the query.
// Returns a *NotFoundError when no GithubRelease ID was found.
func (grq *GithubReleaseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = grq.Limit(1).IDs(setContextOp(ctx, grq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{githubrelease.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (grq *GithubReleaseQuery) FirstIDX(ctx context.Context) int {
	id, err := grq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GithubRelease entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GithubRelease entity is found.
// Returns a *NotFoundError when no GithubRelease entities are found.
func (grq *GithubReleaseQuery) Only(ctx context.Context) (*GithubRelease, error) {
	nodes, err := grq.Limit(2).All(setContextOp(ctx, grq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{githubrelease.Label}
	default:
		return nil, &NotSingularError{githubrelease.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (grq *GithubReleaseQuery) OnlyX(ctx context.Context) *GithubRelease {
	node, err := grq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GithubRelease ID in the query.
// Returns a *NotSingularError when more than one GithubRelease ID is found.
// Returns a *NotFoundError when no entities are found.
func (grq *GithubReleaseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = grq.Limit(2).IDs(setContextOp(ctx, grq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{githubrelease.Label}
	default:
		err = &NotSingularError{githubrelease.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (grq *GithubReleaseQuery) OnlyIDX(ctx context.Context) int {
	id, err := grq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GithubReleases.
func (grq *GithubReleaseQuery) All(ctx context.Context) ([]*GithubRelease, error) {
	ctx = setContextOp(ctx, grq.ctx, "All")
	if err := grq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GithubRelease, *GithubReleaseQuery]()
	return withInterceptors[[]*GithubRelease](ctx, grq, qr, grq.inters)
}

// AllX is like All, but panics if an error occurs.
func (grq *GithubReleaseQuery) AllX(ctx context.Context) []*GithubRelease {
	nodes, err := grq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GithubRelease IDs.
func (grq *GithubReleaseQuery) IDs(ctx context.Context) (ids []int, err error) {
	if grq.ctx.Unique == nil && grq.path != nil {
		grq.Unique(true)
	}
	ctx = setContextOp(ctx, grq.ctx, "IDs")
	if err = grq.Select(githubrelease.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (grq *GithubReleaseQuery) IDsX(ctx context.Context) []int {
	ids, err := grq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (grq *GithubReleaseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, grq.ctx, "Count")
	if err := grq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, grq, querierCount[*GithubReleaseQuery](), grq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (grq *GithubReleaseQuery) CountX(ctx context.Context) int {
	count, err := grq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (grq *GithubReleaseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, grq.ctx, "Exist")
	switch _, err := grq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (grq *GithubReleaseQuery) ExistX(ctx context.Context) bool {
	exist, err := grq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GithubReleaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (grq *GithubReleaseQuery) Clone() *GithubReleaseQuery {
	if grq == nil {
		return nil
	}
	return &GithubReleaseQuery{
		config:         grq.config,
		ctx:            grq.ctx.Clone(),
		order:          append([]githubrelease.OrderOption{}, grq.order...),
		inters:         append([]Interceptor{}, grq.inters...),
		predicates:     append([]predicate.GithubRelease{}, grq.predicates...),
		withRepository: grq.withRepository.Clone(),
		withAssets:     grq.withAssets.Clone(),
		// clone intermediate query.
		sql:  grq.sql.Clone(),
		path: grq.path,
	}
}

// WithRepository tells the query-builder to eager-load the nodes that are connected to
// the "repository" edge. The optional arguments are used to configure the query builder of the edge.
func (grq *GithubReleaseQuery) WithRepository(opts ...func(*GithubRepositoryQuery)) *GithubReleaseQuery {
	query := (&GithubRepositoryClient{config: grq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	grq.withRepository = query
	return grq
}

// WithAssets tells the query-builder to eager-load the nodes that are connected to
// the "assets" edge. The optional arguments are used to configure the query builder of the edge.
func (grq *GithubReleaseQuery) WithAssets(opts ...func(*GithubAssetQuery)) *GithubReleaseQuery {
	query := (&GithubAssetClient{config: grq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	grq.withAssets = query
	return grq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ReleaseID int64 `json:"release_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GithubRelease.Query().
//		GroupBy(githubrelease.FieldReleaseID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (grq *GithubReleaseQuery) GroupBy(field string, fields ...string) *GithubReleaseGroupBy {
	grq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GithubReleaseGroupBy{build: grq}
	grbuild.flds = &grq.ctx.Fields
	grbuild.label = githubrelease.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ReleaseID int64 `json:"release_id,omitempty"`
//	}
//
//	client.GithubRelease.Query().
//		Select(githubrelease.FieldReleaseID).
//		Scan(ctx, &v)
func (grq *GithubReleaseQuery) Select(fields ...string) *GithubReleaseSelect {
	grq.ctx.Fields = append(grq.ctx.Fields, fields...)
	sbuild := &GithubReleaseSelect{GithubReleaseQuery: grq}
	sbuild.label = githubrelease.Label
	sbuild.flds, sbuild.scan = &grq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GithubReleaseSelect configured with the given aggregations.
func (grq *GithubReleaseQuery) Aggregate(fns ...AggregateFunc) *GithubReleaseSelect {
	return grq.Select().Aggregate(fns...)
}

func (grq *GithubReleaseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range grq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, grq); err != nil {
				return err
			}
		}
	}
	for _, f := range grq.ctx.Fields {
		if !githubrelease.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if grq.path != nil {
		prev, err := grq.path(ctx)
		if err != nil {
			return err
		}
		grq.sql = prev
	}
	if githubrelease.Policy == nil {
		return errors.New("ent: uninitialized githubrelease.Policy (forgotten import ent/runtime?)")
	}
	if err := githubrelease.Policy.EvalQuery(ctx, grq); err != nil {
		return err
	}
	return nil
}

func (grq *GithubReleaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GithubRelease, error) {
	var (
		nodes       = []*GithubRelease{}
		withFKs     = grq.withFKs
		_spec       = grq.querySpec()
		loadedTypes = [2]bool{
			grq.withRepository != nil,
			grq.withAssets != nil,
		}
	)
	if grq.withRepository != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, githubrelease.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GithubRelease).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GithubRelease{config: grq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(grq.modifiers) > 0 {
		_spec.Modifiers = grq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, grq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := grq.withRepository; query != nil {
		if err := grq.loadRepository(ctx, query, nodes, nil,
			func(n *GithubRelease, e *GithubRepository) { n.Edges.Repository = e }); err != nil {
			return nil, err
		}
	}
	if query := grq.withAssets; query != nil {
		if err := grq.loadAssets(ctx, query, nodes,
			func(n *GithubRelease) { n.Edges.Assets = []*GithubAsset{} },
			func(n *GithubRelease, e *GithubAsset) { n.Edges.Assets = append(n.Edges.Assets, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range grq.withNamedAssets {
		if err := grq.loadAssets(ctx, query, nodes,
			func(n *GithubRelease) { n.appendNamedAssets(name) },
			func(n *GithubRelease, e *GithubAsset) { n.appendNamedAssets(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range grq.loadTotal {
		if err := grq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (grq *GithubReleaseQuery) loadRepository(ctx context.Context, query *GithubRepositoryQuery, nodes []*GithubRelease, init func(*GithubRelease), assign func(*GithubRelease, *GithubRepository)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GithubRelease)
	for i := range nodes {
		if nodes[i].github_repository_releases == nil {
			continue
		}
		fk := *nodes[i].github_repository_releases
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(githubrepository.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "github_repository_releases" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (grq *GithubReleaseQuery) loadAssets(ctx context.Context, query *GithubAssetQuery, nodes []*GithubRelease, init func(*GithubRelease), assign func(*GithubRelease, *GithubAsset)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*GithubRelease)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.GithubAsset(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(githubrelease.AssetsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.github_release_assets
		if fk == nil {
			return fmt.Errorf(`foreign-key "github_release_assets" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "github_release_assets" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (grq *GithubReleaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := grq.querySpec()
	if len(grq.modifiers) > 0 {
		_spec.Modifiers = grq.modifiers
	}
	_spec.Node.Columns = grq.ctx.Fields
	if len(grq.ctx.Fields) > 0 {
		_spec.Unique = grq.ctx.Unique != nil && *grq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, grq.driver, _spec)
}

func (grq *GithubReleaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(githubrelease.Table, githubrelease.Columns, sqlgraph.NewFieldSpec(githubrelease.FieldID, field.TypeInt))
	_spec.From = grq.sql
	if unique := grq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if grq.path != nil {
		_spec.Unique = true
	}
	if fields := grq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githubrelease.FieldID)
		for i := range fields {
			if fields[i] != githubrelease.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := grq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := grq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := grq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := grq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (grq *GithubReleaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(grq.driver.Dialect())
	t1 := builder.Table(githubrelease.Table)
	columns := grq.ctx.Fields
	if len(columns) == 0 {
		columns = githubrelease.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if grq.sql != nil {
		selector = grq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if grq.ctx.Unique != nil && *grq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range grq.predicates {
		p(selector)
	}
	for _, p := range grq.order {
		p(selector)
	}
	if offset := grq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := grq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAssets tells the query-builder to eager-load the nodes that are connected to the "assets"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (grq *GithubReleaseQuery) WithNamedAssets(name string, opts ...func(*GithubAssetQuery)) *GithubReleaseQuery {
	query := (&GithubAssetClient{config: grq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if grq.withNamedAssets == nil {
		grq.withNamedAssets = make(map[string]*GithubAssetQuery)
	}
	grq.withNamedAssets[name] = query
	return grq
}

// GithubReleaseGroupBy is the group-by builder for GithubRelease entities.
type GithubReleaseGroupBy struct {
	selector
	build *GithubReleaseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (grgb *GithubReleaseGroupBy) Aggregate(fns ...AggregateFunc) *GithubReleaseGroupBy {
	grgb.fns = append(grgb.fns, fns...)
	return grgb
}

// Scan applies the selector query and scans the result into the given value.
func (grgb *GithubReleaseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, grgb.build.ctx, "GroupBy")
	if err := grgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GithubReleaseQuery, *GithubReleaseGroupBy](ctx, grgb.build, grgb, grgb.build.inters, v)
}

func (grgb *GithubReleaseGroupBy) sqlScan(ctx context.Context, root *GithubReleaseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(grgb.fns))
	for _, fn := range grgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*grgb.flds)+len(grgb.fns))
		for _, f := range *grgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*grgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := grgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GithubReleaseSelect is the builder for selecting fields of GithubRelease entities.
type GithubReleaseSelect struct {
	*GithubReleaseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (grs *GithubReleaseSelect) Aggregate(fns ...AggregateFunc) *GithubReleaseSelect {
	grs.fns = append(grs.fns, fns...)
	return grs
}

// Scan applies the selector query and scans the result into the given value.
func (grs *GithubReleaseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, grs.ctx, "Select")
	if err := grs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GithubReleaseQuery, *GithubReleaseSelect](ctx, grs.GithubReleaseQuery, grs, grs.inters, v)
}

func (grs *GithubReleaseSelect) sqlScan(ctx context.Context, root *GithubReleaseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(grs.fns))
	for _, fn := range grs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*grs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := grs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
