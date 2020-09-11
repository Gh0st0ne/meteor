// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/degenerat3/meteor/meteor/core/ent/migrate"

	"github.com/degenerat3/meteor/meteor/core/ent/action"
	"github.com/degenerat3/meteor/meteor/core/ent/bot"
	"github.com/degenerat3/meteor/meteor/core/ent/group"
	"github.com/degenerat3/meteor/meteor/core/ent/host"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Action is the client for interacting with the Action builders.
	Action *ActionClient
	// Bot is the client for interacting with the Bot builders.
	Bot *BotClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Host is the client for interacting with the Host builders.
	Host *HostClient
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
	c.Action = NewActionClient(c.config)
	c.Bot = NewBotClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.Host = NewHostClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Action: NewActionClient(cfg),
		Bot:    NewBotClient(cfg),
		Group:  NewGroupClient(cfg),
		Host:   NewHostClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config: cfg,
		Action: NewActionClient(cfg),
		Bot:    NewBotClient(cfg),
		Group:  NewGroupClient(cfg),
		Host:   NewHostClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Action.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Action.Use(hooks...)
	c.Bot.Use(hooks...)
	c.Group.Use(hooks...)
	c.Host.Use(hooks...)
}

// ActionClient is a client for the Action schema.
type ActionClient struct {
	config
}

