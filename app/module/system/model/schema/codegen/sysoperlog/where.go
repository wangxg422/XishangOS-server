// Code generated by ent, DO NOT EDIT.

package sysoperlog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldTitle, v))
}

// BusinessType applies equality check predicate on the "business_type" field. It's identical to BusinessTypeEQ.
func BusinessType(v int) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldBusinessType, v))
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldMethod, v))
}

// RequestMethod applies equality check predicate on the "request_method" field. It's identical to RequestMethodEQ.
func RequestMethod(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldRequestMethod, v))
}

// OperType applies equality check predicate on the "oper_type" field. It's identical to OperTypeEQ.
func OperType(v int8) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperType, v))
}

// OperName applies equality check predicate on the "oper_name" field. It's identical to OperNameEQ.
func OperName(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperName, v))
}

// DeptName applies equality check predicate on the "dept_name" field. It's identical to DeptNameEQ.
func DeptName(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldDeptName, v))
}

// OperURL applies equality check predicate on the "oper_url" field. It's identical to OperURLEQ.
func OperURL(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperURL, v))
}

// OperIP applies equality check predicate on the "oper_ip" field. It's identical to OperIPEQ.
func OperIP(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperIP, v))
}

// OperLocation applies equality check predicate on the "oper_location" field. It's identical to OperLocationEQ.
func OperLocation(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperLocation, v))
}

// OperParam applies equality check predicate on the "oper_param" field. It's identical to OperParamEQ.
func OperParam(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperParam, v))
}

// ErrorMsg applies equality check predicate on the "error_msg" field. It's identical to ErrorMsgEQ.
func ErrorMsg(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldErrorMsg, v))
}

// OperTime applies equality check predicate on the "oper_time" field. It's identical to OperTimeEQ.
func OperTime(v time.Time) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperTime, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldTitle))
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldTitle))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContainsFold(FieldTitle, v))
}

// BusinessTypeEQ applies the EQ predicate on the "business_type" field.
func BusinessTypeEQ(v int) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldBusinessType, v))
}

// BusinessTypeNEQ applies the NEQ predicate on the "business_type" field.
func BusinessTypeNEQ(v int) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldBusinessType, v))
}

// BusinessTypeIn applies the In predicate on the "business_type" field.
func BusinessTypeIn(vs ...int) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldBusinessType, vs...))
}

// BusinessTypeNotIn applies the NotIn predicate on the "business_type" field.
func BusinessTypeNotIn(vs ...int) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldBusinessType, vs...))
}

// BusinessTypeGT applies the GT predicate on the "business_type" field.
func BusinessTypeGT(v int) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldBusinessType, v))
}

// BusinessTypeGTE applies the GTE predicate on the "business_type" field.
func BusinessTypeGTE(v int) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldBusinessType, v))
}

// BusinessTypeLT applies the LT predicate on the "business_type" field.
func BusinessTypeLT(v int) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldBusinessType, v))
}

// BusinessTypeLTE applies the LTE predicate on the "business_type" field.
func BusinessTypeLTE(v int) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldBusinessType, v))
}

// BusinessTypeIsNil applies the IsNil predicate on the "business_type" field.
func BusinessTypeIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldBusinessType))
}

// BusinessTypeNotNil applies the NotNil predicate on the "business_type" field.
func BusinessTypeNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldBusinessType))
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldMethod, v))
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldMethod, v))
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldMethod, vs...))
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldMethod, vs...))
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldMethod, v))
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldMethod, v))
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldMethod, v))
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldMethod, v))
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContains(FieldMethod, v))
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasPrefix(FieldMethod, v))
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasSuffix(FieldMethod, v))
}

// MethodIsNil applies the IsNil predicate on the "method" field.
func MethodIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldMethod))
}

// MethodNotNil applies the NotNil predicate on the "method" field.
func MethodNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldMethod))
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEqualFold(FieldMethod, v))
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContainsFold(FieldMethod, v))
}

