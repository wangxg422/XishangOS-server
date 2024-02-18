// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdept"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysrole"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuser"
)

// SysDeptCreate is the builder for creating a SysDept entity.
type SysDeptCreate struct {
	config
	mutation *SysDeptMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sdc *SysDeptCreate) SetCreatedAt(t time.Time) *SysDeptCreate {
	sdc.mutation.SetCreatedAt(t)
	return sdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableCreatedAt(t *time.Time) *SysDeptCreate {
	if t != nil {
		sdc.SetCreatedAt(*t)
	}
	return sdc
}

// SetUpdatedAt sets the "updated_at" field.
func (sdc *SysDeptCreate) SetUpdatedAt(t time.Time) *SysDeptCreate {
	sdc.mutation.SetUpdatedAt(t)
	return sdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableUpdatedAt(t *time.Time) *SysDeptCreate {
	if t != nil {
		sdc.SetUpdatedAt(*t)
	}
	return sdc
}

// SetDeleteAt sets the "delete_at" field.
func (sdc *SysDeptCreate) SetDeleteAt(t time.Time) *SysDeptCreate {
	sdc.mutation.SetDeleteAt(t)
	return sdc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableDeleteAt(t *time.Time) *SysDeptCreate {
	if t != nil {
		sdc.SetDeleteAt(*t)
	}
	return sdc
}

// SetCreatedBy sets the "created_by" field.
func (sdc *SysDeptCreate) SetCreatedBy(i int64) *SysDeptCreate {
	sdc.mutation.SetCreatedBy(i)
	return sdc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableCreatedBy(i *int64) *SysDeptCreate {
	if i != nil {
		sdc.SetCreatedBy(*i)
	}
	return sdc
}

// SetUpdatedBy sets the "updated_by" field.
func (sdc *SysDeptCreate) SetUpdatedBy(i int64) *SysDeptCreate {
	sdc.mutation.SetUpdatedBy(i)
	return sdc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableUpdatedBy(i *int64) *SysDeptCreate {
	if i != nil {
		sdc.SetUpdatedBy(*i)
	}
	return sdc
}

// SetDeleteBy sets the "delete_by" field.
func (sdc *SysDeptCreate) SetDeleteBy(i int64) *SysDeptCreate {
	sdc.mutation.SetDeleteBy(i)
	return sdc
}

// SetNillableDeleteBy sets the "delete_by" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableDeleteBy(i *int64) *SysDeptCreate {
	if i != nil {
		sdc.SetDeleteBy(*i)
	}
	return sdc
}

// SetRemark sets the "remark" field.
func (sdc *SysDeptCreate) SetRemark(s string) *SysDeptCreate {
	sdc.mutation.SetRemark(s)
	return sdc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableRemark(s *string) *SysDeptCreate {
	if s != nil {
		sdc.SetRemark(*s)
	}
	return sdc
}

// SetStatus sets the "status" field.
func (sdc *SysDeptCreate) SetStatus(i int8) *SysDeptCreate {
	sdc.mutation.SetStatus(i)
	return sdc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableStatus(i *int8) *SysDeptCreate {
	if i != nil {
		sdc.SetStatus(*i)
	}
	return sdc
}

// SetSort sets the "sort" field.
func (sdc *SysDeptCreate) SetSort(i int) *SysDeptCreate {
	sdc.mutation.SetSort(i)
	return sdc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableSort(i *int) *SysDeptCreate {
	if i != nil {
		sdc.SetSort(*i)
	}
	return sdc
}

// SetParentID sets the "parent_id" field.
func (sdc *SysDeptCreate) SetParentID(i int64) *SysDeptCreate {
	sdc.mutation.SetParentID(i)
	return sdc
}

// SetAncestors sets the "ancestors" field.
func (sdc *SysDeptCreate) SetAncestors(s string) *SysDeptCreate {
	sdc.mutation.SetAncestors(s)
	return sdc
}

// SetNillableAncestors sets the "ancestors" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableAncestors(s *string) *SysDeptCreate {
	if s != nil {
		sdc.SetAncestors(*s)
	}
	return sdc
}

// SetDeptName sets the "dept_name" field.
func (sdc *SysDeptCreate) SetDeptName(s string) *SysDeptCreate {
	sdc.mutation.SetDeptName(s)
	return sdc
}

// SetNillableDeptName sets the "dept_name" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableDeptName(s *string) *SysDeptCreate {
	if s != nil {
		sdc.SetDeptName(*s)
	}
	return sdc
}

// SetDeptCode sets the "dept_code" field.
func (sdc *SysDeptCreate) SetDeptCode(s string) *SysDeptCreate {
	sdc.mutation.SetDeptCode(s)
	return sdc
}

// SetNillableDeptCode sets the "dept_code" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableDeptCode(s *string) *SysDeptCreate {
	if s != nil {
		sdc.SetDeptCode(*s)
	}
	return sdc
}

// SetLeader sets the "leader" field.
func (sdc *SysDeptCreate) SetLeader(s string) *SysDeptCreate {
	sdc.mutation.SetLeader(s)
	return sdc
}

// SetNillableLeader sets the "leader" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableLeader(s *string) *SysDeptCreate {
	if s != nil {
		sdc.SetLeader(*s)
	}
	return sdc
}

// SetPhone sets the "phone" field.
func (sdc *SysDeptCreate) SetPhone(s string) *SysDeptCreate {
	sdc.mutation.SetPhone(s)
	return sdc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillablePhone(s *string) *SysDeptCreate {
	if s != nil {
		sdc.SetPhone(*s)
	}
	return sdc
}

// SetEmail sets the "email" field.
func (sdc *SysDeptCreate) SetEmail(s string) *SysDeptCreate {
	sdc.mutation.SetEmail(s)
	return sdc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableEmail(s *string) *SysDeptCreate {
	if s != nil {
		sdc.SetEmail(*s)
	}
	return sdc
}

// SetID sets the "id" field.
func (sdc *SysDeptCreate) SetID(i int64) *SysDeptCreate {
	sdc.mutation.SetID(i)
	return sdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sdc *SysDeptCreate) SetNillableID(i *int64) *SysDeptCreate {
	if i != nil {
		sdc.SetID(*i)
	}
	return sdc
}

// AddSysUserIDs adds the "sys_user" edge to the SysUser entity by IDs.
func (sdc *SysDeptCreate) AddSysUserIDs(ids ...int64) *SysDeptCreate {
	sdc.mutation.AddSysUserIDs(ids...)
	return sdc
}

// AddSysUser adds the "sys_user" edges to the SysUser entity.
func (sdc *SysDeptCreate) AddSysUser(s ...*SysUser) *SysDeptCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdc.AddSysUserIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the SysRole entity by IDs.
func (sdc *SysDeptCreate) AddRoleIDs(ids ...int64) *SysDeptCreate {
	sdc.mutation.AddRoleIDs(ids...)
	return sdc
}

// AddRoles adds the "roles" edges to the SysRole entity.
func (sdc *SysDeptCreate) AddRoles(s ...*SysRole) *SysDeptCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sdc.AddRoleIDs(ids...)
}

// Mutation returns the SysDeptMutation object of the builder.
func (sdc *SysDeptCreate) Mutation() *SysDeptMutation {
	return sdc.mutation
}

// Save creates the SysDept in the database.
func (sdc *SysDeptCreate) Save(ctx context.Context) (*SysDept, error) {
	sdc.defaults()
	return withHooks(ctx, sdc.sqlSave, sdc.mutation, sdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *SysDeptCreate) SaveX(ctx context.Context) *SysDept {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdc *SysDeptCreate) Exec(ctx context.Context) error {
	_, err := sdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdc *SysDeptCreate) ExecX(ctx context.Context) {
	if err := sdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdc *SysDeptCreate) defaults() {
	if _, ok := sdc.mutation.CreatedAt(); !ok {
		v := sysdept.DefaultCreatedAt()
		sdc.mutation.SetCreatedAt(v)
	}
	if _, ok := sdc.mutation.UpdatedAt(); !ok {
		v := sysdept.DefaultUpdatedAt()
		sdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sdc.mutation.DeleteAt(); !ok {
		v := sysdept.DefaultDeleteAt()
		sdc.mutation.SetDeleteAt(v)
	}
	if _, ok := sdc.mutation.Status(); !ok {
		v := sysdept.DefaultStatus
		sdc.mutation.SetStatus(v)
	}
	if _, ok := sdc.mutation.ID(); !ok {
		v := sysdept.DefaultID()
		sdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdc *SysDeptCreate) check() error {
	if _, ok := sdc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`codegen: missing required field "SysDept.parent_id"`)}
	}
	return nil
}

func (sdc *SysDeptCreate) sqlSave(ctx context.Context) (*SysDept, error) {
	if err := sdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	sdc.mutation.id = &_node.ID
	sdc.mutation.done = true
	return _node, nil
}

func (sdc *SysDeptCreate) createSpec() (*SysDept, *sqlgraph.CreateSpec) {
	var (
		_node = &SysDept{config: sdc.config}
		_spec = sqlgraph.NewCreateSpec(sysdept.Table, sqlgraph.NewFieldSpec(sysdept.FieldID, field.TypeInt64))
	)
	if id, ok := sdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sdc.mutation.CreatedAt(); ok {
		_spec.SetField(sysdept.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sdc.mutation.UpdatedAt(); ok {
		_spec.SetField(sysdept.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sdc.mutation.DeleteAt(); ok {
		_spec.SetField(sysdept.FieldDeleteAt, field.TypeTime, value)
		_node.DeleteAt = value
	}
	if value, ok := sdc.mutation.CreatedBy(); ok {
		_spec.SetField(sysdept.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := sdc.mutation.UpdatedBy(); ok {
		_spec.SetField(sysdept.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := sdc.mutation.DeleteBy(); ok {
		_spec.SetField(sysdept.FieldDeleteBy, field.TypeInt64, value)
		_node.DeleteBy = value
	}
	if value, ok := sdc.mutation.Remark(); ok {
		_spec.SetField(sysdept.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := sdc.mutation.Status(); ok {
		_spec.SetField(sysdept.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := sdc.mutation.Sort(); ok {
		_spec.SetField(sysdept.FieldSort, field.TypeInt, value)
		_node.Sort = value
	}
	if value, ok := sdc.mutation.ParentID(); ok {
		_spec.SetField(sysdept.FieldParentID, field.TypeInt64, value)
		_node.ParentID = value
	}
	if value, ok := sdc.mutation.Ancestors(); ok {
		_spec.SetField(sysdept.FieldAncestors, field.TypeString, value)
		_node.Ancestors = value
	}
	if value, ok := sdc.mutation.DeptName(); ok {
		_spec.SetField(sysdept.FieldDeptName, field.TypeString, value)
		_node.DeptName = value
	}
	if value, ok := sdc.mutation.DeptCode(); ok {
		_spec.SetField(sysdept.FieldDeptCode, field.TypeString, value)
		_node.DeptCode = value
	}
	if value, ok := sdc.mutation.Leader(); ok {
		_spec.SetField(sysdept.FieldLeader, field.TypeString, value)
		_node.Leader = value
	}
	if value, ok := sdc.mutation.Phone(); ok {
		_spec.SetField(sysdept.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := sdc.mutation.Email(); ok {
		_spec.SetField(sysdept.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if nodes := sdc.mutation.SysUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sysdept.SysUserTable,
			Columns: []string{sysdept.SysUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sdc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sysdept.RolesTable,
			Columns: sysdept.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysrole.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SysDeptCreateBulk is the builder for creating many SysDept entities in bulk.
type SysDeptCreateBulk struct {
	config
	err      error
	builders []*SysDeptCreate
}

// Save creates the SysDept entities in the database.
func (sdcb *SysDeptCreateBulk) Save(ctx context.Context) ([]*SysDept, error) {
	if sdcb.err != nil {
		return nil, sdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*SysDept, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysDeptMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *SysDeptCreateBulk) SaveX(ctx context.Context) []*SysDept {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdcb *SysDeptCreateBulk) Exec(ctx context.Context) error {
	_, err := sdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdcb *SysDeptCreateBulk) ExecX(ctx context.Context) {
	if err := sdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
