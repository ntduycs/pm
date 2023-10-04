// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"project-management/ent/migrate"

	"project-management/ent/member"
	"project-management/ent/papc"
	"project-management/ent/papctechnicalscore"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Member is the client for interacting with the Member builders.
	Member *MemberClient
	// PaPc is the client for interacting with the PaPc builders.
	PaPc *PaPcClient
	// PaPcTechnicalScore is the client for interacting with the PaPcTechnicalScore builders.
	PaPcTechnicalScore *PaPcTechnicalScoreClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Member = NewMemberClient(c.config)
	c.PaPc = NewPaPcClient(c.config)
	c.PaPcTechnicalScore = NewPaPcTechnicalScoreClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		Member:             NewMemberClient(cfg),
		PaPc:               NewPaPcClient(cfg),
		PaPcTechnicalScore: NewPaPcTechnicalScoreClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:                ctx,
		config:             cfg,
		Member:             NewMemberClient(cfg),
		PaPc:               NewPaPcClient(cfg),
		PaPcTechnicalScore: NewPaPcTechnicalScoreClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Member.
//		Query().
//		Count(ctx)
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
	c.Member.Use(hooks...)
	c.PaPc.Use(hooks...)
	c.PaPcTechnicalScore.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Member.Intercept(interceptors...)
	c.PaPc.Intercept(interceptors...)
	c.PaPcTechnicalScore.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *MemberMutation:
		return c.Member.mutate(ctx, m)
	case *PaPcMutation:
		return c.PaPc.mutate(ctx, m)
	case *PaPcTechnicalScoreMutation:
		return c.PaPcTechnicalScore.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// MemberClient is a client for the Member schema.
type MemberClient struct {
	config
}

