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
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysoperlog"
)

// SysOperLogUpdate is the builder for updating SysOperLog entities.
type SysOperLogUpdate struct {
	config
	hooks    []Hook
	mutation *SysOperLogMutation
}

// Where appends a list predicates to the SysOperLogUpdate builder.
func (solu *SysOperLogUpdate) Where(ps ...predicate.SysOperLog) *SysOperLogUpdate {
	solu.mutation.Where(ps...)
	return solu
}

// SetTitle sets the "title" field.
func (solu *SysOperLogUpdate) SetTitle(s string) *SysOperLogUpdate {
	solu.mutation.SetTitle(s)
	return solu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableTitle(s *string) *SysOperLogUpdate {
	if s != nil {
		solu.SetTitle(*s)
	}
	return solu
}

// ClearTitle clears the value of the "title" field.
func (solu *SysOperLogUpdate) ClearTitle() *SysOperLogUpdate {
	solu.mutation.ClearTitle()
	return solu
}

// SetBusinessType sets the "business_type" field.
func (solu *SysOperLogUpdate) SetBusinessType(i int) *SysOperLogUpdate {
	solu.mutation.ResetBusinessType()
	solu.mutation.SetBusinessType(i)
	return solu
}

// SetNillableBusinessType sets the "business_type" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableBusinessType(i *int) *SysOperLogUpdate {
	if i != nil {
		solu.SetBusinessType(*i)
	}
	return solu
}

// AddBusinessType adds i to the "business_type" field.
func (solu *SysOperLogUpdate) AddBusinessType(i int) *SysOperLogUpdate {
	solu.mutation.AddBusinessType(i)
	return solu
}

// ClearBusinessType clears the value of the "business_type" field.
func (solu *SysOperLogUpdate) ClearBusinessType() *SysOperLogUpdate {
	solu.mutation.ClearBusinessType()
	return solu
}

// SetMethod sets the "method" field.
func (solu *SysOperLogUpdate) SetMethod(s string) *SysOperLogUpdate {
	solu.mutation.SetMethod(s)
	return solu
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableMethod(s *string) *SysOperLogUpdate {
	if s != nil {
		solu.SetMethod(*s)
	}
	return solu
}

// ClearMethod clears the value of the "method" field.
func (solu *SysOperLogUpdate) ClearMethod() *SysOperLogUpdate {
	solu.mutation.ClearMethod()
	return solu
}

// SetRequestMethod sets the "request_method" field.
func (solu *SysOperLogUpdate) SetRequestMethod(s string) *SysOperLogUpdate {
	solu.mutation.SetRequestMethod(s)
	return solu
}

// SetNillableRequestMethod sets the "request_method" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableRequestMethod(s *string) *SysOperLogUpdate {
	if s != nil {
		solu.SetRequestMethod(*s)
	}
	return solu
}

// ClearRequestMethod clears the value of the "request_method" field.
func (solu *SysOperLogUpdate) ClearRequestMethod() *SysOperLogUpdate {
	solu.mutation.ClearRequestMethod()
	return solu
}

// SetOperType sets the "oper_type" field.
func (solu *SysOperLogUpdate) SetOperType(i int8) *SysOperLogUpdate {
	solu.mutation.ResetOperType()
	solu.mutation.SetOperType(i)
	return solu
}

// SetNillableOperType sets the "oper_type" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableOperType(i *int8) *SysOperLogUpdate {
	if i != nil {
		solu.SetOperType(*i)
	}
	return solu
}

// AddOperType adds i to the "oper_type" field.
func (solu *SysOperLogUpdate) AddOperType(i int8) *SysOperLogUpdate {
	solu.mutation.AddOperType(i)
	return solu
}

// ClearOperType clears the value of the "oper_type" field.
func (solu *SysOperLogUpdate) ClearOperType() *SysOperLogUpdate {
	solu.mutation.ClearOperType()
	return solu
}

// SetOperName sets the "oper_name" field.
func (solu *SysOperLogUpdate) SetOperName(s string) *SysOperLogUpdate {
	solu.mutation.SetOperName(s)
	return solu
}

