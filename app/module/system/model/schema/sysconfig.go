package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	mixin "github.com/wangxg422/XishangOS-backend/app/module/base/schema/mixin"
)

// SysConfig holds the schema definition for the SysConfig entity.
type SysConfig struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_config"},
		entsql.WithComments(true),
		schema.Comment("系统配置表"),
	}
}

// Mixin 嵌入字段
func (SysConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
		mixin.ByMixin{},
		mixin.RemarkMixin{},
	}
}

// Fields of the SysConfig.
func (SysConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("config_name").Optional().Comment("配置名称"),
		field.String("config_key").Optional().Comment("配置项"),
		field.String("config_value").Optional().Comment("配置值"),
		field.Int8("config_type").Optional().Comment("系统内置配置(0是1否)"),
	}
}

// Edges of the SysConfig.
func (SysConfig) Edges() []ent.Edge {
	return nil
}
