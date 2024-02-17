// Code generated by ent, DO NOT EDIT.

package appinstance

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldDeleteAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldStatus, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldRemark, v))
}

// InstanceName applies equality check predicate on the "instance_name" field. It's identical to InstanceNameEQ.
func InstanceName(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstanceName, v))
}

// InstanceCode applies equality check predicate on the "instance_code" field. It's identical to InstanceCodeEQ.
func InstanceCode(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstanceCode, v))
}

// InstancePackage applies equality check predicate on the "instance_package" field. It's identical to InstancePackageEQ.
func InstancePackage(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstancePackage, v))
}

// InstanceIcon applies equality check predicate on the "instance_icon" field. It's identical to InstanceIconEQ.
func InstanceIcon(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstanceIcon, v))
}

// InstanceAddress applies equality check predicate on the "instance_address" field. It's identical to InstanceAddressEQ.
func InstanceAddress(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstanceAddress, v))
}

// InstanceType applies equality check predicate on the "instance_type" field. It's identical to InstanceTypeEQ.
func InstanceType(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstanceType, v))
}

// Installer applies equality check predicate on the "installer" field. It's identical to InstallerEQ.
func Installer(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstaller, v))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldDesc, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldDeleteAt, v))
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldDeleteAt, v))
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldDeleteAt, vs...))
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldDeleteAt, vs...))
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldDeleteAt, v))
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldDeleteAt, v))
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldDeleteAt, v))
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v time.Time) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldDeleteAt, v))
}

// DeleteAtIsNil applies the IsNil predicate on the "delete_at" field.
func DeleteAtIsNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIsNull(FieldDeleteAt))
}

// DeleteAtNotNil applies the NotNil predicate on the "delete_at" field.
func DeleteAtNotNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotNull(FieldDeleteAt))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotNull(FieldStatus))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotNull(FieldRemark))
}

// InstanceNameEQ applies the EQ predicate on the "instance_name" field.
func InstanceNameEQ(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstanceName, v))
}

// InstanceNameNEQ applies the NEQ predicate on the "instance_name" field.
func InstanceNameNEQ(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldInstanceName, v))
}

// InstanceNameIn applies the In predicate on the "instance_name" field.
func InstanceNameIn(vs ...string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldInstanceName, vs...))
}

// InstanceNameNotIn applies the NotIn predicate on the "instance_name" field.
func InstanceNameNotIn(vs ...string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldInstanceName, vs...))
}

// InstanceNameGT applies the GT predicate on the "instance_name" field.
func InstanceNameGT(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldInstanceName, v))
}

// InstanceNameGTE applies the GTE predicate on the "instance_name" field.
func InstanceNameGTE(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldInstanceName, v))
}

// InstanceNameLT applies the LT predicate on the "instance_name" field.
func InstanceNameLT(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldInstanceName, v))
}

// InstanceNameLTE applies the LTE predicate on the "instance_name" field.
func InstanceNameLTE(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldInstanceName, v))
}

// InstanceNameContains applies the Contains predicate on the "instance_name" field.
func InstanceNameContains(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldContains(FieldInstanceName, v))
}

// InstanceNameHasPrefix applies the HasPrefix predicate on the "instance_name" field.
func InstanceNameHasPrefix(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldHasPrefix(FieldInstanceName, v))
}

// InstanceNameHasSuffix applies the HasSuffix predicate on the "instance_name" field.
func InstanceNameHasSuffix(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldHasSuffix(FieldInstanceName, v))
}

// InstanceNameEqualFold applies the EqualFold predicate on the "instance_name" field.
func InstanceNameEqualFold(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEqualFold(FieldInstanceName, v))
}

// InstanceNameContainsFold applies the ContainsFold predicate on the "instance_name" field.
func InstanceNameContainsFold(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldContainsFold(FieldInstanceName, v))
}

// InstanceCodeEQ applies the EQ predicate on the "instance_code" field.
func InstanceCodeEQ(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstanceCode, v))
}

