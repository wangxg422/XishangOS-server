// Code generated by ent, DO NOT EDIT.

package apppackage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldDeleteAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldStatus, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeleteBy applies equality check predicate on the "delete_by" field. It's identical to DeleteByEQ.
func DeleteBy(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldDeleteBy, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldRemark, v))
}

// DelFlag applies equality check predicate on the "del_flag" field. It's identical to DelFlagEQ.
func DelFlag(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldDelFlag, v))
}

// PkgName applies equality check predicate on the "pkg_name" field. It's identical to PkgNameEQ.
func PkgName(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPkgName, v))
}

// PkgCode applies equality check predicate on the "pkg_code" field. It's identical to PkgCodeEQ.
func PkgCode(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPkgCode, v))
}

// PkgVersion applies equality check predicate on the "pkg_version" field. It's identical to PkgVersionEQ.
func PkgVersion(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPkgVersion, v))
}

// PkgType applies equality check predicate on the "pkg_type" field. It's identical to PkgTypeEQ.
func PkgType(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPkgType, v))
}

// PkgKind applies equality check predicate on the "pkg_kind" field. It's identical to PkgKindEQ.
func PkgKind(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPkgKind, v))
}

// Uploader applies equality check predicate on the "uploader" field. It's identical to UploaderEQ.
func Uploader(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldUploader, v))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldDesc, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldDeleteAt, v))
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldDeleteAt, v))
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldDeleteAt, vs...))
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldDeleteAt, vs...))
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldDeleteAt, v))
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldDeleteAt, v))
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldDeleteAt, v))
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldDeleteAt, v))
}

// DeleteAtIsNil applies the IsNil predicate on the "delete_at" field.
func DeleteAtIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldDeleteAt))
}

// DeleteAtNotNil applies the NotNil predicate on the "delete_at" field.
func DeleteAtNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldDeleteAt))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldStatus, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldCreatedBy))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldUpdatedBy))
}

// DeleteByEQ applies the EQ predicate on the "delete_by" field.
func DeleteByEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldDeleteBy, v))
}

// DeleteByNEQ applies the NEQ predicate on the "delete_by" field.
func DeleteByNEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldDeleteBy, v))
}

// DeleteByIn applies the In predicate on the "delete_by" field.
func DeleteByIn(vs ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldDeleteBy, vs...))
}

// DeleteByNotIn applies the NotIn predicate on the "delete_by" field.
func DeleteByNotIn(vs ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldDeleteBy, vs...))
}

// DeleteByGT applies the GT predicate on the "delete_by" field.
func DeleteByGT(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldDeleteBy, v))
}

// DeleteByGTE applies the GTE predicate on the "delete_by" field.
func DeleteByGTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldDeleteBy, v))
}

// DeleteByLT applies the LT predicate on the "delete_by" field.
func DeleteByLT(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldDeleteBy, v))
}

// DeleteByLTE applies the LTE predicate on the "delete_by" field.
func DeleteByLTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldDeleteBy, v))
}

// DeleteByIsNil applies the IsNil predicate on the "delete_by" field.
func DeleteByIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldDeleteBy))
}

// DeleteByNotNil applies the NotNil predicate on the "delete_by" field.
func DeleteByNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldDeleteBy))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContainsFold(FieldRemark, v))
}

// DelFlagEQ applies the EQ predicate on the "del_flag" field.
func DelFlagEQ(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldDelFlag, v))
}

// DelFlagNEQ applies the NEQ predicate on the "del_flag" field.
func DelFlagNEQ(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldDelFlag, v))
}

// DelFlagIn applies the In predicate on the "del_flag" field.
func DelFlagIn(vs ...int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldDelFlag, vs...))
}

// DelFlagNotIn applies the NotIn predicate on the "del_flag" field.
func DelFlagNotIn(vs ...int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldDelFlag, vs...))
}

// DelFlagGT applies the GT predicate on the "del_flag" field.
func DelFlagGT(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldDelFlag, v))
}

// DelFlagGTE applies the GTE predicate on the "del_flag" field.
func DelFlagGTE(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldDelFlag, v))
}

// DelFlagLT applies the LT predicate on the "del_flag" field.
func DelFlagLT(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldDelFlag, v))
}

// DelFlagLTE applies the LTE predicate on the "del_flag" field.
func DelFlagLTE(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldDelFlag, v))
}

// PkgNameEQ applies the EQ predicate on the "pkg_name" field.
func PkgNameEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPkgName, v))
}

// PkgNameNEQ applies the NEQ predicate on the "pkg_name" field.
func PkgNameNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldPkgName, v))
}

// PkgNameIn applies the In predicate on the "pkg_name" field.
func PkgNameIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldPkgName, vs...))
}

// PkgNameNotIn applies the NotIn predicate on the "pkg_name" field.
func PkgNameNotIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldPkgName, vs...))
}

// PkgNameGT applies the GT predicate on the "pkg_name" field.
func PkgNameGT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldPkgName, v))
}

