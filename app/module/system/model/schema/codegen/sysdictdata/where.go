// Code generated by ent, DO NOT EDIT.

package sysdictdata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDeleteAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeleteBy applies equality check predicate on the "delete_by" field. It's identical to DeleteByEQ.
func DeleteBy(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDeleteBy, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldRemark, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldStatus, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldSort, v))
}

// DelFlag applies equality check predicate on the "del_flag" field. It's identical to DelFlagEQ.
func DelFlag(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDelFlag, v))
}

// DictLabel applies equality check predicate on the "dict_label" field. It's identical to DictLabelEQ.
func DictLabel(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDictLabel, v))
}

// DictValue applies equality check predicate on the "dict_value" field. It's identical to DictValueEQ.
func DictValue(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDictValue, v))
}

// DictTypeID applies equality check predicate on the "dict_type_id" field. It's identical to DictTypeIDEQ.
func DictTypeID(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDictTypeID, v))
}

// CSSClass applies equality check predicate on the "css_class" field. It's identical to CSSClassEQ.
func CSSClass(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldCSSClass, v))
}

// ListClass applies equality check predicate on the "list_class" field. It's identical to ListClassEQ.
func ListClass(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldListClass, v))
}

// IsDefault applies equality check predicate on the "is_default" field. It's identical to IsDefaultEQ.
func IsDefault(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldIsDefault, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDeleteAt, v))
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldDeleteAt, v))
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldDeleteAt, vs...))
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldDeleteAt, vs...))
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldDeleteAt, v))
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldDeleteAt, v))
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldDeleteAt, v))
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v time.Time) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldDeleteAt, v))
}

// DeleteAtIsNil applies the IsNil predicate on the "delete_at" field.
func DeleteAtIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldDeleteAt))
}

// DeleteAtNotNil applies the NotNil predicate on the "delete_at" field.
func DeleteAtNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldDeleteAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldCreatedBy))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldUpdatedBy))
}

// DeleteByEQ applies the EQ predicate on the "delete_by" field.
func DeleteByEQ(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDeleteBy, v))
}

// DeleteByNEQ applies the NEQ predicate on the "delete_by" field.
func DeleteByNEQ(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldDeleteBy, v))
}

// DeleteByIn applies the In predicate on the "delete_by" field.
func DeleteByIn(vs ...int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldDeleteBy, vs...))
}

// DeleteByNotIn applies the NotIn predicate on the "delete_by" field.
func DeleteByNotIn(vs ...int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldDeleteBy, vs...))
}

// DeleteByGT applies the GT predicate on the "delete_by" field.
func DeleteByGT(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldDeleteBy, v))
}

// DeleteByGTE applies the GTE predicate on the "delete_by" field.
func DeleteByGTE(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldDeleteBy, v))
}

// DeleteByLT applies the LT predicate on the "delete_by" field.
func DeleteByLT(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldDeleteBy, v))
}

// DeleteByLTE applies the LTE predicate on the "delete_by" field.
func DeleteByLTE(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldDeleteBy, v))
}

// DeleteByIsNil applies the IsNil predicate on the "delete_by" field.
func DeleteByIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldDeleteBy))
}

// DeleteByNotNil applies the NotNil predicate on the "delete_by" field.
func DeleteByNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldDeleteBy))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldContainsFold(FieldRemark, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldStatus, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldSort, v))
}

// SortIsNil applies the IsNil predicate on the "sort" field.
func SortIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldSort))
}

// SortNotNil applies the NotNil predicate on the "sort" field.
func SortNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldSort))
}

// DelFlagEQ applies the EQ predicate on the "del_flag" field.
func DelFlagEQ(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDelFlag, v))
}

// DelFlagNEQ applies the NEQ predicate on the "del_flag" field.
func DelFlagNEQ(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldDelFlag, v))
}

// DelFlagIn applies the In predicate on the "del_flag" field.
func DelFlagIn(vs ...int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldDelFlag, vs...))
}

// DelFlagNotIn applies the NotIn predicate on the "del_flag" field.
func DelFlagNotIn(vs ...int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldDelFlag, vs...))
}

