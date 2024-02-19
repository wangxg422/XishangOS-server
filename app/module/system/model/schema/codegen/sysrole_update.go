// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/predicate"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdept"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysmenu"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysrole"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuser"
)

// SysRoleUpdate is the builder for updating SysRole entities.
type SysRoleUpdate struct {
	config
	hooks    []Hook
	mutation *SysRoleMutation
}

// Where appends a list predicates to the SysRoleUpdate builder.
func (sru *SysRoleUpdate) Where(ps ...predicate.SysRole) *SysRoleUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// SetUpdatedAt sets the "updated_at" field.
func (sru *SysRoleUpdate) SetUpdatedAt(t time.Time) *SysRoleUpdate {
	sru.mutation.SetUpdatedAt(t)
	return sru
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (sru *SysRoleUpdate) ClearUpdatedAt() *SysRoleUpdate {
	sru.mutation.ClearUpdatedAt()
	return sru
}

// SetDeleteAt sets the "delete_at" field.
func (sru *SysRoleUpdate) SetDeleteAt(t time.Time) *SysRoleUpdate {
	sru.mutation.SetDeleteAt(t)
	return sru
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (sru *SysRoleUpdate) ClearDeleteAt() *SysRoleUpdate {
	sru.mutation.ClearDeleteAt()
	return sru
}

// SetRemark sets the "remark" field.
func (sru *SysRoleUpdate) SetRemark(s string) *SysRoleUpdate {
	sru.mutation.SetRemark(s)
	return sru
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableRemark(s *string) *SysRoleUpdate {
	if s != nil {
		sru.SetRemark(*s)
	}
	return sru
}

// ClearRemark clears the value of the "remark" field.
func (sru *SysRoleUpdate) ClearRemark() *SysRoleUpdate {
	sru.mutation.ClearRemark()
	return sru
}

// SetListOrder sets the "list_order" field.
func (sru *SysRoleUpdate) SetListOrder(i int64) *SysRoleUpdate {
	sru.mutation.ResetListOrder()
	sru.mutation.SetListOrder(i)
	return sru
}

// SetNillableListOrder sets the "list_order" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableListOrder(i *int64) *SysRoleUpdate {
	if i != nil {
		sru.SetListOrder(*i)
	}
	return sru
}

// AddListOrder adds i to the "list_order" field.
func (sru *SysRoleUpdate) AddListOrder(i int64) *SysRoleUpdate {
	sru.mutation.AddListOrder(i)
	return sru
}

// ClearListOrder clears the value of the "list_order" field.
func (sru *SysRoleUpdate) ClearListOrder() *SysRoleUpdate {
	sru.mutation.ClearListOrder()
	return sru
}

// SetName sets the "name" field.
func (sru *SysRoleUpdate) SetName(s string) *SysRoleUpdate {
	sru.mutation.SetName(s)
	return sru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableName(s *string) *SysRoleUpdate {
	if s != nil {
		sru.SetName(*s)
	}
	return sru
}

// ClearName clears the value of the "name" field.
func (sru *SysRoleUpdate) ClearName() *SysRoleUpdate {
	sru.mutation.ClearName()
	return sru
}

// SetDataScope sets the "data_scope" field.
func (sru *SysRoleUpdate) SetDataScope(i int8) *SysRoleUpdate {
	sru.mutation.ResetDataScope()
	sru.mutation.SetDataScope(i)
	return sru
}

// SetNillableDataScope sets the "data_scope" field if the given value is not nil.
func (sru *SysRoleUpdate) SetNillableDataScope(i *int8) *SysRoleUpdate {
	if i != nil {
		sru.SetDataScope(*i)
	}
	return sru
}

// AddDataScope adds i to the "data_scope" field.
func (sru *SysRoleUpdate) AddDataScope(i int8) *SysRoleUpdate {
	sru.mutation.AddDataScope(i)
	return sru
}

// ClearDataScope clears the value of the "data_scope" field.
func (sru *SysRoleUpdate) ClearDataScope() *SysRoleUpdate {
	sru.mutation.ClearDataScope()
	return sru
}

// AddDeptIDs adds the "depts" edge to the SysDept entity by IDs.
func (sru *SysRoleUpdate) AddDeptIDs(ids ...int64) *SysRoleUpdate {
	sru.mutation.AddDeptIDs(ids...)
	return sru
}

// AddDepts adds the "depts" edges to the SysDept entity.
func (sru *SysRoleUpdate) AddDepts(s ...*SysDept) *SysRoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.AddDeptIDs(ids...)
}

// AddSysUserIDs adds the "sysUsers" edge to the SysUser entity by IDs.
func (sru *SysRoleUpdate) AddSysUserIDs(ids ...int64) *SysRoleUpdate {
	sru.mutation.AddSysUserIDs(ids...)
	return sru
}

// AddSysUsers adds the "sysUsers" edges to the SysUser entity.
func (sru *SysRoleUpdate) AddSysUsers(s ...*SysUser) *SysRoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.AddSysUserIDs(ids...)
}

// AddSysMenuIDs adds the "sysMenus" edge to the SysMenu entity by IDs.
func (sru *SysRoleUpdate) AddSysMenuIDs(ids ...int64) *SysRoleUpdate {
	sru.mutation.AddSysMenuIDs(ids...)
	return sru
}

// AddSysMenus adds the "sysMenus" edges to the SysMenu entity.
func (sru *SysRoleUpdate) AddSysMenus(s ...*SysMenu) *SysRoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.AddSysMenuIDs(ids...)
}

