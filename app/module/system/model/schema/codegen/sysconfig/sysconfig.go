// Code generated by ent, DO NOT EDIT.

package sysconfig

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the sysconfig type in the database.
	Label = "sys_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeleteBy holds the string denoting the delete_by field in the database.
	FieldDeleteBy = "delete_by"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDelFlag holds the string denoting the del_flag field in the database.
	FieldDelFlag = "del_flag"
	// FieldConfigName holds the string denoting the config_name field in the database.
	FieldConfigName = "config_name"
	// FieldConfigKey holds the string denoting the config_key field in the database.
	FieldConfigKey = "config_key"
	// FieldConfigValue holds the string denoting the config_value field in the database.
	FieldConfigValue = "config_value"
	// FieldConfigType holds the string denoting the config_type field in the database.
	FieldConfigType = "config_type"
	// Table holds the table name of the sysconfig in the database.
	Table = "sys_config"
)

// Columns holds all SQL columns for sysconfig fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleteAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeleteBy,
	FieldRemark,
	FieldStatus,
	FieldDelFlag,
	FieldConfigName,
	FieldConfigKey,
	FieldConfigValue,
	FieldConfigType,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// DefaultDelFlag holds the default value on creation for the "del_flag" field.
	DefaultDelFlag int8
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the SysConfig queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeleteAt orders the results by the delete_at field.
func ByDeleteAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeleteBy orders the results by the delete_by field.
func ByDeleteBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteBy, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByDelFlag orders the results by the del_flag field.
func ByDelFlag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelFlag, opts...).ToFunc()
}

// ByConfigName orders the results by the config_name field.
func ByConfigName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConfigName, opts...).ToFunc()
}

// ByConfigKey orders the results by the config_key field.
func ByConfigKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConfigKey, opts...).ToFunc()
}

// ByConfigValue orders the results by the config_value field.
func ByConfigValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConfigValue, opts...).ToFunc()
}

// ByConfigType orders the results by the config_type field.
func ByConfigType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConfigType, opts...).ToFunc()
}
