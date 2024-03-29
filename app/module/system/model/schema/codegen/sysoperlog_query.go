// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/predicate"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysoperlog"
)

// SysOperLogQuery is the builder for querying SysOperLog entities.
type SysOperLogQuery struct {
	config
	ctx        *QueryContext
	order      []sysoperlog.OrderOption
	inters     []Interceptor
	predicates []predicate.SysOperLog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysOperLogQuery builder.
func (solq *SysOperLogQuery) Where(ps ...predicate.SysOperLog) *SysOperLogQuery {
	solq.predicates = append(solq.predicates, ps...)
	return solq
}

// Limit the number of records to be returned by this query.
func (solq *SysOperLogQuery) Limit(limit int) *SysOperLogQuery {
	solq.ctx.Limit = &limit
	return solq
}

// Offset to start from.
func (solq *SysOperLogQuery) Offset(offset int) *SysOperLogQuery {
	solq.ctx.Offset = &offset
	return solq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (solq *SysOperLogQuery) Unique(unique bool) *SysOperLogQuery {
	solq.ctx.Unique = &unique
	return solq
}

// Order specifies how the records should be ordered.
func (solq *SysOperLogQuery) Order(o ...sysoperlog.OrderOption) *SysOperLogQuery {
	solq.order = append(solq.order, o...)
	return solq
}

// First returns the first SysOperLog entity from the query.
// Returns a *NotFoundError when no SysOperLog was found.
func (solq *SysOperLogQuery) First(ctx context.Context) (*SysOperLog, error) {
	nodes, err := solq.Limit(1).All(setContextOp(ctx, solq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysoperlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (solq *SysOperLogQuery) FirstX(ctx context.Context) *SysOperLog {
	node, err := solq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysOperLog ID from the query.
// Returns a *NotFoundError when no SysOperLog ID was found.
func (solq *SysOperLogQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = solq.Limit(1).IDs(setContextOp(ctx, solq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysoperlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (solq *SysOperLogQuery) FirstIDX(ctx context.Context) int64 {
	id, err := solq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysOperLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysOperLog entity is found.
// Returns a *NotFoundError when no SysOperLog entities are found.
func (solq *SysOperLogQuery) Only(ctx context.Context) (*SysOperLog, error) {
	nodes, err := solq.Limit(2).All(setContextOp(ctx, solq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysoperlog.Label}
	default:
		return nil, &NotSingularError{sysoperlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (solq *SysOperLogQuery) OnlyX(ctx context.Context) *SysOperLog {
	node, err := solq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysOperLog ID in the query.
// Returns a *NotSingularError when more than one SysOperLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (solq *SysOperLogQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = solq.Limit(2).IDs(setContextOp(ctx, solq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysoperlog.Label}
	default:
		err = &NotSingularError{sysoperlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (solq *SysOperLogQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := solq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysOperLogs.
func (solq *SysOperLogQuery) All(ctx context.Context) ([]*SysOperLog, error) {
	ctx = setContextOp(ctx, solq.ctx, "All")
	if err := solq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysOperLog, *SysOperLogQuery]()
	return withInterceptors[[]*SysOperLog](ctx, solq, qr, solq.inters)
}

// AllX is like All, but panics if an error occurs.
func (solq *SysOperLogQuery) AllX(ctx context.Context) []*SysOperLog {
	nodes, err := solq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysOperLog IDs.
func (solq *SysOperLogQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if solq.ctx.Unique == nil && solq.path != nil {
		solq.Unique(true)
	}
	ctx = setContextOp(ctx, solq.ctx, "IDs")
	if err = solq.Select(sysoperlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (solq *SysOperLogQuery) IDsX(ctx context.Context) []int64 {
	ids, err := solq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (solq *SysOperLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, solq.ctx, "Count")
	if err := solq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, solq, querierCount[*SysOperLogQuery](), solq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (solq *SysOperLogQuery) CountX(ctx context.Context) int {
	count, err := solq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (solq *SysOperLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, solq.ctx, "Exist")
	switch _, err := solq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("codegen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (solq *SysOperLogQuery) ExistX(ctx context.Context) bool {
	exist, err := solq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysOperLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (solq *SysOperLogQuery) Clone() *SysOperLogQuery {
	if solq == nil {
		return nil
	}
	return &SysOperLogQuery{
		config:     solq.config,
		ctx:        solq.ctx.Clone(),
		order:      append([]sysoperlog.OrderOption{}, solq.order...),
		inters:     append([]Interceptor{}, solq.inters...),
		predicates: append([]predicate.SysOperLog{}, solq.predicates...),
		// clone intermediate query.
		sql:  solq.sql.Clone(),
		path: solq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysOperLog.Query().
//		GroupBy(sysoperlog.FieldTitle).
//		Aggregate(codegen.Count()).
//		Scan(ctx, &v)
func (solq *SysOperLogQuery) GroupBy(field string, fields ...string) *SysOperLogGroupBy {
	solq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysOperLogGroupBy{build: solq}
	grbuild.flds = &solq.ctx.Fields
	grbuild.label = sysoperlog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.SysOperLog.Query().
//		Select(sysoperlog.FieldTitle).
//		Scan(ctx, &v)
func (solq *SysOperLogQuery) Select(fields ...string) *SysOperLogSelect {
	solq.ctx.Fields = append(solq.ctx.Fields, fields...)
	sbuild := &SysOperLogSelect{SysOperLogQuery: solq}
	sbuild.label = sysoperlog.Label
	sbuild.flds, sbuild.scan = &solq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysOperLogSelect configured with the given aggregations.
func (solq *SysOperLogQuery) Aggregate(fns ...AggregateFunc) *SysOperLogSelect {
	return solq.Select().Aggregate(fns...)
}

func (solq *SysOperLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range solq.inters {
		if inter == nil {
			return fmt.Errorf("codegen: uninitialized interceptor (forgotten import codegen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, solq); err != nil {
				return err
			}
		}
	}
	for _, f := range solq.ctx.Fields {
		if !sysoperlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
		}
	}
	if solq.path != nil {
		prev, err := solq.path(ctx)
		if err != nil {
			return err
		}
		solq.sql = prev
	}
	return nil
}

func (solq *SysOperLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysOperLog, error) {
	var (
		nodes = []*SysOperLog{}
		_spec = solq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysOperLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysOperLog{config: solq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, solq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (solq *SysOperLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := solq.querySpec()
	_spec.Node.Columns = solq.ctx.Fields
	if len(solq.ctx.Fields) > 0 {
		_spec.Unique = solq.ctx.Unique != nil && *solq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, solq.driver, _spec)
}

func (solq *SysOperLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysoperlog.Table, sysoperlog.Columns, sqlgraph.NewFieldSpec(sysoperlog.FieldID, field.TypeInt64))
	_spec.From = solq.sql
	if unique := solq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if solq.path != nil {
		_spec.Unique = true
	}
	if fields := solq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysoperlog.FieldID)
		for i := range fields {
			if fields[i] != sysoperlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := solq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := solq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := solq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := solq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (solq *SysOperLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(solq.driver.Dialect())
	t1 := builder.Table(sysoperlog.Table)
	columns := solq.ctx.Fields
	if len(columns) == 0 {
		columns = sysoperlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if solq.sql != nil {
		selector = solq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if solq.ctx.Unique != nil && *solq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range solq.predicates {
		p(selector)
	}
	for _, p := range solq.order {
		p(selector)
	}
	if offset := solq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := solq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysOperLogGroupBy is the group-by builder for SysOperLog entities.
type SysOperLogGroupBy struct {
	selector
	build *SysOperLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (solgb *SysOperLogGroupBy) Aggregate(fns ...AggregateFunc) *SysOperLogGroupBy {
	solgb.fns = append(solgb.fns, fns...)
	return solgb
}

// Scan applies the selector query and scans the result into the given value.
func (solgb *SysOperLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, solgb.build.ctx, "GroupBy")
	if err := solgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysOperLogQuery, *SysOperLogGroupBy](ctx, solgb.build, solgb, solgb.build.inters, v)
}

func (solgb *SysOperLogGroupBy) sqlScan(ctx context.Context, root *SysOperLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(solgb.fns))
	for _, fn := range solgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*solgb.flds)+len(solgb.fns))
		for _, f := range *solgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*solgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := solgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysOperLogSelect is the builder for selecting fields of SysOperLog entities.
type SysOperLogSelect struct {
	*SysOperLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sols *SysOperLogSelect) Aggregate(fns ...AggregateFunc) *SysOperLogSelect {
	sols.fns = append(sols.fns, fns...)
	return sols
}

// Scan applies the selector query and scans the result into the given value.
func (sols *SysOperLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sols.ctx, "Select")
	if err := sols.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysOperLogQuery, *SysOperLogSelect](ctx, sols.SysOperLogQuery, sols, sols.inters, v)
}

func (sols *SysOperLogSelect) sqlScan(ctx context.Context, root *SysOperLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sols.fns))
	for _, fn := range sols.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sols.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sols.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