// PkgNameGTE applies the GTE predicate on the "pkg_name" field.
func PkgNameGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldPkgName, v))
}

// PkgNameLT applies the LT predicate on the "pkg_name" field.
func PkgNameLT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldPkgName, v))
}

// PkgNameLTE applies the LTE predicate on the "pkg_name" field.
func PkgNameLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldPkgName, v))
}

// PkgNameContains applies the Contains predicate on the "pkg_name" field.
func PkgNameContains(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContains(FieldPkgName, v))
}

// PkgNameHasPrefix applies the HasPrefix predicate on the "pkg_name" field.
func PkgNameHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasPrefix(FieldPkgName, v))
}

// PkgNameHasSuffix applies the HasSuffix predicate on the "pkg_name" field.
func PkgNameHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasSuffix(FieldPkgName, v))
}

// PkgNameEqualFold applies the EqualFold predicate on the "pkg_name" field.
func PkgNameEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEqualFold(FieldPkgName, v))
}

// PkgNameContainsFold applies the ContainsFold predicate on the "pkg_name" field.
func PkgNameContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContainsFold(FieldPkgName, v))
}

// PkgCodeEQ applies the EQ predicate on the "pkg_code" field.
func PkgCodeEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPkgCode, v))
}

// PkgCodeNEQ applies the NEQ predicate on the "pkg_code" field.
func PkgCodeNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldPkgCode, v))
}

// PkgCodeIn applies the In predicate on the "pkg_code" field.
func PkgCodeIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldPkgCode, vs...))
}

// PkgCodeNotIn applies the NotIn predicate on the "pkg_code" field.
func PkgCodeNotIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldPkgCode, vs...))
}

// PkgCodeGT applies the GT predicate on the "pkg_code" field.
func PkgCodeGT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldPkgCode, v))
}

// PkgCodeGTE applies the GTE predicate on the "pkg_code" field.
func PkgCodeGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldPkgCode, v))
}

// PkgCodeLT applies the LT predicate on the "pkg_code" field.
func PkgCodeLT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldPkgCode, v))
}

// PkgCodeLTE applies the LTE predicate on the "pkg_code" field.
func PkgCodeLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldPkgCode, v))
}

// PkgCodeContains applies the Contains predicate on the "pkg_code" field.
func PkgCodeContains(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContains(FieldPkgCode, v))
}

// PkgCodeHasPrefix applies the HasPrefix predicate on the "pkg_code" field.
func PkgCodeHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasPrefix(FieldPkgCode, v))
}

// PkgCodeHasSuffix applies the HasSuffix predicate on the "pkg_code" field.
func PkgCodeHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasSuffix(FieldPkgCode, v))
}

// PkgCodeEqualFold applies the EqualFold predicate on the "pkg_code" field.
func PkgCodeEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEqualFold(FieldPkgCode, v))
}

// PkgCodeContainsFold applies the ContainsFold predicate on the "pkg_code" field.
func PkgCodeContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContainsFold(FieldPkgCode, v))
}

// PkgVersionEQ applies the EQ predicate on the "pkg_version" field.
func PkgVersionEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPkgVersion, v))
}

// PkgVersionNEQ applies the NEQ predicate on the "pkg_version" field.
func PkgVersionNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldPkgVersion, v))
}

// PkgVersionIn applies the In predicate on the "pkg_version" field.
func PkgVersionIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldPkgVersion, vs...))
}

// PkgVersionNotIn applies the NotIn predicate on the "pkg_version" field.
func PkgVersionNotIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldPkgVersion, vs...))
}

// PkgVersionGT applies the GT predicate on the "pkg_version" field.
func PkgVersionGT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldPkgVersion, v))
}

// PkgVersionGTE applies the GTE predicate on the "pkg_version" field.
func PkgVersionGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldPkgVersion, v))
}

// PkgVersionLT applies the LT predicate on the "pkg_version" field.
func PkgVersionLT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldPkgVersion, v))
}

// PkgVersionLTE applies the LTE predicate on the "pkg_version" field.
func PkgVersionLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldPkgVersion, v))
}

// PkgVersionContains applies the Contains predicate on the "pkg_version" field.
func PkgVersionContains(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContains(FieldPkgVersion, v))
}

// PkgVersionHasPrefix applies the HasPrefix predicate on the "pkg_version" field.
func PkgVersionHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasPrefix(FieldPkgVersion, v))
}

// PkgVersionHasSuffix applies the HasSuffix predicate on the "pkg_version" field.
func PkgVersionHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasSuffix(FieldPkgVersion, v))
}

// PkgVersionIsNil applies the IsNil predicate on the "pkg_version" field.
func PkgVersionIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldPkgVersion))
}

// PkgVersionNotNil applies the NotNil predicate on the "pkg_version" field.
func PkgVersionNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldPkgVersion))
}

// PkgVersionEqualFold applies the EqualFold predicate on the "pkg_version" field.
func PkgVersionEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEqualFold(FieldPkgVersion, v))
}

