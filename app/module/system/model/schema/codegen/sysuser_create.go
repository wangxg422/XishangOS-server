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
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/syspost"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuser"
)

// SysUserCreate is the builder for creating a SysUser entity.
type SysUserCreate struct {
	config
	mutation *SysUserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (suc *SysUserCreate) SetCreatedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetCreatedAt(t)
	return suc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableCreatedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetCreatedAt(*t)
	}
	return suc
}

// SetUpdatedAt sets the "updated_at" field.
func (suc *SysUserCreate) SetUpdatedAt(t time.Time) *SysUserCreate {
	suc.mutation.SetUpdatedAt(t)
	return suc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUpdatedAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetUpdatedAt(*t)
	}
	return suc
}

// SetDeleteAt sets the "delete_at" field.
func (suc *SysUserCreate) SetDeleteAt(t time.Time) *SysUserCreate {
	suc.mutation.SetDeleteAt(t)
	return suc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableDeleteAt(t *time.Time) *SysUserCreate {
	if t != nil {
		suc.SetDeleteAt(*t)
	}
	return suc
}

// SetRemark sets the "remark" field.
func (suc *SysUserCreate) SetRemark(s string) *SysUserCreate {
	suc.mutation.SetRemark(s)
	return suc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableRemark(s *string) *SysUserCreate {
	if s != nil {
		suc.SetRemark(*s)
	}
	return suc
}

// SetUserName sets the "user_name" field.
func (suc *SysUserCreate) SetUserName(s string) *SysUserCreate {
	suc.mutation.SetUserName(s)
	return suc
}

// SetUserNickname sets the "user_nickname" field.
func (suc *SysUserCreate) SetUserNickname(s string) *SysUserCreate {
	suc.mutation.SetUserNickname(s)
	return suc
}

// SetNillableUserNickname sets the "user_nickname" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUserNickname(s *string) *SysUserCreate {
	if s != nil {
		suc.SetUserNickname(*s)
	}
	return suc
}

// SetMobile sets the "mobile" field.
func (suc *SysUserCreate) SetMobile(s string) *SysUserCreate {
	suc.mutation.SetMobile(s)
	return suc
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableMobile(s *string) *SysUserCreate {
	if s != nil {
		suc.SetMobile(*s)
	}
	return suc
}

// SetBirthday sets the "birthday" field.
func (suc *SysUserCreate) SetBirthday(s string) *SysUserCreate {
	suc.mutation.SetBirthday(s)
	return suc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableBirthday(s *string) *SysUserCreate {
	if s != nil {
		suc.SetBirthday(*s)
	}
	return suc
}

// SetUserPassword sets the "user_password" field.
func (suc *SysUserCreate) SetUserPassword(s string) *SysUserCreate {
	suc.mutation.SetUserPassword(s)
	return suc
}

// SetNillableUserPassword sets the "user_password" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUserPassword(s *string) *SysUserCreate {
	if s != nil {
		suc.SetUserPassword(*s)
	}
	return suc
}

// SetUserSalt sets the "user_salt" field.
func (suc *SysUserCreate) SetUserSalt(s string) *SysUserCreate {
	suc.mutation.SetUserSalt(s)
	return suc
}

// SetNillableUserSalt sets the "user_salt" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUserSalt(s *string) *SysUserCreate {
	if s != nil {
		suc.SetUserSalt(*s)
	}
	return suc
}

// SetUserEmail sets the "user_email" field.
func (suc *SysUserCreate) SetUserEmail(s string) *SysUserCreate {
	suc.mutation.SetUserEmail(s)
	return suc
}

// SetNillableUserEmail sets the "user_email" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUserEmail(s *string) *SysUserCreate {
	if s != nil {
		suc.SetUserEmail(*s)
	}
	return suc
}

