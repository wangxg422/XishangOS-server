// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdept"
)

// 系统部门表
type SysDept struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,string"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"UpdateAt"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt time.Time `json:"deleteAt"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int64 `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int64 `json:"updated_by,omitempty"`
	// DeleteBy holds the value of the "delete_by" field.
	DeleteBy int64 `json:"delete_by,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,string"`
	// Sort holds the value of the "sort" field.
	Sort int `json:"sort,omitempty"`
	// DelFlag holds the value of the "del_flag" field.
	DelFlag int8 `json:"-"`
	// 父级部门id
	ParentID int64 `json:"parentId,string"`
	// 祖先部门列表
	Ancestors string `json:"ancestors"`
	// 部门名称
	DeptName string `json:"deptName"`
	// 部门编码
	DeptCode string `json:"deptCode"`
	// 负责人
	Leader string `json:"leader"`
	// 部门联系电话
	Phone string `json:"phone"`
	// 部门电子邮箱
	Email string `json:"email"`
	// 部门地址
	Address string `json:"address"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysDeptQuery when eager-loading is set.
	Edges        SysDeptEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SysDeptEdges holds the relations/edges for other nodes in the graph.
type SysDeptEdges struct {
	// SysUsers holds the value of the sysUsers edge.
	SysUsers []*SysUser `json:"sysUsers,omitempty"`
	// SysRoles holds the value of the sysRoles edge.
	SysRoles []*SysRole `json:"sysRoles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SysUsersOrErr returns the SysUsers value or an error if the edge
// was not loaded in eager-loading.
func (e SysDeptEdges) SysUsersOrErr() ([]*SysUser, error) {
	if e.loadedTypes[0] {
		return e.SysUsers, nil
	}
	return nil, &NotLoadedError{edge: "sysUsers"}
}

// SysRolesOrErr returns the SysRoles value or an error if the edge
// was not loaded in eager-loading.
func (e SysDeptEdges) SysRolesOrErr() ([]*SysRole, error) {
	if e.loadedTypes[1] {
		return e.SysRoles, nil
	}
	return nil, &NotLoadedError{edge: "sysRoles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysDept) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysdept.FieldID, sysdept.FieldCreatedBy, sysdept.FieldUpdatedBy, sysdept.FieldDeleteBy, sysdept.FieldStatus, sysdept.FieldSort, sysdept.FieldDelFlag, sysdept.FieldParentID:
			values[i] = new(sql.NullInt64)
		case sysdept.FieldRemark, sysdept.FieldAncestors, sysdept.FieldDeptName, sysdept.FieldDeptCode, sysdept.FieldLeader, sysdept.FieldPhone, sysdept.FieldEmail, sysdept.FieldAddress:
			values[i] = new(sql.NullString)
		case sysdept.FieldCreatedAt, sysdept.FieldUpdatedAt, sysdept.FieldDeleteAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysDept fields.
func (sd *SysDept) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysdept.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sd.ID = int64(value.Int64)
		case sysdept.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sd.CreatedAt = value.Time
			}
		case sysdept.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sd.UpdatedAt = value.Time
			}
		case sysdept.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				sd.DeleteAt = value.Time
			}
		case sysdept.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				sd.CreatedBy = value.Int64
			}
		case sysdept.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				sd.UpdatedBy = value.Int64
			}
		case sysdept.FieldDeleteBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_by", values[i])
			} else if value.Valid {
				sd.DeleteBy = value.Int64
			}
		case sysdept.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				sd.Remark = value.String
			}
		case sysdept.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sd.Status = int8(value.Int64)
			}
		case sysdept.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sd.Sort = int(value.Int64)
			}
		case sysdept.FieldDelFlag:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field del_flag", values[i])
			} else if value.Valid {
				sd.DelFlag = int8(value.Int64)
			}
		case sysdept.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				sd.ParentID = value.Int64
			}
		case sysdept.FieldAncestors:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ancestors", values[i])
			} else if value.Valid {
				sd.Ancestors = value.String
			}
		case sysdept.FieldDeptName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dept_name", values[i])
			} else if value.Valid {
				sd.DeptName = value.String
			}
		case sysdept.FieldDeptCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dept_code", values[i])
			} else if value.Valid {
				sd.DeptCode = value.String
			}
		case sysdept.FieldLeader:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field leader", values[i])
			} else if value.Valid {
				sd.Leader = value.String
			}
		case sysdept.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				sd.Phone = value.String
			}
		case sysdept.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				sd.Email = value.String
			}
		case sysdept.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				sd.Address = value.String
			}
		default:
			sd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SysDept.
// This includes values selected through modifiers, order, etc.
func (sd *SysDept) Value(name string) (ent.Value, error) {
	return sd.selectValues.Get(name)
}

// QuerySysUsers queries the "sysUsers" edge of the SysDept entity.
func (sd *SysDept) QuerySysUsers() *SysUserQuery {
	return NewSysDeptClient(sd.config).QuerySysUsers(sd)
}

// QuerySysRoles queries the "sysRoles" edge of the SysDept entity.
func (sd *SysDept) QuerySysRoles() *SysRoleQuery {
	return NewSysDeptClient(sd.config).QuerySysRoles(sd)
}

// Update returns a builder for updating this SysDept.
// Note that you need to call SysDept.Unwrap() before calling this method if this SysDept
// was returned from a transaction, and the transaction was committed or rolled back.
func (sd *SysDept) Update() *SysDeptUpdateOne {
	return NewSysDeptClient(sd.config).UpdateOne(sd)
}

// Unwrap unwraps the SysDept entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sd *SysDept) Unwrap() *SysDept {
	_tx, ok := sd.config.driver.(*txDriver)
	if !ok {
		panic("codegen: SysDept is not a transactional entity")
	}
	sd.config.driver = _tx.drv
	return sd
}

// String implements the fmt.Stringer.
func (sd *SysDept) String() string {
	var builder strings.Builder
	builder.WriteString("SysDept(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sd.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(sd.DeleteAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", sd.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", sd.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("delete_by=")
	builder.WriteString(fmt.Sprintf("%v", sd.DeleteBy))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(sd.Remark)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", sd.Status))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", sd.Sort))
	builder.WriteString(", ")
	builder.WriteString("del_flag=")
	builder.WriteString(fmt.Sprintf("%v", sd.DelFlag))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", sd.ParentID))
	builder.WriteString(", ")
	builder.WriteString("ancestors=")
	builder.WriteString(sd.Ancestors)
	builder.WriteString(", ")
	builder.WriteString("dept_name=")
	builder.WriteString(sd.DeptName)
	builder.WriteString(", ")
	builder.WriteString("dept_code=")
	builder.WriteString(sd.DeptCode)
	builder.WriteString(", ")
	builder.WriteString("leader=")
	builder.WriteString(sd.Leader)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(sd.Phone)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(sd.Email)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(sd.Address)
	builder.WriteByte(')')
	return builder.String()
}

// SysDepts is a parsable slice of SysDept.
type SysDepts []*SysDept
