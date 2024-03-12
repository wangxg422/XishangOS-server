package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	baseMixin "github.com/wangxg422/XishangOS-backend/app/module/base/schema/mixin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/mixin"
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
		baseMixin.IdMixin{},
		baseMixin.TimeMixin{},
		baseMixin.DeleteTimeMixin{},
		baseMixin.ByMixin{},
		baseMixin.RemarkMixin{},
		baseMixin.StatusMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the SysDictType.
func (SysDictType) Fields() []ent.Field {
	return []ent.Field{
		field.String("dict_name").StructTag(`json:"dictName"`).Optional().Comment("字典类型名称"),
		field.String("dict_type").StructTag(`json:"dictType"`).Optional().Comment("字典类型").Unique(),
	}
}

// Edges of the SysDictType.
func (SysDictType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sysDictDatas", SysDictData.Type),
	}
}
