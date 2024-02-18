package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/base/schema/mixin"
)

// SysDictData holds the schema definition for the SysDictData entity.
type SysDictData struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysDictData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dict_data"},
		entsql.WithComments(true),
		schema.Comment("系统字典数据表"),
	}
}

// Mixin 嵌入字段
func (SysDictData) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
		mixin.ByMixin{},
		mixin.RemarkMixin{},
		mixin.StatusMixin{},
		mixin.SortMixin{},
	}
}

// Fields of the SysDictData.
func (SysDictData) Fields() []ent.Field {
	return []ent.Field{
		field.String("dict_label").Optional().Comment("字典标签"),
		field.String("dict_value").Optional().Comment("字典值"),
		field.String("dict_type").Optional().Comment("字典类型"),
		field.String("css_class").Optional().Comment("样式属性"),
		field.String("list_class").Optional().Comment("表格回显样式"),
		field.Int8("is_default").Optional().Comment("是否默认(1是0否)"),
	}
}

// Edges of the SysDictData.
func (SysDictData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", SysDictType.Type).
			Ref("sys_dict_data").
			Unique(),
	}
}