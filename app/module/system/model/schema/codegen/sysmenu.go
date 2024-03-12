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
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// Sort holds the value of the "sort" field.
	Sort int `json:"sort,omitempty"`
	// DelFlag holds the value of the "del_flag" field.
	DelFlag int8 `json:"-"`
	// 父菜单ID
	Pid int64 `json:"pId,string"`
	// 路由地址
	Path string `json:"path"`
	// 菜单名称
	Name string `json:"name"`
	// 组件路径
	Component string `json:"component"`
	// 路由重定向地址
	Redirect string `json:"redirect"`
	// 菜单标题
	Title string `json:"title"`
	// 菜单图标
	Icon string `json:"icon"`
	// 是否是外部链接(1是,2否)
	IsLink int8 `json:"isLink"`
	// 显示状态
	IsHide int8 `json:"isHide"`
	// 是否固定(1是,2否)
	IsAffix int8 `json:"isAffix"`
	// 是否缓存(1是,2否)
	IsKeepAlive int8 `json:"isKeepAlive"`
	// 是否是内嵌iframe(1是,0否)
	IsIframe int8 `json:"isIframe"`
	// 菜单类型(1目录,2菜单,3按钮)
	Type int8 `json:"type"`
	// 所属模块
	ModuleType string `json:"moduleType"`
	// 链接地址
	LinkURL string `json:"linkUrl"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysMenuQuery when eager-loading is set.
	Edges        SysMenuEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SysMenuEdges holds the relations/edges for other nodes in the graph.
type SysMenuEdges struct {
	// SysRoles holds the value of the sysRoles edge.
	SysRoles []*SysRole `json:"sysRoles,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *SysMenu `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*SysMenu `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SysRolesOrErr returns the SysRoles value or an error if the edge
// was not loaded in eager-loading.
func (e SysMenuEdges) SysRolesOrErr() ([]*SysRole, error) {
	if e.loadedTypes[0] {
		return e.SysRoles, nil
	}
	return nil, &NotLoadedError{edge: "sysRoles"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysMenuEdges) ParentOrErr() (*SysMenu, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: sysmenu.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e SysMenuEdges) ChildrenOrErr() ([]*SysMenu, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysMenu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysmenu.FieldID, sysmenu.FieldCreatedBy, sysmenu.FieldUpdatedBy, sysmenu.FieldDeleteBy, sysmenu.FieldStatus, sysmenu.FieldSort, sysmenu.FieldDelFlag, sysmenu.FieldPid, sysmenu.FieldIsLink, sysmenu.FieldIsHide, sysmenu.FieldIsAffix, sysmenu.FieldIsKeepAlive, sysmenu.FieldIsIframe, sysmenu.FieldType:
			values[i] = new(sql.NullInt64)
		case sysmenu.FieldRemark, sysmenu.FieldPath, sysmenu.FieldName, sysmenu.FieldComponent, sysmenu.FieldRedirect, sysmenu.FieldTitle, sysmenu.FieldIcon, sysmenu.FieldModuleType, sysmenu.FieldLinkURL:
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
		case sysmenu.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				sm.CreatedBy = value.Int64
			}
		case sysmenu.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				sm.UpdatedBy = value.Int64
			}
		case sysmenu.FieldDeleteBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_by", values[i])
			} else if value.Valid {
				sm.DeleteBy = value.Int64
			}
		case sysmenu.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sm.Status = int8(value.Int64)
			}
		case sysmenu.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				sm.Remark = value.String
			}
		case sysmenu.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sm.Sort = int(value.Int64)
			}
		case sysmenu.FieldDelFlag:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field del_flag", values[i])
			} else if value.Valid {
				sm.DelFlag = int8(value.Int64)
			}
		case sysmenu.FieldPid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				sm.Pid = value.Int64
			}
		case sysmenu.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				sm.Path = value.String
			}
		case sysmenu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sm.Name = value.String
			}
		case sysmenu.FieldComponent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field component", values[i])
			} else if value.Valid {
				sm.Component = value.String
			}
		case sysmenu.FieldRedirect:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect", values[i])
			} else if value.Valid {
				sm.Redirect = value.String
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
		case sysmenu.FieldIsLink:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_link", values[i])
			} else if value.Valid {
				sm.IsLink = int8(value.Int64)
			}
		case sysmenu.FieldIsHide:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_hide", values[i])
			} else if value.Valid {
				sm.IsHide = int8(value.Int64)
			}
		case sysmenu.FieldIsAffix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_affix", values[i])
			} else if value.Valid {
				sm.IsAffix = int8(value.Int64)
			}
		case sysmenu.FieldIsKeepAlive:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_keep_alive", values[i])
			} else if value.Valid {
				sm.IsKeepAlive = int8(value.Int64)
			}
		case sysmenu.FieldIsIframe:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_iframe", values[i])
			} else if value.Valid {
				sm.IsIframe = int8(value.Int64)
			}
		case sysmenu.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				sm.Type = int8(value.Int64)
			}
		case sysmenu.FieldModuleType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field module_type", values[i])
			} else if value.Valid {
				sm.ModuleType = value.String
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

// QuerySysRoles queries the "sysRoles" edge of the SysMenu entity.
func (sm *SysMenu) QuerySysRoles() *SysRoleQuery {
	return NewSysMenuClient(sm.config).QuerySysRoles(sm)
}

// QueryParent queries the "parent" edge of the SysMenu entity.
func (sm *SysMenu) QueryParent() *SysMenuQuery {
	return NewSysMenuClient(sm.config).QueryParent(sm)
}

// QueryChildren queries the "children" edge of the SysMenu entity.
func (sm *SysMenu) QueryChildren() *SysMenuQuery {
	return NewSysMenuClient(sm.config).QueryChildren(sm)
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
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", sm.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", sm.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("delete_by=")
	builder.WriteString(fmt.Sprintf("%v", sm.DeleteBy))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", sm.Status))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(sm.Remark)
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", sm.Sort))
	builder.WriteString(", ")
	builder.WriteString("del_flag=")
	builder.WriteString(fmt.Sprintf("%v", sm.DelFlag))
	builder.WriteString(", ")
	builder.WriteString("pid=")
	builder.WriteString(fmt.Sprintf("%v", sm.Pid))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(sm.Path)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sm.Name)
	builder.WriteString(", ")
	builder.WriteString("component=")
	builder.WriteString(sm.Component)
	builder.WriteString(", ")
	builder.WriteString("redirect=")
	builder.WriteString(sm.Redirect)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(sm.Title)
	builder.WriteString(", ")
	builder.WriteString("icon=")
	builder.WriteString(sm.Icon)
	builder.WriteString(", ")
	builder.WriteString("is_link=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsLink))
	builder.WriteString(", ")
	builder.WriteString("is_hide=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsHide))
	builder.WriteString(", ")
	builder.WriteString("is_affix=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsAffix))
	builder.WriteString(", ")
	builder.WriteString("is_keep_alive=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsKeepAlive))
	builder.WriteString(", ")
	builder.WriteString("is_iframe=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsIframe))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", sm.Type))
	builder.WriteString(", ")
	builder.WriteString("module_type=")
	builder.WriteString(sm.ModuleType)
	builder.WriteString(", ")
	builder.WriteString("link_url=")
	builder.WriteString(sm.LinkURL)
	builder.WriteByte(')')
	return builder.String()
}

// SysMenus is a parsable slice of SysMenu.
type SysMenus []*SysMenu
