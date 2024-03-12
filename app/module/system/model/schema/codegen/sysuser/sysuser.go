// Code generated by ent, DO NOT EDIT.

package sysuser

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sysuser type in the database.
	Label = "sys_user"
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
	// FieldDelFlag holds the string denoting the del_flag field in the database.
	FieldDelFlag = "del_flag"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldUserStatus holds the string denoting the user_status field in the database.
	FieldUserStatus = "user_status"
	// FieldDeptID holds the string denoting the dept_id field in the database.
	FieldDeptID = "dept_id"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldLastLoginIP holds the string denoting the last_login_ip field in the database.
	FieldLastLoginIP = "last_login_ip"
	// FieldLastLoginTime holds the string denoting the last_login_time field in the database.
	FieldLastLoginTime = "last_login_time"
	// EdgeSysDept holds the string denoting the sysdept edge name in mutations.
	EdgeSysDept = "sysDept"
	// EdgeSysPosts holds the string denoting the sysposts edge name in mutations.
	EdgeSysPosts = "sysPosts"
	// EdgeSysRoles holds the string denoting the sysroles edge name in mutations.
	EdgeSysRoles = "sysRoles"
	// Table holds the table name of the sysuser in the database.
	Table = "sys_user"
	// SysDeptTable is the table that holds the sysDept relation/edge.
	SysDeptTable = "sys_user"
	// SysDeptInverseTable is the table name for the SysDept entity.
	// It exists in this package in order to avoid circular dependency with the "sysdept" package.
	SysDeptInverseTable = "sys_dept"
	// SysDeptColumn is the table column denoting the sysDept relation/edge.
	SysDeptColumn = "dept_id"
	// SysPostsTable is the table that holds the sysPosts relation/edge. The primary key declared below.
	SysPostsTable = "sys_user_post"
	// SysPostsInverseTable is the table name for the SysPost entity.
	// It exists in this package in order to avoid circular dependency with the "syspost" package.
	SysPostsInverseTable = "sys_post"
	// SysRolesTable is the table that holds the sysRoles relation/edge. The primary key declared below.
	SysRolesTable = "sys_user_role"
	// SysRolesInverseTable is the table name for the SysRole entity.
	// It exists in this package in order to avoid circular dependency with the "sysrole" package.
	SysRolesInverseTable = "sys_role"
)

// Columns holds all SQL columns for sysuser fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleteAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeleteBy,
	FieldRemark,
	FieldDelFlag,
	FieldUserName,
	FieldNickname,
	FieldMobile,
	FieldBirthday,
	FieldPassword,
	FieldSalt,
	FieldEmail,
	FieldSex,
	FieldAvatar,
	FieldUserStatus,
	FieldDeptID,
	FieldAddress,
	FieldDesc,
	FieldLastLoginIP,
	FieldLastLoginTime,
}

var (
	// SysPostsPrimaryKey and SysPostsColumn2 are the table columns denoting the
	// primary key for the sysPosts relation (M2M).
	SysPostsPrimaryKey = []string{"user_id", "post_id"}
	// SysRolesPrimaryKey and SysRolesColumn2 are the table columns denoting the
	// primary key for the sysRoles relation (M2M).
	SysRolesPrimaryKey = []string{"user_id", "role_id"}
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
	// DefaultDelFlag holds the default value on creation for the "del_flag" field.
	DefaultDelFlag int8
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the SysUser queries.
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

// ByDelFlag orders the results by the del_flag field.
func ByDelFlag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelFlag, opts...).ToFunc()
}

// ByUserName orders the results by the user_name field.
func ByUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserName, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByBirthday orders the results by the birthday field.
func ByBirthday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthday, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// BySalt orders the results by the salt field.
func BySalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSalt, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// BySex orders the results by the sex field.
func BySex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSex, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByUserStatus orders the results by the user_status field.
func ByUserStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserStatus, opts...).ToFunc()
}

// ByDeptID orders the results by the dept_id field.
func ByDeptID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeptID, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByLastLoginIP orders the results by the last_login_ip field.
func ByLastLoginIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginIP, opts...).ToFunc()
}

// ByLastLoginTime orders the results by the last_login_time field.
func ByLastLoginTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginTime, opts...).ToFunc()
}

// BySysDeptField orders the results by sysDept field.
func BySysDeptField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSysDeptStep(), sql.OrderByField(field, opts...))
	}
}

// BySysPostsCount orders the results by sysPosts count.
func BySysPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSysPostsStep(), opts...)
	}
}

// BySysPosts orders the results by sysPosts terms.
func BySysPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSysPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySysRolesCount orders the results by sysRoles count.
func BySysRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSysRolesStep(), opts...)
	}
}

// BySysRoles orders the results by sysRoles terms.
func BySysRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSysRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSysDeptStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SysDeptInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SysDeptTable, SysDeptColumn),
	)
}
func newSysPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SysPostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SysPostsTable, SysPostsPrimaryKey...),
	)
}
func newSysRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SysRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SysRolesTable, SysRolesPrimaryKey...),
	)
}