// SetNillableOperName sets the "oper_name" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableOperName(s *string) *SysOperLogUpdate {
	if s != nil {
		solu.SetOperName(*s)
	}
	return solu
}

// ClearOperName clears the value of the "oper_name" field.
func (solu *SysOperLogUpdate) ClearOperName() *SysOperLogUpdate {
	solu.mutation.ClearOperName()
	return solu
}

// SetDeptName sets the "dept_name" field.
func (solu *SysOperLogUpdate) SetDeptName(s string) *SysOperLogUpdate {
	solu.mutation.SetDeptName(s)
	return solu
}

// SetNillableDeptName sets the "dept_name" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableDeptName(s *string) *SysOperLogUpdate {
	if s != nil {
		solu.SetDeptName(*s)
	}
	return solu
}

// ClearDeptName clears the value of the "dept_name" field.
func (solu *SysOperLogUpdate) ClearDeptName() *SysOperLogUpdate {
	solu.mutation.ClearDeptName()
	return solu
}

// SetOperURL sets the "oper_url" field.
func (solu *SysOperLogUpdate) SetOperURL(s string) *SysOperLogUpdate {
	solu.mutation.SetOperURL(s)
	return solu
}

// SetNillableOperURL sets the "oper_url" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableOperURL(s *string) *SysOperLogUpdate {
	if s != nil {
		solu.SetOperURL(*s)
	}
	return solu
}

// ClearOperURL clears the value of the "oper_url" field.
func (solu *SysOperLogUpdate) ClearOperURL() *SysOperLogUpdate {
	solu.mutation.ClearOperURL()
	return solu
}

// SetOperIP sets the "oper_ip" field.
func (solu *SysOperLogUpdate) SetOperIP(s string) *SysOperLogUpdate {
	solu.mutation.SetOperIP(s)
	return solu
}

// SetNillableOperIP sets the "oper_ip" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableOperIP(s *string) *SysOperLogUpdate {
	if s != nil {
		solu.SetOperIP(*s)
	}
	return solu
}

// ClearOperIP clears the value of the "oper_ip" field.
func (solu *SysOperLogUpdate) ClearOperIP() *SysOperLogUpdate {
	solu.mutation.ClearOperIP()
	return solu
}

// SetOperLocation sets the "oper_location" field.
func (solu *SysOperLogUpdate) SetOperLocation(s string) *SysOperLogUpdate {
	solu.mutation.SetOperLocation(s)
	return solu
}

// SetNillableOperLocation sets the "oper_location" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableOperLocation(s *string) *SysOperLogUpdate {
	if s != nil {
		solu.SetOperLocation(*s)
	}
	return solu
}

// ClearOperLocation clears the value of the "oper_location" field.
func (solu *SysOperLogUpdate) ClearOperLocation() *SysOperLogUpdate {
	solu.mutation.ClearOperLocation()
	return solu
}

// SetOperParam sets the "oper_param" field.
func (solu *SysOperLogUpdate) SetOperParam(s string) *SysOperLogUpdate {
	solu.mutation.SetOperParam(s)
	return solu
}

// SetNillableOperParam sets the "oper_param" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableOperParam(s *string) *SysOperLogUpdate {
	if s != nil {
		solu.SetOperParam(*s)
	}
	return solu
}

// ClearOperParam clears the value of the "oper_param" field.
func (solu *SysOperLogUpdate) ClearOperParam() *SysOperLogUpdate {
	solu.mutation.ClearOperParam()
	return solu
}

// SetErrorMsg sets the "error_msg" field.
func (solu *SysOperLogUpdate) SetErrorMsg(s string) *SysOperLogUpdate {
	solu.mutation.SetErrorMsg(s)
	return solu
}

// SetNillableErrorMsg sets the "error_msg" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableErrorMsg(s *string) *SysOperLogUpdate {
	if s != nil {
		solu.SetErrorMsg(*s)
	}
	return solu
}

// ClearErrorMsg clears the value of the "error_msg" field.
func (solu *SysOperLogUpdate) ClearErrorMsg() *SysOperLogUpdate {
	solu.mutation.ClearErrorMsg()
	return solu
}

