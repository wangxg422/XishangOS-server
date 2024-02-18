// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysconfig"
)

// SysConfigCreate is the builder for creating a SysConfig entity.
type SysConfigCreate struct {
	config
	mutation *SysConfigMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (scc *SysConfigCreate) SetCreatedAt(t time.Time) *SysConfigCreate {
	scc.mutation.SetCreatedAt(t)
	return scc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableCreatedAt(t *time.Time) *SysConfigCreate {
	if t != nil {
		scc.SetCreatedAt(*t)
	}
	return scc
}

// SetUpdatedAt sets the "updated_at" field.
func (scc *SysConfigCreate) SetUpdatedAt(t time.Time) *SysConfigCreate {
	scc.mutation.SetUpdatedAt(t)
	return scc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableUpdatedAt(t *time.Time) *SysConfigCreate {
	if t != nil {
		scc.SetUpdatedAt(*t)
	}
	return scc
}

// SetDeleteAt sets the "delete_at" field.
func (scc *SysConfigCreate) SetDeleteAt(t time.Time) *SysConfigCreate {
	scc.mutation.SetDeleteAt(t)
	return scc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableDeleteAt(t *time.Time) *SysConfigCreate {
	if t != nil {
		scc.SetDeleteAt(*t)
	}
	return scc
}

// SetCreatedBy sets the "created_by" field.
func (scc *SysConfigCreate) SetCreatedBy(i int64) *SysConfigCreate {
	scc.mutation.SetCreatedBy(i)
	return scc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableCreatedBy(i *int64) *SysConfigCreate {
	if i != nil {
		scc.SetCreatedBy(*i)
	}
	return scc
}

// SetUpdatedBy sets the "updated_by" field.
func (scc *SysConfigCreate) SetUpdatedBy(i int64) *SysConfigCreate {
	scc.mutation.SetUpdatedBy(i)
	return scc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableUpdatedBy(i *int64) *SysConfigCreate {
	if i != nil {
		scc.SetUpdatedBy(*i)
	}
	return scc
}

// SetDeleteBy sets the "delete_by" field.
func (scc *SysConfigCreate) SetDeleteBy(i int64) *SysConfigCreate {
	scc.mutation.SetDeleteBy(i)
	return scc
}

// SetNillableDeleteBy sets the "delete_by" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableDeleteBy(i *int64) *SysConfigCreate {
	if i != nil {
		scc.SetDeleteBy(*i)
	}
	return scc
}

// SetRemark sets the "remark" field.
func (scc *SysConfigCreate) SetRemark(s string) *SysConfigCreate {
	scc.mutation.SetRemark(s)
	return scc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableRemark(s *string) *SysConfigCreate {
	if s != nil {
		scc.SetRemark(*s)
	}
	return scc
}

// SetConfigName sets the "config_name" field.
func (scc *SysConfigCreate) SetConfigName(s string) *SysConfigCreate {
	scc.mutation.SetConfigName(s)
	return scc
}

// SetNillableConfigName sets the "config_name" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableConfigName(s *string) *SysConfigCreate {
	if s != nil {
		scc.SetConfigName(*s)
	}
	return scc
}

// SetConfigKey sets the "config_key" field.
func (scc *SysConfigCreate) SetConfigKey(s string) *SysConfigCreate {
	scc.mutation.SetConfigKey(s)
	return scc
}

// SetNillableConfigKey sets the "config_key" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableConfigKey(s *string) *SysConfigCreate {
	if s != nil {
		scc.SetConfigKey(*s)
	}
	return scc
}

// SetConfigValue sets the "config_value" field.
func (scc *SysConfigCreate) SetConfigValue(s string) *SysConfigCreate {
	scc.mutation.SetConfigValue(s)
	return scc
}

// SetNillableConfigValue sets the "config_value" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableConfigValue(s *string) *SysConfigCreate {
	if s != nil {
		scc.SetConfigValue(*s)
	}
	return scc
}

// SetConfigType sets the "config_type" field.
func (scc *SysConfigCreate) SetConfigType(s string) *SysConfigCreate {
	scc.mutation.SetConfigType(s)
	return scc
}

// SetNillableConfigType sets the "config_type" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableConfigType(s *string) *SysConfigCreate {
	if s != nil {
		scc.SetConfigType(*s)
	}
	return scc
}

// SetID sets the "id" field.
func (scc *SysConfigCreate) SetID(i int64) *SysConfigCreate {
	scc.mutation.SetID(i)
	return scc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (scc *SysConfigCreate) SetNillableID(i *int64) *SysConfigCreate {
	if i != nil {
		scc.SetID(*i)
	}
	return scc
}

// Mutation returns the SysConfigMutation object of the builder.
func (scc *SysConfigCreate) Mutation() *SysConfigMutation {
	return scc.mutation
}

// Save creates the SysConfig in the database.
func (scc *SysConfigCreate) Save(ctx context.Context) (*SysConfig, error) {
	scc.defaults()
	return withHooks(ctx, scc.sqlSave, scc.mutation, scc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (scc *SysConfigCreate) SaveX(ctx context.Context) *SysConfig {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *SysConfigCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *SysConfigCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *SysConfigCreate) defaults() {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		v := sysconfig.DefaultCreatedAt()
		scc.mutation.SetCreatedAt(v)
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		v := sysconfig.DefaultUpdatedAt()
		scc.mutation.SetUpdatedAt(v)
	}
	if _, ok := scc.mutation.DeleteAt(); !ok {
		v := sysconfig.DefaultDeleteAt()
		scc.mutation.SetDeleteAt(v)
	}
	if _, ok := scc.mutation.ID(); !ok {
		v := sysconfig.DefaultID()
		scc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *SysConfigCreate) check() error {
	return nil
}

func (scc *SysConfigCreate) sqlSave(ctx context.Context) (*SysConfig, error) {
	if err := scc.check(); err != nil {
		return nil, err
	}
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	scc.mutation.id = &_node.ID
	scc.mutation.done = true
	return _node, nil
}

func (scc *SysConfigCreate) createSpec() (*SysConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &SysConfig{config: scc.config}
		_spec = sqlgraph.NewCreateSpec(sysconfig.Table, sqlgraph.NewFieldSpec(sysconfig.FieldID, field.TypeInt64))
	)
	if id, ok := scc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := scc.mutation.CreatedAt(); ok {
		_spec.SetField(sysconfig.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := scc.mutation.UpdatedAt(); ok {
		_spec.SetField(sysconfig.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := scc.mutation.DeleteAt(); ok {
		_spec.SetField(sysconfig.FieldDeleteAt, field.TypeTime, value)
		_node.DeleteAt = value
	}
	if value, ok := scc.mutation.CreatedBy(); ok {
		_spec.SetField(sysconfig.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := scc.mutation.UpdatedBy(); ok {
		_spec.SetField(sysconfig.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := scc.mutation.DeleteBy(); ok {
		_spec.SetField(sysconfig.FieldDeleteBy, field.TypeInt64, value)
		_node.DeleteBy = value
	}
	if value, ok := scc.mutation.Remark(); ok {
		_spec.SetField(sysconfig.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := scc.mutation.ConfigName(); ok {
		_spec.SetField(sysconfig.FieldConfigName, field.TypeString, value)
		_node.ConfigName = value
	}
	if value, ok := scc.mutation.ConfigKey(); ok {
		_spec.SetField(sysconfig.FieldConfigKey, field.TypeString, value)
		_node.ConfigKey = value
	}
	if value, ok := scc.mutation.ConfigValue(); ok {
		_spec.SetField(sysconfig.FieldConfigValue, field.TypeString, value)
		_node.ConfigValue = value
	}
	if value, ok := scc.mutation.ConfigType(); ok {
		_spec.SetField(sysconfig.FieldConfigType, field.TypeString, value)
		_node.ConfigType = value
	}
	return _node, _spec
}

// SysConfigCreateBulk is the builder for creating many SysConfig entities in bulk.
type SysConfigCreateBulk struct {
	config
	err      error
	builders []*SysConfigCreate
}

// Save creates the SysConfig entities in the database.
func (sccb *SysConfigCreateBulk) Save(ctx context.Context) ([]*SysConfig, error) {
	if sccb.err != nil {
		return nil, sccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*SysConfig, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SysConfigMutation)
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
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *SysConfigCreateBulk) SaveX(ctx context.Context) []*SysConfig {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *SysConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *SysConfigCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}