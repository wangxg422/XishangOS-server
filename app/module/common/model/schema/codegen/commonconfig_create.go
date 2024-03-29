// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/common/model/schema/codegen/commonconfig"
)

// CommonConfigCreate is the builder for creating a CommonConfig entity.
type CommonConfigCreate struct {
	config
	mutation *CommonConfigMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ccc *CommonConfigCreate) SetCreatedAt(t time.Time) *CommonConfigCreate {
	ccc.mutation.SetCreatedAt(t)
	return ccc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableCreatedAt(t *time.Time) *CommonConfigCreate {
	if t != nil {
		ccc.SetCreatedAt(*t)
	}
	return ccc
}

// SetUpdatedAt sets the "updated_at" field.
func (ccc *CommonConfigCreate) SetUpdatedAt(t time.Time) *CommonConfigCreate {
	ccc.mutation.SetUpdatedAt(t)
	return ccc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableUpdatedAt(t *time.Time) *CommonConfigCreate {
	if t != nil {
		ccc.SetUpdatedAt(*t)
	}
	return ccc
}

// SetDeleteAt sets the "delete_at" field.
func (ccc *CommonConfigCreate) SetDeleteAt(t time.Time) *CommonConfigCreate {
	ccc.mutation.SetDeleteAt(t)
	return ccc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableDeleteAt(t *time.Time) *CommonConfigCreate {
	if t != nil {
		ccc.SetDeleteAt(*t)
	}
	return ccc
}

// SetStatus sets the "status" field.
func (ccc *CommonConfigCreate) SetStatus(i int8) *CommonConfigCreate {
	ccc.mutation.SetStatus(i)
	return ccc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableStatus(i *int8) *CommonConfigCreate {
	if i != nil {
		ccc.SetStatus(*i)
	}
	return ccc
}

// SetCreatedBy sets the "created_by" field.
func (ccc *CommonConfigCreate) SetCreatedBy(i int64) *CommonConfigCreate {
	ccc.mutation.SetCreatedBy(i)
	return ccc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableCreatedBy(i *int64) *CommonConfigCreate {
	if i != nil {
		ccc.SetCreatedBy(*i)
	}
	return ccc
}

// SetUpdatedBy sets the "updated_by" field.
func (ccc *CommonConfigCreate) SetUpdatedBy(i int64) *CommonConfigCreate {
	ccc.mutation.SetUpdatedBy(i)
	return ccc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableUpdatedBy(i *int64) *CommonConfigCreate {
	if i != nil {
		ccc.SetUpdatedBy(*i)
	}
	return ccc
}

// SetDeleteBy sets the "delete_by" field.
func (ccc *CommonConfigCreate) SetDeleteBy(i int64) *CommonConfigCreate {
	ccc.mutation.SetDeleteBy(i)
	return ccc
}

// SetNillableDeleteBy sets the "delete_by" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableDeleteBy(i *int64) *CommonConfigCreate {
	if i != nil {
		ccc.SetDeleteBy(*i)
	}
	return ccc
}

// SetRemark sets the "remark" field.
func (ccc *CommonConfigCreate) SetRemark(s string) *CommonConfigCreate {
	ccc.mutation.SetRemark(s)
	return ccc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableRemark(s *string) *CommonConfigCreate {
	if s != nil {
		ccc.SetRemark(*s)
	}
	return ccc
}

// SetDelFlag sets the "del_flag" field.
func (ccc *CommonConfigCreate) SetDelFlag(i int8) *CommonConfigCreate {
	ccc.mutation.SetDelFlag(i)
	return ccc
}

// SetNillableDelFlag sets the "del_flag" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableDelFlag(i *int8) *CommonConfigCreate {
	if i != nil {
		ccc.SetDelFlag(*i)
	}
	return ccc
}

// SetConfigName sets the "config_name" field.
func (ccc *CommonConfigCreate) SetConfigName(s string) *CommonConfigCreate {
	ccc.mutation.SetConfigName(s)
	return ccc
}

// SetNillableConfigName sets the "config_name" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableConfigName(s *string) *CommonConfigCreate {
	if s != nil {
		ccc.SetConfigName(*s)
	}
	return ccc
}

// SetConfigKey sets the "config_key" field.
func (ccc *CommonConfigCreate) SetConfigKey(s string) *CommonConfigCreate {
	ccc.mutation.SetConfigKey(s)
	return ccc
}

// SetNillableConfigKey sets the "config_key" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableConfigKey(s *string) *CommonConfigCreate {
	if s != nil {
		ccc.SetConfigKey(*s)
	}
	return ccc
}

// SetConfigValue sets the "config_value" field.
func (ccc *CommonConfigCreate) SetConfigValue(s string) *CommonConfigCreate {
	ccc.mutation.SetConfigValue(s)
	return ccc
}

// SetNillableConfigValue sets the "config_value" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableConfigValue(s *string) *CommonConfigCreate {
	if s != nil {
		ccc.SetConfigValue(*s)
	}
	return ccc
}

// SetConfigType sets the "config_type" field.
func (ccc *CommonConfigCreate) SetConfigType(i int8) *CommonConfigCreate {
	ccc.mutation.SetConfigType(i)
	return ccc
}

// SetNillableConfigType sets the "config_type" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableConfigType(i *int8) *CommonConfigCreate {
	if i != nil {
		ccc.SetConfigType(*i)
	}
	return ccc
}

// SetID sets the "id" field.
func (ccc *CommonConfigCreate) SetID(i int64) *CommonConfigCreate {
	ccc.mutation.SetID(i)
	return ccc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ccc *CommonConfigCreate) SetNillableID(i *int64) *CommonConfigCreate {
	if i != nil {
		ccc.SetID(*i)
	}
	return ccc
}

// Mutation returns the CommonConfigMutation object of the builder.
func (ccc *CommonConfigCreate) Mutation() *CommonConfigMutation {
	return ccc.mutation
}

// Save creates the CommonConfig in the database.
func (ccc *CommonConfigCreate) Save(ctx context.Context) (*CommonConfig, error) {
	if err := ccc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ccc.sqlSave, ccc.mutation, ccc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ccc *CommonConfigCreate) SaveX(ctx context.Context) *CommonConfig {
	v, err := ccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccc *CommonConfigCreate) Exec(ctx context.Context) error {
	_, err := ccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccc *CommonConfigCreate) ExecX(ctx context.Context) {
	if err := ccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccc *CommonConfigCreate) defaults() error {
	if _, ok := ccc.mutation.CreatedAt(); !ok {
		if commonconfig.DefaultCreatedAt == nil {
			return fmt.Errorf("codegen: uninitialized commonconfig.DefaultCreatedAt (forgotten import codegen/runtime?)")
		}
		v := commonconfig.DefaultCreatedAt()
		ccc.mutation.SetCreatedAt(v)
	}
	if _, ok := ccc.mutation.UpdatedAt(); !ok {
		if commonconfig.DefaultUpdatedAt == nil {
			return fmt.Errorf("codegen: uninitialized commonconfig.DefaultUpdatedAt (forgotten import codegen/runtime?)")
		}
		v := commonconfig.DefaultUpdatedAt()
		ccc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ccc.mutation.Status(); !ok {
		v := commonconfig.DefaultStatus
		ccc.mutation.SetStatus(v)
	}
	if _, ok := ccc.mutation.DelFlag(); !ok {
		v := commonconfig.DefaultDelFlag
		ccc.mutation.SetDelFlag(v)
	}
	if _, ok := ccc.mutation.ID(); !ok {
		if commonconfig.DefaultID == nil {
			return fmt.Errorf("codegen: uninitialized commonconfig.DefaultID (forgotten import codegen/runtime?)")
		}
		v := commonconfig.DefaultID()
		ccc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ccc *CommonConfigCreate) check() error {
	if _, ok := ccc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`codegen: missing required field "CommonConfig.status"`)}
	}
	if _, ok := ccc.mutation.DelFlag(); !ok {
		return &ValidationError{Name: "del_flag", err: errors.New(`codegen: missing required field "CommonConfig.del_flag"`)}
	}
	return nil
}

