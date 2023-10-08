// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"project-management/ent/member"
	"project-management/ent/papc"
	"project-management/ent/papctechnicalscore"
	"project-management/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaPcQuery is the builder for querying PaPc entities.
type PaPcQuery struct {
	config
	ctx                       *QueryContext
	order                     []papc.OrderOption
	inters                    []Interceptor
	predicates                []predicate.PaPc
	withMember                *MemberQuery
	withTechnicalScoreDetails *PaPcTechnicalScoreQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PaPcQuery builder.
func (ppq *PaPcQuery) Where(ps ...predicate.PaPc) *PaPcQuery {
	ppq.predicates = append(ppq.predicates, ps...)
	return ppq
}

// Limit the number of records to be returned by this query.
func (ppq *PaPcQuery) Limit(limit int) *PaPcQuery {
	ppq.ctx.Limit = &limit
	return ppq
}

// Offset to start from.
func (ppq *PaPcQuery) Offset(offset int) *PaPcQuery {
	ppq.ctx.Offset = &offset
	return ppq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ppq *PaPcQuery) Unique(unique bool) *PaPcQuery {
	ppq.ctx.Unique = &unique
	return ppq
}

// Order specifies how the records should be ordered.
func (ppq *PaPcQuery) Order(o ...papc.OrderOption) *PaPcQuery {
	ppq.order = append(ppq.order, o...)
	return ppq
}

// QueryMember chains the current query on the "member" edge.
func (ppq *PaPcQuery) QueryMember() *MemberQuery {
	query := (&MemberClient{config: ppq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(papc.Table, papc.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, papc.MemberTable, papc.MemberColumn),
		)
		fromU = sqlgraph.SetNeighbors(ppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTechnicalScoreDetails chains the current query on the "technical_score_details" edge.
func (ppq *PaPcQuery) QueryTechnicalScoreDetails() *PaPcTechnicalScoreQuery {
	query := (&PaPcTechnicalScoreClient{config: ppq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(papc.Table, papc.FieldID, selector),
			sqlgraph.To(papctechnicalscore.Table, papctechnicalscore.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, papc.TechnicalScoreDetailsTable, papc.TechnicalScoreDetailsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PaPc entity from the query.
// Returns a *NotFoundError when no PaPc was found.
func (ppq *PaPcQuery) First(ctx context.Context) (*PaPc, error) {
	nodes, err := ppq.Limit(1).All(setContextOp(ctx, ppq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{papc.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ppq *PaPcQuery) FirstX(ctx context.Context) *PaPc {
	node, err := ppq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PaPc ID from the query.
// Returns a *NotFoundError when no PaPc ID was found.
func (ppq *PaPcQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ppq.Limit(1).IDs(setContextOp(ctx, ppq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{papc.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ppq *PaPcQuery) FirstIDX(ctx context.Context) int {
	id, err := ppq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PaPc entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PaPc entity is found.
// Returns a *NotFoundError when no PaPc entities are found.
func (ppq *PaPcQuery) Only(ctx context.Context) (*PaPc, error) {
	nodes, err := ppq.Limit(2).All(setContextOp(ctx, ppq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{papc.Label}
	default:
		return nil, &NotSingularError{papc.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ppq *PaPcQuery) OnlyX(ctx context.Context) *PaPc {
	node, err := ppq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PaPc ID in the query.
// Returns a *NotSingularError when more than one PaPc ID is found.
// Returns a *NotFoundError when no entities are found.
func (ppq *PaPcQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ppq.Limit(2).IDs(setContextOp(ctx, ppq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{papc.Label}
	default:
		err = &NotSingularError{papc.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ppq *PaPcQuery) OnlyIDX(ctx context.Context) int {
	id, err := ppq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PaPcs.
func (ppq *PaPcQuery) All(ctx context.Context) ([]*PaPc, error) {
	ctx = setContextOp(ctx, ppq.ctx, "All")
	if err := ppq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PaPc, *PaPcQuery]()
	return withInterceptors[[]*PaPc](ctx, ppq, qr, ppq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ppq *PaPcQuery) AllX(ctx context.Context) []*PaPc {
	nodes, err := ppq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PaPc IDs.
func (ppq *PaPcQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ppq.ctx.Unique == nil && ppq.path != nil {
		ppq.Unique(true)
	}
	ctx = setContextOp(ctx, ppq.ctx, "IDs")
	if err = ppq.Select(papc.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ppq *PaPcQuery) IDsX(ctx context.Context) []int {
	ids, err := ppq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ppq *PaPcQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ppq.ctx, "Count")
	if err := ppq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ppq, querierCount[*PaPcQuery](), ppq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ppq *PaPcQuery) CountX(ctx context.Context) int {
	count, err := ppq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ppq *PaPcQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ppq.ctx, "Exist")
	switch _, err := ppq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ppq *PaPcQuery) ExistX(ctx context.Context) bool {
	exist, err := ppq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PaPcQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ppq *PaPcQuery) Clone() *PaPcQuery {
	if ppq == nil {
		return nil
	}
	return &PaPcQuery{
		config:                    ppq.config,
		ctx:                       ppq.ctx.Clone(),
		order:                     append([]papc.OrderOption{}, ppq.order...),
		inters:                    append([]Interceptor{}, ppq.inters...),
		predicates:                append([]predicate.PaPc{}, ppq.predicates...),
		withMember:                ppq.withMember.Clone(),
		withTechnicalScoreDetails: ppq.withTechnicalScoreDetails.Clone(),
		// clone intermediate query.
		sql:  ppq.sql.Clone(),
		path: ppq.path,
	}
}

// WithMember tells the query-builder to eager-load the nodes that are connected to
// the "member" edge. The optional arguments are used to configure the query builder of the edge.
func (ppq *PaPcQuery) WithMember(opts ...func(*MemberQuery)) *PaPcQuery {
	query := (&MemberClient{config: ppq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ppq.withMember = query
	return ppq
}

// WithTechnicalScoreDetails tells the query-builder to eager-load the nodes that are connected to
// the "technical_score_details" edge. The optional arguments are used to configure the query builder of the edge.
func (ppq *PaPcQuery) WithTechnicalScoreDetails(opts ...func(*PaPcTechnicalScoreQuery)) *PaPcQuery {
	query := (&PaPcTechnicalScoreClient{config: ppq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ppq.withTechnicalScoreDetails = query
	return ppq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MemberID int `json:"member_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PaPc.Query().
//		GroupBy(papc.FieldMemberID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ppq *PaPcQuery) GroupBy(field string, fields ...string) *PaPcGroupBy {
	ppq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PaPcGroupBy{build: ppq}
	grbuild.flds = &ppq.ctx.Fields
	grbuild.label = papc.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MemberID int `json:"member_id,omitempty"`
//	}
//
//	client.PaPc.Query().
//		Select(papc.FieldMemberID).
//		Scan(ctx, &v)
func (ppq *PaPcQuery) Select(fields ...string) *PaPcSelect {
	ppq.ctx.Fields = append(ppq.ctx.Fields, fields...)
	sbuild := &PaPcSelect{PaPcQuery: ppq}
	sbuild.label = papc.Label
	sbuild.flds, sbuild.scan = &ppq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PaPcSelect configured with the given aggregations.
func (ppq *PaPcQuery) Aggregate(fns ...AggregateFunc) *PaPcSelect {
	return ppq.Select().Aggregate(fns...)
}

func (ppq *PaPcQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ppq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ppq); err != nil {
				return err
			}
		}
	}
	for _, f := range ppq.ctx.Fields {
		if !papc.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ppq.path != nil {
		prev, err := ppq.path(ctx)
		if err != nil {
			return err
		}
		ppq.sql = prev
	}
	return nil
}

func (ppq *PaPcQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PaPc, error) {
	var (
		nodes       = []*PaPc{}
		_spec       = ppq.querySpec()
		loadedTypes = [2]bool{
			ppq.withMember != nil,
			ppq.withTechnicalScoreDetails != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PaPc).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PaPc{config: ppq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ppq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ppq.withMember; query != nil {
		if err := ppq.loadMember(ctx, query, nodes, nil,
			func(n *PaPc, e *Member) { n.Edges.Member = e }); err != nil {
			return nil, err
		}
	}
	if query := ppq.withTechnicalScoreDetails; query != nil {
		if err := ppq.loadTechnicalScoreDetails(ctx, query, nodes,
			func(n *PaPc) { n.Edges.TechnicalScoreDetails = []*PaPcTechnicalScore{} },
			func(n *PaPc, e *PaPcTechnicalScore) {
				n.Edges.TechnicalScoreDetails = append(n.Edges.TechnicalScoreDetails, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ppq *PaPcQuery) loadMember(ctx context.Context, query *MemberQuery, nodes []*PaPc, init func(*PaPc), assign func(*PaPc, *Member)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PaPc)
	for i := range nodes {
		fk := nodes[i].MemberID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(member.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ppq *PaPcQuery) loadTechnicalScoreDetails(ctx context.Context, query *PaPcTechnicalScoreQuery, nodes []*PaPc, init func(*PaPc), assign func(*PaPc, *PaPcTechnicalScore)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*PaPc)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(papctechnicalscore.FieldPaPcID)
	}
	query.Where(predicate.PaPcTechnicalScore(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(papc.TechnicalScoreDetailsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PaPcID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "pa_pc_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ppq *PaPcQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ppq.querySpec()
	_spec.Node.Columns = ppq.ctx.Fields
	if len(ppq.ctx.Fields) > 0 {
		_spec.Unique = ppq.ctx.Unique != nil && *ppq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ppq.driver, _spec)
}

func (ppq *PaPcQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(papc.Table, papc.Columns, sqlgraph.NewFieldSpec(papc.FieldID, field.TypeInt))
	_spec.From = ppq.sql
	if unique := ppq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ppq.path != nil {
		_spec.Unique = true
	}
	if fields := ppq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, papc.FieldID)
		for i := range fields {
			if fields[i] != papc.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ppq.withMember != nil {
			_spec.Node.AddColumnOnce(papc.FieldMemberID)
		}
	}
	if ps := ppq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ppq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ppq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ppq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ppq *PaPcQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ppq.driver.Dialect())
	t1 := builder.Table(papc.Table)
	columns := ppq.ctx.Fields
	if len(columns) == 0 {
		columns = papc.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ppq.sql != nil {
		selector = ppq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ppq.ctx.Unique != nil && *ppq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ppq.predicates {
		p(selector)
	}
	for _, p := range ppq.order {
		p(selector)
	}
	if offset := ppq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ppq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PaPcGroupBy is the group-by builder for PaPc entities.
type PaPcGroupBy struct {
	selector
	build *PaPcQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ppgb *PaPcGroupBy) Aggregate(fns ...AggregateFunc) *PaPcGroupBy {
	ppgb.fns = append(ppgb.fns, fns...)
	return ppgb
}

// Scan applies the selector query and scans the result into the given value.
func (ppgb *PaPcGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ppgb.build.ctx, "GroupBy")
	if err := ppgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaPcQuery, *PaPcGroupBy](ctx, ppgb.build, ppgb, ppgb.build.inters, v)
}

func (ppgb *PaPcGroupBy) sqlScan(ctx context.Context, root *PaPcQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ppgb.fns))
	for _, fn := range ppgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ppgb.flds)+len(ppgb.fns))
		for _, f := range *ppgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ppgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ppgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PaPcSelect is the builder for selecting fields of PaPc entities.
type PaPcSelect struct {
	*PaPcQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pps *PaPcSelect) Aggregate(fns ...AggregateFunc) *PaPcSelect {
	pps.fns = append(pps.fns, fns...)
	return pps
}

// Scan applies the selector query and scans the result into the given value.
func (pps *PaPcSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pps.ctx, "Select")
	if err := pps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaPcQuery, *PaPcSelect](ctx, pps.PaPcQuery, pps, pps.inters, v)
}

func (pps *PaPcSelect) sqlScan(ctx context.Context, root *PaPcQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pps.fns))
	for _, fn := range pps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}