// SetOperTime sets the "oper_time" field.
func (solu *SysOperLogUpdate) SetOperTime(t time.Time) *SysOperLogUpdate {
	solu.mutation.SetOperTime(t)
	return solu
}

// SetNillableOperTime sets the "oper_time" field if the given value is not nil.
func (solu *SysOperLogUpdate) SetNillableOperTime(t *time.Time) *SysOperLogUpdate {
	if t != nil {
		solu.SetOperTime(*t)
	}
	return solu
}

// ClearOperTime clears the value of the "oper_time" field.
func (solu *SysOperLogUpdate) ClearOperTime() *SysOperLogUpdate {
	solu.mutation.ClearOperTime()
	return solu
}

// Mutation returns the SysOperLogMutation object of the builder.
func (solu *SysOperLogUpdate) Mutation() *SysOperLogMutation {
	return solu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (solu *SysOperLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, solu.sqlSave, solu.mutation, solu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (solu *SysOperLogUpdate) SaveX(ctx context.Context) int {
	affected, err := solu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (solu *SysOperLogUpdate) Exec(ctx context.Context) error {
	_, err := solu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (solu *SysOperLogUpdate) ExecX(ctx context.Context) {
	if err := solu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (solu *SysOperLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sysoperlog.Table, sysoperlog.Columns, sqlgraph.NewFieldSpec(sysoperlog.FieldID, field.TypeInt64))
	if ps := solu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := solu.mutation.Title(); ok {
		_spec.SetField(sysoperlog.FieldTitle, field.TypeString, value)
	}
	if solu.mutation.TitleCleared() {
		_spec.ClearField(sysoperlog.FieldTitle, field.TypeString)
	}
	if value, ok := solu.mutation.BusinessType(); ok {
		_spec.SetField(sysoperlog.FieldBusinessType, field.TypeInt, value)
	}
	if value, ok := solu.mutation.AddedBusinessType(); ok {
		_spec.AddField(sysoperlog.FieldBusinessType, field.TypeInt, value)
	}
	if solu.mutation.BusinessTypeCleared() {
		_spec.ClearField(sysoperlog.FieldBusinessType, field.TypeInt)
	}
	if value, ok := solu.mutation.Method(); ok {
		_spec.SetField(sysoperlog.FieldMethod, field.TypeString, value)
	}
	if solu.mutation.MethodCleared() {
		_spec.ClearField(sysoperlog.FieldMethod, field.TypeString)
	}
	if value, ok := solu.mutation.RequestMethod(); ok {
		_spec.SetField(sysoperlog.FieldRequestMethod, field.TypeString, value)
	}
	if solu.mutation.RequestMethodCleared() {
		_spec.ClearField(sysoperlog.FieldRequestMethod, field.TypeString)
	}
	if value, ok := solu.mutation.OperType(); ok {
		_spec.SetField(sysoperlog.FieldOperType, field.TypeInt8, value)
	}
	if value, ok := solu.mutation.AddedOperType(); ok {
		_spec.AddField(sysoperlog.FieldOperType, field.TypeInt8, value)
	}
	if solu.mutation.OperTypeCleared() {
		_spec.ClearField(sysoperlog.FieldOperType, field.TypeInt8)
	}
	if value, ok := solu.mutation.OperName(); ok {
		_spec.SetField(sysoperlog.FieldOperName, field.TypeString, value)
	}
	if solu.mutation.OperNameCleared() {
		_spec.ClearField(sysoperlog.FieldOperName, field.TypeString)
	}
	if value, ok := solu.mutation.DeptName(); ok {
		_spec.SetField(sysoperlog.FieldDeptName, field.TypeString, value)
	}
	if solu.mutation.DeptNameCleared() {
		_spec.ClearField(sysoperlog.FieldDeptName, field.TypeString)
	}
	if value, ok := solu.mutation.OperURL(); ok {
		_spec.SetField(sysoperlog.FieldOperURL, field.TypeString, value)
	}
	if solu.mutation.OperURLCleared() {
		_spec.ClearField(sysoperlog.FieldOperURL, field.TypeString)
	}
	if value, ok := solu.mutation.OperIP(); ok {
		_spec.SetField(sysoperlog.FieldOperIP, field.TypeString, value)
	}
	if solu.mutation.OperIPCleared() {
		_spec.ClearField(sysoperlog.FieldOperIP, field.TypeString)
	}
	if value, ok := solu.mutation.OperLocation(); ok {
		_spec.SetField(sysoperlog.FieldOperLocation, field.TypeString, value)
	}
	if solu.mutation.OperLocationCleared() {
		_spec.ClearField(sysoperlog.FieldOperLocation, field.TypeString)
	}
	if value, ok := solu.mutation.OperParam(); ok {
		_spec.SetField(sysoperlog.FieldOperParam, field.TypeString, value)
	}
	if solu.mutation.OperParamCleared() {
		_spec.ClearField(sysoperlog.FieldOperParam, field.TypeString)
	}
	if value, ok := solu.mutation.ErrorMsg(); ok {
		_spec.SetField(sysoperlog.FieldErrorMsg, field.TypeString, value)
	}
	if solu.mutation.ErrorMsgCleared() {
		_spec.ClearField(sysoperlog.FieldErrorMsg, field.TypeString)
	}
	if value, ok := solu.mutation.OperTime(); ok {
		_spec.SetField(sysoperlog.FieldOperTime, field.TypeTime, value)
	}
	if solu.mutation.OperTimeCleared() {
		_spec.ClearField(sysoperlog.FieldOperTime, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, solu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysoperlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	solu.mutation.done = true
	return n, nil
}

// SysOperLogUpdateOne is the builder for updating a single SysOperLog entity.
type SysOperLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysOperLogMutation
}

// SetTitle sets the "title" field.
func (soluo *SysOperLogUpdateOne) SetTitle(s string) *SysOperLogUpdateOne {
	soluo.mutation.SetTitle(s)
	return soluo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableTitle(s *string) *SysOperLogUpdateOne {
	if s != nil {
		soluo.SetTitle(*s)
	}
	return soluo
}

// ClearTitle clears the value of the "title" field.
func (soluo *SysOperLogUpdateOne) ClearTitle() *SysOperLogUpdateOne {
	soluo.mutation.ClearTitle()
	return soluo
}

// SetBusinessType sets the "business_type" field.
func (soluo *SysOperLogUpdateOne) SetBusinessType(i int) *SysOperLogUpdateOne {
	soluo.mutation.ResetBusinessType()
	soluo.mutation.SetBusinessType(i)
	return soluo
}

// SetNillableBusinessType sets the "business_type" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableBusinessType(i *int) *SysOperLogUpdateOne {
	if i != nil {
		soluo.SetBusinessType(*i)
	}
	return soluo
}

// AddBusinessType adds i to the "business_type" field.
func (soluo *SysOperLogUpdateOne) AddBusinessType(i int) *SysOperLogUpdateOne {
	soluo.mutation.AddBusinessType(i)
	return soluo
}

// ClearBusinessType clears the value of the "business_type" field.
func (soluo *SysOperLogUpdateOne) ClearBusinessType() *SysOperLogUpdateOne {
	soluo.mutation.ClearBusinessType()
	return soluo
}

// SetMethod sets the "method" field.
func (soluo *SysOperLogUpdateOne) SetMethod(s string) *SysOperLogUpdateOne {
	soluo.mutation.SetMethod(s)
	return soluo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableMethod(s *string) *SysOperLogUpdateOne {
	if s != nil {
		soluo.SetMethod(*s)
	}
	return soluo
}

// ClearMethod clears the value of the "method" field.
func (soluo *SysOperLogUpdateOne) ClearMethod() *SysOperLogUpdateOne {
	soluo.mutation.ClearMethod()
	return soluo
}

// SetRequestMethod sets the "request_method" field.
func (soluo *SysOperLogUpdateOne) SetRequestMethod(s string) *SysOperLogUpdateOne {
	soluo.mutation.SetRequestMethod(s)
	return soluo
}

// SetNillableRequestMethod sets the "request_method" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableRequestMethod(s *string) *SysOperLogUpdateOne {
	if s != nil {
		soluo.SetRequestMethod(*s)
	}
	return soluo
}

// ClearRequestMethod clears the value of the "request_method" field.
func (soluo *SysOperLogUpdateOne) ClearRequestMethod() *SysOperLogUpdateOne {
	soluo.mutation.ClearRequestMethod()
	return soluo
}

// SetOperType sets the "oper_type" field.
func (soluo *SysOperLogUpdateOne) SetOperType(i int8) *SysOperLogUpdateOne {
	soluo.mutation.ResetOperType()
	soluo.mutation.SetOperType(i)
	return soluo
}

// SetNillableOperType sets the "oper_type" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableOperType(i *int8) *SysOperLogUpdateOne {
	if i != nil {
		soluo.SetOperType(*i)
	}
	return soluo
}

// AddOperType adds i to the "oper_type" field.
func (soluo *SysOperLogUpdateOne) AddOperType(i int8) *SysOperLogUpdateOne {
	soluo.mutation.AddOperType(i)
	return soluo
}

// ClearOperType clears the value of the "oper_type" field.
func (soluo *SysOperLogUpdateOne) ClearOperType() *SysOperLogUpdateOne {
	soluo.mutation.ClearOperType()
	return soluo
}

// SetOperName sets the "oper_name" field.
func (soluo *SysOperLogUpdateOne) SetOperName(s string) *SysOperLogUpdateOne {
	soluo.mutation.SetOperName(s)
	return soluo
}

// SetNillableOperName sets the "oper_name" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableOperName(s *string) *SysOperLogUpdateOne {
	if s != nil {
		soluo.SetOperName(*s)
	}
	return soluo
}

// ClearOperName clears the value of the "oper_name" field.
func (soluo *SysOperLogUpdateOne) ClearOperName() *SysOperLogUpdateOne {
	soluo.mutation.ClearOperName()
	return soluo
}

// SetDeptName sets the "dept_name" field.
func (soluo *SysOperLogUpdateOne) SetDeptName(s string) *SysOperLogUpdateOne {
	soluo.mutation.SetDeptName(s)
	return soluo
}

// SetNillableDeptName sets the "dept_name" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableDeptName(s *string) *SysOperLogUpdateOne {
	if s != nil {
		soluo.SetDeptName(*s)
	}
	return soluo
}

// ClearDeptName clears the value of the "dept_name" field.
func (soluo *SysOperLogUpdateOne) ClearDeptName() *SysOperLogUpdateOne {
	soluo.mutation.ClearDeptName()
	return soluo
}

// SetOperURL sets the "oper_url" field.
func (soluo *SysOperLogUpdateOne) SetOperURL(s string) *SysOperLogUpdateOne {
	soluo.mutation.SetOperURL(s)
	return soluo
}

// SetNillableOperURL sets the "oper_url" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableOperURL(s *string) *SysOperLogUpdateOne {
	if s != nil {
		soluo.SetOperURL(*s)
	}
	return soluo
}

// ClearOperURL clears the value of the "oper_url" field.
func (soluo *SysOperLogUpdateOne) ClearOperURL() *SysOperLogUpdateOne {
	soluo.mutation.ClearOperURL()
	return soluo
}

// SetOperIP sets the "oper_ip" field.
func (soluo *SysOperLogUpdateOne) SetOperIP(s string) *SysOperLogUpdateOne {
	soluo.mutation.SetOperIP(s)
	return soluo
}

// SetNillableOperIP sets the "oper_ip" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableOperIP(s *string) *SysOperLogUpdateOne {
	if s != nil {
		soluo.SetOperIP(*s)
	}
	return soluo
}