// PkgVersionContainsFold applies the ContainsFold predicate on the "pkg_version" field.
func PkgVersionContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContainsFold(FieldPkgVersion, v))
}

// PkgTypeEQ applies the EQ predicate on the "pkg_type" field.
func PkgTypeEQ(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPkgType, v))
}

// PkgTypeNEQ applies the NEQ predicate on the "pkg_type" field.
func PkgTypeNEQ(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldPkgType, v))
}

// PkgTypeIn applies the In predicate on the "pkg_type" field.
func PkgTypeIn(vs ...int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldPkgType, vs...))
}

// PkgTypeNotIn applies the NotIn predicate on the "pkg_type" field.
func PkgTypeNotIn(vs ...int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldPkgType, vs...))
}

// PkgTypeGT applies the GT predicate on the "pkg_type" field.
func PkgTypeGT(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldPkgType, v))
}

// PkgTypeGTE applies the GTE predicate on the "pkg_type" field.
func PkgTypeGTE(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldPkgType, v))
}

// PkgTypeLT applies the LT predicate on the "pkg_type" field.
func PkgTypeLT(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldPkgType, v))
}

// PkgTypeLTE applies the LTE predicate on the "pkg_type" field.
func PkgTypeLTE(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldPkgType, v))
}

// PkgTypeIsNil applies the IsNil predicate on the "pkg_type" field.
func PkgTypeIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldPkgType))
}

// PkgTypeNotNil applies the NotNil predicate on the "pkg_type" field.
func PkgTypeNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldPkgType))
}

// PkgKindEQ applies the EQ predicate on the "pkg_kind" field.
func PkgKindEQ(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldPkgKind, v))
}

// PkgKindNEQ applies the NEQ predicate on the "pkg_kind" field.
func PkgKindNEQ(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldPkgKind, v))
}

// PkgKindIn applies the In predicate on the "pkg_kind" field.
func PkgKindIn(vs ...int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldPkgKind, vs...))
}

// PkgKindNotIn applies the NotIn predicate on the "pkg_kind" field.
func PkgKindNotIn(vs ...int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldPkgKind, vs...))
}

// PkgKindGT applies the GT predicate on the "pkg_kind" field.
func PkgKindGT(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldPkgKind, v))
}

// PkgKindGTE applies the GTE predicate on the "pkg_kind" field.
func PkgKindGTE(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldPkgKind, v))
}

// PkgKindLT applies the LT predicate on the "pkg_kind" field.
func PkgKindLT(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldPkgKind, v))
}

// PkgKindLTE applies the LTE predicate on the "pkg_kind" field.
func PkgKindLTE(v int8) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldPkgKind, v))
}

// PkgKindIsNil applies the IsNil predicate on the "pkg_kind" field.
func PkgKindIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldPkgKind))
}

// PkgKindNotNil applies the NotNil predicate on the "pkg_kind" field.
func PkgKindNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldPkgKind))
}

// UploaderEQ applies the EQ predicate on the "uploader" field.
func UploaderEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldUploader, v))
}

// UploaderNEQ applies the NEQ predicate on the "uploader" field.
func UploaderNEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldUploader, v))
}

// UploaderIn applies the In predicate on the "uploader" field.
func UploaderIn(vs ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldUploader, vs...))
}

// UploaderNotIn applies the NotIn predicate on the "uploader" field.
func UploaderNotIn(vs ...int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldUploader, vs...))
}

// UploaderGT applies the GT predicate on the "uploader" field.
func UploaderGT(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldUploader, v))
}

// UploaderGTE applies the GTE predicate on the "uploader" field.
func UploaderGTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldUploader, v))
}

// UploaderLT applies the LT predicate on the "uploader" field.
func UploaderLT(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldUploader, v))
}

// UploaderLTE applies the LTE predicate on the "uploader" field.
func UploaderLTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldUploader, v))
}

// UploaderIsNil applies the IsNil predicate on the "uploader" field.
func UploaderIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldUploader))
}

// UploaderNotNil applies the NotNil predicate on the "uploader" field.
func UploaderNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldUploader))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldHasSuffix(FieldDesc, v))
}

// DescIsNil applies the IsNil predicate on the "desc" field.
func DescIsNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldIsNull(FieldDesc))
}

// DescNotNil applies the NotNil predicate on the "desc" field.
func DescNotNil() predicate.AppPackage {
	return predicate.AppPackage(sql.FieldNotNull(FieldDesc))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(sql.FieldContainsFold(FieldDesc, v))
}

// HasAppInstance applies the HasEdge predicate on the "app_instance" edge.
func HasAppInstance() predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AppInstanceTable, AppInstanceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppInstanceWith applies the HasEdge predicate on the "app_instance" edge with a given conditions (other predicates).
func HasAppInstanceWith(preds ...predicate.AppInstance) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		step := newAppInstanceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppPackage) predicate.AppPackage {
	return predicate.AppPackage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppPackage) predicate.AppPackage {
	return predicate.AppPackage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppPackage) predicate.AppPackage {
	return predicate.AppPackage(sql.NotPredicates(p))
}