// InstanceCodeNEQ applies the NEQ predicate on the "instance_code" field.
func InstanceCodeNEQ(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldInstanceCode, v))
}

// InstanceCodeIn applies the In predicate on the "instance_code" field.
func InstanceCodeIn(vs ...string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldInstanceCode, vs...))
}

// InstanceCodeNotIn applies the NotIn predicate on the "instance_code" field.
func InstanceCodeNotIn(vs ...string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldInstanceCode, vs...))
}

// InstanceCodeGT applies the GT predicate on the "instance_code" field.
func InstanceCodeGT(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldInstanceCode, v))
}

// InstanceCodeGTE applies the GTE predicate on the "instance_code" field.
func InstanceCodeGTE(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldInstanceCode, v))
}

// InstanceCodeLT applies the LT predicate on the "instance_code" field.
func InstanceCodeLT(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldInstanceCode, v))
}

// InstanceCodeLTE applies the LTE predicate on the "instance_code" field.
func InstanceCodeLTE(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldInstanceCode, v))
}

// InstanceCodeContains applies the Contains predicate on the "instance_code" field.
func InstanceCodeContains(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldContains(FieldInstanceCode, v))
}

// InstanceCodeHasPrefix applies the HasPrefix predicate on the "instance_code" field.
func InstanceCodeHasPrefix(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldHasPrefix(FieldInstanceCode, v))
}

// InstanceCodeHasSuffix applies the HasSuffix predicate on the "instance_code" field.
func InstanceCodeHasSuffix(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldHasSuffix(FieldInstanceCode, v))
}

// InstanceCodeEqualFold applies the EqualFold predicate on the "instance_code" field.
func InstanceCodeEqualFold(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEqualFold(FieldInstanceCode, v))
}

// InstanceCodeContainsFold applies the ContainsFold predicate on the "instance_code" field.
func InstanceCodeContainsFold(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldContainsFold(FieldInstanceCode, v))
}

// InstancePackageEQ applies the EQ predicate on the "instance_package" field.
func InstancePackageEQ(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstancePackage, v))
}

// InstancePackageNEQ applies the NEQ predicate on the "instance_package" field.
func InstancePackageNEQ(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldInstancePackage, v))
}

// InstancePackageIn applies the In predicate on the "instance_package" field.
func InstancePackageIn(vs ...int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldInstancePackage, vs...))
}

// InstancePackageNotIn applies the NotIn predicate on the "instance_package" field.
func InstancePackageNotIn(vs ...int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldInstancePackage, vs...))
}

// InstancePackageGT applies the GT predicate on the "instance_package" field.
func InstancePackageGT(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldInstancePackage, v))
}

// InstancePackageGTE applies the GTE predicate on the "instance_package" field.
func InstancePackageGTE(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldInstancePackage, v))
}

// InstancePackageLT applies the LT predicate on the "instance_package" field.
func InstancePackageLT(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldInstancePackage, v))
}

// InstancePackageLTE applies the LTE predicate on the "instance_package" field.
func InstancePackageLTE(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldInstancePackage, v))
}

// InstanceIconEQ applies the EQ predicate on the "instance_icon" field.
func InstanceIconEQ(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstanceIcon, v))
}

// InstanceIconNEQ applies the NEQ predicate on the "instance_icon" field.
func InstanceIconNEQ(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldInstanceIcon, v))
}

// InstanceIconIn applies the In predicate on the "instance_icon" field.
func InstanceIconIn(vs ...string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldInstanceIcon, vs...))
}

// InstanceIconNotIn applies the NotIn predicate on the "instance_icon" field.
func InstanceIconNotIn(vs ...string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldInstanceIcon, vs...))
}

// InstanceIconGT applies the GT predicate on the "instance_icon" field.
func InstanceIconGT(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldInstanceIcon, v))
}

// InstanceIconGTE applies the GTE predicate on the "instance_icon" field.
func InstanceIconGTE(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldInstanceIcon, v))
}