// ClearOperIP clears the value of the "oper_ip" field.
func (soluo *SysOperLogUpdateOne) ClearOperIP() *SysOperLogUpdateOne {
	soluo.mutation.ClearOperIP()
	return soluo
}

// SetOperLocation sets the "oper_location" field.
func (soluo *SysOperLogUpdateOne) SetOperLocation(s string) *SysOperLogUpdateOne {
	soluo.mutation.SetOperLocation(s)
	return soluo
}

// SetNillableOperLocation sets the "oper_location" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableOperLocation(s *string) *SysOperLogUpdateOne {
	if s != nil {
		soluo.SetOperLocation(*s)
	}
	return soluo
}

// ClearOperLocation clears the value of the "oper_location" field.
func (soluo *SysOperLogUpdateOne) ClearOperLocation() *SysOperLogUpdateOne {
	soluo.mutation.ClearOperLocation()
	return soluo
}

// SetOperParam sets the "oper_param" field.
func (soluo *SysOperLogUpdateOne) SetOperParam(s string) *SysOperLogUpdateOne {
	soluo.mutation.SetOperParam(s)
	return soluo
}

// SetNillableOperParam sets the "oper_param" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableOperParam(s *string) *SysOperLogUpdateOne {
	if s != nil {
		soluo.SetOperParam(*s)
	}
	return soluo
}

