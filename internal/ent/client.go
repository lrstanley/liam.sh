// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// DO NOT EDIT, CODE GENERATED BY entc.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/lrstanley/liam.sh/internal/ent/migrate"

	"github.com/lrstanley/liam.sh/internal/ent/githubasset"
	"github.com/lrstanley/liam.sh/internal/ent/githubevent"
	"github.com/lrstanley/liam.sh/internal/ent/githubrelease"
	"github.com/lrstanley/liam.sh/internal/ent/githubrepository"
	"github.com/lrstanley/liam.sh/internal/ent/label"
	"github.com/lrstanley/liam.sh/internal/ent/post"
	"github.com/lrstanley/liam.sh/internal/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// GithubAsset is the client for interacting with the GithubAsset builders.
	GithubAsset *GithubAssetClient
	// GithubEvent is the client for interacting with the GithubEvent builders.
	GithubEvent *GithubEventClient
	// GithubRelease is the client for interacting with the GithubRelease builders.
	GithubRelease *GithubReleaseClient
	// GithubRepository is the client for interacting with the GithubRepository builders.
	GithubRepository *GithubRepositoryClient
	// Label is the client for interacting with the Label builders.
	Label *LabelClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.GithubAsset = NewGithubAssetClient(c.config)
	c.GithubEvent = NewGithubEventClient(c.config)
	c.GithubRelease = NewGithubReleaseClient(c.config)
	c.GithubRepository = NewGithubRepositoryClient(c.config)
	c.Label = NewLabelClient(c.config)
	c.Post = NewPostClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		GithubAsset:      NewGithubAssetClient(cfg),
		GithubEvent:      NewGithubEventClient(cfg),
		GithubRelease:    NewGithubReleaseClient(cfg),
		GithubRepository: NewGithubRepositoryClient(cfg),
		Label:            NewLabelClient(cfg),
		Post:             NewPostClient(cfg),
		User:             NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		GithubAsset:      NewGithubAssetClient(cfg),
		GithubEvent:      NewGithubEventClient(cfg),
		GithubRelease:    NewGithubReleaseClient(cfg),
		GithubRepository: NewGithubRepositoryClient(cfg),
		Label:            NewLabelClient(cfg),
		Post:             NewPostClient(cfg),
		User:             NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		GithubAsset.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.GithubAsset.Use(hooks...)
	c.GithubEvent.Use(hooks...)
	c.GithubRelease.Use(hooks...)
	c.GithubRepository.Use(hooks...)
	c.Label.Use(hooks...)
	c.Post.Use(hooks...)
	c.User.Use(hooks...)
}

// GithubAssetClient is a client for the GithubAsset schema.
type GithubAssetClient struct {
	config
}