// Mutation returns the SysRoleMutation object of the builder.
func (sru *SysRoleUpdate) Mutation() *SysRoleMutation {
	return sru.mutation
}

// ClearDepts clears all "depts" edges to the SysDept entity.
func (sru *SysRoleUpdate) ClearDepts() *SysRoleUpdate {
	sru.mutation.ClearDepts()
	return sru
}

// RemoveDeptIDs removes the "depts" edge to SysDept entities by IDs.
func (sru *SysRoleUpdate) RemoveDeptIDs(ids ...int64) *SysRoleUpdate {
	sru.mutation.RemoveDeptIDs(ids...)
	return sru
}

// RemoveDepts removes "depts" edges to SysDept entities.
func (sru *SysRoleUpdate) RemoveDepts(s ...*SysDept) *SysRoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.RemoveDeptIDs(ids...)
}

// ClearSysUsers clears all "sysUsers" edges to the SysUser entity.
func (sru *SysRoleUpdate) ClearSysUsers() *SysRoleUpdate {
	sru.mutation.ClearSysUsers()
	return sru
}

// RemoveSysUserIDs removes the "sysUsers" edge to SysUser entities by IDs.
func (sru *SysRoleUpdate) RemoveSysUserIDs(ids ...int64) *SysRoleUpdate {
	sru.mutation.RemoveSysUserIDs(ids...)
	return sru
}

// RemoveSysUsers removes "sysUsers" edges to SysUser entities.
func (sru *SysRoleUpdate) RemoveSysUsers(s ...*SysUser) *SysRoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.RemoveSysUserIDs(ids...)
}

// ClearSysMenus clears all "sysMenus" edges to the SysMenu entity.
func (sru *SysRoleUpdate) ClearSysMenus() *SysRoleUpdate {
	sru.mutation.ClearSysMenus()
	return sru
}

// RemoveSysMenuIDs removes the "sysMenus" edge to SysMenu entities by IDs.
func (sru *SysRoleUpdate) RemoveSysMenuIDs(ids ...int64) *SysRoleUpdate {
	sru.mutation.RemoveSysMenuIDs(ids...)
	return sru
}

