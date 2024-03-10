// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/apppackage"
)

// 应用安装包信息表
type AppPackage struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,string"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"UpdateAt"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt time.Time `json:"deleteAt"`
	// Status holds the value of the "status" field.
	Status int8 `json:"status,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int64 `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int64 `json:"updated_by,omitempty"`
	// DeleteBy holds the value of the "delete_by" field.
	DeleteBy int64 `json:"delete_by,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// DelFlag holds the value of the "del_flag" field.
	DelFlag int8 `json:"-"`
	// PkgName holds the value of the "pkg_name" field.
	PkgName string `json:"pkg_name,omitempty"`
	// PkgCode holds the value of the "pkg_code" field.
	PkgCode string `json:"pkg_code,omitempty"`
	// 安装包版本
	PkgVersion string `json:"pkg_version,omitempty"`
	// PkgType holds the value of the "pkg_type" field.
	PkgType int8 `json:"pkg_type,omitempty"`
	// PkgKind holds the value of the "pkg_kind" field.
	PkgKind int8 `json:"pkg_kind,omitempty"`
	// 上传应用的用户
	Uploader int64 `json:"uploader,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppPackageQuery when eager-loading is set.
	Edges        AppPackageEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AppPackageEdges holds the relations/edges for other nodes in the graph.
type AppPackageEdges struct {
	// AppInstance holds the value of the app_instance edge.
	AppInstance []*AppInstance `json:"app_instance,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AppInstanceOrErr returns the AppInstance value or an error if the edge
// was not loaded in eager-loading.
func (e AppPackageEdges) AppInstanceOrErr() ([]*AppInstance, error) {
	if e.loadedTypes[0] {
		return e.AppInstance, nil
	}
	return nil, &NotLoadedError{edge: "app_instance"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppPackage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apppackage.FieldID, apppackage.FieldStatus, apppackage.FieldCreatedBy, apppackage.FieldUpdatedBy, apppackage.FieldDeleteBy, apppackage.FieldDelFlag, apppackage.FieldPkgType, apppackage.FieldPkgKind, apppackage.FieldUploader:
			values[i] = new(sql.NullInt64)
		case apppackage.FieldRemark, apppackage.FieldPkgName, apppackage.FieldPkgCode, apppackage.FieldPkgVersion, apppackage.FieldDesc:
			values[i] = new(sql.NullString)
		case apppackage.FieldCreatedAt, apppackage.FieldUpdatedAt, apppackage.FieldDeleteAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppPackage fields.
func (ap *AppPackage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apppackage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ap.ID = int64(value.Int64)
		case apppackage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ap.CreatedAt = value.Time
			}
		case apppackage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ap.UpdatedAt = value.Time
			}
		case apppackage.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ap.DeleteAt = value.Time
			}
		case apppackage.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ap.Status = int8(value.Int64)
			}
		case apppackage.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ap.CreatedBy = value.Int64
			}
		case apppackage.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ap.UpdatedBy = value.Int64
			}
		case apppackage.FieldDeleteBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_by", values[i])
			} else if value.Valid {
				ap.DeleteBy = value.Int64
			}
		case apppackage.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				ap.Remark = value.String
			}
		case apppackage.FieldDelFlag:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field del_flag", values[i])
			} else if value.Valid {
				ap.DelFlag = int8(value.Int64)
			}
		case apppackage.FieldPkgName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pkg_name", values[i])
			} else if value.Valid {
				ap.PkgName = value.String
			}
		case apppackage.FieldPkgCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pkg_code", values[i])
			} else if value.Valid {
				ap.PkgCode = value.String
			}
		case apppackage.FieldPkgVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pkg_version", values[i])
			} else if value.Valid {
				ap.PkgVersion = value.String
			}
		case apppackage.FieldPkgType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pkg_type", values[i])
			} else if value.Valid {
				ap.PkgType = int8(value.Int64)
			}
		case apppackage.FieldPkgKind:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pkg_kind", values[i])
			} else if value.Valid {
				ap.PkgKind = int8(value.Int64)
			}
		case apppackage.FieldUploader:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uploader", values[i])
			} else if value.Valid {
				ap.Uploader = value.Int64
			}
		case apppackage.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				ap.Desc = value.String
			}
		default:
			ap.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AppPackage.
// This includes values selected through modifiers, order, etc.
func (ap *AppPackage) Value(name string) (ent.Value, error) {
	return ap.selectValues.Get(name)
}

// QueryAppInstance queries the "app_instance" edge of the AppPackage entity.
func (ap *AppPackage) QueryAppInstance() *AppInstanceQuery {
	return NewAppPackageClient(ap.config).QueryAppInstance(ap)
}

// Update returns a builder for updating this AppPackage.
// Note that you need to call AppPackage.Unwrap() before calling this method if this AppPackage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ap *AppPackage) Update() *AppPackageUpdateOne {
	return NewAppPackageClient(ap.config).UpdateOne(ap)
}

// Unwrap unwraps the AppPackage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ap *AppPackage) Unwrap() *AppPackage {
	_tx, ok := ap.config.driver.(*txDriver)
	if !ok {
		panic("codegen: AppPackage is not a transactional entity")
	}
	ap.config.driver = _tx.drv
	return ap
}

// String implements the fmt.Stringer.
func (ap *AppPackage) String() string {
	var builder strings.Builder
	builder.WriteString("AppPackage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ap.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ap.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ap.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(ap.DeleteAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ap.Status))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ap.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ap.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("delete_by=")
	builder.WriteString(fmt.Sprintf("%v", ap.DeleteBy))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(ap.Remark)
	builder.WriteString(", ")
	builder.WriteString("del_flag=")
	builder.WriteString(fmt.Sprintf("%v", ap.DelFlag))
	builder.WriteString(", ")
	builder.WriteString("pkg_name=")
	builder.WriteString(ap.PkgName)
	builder.WriteString(", ")
	builder.WriteString("pkg_code=")
	builder.WriteString(ap.PkgCode)
	builder.WriteString(", ")
	builder.WriteString("pkg_version=")
	builder.WriteString(ap.PkgVersion)
	builder.WriteString(", ")
	builder.WriteString("pkg_type=")
	builder.WriteString(fmt.Sprintf("%v", ap.PkgType))
	builder.WriteString(", ")
	builder.WriteString("pkg_kind=")
	builder.WriteString(fmt.Sprintf("%v", ap.PkgKind))
	builder.WriteString(", ")
	builder.WriteString("uploader=")
	builder.WriteString(fmt.Sprintf("%v", ap.Uploader))
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(ap.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// AppPackages is a parsable slice of AppPackage.
type AppPackages []*AppPackage
