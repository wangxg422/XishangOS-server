// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/predicate"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/syspost"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuser"
)

// SysPostQuery is the builder for querying SysPost entities.
type SysPostQuery struct {
	config
	ctx          *QueryContext
	order        []syspost.OrderOption
	inters       []Interceptor
	predicates   []predicate.SysPost
	withSysUsers *SysUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysPostQuery builder.
func (spq *SysPostQuery) Where(ps ...predicate.SysPost) *SysPostQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit the number of records to be returned by this query.
func (spq *SysPostQuery) Limit(limit int) *SysPostQuery {
	spq.ctx.Limit = &limit
	return spq
}

// Offset to start from.
func (spq *SysPostQuery) Offset(offset int) *SysPostQuery {
	spq.ctx.Offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *SysPostQuery) Unique(unique bool) *SysPostQuery {
	spq.ctx.Unique = &unique
	return spq
}

// Order specifies how the records should be ordered.
func (spq *SysPostQuery) Order(o ...syspost.OrderOption) *SysPostQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// QuerySysUsers chains the current query on the "sysUsers" edge.
func (spq *SysPostQuery) QuerySysUsers() *SysUserQuery {
	query := (&SysUserClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(syspost.Table, syspost.FieldID, selector),
			sqlgraph.To(sysuser.Table, sysuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, syspost.SysUsersTable, syspost.SysUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysPost entity from the query.
// Returns a *NotFoundError when no SysPost was found.
func (spq *SysPostQuery) First(ctx context.Context) (*SysPost, error) {
	nodes, err := spq.Limit(1).All(setContextOp(ctx, spq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{syspost.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *SysPostQuery) FirstX(ctx context.Context) *SysPost {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysPost ID from the query.
// Returns a *NotFoundError when no SysPost ID was found.
func (spq *SysPostQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = spq.Limit(1).IDs(setContextOp(ctx, spq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{syspost.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *SysPostQuery) FirstIDX(ctx context.Context) int64 {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysPost entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysPost entity is found.
// Returns a *NotFoundError when no SysPost entities are found.
func (spq *SysPostQuery) Only(ctx context.Context) (*SysPost, error) {
	nodes, err := spq.Limit(2).All(setContextOp(ctx, spq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{syspost.Label}
	default:
		return nil, &NotSingularError{syspost.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *SysPostQuery) OnlyX(ctx context.Context) *SysPost {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysPost ID in the query.
// Returns a *NotSingularError when more than one SysPost ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *SysPostQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = spq.Limit(2).IDs(setContextOp(ctx, spq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{syspost.Label}
	default:
		err = &NotSingularError{syspost.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *SysPostQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysPosts.
func (spq *SysPostQuery) All(ctx context.Context) ([]*SysPost, error) {
	ctx = setContextOp(ctx, spq.ctx, "All")
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysPost, *SysPostQuery]()
	return withInterceptors[[]*SysPost](ctx, spq, qr, spq.inters)
}

// AllX is like All, but panics if an error occurs.
func (spq *SysPostQuery) AllX(ctx context.Context) []*SysPost {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysPost IDs.
func (spq *SysPostQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if spq.ctx.Unique == nil && spq.path != nil {
		spq.Unique(true)
	}
	ctx = setContextOp(ctx, spq.ctx, "IDs")
	if err = spq.Select(syspost.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *SysPostQuery) IDsX(ctx context.Context) []int64 {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *SysPostQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, spq.ctx, "Count")
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, spq, querierCount[*SysPostQuery](), spq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (spq *SysPostQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *SysPostQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, spq.ctx, "Exist")
	switch _, err := spq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("codegen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *SysPostQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysPostQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *SysPostQuery) Clone() *SysPostQuery {
	if spq == nil {
		return nil
	}
	return &SysPostQuery{
		config:       spq.config,
		ctx:          spq.ctx.Clone(),
		order:        append([]syspost.OrderOption{}, spq.order...),
		inters:       append([]Interceptor{}, spq.inters...),
		predicates:   append([]predicate.SysPost{}, spq.predicates...),
		withSysUsers: spq.withSysUsers.Clone(),
		// clone intermediate query.
		sql:  spq.sql.Clone(),
		path: spq.path,
	}
}

// WithSysUsers tells the query-builder to eager-load the nodes that are connected to
// the "sysUsers" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *SysPostQuery) WithSysUsers(opts ...func(*SysUserQuery)) *SysPostQuery {
	query := (&SysUserClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withSysUsers = query
	return spq
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
//	client.SysPost.Query().
//		GroupBy(syspost.FieldCreatedAt).
//		Aggregate(codegen.Count()).
//		Scan(ctx, &v)
func (spq *SysPostQuery) GroupBy(field string, fields ...string) *SysPostGroupBy {
	spq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysPostGroupBy{build: spq}
	grbuild.flds = &spq.ctx.Fields
	grbuild.label = syspost.Label
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
//	client.SysPost.Query().
//		Select(syspost.FieldCreatedAt).
//		Scan(ctx, &v)
func (spq *SysPostQuery) Select(fields ...string) *SysPostSelect {
	spq.ctx.Fields = append(spq.ctx.Fields, fields...)
	sbuild := &SysPostSelect{SysPostQuery: spq}
	sbuild.label = syspost.Label
	sbuild.flds, sbuild.scan = &spq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysPostSelect configured with the given aggregations.
func (spq *SysPostQuery) Aggregate(fns ...AggregateFunc) *SysPostSelect {
	return spq.Select().Aggregate(fns...)
}

func (spq *SysPostQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range spq.inters {
		if inter == nil {
			return fmt.Errorf("codegen: uninitialized interceptor (forgotten import codegen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, spq); err != nil {
				return err
			}
		}
	}
	for _, f := range spq.ctx.Fields {
		if !syspost.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *SysPostQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysPost, error) {
	var (
		nodes       = []*SysPost{}
		_spec       = spq.querySpec()
		loadedTypes = [1]bool{
			spq.withSysUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysPost).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysPost{config: spq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := spq.withSysUsers; query != nil {
		if err := spq.loadSysUsers(ctx, query, nodes,
			func(n *SysPost) { n.Edges.SysUsers = []*SysUser{} },
			func(n *SysPost, e *SysUser) { n.Edges.SysUsers = append(n.Edges.SysUsers, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (spq *SysPostQuery) loadSysUsers(ctx context.Context, query *SysUserQuery, nodes []*SysPost, init func(*SysPost), assign func(*SysPost, *SysUser)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*SysPost)
	nids := make(map[int64]map[*SysPost]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(syspost.SysUsersTable)
		s.Join(joinT).On(s.C(sysuser.FieldID), joinT.C(syspost.SysUsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(syspost.SysUsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(syspost.SysUsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullInt64).Int64
				inValue := values[1].(*sql.NullInt64).Int64
				if nids[inValue] == nil {
					nids[inValue] = map[*SysPost]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*SysUser](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "sysUsers" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (spq *SysPostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.ctx.Fields
	if len(spq.ctx.Fields) > 0 {
		_spec.Unique = spq.ctx.Unique != nil && *spq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *SysPostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(syspost.Table, syspost.Columns, sqlgraph.NewFieldSpec(syspost.FieldID, field.TypeInt64))
	_spec.From = spq.sql
	if unique := spq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if spq.path != nil {
		_spec.Unique = true
	}
	if fields := spq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, syspost.FieldID)
		for i := range fields {
			if fields[i] != syspost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *SysPostQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(syspost.Table)
	columns := spq.ctx.Fields
	if len(columns) == 0 {
		columns = syspost.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.ctx.Unique != nil && *spq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysPostGroupBy is the group-by builder for SysPost entities.
type SysPostGroupBy struct {
	selector
	build *SysPostQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *SysPostGroupBy) Aggregate(fns ...AggregateFunc) *SysPostGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the selector query and scans the result into the given value.
func (spgb *SysPostGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, spgb.build.ctx, "GroupBy")
	if err := spgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysPostQuery, *SysPostGroupBy](ctx, spgb.build, spgb, spgb.build.inters, v)
}

func (spgb *SysPostGroupBy) sqlScan(ctx context.Context, root *SysPostQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*spgb.flds)+len(spgb.fns))
		for _, f := range *spgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*spgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysPostSelect is the builder for selecting fields of SysPost entities.
type SysPostSelect struct {
	*SysPostQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sps *SysPostSelect) Aggregate(fns ...AggregateFunc) *SysPostSelect {
	sps.fns = append(sps.fns, fns...)
	return sps
}

// Scan applies the selector query and scans the result into the given value.
func (sps *SysPostSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sps.ctx, "Select")
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysPostQuery, *SysPostSelect](ctx, sps.SysPostQuery, sps, sps.inters, v)
}

func (sps *SysPostSelect) sqlScan(ctx context.Context, root *SysPostQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sps.fns))
	for _, fn := range sps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
