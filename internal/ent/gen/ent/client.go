// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/migrate"
	"github.com/google/uuid"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/dnsblquery"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/dnsblresponse"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/ip"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/user"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// DNSBLQuery is the client for interacting with the DNSBLQuery builders.
	DNSBLQuery *DNSBLQueryClient
	// DNSBLResponse is the client for interacting with the DNSBLResponse builders.
	DNSBLResponse *DNSBLResponseClient
	// IP is the client for interacting with the IP builders.
	IP *IPClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.DNSBLQuery = NewDNSBLQueryClient(c.config)
	c.DNSBLResponse = NewDNSBLResponseClient(c.config)
	c.IP = NewIPClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		DNSBLQuery:    NewDNSBLQueryClient(cfg),
		DNSBLResponse: NewDNSBLResponseClient(cfg),
		IP:            NewIPClient(cfg),
		User:          NewUserClient(cfg),
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
		config:        cfg,
		DNSBLQuery:    NewDNSBLQueryClient(cfg),
		DNSBLResponse: NewDNSBLResponseClient(cfg),
		IP:            NewIPClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		DNSBLQuery.
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
	c.DNSBLQuery.Use(hooks...)
	c.DNSBLResponse.Use(hooks...)
	c.IP.Use(hooks...)
	c.User.Use(hooks...)
}

// DNSBLQueryClient is a client for the DNSBLQuery schema.
type DNSBLQueryClient struct {
	config
}