// DelFlagGT applies the GT predicate on the "del_flag" field.
func DelFlagGT(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldDelFlag, v))
}

// DelFlagGTE applies the GTE predicate on the "del_flag" field.
func DelFlagGTE(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldDelFlag, v))
}

// DelFlagLT applies the LT predicate on the "del_flag" field.
func DelFlagLT(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldDelFlag, v))
}

// DelFlagLTE applies the LTE predicate on the "del_flag" field.
func DelFlagLTE(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldDelFlag, v))
}

// DictLabelEQ applies the EQ predicate on the "dict_label" field.
func DictLabelEQ(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDictLabel, v))
}

// DictLabelNEQ applies the NEQ predicate on the "dict_label" field.
func DictLabelNEQ(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldDictLabel, v))
}

// DictLabelIn applies the In predicate on the "dict_label" field.
func DictLabelIn(vs ...string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldDictLabel, vs...))
}

// DictLabelNotIn applies the NotIn predicate on the "dict_label" field.
func DictLabelNotIn(vs ...string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldDictLabel, vs...))
}

// DictLabelGT applies the GT predicate on the "dict_label" field.
func DictLabelGT(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldDictLabel, v))
}

// DictLabelGTE applies the GTE predicate on the "dict_label" field.
func DictLabelGTE(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldDictLabel, v))
}

// DictLabelLT applies the LT predicate on the "dict_label" field.
func DictLabelLT(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldDictLabel, v))
}

// DictLabelLTE applies the LTE predicate on the "dict_label" field.
func DictLabelLTE(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldDictLabel, v))
}

// DictLabelContains applies the Contains predicate on the "dict_label" field.
func DictLabelContains(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldContains(FieldDictLabel, v))
}

// DictLabelHasPrefix applies the HasPrefix predicate on the "dict_label" field.
func DictLabelHasPrefix(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldHasPrefix(FieldDictLabel, v))
}

// DictLabelHasSuffix applies the HasSuffix predicate on the "dict_label" field.
func DictLabelHasSuffix(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldHasSuffix(FieldDictLabel, v))
}

// DictLabelIsNil applies the IsNil predicate on the "dict_label" field.
func DictLabelIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldDictLabel))
}

// DictLabelNotNil applies the NotNil predicate on the "dict_label" field.
func DictLabelNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldDictLabel))
}

// DictLabelEqualFold applies the EqualFold predicate on the "dict_label" field.
func DictLabelEqualFold(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEqualFold(FieldDictLabel, v))
}

// DictLabelContainsFold applies the ContainsFold predicate on the "dict_label" field.
func DictLabelContainsFold(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldContainsFold(FieldDictLabel, v))
}

// DictValueEQ applies the EQ predicate on the "dict_value" field.
func DictValueEQ(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDictValue, v))
}

// DictValueNEQ applies the NEQ predicate on the "dict_value" field.
func DictValueNEQ(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldDictValue, v))
}

// DictValueIn applies the In predicate on the "dict_value" field.
func DictValueIn(vs ...string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldDictValue, vs...))
}

// DictValueNotIn applies the NotIn predicate on the "dict_value" field.
func DictValueNotIn(vs ...string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldDictValue, vs...))
}

// DictValueGT applies the GT predicate on the "dict_value" field.
func DictValueGT(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldDictValue, v))
}

// DictValueGTE applies the GTE predicate on the "dict_value" field.
func DictValueGTE(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldDictValue, v))
}

// DictValueLT applies the LT predicate on the "dict_value" field.
func DictValueLT(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldDictValue, v))
}

// DictValueLTE applies the LTE predicate on the "dict_value" field.
func DictValueLTE(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldDictValue, v))
}

// DictValueContains applies the Contains predicate on the "dict_value" field.
func DictValueContains(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldContains(FieldDictValue, v))
}

// DictValueHasPrefix applies the HasPrefix predicate on the "dict_value" field.
func DictValueHasPrefix(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldHasPrefix(FieldDictValue, v))
}

