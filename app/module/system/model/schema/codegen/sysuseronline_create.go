// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuseronline"
)

// SysUserOnlineCreate is the builder for creating a SysUserOnline entity.
type SysUserOnlineCreate struct {
	config
	mutation *SysUserOnlineMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (suoc *SysUserOnlineCreate) SetUUID(i int64) *SysUserOnlineCreate {
	suoc.mutation.SetUUID(i)
	return suoc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (suoc *SysUserOnlineCreate) SetNillableUUID(i *int64) *SysUserOnlineCreate {
	if i != nil {
		suoc.SetUUID(*i)
	}
	return suoc
}

// SetToken sets the "token" field.
func (suoc *SysUserOnlineCreate) SetToken(i int64) *SysUserOnlineCreate {
	suoc.mutation.SetToken(i)
	return suoc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (suoc *SysUserOnlineCreate) SetNillableToken(i *int64) *SysUserOnlineCreate {
	if i != nil {
		suoc.SetToken(*i)
	}
	return suoc
}

// SetCreateTime sets the "create_time" field.
func (suoc *SysUserOnlineCreate) SetCreateTime(i int64) *SysUserOnlineCreate {
	suoc.mutation.SetCreateTime(i)
	return suoc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (suoc *SysUserOnlineCreate) SetNillableCreateTime(i *int64) *SysUserOnlineCreate {
	if i != nil {
		suoc.SetCreateTime(*i)
	}
	return suoc
}

// SetUserName sets the "user_name" field.
func (suoc *SysUserOnlineCreate) SetUserName(i int64) *SysUserOnlineCreate {
	suoc.mutation.SetUserName(i)
	return suoc
}

// SetNillableUserName sets the "user_name" field if the given value is not nil.
func (suoc *SysUserOnlineCreate) SetNillableUserName(i *int64) *SysUserOnlineCreate {
	if i != nil {
		suoc.SetUserName(*i)
	}
	return suoc
}

// SetIPAddr sets the "ip_addr" field.
func (suoc *SysUserOnlineCreate) SetIPAddr(i int64) *SysUserOnlineCreate {
	suoc.mutation.SetIPAddr(i)
	return suoc
}

// SetNillableIPAddr sets the "ip_addr" field if the given value is not nil.
func (suoc *SysUserOnlineCreate) SetNillableIPAddr(i *int64) *SysUserOnlineCreate {
	if i != nil {
		suoc.SetIPAddr(*i)
	}
	return suoc
}

// SetBrowser sets the "browser" field.
func (suoc *SysUserOnlineCreate) SetBrowser(i int64) *SysUserOnlineCreate {
	suoc.mutation.SetBrowser(i)
	return suoc
}

// SetNillableBrowser sets the "browser" field if the given value is not nil.
func (suoc *SysUserOnlineCreate) SetNillableBrowser(i *int64) *SysUserOnlineCreate {
	if i != nil {
		suoc.SetBrowser(*i)
	}
	return suoc
}

// SetOs sets the "os" field.
func (suoc *SysUserOnlineCreate) SetOs(i int64) *SysUserOnlineCreate {
	suoc.mutation.SetOs(i)
	return suoc
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (suoc *SysUserOnlineCreate) SetNillableOs(i *int64) *SysUserOnlineCreate {
	if i != nil {
		suoc.SetOs(*i)
	}
	return suoc
}

// SetID sets the "id" field.
func (suoc *SysUserOnlineCreate) SetID(i int64) *SysUserOnlineCreate {
	suoc.mutation.SetID(i)
	return suoc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (suoc *SysUserOnlineCreate) SetNillableID(i *int64) *SysUserOnlineCreate {
	if i != nil {
		suoc.SetID(*i)
	}
	return suoc
}

// Mutation returns the SysUserOnlineMutation object of the builder.
func (suoc *SysUserOnlineCreate) Mutation() *SysUserOnlineMutation {
	return suoc.mutation
}

// Save creates the SysUserOnline in the database.
func (suoc *SysUserOnlineCreate) Save(ctx context.Context) (*SysUserOnline, error) {
	suoc.defaults()
	return withHooks(ctx, suoc.sqlSave, suoc.mutation, suoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (suoc *SysUserOnlineCreate) SaveX(ctx context.Context) *SysUserOnline {
	v, err := suoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (suoc *SysUserOnlineCreate) Exec(ctx context.Context) error {
	_, err := suoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suoc *SysUserOnlineCreate) ExecX(ctx context.Context) {
	if err := suoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suoc *SysUserOnlineCreate) defaults() {
	if _, ok := suoc.mutation.ID(); !ok {
		v := sysuseronline.DefaultID()
		suoc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suoc *SysUserOnlineCreate) check() error {
	return nil
}

func (suoc *SysUserOnlineCreate) sqlSave(ctx context.Context) (*SysUserOnline, error) {
	if err := suoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := suoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, suoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	suoc.mutation.id = &_node.ID
	suoc.mutation.done = true
	return _node, nil
}

func (suoc *SysUserOnlineCreate) createSpec() (*SysUserOnline, *sqlgraph.CreateSpec) {
	var (
		_node = &SysUserOnline{config: suoc.config}
		_spec = sqlgraph.NewCreateSpec(sysuseronline.Table, sqlgraph.NewFieldSpec(sysuseronline.FieldID, field.TypeInt64))
	)
	if id, ok := suoc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := suoc.mutation.UUID(); ok {
		_spec.SetField(sysuseronline.FieldUUID, field.TypeInt64, value)
		_node.UUID = value
	}
	if value, ok := suoc.mutation.Token(); ok {
		_spec.SetField(sysuseronline.FieldToken, field.TypeInt64, value)
		_node.Token = value
	}
	if value, ok := suoc.mutation.CreateTime(); ok {
		_spec.SetField(sysuseronline.FieldCreateTime, field.TypeInt64, value)
		_node.CreateTime = value
	}
	if value, ok := suoc.mutation.UserName(); ok {
		_spec.SetField(sysuseronline.FieldUserName, field.TypeInt64, value)
		_node.UserName = value
	}
	if value, ok := suoc.mutation.IPAddr(); ok {
		_spec.SetField(sysuseronline.FieldIPAddr, field.TypeInt64, value)
		_node.IPAddr = value
	}
	if value, ok := suoc.mutation.Browser(); ok {
		_spec.SetField(sysuseronline.FieldBrowser, field.TypeInt64, value)
		_node.Browser = value
	}
	if value, ok := suoc.mutation.Os(); ok {
		_spec.SetField(sysuseronline.FieldOs, field.TypeInt64, value)
		_node.Os = value
	}
	return _node, _spec
}

// SysUserOnlineCreateBulk is the builder for creating many SysUserOnline entities in bulk.
type SysUserOnlineCreateBulk struct {
	config
	err      error
	builders []*SysUserOnlineCreate
}

// Save creates the SysUserOnline entities in the database.
func (suocb *SysUserOnlineCreateBulk) Save(ctx context.Context) ([]*SysUserOnline, error) {
	if suocb.err != nil {
		return nil, suocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(suocb.builders))
	nodes := make([]*SysUserOnline, len(suocb.builders))
	mutators := make([]Mutator, len(suocb.builders))
	for i := range suocb.builders {
		func(i int, root context.Context) {
			builder := suocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysUserOnlineMutation)
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
					_, err = mutators[i+1].Mutate(root, suocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, suocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, suocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (suocb *SysUserOnlineCreateBulk) SaveX(ctx context.Context) []*SysUserOnline {
	v, err := suocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (suocb *SysUserOnlineCreateBulk) Exec(ctx context.Context) error {
	_, err := suocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suocb *SysUserOnlineCreateBulk) ExecX(ctx context.Context) {
	if err := suocb.Exec(ctx); err != nil {
		panic(err)
	}
}
