package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/base/schema/mixin"
)

// SysPost holds the schema definition for the SysPost entity.
type SysPost struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_post"},
		entsql.WithComments(true),
		schema.Comment("系统岗位表"),
	}
}

// Mixin 嵌入字段
func (SysPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
		mixin.ByMixin{},
		mixin.RemarkMixin{},
		mixin.StatusMixin{},
		mixin.SortMixin{},
	}
}

// Fields of the SysPost.
func (SysPost) Fields() []ent.Field {
	return []ent.Field{
		field.String("post_code").Optional().Comment("岗位编码"),
		field.String("post_name").Optional().Comment("岗位名称"),
	}
}

// Edges of the SysPost.
func (SysPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", SysUser.Type).
			Ref("posts"),
	}
}