// DictValueHasSuffix applies the HasSuffix predicate on the "dict_value" field.
func DictValueHasSuffix(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldHasSuffix(FieldDictValue, v))
}

// DictValueIsNil applies the IsNil predicate on the "dict_value" field.
func DictValueIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldDictValue))
}

// DictValueNotNil applies the NotNil predicate on the "dict_value" field.
func DictValueNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldDictValue))
}

// DictValueEqualFold applies the EqualFold predicate on the "dict_value" field.
func DictValueEqualFold(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEqualFold(FieldDictValue, v))
}

// DictValueContainsFold applies the ContainsFold predicate on the "dict_value" field.
func DictValueContainsFold(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldContainsFold(FieldDictValue, v))
}

// DictTypeIDEQ applies the EQ predicate on the "dict_type_id" field.
func DictTypeIDEQ(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldDictTypeID, v))
}

// DictTypeIDNEQ applies the NEQ predicate on the "dict_type_id" field.
func DictTypeIDNEQ(v int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldDictTypeID, v))
}

// DictTypeIDIn applies the In predicate on the "dict_type_id" field.
func DictTypeIDIn(vs ...int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldDictTypeID, vs...))
}

// DictTypeIDNotIn applies the NotIn predicate on the "dict_type_id" field.
func DictTypeIDNotIn(vs ...int64) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldDictTypeID, vs...))
}

// DictTypeIDIsNil applies the IsNil predicate on the "dict_type_id" field.
func DictTypeIDIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldDictTypeID))
}

// DictTypeIDNotNil applies the NotNil predicate on the "dict_type_id" field.
func DictTypeIDNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldDictTypeID))
}

// CSSClassEQ applies the EQ predicate on the "css_class" field.
func CSSClassEQ(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldCSSClass, v))
}

// CSSClassNEQ applies the NEQ predicate on the "css_class" field.
func CSSClassNEQ(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldCSSClass, v))
}

// CSSClassIn applies the In predicate on the "css_class" field.
func CSSClassIn(vs ...string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldCSSClass, vs...))
}

// CSSClassNotIn applies the NotIn predicate on the "css_class" field.
func CSSClassNotIn(vs ...string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldCSSClass, vs...))
}

// CSSClassGT applies the GT predicate on the "css_class" field.
func CSSClassGT(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldCSSClass, v))
}

// CSSClassGTE applies the GTE predicate on the "css_class" field.
func CSSClassGTE(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldCSSClass, v))
}

// CSSClassLT applies the LT predicate on the "css_class" field.
func CSSClassLT(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldCSSClass, v))
}

// CSSClassLTE applies the LTE predicate on the "css_class" field.
func CSSClassLTE(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldCSSClass, v))
}

// CSSClassContains applies the Contains predicate on the "css_class" field.
func CSSClassContains(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldContains(FieldCSSClass, v))
}

// CSSClassHasPrefix applies the HasPrefix predicate on the "css_class" field.
func CSSClassHasPrefix(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldHasPrefix(FieldCSSClass, v))
}

// CSSClassHasSuffix applies the HasSuffix predicate on the "css_class" field.
func CSSClassHasSuffix(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldHasSuffix(FieldCSSClass, v))
}

// CSSClassIsNil applies the IsNil predicate on the "css_class" field.
func CSSClassIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldCSSClass))
}

// CSSClassNotNil applies the NotNil predicate on the "css_class" field.
func CSSClassNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldCSSClass))
}

// CSSClassEqualFold applies the EqualFold predicate on the "css_class" field.
func CSSClassEqualFold(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEqualFold(FieldCSSClass, v))
}

// CSSClassContainsFold applies the ContainsFold predicate on the "css_class" field.
func CSSClassContainsFold(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldContainsFold(FieldCSSClass, v))
}

// ListClassEQ applies the EQ predicate on the "list_class" field.
func ListClassEQ(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldListClass, v))
}

// ListClassNEQ applies the NEQ predicate on the "list_class" field.
func ListClassNEQ(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldListClass, v))
}

// ListClassIn applies the In predicate on the "list_class" field.
func ListClassIn(vs ...string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldListClass, vs...))
}

