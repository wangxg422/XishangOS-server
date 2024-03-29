// Code generated by ent, DO NOT EDIT.

package syspost

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldDeleteAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeleteBy applies equality check predicate on the "delete_by" field. It's identical to DeleteByEQ.
func DeleteBy(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldDeleteBy, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldRemark, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldStatus, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldSort, v))
}

// DelFlag applies equality check predicate on the "del_flag" field. It's identical to DelFlagEQ.
func DelFlag(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldDelFlag, v))
}

// PostCode applies equality check predicate on the "post_code" field. It's identical to PostCodeEQ.
func PostCode(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldPostCode, v))
}

// PostName applies equality check predicate on the "post_name" field. It's identical to PostNameEQ.
func PostName(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldPostName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldDeleteAt, v))
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldDeleteAt, v))
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldDeleteAt, vs...))
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldDeleteAt, vs...))
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldDeleteAt, v))
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldDeleteAt, v))
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldDeleteAt, v))
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v time.Time) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldDeleteAt, v))
}

// DeleteAtIsNil applies the IsNil predicate on the "delete_at" field.
func DeleteAtIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldDeleteAt))
}

// DeleteAtNotNil applies the NotNil predicate on the "delete_at" field.
func DeleteAtNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldDeleteAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldCreatedBy))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldUpdatedBy))
}

// DeleteByEQ applies the EQ predicate on the "delete_by" field.
func DeleteByEQ(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldDeleteBy, v))
}

// DeleteByNEQ applies the NEQ predicate on the "delete_by" field.
func DeleteByNEQ(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldDeleteBy, v))
}

// DeleteByIn applies the In predicate on the "delete_by" field.
func DeleteByIn(vs ...int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldDeleteBy, vs...))
}

// DeleteByNotIn applies the NotIn predicate on the "delete_by" field.
func DeleteByNotIn(vs ...int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldDeleteBy, vs...))
}

// DeleteByGT applies the GT predicate on the "delete_by" field.
func DeleteByGT(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldDeleteBy, v))
}

// DeleteByGTE applies the GTE predicate on the "delete_by" field.
func DeleteByGTE(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldDeleteBy, v))
}

// DeleteByLT applies the LT predicate on the "delete_by" field.
func DeleteByLT(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldDeleteBy, v))
}

// DeleteByLTE applies the LTE predicate on the "delete_by" field.
func DeleteByLTE(v int64) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldDeleteBy, v))
}

// DeleteByIsNil applies the IsNil predicate on the "delete_by" field.
func DeleteByIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldDeleteBy))
}

// DeleteByNotNil applies the NotNil predicate on the "delete_by" field.
func DeleteByNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldDeleteBy))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContainsFold(FieldRemark, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldStatus, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldSort, v))
}

// SortIsNil applies the IsNil predicate on the "sort" field.
func SortIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldSort))
}

// SortNotNil applies the NotNil predicate on the "sort" field.
func SortNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldSort))
}

// DelFlagEQ applies the EQ predicate on the "del_flag" field.
func DelFlagEQ(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldDelFlag, v))
}

// DelFlagNEQ applies the NEQ predicate on the "del_flag" field.
func DelFlagNEQ(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldDelFlag, v))
}

// DelFlagIn applies the In predicate on the "del_flag" field.
func DelFlagIn(vs ...int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldDelFlag, vs...))
}

// DelFlagNotIn applies the NotIn predicate on the "del_flag" field.
func DelFlagNotIn(vs ...int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldDelFlag, vs...))
}

// DelFlagGT applies the GT predicate on the "del_flag" field.
func DelFlagGT(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldDelFlag, v))
}

// DelFlagGTE applies the GTE predicate on the "del_flag" field.
func DelFlagGTE(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldDelFlag, v))
}

// DelFlagLT applies the LT predicate on the "del_flag" field.
func DelFlagLT(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldDelFlag, v))
}

// DelFlagLTE applies the LTE predicate on the "del_flag" field.
func DelFlagLTE(v int8) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldDelFlag, v))
}

// PostCodeEQ applies the EQ predicate on the "post_code" field.
func PostCodeEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldPostCode, v))
}

// PostCodeNEQ applies the NEQ predicate on the "post_code" field.
func PostCodeNEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldPostCode, v))
}

// PostCodeIn applies the In predicate on the "post_code" field.
func PostCodeIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldPostCode, vs...))
}

// PostCodeNotIn applies the NotIn predicate on the "post_code" field.
func PostCodeNotIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldPostCode, vs...))
}