// NewMemberClient returns a client for the Member from the given config.
func NewMemberClient(c config) *MemberClient {
	return &MemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `member.Hooks(f(g(h())))`.
func (c *MemberClient) Use(hooks ...Hook) {
	c.hooks.Member = append(c.hooks.Member, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `member.Intercept(f(g(h())))`.
func (c *MemberClient) Intercept(interceptors ...Interceptor) {
	c.inters.Member = append(c.inters.Member, interceptors...)
}

// Create returns a builder for creating a Member entity.
func (c *MemberClient) Create() *MemberCreate {
	mutation := newMemberMutation(c.config, OpCreate)
	return &MemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Member entities.
func (c *MemberClient) CreateBulk(builders ...*MemberCreate) *MemberCreateBulk {
	return &MemberCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MemberClient) MapCreateBulk(slice any, setFunc func(*MemberCreate, int)) *MemberCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MemberCreateBulk{err: fmt.Errorf("calling to MemberClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MemberCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Member.
func (c *MemberClient) Update() *MemberUpdate {
	mutation := newMemberMutation(c.config, OpUpdate)
	return &MemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MemberClient) UpdateOne(m *Member) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMember(m))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MemberClient) UpdateOneID(id int) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMemberID(id))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Member.
func (c *MemberClient) Delete() *MemberDelete {
	mutation := newMemberMutation(c.config, OpDelete)
	return &MemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MemberClient) DeleteOne(m *Member) *MemberDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MemberClient) DeleteOneID(id int) *MemberDeleteOne {
	builder := c.Delete().Where(member.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MemberDeleteOne{builder}
}

// Query returns a query builder for Member.
func (c *MemberClient) Query() *MemberQuery {
	return &MemberQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMember},
		inters: c.Interceptors(),
	}
}

// Get returns a Member entity by its id.
func (c *MemberClient) Get(ctx context.Context, id int) (*Member, error) {
	return c.Query().Where(member.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MemberClient) GetX(ctx context.Context, id int) *Member {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPaPcResults queries the pa_pc_results edge of a Member.
func (c *MemberClient) QueryPaPcResults(m *Member) *PaPcQuery {
	query := (&PaPcClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(papc.Table, papc.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.PaPcResultsTable, member.PaPcResultsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MemberClient) Hooks() []Hook {
	return c.hooks.Member
}

// Interceptors returns the client interceptors.
func (c *MemberClient) Interceptors() []Interceptor {
	return c.inters.Member
}

func (c *MemberClient) mutate(ctx context.Context, m *MemberMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MemberCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MemberUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MemberDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Member mutation op: %q", m.Op())
	}
}

// PaPcClient is a client for the PaPc schema.
type PaPcClient struct {
	config
}

// NewPaPcClient returns a client for the PaPc from the given config.
func NewPaPcClient(c config) *PaPcClient {
	return &PaPcClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `papc.Hooks(f(g(h())))`.
func (c *PaPcClient) Use(hooks ...Hook) {
	c.hooks.PaPc = append(c.hooks.PaPc, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `papc.Intercept(f(g(h())))`.
func (c *PaPcClient) Intercept(interceptors ...Interceptor) {
	c.inters.PaPc = append(c.inters.PaPc, interceptors...)
}

// Create returns a builder for creating a PaPc entity.
func (c *PaPcClient) Create() *PaPcCreate {
	mutation := newPaPcMutation(c.config, OpCreate)
	return &PaPcCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PaPc entities.
func (c *PaPcClient) CreateBulk(builders ...*PaPcCreate) *PaPcCreateBulk {
	return &PaPcCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PaPcClient) MapCreateBulk(slice any, setFunc func(*PaPcCreate, int)) *PaPcCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PaPcCreateBulk{err: fmt.Errorf("calling to PaPcClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PaPcCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PaPcCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PaPc.
func (c *PaPcClient) Update() *PaPcUpdate {
	mutation := newPaPcMutation(c.config, OpUpdate)
	return &PaPcUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PaPcClient) UpdateOne(pp *PaPc) *PaPcUpdateOne {
	mutation := newPaPcMutation(c.config, OpUpdateOne, withPaPc(pp))
	return &PaPcUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PaPcClient) UpdateOneID(id int) *PaPcUpdateOne {
	mutation := newPaPcMutation(c.config, OpUpdateOne, withPaPcID(id))
	return &PaPcUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PaPc.
func (c *PaPcClient) Delete() *PaPcDelete {
	mutation := newPaPcMutation(c.config, OpDelete)
	return &PaPcDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PaPcClient) DeleteOne(pp *PaPc) *PaPcDeleteOne {
	return c.DeleteOneID(pp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PaPcClient) DeleteOneID(id int) *PaPcDeleteOne {
	builder := c.Delete().Where(papc.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PaPcDeleteOne{builder}
}

// Query returns a query builder for PaPc.
func (c *PaPcClient) Query() *PaPcQuery {
	return &PaPcQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePaPc},
		inters: c.Interceptors(),
	}
}

// Get returns a PaPc entity by its id.
func (c *PaPcClient) Get(ctx context.Context, id int) (*PaPc, error) {
	return c.Query().Where(papc.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PaPcClient) GetX(ctx context.Context, id int) *PaPc {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMember queries the member edge of a PaPc.
func (c *PaPcClient) QueryMember(pp *PaPc) *MemberQuery {
	query := (&MemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(papc.Table, papc.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, papc.MemberTable, papc.MemberColumn),
		)
		fromV = sqlgraph.Neighbors(pp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTechnicalScoreDetails queries the technical_score_details edge of a PaPc.
func (c *PaPcClient) QueryTechnicalScoreDetails(pp *PaPc) *PaPcTechnicalScoreQuery {
	query := (&PaPcTechnicalScoreClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(papc.Table, papc.FieldID, id),
			sqlgraph.To(papctechnicalscore.Table, papctechnicalscore.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, papc.TechnicalScoreDetailsTable, papc.TechnicalScoreDetailsColumn),
		)
		fromV = sqlgraph.Neighbors(pp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PaPcClient) Hooks() []Hook {
	return c.hooks.PaPc
}

// Interceptors returns the client interceptors.
func (c *PaPcClient) Interceptors() []Interceptor {
	return c.inters.PaPc
}

func (c *PaPcClient) mutate(ctx context.Context, m *PaPcMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PaPcCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PaPcUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PaPcUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PaPcDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PaPc mutation op: %q", m.Op())
	}
}

// PaPcTechnicalScoreClient is a client for the PaPcTechnicalScore schema.
type PaPcTechnicalScoreClient struct {
	config
}

// NewPaPcTechnicalScoreClient returns a client for the PaPcTechnicalScore from the given config.
func NewPaPcTechnicalScoreClient(c config) *PaPcTechnicalScoreClient {
	return &PaPcTechnicalScoreClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `papctechnicalscore.Hooks(f(g(h())))`.
func (c *PaPcTechnicalScoreClient) Use(hooks ...Hook) {
	c.hooks.PaPcTechnicalScore = append(c.hooks.PaPcTechnicalScore, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `papctechnicalscore.Intercept(f(g(h())))`.
func (c *PaPcTechnicalScoreClient) Intercept(interceptors ...Interceptor) {
	c.inters.PaPcTechnicalScore = append(c.inters.PaPcTechnicalScore, interceptors...)
}

// Create returns a builder for creating a PaPcTechnicalScore entity.
func (c *PaPcTechnicalScoreClient) Create() *PaPcTechnicalScoreCreate {
	mutation := newPaPcTechnicalScoreMutation(c.config, OpCreate)
	return &PaPcTechnicalScoreCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PaPcTechnicalScore entities.
func (c *PaPcTechnicalScoreClient) CreateBulk(builders ...*PaPcTechnicalScoreCreate) *PaPcTechnicalScoreCreateBulk {
	return &PaPcTechnicalScoreCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PaPcTechnicalScoreClient) MapCreateBulk(slice any, setFunc func(*PaPcTechnicalScoreCreate, int)) *PaPcTechnicalScoreCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PaPcTechnicalScoreCreateBulk{err: fmt.Errorf("calling to PaPcTechnicalScoreClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PaPcTechnicalScoreCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PaPcTechnicalScoreCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PaPcTechnicalScore.
func (c *PaPcTechnicalScoreClient) Update() *PaPcTechnicalScoreUpdate {
	mutation := newPaPcTechnicalScoreMutation(c.config, OpUpdate)
	return &PaPcTechnicalScoreUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PaPcTechnicalScoreClient) UpdateOne(ppts *PaPcTechnicalScore) *PaPcTechnicalScoreUpdateOne {
	mutation := newPaPcTechnicalScoreMutation(c.config, OpUpdateOne, withPaPcTechnicalScore(ppts))
	return &PaPcTechnicalScoreUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PaPcTechnicalScoreClient) UpdateOneID(id int) *PaPcTechnicalScoreUpdateOne {
	mutation := newPaPcTechnicalScoreMutation(c.config, OpUpdateOne, withPaPcTechnicalScoreID(id))
	return &PaPcTechnicalScoreUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PaPcTechnicalScore.
func (c *PaPcTechnicalScoreClient) Delete() *PaPcTechnicalScoreDelete {
	mutation := newPaPcTechnicalScoreMutation(c.config, OpDelete)
	return &PaPcTechnicalScoreDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PaPcTechnicalScoreClient) DeleteOne(ppts *PaPcTechnicalScore) *PaPcTechnicalScoreDeleteOne {
	return c.DeleteOneID(ppts.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PaPcTechnicalScoreClient) DeleteOneID(id int) *PaPcTechnicalScoreDeleteOne {
	builder := c.Delete().Where(papctechnicalscore.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PaPcTechnicalScoreDeleteOne{builder}
}

// Query returns a query builder for PaPcTechnicalScore.
func (c *PaPcTechnicalScoreClient) Query() *PaPcTechnicalScoreQuery {
	return &PaPcTechnicalScoreQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePaPcTechnicalScore},
		inters: c.Interceptors(),
	}
}

// Get returns a PaPcTechnicalScore entity by its id.
func (c *PaPcTechnicalScoreClient) Get(ctx context.Context, id int) (*PaPcTechnicalScore, error) {
	return c.Query().Where(papctechnicalscore.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PaPcTechnicalScoreClient) GetX(ctx context.Context, id int) *PaPcTechnicalScore {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPaPc queries the pa_pc edge of a PaPcTechnicalScore.
func (c *PaPcTechnicalScoreClient) QueryPaPc(ppts *PaPcTechnicalScore) *PaPcQuery {
	query := (&PaPcClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ppts.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(papctechnicalscore.Table, papctechnicalscore.FieldID, id),
			sqlgraph.To(papc.Table, papc.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, papctechnicalscore.PaPcTable, papctechnicalscore.PaPcColumn),
		)
		fromV = sqlgraph.Neighbors(ppts.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PaPcTechnicalScoreClient) Hooks() []Hook {
	return c.hooks.PaPcTechnicalScore
}

// Interceptors returns the client interceptors.
func (c *PaPcTechnicalScoreClient) Interceptors() []Interceptor {
	return c.inters.PaPcTechnicalScore
}

func (c *PaPcTechnicalScoreClient) mutate(ctx context.Context, m *PaPcTechnicalScoreMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PaPcTechnicalScoreCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PaPcTechnicalScoreUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PaPcTechnicalScoreUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PaPcTechnicalScoreDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PaPcTechnicalScore mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Member, PaPc, PaPcTechnicalScore []ent.Hook
	}
	inters struct {
		Member, PaPc, PaPcTechnicalScore []ent.Interceptor
	}
)