// InstanceIconLT applies the LT predicate on the "instance_icon" field.
func InstanceIconLT(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldInstanceIcon, v))
}

// InstanceIconLTE applies the LTE predicate on the "instance_icon" field.
func InstanceIconLTE(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldInstanceIcon, v))
}

// InstanceIconContains applies the Contains predicate on the "instance_icon" field.
func InstanceIconContains(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldContains(FieldInstanceIcon, v))
}

// InstanceIconHasPrefix applies the HasPrefix predicate on the "instance_icon" field.
func InstanceIconHasPrefix(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldHasPrefix(FieldInstanceIcon, v))
}

// InstanceIconHasSuffix applies the HasSuffix predicate on the "instance_icon" field.
func InstanceIconHasSuffix(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldHasSuffix(FieldInstanceIcon, v))
}

// InstanceIconIsNil applies the IsNil predicate on the "instance_icon" field.
func InstanceIconIsNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIsNull(FieldInstanceIcon))
}

// InstanceIconNotNil applies the NotNil predicate on the "instance_icon" field.
func InstanceIconNotNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotNull(FieldInstanceIcon))
}

// InstanceIconEqualFold applies the EqualFold predicate on the "instance_icon" field.
func InstanceIconEqualFold(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEqualFold(FieldInstanceIcon, v))
}

// InstanceIconContainsFold applies the ContainsFold predicate on the "instance_icon" field.
func InstanceIconContainsFold(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldContainsFold(FieldInstanceIcon, v))
}

// InstanceAddressEQ applies the EQ predicate on the "instance_address" field.
func InstanceAddressEQ(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstanceAddress, v))
}

// InstanceAddressNEQ applies the NEQ predicate on the "instance_address" field.
func InstanceAddressNEQ(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldInstanceAddress, v))
}

// InstanceAddressIn applies the In predicate on the "instance_address" field.
func InstanceAddressIn(vs ...string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldInstanceAddress, vs...))
}

// InstanceAddressNotIn applies the NotIn predicate on the "instance_address" field.
func InstanceAddressNotIn(vs ...string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldInstanceAddress, vs...))
}

// InstanceAddressGT applies the GT predicate on the "instance_address" field.
func InstanceAddressGT(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldInstanceAddress, v))
}

// InstanceAddressGTE applies the GTE predicate on the "instance_address" field.
func InstanceAddressGTE(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldInstanceAddress, v))
}

// InstanceAddressLT applies the LT predicate on the "instance_address" field.
func InstanceAddressLT(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldInstanceAddress, v))
}

// InstanceAddressLTE applies the LTE predicate on the "instance_address" field.
func InstanceAddressLTE(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldInstanceAddress, v))
}

// InstanceAddressContains applies the Contains predicate on the "instance_address" field.
func InstanceAddressContains(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldContains(FieldInstanceAddress, v))
}

// InstanceAddressHasPrefix applies the HasPrefix predicate on the "instance_address" field.
func InstanceAddressHasPrefix(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldHasPrefix(FieldInstanceAddress, v))
}

// InstanceAddressHasSuffix applies the HasSuffix predicate on the "instance_address" field.
func InstanceAddressHasSuffix(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldHasSuffix(FieldInstanceAddress, v))
}

// InstanceAddressIsNil applies the IsNil predicate on the "instance_address" field.
func InstanceAddressIsNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIsNull(FieldInstanceAddress))
}

// InstanceAddressNotNil applies the NotNil predicate on the "instance_address" field.
func InstanceAddressNotNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotNull(FieldInstanceAddress))
}

// InstanceAddressEqualFold applies the EqualFold predicate on the "instance_address" field.
func InstanceAddressEqualFold(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEqualFold(FieldInstanceAddress, v))
}

// InstanceAddressContainsFold applies the ContainsFold predicate on the "instance_address" field.
func InstanceAddressContainsFold(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldContainsFold(FieldInstanceAddress, v))
}