// ClearOperParam clears the value of the "oper_param" field.
func (soluo *SysOperLogUpdateOne) ClearOperParam() *SysOperLogUpdateOne {
	soluo.mutation.ClearOperParam()
	return soluo
}

// SetErrorMsg sets the "error_msg" field.
func (soluo *SysOperLogUpdateOne) SetErrorMsg(s string) *SysOperLogUpdateOne {
	soluo.mutation.SetErrorMsg(s)
	return soluo
}

// SetNillableErrorMsg sets the "error_msg" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableErrorMsg(s *string) *SysOperLogUpdateOne {
	if s != nil {
		soluo.SetErrorMsg(*s)
	}
	return soluo
}

// ClearErrorMsg clears the value of the "error_msg" field.
func (soluo *SysOperLogUpdateOne) ClearErrorMsg() *SysOperLogUpdateOne {
	soluo.mutation.ClearErrorMsg()
	return soluo
}

// SetOperTime sets the "oper_time" field.
func (soluo *SysOperLogUpdateOne) SetOperTime(t time.Time) *SysOperLogUpdateOne {
	soluo.mutation.SetOperTime(t)
	return soluo
}

// SetNillableOperTime sets the "oper_time" field if the given value is not nil.
func (soluo *SysOperLogUpdateOne) SetNillableOperTime(t *time.Time) *SysOperLogUpdateOne {
	if t != nil {
		soluo.SetOperTime(*t)
	}
	return soluo
}