// RequestMethodEQ applies the EQ predicate on the "request_method" field.
func RequestMethodEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldRequestMethod, v))
}

// RequestMethodNEQ applies the NEQ predicate on the "request_method" field.
func RequestMethodNEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldRequestMethod, v))
}

// RequestMethodIn applies the In predicate on the "request_method" field.
func RequestMethodIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldRequestMethod, vs...))
}

// RequestMethodNotIn applies the NotIn predicate on the "request_method" field.
func RequestMethodNotIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldRequestMethod, vs...))
}

// RequestMethodGT applies the GT predicate on the "request_method" field.
func RequestMethodGT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldRequestMethod, v))
}

// RequestMethodGTE applies the GTE predicate on the "request_method" field.
func RequestMethodGTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldRequestMethod, v))
}

// RequestMethodLT applies the LT predicate on the "request_method" field.
func RequestMethodLT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldRequestMethod, v))
}

// RequestMethodLTE applies the LTE predicate on the "request_method" field.
func RequestMethodLTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldRequestMethod, v))
}

// RequestMethodContains applies the Contains predicate on the "request_method" field.
func RequestMethodContains(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContains(FieldRequestMethod, v))
}

// RequestMethodHasPrefix applies the HasPrefix predicate on the "request_method" field.
func RequestMethodHasPrefix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasPrefix(FieldRequestMethod, v))
}

// RequestMethodHasSuffix applies the HasSuffix predicate on the "request_method" field.
func RequestMethodHasSuffix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasSuffix(FieldRequestMethod, v))
}

// RequestMethodIsNil applies the IsNil predicate on the "request_method" field.
func RequestMethodIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldRequestMethod))
}

// RequestMethodNotNil applies the NotNil predicate on the "request_method" field.
func RequestMethodNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldRequestMethod))
}

// RequestMethodEqualFold applies the EqualFold predicate on the "request_method" field.
func RequestMethodEqualFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEqualFold(FieldRequestMethod, v))
}

// RequestMethodContainsFold applies the ContainsFold predicate on the "request_method" field.
func RequestMethodContainsFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContainsFold(FieldRequestMethod, v))
}

// OperTypeEQ applies the EQ predicate on the "oper_type" field.
func OperTypeEQ(v int8) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperType, v))
}

// OperTypeNEQ applies the NEQ predicate on the "oper_type" field.
func OperTypeNEQ(v int8) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldOperType, v))
}

// OperTypeIn applies the In predicate on the "oper_type" field.
func OperTypeIn(vs ...int8) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldOperType, vs...))
}

// OperTypeNotIn applies the NotIn predicate on the "oper_type" field.
func OperTypeNotIn(vs ...int8) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldOperType, vs...))
}

// OperTypeGT applies the GT predicate on the "oper_type" field.
func OperTypeGT(v int8) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldOperType, v))
}

// OperTypeGTE applies the GTE predicate on the "oper_type" field.
func OperTypeGTE(v int8) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldOperType, v))
}

// OperTypeLT applies the LT predicate on the "oper_type" field.
func OperTypeLT(v int8) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldOperType, v))
}

// OperTypeLTE applies the LTE predicate on the "oper_type" field.
func OperTypeLTE(v int8) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldOperType, v))
}

// OperTypeIsNil applies the IsNil predicate on the "oper_type" field.
func OperTypeIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldOperType))
}

// OperTypeNotNil applies the NotNil predicate on the "oper_type" field.
func OperTypeNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldOperType))
}

// OperNameEQ applies the EQ predicate on the "oper_name" field.
func OperNameEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperName, v))
}

// OperNameNEQ applies the NEQ predicate on the "oper_name" field.
func OperNameNEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldOperName, v))
}

// OperNameIn applies the In predicate on the "oper_name" field.
func OperNameIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldOperName, vs...))
}

