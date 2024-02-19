package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	mixin2 "github.com/wangxg422/XishangOS-backend/app/base/model/schema/mixin"
)

// SysUser holds the schema definition for the SysUser entity.
type SysUser struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_user"},
		entsql.WithComments(true),
		schema.Comment("系统用户表"),
	}
}

// Mixin 嵌入字段
func (SysUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin2.IdMixin{},
		mixin2.TimeMixin{},
		mixin2.RemarkMixin{},
	}
}

// Fields of the SysUser.
func (SysUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name"),
		field.String("user_nickname").Optional().Comment("用户昵称"),
		field.String("mobile").Optional(),
		field.String("birthday").Optional(),
		field.String("user_password").Optional(),
		field.String("user_salt").Optional(),
		field.String("user_email").Optional(),
		field.Int8("sex").Optional(),
		field.String("avatar").Optional().Comment("用户头像地址"),
		field.Int8("is_admin").Optional(),
		field.Int8("user_status").Optional().Comment("用户状态(0禁用,1正常,2未知)"),
		field.Int64("dept_id").Optional().Comment("用户所属部门"),
		field.String("address").Optional().Comment("用户联系地址"),
		field.String("describe").Optional(),
		field.String("last_login_ip").Optional(),
		field.String("last_login_time").Optional(),
	}
}

// Edges of the SysUser.
func (SysUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sysDept", SysDept.Type).
			Ref("sysUsers").
			Field("dept_id").
			Unique(),
		edge.To("sysPosts", SysPost.Type).
			StorageKey(edge.Table("sys_user_post"), edge.Columns("user_id", "post_id")),
		edge.To("sysRoles", SysRole.Type).
			StorageKey(edge.Table("sys_user_role"), edge.Columns("user_id", "role_id")),
	}
}
