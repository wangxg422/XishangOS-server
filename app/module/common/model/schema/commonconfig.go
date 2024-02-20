package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/mixin"
	baseMixin "github.com/wangxg422/XishangOS-backend/app/module/base/schema/mixin"
)

// CommonConfig holds the schema definition for the CommonConfig entity.
type CommonConfig struct {
	ent.Schema
}

// Annotations 修改表名称
func (CommonConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "common_config"},
		entsql.WithComments(true),
		schema.Comment("系统配置表"),
	}
}

// Mixin 嵌入字段
func (CommonConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		baseMixin.IdMixin{},
		baseMixin.TimeMixin{},
		baseMixin.DeleteTimeMixin{},
		baseMixin.StatusMixin{},
		baseMixin.ByMixin{},
		baseMixin.RemarkMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the CommonConfig.
func (CommonConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("config_name").Optional().Comment("配置名称"),
		field.String("config_key").Optional().Comment("配置项"),
		field.String("config_value").Optional().Comment("配置值"),
		field.Int8("config_type").Optional().Comment("系统内置配置(0是1否)"),
	}
}

// Edges of the CommonConfig.
func (CommonConfig) Edges() []ent.Edge {
	return nil
}
