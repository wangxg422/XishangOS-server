// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysmenu"
)

// 系统菜单权限表
type SysMenu struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt time.Time `json:"delete_at,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// 父ID
	Pid int64 `json:"pid,omitempty"`
	// 菜单名称
	Name string `json:"name,omitempty"`
	// 菜单标题
	Title string `json:"title,omitempty"`
	// 菜单图标
	Icon string `json:"icon,omitempty"`
	// 条件
	Condition string `json:"condition,omitempty"`
	// 菜单类型(0目录,1菜单,2按钮)
	MenuType int8 `json:"menu_type,omitempty"`
	// 菜单权重
	Weigh int64 `json:"weigh,omitempty"`
	// 显示状态
	IsHide int8 `json:"is_hide,omitempty"`
	// 路由地址
	Path string `json:"path,omitempty"`
	// 组件路径
	Component string `json:"component,omitempty"`
	// 是否是外部链接(1是,0否)
	IsLink int8 `json:"is_link,omitempty"`
	// 所属模块
	ModuleType string `json:"module_type,omitempty"`
	// 模型id
	ModelID int64 `json:"model_id,omitempty"`
	// 是否是内嵌iframe(1是,0否)
	IsIframe int8 `json:"is_iframe,omitempty"`
	// 是否缓存(1是,0否)
	IsCached int8 `json:"is_cached,omitempty"`
	// 路由重定向地址
	Redirect string `json:"redirect,omitempty"`
	// 是否固定(1是,0否)
	IsAffix int8 `json:"is_affix,omitempty"`
	// 链接地址
	LinkURL      string `json:"link_url,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysMenu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysmenu.FieldID, sysmenu.FieldPid, sysmenu.FieldMenuType, sysmenu.FieldWeigh, sysmenu.FieldIsHide, sysmenu.FieldIsLink, sysmenu.FieldModelID, sysmenu.FieldIsIframe, sysmenu.FieldIsCached, sysmenu.FieldIsAffix:
			values[i] = new(sql.NullInt64)
		case sysmenu.FieldRemark, sysmenu.FieldName, sysmenu.FieldTitle, sysmenu.FieldIcon, sysmenu.FieldCondition, sysmenu.FieldPath, sysmenu.FieldComponent, sysmenu.FieldModuleType, sysmenu.FieldRedirect, sysmenu.FieldLinkURL:
			values[i] = new(sql.NullString)
		case sysmenu.FieldCreatedAt, sysmenu.FieldUpdatedAt, sysmenu.FieldDeleteAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysMenu fields.
func (sm *SysMenu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysmenu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sm.ID = int64(value.Int64)
		case sysmenu.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sm.CreatedAt = value.Time
			}
		case sysmenu.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sm.UpdatedAt = value.Time
			}
		case sysmenu.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				sm.DeleteAt = value.Time
			}
		case sysmenu.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				sm.Remark = value.String
			}
		case sysmenu.FieldPid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				sm.Pid = value.Int64
			}
		case sysmenu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sm.Name = value.String
			}
		case sysmenu.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				sm.Title = value.String
			}
		case sysmenu.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				sm.Icon = value.String
			}
		case sysmenu.FieldCondition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field condition", values[i])
			} else if value.Valid {
				sm.Condition = value.String
			}
		case sysmenu.FieldMenuType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field menu_type", values[i])
			} else if value.Valid {
				sm.MenuType = int8(value.Int64)
			}
		case sysmenu.FieldWeigh:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weigh", values[i])
			} else if value.Valid {
				sm.Weigh = value.Int64
			}
		case sysmenu.FieldIsHide:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_hide", values[i])
			} else if value.Valid {
				sm.IsHide = int8(value.Int64)
			}
		case sysmenu.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				sm.Path = value.String
			}
		case sysmenu.FieldComponent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field component", values[i])
			} else if value.Valid {
				sm.Component = value.String
			}
		case sysmenu.FieldIsLink:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_link", values[i])
			} else if value.Valid {
				sm.IsLink = int8(value.Int64)
			}
		case sysmenu.FieldModuleType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field module_type", values[i])
			} else if value.Valid {
				sm.ModuleType = value.String
			}
		case sysmenu.FieldModelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field model_id", values[i])
			} else if value.Valid {
				sm.ModelID = value.Int64
			}
		case sysmenu.FieldIsIframe:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_iframe", values[i])
			} else if value.Valid {
				sm.IsIframe = int8(value.Int64)
			}
		case sysmenu.FieldIsCached:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_cached", values[i])
			} else if value.Valid {
				sm.IsCached = int8(value.Int64)
			}
		case sysmenu.FieldRedirect:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect", values[i])
			} else if value.Valid {
				sm.Redirect = value.String
			}
		case sysmenu.FieldIsAffix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_affix", values[i])
			} else if value.Valid {
				sm.IsAffix = int8(value.Int64)
			}
		case sysmenu.FieldLinkURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link_url", values[i])
			} else if value.Valid {
				sm.LinkURL = value.String
			}
		default:
			sm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SysMenu.