// RemoveSysMenus removes "sysMenus" edges to SysMenu entities.
func (sru *SysRoleUpdate) RemoveSysMenus(s ...*SysMenu) *SysRoleUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sru.RemoveSysMenuIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *SysRoleUpdate) Save(ctx context.Context) (int, error) {
	sru.defaults()
	return withHooks(ctx, sru.sqlSave, sru.mutation, sru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sru *SysRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *SysRoleUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *SysRoleUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sru *SysRoleUpdate) defaults() {
	if _, ok := sru.mutation.UpdatedAt(); !ok && !sru.mutation.UpdatedAtCleared() {
		v := sysrole.UpdateDefaultUpdatedAt()
		sru.mutation.SetUpdatedAt(v)
	}
	if _, ok := sru.mutation.DeleteAt(); !ok && !sru.mutation.DeleteAtCleared() {
		v := sysrole.UpdateDefaultDeleteAt()
		sru.mutation.SetDeleteAt(v)
	}
}

func (sru *SysRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sysrole.Table, sysrole.Columns, sqlgraph.NewFieldSpec(sysrole.FieldID, field.TypeInt64))
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if sru.mutation.CreatedAtCleared() {
		_spec.ClearField(sysrole.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := sru.mutation.UpdatedAt(); ok {
		_spec.SetField(sysrole.FieldUpdatedAt, field.TypeTime, value)
	}
	if sru.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysrole.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := sru.mutation.DeleteAt(); ok {
		_spec.SetField(sysrole.FieldDeleteAt, field.TypeTime, value)
	}
	if sru.mutation.DeleteAtCleared() {
		_spec.ClearField(sysrole.FieldDeleteAt, field.TypeTime)
	}
	if sru.mutation.StatusCleared() {
		_spec.ClearField(sysrole.FieldStatus, field.TypeInt8)
	}
	if value, ok := sru.mutation.Remark(); ok {
		_spec.SetField(sysrole.FieldRemark, field.TypeString, value)
	}
	if sru.mutation.RemarkCleared() {
		_spec.ClearField(sysrole.FieldRemark, field.TypeString)
	}
	if value, ok := sru.mutation.ListOrder(); ok {
		_spec.SetField(sysrole.FieldListOrder, field.TypeInt64, value)
	}
	if value, ok := sru.mutation.AddedListOrder(); ok {
		_spec.AddField(sysrole.FieldListOrder, field.TypeInt64, value)
	}
	if sru.mutation.ListOrderCleared() {
		_spec.ClearField(sysrole.FieldListOrder, field.TypeInt64)
	}
	if value, ok := sru.mutation.Name(); ok {
		_spec.SetField(sysrole.FieldName, field.TypeString, value)
	}
	if sru.mutation.NameCleared() {
		_spec.ClearField(sysrole.FieldName, field.TypeString)
	}
	if value, ok := sru.mutation.DataScope(); ok {
		_spec.SetField(sysrole.FieldDataScope, field.TypeInt8, value)
	}
	if value, ok := sru.mutation.AddedDataScope(); ok {
		_spec.AddField(sysrole.FieldDataScope, field.TypeInt8, value)
	}
	if sru.mutation.DataScopeCleared() {
		_spec.ClearField(sysrole.FieldDataScope, field.TypeInt8)
	}
	if sru.mutation.DeptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.DeptsTable,
			Columns: sysrole.DeptsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdept.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.RemovedDeptsIDs(); len(nodes) > 0 && !sru.mutation.DeptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.DeptsTable,
			Columns: sysrole.DeptsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdept.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.DeptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.DeptsTable,
			Columns: sysrole.DeptsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdept.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sru.mutation.SysUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.SysUsersTable,
			Columns: sysrole.SysUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.RemovedSysUsersIDs(); len(nodes) > 0 && !sru.mutation.SysUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.SysUsersTable,
			Columns: sysrole.SysUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.SysUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.SysUsersTable,
			Columns: sysrole.SysUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sru.mutation.SysMenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.SysMenusTable,
			Columns: sysrole.SysMenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysmenu.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.RemovedSysMenusIDs(); len(nodes) > 0 && !sru.mutation.SysMenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.SysMenusTable,
			Columns: sysrole.SysMenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysmenu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.SysMenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.SysMenusTable,
			Columns: sysrole.SysMenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysmenu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sru.mutation.done = true
	return n, nil
}

// SysRoleUpdateOne is the builder for updating a single SysRole entity.
type SysRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysRoleMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (sruo *SysRoleUpdateOne) SetUpdatedAt(t time.Time) *SysRoleUpdateOne {
	sruo.mutation.SetUpdatedAt(t)
	return sruo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (sruo *SysRoleUpdateOne) ClearUpdatedAt() *SysRoleUpdateOne {
	sruo.mutation.ClearUpdatedAt()
	return sruo
}

// SetDeleteAt sets the "delete_at" field.
func (sruo *SysRoleUpdateOne) SetDeleteAt(t time.Time) *SysRoleUpdateOne {
	sruo.mutation.SetDeleteAt(t)
	return sruo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (sruo *SysRoleUpdateOne) ClearDeleteAt() *SysRoleUpdateOne {
	sruo.mutation.ClearDeleteAt()
	return sruo
}

// SetRemark sets the "remark" field.
func (sruo *SysRoleUpdateOne) SetRemark(s string) *SysRoleUpdateOne {
	sruo.mutation.SetRemark(s)
	return sruo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableRemark(s *string) *SysRoleUpdateOne {
	if s != nil {
		sruo.SetRemark(*s)
	}
	return sruo
}

// ClearRemark clears the value of the "remark" field.
func (sruo *SysRoleUpdateOne) ClearRemark() *SysRoleUpdateOne {
	sruo.mutation.ClearRemark()
	return sruo
}

// SetListOrder sets the "list_order" field.
func (sruo *SysRoleUpdateOne) SetListOrder(i int64) *SysRoleUpdateOne {
	sruo.mutation.ResetListOrder()
	sruo.mutation.SetListOrder(i)
	return sruo
}

// SetNillableListOrder sets the "list_order" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableListOrder(i *int64) *SysRoleUpdateOne {
	if i != nil {
		sruo.SetListOrder(*i)
	}
	return sruo
}

// AddListOrder adds i to the "list_order" field.
func (sruo *SysRoleUpdateOne) AddListOrder(i int64) *SysRoleUpdateOne {
	sruo.mutation.AddListOrder(i)
	return sruo
}

// ClearListOrder clears the value of the "list_order" field.
func (sruo *SysRoleUpdateOne) ClearListOrder() *SysRoleUpdateOne {
	sruo.mutation.ClearListOrder()
	return sruo
}

// SetName sets the "name" field.
func (sruo *SysRoleUpdateOne) SetName(s string) *SysRoleUpdateOne {
	sruo.mutation.SetName(s)
	return sruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableName(s *string) *SysRoleUpdateOne {
	if s != nil {
		sruo.SetName(*s)
	}
	return sruo
}

// ClearName clears the value of the "name" field.
func (sruo *SysRoleUpdateOne) ClearName() *SysRoleUpdateOne {
	sruo.mutation.ClearName()
	return sruo
}

// SetDataScope sets the "data_scope" field.
func (sruo *SysRoleUpdateOne) SetDataScope(i int8) *SysRoleUpdateOne {
	sruo.mutation.ResetDataScope()
	sruo.mutation.SetDataScope(i)
	return sruo
}

// SetNillableDataScope sets the "data_scope" field if the given value is not nil.
func (sruo *SysRoleUpdateOne) SetNillableDataScope(i *int8) *SysRoleUpdateOne {
	if i != nil {
		sruo.SetDataScope(*i)
	}
	return sruo
}

// AddDataScope adds i to the "data_scope" field.
func (sruo *SysRoleUpdateOne) AddDataScope(i int8) *SysRoleUpdateOne {
	sruo.mutation.AddDataScope(i)
	return sruo
}

// ClearDataScope clears the value of the "data_scope" field.
func (sruo *SysRoleUpdateOne) ClearDataScope() *SysRoleUpdateOne {
	sruo.mutation.ClearDataScope()
	return sruo
}

// AddDeptIDs adds the "depts" edge to the SysDept entity by IDs.
func (sruo *SysRoleUpdateOne) AddDeptIDs(ids ...int64) *SysRoleUpdateOne {
	sruo.mutation.AddDeptIDs(ids...)
	return sruo
}

// AddDepts adds the "depts" edges to the SysDept entity.
func (sruo *SysRoleUpdateOne) AddDepts(s ...*SysDept) *SysRoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.AddDeptIDs(ids...)
}

// AddSysUserIDs adds the "sysUsers" edge to the SysUser entity by IDs.
func (sruo *SysRoleUpdateOne) AddSysUserIDs(ids ...int64) *SysRoleUpdateOne {
	sruo.mutation.AddSysUserIDs(ids...)
	return sruo
}

// AddSysUsers adds the "sysUsers" edges to the SysUser entity.
func (sruo *SysRoleUpdateOne) AddSysUsers(s ...*SysUser) *SysRoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.AddSysUserIDs(ids...)
}

// AddSysMenuIDs adds the "sysMenus" edge to the SysMenu entity by IDs.
func (sruo *SysRoleUpdateOne) AddSysMenuIDs(ids ...int64) *SysRoleUpdateOne {
	sruo.mutation.AddSysMenuIDs(ids...)
	return sruo
}

// AddSysMenus adds the "sysMenus" edges to the SysMenu entity.
func (sruo *SysRoleUpdateOne) AddSysMenus(s ...*SysMenu) *SysRoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.AddSysMenuIDs(ids...)
}

// Mutation returns the SysRoleMutation object of the builder.
func (sruo *SysRoleUpdateOne) Mutation() *SysRoleMutation {
	return sruo.mutation
}

// ClearDepts clears all "depts" edges to the SysDept entity.
func (sruo *SysRoleUpdateOne) ClearDepts() *SysRoleUpdateOne {
	sruo.mutation.ClearDepts()
	return sruo
}

// RemoveDeptIDs removes the "depts" edge to SysDept entities by IDs.
func (sruo *SysRoleUpdateOne) RemoveDeptIDs(ids ...int64) *SysRoleUpdateOne {
	sruo.mutation.RemoveDeptIDs(ids...)
	return sruo
}

// RemoveDepts removes "depts" edges to SysDept entities.
func (sruo *SysRoleUpdateOne) RemoveDepts(s ...*SysDept) *SysRoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.RemoveDeptIDs(ids...)
}

// ClearSysUsers clears all "sysUsers" edges to the SysUser entity.
func (sruo *SysRoleUpdateOne) ClearSysUsers() *SysRoleUpdateOne {
	sruo.mutation.ClearSysUsers()
	return sruo
}

// RemoveSysUserIDs removes the "sysUsers" edge to SysUser entities by IDs.
func (sruo *SysRoleUpdateOne) RemoveSysUserIDs(ids ...int64) *SysRoleUpdateOne {
	sruo.mutation.RemoveSysUserIDs(ids...)
	return sruo
}

// RemoveSysUsers removes "sysUsers" edges to SysUser entities.
func (sruo *SysRoleUpdateOne) RemoveSysUsers(s ...*SysUser) *SysRoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.RemoveSysUserIDs(ids...)
}

