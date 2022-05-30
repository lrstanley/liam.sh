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
	"github.com/lrstanley/liam.sh/internal/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/ent/label"
	"github.com/lrstanley/liam.sh/internal/ent/predicate"
)

// GithubRepositoryQuery is the builder for querying GithubRepository entities.
type GithubRepositoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GithubRepository
	// eager-loading edges.
	withLabels *LabelQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*GithubRepository) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GithubRepositoryQuery builder.
func (grq *GithubRepositoryQuery) Where(ps ...predicate.GithubRepository) *GithubRepositoryQuery {
	grq.predicates = append(grq.predicates, ps...)
	return grq
}

// Limit adds a limit step to the query.
func (grq *GithubRepositoryQuery) Limit(limit int) *GithubRepositoryQuery {
	grq.limit = &limit
	return grq
}

// Offset adds an offset step to the query.
func (grq *GithubRepositoryQuery) Offset(offset int) *GithubRepositoryQuery {
	grq.offset = &offset
	return grq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (grq *GithubRepositoryQuery) Unique(unique bool) *GithubRepositoryQuery {
	grq.unique = &unique
	return grq
}

// Order adds an order step to the query.
func (grq *GithubRepositoryQuery) Order(o ...OrderFunc) *GithubRepositoryQuery {
	grq.order = append(grq.order, o...)
	return grq
}

// QueryLabels chains the current query on the "labels" edge.
func (grq *GithubRepositoryQuery) QueryLabels() *LabelQuery {
	query := &LabelQuery{config: grq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := grq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(githubrepository.Table, githubrepository.FieldID, selector),
			sqlgraph.To(label.Table, label.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, githubrepository.LabelsTable, githubrepository.LabelsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(grq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GithubRepository entity from the query.
// Returns a *NotFoundError when no GithubRepository was found.
func (grq *GithubRepositoryQuery) First(ctx context.Context) (*GithubRepository, error) {
	nodes, err := grq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{githubrepository.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (grq *GithubRepositoryQuery) FirstX(ctx context.Context) *GithubRepository {
	node, err := grq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GithubRepository ID from the query.
// Returns a *NotFoundError when no GithubRepository ID was found.
func (grq *GithubRepositoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = grq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{githubrepository.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (grq *GithubRepositoryQuery) FirstIDX(ctx context.Context) int {
	id, err := grq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GithubRepository entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GithubRepository entity is found.
// Returns a *NotFoundError when no GithubRepository entities are found.
func (grq *GithubRepositoryQuery) Only(ctx context.Context) (*GithubRepository, error) {
	nodes, err := grq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{githubrepository.Label}
	default:
		return nil, &NotSingularError{githubrepository.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (grq *GithubRepositoryQuery) OnlyX(ctx context.Context) *GithubRepository {
	node, err := grq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GithubRepository ID in the query.
// Returns a *NotSingularError when more than one GithubRepository ID is found.
// Returns a *NotFoundError when no entities are found.
func (grq *GithubRepositoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = grq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{githubrepository.Label}
	default:
		err = &NotSingularError{githubrepository.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (grq *GithubRepositoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := grq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GithubRepositories.
func (grq *GithubRepositoryQuery) All(ctx context.Context) ([]*GithubRepository, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return grq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (grq *GithubRepositoryQuery) AllX(ctx context.Context) []*GithubRepository {
	nodes, err := grq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GithubRepository IDs.
func (grq *GithubRepositoryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := grq.Select(githubrepository.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (grq *GithubRepositoryQuery) IDsX(ctx context.Context) []int {
	ids, err := grq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (grq *GithubRepositoryQuery) Count(ctx context.Context) (int, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return grq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (grq *GithubRepositoryQuery) CountX(ctx context.Context) int {
	count, err := grq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (grq *GithubRepositoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return grq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (grq *GithubRepositoryQuery) ExistX(ctx context.Context) bool {
	exist, err := grq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GithubRepositoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (grq *GithubRepositoryQuery) Clone() *GithubRepositoryQuery {
	if grq == nil {
		return nil
	}
	return &GithubRepositoryQuery{
		config:     grq.config,
		limit:      grq.limit,
		offset:     grq.offset,
		order:      append([]OrderFunc{}, grq.order...),
		predicates: append([]predicate.GithubRepository{}, grq.predicates...),
		withLabels: grq.withLabels.Clone(),
		// clone intermediate query.
		sql:    grq.sql.Clone(),
		path:   grq.path,
		unique: grq.unique,
	}
}

// WithLabels tells the query-builder to eager-load the nodes that are connected to
// the "labels" edge. The optional arguments are used to configure the query builder of the edge.
func (grq *GithubRepositoryQuery) WithLabels(opts ...func(*LabelQuery)) *GithubRepositoryQuery {
	query := &LabelQuery{config: grq.config}
	for _, opt := range opts {
		opt(query)
	}
	grq.withLabels = query
	return grq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RepoID int64 `json:"repo_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GithubRepository.Query().
//		GroupBy(githubrepository.FieldRepoID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (grq *GithubRepositoryQuery) GroupBy(field string, fields ...string) *GithubRepositoryGroupBy {
	grbuild := &GithubRepositoryGroupBy{config: grq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return grq.sqlQuery(ctx), nil
	}
	grbuild.label = githubrepository.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RepoID int64 `json:"repo_id,omitempty"`
//	}
//
//	client.GithubRepository.Query().
//		Select(githubrepository.FieldRepoID).
//		Scan(ctx, &v)
//
func (grq *GithubRepositoryQuery) Select(fields ...string) *GithubRepositorySelect {
	grq.fields = append(grq.fields, fields...)
	selbuild := &GithubRepositorySelect{GithubRepositoryQuery: grq}
	selbuild.label = githubrepository.Label
	selbuild.flds, selbuild.scan = &grq.fields, selbuild.Scan
	return selbuild
}

func (grq *GithubRepositoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range grq.fields {
		if !githubrepository.ValidColumn(f) {
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
	if githubrepository.Policy == nil {
		return errors.New("ent: uninitialized githubrepository.Policy (forgotten import ent/runtime?)")
	}
	if err := githubrepository.Policy.EvalQuery(ctx, grq); err != nil {
		return err
	}
	return nil
}

func (grq *GithubRepositoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GithubRepository, error) {
	var (
		nodes       = []*GithubRepository{}
		_spec       = grq.querySpec()
		loadedTypes = [1]bool{
			grq.withLabels != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*GithubRepository).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &GithubRepository{config: grq.config}
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

	if query := grq.withLabels; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[int]*GithubRepository)
		nids := make(map[int]map[*GithubRepository]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Labels = []*Label{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(githubrepository.LabelsTable)
			s.Join(joinT).On(s.C(label.FieldID), joinT.C(githubrepository.LabelsPrimaryKey[0]))
			s.Where(sql.InValues(joinT.C(githubrepository.LabelsPrimaryKey[1]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(githubrepository.LabelsPrimaryKey[1]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*GithubRepository]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "labels" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Labels = append(kn.Edges.Labels, n)
			}
		}
	}

	for i := range grq.loadTotal {
		if err := grq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (grq *GithubRepositoryQuery) sqlCount(ctx context.Context) (int, error) {
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

func (grq *GithubRepositoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := grq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (grq *GithubRepositoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   githubrepository.Table,
			Columns: githubrepository.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githubrepository.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, githubrepository.FieldID)
		for i := range fields {
			if fields[i] != githubrepository.FieldID {
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

func (grq *GithubRepositoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(grq.driver.Dialect())
	t1 := builder.Table(githubrepository.Table)
	columns := grq.fields
	if len(columns) == 0 {
		columns = githubrepository.Columns
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

// GithubRepositoryGroupBy is the group-by builder for GithubRepository entities.
type GithubRepositoryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (grgb *GithubRepositoryGroupBy) Aggregate(fns ...AggregateFunc) *GithubRepositoryGroupBy {
	grgb.fns = append(grgb.fns, fns...)
	return grgb
}

// Scan applies the group-by query and scans the result into the given value.
func (grgb *GithubRepositoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := grgb.path(ctx)
	if err != nil {
		return err
	}
	grgb.sql = query
	return grgb.sqlScan(ctx, v)
}

func (grgb *GithubRepositoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range grgb.fields {
		if !githubrepository.ValidColumn(f) {
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

func (grgb *GithubRepositoryGroupBy) sqlQuery() *sql.Selector {
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

// GithubRepositorySelect is the builder for selecting fields of GithubRepository entities.
type GithubRepositorySelect struct {
	*GithubRepositoryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (grs *GithubRepositorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := grs.prepareQuery(ctx); err != nil {
		return err
	}
	grs.sql = grs.GithubRepositoryQuery.sqlQuery(ctx)
	return grs.sqlScan(ctx, v)
}

func (grs *GithubRepositorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := grs.sql.Query()
	if err := grs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
