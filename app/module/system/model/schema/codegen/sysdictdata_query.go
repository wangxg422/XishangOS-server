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
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdictdata"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdicttype"
)

// SysDictDataQuery is the builder for querying SysDictData entities.
type SysDictDataQuery struct {
	config
	ctx             *QueryContext
	order           []sysdictdata.OrderOption
	inters          []Interceptor
	predicates      []predicate.SysDictData
	withSysDictType *SysDictTypeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysDictDataQuery builder.
func (sddq *SysDictDataQuery) Where(ps ...predicate.SysDictData) *SysDictDataQuery {
	sddq.predicates = append(sddq.predicates, ps...)
	return sddq
}

// Limit the number of records to be returned by this query.
func (sddq *SysDictDataQuery) Limit(limit int) *SysDictDataQuery {
	sddq.ctx.Limit = &limit
	return sddq
}

// Offset to start from.
func (sddq *SysDictDataQuery) Offset(offset int) *SysDictDataQuery {
	sddq.ctx.Offset = &offset
	return sddq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sddq *SysDictDataQuery) Unique(unique bool) *SysDictDataQuery {
	sddq.ctx.Unique = &unique
	return sddq
}

// Order specifies how the records should be ordered.
func (sddq *SysDictDataQuery) Order(o ...sysdictdata.OrderOption) *SysDictDataQuery {
	sddq.order = append(sddq.order, o...)
	return sddq
}

// QuerySysDictType chains the current query on the "sysDictType" edge.
func (sddq *SysDictDataQuery) QuerySysDictType() *SysDictTypeQuery {
	query := (&SysDictTypeClient{config: sddq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sddq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sddq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysdictdata.Table, sysdictdata.FieldID, selector),
			sqlgraph.To(sysdicttype.Table, sysdicttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sysdictdata.SysDictTypeTable, sysdictdata.SysDictTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(sddq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysDictData entity from the query.
// Returns a *NotFoundError when no SysDictData was found.
func (sddq *SysDictDataQuery) First(ctx context.Context) (*SysDictData, error) {
	nodes, err := sddq.Limit(1).All(setContextOp(ctx, sddq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysdictdata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sddq *SysDictDataQuery) FirstX(ctx context.Context) *SysDictData {
	node, err := sddq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysDictData ID from the query.
// Returns a *NotFoundError when no SysDictData ID was found.
func (sddq *SysDictDataQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sddq.Limit(1).IDs(setContextOp(ctx, sddq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysdictdata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sddq *SysDictDataQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sddq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysDictData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysDictData entity is found.
// Returns a *NotFoundError when no SysDictData entities are found.
func (sddq *SysDictDataQuery) Only(ctx context.Context) (*SysDictData, error) {
	nodes, err := sddq.Limit(2).All(setContextOp(ctx, sddq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysdictdata.Label}
	default:
		return nil, &NotSingularError{sysdictdata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sddq *SysDictDataQuery) OnlyX(ctx context.Context) *SysDictData {
	node, err := sddq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysDictData ID in the query.
// Returns a *NotSingularError when more than one SysDictData ID is found.
// Returns a *NotFoundError when no entities are found.
func (sddq *SysDictDataQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sddq.Limit(2).IDs(setContextOp(ctx, sddq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysdictdata.Label}
	default:
		err = &NotSingularError{sysdictdata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sddq *SysDictDataQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sddq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysDictDataSlice.
func (sddq *SysDictDataQuery) All(ctx context.Context) ([]*SysDictData, error) {
	ctx = setContextOp(ctx, sddq.ctx, "All")
	if err := sddq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysDictData, *SysDictDataQuery]()
	return withInterceptors[[]*SysDictData](ctx, sddq, qr, sddq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sddq *SysDictDataQuery) AllX(ctx context.Context) []*SysDictData {
	nodes, err := sddq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysDictData IDs.
func (sddq *SysDictDataQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if sddq.ctx.Unique == nil && sddq.path != nil {
		sddq.Unique(true)
	}
	ctx = setContextOp(ctx, sddq.ctx, "IDs")
	if err = sddq.Select(sysdictdata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sddq *SysDictDataQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sddq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sddq *SysDictDataQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sddq.ctx, "Count")
	if err := sddq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sddq, querierCount[*SysDictDataQuery](), sddq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sddq *SysDictDataQuery) CountX(ctx context.Context) int {
	count, err := sddq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sddq *SysDictDataQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sddq.ctx, "Exist")
	switch _, err := sddq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("codegen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sddq *SysDictDataQuery) ExistX(ctx context.Context) bool {
	exist, err := sddq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysDictDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sddq *SysDictDataQuery) Clone() *SysDictDataQuery {
	if sddq == nil {
		return nil
	}
	return &SysDictDataQuery{
		config:          sddq.config,
		ctx:             sddq.ctx.Clone(),
		order:           append([]sysdictdata.OrderOption{}, sddq.order...),
		inters:          append([]Interceptor{}, sddq.inters...),
		predicates:      append([]predicate.SysDictData{}, sddq.predicates...),
		withSysDictType: sddq.withSysDictType.Clone(),
		// clone intermediate query.
		sql:  sddq.sql.Clone(),
		path: sddq.path,
	}
}

// WithSysDictType tells the query-builder to eager-load the nodes that are connected to
// the "sysDictType" edge. The optional arguments are used to configure the query builder of the edge.
func (sddq *SysDictDataQuery) WithSysDictType(opts ...func(*SysDictTypeQuery)) *SysDictDataQuery {
	query := (&SysDictTypeClient{config: sddq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sddq.withSysDictType = query
	return sddq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysDictData.Query().
//		GroupBy(sysdictdata.FieldCreatedAt).
//		Aggregate(codegen.Count()).
//		Scan(ctx, &v)
func (sddq *SysDictDataQuery) GroupBy(field string, fields ...string) *SysDictDataGroupBy {
	sddq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysDictDataGroupBy{build: sddq}
	grbuild.flds = &sddq.ctx.Fields
	grbuild.label = sysdictdata.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.SysDictData.Query().
//		Select(sysdictdata.FieldCreatedAt).
//		Scan(ctx, &v)
func (sddq *SysDictDataQuery) Select(fields ...string) *SysDictDataSelect {
	sddq.ctx.Fields = append(sddq.ctx.Fields, fields...)
	sbuild := &SysDictDataSelect{SysDictDataQuery: sddq}
	sbuild.label = sysdictdata.Label
	sbuild.flds, sbuild.scan = &sddq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysDictDataSelect configured with the given aggregations.
func (sddq *SysDictDataQuery) Aggregate(fns ...AggregateFunc) *SysDictDataSelect {
	return sddq.Select().Aggregate(fns...)
}

func (sddq *SysDictDataQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sddq.inters {
		if inter == nil {
			return fmt.Errorf("codegen: uninitialized interceptor (forgotten import codegen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sddq); err != nil {
				return err
			}
		}
	}
	for _, f := range sddq.ctx.Fields {
		if !sysdictdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
		}
	}
	if sddq.path != nil {
		prev, err := sddq.path(ctx)
		if err != nil {
			return err
		}
		sddq.sql = prev
	}
	return nil
}

func (sddq *SysDictDataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysDictData, error) {
	var (
		nodes       = []*SysDictData{}
		_spec       = sddq.querySpec()
		loadedTypes = [1]bool{
			sddq.withSysDictType != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysDictData).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysDictData{config: sddq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sddq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sddq.withSysDictType; query != nil {
		if err := sddq.loadSysDictType(ctx, query, nodes, nil,
			func(n *SysDictData, e *SysDictType) { n.Edges.SysDictType = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sddq *SysDictDataQuery) loadSysDictType(ctx context.Context, query *SysDictTypeQuery, nodes []*SysDictData, init func(*SysDictData), assign func(*SysDictData, *SysDictType)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*SysDictData)
	for i := range nodes {
		fk := nodes[i].DictTypeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(sysdicttype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dict_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sddq *SysDictDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sddq.querySpec()
	_spec.Node.Columns = sddq.ctx.Fields
	if len(sddq.ctx.Fields) > 0 {
		_spec.Unique = sddq.ctx.Unique != nil && *sddq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sddq.driver, _spec)
}

func (sddq *SysDictDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysdictdata.Table, sysdictdata.Columns, sqlgraph.NewFieldSpec(sysdictdata.FieldID, field.TypeInt64))
	_spec.From = sddq.sql
	if unique := sddq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sddq.path != nil {
		_spec.Unique = true
	}
	if fields := sddq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdictdata.FieldID)
		for i := range fields {
			if fields[i] != sysdictdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if sddq.withSysDictType != nil {
			_spec.Node.AddColumnOnce(sysdictdata.FieldDictTypeID)
		}
	}
	if ps := sddq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sddq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sddq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sddq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sddq *SysDictDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sddq.driver.Dialect())
	t1 := builder.Table(sysdictdata.Table)
	columns := sddq.ctx.Fields
	if len(columns) == 0 {
		columns = sysdictdata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sddq.sql != nil {
		selector = sddq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sddq.ctx.Unique != nil && *sddq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sddq.predicates {
		p(selector)
	}
	for _, p := range sddq.order {
		p(selector)
	}
	if offset := sddq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sddq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysDictDataGroupBy is the group-by builder for SysDictData entities.
type SysDictDataGroupBy struct {
	selector
	build *SysDictDataQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sddgb *SysDictDataGroupBy) Aggregate(fns ...AggregateFunc) *SysDictDataGroupBy {
	sddgb.fns = append(sddgb.fns, fns...)
	return sddgb
}

// Scan applies the selector query and scans the result into the given value.
func (sddgb *SysDictDataGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sddgb.build.ctx, "GroupBy")
	if err := sddgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysDictDataQuery, *SysDictDataGroupBy](ctx, sddgb.build, sddgb, sddgb.build.inters, v)
}

func (sddgb *SysDictDataGroupBy) sqlScan(ctx context.Context, root *SysDictDataQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sddgb.fns))
	for _, fn := range sddgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sddgb.flds)+len(sddgb.fns))
		for _, f := range *sddgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sddgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sddgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysDictDataSelect is the builder for selecting fields of SysDictData entities.
type SysDictDataSelect struct {
	*SysDictDataQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sdds *SysDictDataSelect) Aggregate(fns ...AggregateFunc) *SysDictDataSelect {
	sdds.fns = append(sdds.fns, fns...)
	return sdds
}

// Scan applies the selector query and scans the result into the given value.
func (sdds *SysDictDataSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sdds.ctx, "Select")
	if err := sdds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysDictDataQuery, *SysDictDataSelect](ctx, sdds.SysDictDataQuery, sdds, sdds.inters, v)
}

func (sdds *SysDictDataSelect) sqlScan(ctx context.Context, root *SysDictDataQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sdds.fns))
	for _, fn := range sdds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sdds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sdds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