// OperNameNotIn applies the NotIn predicate on the "oper_name" field.
func OperNameNotIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldOperName, vs...))
}

// OperNameGT applies the GT predicate on the "oper_name" field.
func OperNameGT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldOperName, v))
}

// OperNameGTE applies the GTE predicate on the "oper_name" field.
func OperNameGTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldOperName, v))
}

// OperNameLT applies the LT predicate on the "oper_name" field.
func OperNameLT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldOperName, v))
}

// OperNameLTE applies the LTE predicate on the "oper_name" field.
func OperNameLTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldOperName, v))
}

// OperNameContains applies the Contains predicate on the "oper_name" field.
func OperNameContains(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContains(FieldOperName, v))
}

// OperNameHasPrefix applies the HasPrefix predicate on the "oper_name" field.
func OperNameHasPrefix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasPrefix(FieldOperName, v))
}

// OperNameHasSuffix applies the HasSuffix predicate on the "oper_name" field.
func OperNameHasSuffix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasSuffix(FieldOperName, v))
}

// OperNameIsNil applies the IsNil predicate on the "oper_name" field.
func OperNameIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldOperName))
}

// OperNameNotNil applies the NotNil predicate on the "oper_name" field.
func OperNameNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldOperName))
}

// OperNameEqualFold applies the EqualFold predicate on the "oper_name" field.
func OperNameEqualFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEqualFold(FieldOperName, v))
}

// OperNameContainsFold applies the ContainsFold predicate on the "oper_name" field.
func OperNameContainsFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContainsFold(FieldOperName, v))
}

// DeptNameEQ applies the EQ predicate on the "dept_name" field.
func DeptNameEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldDeptName, v))
}

// DeptNameNEQ applies the NEQ predicate on the "dept_name" field.
func DeptNameNEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldDeptName, v))
}

// DeptNameIn applies the In predicate on the "dept_name" field.
func DeptNameIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldDeptName, vs...))
}

// DeptNameNotIn applies the NotIn predicate on the "dept_name" field.
func DeptNameNotIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldDeptName, vs...))
}

// DeptNameGT applies the GT predicate on the "dept_name" field.
func DeptNameGT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldDeptName, v))
}

// DeptNameGTE applies the GTE predicate on the "dept_name" field.
func DeptNameGTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldDeptName, v))
}

// DeptNameLT applies the LT predicate on the "dept_name" field.
func DeptNameLT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldDeptName, v))
}

// DeptNameLTE applies the LTE predicate on the "dept_name" field.
func DeptNameLTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldDeptName, v))
}

// DeptNameContains applies the Contains predicate on the "dept_name" field.
func DeptNameContains(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContains(FieldDeptName, v))
}

// DeptNameHasPrefix applies the HasPrefix predicate on the "dept_name" field.
func DeptNameHasPrefix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasPrefix(FieldDeptName, v))
}

// DeptNameHasSuffix applies the HasSuffix predicate on the "dept_name" field.
func DeptNameHasSuffix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasSuffix(FieldDeptName, v))
}

// DeptNameIsNil applies the IsNil predicate on the "dept_name" field.
func DeptNameIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldDeptName))
}

// DeptNameNotNil applies the NotNil predicate on the "dept_name" field.
func DeptNameNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldDeptName))
}

// DeptNameEqualFold applies the EqualFold predicate on the "dept_name" field.
func DeptNameEqualFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEqualFold(FieldDeptName, v))
}

// DeptNameContainsFold applies the ContainsFold predicate on the "dept_name" field.
func DeptNameContainsFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContainsFold(FieldDeptName, v))
}

// OperURLEQ applies the EQ predicate on the "oper_url" field.
func OperURLEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperURL, v))
}

// OperURLNEQ applies the NEQ predicate on the "oper_url" field.
func OperURLNEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldOperURL, v))
}

// OperURLIn applies the In predicate on the "oper_url" field.
func OperURLIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldOperURL, vs...))
}

