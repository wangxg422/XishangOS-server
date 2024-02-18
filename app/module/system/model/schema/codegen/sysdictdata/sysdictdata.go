// Code generated by ent, DO NOT EDIT.

package sysdictdata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sysdictdata type in the database.
	Label = "sys_dict_data"
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
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldDictLabel holds the string denoting the dict_label field in the database.
	FieldDictLabel = "dict_label"
	// FieldDictValue holds the string denoting the dict_value field in the database.
	FieldDictValue = "dict_value"
	// FieldDictType holds the string denoting the dict_type field in the database.
	FieldDictType = "dict_type"
	// FieldCSSClass holds the string denoting the css_class field in the database.
	FieldCSSClass = "css_class"
	// FieldListClass holds the string denoting the list_class field in the database.
	FieldListClass = "list_class"
	// FieldIsDefault holds the string denoting the is_default field in the database.
	FieldIsDefault = "is_default"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the sysdictdata in the database.
	Table = "sys_dict_data"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "sys_dict_data"
	// OwnerInverseTable is the table name for the SysDictType entity.
	// It exists in this package in order to avoid circular dependency with the "sysdicttype" package.
	OwnerInverseTable = "sys_dict_type"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "sys_dict_type_sys_dict_datas"
)

// Columns holds all SQL columns for sysdictdata fields.
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
	FieldSort,
	FieldDictLabel,
	FieldDictValue,
	FieldDictType,
	FieldCSSClass,
	FieldListClass,
	FieldIsDefault,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sys_dict_data"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"sys_dict_type_sys_dict_datas",
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the SysDictData queries.
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

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByDictLabel orders the results by the dict_label field.
func ByDictLabel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDictLabel, opts...).ToFunc()
}

// ByDictValue orders the results by the dict_value field.
func ByDictValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDictValue, opts...).ToFunc()
}

// ByDictType orders the results by the dict_type field.
func ByDictType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDictType, opts...).ToFunc()
}

// ByCSSClass orders the results by the css_class field.
func ByCSSClass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCSSClass, opts...).ToFunc()
}

// ByListClass orders the results by the list_class field.
func ByListClass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldListClass, opts...).ToFunc()
}

// ByIsDefault orders the results by the is_default field.
func ByIsDefault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDefault, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
