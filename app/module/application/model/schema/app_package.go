package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/common/model/schema/mixin"
)

// AppPackage holds the schema definition for the AppPackage entity.
type AppPackage struct {
	ent.Schema
}

// Annotations 修改表名称
func (AppPackage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_package"},
		entsql.WithComments(true),
		schema.Comment("应用安装包信息表"),
	}
}

// Mixin 嵌入字段
func (AppPackage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
		mixin.StatusMixin{},
	}
}

// Fields of the AppPackage.
func (AppPackage) Fields() []ent.Field {
	return []ent.Field{
		field.String("pkg_name"),
		field.String("pkg_code"),
		field.String("pkg_version").Optional().Comment("安装包版本"),
		field.Int8("pkg_type").Optional(),
		field.Int8("pkg_kind").Optional(),
		field.Int64("uploader").Optional().Comment("上传应用的用户"),
		field.String("desc").Optional(),
		field.String("remark").Optional(),
	}
}

// Edges of the AppPackage.
func (AppPackage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app_instance", AppInstance.Type),
	}
}