// SetSex sets the "sex" field.
func (suc *SysUserCreate) SetSex(i int8) *SysUserCreate {
	suc.mutation.SetSex(i)
	return suc
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableSex(i *int8) *SysUserCreate {
	if i != nil {
		suc.SetSex(*i)
	}
	return suc
}

// SetAvatar sets the "avatar" field.
func (suc *SysUserCreate) SetAvatar(s string) *SysUserCreate {
	suc.mutation.SetAvatar(s)
	return suc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableAvatar(s *string) *SysUserCreate {
	if s != nil {
		suc.SetAvatar(*s)
	}
	return suc
}

// SetIsAdmin sets the "is_admin" field.
func (suc *SysUserCreate) SetIsAdmin(i int8) *SysUserCreate {
	suc.mutation.SetIsAdmin(i)
	return suc
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableIsAdmin(i *int8) *SysUserCreate {
	if i != nil {
		suc.SetIsAdmin(*i)
	}
	return suc
}

// SetUserStatus sets the "user_status" field.
func (suc *SysUserCreate) SetUserStatus(i int8) *SysUserCreate {
	suc.mutation.SetUserStatus(i)
	return suc
}

// SetNillableUserStatus sets the "user_status" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableUserStatus(i *int8) *SysUserCreate {
	if i != nil {
		suc.SetUserStatus(*i)
	}
	return suc
}

// SetDeptID sets the "dept_id" field.
func (suc *SysUserCreate) SetDeptID(i int64) *SysUserCreate {
	suc.mutation.SetDeptID(i)
	return suc
}

// SetNillableDeptID sets the "dept_id" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableDeptID(i *int64) *SysUserCreate {
	if i != nil {
		suc.SetDeptID(*i)
	}
	return suc
}

// SetAddress sets the "address" field.
func (suc *SysUserCreate) SetAddress(s string) *SysUserCreate {
	suc.mutation.SetAddress(s)
	return suc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableAddress(s *string) *SysUserCreate {
	if s != nil {
		suc.SetAddress(*s)
	}
	return suc
}

// SetDescribe sets the "describe" field.
func (suc *SysUserCreate) SetDescribe(s string) *SysUserCreate {
	suc.mutation.SetDescribe(s)
	return suc
}

// SetNillableDescribe sets the "describe" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableDescribe(s *string) *SysUserCreate {
	if s != nil {
		suc.SetDescribe(*s)
	}
	return suc
}

// SetLastLoginIP sets the "last_login_ip" field.
func (suc *SysUserCreate) SetLastLoginIP(s string) *SysUserCreate {
	suc.mutation.SetLastLoginIP(s)
	return suc
}

// SetNillableLastLoginIP sets the "last_login_ip" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableLastLoginIP(s *string) *SysUserCreate {
	if s != nil {
		suc.SetLastLoginIP(*s)
	}
	return suc
}

// SetLastLoginTime sets the "last_login_time" field.
func (suc *SysUserCreate) SetLastLoginTime(s string) *SysUserCreate {
	suc.mutation.SetLastLoginTime(s)
	return suc
}

// SetNillableLastLoginTime sets the "last_login_time" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableLastLoginTime(s *string) *SysUserCreate {
	if s != nil {
		suc.SetLastLoginTime(*s)
	}
	return suc
}

// SetID sets the "id" field.
func (suc *SysUserCreate) SetID(i int64) *SysUserCreate {
	suc.mutation.SetID(i)
	return suc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (suc *SysUserCreate) SetNillableID(i *int64) *SysUserCreate {
	if i != nil {
		suc.SetID(*i)
	}
	return suc
}

// SetBelongToID sets the "belongTo" edge to the SysDept entity by ID.
func (suc *SysUserCreate) SetBelongToID(id int64) *SysUserCreate {
	suc.mutation.SetBelongToID(id)
	return suc
}

// SetNillableBelongToID sets the "belongTo" edge to the SysDept entity by ID if the given value is not nil.
func (suc *SysUserCreate) SetNillableBelongToID(id *int64) *SysUserCreate {
	if id != nil {
		suc = suc.SetBelongToID(*id)
	}
	return suc
}

// SetBelongTo sets the "belongTo" edge to the SysDept entity.
func (suc *SysUserCreate) SetBelongTo(s *SysDept) *SysUserCreate {
	return suc.SetBelongToID(s.ID)
}

// AddPostIDs adds the "posts" edge to the SysPost entity by IDs.
func (suc *SysUserCreate) AddPostIDs(ids ...int64) *SysUserCreate {
	suc.mutation.AddPostIDs(ids...)
	return suc
}

// AddPosts adds the "posts" edges to the SysPost entity.
func (suc *SysUserCreate) AddPosts(s ...*SysPost) *SysUserCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suc.AddPostIDs(ids...)
}

// Mutation returns the SysUserMutation object of the builder.
func (suc *SysUserCreate) Mutation() *SysUserMutation {
	return suc.mutation
}