// OperURLNotIn applies the NotIn predicate on the "oper_url" field.
func OperURLNotIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldOperURL, vs...))
}

// OperURLGT applies the GT predicate on the "oper_url" field.
func OperURLGT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldOperURL, v))
}

// OperURLGTE applies the GTE predicate on the "oper_url" field.
func OperURLGTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldOperURL, v))
}

// OperURLLT applies the LT predicate on the "oper_url" field.
func OperURLLT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldOperURL, v))
}

// OperURLLTE applies the LTE predicate on the "oper_url" field.
func OperURLLTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldOperURL, v))
}

// OperURLContains applies the Contains predicate on the "oper_url" field.
func OperURLContains(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContains(FieldOperURL, v))
}

// OperURLHasPrefix applies the HasPrefix predicate on the "oper_url" field.
func OperURLHasPrefix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasPrefix(FieldOperURL, v))
}

// OperURLHasSuffix applies the HasSuffix predicate on the "oper_url" field.
func OperURLHasSuffix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasSuffix(FieldOperURL, v))
}

// OperURLIsNil applies the IsNil predicate on the "oper_url" field.
func OperURLIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldOperURL))
}

// OperURLNotNil applies the NotNil predicate on the "oper_url" field.
func OperURLNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldOperURL))
}

// OperURLEqualFold applies the EqualFold predicate on the "oper_url" field.
func OperURLEqualFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEqualFold(FieldOperURL, v))
}

// OperURLContainsFold applies the ContainsFold predicate on the "oper_url" field.
func OperURLContainsFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContainsFold(FieldOperURL, v))
}

// OperIPEQ applies the EQ predicate on the "oper_ip" field.
func OperIPEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperIP, v))
}

// OperIPNEQ applies the NEQ predicate on the "oper_ip" field.
func OperIPNEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldOperIP, v))
}

// OperIPIn applies the In predicate on the "oper_ip" field.
func OperIPIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldOperIP, vs...))
}

// OperIPNotIn applies the NotIn predicate on the "oper_ip" field.
func OperIPNotIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldOperIP, vs...))
}

// OperIPGT applies the GT predicate on the "oper_ip" field.
func OperIPGT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldOperIP, v))
}

// OperIPGTE applies the GTE predicate on the "oper_ip" field.
func OperIPGTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldOperIP, v))
}

// OperIPLT applies the LT predicate on the "oper_ip" field.
func OperIPLT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldOperIP, v))
}

// OperIPLTE applies the LTE predicate on the "oper_ip" field.
func OperIPLTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldOperIP, v))
}

// OperIPContains applies the Contains predicate on the "oper_ip" field.
func OperIPContains(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContains(FieldOperIP, v))
}

// OperIPHasPrefix applies the HasPrefix predicate on the "oper_ip" field.
func OperIPHasPrefix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasPrefix(FieldOperIP, v))
}

// OperIPHasSuffix applies the HasSuffix predicate on the "oper_ip" field.
func OperIPHasSuffix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasSuffix(FieldOperIP, v))
}

// OperIPIsNil applies the IsNil predicate on the "oper_ip" field.
func OperIPIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldOperIP))
}

// OperIPNotNil applies the NotNil predicate on the "oper_ip" field.
func OperIPNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldOperIP))
}

// OperIPEqualFold applies the EqualFold predicate on the "oper_ip" field.
func OperIPEqualFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEqualFold(FieldOperIP, v))
}

// OperIPContainsFold applies the ContainsFold predicate on the "oper_ip" field.
func OperIPContainsFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContainsFold(FieldOperIP, v))
}

// OperLocationEQ applies the EQ predicate on the "oper_location" field.
func OperLocationEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperLocation, v))
}

// OperLocationNEQ applies the NEQ predicate on the "oper_location" field.
func OperLocationNEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldOperLocation, v))
}

// OperLocationIn applies the In predicate on the "oper_location" field.
func OperLocationIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldOperLocation, vs...))
}