// NewActionClient returns a client for the Action from the given config.
func NewActionClient(c config) *ActionClient {
	return &ActionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `action.Hooks(f(g(h())))`.
func (c *ActionClient) Use(hooks ...Hook) {
	c.hooks.Action = append(c.hooks.Action, hooks...)
}

// Create returns a create builder for Action.
func (c *ActionClient) Create() *ActionCreate {
	mutation := newActionMutation(c.config, OpCreate)
	return &ActionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Action entities.
func (c *ActionClient) CreateBulk(builders ...*ActionCreate) *ActionCreateBulk {
	return &ActionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Action.
func (c *ActionClient) Update() *ActionUpdate {
	mutation := newActionMutation(c.config, OpUpdate)
	return &ActionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ActionClient) UpdateOne(a *Action) *ActionUpdateOne {
	mutation := newActionMutation(c.config, OpUpdateOne, withAction(a))
	return &ActionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ActionClient) UpdateOneID(id int) *ActionUpdateOne {
	mutation := newActionMutation(c.config, OpUpdateOne, withActionID(id))
	return &ActionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Action.
func (c *ActionClient) Delete() *ActionDelete {
	mutation := newActionMutation(c.config, OpDelete)
	return &ActionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ActionClient) DeleteOne(a *Action) *ActionDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ActionClient) DeleteOneID(id int) *ActionDeleteOne {
	builder := c.Delete().Where(action.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ActionDeleteOne{builder}
}

// Query returns a query builder for Action.
func (c *ActionClient) Query() *ActionQuery {
	return &ActionQuery{config: c.config}
}

// Get returns a Action entity by its id.
func (c *ActionClient) Get(ctx context.Context, id int) (*Action, error) {
	return c.Query().Where(action.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ActionClient) GetX(ctx context.Context, id int) *Action {
	a, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return a
}

// QueryTargeting queries the targeting edge of a Action.
func (c *ActionClient) QueryTargeting(a *Action) *HostQuery {
	query := &HostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(action.Table, action.FieldID, id),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, action.TargetingTable, action.TargetingColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ActionClient) Hooks() []Hook {
	return c.hooks.Action
}

// BotClient is a client for the Bot schema.
type BotClient struct {
	config
}

// NewBotClient returns a client for the Bot from the given config.
func NewBotClient(c config) *BotClient {
	return &BotClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bot.Hooks(f(g(h())))`.
func (c *BotClient) Use(hooks ...Hook) {
	c.hooks.Bot = append(c.hooks.Bot, hooks...)
}

// Create returns a create builder for Bot.
func (c *BotClient) Create() *BotCreate {
	mutation := newBotMutation(c.config, OpCreate)
	return &BotCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Bot entities.
func (c *BotClient) CreateBulk(builders ...*BotCreate) *BotCreateBulk {
	return &BotCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Bot.
func (c *BotClient) Update() *BotUpdate {
	mutation := newBotMutation(c.config, OpUpdate)
	return &BotUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BotClient) UpdateOne(b *Bot) *BotUpdateOne {
	mutation := newBotMutation(c.config, OpUpdateOne, withBot(b))
	return &BotUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BotClient) UpdateOneID(id int) *BotUpdateOne {
	mutation := newBotMutation(c.config, OpUpdateOne, withBotID(id))
	return &BotUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Bot.
func (c *BotClient) Delete() *BotDelete {
	mutation := newBotMutation(c.config, OpDelete)
	return &BotDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BotClient) DeleteOne(b *Bot) *BotDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BotClient) DeleteOneID(id int) *BotDeleteOne {
	builder := c.Delete().Where(bot.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BotDeleteOne{builder}
}

// Query returns a query builder for Bot.
func (c *BotClient) Query() *BotQuery {
	return &BotQuery{config: c.config}
}

// Get returns a Bot entity by its id.
func (c *BotClient) Get(ctx context.Context, id int) (*Bot, error) {
	return c.Query().Where(bot.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BotClient) GetX(ctx context.Context, id int) *Bot {
	b, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return b
}

// QueryInfecting queries the infecting edge of a Bot.
func (c *BotClient) QueryInfecting(b *Bot) *HostQuery {
	query := &HostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bot.Table, bot.FieldID, id),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bot.InfectingTable, bot.InfectingColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BotClient) Hooks() []Hook {
	return c.hooks.Bot
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a create builder for Group.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{config: c.config}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	gr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return gr
}

// QueryMembers queries the members edge of a Group.
func (c *GroupClient) QueryMembers(gr *Group) *HostQuery {
	query := &HostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.MembersTable, group.MembersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// HostClient is a client for the Host schema.
type HostClient struct {
	config
}

// NewHostClient returns a client for the Host from the given config.
func NewHostClient(c config) *HostClient {
	return &HostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `host.Hooks(f(g(h())))`.
func (c *HostClient) Use(hooks ...Hook) {
	c.hooks.Host = append(c.hooks.Host, hooks...)
}

// Create returns a create builder for Host.
func (c *HostClient) Create() *HostCreate {
	mutation := newHostMutation(c.config, OpCreate)
	return &HostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Host entities.
func (c *HostClient) CreateBulk(builders ...*HostCreate) *HostCreateBulk {
	return &HostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Host.
func (c *HostClient) Update() *HostUpdate {
	mutation := newHostMutation(c.config, OpUpdate)
	return &HostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HostClient) UpdateOne(h *Host) *HostUpdateOne {
	mutation := newHostMutation(c.config, OpUpdateOne, withHost(h))
	return &HostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HostClient) UpdateOneID(id int) *HostUpdateOne {
	mutation := newHostMutation(c.config, OpUpdateOne, withHostID(id))
	return &HostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Host.
func (c *HostClient) Delete() *HostDelete {
	mutation := newHostMutation(c.config, OpDelete)
	return &HostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *HostClient) DeleteOne(h *Host) *HostDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *HostClient) DeleteOneID(id int) *HostDeleteOne {
	builder := c.Delete().Where(host.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HostDeleteOne{builder}
}

// Query returns a query builder for Host.
func (c *HostClient) Query() *HostQuery {
	return &HostQuery{config: c.config}
}

// Get returns a Host entity by its id.
func (c *HostClient) Get(ctx context.Context, id int) (*Host, error) {
	return c.Query().Where(host.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HostClient) GetX(ctx context.Context, id int) *Host {
	h, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return h
}

// QueryBots queries the bots edge of a Host.
func (c *HostClient) QueryBots(h *Host) *BotQuery {
	query := &BotQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, id),
			sqlgraph.To(bot.Table, bot.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, host.BotsTable, host.BotsColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryActions queries the actions edge of a Host.
func (c *HostClient) QueryActions(h *Host) *ActionQuery {
	query := &ActionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, id),
			sqlgraph.To(action.Table, action.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, host.ActionsTable, host.ActionsColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMember queries the member edge of a Host.
func (c *HostClient) QueryMember(h *Host) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(host.Table, host.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, host.MemberTable, host.MemberPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HostClient) Hooks() []Hook {
	return c.hooks.Host
}
