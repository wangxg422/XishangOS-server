// Code generated by ent, DO NOT EDIT.

package sysoperlog

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the sysoperlog type in the database.
	Label = "sys_oper_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldBusinessType holds the string denoting the business_type field in the database.
	FieldBusinessType = "business_type"
	// FieldMethod holds the string denoting the method field in the database.
	FieldMethod = "method"
	// FieldRequestMethod holds the string denoting the request_method field in the database.
	FieldRequestMethod = "request_method"
	// FieldOperType holds the string denoting the oper_type field in the database.
	FieldOperType = "oper_type"
	// FieldOperName holds the string denoting the oper_name field in the database.
	FieldOperName = "oper_name"
	// FieldDeptName holds the string denoting the dept_name field in the database.
	FieldDeptName = "dept_name"
	// FieldOperURL holds the string denoting the oper_url field in the database.
	FieldOperURL = "oper_url"
	// FieldOperIP holds the string denoting the oper_ip field in the database.
	FieldOperIP = "oper_ip"
	// FieldOperLocation holds the string denoting the oper_location field in the database.
	FieldOperLocation = "oper_location"
	// FieldOperParam holds the string denoting the oper_param field in the database.
	FieldOperParam = "oper_param"
	// FieldErrorMsg holds the string denoting the error_msg field in the database.
	FieldErrorMsg = "error_msg"
	// FieldOperTime holds the string denoting the oper_time field in the database.
	FieldOperTime = "oper_time"
	// Table holds the table name of the sysoperlog in the database.
	Table = "sys_oper_log"
)

// Columns holds all SQL columns for sysoperlog fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldBusinessType,
	FieldMethod,
	FieldRequestMethod,
	FieldOperType,
	FieldOperName,
	FieldDeptName,
	FieldOperURL,
	FieldOperIP,
	FieldOperLocation,
	FieldOperParam,
	FieldErrorMsg,
	FieldOperTime,
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

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

// OrderOption defines the ordering options for the SysOperLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByBusinessType orders the results by the business_type field.
func ByBusinessType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessType, opts...).ToFunc()
}

// ByMethod orders the results by the method field.
func ByMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMethod, opts...).ToFunc()
}

// ByRequestMethod orders the results by the request_method field.
func ByRequestMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestMethod, opts...).ToFunc()
}

// ByOperType orders the results by the oper_type field.
func ByOperType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperType, opts...).ToFunc()
}

// ByOperName orders the results by the oper_name field.
func ByOperName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperName, opts...).ToFunc()
}

// ByDeptName orders the results by the dept_name field.
func ByDeptName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeptName, opts...).ToFunc()
}

// ByOperURL orders the results by the oper_url field.
func ByOperURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperURL, opts...).ToFunc()
}

// ByOperIP orders the results by the oper_ip field.
func ByOperIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperIP, opts...).ToFunc()
}

// ByOperLocation orders the results by the oper_location field.
func ByOperLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperLocation, opts...).ToFunc()
}

// ByOperParam orders the results by the oper_param field.
func ByOperParam(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperParam, opts...).ToFunc()
}

// ByErrorMsg orders the results by the error_msg field.
func ByErrorMsg(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldErrorMsg, opts...).ToFunc()
}

// ByOperTime orders the results by the oper_time field.
func ByOperTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperTime, opts...).ToFunc()
}
