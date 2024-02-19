// Code generated by ent, DO NOT EDIT.

package sysuser

import (
	"time"

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
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldUserNickname holds the string denoting the user_nickname field in the database.
	FieldUserNickname = "user_nickname"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldUserPassword holds the string denoting the user_password field in the database.
	FieldUserPassword = "user_password"
	// FieldUserSalt holds the string denoting the user_salt field in the database.
	FieldUserSalt = "user_salt"
	// FieldUserEmail holds the string denoting the user_email field in the database.
	FieldUserEmail = "user_email"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldIsAdmin holds the string denoting the is_admin field in the database.
	FieldIsAdmin = "is_admin"
	// FieldUserStatus holds the string denoting the user_status field in the database.
	FieldUserStatus = "user_status"
	// FieldDeptID holds the string denoting the dept_id field in the database.
	FieldDeptID = "dept_id"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldDescribe holds the string denoting the describe field in the database.
	FieldDescribe = "describe"
	// FieldLastLoginIP holds the string denoting the last_login_ip field in the database.
	FieldLastLoginIP = "last_login_ip"
	// FieldLastLoginTime holds the string denoting the last_login_time field in the database.
	FieldLastLoginTime = "last_login_time"
	// EdgeDept holds the string denoting the dept edge name in mutations.
	EdgeDept = "dept"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// Table holds the table name of the sysuser in the database.
	Table = "sys_user"
	// DeptTable is the table that holds the dept relation/edge.
	DeptTable = "sys_user"
	// DeptInverseTable is the table name for the SysDept entity.
	// It exists in this package in order to avoid circular dependency with the "sysdept" package.
	DeptInverseTable = "sys_dept"
	// DeptColumn is the table column denoting the dept relation/edge.
	DeptColumn = "dept_id"
	// PostsTable is the table that holds the posts relation/edge. The primary key declared below.
	PostsTable = "sys_user_post"
	// PostsInverseTable is the table name for the SysPost entity.
	// It exists in this package in order to avoid circular dependency with the "syspost" package.
	PostsInverseTable = "sys_post"
)

// Columns holds all SQL columns for sysuser fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleteAt,
	FieldRemark,
	FieldUserName,
	FieldUserNickname,
	FieldMobile,
	FieldBirthday,
	FieldUserPassword,
	FieldUserSalt,
	FieldUserEmail,
	FieldSex,
	FieldAvatar,
	FieldIsAdmin,
	FieldUserStatus,
	FieldDeptID,
	FieldAddress,
	FieldDescribe,
	FieldLastLoginIP,
	FieldLastLoginTime,
}

var (
	// PostsPrimaryKey and PostsColumn2 are the table columns denoting the
	// primary key for the posts relation (M2M).
	PostsPrimaryKey = []string{"user_id", "post_id"}
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

var (
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

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByUserName orders the results by the user_name field.
func ByUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserName, opts...).ToFunc()
}

// ByUserNickname orders the results by the user_nickname field.
func ByUserNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserNickname, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByBirthday orders the results by the birthday field.
func ByBirthday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthday, opts...).ToFunc()
}

// ByUserPassword orders the results by the user_password field.
func ByUserPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserPassword, opts...).ToFunc()
}

// ByUserSalt orders the results by the user_salt field.
func ByUserSalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserSalt, opts...).ToFunc()
}

// ByUserEmail orders the results by the user_email field.
func ByUserEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserEmail, opts...).ToFunc()
}

// BySex orders the results by the sex field.
func BySex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSex, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByIsAdmin orders the results by the is_admin field.
func ByIsAdmin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAdmin, opts...).ToFunc()
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

// ByDescribe orders the results by the describe field.
func ByDescribe(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescribe, opts...).ToFunc()
}

// ByLastLoginIP orders the results by the last_login_ip field.
func ByLastLoginIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginIP, opts...).ToFunc()
}

// ByLastLoginTime orders the results by the last_login_time field.
func ByLastLoginTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginTime, opts...).ToFunc()
}

// ByDeptField orders the results by dept field.
func ByDeptField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeptStep(), sql.OrderByField(field, opts...))
	}
}

// ByPostsCount orders the results by posts count.
func ByPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPostsStep(), opts...)
	}
}

// ByPosts orders the results by posts terms.
func ByPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDeptStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeptInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DeptTable, DeptColumn),
	)
}
func newPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PostsTable, PostsPrimaryKey...),
	)
}
