package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/common/model/schema/mixin"
)

// SysDictType holds the schema definition for the SysDictType entity.
type SysDictType struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysDictType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dict_type"},
		entsql.WithComments(true),
		schema.Comment("系统字典类型表"),
	}
}

// Mixin 嵌入字段
func (SysDictType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
		mixin.ByMixin{},
		mixin.RemarkMixin{},
		mixin.StatusMixin{},
	}
}

// Fields of the SysDictType.
func (SysDictType) Fields() []ent.Field {
	return []ent.Field{
		field.String("dict_name").Optional().Comment("字典类型名称"),
		field.String("dict_type").Optional().Comment("字典类型").Unique(),
	}
}

// Edges of the SysDictType.
func (SysDictType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sysDictDatas", SysDictData.Type),
	}
}