// OperLocationNotIn applies the NotIn predicate on the "oper_location" field.
func OperLocationNotIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldOperLocation, vs...))
}

// OperLocationGT applies the GT predicate on the "oper_location" field.
func OperLocationGT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldOperLocation, v))
}

// OperLocationGTE applies the GTE predicate on the "oper_location" field.
func OperLocationGTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldOperLocation, v))
}

// OperLocationLT applies the LT predicate on the "oper_location" field.
func OperLocationLT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldOperLocation, v))
}

// OperLocationLTE applies the LTE predicate on the "oper_location" field.
func OperLocationLTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldOperLocation, v))
}

// OperLocationContains applies the Contains predicate on the "oper_location" field.
func OperLocationContains(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContains(FieldOperLocation, v))
}

// OperLocationHasPrefix applies the HasPrefix predicate on the "oper_location" field.
func OperLocationHasPrefix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasPrefix(FieldOperLocation, v))
}

// OperLocationHasSuffix applies the HasSuffix predicate on the "oper_location" field.
func OperLocationHasSuffix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasSuffix(FieldOperLocation, v))
}

// OperLocationIsNil applies the IsNil predicate on the "oper_location" field.
func OperLocationIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldOperLocation))
}

// OperLocationNotNil applies the NotNil predicate on the "oper_location" field.
func OperLocationNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldOperLocation))
}

// OperLocationEqualFold applies the EqualFold predicate on the "oper_location" field.
func OperLocationEqualFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEqualFold(FieldOperLocation, v))
}

// OperLocationContainsFold applies the ContainsFold predicate on the "oper_location" field.
func OperLocationContainsFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContainsFold(FieldOperLocation, v))
}

// OperParamEQ applies the EQ predicate on the "oper_param" field.
func OperParamEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperParam, v))
}

// OperParamNEQ applies the NEQ predicate on the "oper_param" field.
func OperParamNEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldOperParam, v))
}

// OperParamIn applies the In predicate on the "oper_param" field.
func OperParamIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldOperParam, vs...))
}

// OperParamNotIn applies the NotIn predicate on the "oper_param" field.
func OperParamNotIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldOperParam, vs...))
}

// OperParamGT applies the GT predicate on the "oper_param" field.
func OperParamGT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldOperParam, v))
}

// OperParamGTE applies the GTE predicate on the "oper_param" field.
func OperParamGTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldOperParam, v))
}

// OperParamLT applies the LT predicate on the "oper_param" field.
func OperParamLT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldOperParam, v))
}

// OperParamLTE applies the LTE predicate on the "oper_param" field.
func OperParamLTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldOperParam, v))
}

// OperParamContains applies the Contains predicate on the "oper_param" field.
func OperParamContains(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContains(FieldOperParam, v))
}

// OperParamHasPrefix applies the HasPrefix predicate on the "oper_param" field.
func OperParamHasPrefix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasPrefix(FieldOperParam, v))
}

// OperParamHasSuffix applies the HasSuffix predicate on the "oper_param" field.
func OperParamHasSuffix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasSuffix(FieldOperParam, v))
}

// OperParamIsNil applies the IsNil predicate on the "oper_param" field.
func OperParamIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldOperParam))
}

// OperParamNotNil applies the NotNil predicate on the "oper_param" field.
func OperParamNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldOperParam))
}

// OperParamEqualFold applies the EqualFold predicate on the "oper_param" field.
func OperParamEqualFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEqualFold(FieldOperParam, v))
}

// OperParamContainsFold applies the ContainsFold predicate on the "oper_param" field.
func OperParamContainsFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContainsFold(FieldOperParam, v))
}

// ErrorMsgEQ applies the EQ predicate on the "error_msg" field.
func ErrorMsgEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldErrorMsg, v))
}

// ErrorMsgNEQ applies the NEQ predicate on the "error_msg" field.
func ErrorMsgNEQ(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldErrorMsg, v))
}

