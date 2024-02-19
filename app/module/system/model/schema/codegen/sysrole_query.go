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
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdept"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysmenu"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysrole"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuser"
)

// SysRoleQuery is the builder for querying SysRole entities.
type SysRoleQuery struct {
	config
	ctx          *QueryContext
	order        []sysrole.OrderOption
	inters       []Interceptor
	predicates   []predicate.SysRole
	withDepts    *SysDeptQuery
	withSysUsers *SysUserQuery
	withSysMenus *SysMenuQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysRoleQuery builder.
func (srq *SysRoleQuery) Where(ps ...predicate.SysRole) *SysRoleQuery {
	srq.predicates = append(srq.predicates, ps...)
	return srq
}

// Limit the number of records to be returned by this query.
func (srq *SysRoleQuery) Limit(limit int) *SysRoleQuery {
	srq.ctx.Limit = &limit
	return srq
}

// Offset to start from.
func (srq *SysRoleQuery) Offset(offset int) *SysRoleQuery {
	srq.ctx.Offset = &offset
	return srq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (srq *SysRoleQuery) Unique(unique bool) *SysRoleQuery {
	srq.ctx.Unique = &unique
	return srq
}

// Order specifies how the records should be ordered.
func (srq *SysRoleQuery) Order(o ...sysrole.OrderOption) *SysRoleQuery {
	srq.order = append(srq.order, o...)
	return srq
}

// QueryDepts chains the current query on the "depts" edge.
func (srq *SysRoleQuery) QueryDepts() *SysDeptQuery {
	query := (&SysDeptClient{config: srq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := srq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysrole.Table, sysrole.FieldID, selector),
			sqlgraph.To(sysdept.Table, sysdept.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, sysrole.DeptsTable, sysrole.DeptsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySysUsers chains the current query on the "sysUsers" edge.
func (srq *SysRoleQuery) QuerySysUsers() *SysUserQuery {
	query := (&SysUserClient{config: srq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := srq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysrole.Table, sysrole.FieldID, selector),
			sqlgraph.To(sysuser.Table, sysuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, sysrole.SysUsersTable, sysrole.SysUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySysMenus chains the current query on the "sysMenus" edge.
func (srq *SysRoleQuery) QuerySysMenus() *SysMenuQuery {
	query := (&SysMenuClient{config: srq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := srq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysrole.Table, sysrole.FieldID, selector),
			sqlgraph.To(sysmenu.Table, sysmenu.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, sysrole.SysMenusTable, sysrole.SysMenusPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysRole entity from the query.
// Returns a *NotFoundError when no SysRole was found.
func (srq *SysRoleQuery) First(ctx context.Context) (*SysRole, error) {
	nodes, err := srq.Limit(1).All(setContextOp(ctx, srq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (srq *SysRoleQuery) FirstX(ctx context.Context) *SysRole {
	node, err := srq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysRole ID from the query.
// Returns a *NotFoundError when no SysRole ID was found.
func (srq *SysRoleQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = srq.Limit(1).IDs(setContextOp(ctx, srq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (srq *SysRoleQuery) FirstIDX(ctx context.Context) int64 {
	id, err := srq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SysRole entity is found.
// Returns a *NotFoundError when no SysRole entities are found.
func (srq *SysRoleQuery) Only(ctx context.Context) (*SysRole, error) {
	nodes, err := srq.Limit(2).All(setContextOp(ctx, srq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysrole.Label}
	default:
		return nil, &NotSingularError{sysrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (srq *SysRoleQuery) OnlyX(ctx context.Context) *SysRole {
	node, err := srq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysRole ID in the query.
// Returns a *NotSingularError when more than one SysRole ID is found.
// Returns a *NotFoundError when no entities are found.
func (srq *SysRoleQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = srq.Limit(2).IDs(setContextOp(ctx, srq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = &NotSingularError{sysrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (srq *SysRoleQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := srq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysRoles.
func (srq *SysRoleQuery) All(ctx context.Context) ([]*SysRole, error) {
	ctx = setContextOp(ctx, srq.ctx, "All")
	if err := srq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SysRole, *SysRoleQuery]()
	return withInterceptors[[]*SysRole](ctx, srq, qr, srq.inters)
}

// AllX is like All, but panics if an error occurs.
func (srq *SysRoleQuery) AllX(ctx context.Context) []*SysRole {
	nodes, err := srq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysRole IDs.
func (srq *SysRoleQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if srq.ctx.Unique == nil && srq.path != nil {
		srq.Unique(true)
	}
	ctx = setContextOp(ctx, srq.ctx, "IDs")
	if err = srq.Select(sysrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (srq *SysRoleQuery) IDsX(ctx context.Context) []int64 {
	ids, err := srq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (srq *SysRoleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, srq.ctx, "Count")
	if err := srq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, srq, querierCount[*SysRoleQuery](), srq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (srq *SysRoleQuery) CountX(ctx context.Context) int {
	count, err := srq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (srq *SysRoleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, srq.ctx, "Exist")
	switch _, err := srq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("codegen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (srq *SysRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := srq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (srq *SysRoleQuery) Clone() *SysRoleQuery {
	if srq == nil {
		return nil
	}
	return &SysRoleQuery{
		config:       srq.config,
		ctx:          srq.ctx.Clone(),
		order:        append([]sysrole.OrderOption{}, srq.order...),
		inters:       append([]Interceptor{}, srq.inters...),
		predicates:   append([]predicate.SysRole{}, srq.predicates...),
		withDepts:    srq.withDepts.Clone(),
		withSysUsers: srq.withSysUsers.Clone(),
		withSysMenus: srq.withSysMenus.Clone(),
		// clone intermediate query.
		sql:  srq.sql.Clone(),
		path: srq.path,
	}
}

// WithDepts tells the query-builder to eager-load the nodes that are connected to
// the "depts" edge. The optional arguments are used to configure the query builder of the edge.
func (srq *SysRoleQuery) WithDepts(opts ...func(*SysDeptQuery)) *SysRoleQuery {
	query := (&SysDeptClient{config: srq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	srq.withDepts = query
	return srq
}

// WithSysUsers tells the query-builder to eager-load the nodes that are connected to
// the "sysUsers" edge. The optional arguments are used to configure the query builder of the edge.
func (srq *SysRoleQuery) WithSysUsers(opts ...func(*SysUserQuery)) *SysRoleQuery {
	query := (&SysUserClient{config: srq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	srq.withSysUsers = query
	return srq
}

// WithSysMenus tells the query-builder to eager-load the nodes that are connected to
// the "sysMenus" edge. The optional arguments are used to configure the query builder of the edge.
func (srq *SysRoleQuery) WithSysMenus(opts ...func(*SysMenuQuery)) *SysRoleQuery {
	query := (&SysMenuClient{config: srq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	srq.withSysMenus = query
	return srq
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
//	client.SysRole.Query().
//		GroupBy(sysrole.FieldCreatedAt).
//		Aggregate(codegen.Count()).
//		Scan(ctx, &v)
func (srq *SysRoleQuery) GroupBy(field string, fields ...string) *SysRoleGroupBy {
	srq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SysRoleGroupBy{build: srq}
	grbuild.flds = &srq.ctx.Fields
	grbuild.label = sysrole.Label
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
//	client.SysRole.Query().
//		Select(sysrole.FieldCreatedAt).
//		Scan(ctx, &v)
func (srq *SysRoleQuery) Select(fields ...string) *SysRoleSelect {
	srq.ctx.Fields = append(srq.ctx.Fields, fields...)
	sbuild := &SysRoleSelect{SysRoleQuery: srq}
	sbuild.label = sysrole.Label
	sbuild.flds, sbuild.scan = &srq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SysRoleSelect configured with the given aggregations.
func (srq *SysRoleQuery) Aggregate(fns ...AggregateFunc) *SysRoleSelect {
	return srq.Select().Aggregate(fns...)
}

func (srq *SysRoleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range srq.inters {
		if inter == nil {
			return fmt.Errorf("codegen: uninitialized interceptor (forgotten import codegen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, srq); err != nil {
				return err
			}
		}
	}
	for _, f := range srq.ctx.Fields {
		if !sysrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
		}
	}
	if srq.path != nil {
		prev, err := srq.path(ctx)
		if err != nil {
			return err
		}
		srq.sql = prev
	}
	return nil
}

func (srq *SysRoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SysRole, error) {
	var (
		nodes       = []*SysRole{}
		_spec       = srq.querySpec()
		loadedTypes = [3]bool{
			srq.withDepts != nil,
			srq.withSysUsers != nil,
			srq.withSysMenus != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SysRole).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SysRole{config: srq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, srq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := srq.withDepts; query != nil {
		if err := srq.loadDepts(ctx, query, nodes,
			func(n *SysRole) { n.Edges.Depts = []*SysDept{} },
			func(n *SysRole, e *SysDept) { n.Edges.Depts = append(n.Edges.Depts, e) }); err != nil {
			return nil, err
		}
	}
	if query := srq.withSysUsers; query != nil {
		if err := srq.loadSysUsers(ctx, query, nodes,
			func(n *SysRole) { n.Edges.SysUsers = []*SysUser{} },
			func(n *SysRole, e *SysUser) { n.Edges.SysUsers = append(n.Edges.SysUsers, e) }); err != nil {
			return nil, err
		}
	}
	if query := srq.withSysMenus; query != nil {
		if err := srq.loadSysMenus(ctx, query, nodes,
			func(n *SysRole) { n.Edges.SysMenus = []*SysMenu{} },
			func(n *SysRole, e *SysMenu) { n.Edges.SysMenus = append(n.Edges.SysMenus, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (srq *SysRoleQuery) loadDepts(ctx context.Context, query *SysDeptQuery, nodes []*SysRole, init func(*SysRole), assign func(*SysRole, *SysDept)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*SysRole)
	nids := make(map[int64]map[*SysRole]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(sysrole.DeptsTable)
		s.Join(joinT).On(s.C(sysdept.FieldID), joinT.C(sysrole.DeptsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(sysrole.DeptsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(sysrole.DeptsPrimaryKey[0]))
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
					nids[inValue] = map[*SysRole]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*SysDept](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "depts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (srq *SysRoleQuery) loadSysUsers(ctx context.Context, query *SysUserQuery, nodes []*SysRole, init func(*SysRole), assign func(*SysRole, *SysUser)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*SysRole)
	nids := make(map[int64]map[*SysRole]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(sysrole.SysUsersTable)
		s.Join(joinT).On(s.C(sysuser.FieldID), joinT.C(sysrole.SysUsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(sysrole.SysUsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(sysrole.SysUsersPrimaryKey[1]))
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
					nids[inValue] = map[*SysRole]struct{}{byID[outValue]: {}}
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
func (srq *SysRoleQuery) loadSysMenus(ctx context.Context, query *SysMenuQuery, nodes []*SysRole, init func(*SysRole), assign func(*SysRole, *SysMenu)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*SysRole)
	nids := make(map[int64]map[*SysRole]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(sysrole.SysMenusTable)
		s.Join(joinT).On(s.C(sysmenu.FieldID), joinT.C(sysrole.SysMenusPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(sysrole.SysMenusPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(sysrole.SysMenusPrimaryKey[0]))
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
					nids[inValue] = map[*SysRole]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*SysMenu](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "sysMenus" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (srq *SysRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := srq.querySpec()
	_spec.Node.Columns = srq.ctx.Fields
	if len(srq.ctx.Fields) > 0 {
		_spec.Unique = srq.ctx.Unique != nil && *srq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, srq.driver, _spec)
}

func (srq *SysRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sysrole.Table, sysrole.Columns, sqlgraph.NewFieldSpec(sysrole.FieldID, field.TypeInt64))
	_spec.From = srq.sql
	if unique := srq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if srq.path != nil {
		_spec.Unique = true
	}
	if fields := srq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrole.FieldID)
		for i := range fields {
			if fields[i] != sysrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := srq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := srq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := srq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := srq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (srq *SysRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(srq.driver.Dialect())
	t1 := builder.Table(sysrole.Table)
	columns := srq.ctx.Fields
	if len(columns) == 0 {
		columns = sysrole.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if srq.sql != nil {
		selector = srq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if srq.ctx.Unique != nil && *srq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range srq.predicates {
		p(selector)
	}
	for _, p := range srq.order {
		p(selector)
	}
	if offset := srq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := srq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysRoleGroupBy is the group-by builder for SysRole entities.
type SysRoleGroupBy struct {
	selector
	build *SysRoleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (srgb *SysRoleGroupBy) Aggregate(fns ...AggregateFunc) *SysRoleGroupBy {
	srgb.fns = append(srgb.fns, fns...)
	return srgb
}

// Scan applies the selector query and scans the result into the given value.
func (srgb *SysRoleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, srgb.build.ctx, "GroupBy")
	if err := srgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysRoleQuery, *SysRoleGroupBy](ctx, srgb.build, srgb, srgb.build.inters, v)
}

func (srgb *SysRoleGroupBy) sqlScan(ctx context.Context, root *SysRoleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(srgb.fns))
	for _, fn := range srgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*srgb.flds)+len(srgb.fns))
		for _, f := range *srgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*srgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SysRoleSelect is the builder for selecting fields of SysRole entities.
type SysRoleSelect struct {
	*SysRoleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (srs *SysRoleSelect) Aggregate(fns ...AggregateFunc) *SysRoleSelect {
	srs.fns = append(srs.fns, fns...)
	return srs
}

// Scan applies the selector query and scans the result into the given value.
func (srs *SysRoleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, srs.ctx, "Select")
	if err := srs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SysRoleQuery, *SysRoleSelect](ctx, srs.SysRoleQuery, srs, srs.inters, v)
}

func (srs *SysRoleSelect) sqlScan(ctx context.Context, root *SysRoleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(srs.fns))
	for _, fn := range srs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*srs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
