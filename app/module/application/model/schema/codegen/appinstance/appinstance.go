// Code generated by ent, DO NOT EDIT.

package appinstance

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the appinstance type in the database.
	Label = "app_instance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldDelFlag holds the string denoting the del_flag field in the database.
	FieldDelFlag = "del_flag"
	// FieldInstanceName holds the string denoting the instance_name field in the database.
	FieldInstanceName = "instance_name"
	// FieldInstanceCode holds the string denoting the instance_code field in the database.
	FieldInstanceCode = "instance_code"
	// FieldInstancePackage holds the string denoting the instance_package field in the database.
	FieldInstancePackage = "instance_package"
	// FieldInstanceIcon holds the string denoting the instance_icon field in the database.
	FieldInstanceIcon = "instance_icon"
	// FieldInstanceAddress holds the string denoting the instance_address field in the database.
	FieldInstanceAddress = "instance_address"
	// FieldInstanceType holds the string denoting the instance_type field in the database.
	FieldInstanceType = "instance_type"
	// FieldInstaller holds the string denoting the installer field in the database.
	FieldInstaller = "installer"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// EdgeInstallFrom holds the string denoting the installfrom edge name in mutations.
	EdgeInstallFrom = "installFrom"
	// Table holds the table name of the appinstance in the database.
	Table = "app_instance"
	// InstallFromTable is the table that holds the installFrom relation/edge.
	InstallFromTable = "app_instance"
	// InstallFromInverseTable is the table name for the AppPackage entity.
	// It exists in this package in order to avoid circular dependency with the "apppackage" package.
	InstallFromInverseTable = "app_package"
	// InstallFromColumn is the table column denoting the installFrom relation/edge.
	InstallFromColumn = "app_package_app_instance"
)

// Columns holds all SQL columns for appinstance fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleteAt,
	FieldStatus,
	FieldRemark,
	FieldDelFlag,
	FieldInstanceName,
	FieldInstanceCode,
	FieldInstancePackage,
	FieldInstanceIcon,
	FieldInstanceAddress,
	FieldInstanceType,
	FieldInstaller,
	FieldDesc,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "app_instance"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"app_package_app_instance",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() time.Time
	// UpdateDefaultDeleteAt holds the default value on update for the "delete_at" field.
	UpdateDefaultDeleteAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the AppInstance queries.
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

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByDelFlag orders the results by the del_flag field.
func ByDelFlag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelFlag, opts...).ToFunc()
}

// ByInstanceName orders the results by the instance_name field.
func ByInstanceName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceName, opts...).ToFunc()
}

// ByInstanceCode orders the results by the instance_code field.
func ByInstanceCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceCode, opts...).ToFunc()
}

// ByInstancePackage orders the results by the instance_package field.
func ByInstancePackage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstancePackage, opts...).ToFunc()
}

// ByInstanceIcon orders the results by the instance_icon field.
func ByInstanceIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceIcon, opts...).ToFunc()
}

// ByInstanceAddress orders the results by the instance_address field.
func ByInstanceAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceAddress, opts...).ToFunc()
}

// ByInstanceType orders the results by the instance_type field.
func ByInstanceType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstanceType, opts...).ToFunc()
}

// ByInstaller orders the results by the installer field.
func ByInstaller(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstaller, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByInstallFromField orders the results by installFrom field.
func ByInstallFromField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInstallFromStep(), sql.OrderByField(field, opts...))
	}
}
func newInstallFromStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InstallFromInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InstallFromTable, InstallFromColumn),
	)
}
