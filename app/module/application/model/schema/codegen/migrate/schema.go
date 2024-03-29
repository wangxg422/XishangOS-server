// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppInstanceColumns holds the columns for the "app_instance" table.
	AppInstanceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "created_by", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_by", Type: field.TypeInt64, Nullable: true},
		{Name: "del_flag", Type: field.TypeInt8, Default: 1},
		{Name: "instance_name", Type: field.TypeString},
		{Name: "instance_code", Type: field.TypeString},
		{Name: "instance_package", Type: field.TypeInt64},
		{Name: "instance_icon", Type: field.TypeString, Nullable: true, Comment: "应用图标"},
		{Name: "instance_address", Type: field.TypeString, Nullable: true, Comment: "应用访问地址"},
		{Name: "instance_type", Type: field.TypeInt8, Nullable: true},
		{Name: "installer", Type: field.TypeInt64, Nullable: true, Comment: "安装应用的用户"},
		{Name: "desc", Type: field.TypeString, Nullable: true},
		{Name: "app_package_app_instance", Type: field.TypeInt64, Nullable: true},
	}
	// AppInstanceTable holds the schema information for the "app_instance" table.
	AppInstanceTable = &schema.Table{
		Name:       "app_instance",
		Comment:    "应用实例表",
		Columns:    AppInstanceColumns,
		PrimaryKey: []*schema.Column{AppInstanceColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_instance_app_package_app_instance",
				Columns:    []*schema.Column{AppInstanceColumns[18]},
				RefColumns: []*schema.Column{AppPackageColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AppPackageColumns holds the columns for the "app_package" table.
	AppPackageColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "created_by", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_by", Type: field.TypeInt64, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "del_flag", Type: field.TypeInt8, Default: 1},
		{Name: "pkg_name", Type: field.TypeString},
		{Name: "pkg_code", Type: field.TypeString},
		{Name: "pkg_version", Type: field.TypeString, Nullable: true, Comment: "安装包版本"},
		{Name: "pkg_type", Type: field.TypeInt8, Nullable: true},
		{Name: "pkg_kind", Type: field.TypeInt8, Nullable: true},
		{Name: "uploader", Type: field.TypeInt64, Nullable: true, Comment: "上传应用的用户"},
		{Name: "desc", Type: field.TypeString, Nullable: true},
	}
	// AppPackageTable holds the schema information for the "app_package" table.
	AppPackageTable = &schema.Table{
		Name:       "app_package",
		Comment:    "应用安装包信息表",
		Columns:    AppPackageColumns,
		PrimaryKey: []*schema.Column{AppPackageColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppInstanceTable,
		AppPackageTable,
	}
)

func init() {
	AppInstanceTable.ForeignKeys[0].RefTable = AppPackageTable
	AppInstanceTable.Annotation = &entsql.Annotation{
		Table: "app_instance",
	}
	AppPackageTable.Annotation = &entsql.Annotation{
		Table: "app_package",
	}
}
