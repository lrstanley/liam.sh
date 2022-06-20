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
	"github.com/lrstanley/liam.sh/internal/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// GithubReleaseQuery is the builder for querying GithubRelease entities.
type GithubReleaseQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GithubRelease
	// eager-loading edges.
	withRepository *GithubRepositoryQuery
	withAssets     *GithubAssetQuery
	withFKs        bool
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*GithubRelease) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GithubReleaseQuery builder.
func (grq *GithubReleaseQuery) Where(ps ...predicate.GithubRelease) *GithubReleaseQuery {
	grq.predicates = append(grq.predicates, ps...)
	return grq
}

// Limit adds a limit step to the query.
func (grq *GithubReleaseQuery) Limit(limit int) *GithubReleaseQuery {
	grq.limit = &limit
	return grq
}

// Offset adds an offset step to the query.
func (grq *GithubReleaseQuery) Offset(offset int) *GithubReleaseQuery {
	grq.offset = &offset
	return grq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (grq *GithubReleaseQuery) Unique(unique bool) *GithubReleaseQuery {
	grq.unique = &unique
	return grq
}

// Order adds an order step to the query.
func (grq *GithubReleaseQuery) Order(o ...OrderFunc) *GithubReleaseQuery {
	grq.order = append(grq.order, o...)
	return grq
}

// QueryRepository chains the current query on the "repository" edge.
func (grq *GithubReleaseQuery) QueryRepository() *GithubRepositoryQuery {
	query := &GithubRepositoryQuery{config: grq.config}
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
	query := &GithubAssetQuery{config: grq.config}
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
	nodes, err := grq.Limit(1).All(ctx)
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
	if ids, err = grq.Limit(1).IDs(ctx); err != nil {
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
	nodes, err := grq.Limit(2).All(ctx)
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
	if ids, err = grq.Limit(2).IDs(ctx); err != nil {
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
	if err := grq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return grq.sqlAll(ctx)
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
func (grq *GithubReleaseQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := grq.Select(githubrelease.FieldID).Scan(ctx, &ids); err != nil {
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
	if err := grq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return grq.sqlCount(ctx)
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
	if err := grq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return grq.sqlExist(ctx)
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
		limit:          grq.limit,
		offset:         grq.offset,
		order:          append([]OrderFunc{}, grq.order...),
		predicates:     append([]predicate.GithubRelease{}, grq.predicates...),
		withRepository: grq.withRepository.Clone(),
		withAssets:     grq.withAssets.Clone(),
		// clone intermediate query.
		sql:    grq.sql.Clone(),
		path:   grq.path,
		unique: grq.unique,
	}
}

// WithRepository tells the query-builder to eager-load the nodes that are connected to
// the "repository" edge. The optional arguments are used to configure the query builder of the edge.
func (grq *GithubReleaseQuery) WithRepository(opts ...func(*GithubRepositoryQuery)) *GithubReleaseQuery {
	query := &GithubRepositoryQuery{config: grq.config}
	for _, opt := range opts {
		opt(query)
	}
	grq.withRepository = query
	return grq
}

// WithAssets tells the query-builder to eager-load the nodes that are connected to
// the "assets" edge. The optional arguments are used to configure the query builder of the edge.
func (grq *GithubReleaseQuery) WithAssets(opts ...func(*GithubAssetQuery)) *GithubReleaseQuery {
	query := &GithubAssetQuery{config: grq.config}
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
//
func (grq *GithubReleaseQuery) GroupBy(field string, fields ...string) *GithubReleaseGroupBy {
	grbuild := &GithubReleaseGroupBy{config: grq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return grq.sqlQuery(ctx), nil
	}
	grbuild.label = githubrelease.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
//
func (grq *GithubReleaseQuery) Select(fields ...string) *GithubReleaseSelect {
	grq.fields = append(grq.fields, fields...)
	selbuild := &GithubReleaseSelect{GithubReleaseQuery: grq}
	selbuild.label = githubrelease.Label
	selbuild.flds, selbuild.scan = &grq.fields, selbuild.Scan
	return selbuild
}

func (grq *GithubReleaseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range grq.fields {
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
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*GithubRelease).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
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
		query.Where(githubrepository.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "github_repository_releases" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Repository = n
			}
		}
	}

	if query := grq.withAssets; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*GithubRelease)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Assets = []*GithubAsset{}
		}
		query.withFKs = true
		query.Where(predicate.GithubAsset(func(s *sql.Selector) {
			s.Where(sql.InValues(githubrelease.AssetsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.github_release_assets
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "github_release_assets" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "github_release_assets" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Assets = append(node.Edges.Assets, n)
		}
	}

	for i := range grq.loadTotal {
		if err := grq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (grq *GithubReleaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := grq.querySpec()
	if len(grq.modifiers) > 0 {
		_spec.Modifiers = grq.modifiers
	}
	_spec.Node.Columns = grq.fields
	if len(grq.fields) > 0 {
		_spec.Unique = grq.unique != nil && *grq.unique
	}
	return sqlgraph.CountNodes(ctx, grq.driver, _spec)
}

func (grq *GithubReleaseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := grq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (grq *GithubReleaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   githubrelease.Table,
			Columns: githubrelease.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githubrelease.FieldID,
			},
		},
		From:   grq.sql,
		Unique: true,
	}
	if unique := grq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := grq.fields; len(fields) > 0 {
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
	if limit := grq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := grq.offset; offset != nil {
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
	columns := grq.fields
	if len(columns) == 0 {
		columns = githubrelease.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if grq.sql != nil {
		selector = grq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if grq.unique != nil && *grq.unique {
		selector.Distinct()
	}
	for _, p := range grq.predicates {
		p(selector)
	}
	for _, p := range grq.order {
		p(selector)
	}
	if offset := grq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := grq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GithubReleaseGroupBy is the group-by builder for GithubRelease entities.
type GithubReleaseGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (grgb *GithubReleaseGroupBy) Aggregate(fns ...AggregateFunc) *GithubReleaseGroupBy {
	grgb.fns = append(grgb.fns, fns...)
	return grgb
}

// Scan applies the group-by query and scans the result into the given value.
func (grgb *GithubReleaseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := grgb.path(ctx)
	if err != nil {
		return err
	}
	grgb.sql = query
	return grgb.sqlScan(ctx, v)
}

func (grgb *GithubReleaseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range grgb.fields {
		if !githubrelease.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := grgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := grgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (grgb *GithubReleaseGroupBy) sqlQuery() *sql.Selector {
	selector := grgb.sql.Select()
	aggregation := make([]string, 0, len(grgb.fns))
	for _, fn := range grgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(grgb.fields)+len(grgb.fns))
		for _, f := range grgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(grgb.fields...)...)
}

// GithubReleaseSelect is the builder for selecting fields of GithubRelease entities.
type GithubReleaseSelect struct {
	*GithubReleaseQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (grs *GithubReleaseSelect) Scan(ctx context.Context, v interface{}) error {
	if err := grs.prepareQuery(ctx); err != nil {
		return err
	}
	grs.sql = grs.GithubReleaseQuery.sqlQuery(ctx)
	return grs.sqlScan(ctx, v)
}

func (grs *GithubReleaseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := grs.sql.Query()
	if err := grs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