// NewGithubAssetClient returns a client for the GithubAsset from the given config.
func NewGithubAssetClient(c config) *GithubAssetClient {
	return &GithubAssetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `githubasset.Hooks(f(g(h())))`.
func (c *GithubAssetClient) Use(hooks ...Hook) {
	c.hooks.GithubAsset = append(c.hooks.GithubAsset, hooks...)
}

// Create returns a builder for creating a GithubAsset entity.
func (c *GithubAssetClient) Create() *GithubAssetCreate {
	mutation := newGithubAssetMutation(c.config, OpCreate)
	return &GithubAssetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GithubAsset entities.
func (c *GithubAssetClient) CreateBulk(builders ...*GithubAssetCreate) *GithubAssetCreateBulk {
	return &GithubAssetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GithubAsset.
func (c *GithubAssetClient) Update() *GithubAssetUpdate {
	mutation := newGithubAssetMutation(c.config, OpUpdate)
	return &GithubAssetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GithubAssetClient) UpdateOne(ga *GithubAsset) *GithubAssetUpdateOne {
	mutation := newGithubAssetMutation(c.config, OpUpdateOne, withGithubAsset(ga))
	return &GithubAssetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GithubAssetClient) UpdateOneID(id int) *GithubAssetUpdateOne {
	mutation := newGithubAssetMutation(c.config, OpUpdateOne, withGithubAssetID(id))
	return &GithubAssetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GithubAsset.
func (c *GithubAssetClient) Delete() *GithubAssetDelete {
	mutation := newGithubAssetMutation(c.config, OpDelete)
	return &GithubAssetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GithubAssetClient) DeleteOne(ga *GithubAsset) *GithubAssetDeleteOne {
	return c.DeleteOneID(ga.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GithubAssetClient) DeleteOneID(id int) *GithubAssetDeleteOne {
	builder := c.Delete().Where(githubasset.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GithubAssetDeleteOne{builder}
}

// Query returns a query builder for GithubAsset.
func (c *GithubAssetClient) Query() *GithubAssetQuery {
	return &GithubAssetQuery{
		config: c.config,
	}
}

// Get returns a GithubAsset entity by its id.
func (c *GithubAssetClient) Get(ctx context.Context, id int) (*GithubAsset, error) {
	return c.Query().Where(githubasset.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GithubAssetClient) GetX(ctx context.Context, id int) *GithubAsset {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRelease queries the release edge of a GithubAsset.
func (c *GithubAssetClient) QueryRelease(ga *GithubAsset) *GithubReleaseQuery {
	query := &GithubReleaseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(githubasset.Table, githubasset.FieldID, id),
			sqlgraph.To(githubrelease.Table, githubrelease.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, githubasset.ReleaseTable, githubasset.ReleaseColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GithubAssetClient) Hooks() []Hook {
	hooks := c.hooks.GithubAsset
	return append(hooks[:len(hooks):len(hooks)], githubasset.Hooks[:]...)
}

// GithubEventClient is a client for the GithubEvent schema.
type GithubEventClient struct {
	config
}

// NewGithubEventClient returns a client for the GithubEvent from the given config.
func NewGithubEventClient(c config) *GithubEventClient {
	return &GithubEventClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `githubevent.Hooks(f(g(h())))`.
func (c *GithubEventClient) Use(hooks ...Hook) {
	c.hooks.GithubEvent = append(c.hooks.GithubEvent, hooks...)
}

// Create returns a builder for creating a GithubEvent entity.
func (c *GithubEventClient) Create() *GithubEventCreate {
	mutation := newGithubEventMutation(c.config, OpCreate)
	return &GithubEventCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GithubEvent entities.
func (c *GithubEventClient) CreateBulk(builders ...*GithubEventCreate) *GithubEventCreateBulk {
	return &GithubEventCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GithubEvent.
func (c *GithubEventClient) Update() *GithubEventUpdate {
	mutation := newGithubEventMutation(c.config, OpUpdate)
	return &GithubEventUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GithubEventClient) UpdateOne(ge *GithubEvent) *GithubEventUpdateOne {
	mutation := newGithubEventMutation(c.config, OpUpdateOne, withGithubEvent(ge))
	return &GithubEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GithubEventClient) UpdateOneID(id int) *GithubEventUpdateOne {
	mutation := newGithubEventMutation(c.config, OpUpdateOne, withGithubEventID(id))
	return &GithubEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GithubEvent.
func (c *GithubEventClient) Delete() *GithubEventDelete {
	mutation := newGithubEventMutation(c.config, OpDelete)
	return &GithubEventDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GithubEventClient) DeleteOne(ge *GithubEvent) *GithubEventDeleteOne {
	return c.DeleteOneID(ge.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GithubEventClient) DeleteOneID(id int) *GithubEventDeleteOne {
	builder := c.Delete().Where(githubevent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GithubEventDeleteOne{builder}
}

// Query returns a query builder for GithubEvent.
func (c *GithubEventClient) Query() *GithubEventQuery {
	return &GithubEventQuery{
		config: c.config,
	}
}

// Get returns a GithubEvent entity by its id.
func (c *GithubEventClient) Get(ctx context.Context, id int) (*GithubEvent, error) {
	return c.Query().Where(githubevent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GithubEventClient) GetX(ctx context.Context, id int) *GithubEvent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GithubEventClient) Hooks() []Hook {
	hooks := c.hooks.GithubEvent
	return append(hooks[:len(hooks):len(hooks)], githubevent.Hooks[:]...)
}

// GithubReleaseClient is a client for the GithubRelease schema.
type GithubReleaseClient struct {
	config
}

// NewGithubReleaseClient returns a client for the GithubRelease from the given config.
func NewGithubReleaseClient(c config) *GithubReleaseClient {
	return &GithubReleaseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `githubrelease.Hooks(f(g(h())))`.
func (c *GithubReleaseClient) Use(hooks ...Hook) {
	c.hooks.GithubRelease = append(c.hooks.GithubRelease, hooks...)
}

// Create returns a builder for creating a GithubRelease entity.
func (c *GithubReleaseClient) Create() *GithubReleaseCreate {
	mutation := newGithubReleaseMutation(c.config, OpCreate)
	return &GithubReleaseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GithubRelease entities.
func (c *GithubReleaseClient) CreateBulk(builders ...*GithubReleaseCreate) *GithubReleaseCreateBulk {
	return &GithubReleaseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GithubRelease.
func (c *GithubReleaseClient) Update() *GithubReleaseUpdate {
	mutation := newGithubReleaseMutation(c.config, OpUpdate)
	return &GithubReleaseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GithubReleaseClient) UpdateOne(gr *GithubRelease) *GithubReleaseUpdateOne {
	mutation := newGithubReleaseMutation(c.config, OpUpdateOne, withGithubRelease(gr))
	return &GithubReleaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GithubReleaseClient) UpdateOneID(id int) *GithubReleaseUpdateOne {
	mutation := newGithubReleaseMutation(c.config, OpUpdateOne, withGithubReleaseID(id))
	return &GithubReleaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GithubRelease.
func (c *GithubReleaseClient) Delete() *GithubReleaseDelete {
	mutation := newGithubReleaseMutation(c.config, OpDelete)
	return &GithubReleaseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GithubReleaseClient) DeleteOne(gr *GithubRelease) *GithubReleaseDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GithubReleaseClient) DeleteOneID(id int) *GithubReleaseDeleteOne {
	builder := c.Delete().Where(githubrelease.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GithubReleaseDeleteOne{builder}
}

// Query returns a query builder for GithubRelease.
func (c *GithubReleaseClient) Query() *GithubReleaseQuery {
	return &GithubReleaseQuery{
		config: c.config,
	}
}

// Get returns a GithubRelease entity by its id.
func (c *GithubReleaseClient) Get(ctx context.Context, id int) (*GithubRelease, error) {
	return c.Query().Where(githubrelease.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GithubReleaseClient) GetX(ctx context.Context, id int) *GithubRelease {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRepository queries the repository edge of a GithubRelease.
func (c *GithubReleaseClient) QueryRepository(gr *GithubRelease) *GithubRepositoryQuery {
	query := &GithubRepositoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(githubrelease.Table, githubrelease.FieldID, id),
			sqlgraph.To(githubrepository.Table, githubrepository.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, githubrelease.RepositoryTable, githubrelease.RepositoryColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAssets queries the assets edge of a GithubRelease.
func (c *GithubReleaseClient) QueryAssets(gr *GithubRelease) *GithubAssetQuery {
	query := &GithubAssetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(githubrelease.Table, githubrelease.FieldID, id),
			sqlgraph.To(githubasset.Table, githubasset.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, githubrelease.AssetsTable, githubrelease.AssetsColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GithubReleaseClient) Hooks() []Hook {
	hooks := c.hooks.GithubRelease
	return append(hooks[:len(hooks):len(hooks)], githubrelease.Hooks[:]...)
}

// GithubRepositoryClient is a client for the GithubRepository schema.
type GithubRepositoryClient struct {
	config
}

// NewGithubRepositoryClient returns a client for the GithubRepository from the given config.
func NewGithubRepositoryClient(c config) *GithubRepositoryClient {
	return &GithubRepositoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `githubrepository.Hooks(f(g(h())))`.
func (c *GithubRepositoryClient) Use(hooks ...Hook) {
	c.hooks.GithubRepository = append(c.hooks.GithubRepository, hooks...)
}

// Create returns a builder for creating a GithubRepository entity.
func (c *GithubRepositoryClient) Create() *GithubRepositoryCreate {
	mutation := newGithubRepositoryMutation(c.config, OpCreate)
	return &GithubRepositoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GithubRepository entities.
func (c *GithubRepositoryClient) CreateBulk(builders ...*GithubRepositoryCreate) *GithubRepositoryCreateBulk {
	return &GithubRepositoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GithubRepository.
func (c *GithubRepositoryClient) Update() *GithubRepositoryUpdate {
	mutation := newGithubRepositoryMutation(c.config, OpUpdate)
	return &GithubRepositoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GithubRepositoryClient) UpdateOne(gr *GithubRepository) *GithubRepositoryUpdateOne {
	mutation := newGithubRepositoryMutation(c.config, OpUpdateOne, withGithubRepository(gr))
	return &GithubRepositoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GithubRepositoryClient) UpdateOneID(id int) *GithubRepositoryUpdateOne {
	mutation := newGithubRepositoryMutation(c.config, OpUpdateOne, withGithubRepositoryID(id))
	return &GithubRepositoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GithubRepository.
func (c *GithubRepositoryClient) Delete() *GithubRepositoryDelete {
	mutation := newGithubRepositoryMutation(c.config, OpDelete)
	return &GithubRepositoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GithubRepositoryClient) DeleteOne(gr *GithubRepository) *GithubRepositoryDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GithubRepositoryClient) DeleteOneID(id int) *GithubRepositoryDeleteOne {
	builder := c.Delete().Where(githubrepository.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GithubRepositoryDeleteOne{builder}
}

// Query returns a query builder for GithubRepository.
func (c *GithubRepositoryClient) Query() *GithubRepositoryQuery {
	return &GithubRepositoryQuery{
		config: c.config,
	}
}

// Get returns a GithubRepository entity by its id.
func (c *GithubRepositoryClient) Get(ctx context.Context, id int) (*GithubRepository, error) {
	return c.Query().Where(githubrepository.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GithubRepositoryClient) GetX(ctx context.Context, id int) *GithubRepository {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLabels queries the labels edge of a GithubRepository.
func (c *GithubRepositoryClient) QueryLabels(gr *GithubRepository) *LabelQuery {
	query := &LabelQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(githubrepository.Table, githubrepository.FieldID, id),
			sqlgraph.To(label.Table, label.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, githubrepository.LabelsTable, githubrepository.LabelsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReleases queries the releases edge of a GithubRepository.
func (c *GithubRepositoryClient) QueryReleases(gr *GithubRepository) *GithubReleaseQuery {
	query := &GithubReleaseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(githubrepository.Table, githubrepository.FieldID, id),
			sqlgraph.To(githubrelease.Table, githubrelease.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, githubrepository.ReleasesTable, githubrepository.ReleasesColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GithubRepositoryClient) Hooks() []Hook {
	hooks := c.hooks.GithubRepository
	return append(hooks[:len(hooks):len(hooks)], githubrepository.Hooks[:]...)
}

// LabelClient is a client for the Label schema.
type LabelClient struct {
	config
}

// NewLabelClient returns a client for the Label from the given config.
func NewLabelClient(c config) *LabelClient {
	return &LabelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `label.Hooks(f(g(h())))`.
func (c *LabelClient) Use(hooks ...Hook) {
	c.hooks.Label = append(c.hooks.Label, hooks...)
}

// Create returns a builder for creating a Label entity.
func (c *LabelClient) Create() *LabelCreate {
	mutation := newLabelMutation(c.config, OpCreate)
	return &LabelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Label entities.
func (c *LabelClient) CreateBulk(builders ...*LabelCreate) *LabelCreateBulk {
	return &LabelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Label.
func (c *LabelClient) Update() *LabelUpdate {
	mutation := newLabelMutation(c.config, OpUpdate)
	return &LabelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LabelClient) UpdateOne(l *Label) *LabelUpdateOne {
	mutation := newLabelMutation(c.config, OpUpdateOne, withLabel(l))
	return &LabelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LabelClient) UpdateOneID(id int) *LabelUpdateOne {
	mutation := newLabelMutation(c.config, OpUpdateOne, withLabelID(id))
	return &LabelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Label.
func (c *LabelClient) Delete() *LabelDelete {
	mutation := newLabelMutation(c.config, OpDelete)
	return &LabelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LabelClient) DeleteOne(l *Label) *LabelDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *LabelClient) DeleteOneID(id int) *LabelDeleteOne {
	builder := c.Delete().Where(label.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LabelDeleteOne{builder}
}

// Query returns a query builder for Label.
func (c *LabelClient) Query() *LabelQuery {
	return &LabelQuery{
		config: c.config,
	}
}

// Get returns a Label entity by its id.
func (c *LabelClient) Get(ctx context.Context, id int) (*Label, error) {
	return c.Query().Where(label.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LabelClient) GetX(ctx context.Context, id int) *Label {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPosts queries the posts edge of a Label.
func (c *LabelClient) QueryPosts(l *Label) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(label.Table, label.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, label.PostsTable, label.PostsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGithubRepositories queries the github_repositories edge of a Label.
func (c *LabelClient) QueryGithubRepositories(l *Label) *GithubRepositoryQuery {
	query := &GithubRepositoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(label.Table, label.FieldID, id),
			sqlgraph.To(githubrepository.Table, githubrepository.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, label.GithubRepositoriesTable, label.GithubRepositoriesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LabelClient) Hooks() []Hook {
	hooks := c.hooks.Label
	return append(hooks[:len(hooks):len(hooks)], label.Hooks[:]...)
}

// PostClient is a client for the Post schema.
type PostClient struct {
	config
}

// NewPostClient returns a client for the Post from the given config.
func NewPostClient(c config) *PostClient {
	return &PostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `post.Hooks(f(g(h())))`.
func (c *PostClient) Use(hooks ...Hook) {
	c.hooks.Post = append(c.hooks.Post, hooks...)
}

// Create returns a builder for creating a Post entity.
func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Post entities.
func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
	return &PostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Post.
func (c *PostClient) Update() *PostUpdate {
	mutation := newPostMutation(c.config, OpUpdate)
	return &PostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PostClient) UpdateOne(po *Post) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPost(po))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PostClient) UpdateOneID(id int) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPostID(id))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Post.
func (c *PostClient) Delete() *PostDelete {
	mutation := newPostMutation(c.config, OpDelete)
	return &PostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PostClient) DeleteOne(po *Post) *PostDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PostClient) DeleteOneID(id int) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

// Query returns a query builder for Post.
func (c *PostClient) Query() *PostQuery {
	return &PostQuery{
		config: c.config,
	}
}

// Get returns a Post entity by its id.
func (c *PostClient) Get(ctx context.Context, id int) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PostClient) GetX(ctx context.Context, id int) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a Post.
func (c *PostClient) QueryAuthor(po *Post) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.AuthorTable, post.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLabels queries the labels edge of a Post.
func (c *PostClient) QueryLabels(po *Post) *LabelQuery {
	query := &LabelQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(label.Table, label.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, post.LabelsTable, post.LabelsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PostClient) Hooks() []Hook {
	hooks := c.hooks.Post
	return append(hooks[:len(hooks):len(hooks)], post.Hooks[:]...)
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPosts queries the posts edge of a User.
func (c *UserClient) QueryPosts(u *User) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PostsTable, user.PostsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}
