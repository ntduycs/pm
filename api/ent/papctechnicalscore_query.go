// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"project-management/ent/papc"
	"project-management/ent/papctechnicalscore"
	"project-management/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaPcTechnicalScoreQuery is the builder for querying PaPcTechnicalScore entities.
type PaPcTechnicalScoreQuery struct {
	config
	ctx        *QueryContext
	order      []papctechnicalscore.OrderOption
	inters     []Interceptor
	predicates []predicate.PaPcTechnicalScore
	withPaPc   *PaPcQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PaPcTechnicalScoreQuery builder.
func (pptsq *PaPcTechnicalScoreQuery) Where(ps ...predicate.PaPcTechnicalScore) *PaPcTechnicalScoreQuery {
	pptsq.predicates = append(pptsq.predicates, ps...)
	return pptsq
}

// Limit the number of records to be returned by this query.
func (pptsq *PaPcTechnicalScoreQuery) Limit(limit int) *PaPcTechnicalScoreQuery {
	pptsq.ctx.Limit = &limit
	return pptsq
}

// Offset to start from.
func (pptsq *PaPcTechnicalScoreQuery) Offset(offset int) *PaPcTechnicalScoreQuery {
	pptsq.ctx.Offset = &offset
	return pptsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pptsq *PaPcTechnicalScoreQuery) Unique(unique bool) *PaPcTechnicalScoreQuery {
	pptsq.ctx.Unique = &unique
	return pptsq
}

// Order specifies how the records should be ordered.
func (pptsq *PaPcTechnicalScoreQuery) Order(o ...papctechnicalscore.OrderOption) *PaPcTechnicalScoreQuery {
	pptsq.order = append(pptsq.order, o...)
	return pptsq
}

// QueryPaPc chains the current query on the "pa_pc" edge.
func (pptsq *PaPcTechnicalScoreQuery) QueryPaPc() *PaPcQuery {
	query := (&PaPcClient{config: pptsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pptsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pptsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(papctechnicalscore.Table, papctechnicalscore.FieldID, selector),
			sqlgraph.To(papc.Table, papc.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, papctechnicalscore.PaPcTable, papctechnicalscore.PaPcColumn),
		)
		fromU = sqlgraph.SetNeighbors(pptsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PaPcTechnicalScore entity from the query.
// Returns a *NotFoundError when no PaPcTechnicalScore was found.
func (pptsq *PaPcTechnicalScoreQuery) First(ctx context.Context) (*PaPcTechnicalScore, error) {
	nodes, err := pptsq.Limit(1).All(setContextOp(ctx, pptsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{papctechnicalscore.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pptsq *PaPcTechnicalScoreQuery) FirstX(ctx context.Context) *PaPcTechnicalScore {
	node, err := pptsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PaPcTechnicalScore ID from the query.
// Returns a *NotFoundError when no PaPcTechnicalScore ID was found.
func (pptsq *PaPcTechnicalScoreQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pptsq.Limit(1).IDs(setContextOp(ctx, pptsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{papctechnicalscore.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pptsq *PaPcTechnicalScoreQuery) FirstIDX(ctx context.Context) int {
	id, err := pptsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PaPcTechnicalScore entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PaPcTechnicalScore entity is found.
// Returns a *NotFoundError when no PaPcTechnicalScore entities are found.
func (pptsq *PaPcTechnicalScoreQuery) Only(ctx context.Context) (*PaPcTechnicalScore, error) {
	nodes, err := pptsq.Limit(2).All(setContextOp(ctx, pptsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{papctechnicalscore.Label}
	default:
		return nil, &NotSingularError{papctechnicalscore.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pptsq *PaPcTechnicalScoreQuery) OnlyX(ctx context.Context) *PaPcTechnicalScore {
	node, err := pptsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PaPcTechnicalScore ID in the query.
// Returns a *NotSingularError when more than one PaPcTechnicalScore ID is found.
// Returns a *NotFoundError when no entities are found.
func (pptsq *PaPcTechnicalScoreQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pptsq.Limit(2).IDs(setContextOp(ctx, pptsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{papctechnicalscore.Label}
	default:
		err = &NotSingularError{papctechnicalscore.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pptsq *PaPcTechnicalScoreQuery) OnlyIDX(ctx context.Context) int {
	id, err := pptsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PaPcTechnicalScores.
func (pptsq *PaPcTechnicalScoreQuery) All(ctx context.Context) ([]*PaPcTechnicalScore, error) {
	ctx = setContextOp(ctx, pptsq.ctx, "All")
	if err := pptsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PaPcTechnicalScore, *PaPcTechnicalScoreQuery]()
	return withInterceptors[[]*PaPcTechnicalScore](ctx, pptsq, qr, pptsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pptsq *PaPcTechnicalScoreQuery) AllX(ctx context.Context) []*PaPcTechnicalScore {
	nodes, err := pptsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PaPcTechnicalScore IDs.
func (pptsq *PaPcTechnicalScoreQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pptsq.ctx.Unique == nil && pptsq.path != nil {
		pptsq.Unique(true)
	}
	ctx = setContextOp(ctx, pptsq.ctx, "IDs")
	if err = pptsq.Select(papctechnicalscore.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pptsq *PaPcTechnicalScoreQuery) IDsX(ctx context.Context) []int {
	ids, err := pptsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pptsq *PaPcTechnicalScoreQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pptsq.ctx, "Count")
	if err := pptsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pptsq, querierCount[*PaPcTechnicalScoreQuery](), pptsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pptsq *PaPcTechnicalScoreQuery) CountX(ctx context.Context) int {
	count, err := pptsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pptsq *PaPcTechnicalScoreQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pptsq.ctx, "Exist")
	switch _, err := pptsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pptsq *PaPcTechnicalScoreQuery) ExistX(ctx context.Context) bool {
	exist, err := pptsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PaPcTechnicalScoreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pptsq *PaPcTechnicalScoreQuery) Clone() *PaPcTechnicalScoreQuery {
	if pptsq == nil {
		return nil
	}
	return &PaPcTechnicalScoreQuery{
		config:     pptsq.config,
		ctx:        pptsq.ctx.Clone(),
		order:      append([]papctechnicalscore.OrderOption{}, pptsq.order...),
		inters:     append([]Interceptor{}, pptsq.inters...),
		predicates: append([]predicate.PaPcTechnicalScore{}, pptsq.predicates...),
		withPaPc:   pptsq.withPaPc.Clone(),
		// clone intermediate query.
		sql:  pptsq.sql.Clone(),
		path: pptsq.path,
	}
}

// WithPaPc tells the query-builder to eager-load the nodes that are connected to
// the "pa_pc" edge. The optional arguments are used to configure the query builder of the edge.
func (pptsq *PaPcTechnicalScoreQuery) WithPaPc(opts ...func(*PaPcQuery)) *PaPcTechnicalScoreQuery {
	query := (&PaPcClient{config: pptsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pptsq.withPaPc = query
	return pptsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PaPcID int `json:"pa_pc_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PaPcTechnicalScore.Query().
//		GroupBy(papctechnicalscore.FieldPaPcID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pptsq *PaPcTechnicalScoreQuery) GroupBy(field string, fields ...string) *PaPcTechnicalScoreGroupBy {
	pptsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PaPcTechnicalScoreGroupBy{build: pptsq}
	grbuild.flds = &pptsq.ctx.Fields
	grbuild.label = papctechnicalscore.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PaPcID int `json:"pa_pc_id,omitempty"`
//	}
//
//	client.PaPcTechnicalScore.Query().
//		Select(papctechnicalscore.FieldPaPcID).
//		Scan(ctx, &v)
func (pptsq *PaPcTechnicalScoreQuery) Select(fields ...string) *PaPcTechnicalScoreSelect {
	pptsq.ctx.Fields = append(pptsq.ctx.Fields, fields...)
	sbuild := &PaPcTechnicalScoreSelect{PaPcTechnicalScoreQuery: pptsq}
	sbuild.label = papctechnicalscore.Label
	sbuild.flds, sbuild.scan = &pptsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PaPcTechnicalScoreSelect configured with the given aggregations.
func (pptsq *PaPcTechnicalScoreQuery) Aggregate(fns ...AggregateFunc) *PaPcTechnicalScoreSelect {
	return pptsq.Select().Aggregate(fns...)
}

func (pptsq *PaPcTechnicalScoreQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pptsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pptsq); err != nil {
				return err
			}
		}
	}
	for _, f := range pptsq.ctx.Fields {
		if !papctechnicalscore.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pptsq.path != nil {
		prev, err := pptsq.path(ctx)
		if err != nil {
			return err
		}
		pptsq.sql = prev
	}
	return nil
}

func (pptsq *PaPcTechnicalScoreQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PaPcTechnicalScore, error) {
	var (
		nodes       = []*PaPcTechnicalScore{}
		_spec       = pptsq.querySpec()
		loadedTypes = [1]bool{
			pptsq.withPaPc != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PaPcTechnicalScore).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PaPcTechnicalScore{config: pptsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pptsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pptsq.withPaPc; query != nil {
		if err := pptsq.loadPaPc(ctx, query, nodes, nil,
			func(n *PaPcTechnicalScore, e *PaPc) { n.Edges.PaPc = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pptsq *PaPcTechnicalScoreQuery) loadPaPc(ctx context.Context, query *PaPcQuery, nodes []*PaPcTechnicalScore, init func(*PaPcTechnicalScore), assign func(*PaPcTechnicalScore, *PaPc)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PaPcTechnicalScore)
	for i := range nodes {
		fk := nodes[i].PaPcID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(papc.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "pa_pc_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pptsq *PaPcTechnicalScoreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pptsq.querySpec()
	_spec.Node.Columns = pptsq.ctx.Fields
	if len(pptsq.ctx.Fields) > 0 {
		_spec.Unique = pptsq.ctx.Unique != nil && *pptsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pptsq.driver, _spec)
}

func (pptsq *PaPcTechnicalScoreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(papctechnicalscore.Table, papctechnicalscore.Columns, sqlgraph.NewFieldSpec(papctechnicalscore.FieldID, field.TypeInt))
	_spec.From = pptsq.sql
	if unique := pptsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pptsq.path != nil {
		_spec.Unique = true
	}
	if fields := pptsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, papctechnicalscore.FieldID)
		for i := range fields {
			if fields[i] != papctechnicalscore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pptsq.withPaPc != nil {
			_spec.Node.AddColumnOnce(papctechnicalscore.FieldPaPcID)
		}
	}
	if ps := pptsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pptsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pptsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pptsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pptsq *PaPcTechnicalScoreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pptsq.driver.Dialect())
	t1 := builder.Table(papctechnicalscore.Table)
	columns := pptsq.ctx.Fields
	if len(columns) == 0 {
		columns = papctechnicalscore.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pptsq.sql != nil {
		selector = pptsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pptsq.ctx.Unique != nil && *pptsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pptsq.predicates {
		p(selector)
	}
	for _, p := range pptsq.order {
		p(selector)
	}
	if offset := pptsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pptsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PaPcTechnicalScoreGroupBy is the group-by builder for PaPcTechnicalScore entities.
type PaPcTechnicalScoreGroupBy struct {
	selector
	build *PaPcTechnicalScoreQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pptsgb *PaPcTechnicalScoreGroupBy) Aggregate(fns ...AggregateFunc) *PaPcTechnicalScoreGroupBy {
	pptsgb.fns = append(pptsgb.fns, fns...)
	return pptsgb
}

// Scan applies the selector query and scans the result into the given value.
func (pptsgb *PaPcTechnicalScoreGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pptsgb.build.ctx, "GroupBy")
	if err := pptsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaPcTechnicalScoreQuery, *PaPcTechnicalScoreGroupBy](ctx, pptsgb.build, pptsgb, pptsgb.build.inters, v)
}

func (pptsgb *PaPcTechnicalScoreGroupBy) sqlScan(ctx context.Context, root *PaPcTechnicalScoreQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pptsgb.fns))
	for _, fn := range pptsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pptsgb.flds)+len(pptsgb.fns))
		for _, f := range *pptsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pptsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pptsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PaPcTechnicalScoreSelect is the builder for selecting fields of PaPcTechnicalScore entities.
type PaPcTechnicalScoreSelect struct {
	*PaPcTechnicalScoreQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pptss *PaPcTechnicalScoreSelect) Aggregate(fns ...AggregateFunc) *PaPcTechnicalScoreSelect {
	pptss.fns = append(pptss.fns, fns...)
	return pptss
}

// Scan applies the selector query and scans the result into the given value.
func (pptss *PaPcTechnicalScoreSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pptss.ctx, "Select")
	if err := pptss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PaPcTechnicalScoreQuery, *PaPcTechnicalScoreSelect](ctx, pptss.PaPcTechnicalScoreQuery, pptss, pptss.inters, v)
}

func (pptss *PaPcTechnicalScoreSelect) sqlScan(ctx context.Context, root *PaPcTechnicalScoreQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pptss.fns))
	for _, fn := range pptss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pptss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pptss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}