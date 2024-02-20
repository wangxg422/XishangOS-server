// Code generated by ent, DO NOT EDIT.

package commonconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/wangxg422/XishangOS-backend/app/module/common/model/schema/codegen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldDeleteAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldStatus, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeleteBy applies equality check predicate on the "delete_by" field. It's identical to DeleteByEQ.
func DeleteBy(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldDeleteBy, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldRemark, v))
}

// ConfigName applies equality check predicate on the "config_name" field. It's identical to ConfigNameEQ.
func ConfigName(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldConfigName, v))
}

// ConfigKey applies equality check predicate on the "config_key" field. It's identical to ConfigKeyEQ.
func ConfigKey(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldConfigKey, v))
}

// ConfigValue applies equality check predicate on the "config_value" field. It's identical to ConfigValueEQ.
func ConfigValue(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldConfigValue, v))
}

// ConfigType applies equality check predicate on the "config_type" field. It's identical to ConfigTypeEQ.
func ConfigType(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldConfigType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldDeleteAt, v))
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldDeleteAt, v))
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldDeleteAt, vs...))
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldDeleteAt, vs...))
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldDeleteAt, v))
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldDeleteAt, v))
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldDeleteAt, v))
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v time.Time) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldDeleteAt, v))
}

// DeleteAtIsNil applies the IsNil predicate on the "delete_at" field.
func DeleteAtIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldDeleteAt))
}

// DeleteAtNotNil applies the NotNil predicate on the "delete_at" field.
func DeleteAtNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldDeleteAt))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldStatus))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldCreatedBy))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldUpdatedBy))
}

// DeleteByEQ applies the EQ predicate on the "delete_by" field.
func DeleteByEQ(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldDeleteBy, v))
}

// DeleteByNEQ applies the NEQ predicate on the "delete_by" field.
func DeleteByNEQ(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldDeleteBy, v))
}

// DeleteByIn applies the In predicate on the "delete_by" field.
func DeleteByIn(vs ...int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldDeleteBy, vs...))
}

// DeleteByNotIn applies the NotIn predicate on the "delete_by" field.
func DeleteByNotIn(vs ...int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldDeleteBy, vs...))
}

// DeleteByGT applies the GT predicate on the "delete_by" field.
func DeleteByGT(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldDeleteBy, v))
}

// DeleteByGTE applies the GTE predicate on the "delete_by" field.
func DeleteByGTE(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldDeleteBy, v))
}

// DeleteByLT applies the LT predicate on the "delete_by" field.
func DeleteByLT(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldDeleteBy, v))
}

// DeleteByLTE applies the LTE predicate on the "delete_by" field.
func DeleteByLTE(v int64) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldDeleteBy, v))
}

// DeleteByIsNil applies the IsNil predicate on the "delete_by" field.
func DeleteByIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldDeleteBy))
}

// DeleteByNotNil applies the NotNil predicate on the "delete_by" field.
func DeleteByNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldDeleteBy))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldContainsFold(FieldRemark, v))
}

// ConfigNameEQ applies the EQ predicate on the "config_name" field.
func ConfigNameEQ(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldConfigName, v))
}

// ConfigNameNEQ applies the NEQ predicate on the "config_name" field.
func ConfigNameNEQ(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldConfigName, v))
}

// ConfigNameIn applies the In predicate on the "config_name" field.
func ConfigNameIn(vs ...string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldConfigName, vs...))
}

// ConfigNameNotIn applies the NotIn predicate on the "config_name" field.
func ConfigNameNotIn(vs ...string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldConfigName, vs...))
}

// ConfigNameGT applies the GT predicate on the "config_name" field.
func ConfigNameGT(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldConfigName, v))
}

// ConfigNameGTE applies the GTE predicate on the "config_name" field.
func ConfigNameGTE(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldConfigName, v))
}

// ConfigNameLT applies the LT predicate on the "config_name" field.
func ConfigNameLT(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldConfigName, v))
}

// ConfigNameLTE applies the LTE predicate on the "config_name" field.
func ConfigNameLTE(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldConfigName, v))
}

// ConfigNameContains applies the Contains predicate on the "config_name" field.
func ConfigNameContains(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldContains(FieldConfigName, v))
}

// ConfigNameHasPrefix applies the HasPrefix predicate on the "config_name" field.
func ConfigNameHasPrefix(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldHasPrefix(FieldConfigName, v))
}

// ConfigNameHasSuffix applies the HasSuffix predicate on the "config_name" field.
func ConfigNameHasSuffix(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldHasSuffix(FieldConfigName, v))
}

// ConfigNameIsNil applies the IsNil predicate on the "config_name" field.
func ConfigNameIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldConfigName))
}

// ConfigNameNotNil applies the NotNil predicate on the "config_name" field.
func ConfigNameNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldConfigName))
}

// ConfigNameEqualFold applies the EqualFold predicate on the "config_name" field.
func ConfigNameEqualFold(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEqualFold(FieldConfigName, v))
}

// ConfigNameContainsFold applies the ContainsFold predicate on the "config_name" field.
func ConfigNameContainsFold(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldContainsFold(FieldConfigName, v))
}

// ConfigKeyEQ applies the EQ predicate on the "config_key" field.
func ConfigKeyEQ(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldConfigKey, v))
}

// ConfigKeyNEQ applies the NEQ predicate on the "config_key" field.
func ConfigKeyNEQ(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldConfigKey, v))
}

// ConfigKeyIn applies the In predicate on the "config_key" field.
func ConfigKeyIn(vs ...string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldConfigKey, vs...))
}

// ConfigKeyNotIn applies the NotIn predicate on the "config_key" field.
func ConfigKeyNotIn(vs ...string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldConfigKey, vs...))
}

