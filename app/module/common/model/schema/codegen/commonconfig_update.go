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
	"github.com/wangxg422/XishangOS-backend/app/module/common/model/schema/codegen/commonconfig"
	"github.com/wangxg422/XishangOS-backend/app/module/common/model/schema/codegen/predicate"
)

// CommonConfigUpdate is the builder for updating CommonConfig entities.
type CommonConfigUpdate struct {
	config
	hooks    []Hook
	mutation *CommonConfigMutation
}

// Where appends a list predicates to the CommonConfigUpdate builder.
func (ccu *CommonConfigUpdate) Where(ps ...predicate.CommonConfig) *CommonConfigUpdate {
	ccu.mutation.Where(ps...)
	return ccu
}

// SetCreatedAt sets the "created_at" field.
func (ccu *CommonConfigUpdate) SetCreatedAt(t time.Time) *CommonConfigUpdate {
	ccu.mutation.SetCreatedAt(t)
	return ccu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableCreatedAt(t *time.Time) *CommonConfigUpdate {
	if t != nil {
		ccu.SetCreatedAt(*t)
	}
	return ccu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (ccu *CommonConfigUpdate) ClearCreatedAt() *CommonConfigUpdate {
	ccu.mutation.ClearCreatedAt()
	return ccu
}

// SetUpdatedAt sets the "updated_at" field.
func (ccu *CommonConfigUpdate) SetUpdatedAt(t time.Time) *CommonConfigUpdate {
	ccu.mutation.SetUpdatedAt(t)
	return ccu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ccu *CommonConfigUpdate) ClearUpdatedAt() *CommonConfigUpdate {
	ccu.mutation.ClearUpdatedAt()
	return ccu
}

// SetDeleteAt sets the "delete_at" field.
func (ccu *CommonConfigUpdate) SetDeleteAt(t time.Time) *CommonConfigUpdate {
	ccu.mutation.SetDeleteAt(t)
	return ccu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableDeleteAt(t *time.Time) *CommonConfigUpdate {
	if t != nil {
		ccu.SetDeleteAt(*t)
	}
	return ccu
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (ccu *CommonConfigUpdate) ClearDeleteAt() *CommonConfigUpdate {
	ccu.mutation.ClearDeleteAt()
	return ccu
}

// SetStatus sets the "status" field.
func (ccu *CommonConfigUpdate) SetStatus(i int8) *CommonConfigUpdate {
	ccu.mutation.ResetStatus()
	ccu.mutation.SetStatus(i)
	return ccu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableStatus(i *int8) *CommonConfigUpdate {
	if i != nil {
		ccu.SetStatus(*i)
	}
	return ccu
}

// AddStatus adds i to the "status" field.
func (ccu *CommonConfigUpdate) AddStatus(i int8) *CommonConfigUpdate {
	ccu.mutation.AddStatus(i)
	return ccu
}

// SetCreatedBy sets the "created_by" field.
func (ccu *CommonConfigUpdate) SetCreatedBy(i int64) *CommonConfigUpdate {
	ccu.mutation.ResetCreatedBy()
	ccu.mutation.SetCreatedBy(i)
	return ccu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableCreatedBy(i *int64) *CommonConfigUpdate {
	if i != nil {
		ccu.SetCreatedBy(*i)
	}
	return ccu
}

// AddCreatedBy adds i to the "created_by" field.
func (ccu *CommonConfigUpdate) AddCreatedBy(i int64) *CommonConfigUpdate {
	ccu.mutation.AddCreatedBy(i)
	return ccu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (ccu *CommonConfigUpdate) ClearCreatedBy() *CommonConfigUpdate {
	ccu.mutation.ClearCreatedBy()
	return ccu
}

// SetUpdatedBy sets the "updated_by" field.
func (ccu *CommonConfigUpdate) SetUpdatedBy(i int64) *CommonConfigUpdate {
	ccu.mutation.ResetUpdatedBy()
	ccu.mutation.SetUpdatedBy(i)
	return ccu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableUpdatedBy(i *int64) *CommonConfigUpdate {
	if i != nil {
		ccu.SetUpdatedBy(*i)
	}
	return ccu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ccu *CommonConfigUpdate) AddUpdatedBy(i int64) *CommonConfigUpdate {
	ccu.mutation.AddUpdatedBy(i)
	return ccu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ccu *CommonConfigUpdate) ClearUpdatedBy() *CommonConfigUpdate {
	ccu.mutation.ClearUpdatedBy()
	return ccu
}

// SetDeleteBy sets the "delete_by" field.
func (ccu *CommonConfigUpdate) SetDeleteBy(i int64) *CommonConfigUpdate {
	ccu.mutation.ResetDeleteBy()
	ccu.mutation.SetDeleteBy(i)
	return ccu
}

// SetNillableDeleteBy sets the "delete_by" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableDeleteBy(i *int64) *CommonConfigUpdate {
	if i != nil {
		ccu.SetDeleteBy(*i)
	}
	return ccu
}

// AddDeleteBy adds i to the "delete_by" field.
func (ccu *CommonConfigUpdate) AddDeleteBy(i int64) *CommonConfigUpdate {
	ccu.mutation.AddDeleteBy(i)
	return ccu
}

// ClearDeleteBy clears the value of the "delete_by" field.
func (ccu *CommonConfigUpdate) ClearDeleteBy() *CommonConfigUpdate {
	ccu.mutation.ClearDeleteBy()
	return ccu
}

// SetRemark sets the "remark" field.
func (ccu *CommonConfigUpdate) SetRemark(s string) *CommonConfigUpdate {
	ccu.mutation.SetRemark(s)
	return ccu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableRemark(s *string) *CommonConfigUpdate {
	if s != nil {
		ccu.SetRemark(*s)
	}
	return ccu
}

// ClearRemark clears the value of the "remark" field.
func (ccu *CommonConfigUpdate) ClearRemark() *CommonConfigUpdate {
	ccu.mutation.ClearRemark()
	return ccu
}

// SetDelFlag sets the "del_flag" field.
func (ccu *CommonConfigUpdate) SetDelFlag(i int8) *CommonConfigUpdate {
	ccu.mutation.ResetDelFlag()
	ccu.mutation.SetDelFlag(i)
	return ccu
}

// SetNillableDelFlag sets the "del_flag" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableDelFlag(i *int8) *CommonConfigUpdate {
	if i != nil {
		ccu.SetDelFlag(*i)
	}
	return ccu
}

// AddDelFlag adds i to the "del_flag" field.
func (ccu *CommonConfigUpdate) AddDelFlag(i int8) *CommonConfigUpdate {
	ccu.mutation.AddDelFlag(i)
	return ccu
}

// SetConfigName sets the "config_name" field.
func (ccu *CommonConfigUpdate) SetConfigName(s string) *CommonConfigUpdate {
	ccu.mutation.SetConfigName(s)
	return ccu
}

// SetNillableConfigName sets the "config_name" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableConfigName(s *string) *CommonConfigUpdate {
	if s != nil {
		ccu.SetConfigName(*s)
	}
	return ccu
}

// ClearConfigName clears the value of the "config_name" field.
func (ccu *CommonConfigUpdate) ClearConfigName() *CommonConfigUpdate {
	ccu.mutation.ClearConfigName()
	return ccu
}

// SetConfigKey sets the "config_key" field.
func (ccu *CommonConfigUpdate) SetConfigKey(s string) *CommonConfigUpdate {
	ccu.mutation.SetConfigKey(s)
	return ccu
}

// SetNillableConfigKey sets the "config_key" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableConfigKey(s *string) *CommonConfigUpdate {
	if s != nil {
		ccu.SetConfigKey(*s)
	}
	return ccu
}

// ClearConfigKey clears the value of the "config_key" field.
func (ccu *CommonConfigUpdate) ClearConfigKey() *CommonConfigUpdate {
	ccu.mutation.ClearConfigKey()
	return ccu
}

// SetConfigValue sets the "config_value" field.
func (ccu *CommonConfigUpdate) SetConfigValue(s string) *CommonConfigUpdate {
	ccu.mutation.SetConfigValue(s)
	return ccu
}

// SetNillableConfigValue sets the "config_value" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableConfigValue(s *string) *CommonConfigUpdate {
	if s != nil {
		ccu.SetConfigValue(*s)
	}
	return ccu
}

// ClearConfigValue clears the value of the "config_value" field.
func (ccu *CommonConfigUpdate) ClearConfigValue() *CommonConfigUpdate {
	ccu.mutation.ClearConfigValue()
	return ccu
}

// SetConfigType sets the "config_type" field.
func (ccu *CommonConfigUpdate) SetConfigType(i int8) *CommonConfigUpdate {
	ccu.mutation.ResetConfigType()
	ccu.mutation.SetConfigType(i)
	return ccu
}

// SetNillableConfigType sets the "config_type" field if the given value is not nil.
func (ccu *CommonConfigUpdate) SetNillableConfigType(i *int8) *CommonConfigUpdate {
	if i != nil {
		ccu.SetConfigType(*i)
	}
	return ccu
}

// AddConfigType adds i to the "config_type" field.
func (ccu *CommonConfigUpdate) AddConfigType(i int8) *CommonConfigUpdate {
	ccu.mutation.AddConfigType(i)
	return ccu
}

// ClearConfigType clears the value of the "config_type" field.
func (ccu *CommonConfigUpdate) ClearConfigType() *CommonConfigUpdate {
	ccu.mutation.ClearConfigType()
	return ccu
}

// Mutation returns the CommonConfigMutation object of the builder.
func (ccu *CommonConfigUpdate) Mutation() *CommonConfigMutation {
	return ccu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ccu *CommonConfigUpdate) Save(ctx context.Context) (int, error) {
	if err := ccu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ccu.sqlSave, ccu.mutation, ccu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccu *CommonConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := ccu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ccu *CommonConfigUpdate) Exec(ctx context.Context) error {
	_, err := ccu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccu *CommonConfigUpdate) ExecX(ctx context.Context) {
	if err := ccu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccu *CommonConfigUpdate) defaults() error {
	if _, ok := ccu.mutation.UpdatedAt(); !ok && !ccu.mutation.UpdatedAtCleared() {
		if commonconfig.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("codegen: uninitialized commonconfig.UpdateDefaultUpdatedAt (forgotten import codegen/runtime?)")
		}
		v := commonconfig.UpdateDefaultUpdatedAt()
		ccu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ccu *CommonConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(commonconfig.Table, commonconfig.Columns, sqlgraph.NewFieldSpec(commonconfig.FieldID, field.TypeInt64))
	if ps := ccu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccu.mutation.CreatedAt(); ok {
		_spec.SetField(commonconfig.FieldCreatedAt, field.TypeTime, value)
	}
	if ccu.mutation.CreatedAtCleared() {
		_spec.ClearField(commonconfig.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ccu.mutation.UpdatedAt(); ok {
		_spec.SetField(commonconfig.FieldUpdatedAt, field.TypeTime, value)
	}
	if ccu.mutation.UpdatedAtCleared() {
		_spec.ClearField(commonconfig.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ccu.mutation.DeleteAt(); ok {
		_spec.SetField(commonconfig.FieldDeleteAt, field.TypeTime, value)
	}
	if ccu.mutation.DeleteAtCleared() {
		_spec.ClearField(commonconfig.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := ccu.mutation.Status(); ok {
		_spec.SetField(commonconfig.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.AddedStatus(); ok {
		_spec.AddField(commonconfig.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.CreatedBy(); ok {
		_spec.SetField(commonconfig.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ccu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(commonconfig.FieldCreatedBy, field.TypeInt64, value)
	}
	if ccu.mutation.CreatedByCleared() {
		_spec.ClearField(commonconfig.FieldCreatedBy, field.TypeInt64)
	}
	if value, ok := ccu.mutation.UpdatedBy(); ok {
		_spec.SetField(commonconfig.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ccu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(commonconfig.FieldUpdatedBy, field.TypeInt64, value)
	}
	if ccu.mutation.UpdatedByCleared() {
		_spec.ClearField(commonconfig.FieldUpdatedBy, field.TypeInt64)
	}
	if value, ok := ccu.mutation.DeleteBy(); ok {
		_spec.SetField(commonconfig.FieldDeleteBy, field.TypeInt64, value)
	}
	if value, ok := ccu.mutation.AddedDeleteBy(); ok {
		_spec.AddField(commonconfig.FieldDeleteBy, field.TypeInt64, value)
	}
	if ccu.mutation.DeleteByCleared() {
		_spec.ClearField(commonconfig.FieldDeleteBy, field.TypeInt64)
	}
	if value, ok := ccu.mutation.Remark(); ok {
		_spec.SetField(commonconfig.FieldRemark, field.TypeString, value)
	}
	if ccu.mutation.RemarkCleared() {
		_spec.ClearField(commonconfig.FieldRemark, field.TypeString)
	}
	if value, ok := ccu.mutation.DelFlag(); ok {
		_spec.SetField(commonconfig.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.AddedDelFlag(); ok {
		_spec.AddField(commonconfig.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.ConfigName(); ok {
		_spec.SetField(commonconfig.FieldConfigName, field.TypeString, value)
	}
	if ccu.mutation.ConfigNameCleared() {
		_spec.ClearField(commonconfig.FieldConfigName, field.TypeString)
	}
	if value, ok := ccu.mutation.ConfigKey(); ok {
		_spec.SetField(commonconfig.FieldConfigKey, field.TypeString, value)
	}
	if ccu.mutation.ConfigKeyCleared() {
		_spec.ClearField(commonconfig.FieldConfigKey, field.TypeString)
	}
	if value, ok := ccu.mutation.ConfigValue(); ok {
		_spec.SetField(commonconfig.FieldConfigValue, field.TypeString, value)
	}
	if ccu.mutation.ConfigValueCleared() {
		_spec.ClearField(commonconfig.FieldConfigValue, field.TypeString)
	}
	if value, ok := ccu.mutation.ConfigType(); ok {
		_spec.SetField(commonconfig.FieldConfigType, field.TypeInt8, value)
	}
	if value, ok := ccu.mutation.AddedConfigType(); ok {
		_spec.AddField(commonconfig.FieldConfigType, field.TypeInt8, value)
	}
	if ccu.mutation.ConfigTypeCleared() {
		_spec.ClearField(commonconfig.FieldConfigType, field.TypeInt8)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ccu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commonconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ccu.mutation.done = true
	return n, nil
}

// CommonConfigUpdateOne is the builder for updating a single CommonConfig entity.
type CommonConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommonConfigMutation
}

// SetCreatedAt sets the "created_at" field.
func (ccuo *CommonConfigUpdateOne) SetCreatedAt(t time.Time) *CommonConfigUpdateOne {
	ccuo.mutation.SetCreatedAt(t)
	return ccuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableCreatedAt(t *time.Time) *CommonConfigUpdateOne {
	if t != nil {
		ccuo.SetCreatedAt(*t)
	}
	return ccuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (ccuo *CommonConfigUpdateOne) ClearCreatedAt() *CommonConfigUpdateOne {
	ccuo.mutation.ClearCreatedAt()
	return ccuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ccuo *CommonConfigUpdateOne) SetUpdatedAt(t time.Time) *CommonConfigUpdateOne {
	ccuo.mutation.SetUpdatedAt(t)
	return ccuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ccuo *CommonConfigUpdateOne) ClearUpdatedAt() *CommonConfigUpdateOne {
	ccuo.mutation.ClearUpdatedAt()
	return ccuo
}

// SetDeleteAt sets the "delete_at" field.
func (ccuo *CommonConfigUpdateOne) SetDeleteAt(t time.Time) *CommonConfigUpdateOne {
	ccuo.mutation.SetDeleteAt(t)
	return ccuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableDeleteAt(t *time.Time) *CommonConfigUpdateOne {
	if t != nil {
		ccuo.SetDeleteAt(*t)
	}
	return ccuo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (ccuo *CommonConfigUpdateOne) ClearDeleteAt() *CommonConfigUpdateOne {
	ccuo.mutation.ClearDeleteAt()
	return ccuo
}

// SetStatus sets the "status" field.
func (ccuo *CommonConfigUpdateOne) SetStatus(i int8) *CommonConfigUpdateOne {
	ccuo.mutation.ResetStatus()
	ccuo.mutation.SetStatus(i)
	return ccuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableStatus(i *int8) *CommonConfigUpdateOne {
	if i != nil {
		ccuo.SetStatus(*i)
	}
	return ccuo
}

// AddStatus adds i to the "status" field.
func (ccuo *CommonConfigUpdateOne) AddStatus(i int8) *CommonConfigUpdateOne {
	ccuo.mutation.AddStatus(i)
	return ccuo
}

// SetCreatedBy sets the "created_by" field.
func (ccuo *CommonConfigUpdateOne) SetCreatedBy(i int64) *CommonConfigUpdateOne {
	ccuo.mutation.ResetCreatedBy()
	ccuo.mutation.SetCreatedBy(i)
	return ccuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableCreatedBy(i *int64) *CommonConfigUpdateOne {
	if i != nil {
		ccuo.SetCreatedBy(*i)
	}
	return ccuo
}

// AddCreatedBy adds i to the "created_by" field.
func (ccuo *CommonConfigUpdateOne) AddCreatedBy(i int64) *CommonConfigUpdateOne {
	ccuo.mutation.AddCreatedBy(i)
	return ccuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (ccuo *CommonConfigUpdateOne) ClearCreatedBy() *CommonConfigUpdateOne {
	ccuo.mutation.ClearCreatedBy()
	return ccuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ccuo *CommonConfigUpdateOne) SetUpdatedBy(i int64) *CommonConfigUpdateOne {
	ccuo.mutation.ResetUpdatedBy()
	ccuo.mutation.SetUpdatedBy(i)
	return ccuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableUpdatedBy(i *int64) *CommonConfigUpdateOne {
	if i != nil {
		ccuo.SetUpdatedBy(*i)
	}
	return ccuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ccuo *CommonConfigUpdateOne) AddUpdatedBy(i int64) *CommonConfigUpdateOne {
	ccuo.mutation.AddUpdatedBy(i)
	return ccuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ccuo *CommonConfigUpdateOne) ClearUpdatedBy() *CommonConfigUpdateOne {
	ccuo.mutation.ClearUpdatedBy()
	return ccuo
}

// SetDeleteBy sets the "delete_by" field.
func (ccuo *CommonConfigUpdateOne) SetDeleteBy(i int64) *CommonConfigUpdateOne {
	ccuo.mutation.ResetDeleteBy()
	ccuo.mutation.SetDeleteBy(i)
	return ccuo
}

// SetNillableDeleteBy sets the "delete_by" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableDeleteBy(i *int64) *CommonConfigUpdateOne {
	if i != nil {
		ccuo.SetDeleteBy(*i)
	}
	return ccuo
}

// AddDeleteBy adds i to the "delete_by" field.
func (ccuo *CommonConfigUpdateOne) AddDeleteBy(i int64) *CommonConfigUpdateOne {
	ccuo.mutation.AddDeleteBy(i)
	return ccuo
}

// ClearDeleteBy clears the value of the "delete_by" field.
func (ccuo *CommonConfigUpdateOne) ClearDeleteBy() *CommonConfigUpdateOne {
	ccuo.mutation.ClearDeleteBy()
	return ccuo
}

// SetRemark sets the "remark" field.
func (ccuo *CommonConfigUpdateOne) SetRemark(s string) *CommonConfigUpdateOne {
	ccuo.mutation.SetRemark(s)
	return ccuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableRemark(s *string) *CommonConfigUpdateOne {
	if s != nil {
		ccuo.SetRemark(*s)
	}
	return ccuo
}

// ClearRemark clears the value of the "remark" field.
func (ccuo *CommonConfigUpdateOne) ClearRemark() *CommonConfigUpdateOne {
	ccuo.mutation.ClearRemark()
	return ccuo
}

// SetDelFlag sets the "del_flag" field.
func (ccuo *CommonConfigUpdateOne) SetDelFlag(i int8) *CommonConfigUpdateOne {
	ccuo.mutation.ResetDelFlag()
	ccuo.mutation.SetDelFlag(i)
	return ccuo
}

// SetNillableDelFlag sets the "del_flag" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableDelFlag(i *int8) *CommonConfigUpdateOne {
	if i != nil {
		ccuo.SetDelFlag(*i)
	}
	return ccuo
}

// AddDelFlag adds i to the "del_flag" field.
func (ccuo *CommonConfigUpdateOne) AddDelFlag(i int8) *CommonConfigUpdateOne {
	ccuo.mutation.AddDelFlag(i)
	return ccuo
}

// SetConfigName sets the "config_name" field.
func (ccuo *CommonConfigUpdateOne) SetConfigName(s string) *CommonConfigUpdateOne {
	ccuo.mutation.SetConfigName(s)
	return ccuo
}

// SetNillableConfigName sets the "config_name" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableConfigName(s *string) *CommonConfigUpdateOne {
	if s != nil {
		ccuo.SetConfigName(*s)
	}
	return ccuo
}

// ClearConfigName clears the value of the "config_name" field.
func (ccuo *CommonConfigUpdateOne) ClearConfigName() *CommonConfigUpdateOne {
	ccuo.mutation.ClearConfigName()
	return ccuo
}

// SetConfigKey sets the "config_key" field.
func (ccuo *CommonConfigUpdateOne) SetConfigKey(s string) *CommonConfigUpdateOne {
	ccuo.mutation.SetConfigKey(s)
	return ccuo
}

// SetNillableConfigKey sets the "config_key" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableConfigKey(s *string) *CommonConfigUpdateOne {
	if s != nil {
		ccuo.SetConfigKey(*s)
	}
	return ccuo
}

// ClearConfigKey clears the value of the "config_key" field.
func (ccuo *CommonConfigUpdateOne) ClearConfigKey() *CommonConfigUpdateOne {
	ccuo.mutation.ClearConfigKey()
	return ccuo
}

// SetConfigValue sets the "config_value" field.
func (ccuo *CommonConfigUpdateOne) SetConfigValue(s string) *CommonConfigUpdateOne {
	ccuo.mutation.SetConfigValue(s)
	return ccuo
}

// SetNillableConfigValue sets the "config_value" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableConfigValue(s *string) *CommonConfigUpdateOne {
	if s != nil {
		ccuo.SetConfigValue(*s)
	}
	return ccuo
}

// ClearConfigValue clears the value of the "config_value" field.
func (ccuo *CommonConfigUpdateOne) ClearConfigValue() *CommonConfigUpdateOne {
	ccuo.mutation.ClearConfigValue()
	return ccuo
}

// SetConfigType sets the "config_type" field.
func (ccuo *CommonConfigUpdateOne) SetConfigType(i int8) *CommonConfigUpdateOne {
	ccuo.mutation.ResetConfigType()
	ccuo.mutation.SetConfigType(i)
	return ccuo
}

// SetNillableConfigType sets the "config_type" field if the given value is not nil.
func (ccuo *CommonConfigUpdateOne) SetNillableConfigType(i *int8) *CommonConfigUpdateOne {
	if i != nil {
		ccuo.SetConfigType(*i)
	}
	return ccuo
}

// AddConfigType adds i to the "config_type" field.
func (ccuo *CommonConfigUpdateOne) AddConfigType(i int8) *CommonConfigUpdateOne {
	ccuo.mutation.AddConfigType(i)
	return ccuo
}

// ClearConfigType clears the value of the "config_type" field.
func (ccuo *CommonConfigUpdateOne) ClearConfigType() *CommonConfigUpdateOne {
	ccuo.mutation.ClearConfigType()
	return ccuo
}

// Mutation returns the CommonConfigMutation object of the builder.
func (ccuo *CommonConfigUpdateOne) Mutation() *CommonConfigMutation {
	return ccuo.mutation
}

// Where appends a list predicates to the CommonConfigUpdate builder.
func (ccuo *CommonConfigUpdateOne) Where(ps ...predicate.CommonConfig) *CommonConfigUpdateOne {
	ccuo.mutation.Where(ps...)
	return ccuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ccuo *CommonConfigUpdateOne) Select(field string, fields ...string) *CommonConfigUpdateOne {
	ccuo.fields = append([]string{field}, fields...)
	return ccuo
}

// Save executes the query and returns the updated CommonConfig entity.
func (ccuo *CommonConfigUpdateOne) Save(ctx context.Context) (*CommonConfig, error) {
	if err := ccuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ccuo.sqlSave, ccuo.mutation, ccuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ccuo *CommonConfigUpdateOne) SaveX(ctx context.Context) *CommonConfig {
	node, err := ccuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ccuo *CommonConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := ccuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccuo *CommonConfigUpdateOne) ExecX(ctx context.Context) {
	if err := ccuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ccuo *CommonConfigUpdateOne) defaults() error {
	if _, ok := ccuo.mutation.UpdatedAt(); !ok && !ccuo.mutation.UpdatedAtCleared() {
		if commonconfig.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("codegen: uninitialized commonconfig.UpdateDefaultUpdatedAt (forgotten import codegen/runtime?)")
		}
		v := commonconfig.UpdateDefaultUpdatedAt()
		ccuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ccuo *CommonConfigUpdateOne) sqlSave(ctx context.Context) (_node *CommonConfig, err error) {
	_spec := sqlgraph.NewUpdateSpec(commonconfig.Table, commonconfig.Columns, sqlgraph.NewFieldSpec(commonconfig.FieldID, field.TypeInt64))
	id, ok := ccuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`codegen: missing "CommonConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ccuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commonconfig.FieldID)
		for _, f := range fields {
			if !commonconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
			}
			if f != commonconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ccuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ccuo.mutation.CreatedAt(); ok {
		_spec.SetField(commonconfig.FieldCreatedAt, field.TypeTime, value)
	}
	if ccuo.mutation.CreatedAtCleared() {
		_spec.ClearField(commonconfig.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ccuo.mutation.UpdatedAt(); ok {
		_spec.SetField(commonconfig.FieldUpdatedAt, field.TypeTime, value)
	}
	if ccuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(commonconfig.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ccuo.mutation.DeleteAt(); ok {
		_spec.SetField(commonconfig.FieldDeleteAt, field.TypeTime, value)
	}
	if ccuo.mutation.DeleteAtCleared() {
		_spec.ClearField(commonconfig.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := ccuo.mutation.Status(); ok {
		_spec.SetField(commonconfig.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.AddedStatus(); ok {
		_spec.AddField(commonconfig.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.CreatedBy(); ok {
		_spec.SetField(commonconfig.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ccuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(commonconfig.FieldCreatedBy, field.TypeInt64, value)
	}
	if ccuo.mutation.CreatedByCleared() {
		_spec.ClearField(commonconfig.FieldCreatedBy, field.TypeInt64)
	}
	if value, ok := ccuo.mutation.UpdatedBy(); ok {
		_spec.SetField(commonconfig.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ccuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(commonconfig.FieldUpdatedBy, field.TypeInt64, value)
	}
	if ccuo.mutation.UpdatedByCleared() {
		_spec.ClearField(commonconfig.FieldUpdatedBy, field.TypeInt64)
	}
	if value, ok := ccuo.mutation.DeleteBy(); ok {
		_spec.SetField(commonconfig.FieldDeleteBy, field.TypeInt64, value)
	}
	if value, ok := ccuo.mutation.AddedDeleteBy(); ok {
		_spec.AddField(commonconfig.FieldDeleteBy, field.TypeInt64, value)
	}
	if ccuo.mutation.DeleteByCleared() {
		_spec.ClearField(commonconfig.FieldDeleteBy, field.TypeInt64)
	}
	if value, ok := ccuo.mutation.Remark(); ok {
		_spec.SetField(commonconfig.FieldRemark, field.TypeString, value)
	}
	if ccuo.mutation.RemarkCleared() {
		_spec.ClearField(commonconfig.FieldRemark, field.TypeString)
	}
	if value, ok := ccuo.mutation.DelFlag(); ok {
		_spec.SetField(commonconfig.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.AddedDelFlag(); ok {
		_spec.AddField(commonconfig.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.ConfigName(); ok {
		_spec.SetField(commonconfig.FieldConfigName, field.TypeString, value)
	}
	if ccuo.mutation.ConfigNameCleared() {
		_spec.ClearField(commonconfig.FieldConfigName, field.TypeString)
	}
	if value, ok := ccuo.mutation.ConfigKey(); ok {
		_spec.SetField(commonconfig.FieldConfigKey, field.TypeString, value)
	}
	if ccuo.mutation.ConfigKeyCleared() {
		_spec.ClearField(commonconfig.FieldConfigKey, field.TypeString)
	}
	if value, ok := ccuo.mutation.ConfigValue(); ok {
		_spec.SetField(commonconfig.FieldConfigValue, field.TypeString, value)
	}
	if ccuo.mutation.ConfigValueCleared() {
		_spec.ClearField(commonconfig.FieldConfigValue, field.TypeString)
	}
	if value, ok := ccuo.mutation.ConfigType(); ok {
		_spec.SetField(commonconfig.FieldConfigType, field.TypeInt8, value)
	}
	if value, ok := ccuo.mutation.AddedConfigType(); ok {
		_spec.AddField(commonconfig.FieldConfigType, field.TypeInt8, value)
	}
	if ccuo.mutation.ConfigTypeCleared() {
		_spec.ClearField(commonconfig.FieldConfigType, field.TypeInt8)
	}
	_node = &CommonConfig{config: ccuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ccuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commonconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ccuo.mutation.done = true
	return _node, nil
}