func (ccc *CommonConfigCreate) sqlSave(ctx context.Context) (*CommonConfig, error) {
	if err := ccc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ccc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ccc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ccc.mutation.id = &_node.ID
	ccc.mutation.done = true
	return _node, nil
}

func (ccc *CommonConfigCreate) createSpec() (*CommonConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &CommonConfig{config: ccc.config}
		_spec = sqlgraph.NewCreateSpec(commonconfig.Table, sqlgraph.NewFieldSpec(commonconfig.FieldID, field.TypeInt64))
	)
	if id, ok := ccc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ccc.mutation.CreatedAt(); ok {
		_spec.SetField(commonconfig.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ccc.mutation.UpdatedAt(); ok {
		_spec.SetField(commonconfig.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ccc.mutation.DeleteAt(); ok {
		_spec.SetField(commonconfig.FieldDeleteAt, field.TypeTime, value)
		_node.DeleteAt = value
	}
	if value, ok := ccc.mutation.Status(); ok {
		_spec.SetField(commonconfig.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := ccc.mutation.CreatedBy(); ok {
		_spec.SetField(commonconfig.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := ccc.mutation.UpdatedBy(); ok {
		_spec.SetField(commonconfig.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := ccc.mutation.DeleteBy(); ok {
		_spec.SetField(commonconfig.FieldDeleteBy, field.TypeInt64, value)
		_node.DeleteBy = value
	}
	if value, ok := ccc.mutation.Remark(); ok {
		_spec.SetField(commonconfig.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := ccc.mutation.DelFlag(); ok {
		_spec.SetField(commonconfig.FieldDelFlag, field.TypeInt8, value)
		_node.DelFlag = value
	}
	if value, ok := ccc.mutation.ConfigName(); ok {
		_spec.SetField(commonconfig.FieldConfigName, field.TypeString, value)
		_node.ConfigName = value
	}
	if value, ok := ccc.mutation.ConfigKey(); ok {
		_spec.SetField(commonconfig.FieldConfigKey, field.TypeString, value)
		_node.ConfigKey = value
	}
	if value, ok := ccc.mutation.ConfigValue(); ok {
		_spec.SetField(commonconfig.FieldConfigValue, field.TypeString, value)
		_node.ConfigValue = value
	}
	if value, ok := ccc.mutation.ConfigType(); ok {
		_spec.SetField(commonconfig.FieldConfigType, field.TypeInt8, value)
		_node.ConfigType = value
	}
	return _node, _spec
}

// CommonConfigCreateBulk is the builder for creating many CommonConfig entities in bulk.
type CommonConfigCreateBulk struct {
	config
	err      error
	builders []*CommonConfigCreate
}

// Save creates the CommonConfig entities in the database.
func (cccb *CommonConfigCreateBulk) Save(ctx context.Context) ([]*CommonConfig, error) {
	if cccb.err != nil {
		return nil, cccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cccb.builders))
	nodes := make([]*CommonConfig, len(cccb.builders))
	mutators := make([]Mutator, len(cccb.builders))
	for i := range cccb.builders {
		func(i int, root context.Context) {
			builder := cccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommonConfigMutation)
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
					_, err = mutators[i+1].Mutate(root, cccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cccb *CommonConfigCreateBulk) SaveX(ctx context.Context) []*CommonConfig {
	v, err := cccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cccb *CommonConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := cccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cccb *CommonConfigCreateBulk) ExecX(ctx context.Context) {
	if err := cccb.Exec(ctx); err != nil {
		panic(err)
	}
}