// InstanceTypeEQ applies the EQ predicate on the "instance_type" field.
func InstanceTypeEQ(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstanceType, v))
}

// InstanceTypeNEQ applies the NEQ predicate on the "instance_type" field.
func InstanceTypeNEQ(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldInstanceType, v))
}

// InstanceTypeIn applies the In predicate on the "instance_type" field.
func InstanceTypeIn(vs ...int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldInstanceType, vs...))
}

// InstanceTypeNotIn applies the NotIn predicate on the "instance_type" field.
func InstanceTypeNotIn(vs ...int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldInstanceType, vs...))
}

// InstanceTypeGT applies the GT predicate on the "instance_type" field.
func InstanceTypeGT(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldInstanceType, v))
}

// InstanceTypeGTE applies the GTE predicate on the "instance_type" field.
func InstanceTypeGTE(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldInstanceType, v))
}

// InstanceTypeLT applies the LT predicate on the "instance_type" field.
func InstanceTypeLT(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldInstanceType, v))
}

// InstanceTypeLTE applies the LTE predicate on the "instance_type" field.
func InstanceTypeLTE(v int8) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldInstanceType, v))
}

// InstanceTypeIsNil applies the IsNil predicate on the "instance_type" field.
func InstanceTypeIsNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIsNull(FieldInstanceType))
}

// InstanceTypeNotNil applies the NotNil predicate on the "instance_type" field.
func InstanceTypeNotNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotNull(FieldInstanceType))
}

// InstallerEQ applies the EQ predicate on the "installer" field.
func InstallerEQ(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldInstaller, v))
}

// InstallerNEQ applies the NEQ predicate on the "installer" field.
func InstallerNEQ(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldInstaller, v))
}

// InstallerIn applies the In predicate on the "installer" field.
func InstallerIn(vs ...int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldInstaller, vs...))
}

// InstallerNotIn applies the NotIn predicate on the "installer" field.
func InstallerNotIn(vs ...int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldInstaller, vs...))
}

// InstallerGT applies the GT predicate on the "installer" field.
func InstallerGT(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldInstaller, v))
}

// InstallerGTE applies the GTE predicate on the "installer" field.
func InstallerGTE(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldInstaller, v))
}

// InstallerLT applies the LT predicate on the "installer" field.
func InstallerLT(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldInstaller, v))
}

// InstallerLTE applies the LTE predicate on the "installer" field.
func InstallerLTE(v int64) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldInstaller, v))
}

// InstallerIsNil applies the IsNil predicate on the "installer" field.
func InstallerIsNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIsNull(FieldInstaller))
}

// InstallerNotNil applies the NotNil predicate on the "installer" field.
func InstallerNotNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotNull(FieldInstaller))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldHasSuffix(FieldDesc, v))
}

// DescIsNil applies the IsNil predicate on the "desc" field.
func DescIsNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldIsNull(FieldDesc))
}

// DescNotNil applies the NotNil predicate on the "desc" field.
func DescNotNil() predicate.AppInstance {
	return predicate.AppInstance(sql.FieldNotNull(FieldDesc))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.AppInstance {
	return predicate.AppInstance(sql.FieldContainsFold(FieldDesc, v))
}

// HasInstallFrom applies the HasEdge predicate on the "installFrom" edge.
func HasInstallFrom() predicate.AppInstance {
	return predicate.AppInstance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstallFromTable, InstallFromColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstallFromWith applies the HasEdge predicate on the "installFrom" edge with a given conditions (other predicates).
func HasInstallFromWith(preds ...predicate.AppPackage) predicate.AppInstance {
	return predicate.AppInstance(func(s *sql.Selector) {
		step := newInstallFromStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppInstance) predicate.AppInstance {
	return predicate.AppInstance(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppInstance) predicate.AppInstance {
	return predicate.AppInstance(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppInstance) predicate.AppInstance {
	return predicate.AppInstance(sql.NotPredicates(p))
}
