// Code generated by ent, DO NOT EDIT.

package sysuseronline

import (
	"entgo.io/ent/dialect/sql"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldUUID, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldToken, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldCreateTime, v))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldUserName, v))
}

// IPAddr applies equality check predicate on the "ip_addr" field. It's identical to IPAddrEQ.
func IPAddr(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldIPAddr, v))
}

// Browser applies equality check predicate on the "browser" field. It's identical to BrowserEQ.
func Browser(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldBrowser, v))
}

// Os applies equality check predicate on the "os" field. It's identical to OsEQ.
func Os(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldOs, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLTE(FieldUUID, v))
}

// UUIDIsNil applies the IsNil predicate on the "uuid" field.
func UUIDIsNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIsNull(FieldUUID))
}

// UUIDNotNil applies the NotNil predicate on the "uuid" field.
func UUIDNotNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotNull(FieldUUID))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLTE(FieldToken, v))
}

// TokenIsNil applies the IsNil predicate on the "token" field.
func TokenIsNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIsNull(FieldToken))
}

// TokenNotNil applies the NotNil predicate on the "token" field.
func TokenNotNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotNull(FieldToken))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotNull(FieldCreateTime))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLTE(FieldUserName, v))
}

// UserNameIsNil applies the IsNil predicate on the "user_name" field.
func UserNameIsNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIsNull(FieldUserName))
}

// UserNameNotNil applies the NotNil predicate on the "user_name" field.
func UserNameNotNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotNull(FieldUserName))
}

// IPAddrEQ applies the EQ predicate on the "ip_addr" field.
func IPAddrEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldIPAddr, v))
}

// IPAddrNEQ applies the NEQ predicate on the "ip_addr" field.
func IPAddrNEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNEQ(FieldIPAddr, v))
}

// IPAddrIn applies the In predicate on the "ip_addr" field.
func IPAddrIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIn(FieldIPAddr, vs...))
}

// IPAddrNotIn applies the NotIn predicate on the "ip_addr" field.
func IPAddrNotIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotIn(FieldIPAddr, vs...))
}

// IPAddrGT applies the GT predicate on the "ip_addr" field.
func IPAddrGT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGT(FieldIPAddr, v))
}

// IPAddrGTE applies the GTE predicate on the "ip_addr" field.
func IPAddrGTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGTE(FieldIPAddr, v))
}

// IPAddrLT applies the LT predicate on the "ip_addr" field.
func IPAddrLT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLT(FieldIPAddr, v))
}

// IPAddrLTE applies the LTE predicate on the "ip_addr" field.
func IPAddrLTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLTE(FieldIPAddr, v))
}

// IPAddrIsNil applies the IsNil predicate on the "ip_addr" field.
func IPAddrIsNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIsNull(FieldIPAddr))
}

// IPAddrNotNil applies the NotNil predicate on the "ip_addr" field.
func IPAddrNotNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotNull(FieldIPAddr))
}

// BrowserEQ applies the EQ predicate on the "browser" field.
func BrowserEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldBrowser, v))
}

// BrowserNEQ applies the NEQ predicate on the "browser" field.
func BrowserNEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNEQ(FieldBrowser, v))
}

// BrowserIn applies the In predicate on the "browser" field.
func BrowserIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIn(FieldBrowser, vs...))
}

// BrowserNotIn applies the NotIn predicate on the "browser" field.
func BrowserNotIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotIn(FieldBrowser, vs...))
}

// BrowserGT applies the GT predicate on the "browser" field.
func BrowserGT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGT(FieldBrowser, v))
}

// BrowserGTE applies the GTE predicate on the "browser" field.
func BrowserGTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGTE(FieldBrowser, v))
}

// BrowserLT applies the LT predicate on the "browser" field.
func BrowserLT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLT(FieldBrowser, v))
}

// BrowserLTE applies the LTE predicate on the "browser" field.
func BrowserLTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLTE(FieldBrowser, v))
}

// BrowserIsNil applies the IsNil predicate on the "browser" field.
func BrowserIsNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIsNull(FieldBrowser))
}

// BrowserNotNil applies the NotNil predicate on the "browser" field.
func BrowserNotNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotNull(FieldBrowser))
}

// OsEQ applies the EQ predicate on the "os" field.
func OsEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldEQ(FieldOs, v))
}

// OsNEQ applies the NEQ predicate on the "os" field.
func OsNEQ(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNEQ(FieldOs, v))
}

// OsIn applies the In predicate on the "os" field.
func OsIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIn(FieldOs, vs...))
}

// OsNotIn applies the NotIn predicate on the "os" field.
func OsNotIn(vs ...int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotIn(FieldOs, vs...))
}

// OsGT applies the GT predicate on the "os" field.
func OsGT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGT(FieldOs, v))
}

// OsGTE applies the GTE predicate on the "os" field.
func OsGTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldGTE(FieldOs, v))
}

// OsLT applies the LT predicate on the "os" field.
func OsLT(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLT(FieldOs, v))
}

// OsLTE applies the LTE predicate on the "os" field.
func OsLTE(v int64) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldLTE(FieldOs, v))
}

// OsIsNil applies the IsNil predicate on the "os" field.
func OsIsNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldIsNull(FieldOs))
}

// OsNotNil applies the NotNil predicate on the "os" field.
func OsNotNil() predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.FieldNotNull(FieldOs))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SysUserOnline) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SysUserOnline) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SysUserOnline) predicate.SysUserOnline {
	return predicate.SysUserOnline(sql.NotPredicates(p))
}