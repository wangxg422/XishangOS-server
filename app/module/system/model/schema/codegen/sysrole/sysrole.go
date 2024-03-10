// Code generated by ent, DO NOT EDIT.

package sysrole

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sysrole type in the database.
	Label = "sys_role"
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
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldDelFlag holds the string denoting the del_flag field in the database.
	FieldDelFlag = "del_flag"
	// FieldRoleName holds the string denoting the role_name field in the database.
	FieldRoleName = "role_name"
	// FieldRoleCode holds the string denoting the role_code field in the database.
	FieldRoleCode = "role_code"
	// FieldDataScope holds the string denoting the data_scope field in the database.
	FieldDataScope = "data_scope"
	// EdgeSysDepts holds the string denoting the sysdepts edge name in mutations.
	EdgeSysDepts = "sysDepts"
	// EdgeSysUsers holds the string denoting the sysusers edge name in mutations.
	EdgeSysUsers = "sysUsers"
	// EdgeSysMenus holds the string denoting the sysmenus edge name in mutations.
	EdgeSysMenus = "sysMenus"
	// Table holds the table name of the sysrole in the database.
	Table = "sys_role"
	// SysDeptsTable is the table that holds the sysDepts relation/edge. The primary key declared below.
	SysDeptsTable = "sys_role_dept"
	// SysDeptsInverseTable is the table name for the SysDept entity.
	// It exists in this package in order to avoid circular dependency with the "sysdept" package.
	SysDeptsInverseTable = "sys_dept"
	// SysUsersTable is the table that holds the sysUsers relation/edge. The primary key declared below.
	SysUsersTable = "sys_user_role"
	// SysUsersInverseTable is the table name for the SysUser entity.
	// It exists in this package in order to avoid circular dependency with the "sysuser" package.
	SysUsersInverseTable = "sys_user"
	// SysMenusTable is the table that holds the sysMenus relation/edge. The primary key declared below.
	SysMenusTable = "sys_role_menu"
	// SysMenusInverseTable is the table name for the SysMenu entity.
	// It exists in this package in order to avoid circular dependency with the "sysmenu" package.
	SysMenusInverseTable = "sys_menu"
)

// Columns holds all SQL columns for sysrole fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleteAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeleteBy,
	FieldStatus,
	FieldSort,
	FieldRemark,
	FieldDelFlag,
	FieldRoleName,
	FieldRoleCode,
	FieldDataScope,
}

var (
	// SysDeptsPrimaryKey and SysDeptsColumn2 are the table columns denoting the
	// primary key for the sysDepts relation (M2M).
	SysDeptsPrimaryKey = []string{"role_id", "dept_id"}
	// SysUsersPrimaryKey and SysUsersColumn2 are the table columns denoting the
	// primary key for the sysUsers relation (M2M).
	SysUsersPrimaryKey = []string{"user_id", "role_id"}
	// SysMenusPrimaryKey and SysMenusColumn2 are the table columns denoting the
	// primary key for the sysMenus relation (M2M).
	SysMenusPrimaryKey = []string{"role_id", "menu_id"}
)

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
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int
	// DefaultDelFlag holds the default value on creation for the "del_flag" field.
	DefaultDelFlag int8
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the SysRole queries.
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

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByDelFlag orders the results by the del_flag field.
func ByDelFlag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelFlag, opts...).ToFunc()
}

// ByRoleName orders the results by the role_name field.
func ByRoleName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleName, opts...).ToFunc()
}

// ByRoleCode orders the results by the role_code field.
func ByRoleCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleCode, opts...).ToFunc()
}

// ByDataScope orders the results by the data_scope field.
func ByDataScope(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDataScope, opts...).ToFunc()
}

// BySysDeptsCount orders the results by sysDepts count.
func BySysDeptsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSysDeptsStep(), opts...)
	}
}

// BySysDepts orders the results by sysDepts terms.
func BySysDepts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSysDeptsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySysUsersCount orders the results by sysUsers count.
func BySysUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSysUsersStep(), opts...)
	}
}

// BySysUsers orders the results by sysUsers terms.
func BySysUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSysUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySysMenusCount orders the results by sysMenus count.
func BySysMenusCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSysMenusStep(), opts...)
	}
}

// BySysMenus orders the results by sysMenus terms.
func BySysMenus(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSysMenusStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSysDeptsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SysDeptsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SysDeptsTable, SysDeptsPrimaryKey...),
	)
}
func newSysUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SysUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SysUsersTable, SysUsersPrimaryKey...),
	)
}
func newSysMenusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SysMenusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SysMenusTable, SysMenusPrimaryKey...),
	)
}
