package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/base/schema/mixin"
)

// SysUserPost holds the schema definition for the SysUserPost entity.
type SysUserPost struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysUserPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_user_post"},
		entsql.WithComments(true),
		schema.Comment("系统用户岗位关联表"),
	}
}

// Mixin 嵌入字段
func (SysUserPost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
	}
}

// Fields of the SysUserPost.
func (SysUserPost) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Comment("用户id"),
		field.Int64("post_id").Comment("岗位id"),
	}
}

// Edges of the SysUserPost.
func (SysUserPost) Edges() []ent.Edge {
	return nil
}
