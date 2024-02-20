// Code generated by ent, DO NOT EDIT.

package sysmenu

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sysmenu type in the database.
	Label = "sys_menu"
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
	// FieldDelFlag holds the string denoting the del_flag field in the database.
	FieldDelFlag = "del_flag"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldCondition holds the string denoting the condition field in the database.
	FieldCondition = "condition"
	// FieldMenuType holds the string denoting the menu_type field in the database.
	FieldMenuType = "menu_type"
	// FieldWeigh holds the string denoting the weigh field in the database.
	FieldWeigh = "weigh"
	// FieldIsHide holds the string denoting the is_hide field in the database.
	FieldIsHide = "is_hide"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldComponent holds the string denoting the component field in the database.
	FieldComponent = "component"
	// FieldIsLink holds the string denoting the is_link field in the database.
	FieldIsLink = "is_link"
	// FieldModuleType holds the string denoting the module_type field in the database.
	FieldModuleType = "module_type"
	// FieldModelID holds the string denoting the model_id field in the database.
	FieldModelID = "model_id"
	// FieldIsIframe holds the string denoting the is_iframe field in the database.
	FieldIsIframe = "is_iframe"
	// FieldIsCached holds the string denoting the is_cached field in the database.
	FieldIsCached = "is_cached"
	// FieldRedirect holds the string denoting the redirect field in the database.
	FieldRedirect = "redirect"
	// FieldIsAffix holds the string denoting the is_affix field in the database.
	FieldIsAffix = "is_affix"
	// FieldLinkURL holds the string denoting the link_url field in the database.
	FieldLinkURL = "link_url"
	// EdgeSysRoles holds the string denoting the sysroles edge name in mutations.
	EdgeSysRoles = "sysRoles"
	// Table holds the table name of the sysmenu in the database.
	Table = "sys_menu"
	// SysRolesTable is the table that holds the sysRoles relation/edge. The primary key declared below.
	SysRolesTable = "sys_role_menu"
	// SysRolesInverseTable is the table name for the SysRole entity.
	// It exists in this package in order to avoid circular dependency with the "sysrole" package.
	SysRolesInverseTable = "sys_role"
)

// Columns holds all SQL columns for sysmenu fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleteAt,
	FieldRemark,
	FieldDelFlag,
	FieldPid,
	FieldName,
	FieldTitle,
	FieldIcon,
	FieldCondition,
	FieldMenuType,
	FieldWeigh,
	FieldIsHide,
	FieldPath,
	FieldComponent,
	FieldIsLink,
	FieldModuleType,
	FieldModelID,
	FieldIsIframe,
	FieldIsCached,
	FieldRedirect,
	FieldIsAffix,
	FieldLinkURL,
}

var (
	// SysRolesPrimaryKey and SysRolesColumn2 are the table columns denoting the
	// primary key for the sysRoles relation (M2M).
	SysRolesPrimaryKey = []string{"role_id", "menu_id"}
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
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() time.Time
	// UpdateDefaultDeleteAt holds the default value on update for the "delete_at" field.
	UpdateDefaultDeleteAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the SysMenu queries.
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

// ByDelFlag orders the results by the del_flag field.
func ByDelFlag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelFlag, opts...).ToFunc()
}

// ByPid orders the results by the pid field.
func ByPid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPid, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByIcon orders the results by the icon field.
func ByIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIcon, opts...).ToFunc()
}

// ByCondition orders the results by the condition field.
func ByCondition(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCondition, opts...).ToFunc()
}

// ByMenuType orders the results by the menu_type field.
func ByMenuType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMenuType, opts...).ToFunc()
}

// ByWeigh orders the results by the weigh field.
func ByWeigh(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeigh, opts...).ToFunc()
}

// ByIsHide orders the results by the is_hide field.
func ByIsHide(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsHide, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByComponent orders the results by the component field.
func ByComponent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComponent, opts...).ToFunc()
}

// ByIsLink orders the results by the is_link field.
func ByIsLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsLink, opts...).ToFunc()
}

// ByModuleType orders the results by the module_type field.
func ByModuleType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModuleType, opts...).ToFunc()
}

// ByModelID orders the results by the model_id field.
func ByModelID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModelID, opts...).ToFunc()
}

// ByIsIframe orders the results by the is_iframe field.
func ByIsIframe(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsIframe, opts...).ToFunc()
}

// ByIsCached orders the results by the is_cached field.
func ByIsCached(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsCached, opts...).ToFunc()
}

// ByRedirect orders the results by the redirect field.
func ByRedirect(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRedirect, opts...).ToFunc()
}

// ByIsAffix orders the results by the is_affix field.
func ByIsAffix(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAffix, opts...).ToFunc()
}

// ByLinkURL orders the results by the link_url field.
func ByLinkURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLinkURL, opts...).ToFunc()
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
func newSysRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SysRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SysRolesTable, SysRolesPrimaryKey...),
	)
}
