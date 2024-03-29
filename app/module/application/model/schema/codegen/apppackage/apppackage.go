// Code generated by ent, DO NOT EDIT.

package apppackage

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the apppackage type in the database.
	Label = "app_package"
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
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeleteBy holds the string denoting the delete_by field in the database.
	FieldDeleteBy = "delete_by"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldDelFlag holds the string denoting the del_flag field in the database.
	FieldDelFlag = "del_flag"
	// FieldPkgName holds the string denoting the pkg_name field in the database.
	FieldPkgName = "pkg_name"
	// FieldPkgCode holds the string denoting the pkg_code field in the database.
	FieldPkgCode = "pkg_code"
	// FieldPkgVersion holds the string denoting the pkg_version field in the database.
	FieldPkgVersion = "pkg_version"
	// FieldPkgType holds the string denoting the pkg_type field in the database.
	FieldPkgType = "pkg_type"
	// FieldPkgKind holds the string denoting the pkg_kind field in the database.
	FieldPkgKind = "pkg_kind"
	// FieldUploader holds the string denoting the uploader field in the database.
	FieldUploader = "uploader"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// EdgeAppInstance holds the string denoting the app_instance edge name in mutations.
	EdgeAppInstance = "app_instance"
	// Table holds the table name of the apppackage in the database.
	Table = "app_package"
	// AppInstanceTable is the table that holds the app_instance relation/edge.
	AppInstanceTable = "app_instance"
	// AppInstanceInverseTable is the table name for the AppInstance entity.
	// It exists in this package in order to avoid circular dependency with the "appinstance" package.
	AppInstanceInverseTable = "app_instance"
	// AppInstanceColumn is the table column denoting the app_instance relation/edge.
	AppInstanceColumn = "app_package_app_instance"
)

// Columns holds all SQL columns for apppackage fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleteAt,
	FieldStatus,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeleteBy,
	FieldRemark,
	FieldDelFlag,
	FieldPkgName,
	FieldPkgCode,
	FieldPkgVersion,
	FieldPkgType,
	FieldPkgKind,
	FieldUploader,
	FieldDesc,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// DefaultDelFlag holds the default value on creation for the "del_flag" field.
	DefaultDelFlag int8
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the AppPackage queries.
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

// ByDelFlag orders the results by the del_flag field.
func ByDelFlag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelFlag, opts...).ToFunc()
}

// ByPkgName orders the results by the pkg_name field.
func ByPkgName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPkgName, opts...).ToFunc()
}

// ByPkgCode orders the results by the pkg_code field.
func ByPkgCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPkgCode, opts...).ToFunc()
}

// ByPkgVersion orders the results by the pkg_version field.
func ByPkgVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPkgVersion, opts...).ToFunc()
}

// ByPkgType orders the results by the pkg_type field.
func ByPkgType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPkgType, opts...).ToFunc()
}

// ByPkgKind orders the results by the pkg_kind field.
func ByPkgKind(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPkgKind, opts...).ToFunc()
}

// ByUploader orders the results by the uploader field.
func ByUploader(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUploader, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByAppInstanceCount orders the results by app_instance count.
func ByAppInstanceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAppInstanceStep(), opts...)
	}
}

// ByAppInstance orders the results by app_instance terms.
func ByAppInstance(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppInstanceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAppInstanceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppInstanceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AppInstanceTable, AppInstanceColumn),
	)
}
