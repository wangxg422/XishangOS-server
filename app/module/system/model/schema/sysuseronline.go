package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/base/schema/mixin"
)

// SysUserOnline holds the schema definition for the SysUserOnline entity.
type SysUserOnline struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysUserOnline) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_user_online"},
		entsql.WithComments(true),
		schema.Comment("在线用户表"),
	}
}

// Mixin 嵌入字段
func (SysUserOnline) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
	}
}

// Fields of the SysUserOnline.
func (SysUserOnline) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("uuid").Optional().Comment("用户标识"),
		field.Int64("token").Optional().Comment("用户token"),
		field.Int64("create_time").Optional().Comment("登录时间"),
		field.Int64("user_name").Optional().Comment("用户名"),
		field.Int64("ip_addr").Optional().Comment("登录ip"),
		field.Int64("browser").Optional().Comment("浏览器"),
		field.Int64("os").Optional().Comment("操作系统"),
	}
}

// Edges of the SysUserOnline.
func (SysUserOnline) Edges() []ent.Edge {
	return nil
}