// ConfigKeyGT applies the GT predicate on the "config_key" field.
func ConfigKeyGT(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldConfigKey, v))
}

// ConfigKeyGTE applies the GTE predicate on the "config_key" field.
func ConfigKeyGTE(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldConfigKey, v))
}

// ConfigKeyLT applies the LT predicate on the "config_key" field.
func ConfigKeyLT(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldConfigKey, v))
}

// ConfigKeyLTE applies the LTE predicate on the "config_key" field.
func ConfigKeyLTE(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldConfigKey, v))
}

// ConfigKeyContains applies the Contains predicate on the "config_key" field.
func ConfigKeyContains(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldContains(FieldConfigKey, v))
}

// ConfigKeyHasPrefix applies the HasPrefix predicate on the "config_key" field.
func ConfigKeyHasPrefix(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldHasPrefix(FieldConfigKey, v))
}

// ConfigKeyHasSuffix applies the HasSuffix predicate on the "config_key" field.
func ConfigKeyHasSuffix(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldHasSuffix(FieldConfigKey, v))
}

// ConfigKeyIsNil applies the IsNil predicate on the "config_key" field.
func ConfigKeyIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldConfigKey))
}

// ConfigKeyNotNil applies the NotNil predicate on the "config_key" field.
func ConfigKeyNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldConfigKey))
}

// ConfigKeyEqualFold applies the EqualFold predicate on the "config_key" field.
func ConfigKeyEqualFold(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEqualFold(FieldConfigKey, v))
}

// ConfigKeyContainsFold applies the ContainsFold predicate on the "config_key" field.
func ConfigKeyContainsFold(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldContainsFold(FieldConfigKey, v))
}

// ConfigValueEQ applies the EQ predicate on the "config_value" field.
func ConfigValueEQ(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldConfigValue, v))
}

// ConfigValueNEQ applies the NEQ predicate on the "config_value" field.
func ConfigValueNEQ(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldConfigValue, v))
}

// ConfigValueIn applies the In predicate on the "config_value" field.
func ConfigValueIn(vs ...string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldConfigValue, vs...))
}

// ConfigValueNotIn applies the NotIn predicate on the "config_value" field.
func ConfigValueNotIn(vs ...string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldConfigValue, vs...))
}

// ConfigValueGT applies the GT predicate on the "config_value" field.
func ConfigValueGT(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldConfigValue, v))
}

// ConfigValueGTE applies the GTE predicate on the "config_value" field.
func ConfigValueGTE(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldConfigValue, v))
}

// ConfigValueLT applies the LT predicate on the "config_value" field.
func ConfigValueLT(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldConfigValue, v))
}

// ConfigValueLTE applies the LTE predicate on the "config_value" field.
func ConfigValueLTE(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldConfigValue, v))
}

// ConfigValueContains applies the Contains predicate on the "config_value" field.
func ConfigValueContains(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldContains(FieldConfigValue, v))
}

// ConfigValueHasPrefix applies the HasPrefix predicate on the "config_value" field.
func ConfigValueHasPrefix(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldHasPrefix(FieldConfigValue, v))
}

// ConfigValueHasSuffix applies the HasSuffix predicate on the "config_value" field.
func ConfigValueHasSuffix(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldHasSuffix(FieldConfigValue, v))
}

// ConfigValueIsNil applies the IsNil predicate on the "config_value" field.
func ConfigValueIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldConfigValue))
}

// ConfigValueNotNil applies the NotNil predicate on the "config_value" field.
func ConfigValueNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldConfigValue))
}

// ConfigValueEqualFold applies the EqualFold predicate on the "config_value" field.
func ConfigValueEqualFold(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEqualFold(FieldConfigValue, v))
}

// ConfigValueContainsFold applies the ContainsFold predicate on the "config_value" field.
func ConfigValueContainsFold(v string) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldContainsFold(FieldConfigValue, v))
}

// ConfigTypeEQ applies the EQ predicate on the "config_type" field.
func ConfigTypeEQ(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldEQ(FieldConfigType, v))
}

// ConfigTypeNEQ applies the NEQ predicate on the "config_type" field.
func ConfigTypeNEQ(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNEQ(FieldConfigType, v))
}

// ConfigTypeIn applies the In predicate on the "config_type" field.
func ConfigTypeIn(vs ...int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIn(FieldConfigType, vs...))
}

// ConfigTypeNotIn applies the NotIn predicate on the "config_type" field.
func ConfigTypeNotIn(vs ...int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotIn(FieldConfigType, vs...))
}

// ConfigTypeGT applies the GT predicate on the "config_type" field.
func ConfigTypeGT(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGT(FieldConfigType, v))
}

// ConfigTypeGTE applies the GTE predicate on the "config_type" field.
func ConfigTypeGTE(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldGTE(FieldConfigType, v))
}

// ConfigTypeLT applies the LT predicate on the "config_type" field.
func ConfigTypeLT(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLT(FieldConfigType, v))
}

// ConfigTypeLTE applies the LTE predicate on the "config_type" field.
func ConfigTypeLTE(v int8) predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldLTE(FieldConfigType, v))
}

// ConfigTypeIsNil applies the IsNil predicate on the "config_type" field.
func ConfigTypeIsNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldIsNull(FieldConfigType))
}

// ConfigTypeNotNil applies the NotNil predicate on the "config_type" field.
func ConfigTypeNotNil() predicate.CommonConfig {
	return predicate.CommonConfig(sql.FieldNotNull(FieldConfigType))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CommonConfig) predicate.CommonConfig {
	return predicate.CommonConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CommonConfig) predicate.CommonConfig {
	return predicate.CommonConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CommonConfig) predicate.CommonConfig {
	return predicate.CommonConfig(sql.NotPredicates(p))
}