// ClearOperTime clears the value of the "oper_time" field.
func (soluo *SysOperLogUpdateOne) ClearOperTime() *SysOperLogUpdateOne {
	soluo.mutation.ClearOperTime()
	return soluo
}

// Mutation returns the SysOperLogMutation object of the builder.
func (soluo *SysOperLogUpdateOne) Mutation() *SysOperLogMutation {
	return soluo.mutation
}

// Where appends a list predicates to the SysOperLogUpdate builder.
func (soluo *SysOperLogUpdateOne) Where(ps ...predicate.SysOperLog) *SysOperLogUpdateOne {
	soluo.mutation.Where(ps...)
	return soluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (soluo *SysOperLogUpdateOne) Select(field string, fields ...string) *SysOperLogUpdateOne {
	soluo.fields = append([]string{field}, fields...)
	return soluo
}

// Save executes the query and returns the updated SysOperLog entity.
func (soluo *SysOperLogUpdateOne) Save(ctx context.Context) (*SysOperLog, error) {
	return withHooks(ctx, soluo.sqlSave, soluo.mutation, soluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (soluo *SysOperLogUpdateOne) SaveX(ctx context.Context) *SysOperLog {
	node, err := soluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (soluo *SysOperLogUpdateOne) Exec(ctx context.Context) error {
	_, err := soluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (soluo *SysOperLogUpdateOne) ExecX(ctx context.Context) {
	if err := soluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (soluo *SysOperLogUpdateOne) sqlSave(ctx context.Context) (_node *SysOperLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(sysoperlog.Table, sysoperlog.Columns, sqlgraph.NewFieldSpec(sysoperlog.FieldID, field.TypeInt64))
	id, ok := soluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`codegen: missing "SysOperLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := soluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysoperlog.FieldID)
		for _, f := range fields {
			if !sysoperlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
			}
			if f != sysoperlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := soluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := soluo.mutation.Title(); ok {
		_spec.SetField(sysoperlog.FieldTitle, field.TypeString, value)
	}
	if soluo.mutation.TitleCleared() {
		_spec.ClearField(sysoperlog.FieldTitle, field.TypeString)
	}
	if value, ok := soluo.mutation.BusinessType(); ok {
		_spec.SetField(sysoperlog.FieldBusinessType, field.TypeInt, value)
	}
	if value, ok := soluo.mutation.AddedBusinessType(); ok {
		_spec.AddField(sysoperlog.FieldBusinessType, field.TypeInt, value)
	}
	if soluo.mutation.BusinessTypeCleared() {
		_spec.ClearField(sysoperlog.FieldBusinessType, field.TypeInt)
	}
	if value, ok := soluo.mutation.Method(); ok {
		_spec.SetField(sysoperlog.FieldMethod, field.TypeString, value)
	}
	if soluo.mutation.MethodCleared() {
		_spec.ClearField(sysoperlog.FieldMethod, field.TypeString)
	}
	if value, ok := soluo.mutation.RequestMethod(); ok {
		_spec.SetField(sysoperlog.FieldRequestMethod, field.TypeString, value)
	}
	if soluo.mutation.RequestMethodCleared() {
		_spec.ClearField(sysoperlog.FieldRequestMethod, field.TypeString)
	}
	if value, ok := soluo.mutation.OperType(); ok {
		_spec.SetField(sysoperlog.FieldOperType, field.TypeInt8, value)
	}
	if value, ok := soluo.mutation.AddedOperType(); ok {
		_spec.AddField(sysoperlog.FieldOperType, field.TypeInt8, value)
	}
	if soluo.mutation.OperTypeCleared() {
		_spec.ClearField(sysoperlog.FieldOperType, field.TypeInt8)
	}
	if value, ok := soluo.mutation.OperName(); ok {
		_spec.SetField(sysoperlog.FieldOperName, field.TypeString, value)
	}
	if soluo.mutation.OperNameCleared() {
		_spec.ClearField(sysoperlog.FieldOperName, field.TypeString)
	}
	if value, ok := soluo.mutation.DeptName(); ok {
		_spec.SetField(sysoperlog.FieldDeptName, field.TypeString, value)
	}
	if soluo.mutation.DeptNameCleared() {
		_spec.ClearField(sysoperlog.FieldDeptName, field.TypeString)
	}
	if value, ok := soluo.mutation.OperURL(); ok {
		_spec.SetField(sysoperlog.FieldOperURL, field.TypeString, value)
	}
	if soluo.mutation.OperURLCleared() {
		_spec.ClearField(sysoperlog.FieldOperURL, field.TypeString)
	}
	if value, ok := soluo.mutation.OperIP(); ok {
		_spec.SetField(sysoperlog.FieldOperIP, field.TypeString, value)
	}
	if soluo.mutation.OperIPCleared() {
		_spec.ClearField(sysoperlog.FieldOperIP, field.TypeString)
	}
	if value, ok := soluo.mutation.OperLocation(); ok {
		_spec.SetField(sysoperlog.FieldOperLocation, field.TypeString, value)
	}
	if soluo.mutation.OperLocationCleared() {
		_spec.ClearField(sysoperlog.FieldOperLocation, field.TypeString)
	}
	if value, ok := soluo.mutation.OperParam(); ok {
		_spec.SetField(sysoperlog.FieldOperParam, field.TypeString, value)
	}
	if soluo.mutation.OperParamCleared() {
		_spec.ClearField(sysoperlog.FieldOperParam, field.TypeString)
	}
	if value, ok := soluo.mutation.ErrorMsg(); ok {
		_spec.SetField(sysoperlog.FieldErrorMsg, field.TypeString, value)
	}
	if soluo.mutation.ErrorMsgCleared() {
		_spec.ClearField(sysoperlog.FieldErrorMsg, field.TypeString)
	}
	if value, ok := soluo.mutation.OperTime(); ok {
		_spec.SetField(sysoperlog.FieldOperTime, field.TypeTime, value)
	}
	if soluo.mutation.OperTimeCleared() {
		_spec.ClearField(sysoperlog.FieldOperTime, field.TypeTime)
	}
	_node = &SysOperLog{config: soluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, soluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysoperlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	soluo.mutation.done = true
	return _node, nil
}
