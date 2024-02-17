// Code generated by ent, DO NOT EDIT.

package syscasbinrule

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the syscasbinrule type in the database.
	Label = "sys_casbin_rule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPtype holds the string denoting the ptype field in the database.
	FieldPtype = "ptype"
	// FieldV0 holds the string denoting the v0 field in the database.
	FieldV0 = "v0"
	// FieldV1 holds the string denoting the v1 field in the database.
	FieldV1 = "v1"
	// FieldV2 holds the string denoting the v2 field in the database.
	FieldV2 = "v2"
	// FieldV3 holds the string denoting the v3 field in the database.
	FieldV3 = "v3"
	// FieldV4 holds the string denoting the v4 field in the database.
	FieldV4 = "v4"
	// FieldV5 holds the string denoting the v5 field in the database.
	FieldV5 = "v5"
	// Table holds the table name of the syscasbinrule in the database.
	Table = "sys_casbin_rule"
)

// Columns holds all SQL columns for syscasbinrule fields.
var Columns = []string{
	FieldID,
	FieldPtype,
	FieldV0,
	FieldV1,
	FieldV2,
	FieldV3,
	FieldV4,
	FieldV5,
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

// OrderOption defines the ordering options for the SysCasbinRule queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPtype orders the results by the ptype field.
func ByPtype(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPtype, opts...).ToFunc()
}

// ByV0 orders the results by the v0 field.
func ByV0(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV0, opts...).ToFunc()
}

// ByV1 orders the results by the v1 field.
func ByV1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV1, opts...).ToFunc()
}

// ByV2 orders the results by the v2 field.
func ByV2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV2, opts...).ToFunc()
}

// ByV3 orders the results by the v3 field.
func ByV3(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV3, opts...).ToFunc()
}

// ByV4 orders the results by the v4 field.
func ByV4(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV4, opts...).ToFunc()
}

// ByV5 orders the results by the v5 field.
func ByV5(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldV5, opts...).ToFunc()
}
