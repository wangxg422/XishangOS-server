// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/appinstance"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/apppackage"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/predicate"
)

// AppInstanceQuery is the builder for querying AppInstance entities.
type AppInstanceQuery struct {
	config
	ctx             *QueryContext
	order           []appinstance.OrderOption
	inters          []Interceptor
	predicates      []predicate.AppInstance
	withInstallFrom *AppPackageQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppInstanceQuery builder.
func (aiq *AppInstanceQuery) Where(ps ...predicate.AppInstance) *AppInstanceQuery {
	aiq.predicates = append(aiq.predicates, ps...)
	return aiq
}

// Limit the number of records to be returned by this query.
func (aiq *AppInstanceQuery) Limit(limit int) *AppInstanceQuery {
	aiq.ctx.Limit = &limit
	return aiq
}

// Offset to start from.
func (aiq *AppInstanceQuery) Offset(offset int) *AppInstanceQuery {
	aiq.ctx.Offset = &offset
	return aiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aiq *AppInstanceQuery) Unique(unique bool) *AppInstanceQuery {
	aiq.ctx.Unique = &unique
	return aiq
}

// Order specifies how the records should be ordered.
func (aiq *AppInstanceQuery) Order(o ...appinstance.OrderOption) *AppInstanceQuery {
	aiq.order = append(aiq.order, o...)
	return aiq
}

// QueryInstallFrom chains the current query on the "installFrom" edge.
func (aiq *AppInstanceQuery) QueryInstallFrom() *AppPackageQuery {
	query := (&AppPackageClient{config: aiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appinstance.Table, appinstance.FieldID, selector),
			sqlgraph.To(apppackage.Table, apppackage.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, appinstance.InstallFromTable, appinstance.InstallFromColumn),
		)
		fromU = sqlgraph.SetNeighbors(aiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AppInstance entity from the query.
// Returns a *NotFoundError when no AppInstance was found.
func (aiq *AppInstanceQuery) First(ctx context.Context) (*AppInstance, error) {
	nodes, err := aiq.Limit(1).All(setContextOp(ctx, aiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appinstance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aiq *AppInstanceQuery) FirstX(ctx context.Context) *AppInstance {
	node, err := aiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppInstance ID from the query.
// Returns a *NotFoundError when no AppInstance ID was found.
func (aiq *AppInstanceQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aiq.Limit(1).IDs(setContextOp(ctx, aiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appinstance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aiq *AppInstanceQuery) FirstIDX(ctx context.Context) int64 {
	id, err := aiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppInstance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppInstance entity is found.
// Returns a *NotFoundError when no AppInstance entities are found.
func (aiq *AppInstanceQuery) Only(ctx context.Context) (*AppInstance, error) {
	nodes, err := aiq.Limit(2).All(setContextOp(ctx, aiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appinstance.Label}
	default:
		return nil, &NotSingularError{appinstance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aiq *AppInstanceQuery) OnlyX(ctx context.Context) *AppInstance {
	node, err := aiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppInstance ID in the query.
// Returns a *NotSingularError when more than one AppInstance ID is found.
// Returns a *NotFoundError when no entities are found.
func (aiq *AppInstanceQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aiq.Limit(2).IDs(setContextOp(ctx, aiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appinstance.Label}
	default:
		err = &NotSingularError{appinstance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aiq *AppInstanceQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := aiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppInstances.
func (aiq *AppInstanceQuery) All(ctx context.Context) ([]*AppInstance, error) {
	ctx = setContextOp(ctx, aiq.ctx, "All")
	if err := aiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppInstance, *AppInstanceQuery]()
	return withInterceptors[[]*AppInstance](ctx, aiq, qr, aiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aiq *AppInstanceQuery) AllX(ctx context.Context) []*AppInstance {
	nodes, err := aiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppInstance IDs.
func (aiq *AppInstanceQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if aiq.ctx.Unique == nil && aiq.path != nil {
		aiq.Unique(true)
	}
	ctx = setContextOp(ctx, aiq.ctx, "IDs")
	if err = aiq.Select(appinstance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aiq *AppInstanceQuery) IDsX(ctx context.Context) []int64 {
	ids, err := aiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aiq *AppInstanceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aiq.ctx, "Count")
	if err := aiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aiq, querierCount[*AppInstanceQuery](), aiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aiq *AppInstanceQuery) CountX(ctx context.Context) int {
	count, err := aiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aiq *AppInstanceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aiq.ctx, "Exist")
	switch _, err := aiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("codegen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aiq *AppInstanceQuery) ExistX(ctx context.Context) bool {
	exist, err := aiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppInstanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aiq *AppInstanceQuery) Clone() *AppInstanceQuery {
	if aiq == nil {
		return nil
	}
	return &AppInstanceQuery{
		config:          aiq.config,
		ctx:             aiq.ctx.Clone(),
		order:           append([]appinstance.OrderOption{}, aiq.order...),
		inters:          append([]Interceptor{}, aiq.inters...),
		predicates:      append([]predicate.AppInstance{}, aiq.predicates...),
		withInstallFrom: aiq.withInstallFrom.Clone(),
		// clone intermediate query.
		sql:  aiq.sql.Clone(),
		path: aiq.path,
	}
}

// WithInstallFrom tells the query-builder to eager-load the nodes that are connected to
// the "installFrom" edge. The optional arguments are used to configure the query builder of the edge.
func (aiq *AppInstanceQuery) WithInstallFrom(opts ...func(*AppPackageQuery)) *AppInstanceQuery {
	query := (&AppPackageClient{config: aiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aiq.withInstallFrom = query
	return aiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createAt"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppInstance.Query().
//		GroupBy(appinstance.FieldCreatedAt).
//		Aggregate(codegen.Count()).
//		Scan(ctx, &v)
func (aiq *AppInstanceQuery) GroupBy(field string, fields ...string) *AppInstanceGroupBy {
	aiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppInstanceGroupBy{build: aiq}
	grbuild.flds = &aiq.ctx.Fields
	grbuild.label = appinstance.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createAt"`
//	}
//
//	client.AppInstance.Query().
//		Select(appinstance.FieldCreatedAt).
//		Scan(ctx, &v)
func (aiq *AppInstanceQuery) Select(fields ...string) *AppInstanceSelect {
	aiq.ctx.Fields = append(aiq.ctx.Fields, fields...)
	sbuild := &AppInstanceSelect{AppInstanceQuery: aiq}
	sbuild.label = appinstance.Label
	sbuild.flds, sbuild.scan = &aiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppInstanceSelect configured with the given aggregations.
func (aiq *AppInstanceQuery) Aggregate(fns ...AggregateFunc) *AppInstanceSelect {
	return aiq.Select().Aggregate(fns...)
}

func (aiq *AppInstanceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aiq.inters {
		if inter == nil {
			return fmt.Errorf("codegen: uninitialized interceptor (forgotten import codegen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aiq); err != nil {
				return err
			}
		}
	}
	for _, f := range aiq.ctx.Fields {
		if !appinstance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
		}
	}
	if aiq.path != nil {
		prev, err := aiq.path(ctx)
		if err != nil {
			return err
		}
		aiq.sql = prev
	}
	return nil
}

func (aiq *AppInstanceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppInstance, error) {
	var (
		nodes       = []*AppInstance{}
		withFKs     = aiq.withFKs
		_spec       = aiq.querySpec()
		loadedTypes = [1]bool{
			aiq.withInstallFrom != nil,
		}
	)
	if aiq.withInstallFrom != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, appinstance.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppInstance).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppInstance{config: aiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aiq.withInstallFrom; query != nil {
		if err := aiq.loadInstallFrom(ctx, query, nodes, nil,
			func(n *AppInstance, e *AppPackage) { n.Edges.InstallFrom = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aiq *AppInstanceQuery) loadInstallFrom(ctx context.Context, query *AppPackageQuery, nodes []*AppInstance, init func(*AppInstance), assign func(*AppInstance, *AppPackage)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*AppInstance)
	for i := range nodes {
		if nodes[i].app_package_app_instance == nil {
			continue
		}
		fk := *nodes[i].app_package_app_instance
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(apppackage.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "app_package_app_instance" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aiq *AppInstanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aiq.querySpec()
	_spec.Node.Columns = aiq.ctx.Fields
	if len(aiq.ctx.Fields) > 0 {
		_spec.Unique = aiq.ctx.Unique != nil && *aiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aiq.driver, _spec)
}

func (aiq *AppInstanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appinstance.Table, appinstance.Columns, sqlgraph.NewFieldSpec(appinstance.FieldID, field.TypeInt64))
	_spec.From = aiq.sql
	if unique := aiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aiq.path != nil {
		_spec.Unique = true
	}
	if fields := aiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appinstance.FieldID)
		for i := range fields {
			if fields[i] != appinstance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aiq *AppInstanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aiq.driver.Dialect())
	t1 := builder.Table(appinstance.Table)
	columns := aiq.ctx.Fields
	if len(columns) == 0 {
		columns = appinstance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aiq.sql != nil {
		selector = aiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aiq.ctx.Unique != nil && *aiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aiq.predicates {
		p(selector)
	}
	for _, p := range aiq.order {
		p(selector)
	}
	if offset := aiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppInstanceGroupBy is the group-by builder for AppInstance entities.
type AppInstanceGroupBy struct {
	selector
	build *AppInstanceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aigb *AppInstanceGroupBy) Aggregate(fns ...AggregateFunc) *AppInstanceGroupBy {
	aigb.fns = append(aigb.fns, fns...)
	return aigb
}

// Scan applies the selector query and scans the result into the given value.
func (aigb *AppInstanceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aigb.build.ctx, "GroupBy")
	if err := aigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppInstanceQuery, *AppInstanceGroupBy](ctx, aigb.build, aigb, aigb.build.inters, v)
}

func (aigb *AppInstanceGroupBy) sqlScan(ctx context.Context, root *AppInstanceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aigb.fns))
	for _, fn := range aigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aigb.flds)+len(aigb.fns))
		for _, f := range *aigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppInstanceSelect is the builder for selecting fields of AppInstance entities.
type AppInstanceSelect struct {
	*AppInstanceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ais *AppInstanceSelect) Aggregate(fns ...AggregateFunc) *AppInstanceSelect {
	ais.fns = append(ais.fns, fns...)
	return ais
}

// Scan applies the selector query and scans the result into the given value.
func (ais *AppInstanceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ais.ctx, "Select")
	if err := ais.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppInstanceQuery, *AppInstanceSelect](ctx, ais.AppInstanceQuery, ais, ais.inters, v)
}

func (ais *AppInstanceSelect) sqlScan(ctx context.Context, root *AppInstanceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ais.fns))
	for _, fn := range ais.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ais.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ais.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
