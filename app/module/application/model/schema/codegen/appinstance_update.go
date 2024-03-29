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
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/appinstance"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/apppackage"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/predicate"
)

// AppInstanceUpdate is the builder for updating AppInstance entities.
type AppInstanceUpdate struct {
	config
	hooks    []Hook
	mutation *AppInstanceMutation
}

// Where appends a list predicates to the AppInstanceUpdate builder.
func (aiu *AppInstanceUpdate) Where(ps ...predicate.AppInstance) *AppInstanceUpdate {
	aiu.mutation.Where(ps...)
	return aiu
}

// SetCreatedAt sets the "created_at" field.
func (aiu *AppInstanceUpdate) SetCreatedAt(t time.Time) *AppInstanceUpdate {
	aiu.mutation.SetCreatedAt(t)
	return aiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableCreatedAt(t *time.Time) *AppInstanceUpdate {
	if t != nil {
		aiu.SetCreatedAt(*t)
	}
	return aiu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (aiu *AppInstanceUpdate) ClearCreatedAt() *AppInstanceUpdate {
	aiu.mutation.ClearCreatedAt()
	return aiu
}

// SetUpdatedAt sets the "updated_at" field.
func (aiu *AppInstanceUpdate) SetUpdatedAt(t time.Time) *AppInstanceUpdate {
	aiu.mutation.SetUpdatedAt(t)
	return aiu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aiu *AppInstanceUpdate) ClearUpdatedAt() *AppInstanceUpdate {
	aiu.mutation.ClearUpdatedAt()
	return aiu
}

// SetDeleteAt sets the "delete_at" field.
func (aiu *AppInstanceUpdate) SetDeleteAt(t time.Time) *AppInstanceUpdate {
	aiu.mutation.SetDeleteAt(t)
	return aiu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableDeleteAt(t *time.Time) *AppInstanceUpdate {
	if t != nil {
		aiu.SetDeleteAt(*t)
	}
	return aiu
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (aiu *AppInstanceUpdate) ClearDeleteAt() *AppInstanceUpdate {
	aiu.mutation.ClearDeleteAt()
	return aiu
}

// SetStatus sets the "status" field.
func (aiu *AppInstanceUpdate) SetStatus(i int8) *AppInstanceUpdate {
	aiu.mutation.ResetStatus()
	aiu.mutation.SetStatus(i)
	return aiu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableStatus(i *int8) *AppInstanceUpdate {
	if i != nil {
		aiu.SetStatus(*i)
	}
	return aiu
}

// AddStatus adds i to the "status" field.
func (aiu *AppInstanceUpdate) AddStatus(i int8) *AppInstanceUpdate {
	aiu.mutation.AddStatus(i)
	return aiu
}

// SetRemark sets the "remark" field.
func (aiu *AppInstanceUpdate) SetRemark(s string) *AppInstanceUpdate {
	aiu.mutation.SetRemark(s)
	return aiu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableRemark(s *string) *AppInstanceUpdate {
	if s != nil {
		aiu.SetRemark(*s)
	}
	return aiu
}

// ClearRemark clears the value of the "remark" field.
func (aiu *AppInstanceUpdate) ClearRemark() *AppInstanceUpdate {
	aiu.mutation.ClearRemark()
	return aiu
}

// SetCreatedBy sets the "created_by" field.
func (aiu *AppInstanceUpdate) SetCreatedBy(i int64) *AppInstanceUpdate {
	aiu.mutation.ResetCreatedBy()
	aiu.mutation.SetCreatedBy(i)
	return aiu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableCreatedBy(i *int64) *AppInstanceUpdate {
	if i != nil {
		aiu.SetCreatedBy(*i)
	}
	return aiu
}

// AddCreatedBy adds i to the "created_by" field.
func (aiu *AppInstanceUpdate) AddCreatedBy(i int64) *AppInstanceUpdate {
	aiu.mutation.AddCreatedBy(i)
	return aiu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (aiu *AppInstanceUpdate) ClearCreatedBy() *AppInstanceUpdate {
	aiu.mutation.ClearCreatedBy()
	return aiu
}

// SetUpdatedBy sets the "updated_by" field.
func (aiu *AppInstanceUpdate) SetUpdatedBy(i int64) *AppInstanceUpdate {
	aiu.mutation.ResetUpdatedBy()
	aiu.mutation.SetUpdatedBy(i)
	return aiu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableUpdatedBy(i *int64) *AppInstanceUpdate {
	if i != nil {
		aiu.SetUpdatedBy(*i)
	}
	return aiu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (aiu *AppInstanceUpdate) AddUpdatedBy(i int64) *AppInstanceUpdate {
	aiu.mutation.AddUpdatedBy(i)
	return aiu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (aiu *AppInstanceUpdate) ClearUpdatedBy() *AppInstanceUpdate {
	aiu.mutation.ClearUpdatedBy()
	return aiu
}

// SetDeleteBy sets the "delete_by" field.
func (aiu *AppInstanceUpdate) SetDeleteBy(i int64) *AppInstanceUpdate {
	aiu.mutation.ResetDeleteBy()
	aiu.mutation.SetDeleteBy(i)
	return aiu
}

// SetNillableDeleteBy sets the "delete_by" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableDeleteBy(i *int64) *AppInstanceUpdate {
	if i != nil {
		aiu.SetDeleteBy(*i)
	}
	return aiu
}

// AddDeleteBy adds i to the "delete_by" field.
func (aiu *AppInstanceUpdate) AddDeleteBy(i int64) *AppInstanceUpdate {
	aiu.mutation.AddDeleteBy(i)
	return aiu
}

// ClearDeleteBy clears the value of the "delete_by" field.
func (aiu *AppInstanceUpdate) ClearDeleteBy() *AppInstanceUpdate {
	aiu.mutation.ClearDeleteBy()
	return aiu
}

// SetDelFlag sets the "del_flag" field.
func (aiu *AppInstanceUpdate) SetDelFlag(i int8) *AppInstanceUpdate {
	aiu.mutation.ResetDelFlag()
	aiu.mutation.SetDelFlag(i)
	return aiu
}

// SetNillableDelFlag sets the "del_flag" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableDelFlag(i *int8) *AppInstanceUpdate {
	if i != nil {
		aiu.SetDelFlag(*i)
	}
	return aiu
}

// AddDelFlag adds i to the "del_flag" field.
func (aiu *AppInstanceUpdate) AddDelFlag(i int8) *AppInstanceUpdate {
	aiu.mutation.AddDelFlag(i)
	return aiu
}

// SetInstanceName sets the "instance_name" field.
func (aiu *AppInstanceUpdate) SetInstanceName(s string) *AppInstanceUpdate {
	aiu.mutation.SetInstanceName(s)
	return aiu
}

// SetNillableInstanceName sets the "instance_name" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableInstanceName(s *string) *AppInstanceUpdate {
	if s != nil {
		aiu.SetInstanceName(*s)
	}
	return aiu
}

// SetInstanceCode sets the "instance_code" field.
func (aiu *AppInstanceUpdate) SetInstanceCode(s string) *AppInstanceUpdate {
	aiu.mutation.SetInstanceCode(s)
	return aiu
}

// SetNillableInstanceCode sets the "instance_code" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableInstanceCode(s *string) *AppInstanceUpdate {
	if s != nil {
		aiu.SetInstanceCode(*s)
	}
	return aiu
}

// SetInstancePackage sets the "instance_package" field.
func (aiu *AppInstanceUpdate) SetInstancePackage(i int64) *AppInstanceUpdate {
	aiu.mutation.ResetInstancePackage()
	aiu.mutation.SetInstancePackage(i)
	return aiu
}

// SetNillableInstancePackage sets the "instance_package" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableInstancePackage(i *int64) *AppInstanceUpdate {
	if i != nil {
		aiu.SetInstancePackage(*i)
	}
	return aiu
}

// AddInstancePackage adds i to the "instance_package" field.
func (aiu *AppInstanceUpdate) AddInstancePackage(i int64) *AppInstanceUpdate {
	aiu.mutation.AddInstancePackage(i)
	return aiu
}

// SetInstanceIcon sets the "instance_icon" field.
func (aiu *AppInstanceUpdate) SetInstanceIcon(s string) *AppInstanceUpdate {
	aiu.mutation.SetInstanceIcon(s)
	return aiu
}

// SetNillableInstanceIcon sets the "instance_icon" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableInstanceIcon(s *string) *AppInstanceUpdate {
	if s != nil {
		aiu.SetInstanceIcon(*s)
	}
	return aiu
}

// ClearInstanceIcon clears the value of the "instance_icon" field.
func (aiu *AppInstanceUpdate) ClearInstanceIcon() *AppInstanceUpdate {
	aiu.mutation.ClearInstanceIcon()
	return aiu
}

// SetInstanceAddress sets the "instance_address" field.
func (aiu *AppInstanceUpdate) SetInstanceAddress(s string) *AppInstanceUpdate {
	aiu.mutation.SetInstanceAddress(s)
	return aiu
}

// SetNillableInstanceAddress sets the "instance_address" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableInstanceAddress(s *string) *AppInstanceUpdate {
	if s != nil {
		aiu.SetInstanceAddress(*s)
	}
	return aiu
}

// ClearInstanceAddress clears the value of the "instance_address" field.
func (aiu *AppInstanceUpdate) ClearInstanceAddress() *AppInstanceUpdate {
	aiu.mutation.ClearInstanceAddress()
	return aiu
}

// SetInstanceType sets the "instance_type" field.
func (aiu *AppInstanceUpdate) SetInstanceType(i int8) *AppInstanceUpdate {
	aiu.mutation.ResetInstanceType()
	aiu.mutation.SetInstanceType(i)
	return aiu
}

// SetNillableInstanceType sets the "instance_type" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableInstanceType(i *int8) *AppInstanceUpdate {
	if i != nil {
		aiu.SetInstanceType(*i)
	}
	return aiu
}

// AddInstanceType adds i to the "instance_type" field.
func (aiu *AppInstanceUpdate) AddInstanceType(i int8) *AppInstanceUpdate {
	aiu.mutation.AddInstanceType(i)
	return aiu
}

// ClearInstanceType clears the value of the "instance_type" field.
func (aiu *AppInstanceUpdate) ClearInstanceType() *AppInstanceUpdate {
	aiu.mutation.ClearInstanceType()
	return aiu
}

// SetInstaller sets the "installer" field.
func (aiu *AppInstanceUpdate) SetInstaller(i int64) *AppInstanceUpdate {
	aiu.mutation.ResetInstaller()
	aiu.mutation.SetInstaller(i)
	return aiu
}

// SetNillableInstaller sets the "installer" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableInstaller(i *int64) *AppInstanceUpdate {
	if i != nil {
		aiu.SetInstaller(*i)
	}
	return aiu
}

// AddInstaller adds i to the "installer" field.
func (aiu *AppInstanceUpdate) AddInstaller(i int64) *AppInstanceUpdate {
	aiu.mutation.AddInstaller(i)
	return aiu
}

// ClearInstaller clears the value of the "installer" field.
func (aiu *AppInstanceUpdate) ClearInstaller() *AppInstanceUpdate {
	aiu.mutation.ClearInstaller()
	return aiu
}

// SetDesc sets the "desc" field.
func (aiu *AppInstanceUpdate) SetDesc(s string) *AppInstanceUpdate {
	aiu.mutation.SetDesc(s)
	return aiu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableDesc(s *string) *AppInstanceUpdate {
	if s != nil {
		aiu.SetDesc(*s)
	}
	return aiu
}

// ClearDesc clears the value of the "desc" field.
func (aiu *AppInstanceUpdate) ClearDesc() *AppInstanceUpdate {
	aiu.mutation.ClearDesc()
	return aiu
}

// SetInstallFromID sets the "installFrom" edge to the AppPackage entity by ID.
func (aiu *AppInstanceUpdate) SetInstallFromID(id int64) *AppInstanceUpdate {
	aiu.mutation.SetInstallFromID(id)
	return aiu
}

// SetNillableInstallFromID sets the "installFrom" edge to the AppPackage entity by ID if the given value is not nil.
func (aiu *AppInstanceUpdate) SetNillableInstallFromID(id *int64) *AppInstanceUpdate {
	if id != nil {
		aiu = aiu.SetInstallFromID(*id)
	}
	return aiu
}

// SetInstallFrom sets the "installFrom" edge to the AppPackage entity.
func (aiu *AppInstanceUpdate) SetInstallFrom(a *AppPackage) *AppInstanceUpdate {
	return aiu.SetInstallFromID(a.ID)
}

// Mutation returns the AppInstanceMutation object of the builder.
func (aiu *AppInstanceUpdate) Mutation() *AppInstanceMutation {
	return aiu.mutation
}

// ClearInstallFrom clears the "installFrom" edge to the AppPackage entity.
func (aiu *AppInstanceUpdate) ClearInstallFrom() *AppInstanceUpdate {
	aiu.mutation.ClearInstallFrom()
	return aiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aiu *AppInstanceUpdate) Save(ctx context.Context) (int, error) {
	if err := aiu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, aiu.sqlSave, aiu.mutation, aiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiu *AppInstanceUpdate) SaveX(ctx context.Context) int {
	affected, err := aiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aiu *AppInstanceUpdate) Exec(ctx context.Context) error {
	_, err := aiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiu *AppInstanceUpdate) ExecX(ctx context.Context) {
	if err := aiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aiu *AppInstanceUpdate) defaults() error {
	if _, ok := aiu.mutation.UpdatedAt(); !ok && !aiu.mutation.UpdatedAtCleared() {
		if appinstance.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("codegen: uninitialized appinstance.UpdateDefaultUpdatedAt (forgotten import codegen/runtime?)")
		}
		v := appinstance.UpdateDefaultUpdatedAt()
		aiu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (aiu *AppInstanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(appinstance.Table, appinstance.Columns, sqlgraph.NewFieldSpec(appinstance.FieldID, field.TypeInt64))
	if ps := aiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiu.mutation.CreatedAt(); ok {
		_spec.SetField(appinstance.FieldCreatedAt, field.TypeTime, value)
	}
	if aiu.mutation.CreatedAtCleared() {
		_spec.ClearField(appinstance.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := aiu.mutation.UpdatedAt(); ok {
		_spec.SetField(appinstance.FieldUpdatedAt, field.TypeTime, value)
	}
	if aiu.mutation.UpdatedAtCleared() {
		_spec.ClearField(appinstance.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := aiu.mutation.DeleteAt(); ok {
		_spec.SetField(appinstance.FieldDeleteAt, field.TypeTime, value)
	}
	if aiu.mutation.DeleteAtCleared() {
		_spec.ClearField(appinstance.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := aiu.mutation.Status(); ok {
		_spec.SetField(appinstance.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := aiu.mutation.AddedStatus(); ok {
		_spec.AddField(appinstance.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := aiu.mutation.Remark(); ok {
		_spec.SetField(appinstance.FieldRemark, field.TypeString, value)
	}
	if aiu.mutation.RemarkCleared() {
		_spec.ClearField(appinstance.FieldRemark, field.TypeString)
	}
	if value, ok := aiu.mutation.CreatedBy(); ok {
		_spec.SetField(appinstance.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(appinstance.FieldCreatedBy, field.TypeInt64, value)
	}
	if aiu.mutation.CreatedByCleared() {
		_spec.ClearField(appinstance.FieldCreatedBy, field.TypeInt64)
	}
	if value, ok := aiu.mutation.UpdatedBy(); ok {
		_spec.SetField(appinstance.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(appinstance.FieldUpdatedBy, field.TypeInt64, value)
	}
	if aiu.mutation.UpdatedByCleared() {
		_spec.ClearField(appinstance.FieldUpdatedBy, field.TypeInt64)
	}
	if value, ok := aiu.mutation.DeleteBy(); ok {
		_spec.SetField(appinstance.FieldDeleteBy, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.AddedDeleteBy(); ok {
		_spec.AddField(appinstance.FieldDeleteBy, field.TypeInt64, value)
	}
	if aiu.mutation.DeleteByCleared() {
		_spec.ClearField(appinstance.FieldDeleteBy, field.TypeInt64)
	}
	if value, ok := aiu.mutation.DelFlag(); ok {
		_spec.SetField(appinstance.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := aiu.mutation.AddedDelFlag(); ok {
		_spec.AddField(appinstance.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := aiu.mutation.InstanceName(); ok {
		_spec.SetField(appinstance.FieldInstanceName, field.TypeString, value)
	}
	if value, ok := aiu.mutation.InstanceCode(); ok {
		_spec.SetField(appinstance.FieldInstanceCode, field.TypeString, value)
	}
	if value, ok := aiu.mutation.InstancePackage(); ok {
		_spec.SetField(appinstance.FieldInstancePackage, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.AddedInstancePackage(); ok {
		_spec.AddField(appinstance.FieldInstancePackage, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.InstanceIcon(); ok {
		_spec.SetField(appinstance.FieldInstanceIcon, field.TypeString, value)
	}
	if aiu.mutation.InstanceIconCleared() {
		_spec.ClearField(appinstance.FieldInstanceIcon, field.TypeString)
	}
	if value, ok := aiu.mutation.InstanceAddress(); ok {
		_spec.SetField(appinstance.FieldInstanceAddress, field.TypeString, value)
	}
	if aiu.mutation.InstanceAddressCleared() {
		_spec.ClearField(appinstance.FieldInstanceAddress, field.TypeString)
	}
	if value, ok := aiu.mutation.InstanceType(); ok {
		_spec.SetField(appinstance.FieldInstanceType, field.TypeInt8, value)
	}
	if value, ok := aiu.mutation.AddedInstanceType(); ok {
		_spec.AddField(appinstance.FieldInstanceType, field.TypeInt8, value)
	}
	if aiu.mutation.InstanceTypeCleared() {
		_spec.ClearField(appinstance.FieldInstanceType, field.TypeInt8)
	}
	if value, ok := aiu.mutation.Installer(); ok {
		_spec.SetField(appinstance.FieldInstaller, field.TypeInt64, value)
	}
	if value, ok := aiu.mutation.AddedInstaller(); ok {
		_spec.AddField(appinstance.FieldInstaller, field.TypeInt64, value)
	}
	if aiu.mutation.InstallerCleared() {
		_spec.ClearField(appinstance.FieldInstaller, field.TypeInt64)
	}
	if value, ok := aiu.mutation.Desc(); ok {
		_spec.SetField(appinstance.FieldDesc, field.TypeString, value)
	}
	if aiu.mutation.DescCleared() {
		_spec.ClearField(appinstance.FieldDesc, field.TypeString)
	}
	if aiu.mutation.InstallFromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appinstance.InstallFromTable,
			Columns: []string{appinstance.InstallFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppackage.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.InstallFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appinstance.InstallFromTable,
			Columns: []string{appinstance.InstallFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppackage.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appinstance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aiu.mutation.done = true
	return n, nil
}

// AppInstanceUpdateOne is the builder for updating a single AppInstance entity.
type AppInstanceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppInstanceMutation
}

// SetCreatedAt sets the "created_at" field.
func (aiuo *AppInstanceUpdateOne) SetCreatedAt(t time.Time) *AppInstanceUpdateOne {
	aiuo.mutation.SetCreatedAt(t)
	return aiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableCreatedAt(t *time.Time) *AppInstanceUpdateOne {
	if t != nil {
		aiuo.SetCreatedAt(*t)
	}
	return aiuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (aiuo *AppInstanceUpdateOne) ClearCreatedAt() *AppInstanceUpdateOne {
	aiuo.mutation.ClearCreatedAt()
	return aiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (aiuo *AppInstanceUpdateOne) SetUpdatedAt(t time.Time) *AppInstanceUpdateOne {
	aiuo.mutation.SetUpdatedAt(t)
	return aiuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aiuo *AppInstanceUpdateOne) ClearUpdatedAt() *AppInstanceUpdateOne {
	aiuo.mutation.ClearUpdatedAt()
	return aiuo
}

// SetDeleteAt sets the "delete_at" field.
func (aiuo *AppInstanceUpdateOne) SetDeleteAt(t time.Time) *AppInstanceUpdateOne {
	aiuo.mutation.SetDeleteAt(t)
	return aiuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableDeleteAt(t *time.Time) *AppInstanceUpdateOne {
	if t != nil {
		aiuo.SetDeleteAt(*t)
	}
	return aiuo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (aiuo *AppInstanceUpdateOne) ClearDeleteAt() *AppInstanceUpdateOne {
	aiuo.mutation.ClearDeleteAt()
	return aiuo
}

// SetStatus sets the "status" field.
func (aiuo *AppInstanceUpdateOne) SetStatus(i int8) *AppInstanceUpdateOne {
	aiuo.mutation.ResetStatus()
	aiuo.mutation.SetStatus(i)
	return aiuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableStatus(i *int8) *AppInstanceUpdateOne {
	if i != nil {
		aiuo.SetStatus(*i)
	}
	return aiuo
}

// AddStatus adds i to the "status" field.
func (aiuo *AppInstanceUpdateOne) AddStatus(i int8) *AppInstanceUpdateOne {
	aiuo.mutation.AddStatus(i)
	return aiuo
}

// SetRemark sets the "remark" field.
func (aiuo *AppInstanceUpdateOne) SetRemark(s string) *AppInstanceUpdateOne {
	aiuo.mutation.SetRemark(s)
	return aiuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableRemark(s *string) *AppInstanceUpdateOne {
	if s != nil {
		aiuo.SetRemark(*s)
	}
	return aiuo
}

// ClearRemark clears the value of the "remark" field.
func (aiuo *AppInstanceUpdateOne) ClearRemark() *AppInstanceUpdateOne {
	aiuo.mutation.ClearRemark()
	return aiuo
}

// SetCreatedBy sets the "created_by" field.
func (aiuo *AppInstanceUpdateOne) SetCreatedBy(i int64) *AppInstanceUpdateOne {
	aiuo.mutation.ResetCreatedBy()
	aiuo.mutation.SetCreatedBy(i)
	return aiuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableCreatedBy(i *int64) *AppInstanceUpdateOne {
	if i != nil {
		aiuo.SetCreatedBy(*i)
	}
	return aiuo
}

// AddCreatedBy adds i to the "created_by" field.
func (aiuo *AppInstanceUpdateOne) AddCreatedBy(i int64) *AppInstanceUpdateOne {
	aiuo.mutation.AddCreatedBy(i)
	return aiuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (aiuo *AppInstanceUpdateOne) ClearCreatedBy() *AppInstanceUpdateOne {
	aiuo.mutation.ClearCreatedBy()
	return aiuo
}

// SetUpdatedBy sets the "updated_by" field.
func (aiuo *AppInstanceUpdateOne) SetUpdatedBy(i int64) *AppInstanceUpdateOne {
	aiuo.mutation.ResetUpdatedBy()
	aiuo.mutation.SetUpdatedBy(i)
	return aiuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableUpdatedBy(i *int64) *AppInstanceUpdateOne {
	if i != nil {
		aiuo.SetUpdatedBy(*i)
	}
	return aiuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (aiuo *AppInstanceUpdateOne) AddUpdatedBy(i int64) *AppInstanceUpdateOne {
	aiuo.mutation.AddUpdatedBy(i)
	return aiuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (aiuo *AppInstanceUpdateOne) ClearUpdatedBy() *AppInstanceUpdateOne {
	aiuo.mutation.ClearUpdatedBy()
	return aiuo
}

// SetDeleteBy sets the "delete_by" field.
func (aiuo *AppInstanceUpdateOne) SetDeleteBy(i int64) *AppInstanceUpdateOne {
	aiuo.mutation.ResetDeleteBy()
	aiuo.mutation.SetDeleteBy(i)
	return aiuo
}

// SetNillableDeleteBy sets the "delete_by" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableDeleteBy(i *int64) *AppInstanceUpdateOne {
	if i != nil {
		aiuo.SetDeleteBy(*i)
	}
	return aiuo
}

// AddDeleteBy adds i to the "delete_by" field.
func (aiuo *AppInstanceUpdateOne) AddDeleteBy(i int64) *AppInstanceUpdateOne {
	aiuo.mutation.AddDeleteBy(i)
	return aiuo
}

// ClearDeleteBy clears the value of the "delete_by" field.
func (aiuo *AppInstanceUpdateOne) ClearDeleteBy() *AppInstanceUpdateOne {
	aiuo.mutation.ClearDeleteBy()
	return aiuo
}

// SetDelFlag sets the "del_flag" field.
func (aiuo *AppInstanceUpdateOne) SetDelFlag(i int8) *AppInstanceUpdateOne {
	aiuo.mutation.ResetDelFlag()
	aiuo.mutation.SetDelFlag(i)
	return aiuo
}

// SetNillableDelFlag sets the "del_flag" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableDelFlag(i *int8) *AppInstanceUpdateOne {
	if i != nil {
		aiuo.SetDelFlag(*i)
	}
	return aiuo
}

// AddDelFlag adds i to the "del_flag" field.
func (aiuo *AppInstanceUpdateOne) AddDelFlag(i int8) *AppInstanceUpdateOne {
	aiuo.mutation.AddDelFlag(i)
	return aiuo
}

// SetInstanceName sets the "instance_name" field.
func (aiuo *AppInstanceUpdateOne) SetInstanceName(s string) *AppInstanceUpdateOne {
	aiuo.mutation.SetInstanceName(s)
	return aiuo
}

// SetNillableInstanceName sets the "instance_name" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableInstanceName(s *string) *AppInstanceUpdateOne {
	if s != nil {
		aiuo.SetInstanceName(*s)
	}
	return aiuo
}

// SetInstanceCode sets the "instance_code" field.
func (aiuo *AppInstanceUpdateOne) SetInstanceCode(s string) *AppInstanceUpdateOne {
	aiuo.mutation.SetInstanceCode(s)
	return aiuo
}

// SetNillableInstanceCode sets the "instance_code" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableInstanceCode(s *string) *AppInstanceUpdateOne {
	if s != nil {
		aiuo.SetInstanceCode(*s)
	}
	return aiuo
}

// SetInstancePackage sets the "instance_package" field.
func (aiuo *AppInstanceUpdateOne) SetInstancePackage(i int64) *AppInstanceUpdateOne {
	aiuo.mutation.ResetInstancePackage()
	aiuo.mutation.SetInstancePackage(i)
	return aiuo
}

// SetNillableInstancePackage sets the "instance_package" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableInstancePackage(i *int64) *AppInstanceUpdateOne {
	if i != nil {
		aiuo.SetInstancePackage(*i)
	}
	return aiuo
}

// AddInstancePackage adds i to the "instance_package" field.
func (aiuo *AppInstanceUpdateOne) AddInstancePackage(i int64) *AppInstanceUpdateOne {
	aiuo.mutation.AddInstancePackage(i)
	return aiuo
}

// SetInstanceIcon sets the "instance_icon" field.
func (aiuo *AppInstanceUpdateOne) SetInstanceIcon(s string) *AppInstanceUpdateOne {
	aiuo.mutation.SetInstanceIcon(s)
	return aiuo
}

// SetNillableInstanceIcon sets the "instance_icon" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableInstanceIcon(s *string) *AppInstanceUpdateOne {
	if s != nil {
		aiuo.SetInstanceIcon(*s)
	}
	return aiuo
}

// ClearInstanceIcon clears the value of the "instance_icon" field.
func (aiuo *AppInstanceUpdateOne) ClearInstanceIcon() *AppInstanceUpdateOne {
	aiuo.mutation.ClearInstanceIcon()
	return aiuo
}

// SetInstanceAddress sets the "instance_address" field.
func (aiuo *AppInstanceUpdateOne) SetInstanceAddress(s string) *AppInstanceUpdateOne {
	aiuo.mutation.SetInstanceAddress(s)
	return aiuo
}

// SetNillableInstanceAddress sets the "instance_address" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableInstanceAddress(s *string) *AppInstanceUpdateOne {
	if s != nil {
		aiuo.SetInstanceAddress(*s)
	}
	return aiuo
}

// ClearInstanceAddress clears the value of the "instance_address" field.
func (aiuo *AppInstanceUpdateOne) ClearInstanceAddress() *AppInstanceUpdateOne {
	aiuo.mutation.ClearInstanceAddress()
	return aiuo
}

// SetInstanceType sets the "instance_type" field.
func (aiuo *AppInstanceUpdateOne) SetInstanceType(i int8) *AppInstanceUpdateOne {
	aiuo.mutation.ResetInstanceType()
	aiuo.mutation.SetInstanceType(i)
	return aiuo
}

// SetNillableInstanceType sets the "instance_type" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableInstanceType(i *int8) *AppInstanceUpdateOne {
	if i != nil {
		aiuo.SetInstanceType(*i)
	}
	return aiuo
}

// AddInstanceType adds i to the "instance_type" field.
func (aiuo *AppInstanceUpdateOne) AddInstanceType(i int8) *AppInstanceUpdateOne {
	aiuo.mutation.AddInstanceType(i)
	return aiuo
}

// ClearInstanceType clears the value of the "instance_type" field.
func (aiuo *AppInstanceUpdateOne) ClearInstanceType() *AppInstanceUpdateOne {
	aiuo.mutation.ClearInstanceType()
	return aiuo
}

// SetInstaller sets the "installer" field.
func (aiuo *AppInstanceUpdateOne) SetInstaller(i int64) *AppInstanceUpdateOne {
	aiuo.mutation.ResetInstaller()
	aiuo.mutation.SetInstaller(i)
	return aiuo
}

// SetNillableInstaller sets the "installer" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableInstaller(i *int64) *AppInstanceUpdateOne {
	if i != nil {
		aiuo.SetInstaller(*i)
	}
	return aiuo
}

// AddInstaller adds i to the "installer" field.
func (aiuo *AppInstanceUpdateOne) AddInstaller(i int64) *AppInstanceUpdateOne {
	aiuo.mutation.AddInstaller(i)
	return aiuo
}

// ClearInstaller clears the value of the "installer" field.
func (aiuo *AppInstanceUpdateOne) ClearInstaller() *AppInstanceUpdateOne {
	aiuo.mutation.ClearInstaller()
	return aiuo
}

// SetDesc sets the "desc" field.
func (aiuo *AppInstanceUpdateOne) SetDesc(s string) *AppInstanceUpdateOne {
	aiuo.mutation.SetDesc(s)
	return aiuo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableDesc(s *string) *AppInstanceUpdateOne {
	if s != nil {
		aiuo.SetDesc(*s)
	}
	return aiuo
}

// ClearDesc clears the value of the "desc" field.
func (aiuo *AppInstanceUpdateOne) ClearDesc() *AppInstanceUpdateOne {
	aiuo.mutation.ClearDesc()
	return aiuo
}

// SetInstallFromID sets the "installFrom" edge to the AppPackage entity by ID.
func (aiuo *AppInstanceUpdateOne) SetInstallFromID(id int64) *AppInstanceUpdateOne {
	aiuo.mutation.SetInstallFromID(id)
	return aiuo
}

// SetNillableInstallFromID sets the "installFrom" edge to the AppPackage entity by ID if the given value is not nil.
func (aiuo *AppInstanceUpdateOne) SetNillableInstallFromID(id *int64) *AppInstanceUpdateOne {
	if id != nil {
		aiuo = aiuo.SetInstallFromID(*id)
	}
	return aiuo
}

// SetInstallFrom sets the "installFrom" edge to the AppPackage entity.
func (aiuo *AppInstanceUpdateOne) SetInstallFrom(a *AppPackage) *AppInstanceUpdateOne {
	return aiuo.SetInstallFromID(a.ID)
}

// Mutation returns the AppInstanceMutation object of the builder.
func (aiuo *AppInstanceUpdateOne) Mutation() *AppInstanceMutation {
	return aiuo.mutation
}

// ClearInstallFrom clears the "installFrom" edge to the AppPackage entity.
func (aiuo *AppInstanceUpdateOne) ClearInstallFrom() *AppInstanceUpdateOne {
	aiuo.mutation.ClearInstallFrom()
	return aiuo
}

// Where appends a list predicates to the AppInstanceUpdate builder.
func (aiuo *AppInstanceUpdateOne) Where(ps ...predicate.AppInstance) *AppInstanceUpdateOne {
	aiuo.mutation.Where(ps...)
	return aiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aiuo *AppInstanceUpdateOne) Select(field string, fields ...string) *AppInstanceUpdateOne {
	aiuo.fields = append([]string{field}, fields...)
	return aiuo
}

// Save executes the query and returns the updated AppInstance entity.
func (aiuo *AppInstanceUpdateOne) Save(ctx context.Context) (*AppInstance, error) {
	if err := aiuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, aiuo.sqlSave, aiuo.mutation, aiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiuo *AppInstanceUpdateOne) SaveX(ctx context.Context) *AppInstance {
	node, err := aiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aiuo *AppInstanceUpdateOne) Exec(ctx context.Context) error {
	_, err := aiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiuo *AppInstanceUpdateOne) ExecX(ctx context.Context) {
	if err := aiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aiuo *AppInstanceUpdateOne) defaults() error {
	if _, ok := aiuo.mutation.UpdatedAt(); !ok && !aiuo.mutation.UpdatedAtCleared() {
		if appinstance.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("codegen: uninitialized appinstance.UpdateDefaultUpdatedAt (forgotten import codegen/runtime?)")
		}
		v := appinstance.UpdateDefaultUpdatedAt()
		aiuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (aiuo *AppInstanceUpdateOne) sqlSave(ctx context.Context) (_node *AppInstance, err error) {
	_spec := sqlgraph.NewUpdateSpec(appinstance.Table, appinstance.Columns, sqlgraph.NewFieldSpec(appinstance.FieldID, field.TypeInt64))
	id, ok := aiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`codegen: missing "AppInstance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appinstance.FieldID)
		for _, f := range fields {
			if !appinstance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
			}
			if f != appinstance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiuo.mutation.CreatedAt(); ok {
		_spec.SetField(appinstance.FieldCreatedAt, field.TypeTime, value)
	}
	if aiuo.mutation.CreatedAtCleared() {
		_spec.ClearField(appinstance.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := aiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(appinstance.FieldUpdatedAt, field.TypeTime, value)
	}
	if aiuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(appinstance.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := aiuo.mutation.DeleteAt(); ok {
		_spec.SetField(appinstance.FieldDeleteAt, field.TypeTime, value)
	}
	if aiuo.mutation.DeleteAtCleared() {
		_spec.ClearField(appinstance.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := aiuo.mutation.Status(); ok {
		_spec.SetField(appinstance.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := aiuo.mutation.AddedStatus(); ok {
		_spec.AddField(appinstance.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := aiuo.mutation.Remark(); ok {
		_spec.SetField(appinstance.FieldRemark, field.TypeString, value)
	}
	if aiuo.mutation.RemarkCleared() {
		_spec.ClearField(appinstance.FieldRemark, field.TypeString)
	}
	if value, ok := aiuo.mutation.CreatedBy(); ok {
		_spec.SetField(appinstance.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(appinstance.FieldCreatedBy, field.TypeInt64, value)
	}
	if aiuo.mutation.CreatedByCleared() {
		_spec.ClearField(appinstance.FieldCreatedBy, field.TypeInt64)
	}
	if value, ok := aiuo.mutation.UpdatedBy(); ok {
		_spec.SetField(appinstance.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(appinstance.FieldUpdatedBy, field.TypeInt64, value)
	}
	if aiuo.mutation.UpdatedByCleared() {
		_spec.ClearField(appinstance.FieldUpdatedBy, field.TypeInt64)
	}
	if value, ok := aiuo.mutation.DeleteBy(); ok {
		_spec.SetField(appinstance.FieldDeleteBy, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.AddedDeleteBy(); ok {
		_spec.AddField(appinstance.FieldDeleteBy, field.TypeInt64, value)
	}
	if aiuo.mutation.DeleteByCleared() {
		_spec.ClearField(appinstance.FieldDeleteBy, field.TypeInt64)
	}
	if value, ok := aiuo.mutation.DelFlag(); ok {
		_spec.SetField(appinstance.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := aiuo.mutation.AddedDelFlag(); ok {
		_spec.AddField(appinstance.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := aiuo.mutation.InstanceName(); ok {
		_spec.SetField(appinstance.FieldInstanceName, field.TypeString, value)
	}
	if value, ok := aiuo.mutation.InstanceCode(); ok {
		_spec.SetField(appinstance.FieldInstanceCode, field.TypeString, value)
	}
	if value, ok := aiuo.mutation.InstancePackage(); ok {
		_spec.SetField(appinstance.FieldInstancePackage, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.AddedInstancePackage(); ok {
		_spec.AddField(appinstance.FieldInstancePackage, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.InstanceIcon(); ok {
		_spec.SetField(appinstance.FieldInstanceIcon, field.TypeString, value)
	}
	if aiuo.mutation.InstanceIconCleared() {
		_spec.ClearField(appinstance.FieldInstanceIcon, field.TypeString)
	}
	if value, ok := aiuo.mutation.InstanceAddress(); ok {
		_spec.SetField(appinstance.FieldInstanceAddress, field.TypeString, value)
	}
	if aiuo.mutation.InstanceAddressCleared() {
		_spec.ClearField(appinstance.FieldInstanceAddress, field.TypeString)
	}
	if value, ok := aiuo.mutation.InstanceType(); ok {
		_spec.SetField(appinstance.FieldInstanceType, field.TypeInt8, value)
	}
	if value, ok := aiuo.mutation.AddedInstanceType(); ok {
		_spec.AddField(appinstance.FieldInstanceType, field.TypeInt8, value)
	}
	if aiuo.mutation.InstanceTypeCleared() {
		_spec.ClearField(appinstance.FieldInstanceType, field.TypeInt8)
	}
	if value, ok := aiuo.mutation.Installer(); ok {
		_spec.SetField(appinstance.FieldInstaller, field.TypeInt64, value)
	}
	if value, ok := aiuo.mutation.AddedInstaller(); ok {
		_spec.AddField(appinstance.FieldInstaller, field.TypeInt64, value)
	}
	if aiuo.mutation.InstallerCleared() {
		_spec.ClearField(appinstance.FieldInstaller, field.TypeInt64)
	}
	if value, ok := aiuo.mutation.Desc(); ok {
		_spec.SetField(appinstance.FieldDesc, field.TypeString, value)
	}
	if aiuo.mutation.DescCleared() {
		_spec.ClearField(appinstance.FieldDesc, field.TypeString)
	}
	if aiuo.mutation.InstallFromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appinstance.InstallFromTable,
			Columns: []string{appinstance.InstallFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppackage.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.InstallFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appinstance.InstallFromTable,
			Columns: []string{appinstance.InstallFromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppackage.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AppInstance{config: aiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appinstance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aiuo.mutation.done = true
	return _node, nil
}
