package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	mixin "github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/mixin"
	baseMixin "github.com/wangxg422/XishangOS-backend/app/module/base/schema/mixin"
)

// AppInstance holds the schema definition for the AppInstance entity.
type AppInstance struct {
	ent.Schema
}

// Annotations 修改表名称
func (AppInstance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_instance"},
		entsql.WithComments(true),
		schema.Comment("应用实例表"),
	}
}

// Mixin 嵌入字段
func (AppInstance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		baseMixin.IdMixin{},
		baseMixin.TimeMixin{},
		baseMixin.DeleteTimeMixin{},
		baseMixin.StatusMixin{},
		baseMixin.RemarkMixin{},
		baseMixin.ByMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the AppInstance.
func (AppInstance) Fields() []ent.Field {
	return []ent.Field{
		field.String("instance_name").StructTag(`json:"instanceName"`),
		field.String("instance_code").StructTag(`json:"instanceCode"`),
		field.Int64("instance_package").StructTag(`json:"instancePackage"`),
		field.String("instance_icon").StructTag(`json:"instanceIcon"`).Optional().Comment("应用图标"),
		field.String("instance_address").StructTag(`json:"instanceAddress"`).Optional().Comment("应用访问地址"),
		field.Int8("instance_type").StructTag(`json:"instanceType"`).Optional(),
		field.Int64("installer").Optional().Comment("安装应用的用户"),
		field.String("desc").Optional(),
	}
}

// Edges of the AppInstance.
func (AppInstance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("installFrom", AppPackage.Type).
			Ref("app_instance").
			// 设置边为唯一确保每个应用仅由一个安装包安装
			Unique(),
	}
}