// ErrorMsgIn applies the In predicate on the "error_msg" field.
func ErrorMsgIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldErrorMsg, vs...))
}

// ErrorMsgNotIn applies the NotIn predicate on the "error_msg" field.
func ErrorMsgNotIn(vs ...string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldErrorMsg, vs...))
}

// ErrorMsgGT applies the GT predicate on the "error_msg" field.
func ErrorMsgGT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldErrorMsg, v))
}

// ErrorMsgGTE applies the GTE predicate on the "error_msg" field.
func ErrorMsgGTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldErrorMsg, v))
}

// ErrorMsgLT applies the LT predicate on the "error_msg" field.
func ErrorMsgLT(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldErrorMsg, v))
}

// ErrorMsgLTE applies the LTE predicate on the "error_msg" field.
func ErrorMsgLTE(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldErrorMsg, v))
}

// ErrorMsgContains applies the Contains predicate on the "error_msg" field.
func ErrorMsgContains(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContains(FieldErrorMsg, v))
}

// ErrorMsgHasPrefix applies the HasPrefix predicate on the "error_msg" field.
func ErrorMsgHasPrefix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasPrefix(FieldErrorMsg, v))
}

// ErrorMsgHasSuffix applies the HasSuffix predicate on the "error_msg" field.
func ErrorMsgHasSuffix(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldHasSuffix(FieldErrorMsg, v))
}

// ErrorMsgIsNil applies the IsNil predicate on the "error_msg" field.
func ErrorMsgIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldErrorMsg))
}

// ErrorMsgNotNil applies the NotNil predicate on the "error_msg" field.
func ErrorMsgNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldErrorMsg))
}

// ErrorMsgEqualFold applies the EqualFold predicate on the "error_msg" field.
func ErrorMsgEqualFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEqualFold(FieldErrorMsg, v))
}

// ErrorMsgContainsFold applies the ContainsFold predicate on the "error_msg" field.
func ErrorMsgContainsFold(v string) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldContainsFold(FieldErrorMsg, v))
}

// OperTimeEQ applies the EQ predicate on the "oper_time" field.
func OperTimeEQ(v time.Time) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldEQ(FieldOperTime, v))
}

// OperTimeNEQ applies the NEQ predicate on the "oper_time" field.
func OperTimeNEQ(v time.Time) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNEQ(FieldOperTime, v))
}

// OperTimeIn applies the In predicate on the "oper_time" field.
func OperTimeIn(vs ...time.Time) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIn(FieldOperTime, vs...))
}

// OperTimeNotIn applies the NotIn predicate on the "oper_time" field.
func OperTimeNotIn(vs ...time.Time) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotIn(FieldOperTime, vs...))
}

// OperTimeGT applies the GT predicate on the "oper_time" field.
func OperTimeGT(v time.Time) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGT(FieldOperTime, v))
}

// OperTimeGTE applies the GTE predicate on the "oper_time" field.
func OperTimeGTE(v time.Time) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldGTE(FieldOperTime, v))
}

// OperTimeLT applies the LT predicate on the "oper_time" field.
func OperTimeLT(v time.Time) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLT(FieldOperTime, v))
}

// OperTimeLTE applies the LTE predicate on the "oper_time" field.
func OperTimeLTE(v time.Time) predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldLTE(FieldOperTime, v))
}

// OperTimeIsNil applies the IsNil predicate on the "oper_time" field.
func OperTimeIsNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldIsNull(FieldOperTime))
}

// OperTimeNotNil applies the NotNil predicate on the "oper_time" field.
func OperTimeNotNil() predicate.SysOperLog {
	return predicate.SysOperLog(sql.FieldNotNull(FieldOperTime))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysOperLog) predicate.SysOperLog {
	return predicate.SysOperLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysOperLog) predicate.SysOperLog {
	return predicate.SysOperLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysOperLog) predicate.SysOperLog {
	return predicate.SysOperLog(sql.NotPredicates(p))
}