// Save creates the SysUser in the database.
func (suc *SysUserCreate) Save(ctx context.Context) (*SysUser, error) {
	suc.defaults()
	return withHooks(ctx, suc.sqlSave, suc.mutation, suc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (suc *SysUserCreate) SaveX(ctx context.Context) *SysUser {
	v, err := suc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (suc *SysUserCreate) Exec(ctx context.Context) error {
	_, err := suc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suc *SysUserCreate) ExecX(ctx context.Context) {
	if err := suc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suc *SysUserCreate) defaults() {
	if _, ok := suc.mutation.CreatedAt(); !ok {
		v := sysuser.DefaultCreatedAt()
		suc.mutation.SetCreatedAt(v)
	}
	if _, ok := suc.mutation.UpdatedAt(); !ok {
		v := sysuser.DefaultUpdatedAt()
		suc.mutation.SetUpdatedAt(v)
	}
	if _, ok := suc.mutation.DeleteAt(); !ok {
		v := sysuser.DefaultDeleteAt()
		suc.mutation.SetDeleteAt(v)
	}
	if _, ok := suc.mutation.ID(); !ok {
		v := sysuser.DefaultID()
		suc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suc *SysUserCreate) check() error {
	if _, ok := suc.mutation.UserName(); !ok {
		return &ValidationError{Name: "user_name", err: errors.New(`codegen: missing required field "SysUser.user_name"`)}
	}
	return nil
}

func (suc *SysUserCreate) sqlSave(ctx context.Context) (*SysUser, error) {
	if err := suc.check(); err != nil {
		return nil, err
	}
	_node, _spec := suc.createSpec()
	if err := sqlgraph.CreateNode(ctx, suc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	suc.mutation.id = &_node.ID
	suc.mutation.done = true
	return _node, nil
}

func (suc *SysUserCreate) createSpec() (*SysUser, *sqlgraph.CreateSpec) {
	var (
		_node = &SysUser{config: suc.config}
		_spec = sqlgraph.NewCreateSpec(sysuser.Table, sqlgraph.NewFieldSpec(sysuser.FieldID, field.TypeInt64))
	)
	if id, ok := suc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := suc.mutation.CreatedAt(); ok {
		_spec.SetField(sysuser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := suc.mutation.UpdatedAt(); ok {
		_spec.SetField(sysuser.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := suc.mutation.DeleteAt(); ok {
		_spec.SetField(sysuser.FieldDeleteAt, field.TypeTime, value)
		_node.DeleteAt = value
	}
	if value, ok := suc.mutation.Remark(); ok {
		_spec.SetField(sysuser.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := suc.mutation.UserName(); ok {
		_spec.SetField(sysuser.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}
	if value, ok := suc.mutation.UserNickname(); ok {
		_spec.SetField(sysuser.FieldUserNickname, field.TypeString, value)
		_node.UserNickname = value
	}
	if value, ok := suc.mutation.Mobile(); ok {
		_spec.SetField(sysuser.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := suc.mutation.Birthday(); ok {
		_spec.SetField(sysuser.FieldBirthday, field.TypeString, value)
		_node.Birthday = value
	}
	if value, ok := suc.mutation.UserPassword(); ok {
		_spec.SetField(sysuser.FieldUserPassword, field.TypeString, value)
		_node.UserPassword = value
	}
	if value, ok := suc.mutation.UserSalt(); ok {
		_spec.SetField(sysuser.FieldUserSalt, field.TypeString, value)
		_node.UserSalt = value
	}
	if value, ok := suc.mutation.UserEmail(); ok {
		_spec.SetField(sysuser.FieldUserEmail, field.TypeString, value)
		_node.UserEmail = value
	}
	if value, ok := suc.mutation.Sex(); ok {
		_spec.SetField(sysuser.FieldSex, field.TypeInt8, value)
		_node.Sex = value
	}
	if value, ok := suc.mutation.Avatar(); ok {
		_spec.SetField(sysuser.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := suc.mutation.IsAdmin(); ok {
		_spec.SetField(sysuser.FieldIsAdmin, field.TypeInt8, value)
		_node.IsAdmin = value
	}
	if value, ok := suc.mutation.UserStatus(); ok {
		_spec.SetField(sysuser.FieldUserStatus, field.TypeInt8, value)
		_node.UserStatus = value
	}
	if value, ok := suc.mutation.DeptID(); ok {
		_spec.SetField(sysuser.FieldDeptID, field.TypeInt64, value)
		_node.DeptID = value
	}
	if value, ok := suc.mutation.Address(); ok {
		_spec.SetField(sysuser.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := suc.mutation.Describe(); ok {
		_spec.SetField(sysuser.FieldDescribe, field.TypeString, value)
		_node.Describe = value
	}
	if value, ok := suc.mutation.LastLoginIP(); ok {
		_spec.SetField(sysuser.FieldLastLoginIP, field.TypeString, value)
		_node.LastLoginIP = value
	}
	if value, ok := suc.mutation.LastLoginTime(); ok {
		_spec.SetField(sysuser.FieldLastLoginTime, field.TypeString, value)
		_node.LastLoginTime = value
	}
	if nodes := suc.mutation.BelongToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sysuser.BelongToTable,
			Columns: []string{sysuser.BelongToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sysdept.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sys_dept_sys_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := suc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sysuser.PostsTable,
			Columns: sysuser.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(syspost.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SysUserCreateBulk is the builder for creating many SysUser entities in bulk.
type SysUserCreateBulk struct {
	config
	err      error
	builders []*SysUserCreate
}

// Save creates the SysUser entities in the database.
func (sucb *SysUserCreateBulk) Save(ctx context.Context) ([]*SysUser, error) {
	if sucb.err != nil {
		return nil, sucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sucb.builders))
	nodes := make([]*SysUser, len(sucb.builders))
	mutators := make([]Mutator, len(sucb.builders))
	for i := range sucb.builders {
		func(i int, root context.Context) {
			builder := sucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysUserMutation)
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
					_, err = mutators[i+1].Mutate(root, sucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sucb *SysUserCreateBulk) SaveX(ctx context.Context) []*SysUser {
	v, err := sucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sucb *SysUserCreateBulk) Exec(ctx context.Context) error {
	_, err := sucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sucb *SysUserCreateBulk) ExecX(ctx context.Context) {
	if err := sucb.Exec(ctx); err != nil {
		panic(err)
	}
}