// ClearSysMenus clears all "sysMenus" edges to the SysMenu entity.
func (sruo *SysRoleUpdateOne) ClearSysMenus() *SysRoleUpdateOne {
	sruo.mutation.ClearSysMenus()
	return sruo
}

// RemoveSysMenuIDs removes the "sysMenus" edge to SysMenu entities by IDs.
func (sruo *SysRoleUpdateOne) RemoveSysMenuIDs(ids ...int64) *SysRoleUpdateOne {
	sruo.mutation.RemoveSysMenuIDs(ids...)
	return sruo
}

// RemoveSysMenus removes "sysMenus" edges to SysMenu entities.
func (sruo *SysRoleUpdateOne) RemoveSysMenus(s ...*SysMenu) *SysRoleUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sruo.RemoveSysMenuIDs(ids...)
}

// Where appends a list predicates to the SysRoleUpdate builder.
func (sruo *SysRoleUpdateOne) Where(ps ...predicate.SysRole) *SysRoleUpdateOne {
	sruo.mutation.Where(ps...)
	return sruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *SysRoleUpdateOne) Select(field string, fields ...string) *SysRoleUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated SysRole entity.
func (sruo *SysRoleUpdateOne) Save(ctx context.Context) (*SysRole, error) {
	sruo.defaults()
	return withHooks(ctx, sruo.sqlSave, sruo.mutation, sruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *SysRoleUpdateOne) SaveX(ctx context.Context) *SysRole {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *SysRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *SysRoleUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sruo *SysRoleUpdateOne) defaults() {
	if _, ok := sruo.mutation.UpdatedAt(); !ok && !sruo.mutation.UpdatedAtCleared() {
		v := sysrole.UpdateDefaultUpdatedAt()
		sruo.mutation.SetUpdatedAt(v)
	}
	if _, ok := sruo.mutation.DeleteAt(); !ok && !sruo.mutation.DeleteAtCleared() {
		v := sysrole.UpdateDefaultDeleteAt()
		sruo.mutation.SetDeleteAt(v)
	}
}

func (sruo *SysRoleUpdateOne) sqlSave(ctx context.Context) (_node *SysRole, err error) {
	_spec := sqlgraph.NewUpdateSpec(sysrole.Table, sysrole.Columns, sqlgraph.NewFieldSpec(sysrole.FieldID, field.TypeInt64))
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`codegen: missing "SysRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrole.FieldID)
		for _, f := range fields {
			if !sysrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
			}
			if f != sysrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if sruo.mutation.CreatedAtCleared() {
		_spec.ClearField(sysrole.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := sruo.mutation.UpdatedAt(); ok {
		_spec.SetField(sysrole.FieldUpdatedAt, field.TypeTime, value)
	}
	if sruo.mutation.UpdatedAtCleared() {
		_spec.ClearField(sysrole.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := sruo.mutation.DeleteAt(); ok {
		_spec.SetField(sysrole.FieldDeleteAt, field.TypeTime, value)
	}
	if sruo.mutation.DeleteAtCleared() {
		_spec.ClearField(sysrole.FieldDeleteAt, field.TypeTime)
	}
	if sruo.mutation.StatusCleared() {
		_spec.ClearField(sysrole.FieldStatus, field.TypeInt8)
	}
	if value, ok := sruo.mutation.Remark(); ok {
		_spec.SetField(sysrole.FieldRemark, field.TypeString, value)
	}
	if sruo.mutation.RemarkCleared() {
		_spec.ClearField(sysrole.FieldRemark, field.TypeString)
	}
	if value, ok := sruo.mutation.ListOrder(); ok {
		_spec.SetField(sysrole.FieldListOrder, field.TypeInt64, value)
	}
	if value, ok := sruo.mutation.AddedListOrder(); ok {
		_spec.AddField(sysrole.FieldListOrder, field.TypeInt64, value)
	}
	if sruo.mutation.ListOrderCleared() {
		_spec.ClearField(sysrole.FieldListOrder, field.TypeInt64)
	}
	if value, ok := sruo.mutation.Name(); ok {
		_spec.SetField(sysrole.FieldName, field.TypeString, value)
	}
	if sruo.mutation.NameCleared() {
		_spec.ClearField(sysrole.FieldName, field.TypeString)
	}
	if value, ok := sruo.mutation.DataScope(); ok {
		_spec.SetField(sysrole.FieldDataScope, field.TypeInt8, value)
	}
	if value, ok := sruo.mutation.AddedDataScope(); ok {
		_spec.AddField(sysrole.FieldDataScope, field.TypeInt8, value)
	}
	if sruo.mutation.DataScopeCleared() {
		_spec.ClearField(sysrole.FieldDataScope, field.TypeInt8)
	}
	if sruo.mutation.DeptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.DeptsTable,
			Columns: sysrole.DeptsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdept.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.RemovedDeptsIDs(); len(nodes) > 0 && !sruo.mutation.DeptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.DeptsTable,
			Columns: sysrole.DeptsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdept.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.DeptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.DeptsTable,
			Columns: sysrole.DeptsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdept.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sruo.mutation.SysUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.SysUsersTable,
			Columns: sysrole.SysUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.RemovedSysUsersIDs(); len(nodes) > 0 && !sruo.mutation.SysUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.SysUsersTable,
			Columns: sysrole.SysUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.SysUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysrole.SysUsersTable,
			Columns: sysrole.SysUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sruo.mutation.SysMenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.SysMenusTable,
			Columns: sysrole.SysMenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysmenu.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.RemovedSysMenusIDs(); len(nodes) > 0 && !sruo.mutation.SysMenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.SysMenusTable,
			Columns: sysrole.SysMenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysmenu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.SysMenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysrole.SysMenusTable,
			Columns: sysrole.SysMenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysmenu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SysRole{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sruo.mutation.done = true
	return _node, nil
}