// ListClassNotIn applies the NotIn predicate on the "list_class" field.
func ListClassNotIn(vs ...string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldListClass, vs...))
}

// ListClassGT applies the GT predicate on the "list_class" field.
func ListClassGT(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldListClass, v))
}

// ListClassGTE applies the GTE predicate on the "list_class" field.
func ListClassGTE(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldListClass, v))
}

// ListClassLT applies the LT predicate on the "list_class" field.
func ListClassLT(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldListClass, v))
}

// ListClassLTE applies the LTE predicate on the "list_class" field.
func ListClassLTE(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldListClass, v))
}

// ListClassContains applies the Contains predicate on the "list_class" field.
func ListClassContains(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldContains(FieldListClass, v))
}

// ListClassHasPrefix applies the HasPrefix predicate on the "list_class" field.
func ListClassHasPrefix(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldHasPrefix(FieldListClass, v))
}

// ListClassHasSuffix applies the HasSuffix predicate on the "list_class" field.
func ListClassHasSuffix(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldHasSuffix(FieldListClass, v))
}

// ListClassIsNil applies the IsNil predicate on the "list_class" field.
func ListClassIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldListClass))
}

// ListClassNotNil applies the NotNil predicate on the "list_class" field.
func ListClassNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldListClass))
}

// ListClassEqualFold applies the EqualFold predicate on the "list_class" field.
func ListClassEqualFold(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEqualFold(FieldListClass, v))
}

// ListClassContainsFold applies the ContainsFold predicate on the "list_class" field.
func ListClassContainsFold(v string) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldContainsFold(FieldListClass, v))
}

// IsDefaultEQ applies the EQ predicate on the "is_default" field.
func IsDefaultEQ(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldEQ(FieldIsDefault, v))
}

// IsDefaultNEQ applies the NEQ predicate on the "is_default" field.
func IsDefaultNEQ(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNEQ(FieldIsDefault, v))
}

// IsDefaultIn applies the In predicate on the "is_default" field.
func IsDefaultIn(vs ...int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIn(FieldIsDefault, vs...))
}

// IsDefaultNotIn applies the NotIn predicate on the "is_default" field.
func IsDefaultNotIn(vs ...int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotIn(FieldIsDefault, vs...))
}

// IsDefaultGT applies the GT predicate on the "is_default" field.
func IsDefaultGT(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGT(FieldIsDefault, v))
}

// IsDefaultGTE applies the GTE predicate on the "is_default" field.
func IsDefaultGTE(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldGTE(FieldIsDefault, v))
}

// IsDefaultLT applies the LT predicate on the "is_default" field.
func IsDefaultLT(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLT(FieldIsDefault, v))
}

// IsDefaultLTE applies the LTE predicate on the "is_default" field.
func IsDefaultLTE(v int8) predicate.SysDictData {
	return predicate.SysDictData(sql.FieldLTE(FieldIsDefault, v))
}

// IsDefaultIsNil applies the IsNil predicate on the "is_default" field.
func IsDefaultIsNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldIsNull(FieldIsDefault))
}

// IsDefaultNotNil applies the NotNil predicate on the "is_default" field.
func IsDefaultNotNil() predicate.SysDictData {
	return predicate.SysDictData(sql.FieldNotNull(FieldIsDefault))
}

// HasSysDictType applies the HasEdge predicate on the "sysDictType" edge.
func HasSysDictType() predicate.SysDictData {
	return predicate.SysDictData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SysDictTypeTable, SysDictTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSysDictTypeWith applies the HasEdge predicate on the "sysDictType" edge with a given conditions (other predicates).
func HasSysDictTypeWith(preds ...predicate.SysDictType) predicate.SysDictData {
	return predicate.SysDictData(func(s *sql.Selector) {
		step := newSysDictTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysDictData) predicate.SysDictData {
	return predicate.SysDictData(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysDictData) predicate.SysDictData {
	return predicate.SysDictData(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysDictData) predicate.SysDictData {
	return predicate.SysDictData(sql.NotPredicates(p))
}