// PostCodeGT applies the GT predicate on the "post_code" field.
func PostCodeGT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldPostCode, v))
}

// PostCodeGTE applies the GTE predicate on the "post_code" field.
func PostCodeGTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldPostCode, v))
}

// PostCodeLT applies the LT predicate on the "post_code" field.
func PostCodeLT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldPostCode, v))
}

// PostCodeLTE applies the LTE predicate on the "post_code" field.
func PostCodeLTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldPostCode, v))
}

// PostCodeContains applies the Contains predicate on the "post_code" field.
func PostCodeContains(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContains(FieldPostCode, v))
}

// PostCodeHasPrefix applies the HasPrefix predicate on the "post_code" field.
func PostCodeHasPrefix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasPrefix(FieldPostCode, v))
}

// PostCodeHasSuffix applies the HasSuffix predicate on the "post_code" field.
func PostCodeHasSuffix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasSuffix(FieldPostCode, v))
}

// PostCodeIsNil applies the IsNil predicate on the "post_code" field.
func PostCodeIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldPostCode))
}

// PostCodeNotNil applies the NotNil predicate on the "post_code" field.
func PostCodeNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldPostCode))
}

// PostCodeEqualFold applies the EqualFold predicate on the "post_code" field.
func PostCodeEqualFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEqualFold(FieldPostCode, v))
}

// PostCodeContainsFold applies the ContainsFold predicate on the "post_code" field.
func PostCodeContainsFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContainsFold(FieldPostCode, v))
}

// PostNameEQ applies the EQ predicate on the "post_name" field.
func PostNameEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEQ(FieldPostName, v))
}

// PostNameNEQ applies the NEQ predicate on the "post_name" field.
func PostNameNEQ(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNEQ(FieldPostName, v))
}

// PostNameIn applies the In predicate on the "post_name" field.
func PostNameIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldIn(FieldPostName, vs...))
}

// PostNameNotIn applies the NotIn predicate on the "post_name" field.
func PostNameNotIn(vs ...string) predicate.SysPost {
	return predicate.SysPost(sql.FieldNotIn(FieldPostName, vs...))
}

// PostNameGT applies the GT predicate on the "post_name" field.
func PostNameGT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGT(FieldPostName, v))
}

// PostNameGTE applies the GTE predicate on the "post_name" field.
func PostNameGTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldGTE(FieldPostName, v))
}

// PostNameLT applies the LT predicate on the "post_name" field.
func PostNameLT(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLT(FieldPostName, v))
}

// PostNameLTE applies the LTE predicate on the "post_name" field.
func PostNameLTE(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldLTE(FieldPostName, v))
}

// PostNameContains applies the Contains predicate on the "post_name" field.
func PostNameContains(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContains(FieldPostName, v))
}

// PostNameHasPrefix applies the HasPrefix predicate on the "post_name" field.
func PostNameHasPrefix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasPrefix(FieldPostName, v))
}

// PostNameHasSuffix applies the HasSuffix predicate on the "post_name" field.
func PostNameHasSuffix(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldHasSuffix(FieldPostName, v))
}

// PostNameIsNil applies the IsNil predicate on the "post_name" field.
func PostNameIsNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldIsNull(FieldPostName))
}

// PostNameNotNil applies the NotNil predicate on the "post_name" field.
func PostNameNotNil() predicate.SysPost {
	return predicate.SysPost(sql.FieldNotNull(FieldPostName))
}

// PostNameEqualFold applies the EqualFold predicate on the "post_name" field.
func PostNameEqualFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldEqualFold(FieldPostName, v))
}

// PostNameContainsFold applies the ContainsFold predicate on the "post_name" field.
func PostNameContainsFold(v string) predicate.SysPost {
	return predicate.SysPost(sql.FieldContainsFold(FieldPostName, v))
}

// HasSysUsers applies the HasEdge predicate on the "sysUsers" edge.
func HasSysUsers() predicate.SysPost {
	return predicate.SysPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SysUsersTable, SysUsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSysUsersWith applies the HasEdge predicate on the "sysUsers" edge with a given conditions (other predicates).
func HasSysUsersWith(preds ...predicate.SysUser) predicate.SysPost {
	return predicate.SysPost(func(s *sql.Selector) {
		step := newSysUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysPost) predicate.SysPost {
	return predicate.SysPost(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysPost) predicate.SysPost {
	return predicate.SysPost(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysPost) predicate.SysPost {
	return predicate.SysPost(sql.NotPredicates(p))
}
