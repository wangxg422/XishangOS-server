// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysloginlog"
)

// 系统用户登录日志表
type SysLoginLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,string"`
	// 登录账号
	LoginName string `json:"loginName"`
	// 登录IP地址
	IPAddr string `json:"ipAddr"`
	// 登录地点
	LoginLocation string `json:"loginLocation"`
	// 登录浏览器类型
	Browser string `json:"browser"`
	// 登录操作系统
	Os string `json:"os"`
	// 提示消息
	Msg string `json:"msg"`
	// 登录时间
	LoginTime time.Time `json:"loginTime"`
	// 登录是否成功
	LoginSuccess int8 `json:"loginSuccess"`
	// 登录模块
	Module       string `json:"module"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysLoginLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysloginlog.FieldID, sysloginlog.FieldLoginSuccess:
			values[i] = new(sql.NullInt64)
		case sysloginlog.FieldLoginName, sysloginlog.FieldIPAddr, sysloginlog.FieldLoginLocation, sysloginlog.FieldBrowser, sysloginlog.FieldOs, sysloginlog.FieldMsg, sysloginlog.FieldModule:
			values[i] = new(sql.NullString)
		case sysloginlog.FieldLoginTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysLoginLog fields.
func (sll *SysLoginLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysloginlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sll.ID = int64(value.Int64)
		case sysloginlog.FieldLoginName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field login_name", values[i])
			} else if value.Valid {
				sll.LoginName = value.String
			}
		case sysloginlog.FieldIPAddr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_addr", values[i])
			} else if value.Valid {
				sll.IPAddr = value.String
			}
		case sysloginlog.FieldLoginLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field login_location", values[i])
			} else if value.Valid {
				sll.LoginLocation = value.String
			}
		case sysloginlog.FieldBrowser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field browser", values[i])
			} else if value.Valid {
				sll.Browser = value.String
			}
		case sysloginlog.FieldOs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field os", values[i])
			} else if value.Valid {
				sll.Os = value.String
			}
		case sysloginlog.FieldMsg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field msg", values[i])
			} else if value.Valid {
				sll.Msg = value.String
			}
		case sysloginlog.FieldLoginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field login_time", values[i])
			} else if value.Valid {
				sll.LoginTime = value.Time
			}
		case sysloginlog.FieldLoginSuccess:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field login_success", values[i])
			} else if value.Valid {
				sll.LoginSuccess = int8(value.Int64)
			}
		case sysloginlog.FieldModule:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field module", values[i])
			} else if value.Valid {
				sll.Module = value.String
			}
		default:
			sll.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SysLoginLog.
// This includes values selected through modifiers, order, etc.
func (sll *SysLoginLog) Value(name string) (ent.Value, error) {
	return sll.selectValues.Get(name)
}

// Update returns a builder for updating this SysLoginLog.
// Note that you need to call SysLoginLog.Unwrap() before calling this method if this SysLoginLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (sll *SysLoginLog) Update() *SysLoginLogUpdateOne {
	return NewSysLoginLogClient(sll.config).UpdateOne(sll)
}

// Unwrap unwraps the SysLoginLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sll *SysLoginLog) Unwrap() *SysLoginLog {
	_tx, ok := sll.config.driver.(*txDriver)
	if !ok {
		panic("codegen: SysLoginLog is not a transactional entity")
	}
	sll.config.driver = _tx.drv
	return sll
}

// String implements the fmt.Stringer.
func (sll *SysLoginLog) String() string {
	var builder strings.Builder
	builder.WriteString("SysLoginLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sll.ID))
	builder.WriteString("login_name=")
	builder.WriteString(sll.LoginName)
	builder.WriteString(", ")
	builder.WriteString("ip_addr=")
	builder.WriteString(sll.IPAddr)
	builder.WriteString(", ")
	builder.WriteString("login_location=")
	builder.WriteString(sll.LoginLocation)
	builder.WriteString(", ")
	builder.WriteString("browser=")
	builder.WriteString(sll.Browser)
	builder.WriteString(", ")
	builder.WriteString("os=")
	builder.WriteString(sll.Os)
	builder.WriteString(", ")
	builder.WriteString("msg=")
	builder.WriteString(sll.Msg)
	builder.WriteString(", ")
	builder.WriteString("login_time=")
	builder.WriteString(sll.LoginTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("login_success=")
	builder.WriteString(fmt.Sprintf("%v", sll.LoginSuccess))
	builder.WriteString(", ")
	builder.WriteString("module=")
	builder.WriteString(sll.Module)
	builder.WriteByte(')')
	return builder.String()
}

// SysLoginLogs is a parsable slice of SysLoginLog.
type SysLoginLogs []*SysLoginLog
