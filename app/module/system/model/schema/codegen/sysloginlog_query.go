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
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysloginlog"
)

// SysLoginLogQuery is the builder for querying SysLoginLog entities.
type SysLoginLogQuery struct {
	config
	ctx        *QueryContext
	order      []sysloginlog.OrderOption
	inters     []Interceptor
	predicates []predicate.SysLoginLog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysLoginLogQuery builder.
func (sllq *SysLoginLogQuery) Where(ps ...predicate.SysLoginLog) *SysLoginLogQuery {
	sllq.predicates = append(sllq.predicates, ps...)
	return sllq
}

// Limit the number of records to be returned by this query.
func (sllq *SysLoginLogQuery) Limit(limit int) *SysLoginLogQuery {
	sllq.ctx.Limit = &limit
	return sllq
}

// Offset to start from.
func (sllq *SysLoginLogQuery) Offset(offset int) *SysLoginLogQuery {
	sllq.ctx.Offset = &offset
	return sllq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sllq *SysLoginLogQuery) Unique(unique bool) *SysLoginLogQuery {
	sllq.ctx.Unique = &unique
	return sllq
}

// Order specifies how the records should be ordered.
func (sllq *SysLoginLogQuery) Order(o ...sysloginlog.OrderOption) *SysLoginLogQuery {
	sllq.order = append(sllq.order, o...)
	return sllq
}

// First returns the first SysLoginLog entity from the query.
// Returns a *NotFoundError when no SysLoginLog was found.
func (sllq *SysLoginLogQuery) First(ctx context.Context) (*SysLoginLog, error) {
	nodes, err := sllq.Limit(1).All(setContextOp(ctx, sllq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysloginlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sllq *SysLoginLogQuery) FirstX(ctx context.Context) *SysLoginLog {
	node, err := sllq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysLoginLog ID from the query.
// Returns a *NotFoundError when no SysLoginLog ID was found.
func (sllq *SysLoginLogQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sllq.Limit(1).IDs(setContextOp(ctx, sllq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysloginlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sllq *SysLoginLogQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sllq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysLoginLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysLoginLog entity is found.
// Returns a *NotFoundError when no SysLoginLog entities are found.
func (sllq *SysLoginLogQuery) Only(ctx context.Context) (*SysLoginLog, error) {
	nodes, err := sllq.Limit(2).All(setContextOp(ctx, sllq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysloginlog.Label}
	default:
		return nil, &NotSingularError{sysloginlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sllq *SysLoginLogQuery) OnlyX(ctx context.Context) *SysLoginLog {
	node, err := sllq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysLoginLog ID in the query.
// Returns a *NotSingularError when more than one SysLoginLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (sllq *SysLoginLogQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sllq.Limit(2).IDs(setContextOp(ctx, sllq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysloginlog.Label}
	default:
		err = &NotSingularError{sysloginlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sllq *SysLoginLogQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sllq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysLoginLogs.
func (sllq *SysLoginLogQuery) All(ctx context.Context) ([]*SysLoginLog, error) {
	ctx = setContextOp(ctx, sllq.ctx, "All")
	if err := sllq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysLoginLog, *SysLoginLogQuery]()
	return withInterceptors[[]*SysLoginLog](ctx, sllq, qr, sllq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sllq *SysLoginLogQuery) AllX(ctx context.Context) []*SysLoginLog {
	nodes, err := sllq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysLoginLog IDs.
func (sllq *SysLoginLogQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if sllq.ctx.Unique == nil && sllq.path != nil {
		sllq.Unique(true)
	}
	ctx = setContextOp(ctx, sllq.ctx, "IDs")
	if err = sllq.Select(sysloginlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sllq *SysLoginLogQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sllq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sllq *SysLoginLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sllq.ctx, "Count")
	if err := sllq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sllq, querierCount[*SysLoginLogQuery](), sllq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sllq *SysLoginLogQuery) CountX(ctx context.Context) int {
	count, err := sllq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sllq *SysLoginLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sllq.ctx, "Exist")
	switch _, err := sllq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("codegen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sllq *SysLoginLogQuery) ExistX(ctx context.Context) bool {
	exist, err := sllq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysLoginLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sllq *SysLoginLogQuery) Clone() *SysLoginLogQuery {
	if sllq == nil {
		return nil
	}
	return &SysLoginLogQuery{
		config:     sllq.config,
		ctx:        sllq.ctx.Clone(),
		order:      append([]sysloginlog.OrderOption{}, sllq.order...),
		inters:     append([]Interceptor{}, sllq.inters...),
		predicates: append([]predicate.SysLoginLog{}, sllq.predicates...),
		// clone intermediate query.
		sql:  sllq.sql.Clone(),
		path: sllq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status int8 `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysLoginLog.Query().
//		GroupBy(sysloginlog.FieldStatus).
//		Aggregate(codegen.Count()).
//		Scan(ctx, &v)
func (sllq *SysLoginLogQuery) GroupBy(field string, fields ...string) *SysLoginLogGroupBy {
	sllq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysLoginLogGroupBy{build: sllq}
	grbuild.flds = &sllq.ctx.Fields
	grbuild.label = sysloginlog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status int8 `json:"status,omitempty"`
//	}
//
//	client.SysLoginLog.Query().
//		Select(sysloginlog.FieldStatus).
//		Scan(ctx, &v)
func (sllq *SysLoginLogQuery) Select(fields ...string) *SysLoginLogSelect {
	sllq.ctx.Fields = append(sllq.ctx.Fields, fields...)
	sbuild := &SysLoginLogSelect{SysLoginLogQuery: sllq}
	sbuild.label = sysloginlog.Label
	sbuild.flds, sbuild.scan = &sllq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysLoginLogSelect configured with the given aggregations.
func (sllq *SysLoginLogQuery) Aggregate(fns ...AggregateFunc) *SysLoginLogSelect {
	return sllq.Select().Aggregate(fns...)
}

func (sllq *SysLoginLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sllq.inters {
		if inter == nil {
			return fmt.Errorf("codegen: uninitialized interceptor (forgotten import codegen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sllq); err != nil {
				return err
			}
		}
	}
	for _, f := range sllq.ctx.Fields {
		if !sysloginlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
		}
	}
	if sllq.path != nil {
		prev, err := sllq.path(ctx)
		if err != nil {
			return err
		}
		sllq.sql = prev
	}
	return nil
}

func (sllq *SysLoginLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysLoginLog, error) {
	var (
		nodes = []*SysLoginLog{}
		_spec = sllq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysLoginLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysLoginLog{config: sllq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sllq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sllq *SysLoginLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sllq.querySpec()
	_spec.Node.Columns = sllq.ctx.Fields
	if len(sllq.ctx.Fields) > 0 {
		_spec.Unique = sllq.ctx.Unique != nil && *sllq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sllq.driver, _spec)
}

func (sllq *SysLoginLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysloginlog.Table, sysloginlog.Columns, sqlgraph.NewFieldSpec(sysloginlog.FieldID, field.TypeInt64))
	_spec.From = sllq.sql
	if unique := sllq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sllq.path != nil {
		_spec.Unique = true
	}
	if fields := sllq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysloginlog.FieldID)
		for i := range fields {
			if fields[i] != sysloginlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sllq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sllq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sllq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sllq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sllq *SysLoginLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sllq.driver.Dialect())
	t1 := builder.Table(sysloginlog.Table)
	columns := sllq.ctx.Fields
	if len(columns) == 0 {
		columns = sysloginlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sllq.sql != nil {
		selector = sllq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sllq.ctx.Unique != nil && *sllq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sllq.predicates {
		p(selector)
	}
	for _, p := range sllq.order {
		p(selector)
	}
	if offset := sllq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sllq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysLoginLogGroupBy is the group-by builder for SysLoginLog entities.
type SysLoginLogGroupBy struct {
	selector
	build *SysLoginLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sllgb *SysLoginLogGroupBy) Aggregate(fns ...AggregateFunc) *SysLoginLogGroupBy {
	sllgb.fns = append(sllgb.fns, fns...)
	return sllgb
}

// Scan applies the selector query and scans the result into the given value.
func (sllgb *SysLoginLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sllgb.build.ctx, "GroupBy")
	if err := sllgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysLoginLogQuery, *SysLoginLogGroupBy](ctx, sllgb.build, sllgb, sllgb.build.inters, v)
}

func (sllgb *SysLoginLogGroupBy) sqlScan(ctx context.Context, root *SysLoginLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sllgb.fns))
	for _, fn := range sllgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sllgb.flds)+len(sllgb.fns))
		for _, f := range *sllgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sllgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sllgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysLoginLogSelect is the builder for selecting fields of SysLoginLog entities.
type SysLoginLogSelect struct {
	*SysLoginLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (slls *SysLoginLogSelect) Aggregate(fns ...AggregateFunc) *SysLoginLogSelect {
	slls.fns = append(slls.fns, fns...)
	return slls
}

// Scan applies the selector query and scans the result into the given value.
func (slls *SysLoginLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, slls.ctx, "Select")
	if err := slls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysLoginLogQuery, *SysLoginLogSelect](ctx, slls.SysLoginLogQuery, slls, slls.inters, v)
}

func (slls *SysLoginLogSelect) sqlScan(ctx context.Context, root *SysLoginLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(slls.fns))
	for _, fn := range slls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*slls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := slls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}