// NewDNSBLQueryClient returns a client for the DNSBLQuery from the given config.
func NewDNSBLQueryClient(c config) *DNSBLQueryClient {
	return &DNSBLQueryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dnsblquery.Hooks(f(g(h())))`.
func (c *DNSBLQueryClient) Use(hooks ...Hook) {
	c.hooks.DNSBLQuery = append(c.hooks.DNSBLQuery, hooks...)
}

// Create returns a create builder for DNSBLQuery.
func (c *DNSBLQueryClient) Create() *DNSBLQueryCreate {
	mutation := newDNSBLQueryMutation(c.config, OpCreate)
	return &DNSBLQueryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of DNSBLQuery entities.
func (c *DNSBLQueryClient) CreateBulk(builders ...*DNSBLQueryCreate) *DNSBLQueryCreateBulk {
	return &DNSBLQueryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DNSBLQuery.
func (c *DNSBLQueryClient) Update() *DNSBLQueryUpdate {
	mutation := newDNSBLQueryMutation(c.config, OpUpdate)
	return &DNSBLQueryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DNSBLQueryClient) UpdateOne(dq *DNSBLQuery) *DNSBLQueryUpdateOne {
	mutation := newDNSBLQueryMutation(c.config, OpUpdateOne, withDNSBLQuery(dq))
	return &DNSBLQueryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DNSBLQueryClient) UpdateOneID(id uuid.UUID) *DNSBLQueryUpdateOne {
	mutation := newDNSBLQueryMutation(c.config, OpUpdateOne, withDNSBLQueryID(id))
	return &DNSBLQueryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DNSBLQuery.
func (c *DNSBLQueryClient) Delete() *DNSBLQueryDelete {
	mutation := newDNSBLQueryMutation(c.config, OpDelete)
	return &DNSBLQueryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DNSBLQueryClient) DeleteOne(dq *DNSBLQuery) *DNSBLQueryDeleteOne {
	return c.DeleteOneID(dq.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DNSBLQueryClient) DeleteOneID(id uuid.UUID) *DNSBLQueryDeleteOne {
	builder := c.Delete().Where(dnsblquery.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DNSBLQueryDeleteOne{builder}
}

// Query returns a query builder for DNSBLQuery.
func (c *DNSBLQueryClient) Query() *DNSBLQueryQuery {
	return &DNSBLQueryQuery{config: c.config}
}

// Get returns a DNSBLQuery entity by its id.
func (c *DNSBLQueryClient) Get(ctx context.Context, id uuid.UUID) (*DNSBLQuery, error) {
	return c.Query().Where(dnsblquery.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DNSBLQueryClient) GetX(ctx context.Context, id uuid.UUID) *DNSBLQuery {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryResponses queries the responses edge of a DNSBLQuery.
func (c *DNSBLQueryClient) QueryResponses(dq *DNSBLQuery) *DNSBLResponseQuery {
	query := &DNSBLResponseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dq.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dnsblquery.Table, dnsblquery.FieldID, id),
			sqlgraph.To(dnsblresponse.Table, dnsblresponse.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dnsblquery.ResponsesTable, dnsblquery.ResponsesColumn),
		)
		fromV = sqlgraph.Neighbors(dq.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryIPAddress queries the ip_address edge of a DNSBLQuery.
func (c *DNSBLQueryClient) QueryIPAddress(dq *DNSBLQuery) *IPQuery {
	query := &IPQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dq.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dnsblquery.Table, dnsblquery.FieldID, id),
			sqlgraph.To(ip.Table, ip.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dnsblquery.IPAddressTable, dnsblquery.IPAddressColumn),
		)
		fromV = sqlgraph.Neighbors(dq.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DNSBLQueryClient) Hooks() []Hook {
	return c.hooks.DNSBLQuery
}

// DNSBLResponseClient is a client for the DNSBLResponse schema.
type DNSBLResponseClient struct {
	config
}

// NewDNSBLResponseClient returns a client for the DNSBLResponse from the given config.
func NewDNSBLResponseClient(c config) *DNSBLResponseClient {
	return &DNSBLResponseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dnsblresponse.Hooks(f(g(h())))`.
func (c *DNSBLResponseClient) Use(hooks ...Hook) {
	c.hooks.DNSBLResponse = append(c.hooks.DNSBLResponse, hooks...)
}

// Create returns a create builder for DNSBLResponse.
func (c *DNSBLResponseClient) Create() *DNSBLResponseCreate {
	mutation := newDNSBLResponseMutation(c.config, OpCreate)
	return &DNSBLResponseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of DNSBLResponse entities.
func (c *DNSBLResponseClient) CreateBulk(builders ...*DNSBLResponseCreate) *DNSBLResponseCreateBulk {
	return &DNSBLResponseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DNSBLResponse.
func (c *DNSBLResponseClient) Update() *DNSBLResponseUpdate {
	mutation := newDNSBLResponseMutation(c.config, OpUpdate)
	return &DNSBLResponseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DNSBLResponseClient) UpdateOne(dr *DNSBLResponse) *DNSBLResponseUpdateOne {
	mutation := newDNSBLResponseMutation(c.config, OpUpdateOne, withDNSBLResponse(dr))
	return &DNSBLResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DNSBLResponseClient) UpdateOneID(id uuid.UUID) *DNSBLResponseUpdateOne {
	mutation := newDNSBLResponseMutation(c.config, OpUpdateOne, withDNSBLResponseID(id))
	return &DNSBLResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DNSBLResponse.
func (c *DNSBLResponseClient) Delete() *DNSBLResponseDelete {
	mutation := newDNSBLResponseMutation(c.config, OpDelete)
	return &DNSBLResponseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DNSBLResponseClient) DeleteOne(dr *DNSBLResponse) *DNSBLResponseDeleteOne {
	return c.DeleteOneID(dr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DNSBLResponseClient) DeleteOneID(id uuid.UUID) *DNSBLResponseDeleteOne {
	builder := c.Delete().Where(dnsblresponse.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DNSBLResponseDeleteOne{builder}
}

// Query returns a query builder for DNSBLResponse.
func (c *DNSBLResponseClient) Query() *DNSBLResponseQuery {
	return &DNSBLResponseQuery{config: c.config}
}

// Get returns a DNSBLResponse entity by its id.
func (c *DNSBLResponseClient) Get(ctx context.Context, id uuid.UUID) (*DNSBLResponse, error) {
	return c.Query().Where(dnsblresponse.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DNSBLResponseClient) GetX(ctx context.Context, id uuid.UUID) *DNSBLResponse {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryQuery queries the query edge of a DNSBLResponse.
func (c *DNSBLResponseClient) QueryQuery(dr *DNSBLResponse) *DNSBLQueryQuery {
	query := &DNSBLQueryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dnsblresponse.Table, dnsblresponse.FieldID, id),
			sqlgraph.To(dnsblquery.Table, dnsblquery.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dnsblresponse.QueryTable, dnsblresponse.QueryColumn),
		)
		fromV = sqlgraph.Neighbors(dr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DNSBLResponseClient) Hooks() []Hook {
	return c.hooks.DNSBLResponse
}

// IPClient is a client for the IP schema.
type IPClient struct {
	config
}

// NewIPClient returns a client for the IP from the given config.
func NewIPClient(c config) *IPClient {
	return &IPClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ip.Hooks(f(g(h())))`.
func (c *IPClient) Use(hooks ...Hook) {
	c.hooks.IP = append(c.hooks.IP, hooks...)
}

// Create returns a create builder for IP.
func (c *IPClient) Create() *IPCreate {
	mutation := newIPMutation(c.config, OpCreate)
	return &IPCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of IP entities.
func (c *IPClient) CreateBulk(builders ...*IPCreate) *IPCreateBulk {
	return &IPCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for IP.
func (c *IPClient) Update() *IPUpdate {
	mutation := newIPMutation(c.config, OpUpdate)
	return &IPUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IPClient) UpdateOne(i *IP) *IPUpdateOne {
	mutation := newIPMutation(c.config, OpUpdateOne, withIP(i))
	return &IPUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IPClient) UpdateOneID(id uuid.UUID) *IPUpdateOne {
	mutation := newIPMutation(c.config, OpUpdateOne, withIPID(id))
	return &IPUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for IP.
func (c *IPClient) Delete() *IPDelete {
	mutation := newIPMutation(c.config, OpDelete)
	return &IPDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *IPClient) DeleteOne(i *IP) *IPDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *IPClient) DeleteOneID(id uuid.UUID) *IPDeleteOne {
	builder := c.Delete().Where(ip.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IPDeleteOne{builder}
}

// Query returns a query builder for IP.
func (c *IPClient) Query() *IPQuery {
	return &IPQuery{config: c.config}
}

// Get returns a IP entity by its id.
func (c *IPClient) Get(ctx context.Context, id uuid.UUID) (*IP, error) {
	return c.Query().Where(ip.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IPClient) GetX(ctx context.Context, id uuid.UUID) *IP {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryQueries queries the queries edge of a IP.
func (c *IPClient) QueryQueries(i *IP) *DNSBLQueryQuery {
	query := &DNSBLQueryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ip.Table, ip.FieldID, id),
			sqlgraph.To(dnsblquery.Table, dnsblquery.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ip.QueriesTable, ip.QueriesColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *IPClient) Hooks() []Hook {
	return c.hooks.IP
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

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of User entities.
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
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
