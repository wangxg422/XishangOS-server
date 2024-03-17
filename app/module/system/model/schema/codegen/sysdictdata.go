// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdictdata"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdicttype"
)

// 系统字典数据表
type SysDictData struct {
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
	// 字典标签
	DictLabel string `json:"dictLabel"`
	// 字典值
	DictValue string `json:"dictValue"`
	// 字典类型ID
	DictTypeID int64 `json:"dictTypeId,string"`
	// 样式属性
	CSSClass string `json:"cssClass"`
	// 表格回显样式
	ListClass string `json:"listClass"`
	// 是否默认(1是0否)
	IsDefault int8 `json:"isDefault"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SysDictDataQuery when eager-loading is set.
	Edges        SysDictDataEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SysDictDataEdges holds the relations/edges for other nodes in the graph.
type SysDictDataEdges struct {
	// SysDictType holds the value of the sysDictType edge.
	SysDictType *SysDictType `json:"sysDictType,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SysDictTypeOrErr returns the SysDictType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SysDictDataEdges) SysDictTypeOrErr() (*SysDictType, error) {
	if e.loadedTypes[0] {
		if e.SysDictType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: sysdicttype.Label}
		}
		return e.SysDictType, nil
	}
	return nil, &NotLoadedError{edge: "sysDictType"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SysDictData) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sysdictdata.FieldID, sysdictdata.FieldCreatedBy, sysdictdata.FieldUpdatedBy, sysdictdata.FieldDeleteBy, sysdictdata.FieldStatus, sysdictdata.FieldSort, sysdictdata.FieldDelFlag, sysdictdata.FieldDictTypeID, sysdictdata.FieldIsDefault:
			values[i] = new(sql.NullInt64)
		case sysdictdata.FieldRemark, sysdictdata.FieldDictLabel, sysdictdata.FieldDictValue, sysdictdata.FieldCSSClass, sysdictdata.FieldListClass:
			values[i] = new(sql.NullString)
		case sysdictdata.FieldCreatedAt, sysdictdata.FieldUpdatedAt, sysdictdata.FieldDeleteAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SysDictData fields.
func (sdd *SysDictData) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sysdictdata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sdd.ID = int64(value.Int64)
		case sysdictdata.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sdd.CreatedAt = value.Time
			}
		case sysdictdata.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sdd.UpdatedAt = value.Time
			}
		case sysdictdata.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				sdd.DeleteAt = value.Time
			}
		case sysdictdata.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				sdd.CreatedBy = value.Int64
			}
		case sysdictdata.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				sdd.UpdatedBy = value.Int64
			}
		case sysdictdata.FieldDeleteBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_by", values[i])
			} else if value.Valid {
				sdd.DeleteBy = value.Int64
			}
		case sysdictdata.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				sdd.Remark = value.String
			}
		case sysdictdata.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sdd.Status = int8(value.Int64)
			}
		case sysdictdata.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				sdd.Sort = int(value.Int64)
			}
		case sysdictdata.FieldDelFlag:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field del_flag", values[i])
			} else if value.Valid {
				sdd.DelFlag = int8(value.Int64)
			}
		case sysdictdata.FieldDictLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dict_label", values[i])
			} else if value.Valid {
				sdd.DictLabel = value.String
			}
		case sysdictdata.FieldDictValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dict_value", values[i])
			} else if value.Valid {
				sdd.DictValue = value.String
			}
		case sysdictdata.FieldDictTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dict_type_id", values[i])
			} else if value.Valid {
				sdd.DictTypeID = value.Int64
			}
		case sysdictdata.FieldCSSClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field css_class", values[i])
			} else if value.Valid {
				sdd.CSSClass = value.String
			}
		case sysdictdata.FieldListClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field list_class", values[i])
			} else if value.Valid {
				sdd.ListClass = value.String
			}
		case sysdictdata.FieldIsDefault:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				sdd.IsDefault = int8(value.Int64)
			}
		default:
			sdd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SysDictData.
// This includes values selected through modifiers, order, etc.
func (sdd *SysDictData) Value(name string) (ent.Value, error) {
	return sdd.selectValues.Get(name)
}

// QuerySysDictType queries the "sysDictType" edge of the SysDictData entity.
func (sdd *SysDictData) QuerySysDictType() *SysDictTypeQuery {
	return NewSysDictDataClient(sdd.config).QuerySysDictType(sdd)
}

// Update returns a builder for updating this SysDictData.
// Note that you need to call SysDictData.Unwrap() before calling this method if this SysDictData
// was returned from a transaction, and the transaction was committed or rolled back.
func (sdd *SysDictData) Update() *SysDictDataUpdateOne {
	return NewSysDictDataClient(sdd.config).UpdateOne(sdd)
}

// Unwrap unwraps the SysDictData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sdd *SysDictData) Unwrap() *SysDictData {
	_tx, ok := sdd.config.driver.(*txDriver)
	if !ok {
		panic("codegen: SysDictData is not a transactional entity")
	}
	sdd.config.driver = _tx.drv
	return sdd
}

// String implements the fmt.Stringer.
func (sdd *SysDictData) String() string {
	var builder strings.Builder
	builder.WriteString("SysDictData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sdd.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sdd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sdd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(sdd.DeleteAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", sdd.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", sdd.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("delete_by=")
	builder.WriteString(fmt.Sprintf("%v", sdd.DeleteBy))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(sdd.Remark)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", sdd.Status))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", sdd.Sort))
	builder.WriteString(", ")
	builder.WriteString("del_flag=")
	builder.WriteString(fmt.Sprintf("%v", sdd.DelFlag))
	builder.WriteString(", ")
	builder.WriteString("dict_label=")
	builder.WriteString(sdd.DictLabel)
	builder.WriteString(", ")
	builder.WriteString("dict_value=")
	builder.WriteString(sdd.DictValue)
	builder.WriteString(", ")
	builder.WriteString("dict_type_id=")
	builder.WriteString(fmt.Sprintf("%v", sdd.DictTypeID))
	builder.WriteString(", ")
	builder.WriteString("css_class=")
	builder.WriteString(sdd.CSSClass)
	builder.WriteString(", ")
	builder.WriteString("list_class=")
	builder.WriteString(sdd.ListClass)
	builder.WriteString(", ")
	builder.WriteString("is_default=")
	builder.WriteString(fmt.Sprintf("%v", sdd.IsDefault))
	builder.WriteByte(')')
	return builder.String()
}

// SysDictDataSlice is a parsable slice of SysDictData.
type SysDictDataSlice []*SysDictData
