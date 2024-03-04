// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
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
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/database/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/database/ent/label"
	"github.com/lrstanley/liam.sh/internal/database/ent/predicate"
)

// GithubRepositoryQuery is the builder for querying GithubRepository entities.
type GithubRepositoryQuery struct {
	config
	ctx               *QueryContext
	order             []githubrepository.OrderOption
	inters            []Interceptor
	predicates        []predicate.GithubRepository
	withLabels        *LabelQuery
	withReleases      *GithubReleaseQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*GithubRepository) error
	withNamedLabels   map[string]*LabelQuery
	withNamedReleases map[string]*GithubReleaseQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GithubRepositoryQuery builder.
func (grq *GithubRepositoryQuery) Where(ps ...predicate.GithubRepository) *GithubRepositoryQuery {
	grq.predicates = append(grq.predicates, ps...)
	return grq
}

// Limit the number of records to be returned by this query.
func (grq *GithubRepositoryQuery) Limit(limit int) *GithubRepositoryQuery {
	grq.ctx.Limit = &limit
	return grq
}

// Offset to start from.
func (grq *GithubRepositoryQuery) Offset(offset int) *GithubRepositoryQuery {
	grq.ctx.Offset = &offset
	return grq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (grq *GithubRepositoryQuery) Unique(unique bool) *GithubRepositoryQuery {
	grq.ctx.Unique = &unique
	return grq
}

// Order specifies how the records should be ordered.
func (grq *GithubRepositoryQuery) Order(o ...githubrepository.OrderOption) *GithubRepositoryQuery {
	grq.order = append(grq.order, o...)
	return grq
}

// QueryLabels chains the current query on the "labels" edge.
func (grq *GithubRepositoryQuery) QueryLabels() *LabelQuery {
	query := (&LabelClient{config: grq.config}).Query()
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

// QueryReleases chains the current query on the "releases" edge.
func (grq *GithubRepositoryQuery) QueryReleases() *GithubReleaseQuery {
	query := (&GithubReleaseClient{config: grq.config}).Query()
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
			sqlgraph.To(githubrelease.Table, githubrelease.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, githubrepository.ReleasesTable, githubrepository.ReleasesColumn),
		)
		fromU = sqlgraph.SetNeighbors(grq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GithubRepository entity from the query.
// Returns a *NotFoundError when no GithubRepository was found.
func (grq *GithubRepositoryQuery) First(ctx context.Context) (*GithubRepository, error) {
	nodes, err := grq.Limit(1).All(setContextOp(ctx, grq.ctx, "First"))
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
	if ids, err = grq.Limit(1).IDs(setContextOp(ctx, grq.ctx, "FirstID")); err != nil {
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
	nodes, err := grq.Limit(2).All(setContextOp(ctx, grq.ctx, "Only"))
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
	if ids, err = grq.Limit(2).IDs(setContextOp(ctx, grq.ctx, "OnlyID")); err != nil {
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
	ctx = setContextOp(ctx, grq.ctx, "All")
	if err := grq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GithubRepository, *GithubRepositoryQuery]()
	return withInterceptors[[]*GithubRepository](ctx, grq, qr, grq.inters)
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
func (grq *GithubRepositoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if grq.ctx.Unique == nil && grq.path != nil {
		grq.Unique(true)
	}
	ctx = setContextOp(ctx, grq.ctx, "IDs")
	if err = grq.Select(githubrepository.FieldID).Scan(ctx, &ids); err != nil {
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
	ctx = setContextOp(ctx, grq.ctx, "Count")
	if err := grq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, grq, querierCount[*GithubRepositoryQuery](), grq.inters)
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
		config:       grq.config,
		ctx:          grq.ctx.Clone(),
		order:        append([]githubrepository.OrderOption{}, grq.order...),
		inters:       append([]Interceptor{}, grq.inters...),
		predicates:   append([]predicate.GithubRepository{}, grq.predicates...),
		withLabels:   grq.withLabels.Clone(),
		withReleases: grq.withReleases.Clone(),
		// clone intermediate query.
		sql:  grq.sql.Clone(),
		path: grq.path,
	}
}

// WithLabels tells the query-builder to eager-load the nodes that are connected to
// the "labels" edge. The optional arguments are used to configure the query builder of the edge.
func (grq *GithubRepositoryQuery) WithLabels(opts ...func(*LabelQuery)) *GithubRepositoryQuery {
	query := (&LabelClient{config: grq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	grq.withLabels = query
	return grq
}

// WithReleases tells the query-builder to eager-load the nodes that are connected to
// the "releases" edge. The optional arguments are used to configure the query builder of the edge.
func (grq *GithubRepositoryQuery) WithReleases(opts ...func(*GithubReleaseQuery)) *GithubRepositoryQuery {
	query := (&GithubReleaseClient{config: grq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	grq.withReleases = query
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
func (grq *GithubRepositoryQuery) GroupBy(field string, fields ...string) *GithubRepositoryGroupBy {
	grq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GithubRepositoryGroupBy{build: grq}
	grbuild.flds = &grq.ctx.Fields
	grbuild.label = githubrepository.Label
	grbuild.scan = grbuild.Scan
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
func (grq *GithubRepositoryQuery) Select(fields ...string) *GithubRepositorySelect {
	grq.ctx.Fields = append(grq.ctx.Fields, fields...)
	sbuild := &GithubRepositorySelect{GithubRepositoryQuery: grq}
	sbuild.label = githubrepository.Label
	sbuild.flds, sbuild.scan = &grq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GithubRepositorySelect configured with the given aggregations.
func (grq *GithubRepositoryQuery) Aggregate(fns ...AggregateFunc) *GithubRepositorySelect {
	return grq.Select().Aggregate(fns...)
}

func (grq *GithubRepositoryQuery) prepareQuery(ctx context.Context) error {
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
		loadedTypes = [2]bool{
			grq.withLabels != nil,
			grq.withReleases != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GithubRepository).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
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
		if err := grq.loadLabels(ctx, query, nodes,
			func(n *GithubRepository) { n.Edges.Labels = []*Label{} },
			func(n *GithubRepository, e *Label) { n.Edges.Labels = append(n.Edges.Labels, e) }); err != nil {
			return nil, err
		}
	}
	if query := grq.withReleases; query != nil {
		if err := grq.loadReleases(ctx, query, nodes,
			func(n *GithubRepository) { n.Edges.Releases = []*GithubRelease{} },
			func(n *GithubRepository, e *GithubRelease) { n.Edges.Releases = append(n.Edges.Releases, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range grq.withNamedLabels {
		if err := grq.loadLabels(ctx, query, nodes,
			func(n *GithubRepository) { n.appendNamedLabels(name) },
			func(n *GithubRepository, e *Label) { n.appendNamedLabels(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range grq.withNamedReleases {
		if err := grq.loadReleases(ctx, query, nodes,
			func(n *GithubRepository) { n.appendNamedReleases(name) },
			func(n *GithubRepository, e *GithubRelease) { n.appendNamedReleases(name, e) }); err != nil {
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

func (grq *GithubRepositoryQuery) loadLabels(ctx context.Context, query *LabelQuery, nodes []*GithubRepository, init func(*GithubRepository), assign func(*GithubRepository, *Label)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*GithubRepository)
	nids := make(map[int]map[*GithubRepository]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(githubrepository.LabelsTable)
		s.Join(joinT).On(s.C(label.FieldID), joinT.C(githubrepository.LabelsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(githubrepository.LabelsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(githubrepository.LabelsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*GithubRepository]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Label](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "labels" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (grq *GithubRepositoryQuery) loadReleases(ctx context.Context, query *GithubReleaseQuery, nodes []*GithubRepository, init func(*GithubRepository), assign func(*GithubRepository, *GithubRelease)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*GithubRepository)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.GithubRelease(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(githubrepository.ReleasesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.github_repository_releases
		if fk == nil {
			return fmt.Errorf(`foreign-key "github_repository_releases" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "github_repository_releases" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (grq *GithubRepositoryQuery) sqlCount(ctx context.Context) (int, error) {
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

func (grq *GithubRepositoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(githubrepository.Table, githubrepository.Columns, sqlgraph.NewFieldSpec(githubrepository.FieldID, field.TypeInt))
	_spec.From = grq.sql
	if unique := grq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if grq.path != nil {
		_spec.Unique = true
	}
	if fields := grq.ctx.Fields; len(fields) > 0 {
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

func (grq *GithubRepositoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(grq.driver.Dialect())
	t1 := builder.Table(githubrepository.Table)
	columns := grq.ctx.Fields
	if len(columns) == 0 {
		columns = githubrepository.Columns
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

// WithNamedLabels tells the query-builder to eager-load the nodes that are connected to the "labels"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (grq *GithubRepositoryQuery) WithNamedLabels(name string, opts ...func(*LabelQuery)) *GithubRepositoryQuery {
	query := (&LabelClient{config: grq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if grq.withNamedLabels == nil {
		grq.withNamedLabels = make(map[string]*LabelQuery)
	}
	grq.withNamedLabels[name] = query
	return grq
}

// WithNamedReleases tells the query-builder to eager-load the nodes that are connected to the "releases"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (grq *GithubRepositoryQuery) WithNamedReleases(name string, opts ...func(*GithubReleaseQuery)) *GithubRepositoryQuery {
	query := (&GithubReleaseClient{config: grq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if grq.withNamedReleases == nil {
		grq.withNamedReleases = make(map[string]*GithubReleaseQuery)
	}
	grq.withNamedReleases[name] = query
	return grq
}

// GithubRepositoryGroupBy is the group-by builder for GithubRepository entities.
type GithubRepositoryGroupBy struct {
	selector
	build *GithubRepositoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (grgb *GithubRepositoryGroupBy) Aggregate(fns ...AggregateFunc) *GithubRepositoryGroupBy {
	grgb.fns = append(grgb.fns, fns...)
	return grgb
}

// Scan applies the selector query and scans the result into the given value.
func (grgb *GithubRepositoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, grgb.build.ctx, "GroupBy")
	if err := grgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GithubRepositoryQuery, *GithubRepositoryGroupBy](ctx, grgb.build, grgb, grgb.build.inters, v)
}

func (grgb *GithubRepositoryGroupBy) sqlScan(ctx context.Context, root *GithubRepositoryQuery, v any) error {
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

// GithubRepositorySelect is the builder for selecting fields of GithubRepository entities.
type GithubRepositorySelect struct {
	*GithubRepositoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (grs *GithubRepositorySelect) Aggregate(fns ...AggregateFunc) *GithubRepositorySelect {
	grs.fns = append(grs.fns, fns...)
	return grs
}

// Scan applies the selector query and scans the result into the given value.
func (grs *GithubRepositorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, grs.ctx, "Select")
	if err := grs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GithubRepositoryQuery, *GithubRepositorySelect](ctx, grs.GithubRepositoryQuery, grs, grs.inters, v)
}

func (grs *GithubRepositorySelect) sqlScan(ctx context.Context, root *GithubRepositoryQuery, v any) error {
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