// This includes values selected through modifiers, order, etc.
func (sm *SysMenu) Value(name string) (ent.Value, error) {
	return sm.selectValues.Get(name)
}

// Update returns a builder for updating this SysMenu.
// Note that you need to call SysMenu.Unwrap() before calling this method if this SysMenu
// was returned from a transaction, and the transaction was committed or rolled back.
func (sm *SysMenu) Update() *SysMenuUpdateOne {
	return NewSysMenuClient(sm.config).UpdateOne(sm)
}

// Unwrap unwraps the SysMenu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sm *SysMenu) Unwrap() *SysMenu {
	_tx, ok := sm.config.driver.(*txDriver)
	if !ok {
		panic("codegen: SysMenu is not a transactional entity")
	}
	sm.config.driver = _tx.drv
	return sm
}

// String implements the fmt.Stringer.
func (sm *SysMenu) String() string {
	var builder strings.Builder
	builder.WriteString("SysMenu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(sm.DeleteAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(sm.Remark)
	builder.WriteString(", ")
	builder.WriteString("pid=")
	builder.WriteString(fmt.Sprintf("%v", sm.Pid))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sm.Name)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(sm.Title)
	builder.WriteString(", ")
	builder.WriteString("icon=")
	builder.WriteString(sm.Icon)
	builder.WriteString(", ")
	builder.WriteString("condition=")
	builder.WriteString(sm.Condition)
	builder.WriteString(", ")
	builder.WriteString("menu_type=")
	builder.WriteString(fmt.Sprintf("%v", sm.MenuType))
	builder.WriteString(", ")
	builder.WriteString("weigh=")
	builder.WriteString(fmt.Sprintf("%v", sm.Weigh))
	builder.WriteString(", ")
	builder.WriteString("is_hide=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsHide))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(sm.Path)
	builder.WriteString(", ")
	builder.WriteString("component=")
	builder.WriteString(sm.Component)
	builder.WriteString(", ")
	builder.WriteString("is_link=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsLink))
	builder.WriteString(", ")
	builder.WriteString("module_type=")
	builder.WriteString(sm.ModuleType)
	builder.WriteString(", ")
	builder.WriteString("model_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.ModelID))
	builder.WriteString(", ")
	builder.WriteString("is_iframe=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsIframe))
	builder.WriteString(", ")
	builder.WriteString("is_cached=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsCached))
	builder.WriteString(", ")
	builder.WriteString("redirect=")
	builder.WriteString(sm.Redirect)
	builder.WriteString(", ")
	builder.WriteString("is_affix=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsAffix))
	builder.WriteString(", ")
	builder.WriteString("link_url=")
	builder.WriteString(sm.LinkURL)
	builder.WriteByte(')')
	return builder.String()
}

// SysMenus is a parsable slice of SysMenu.
type SysMenus []*SysMenu
