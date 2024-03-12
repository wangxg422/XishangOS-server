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
		baseMixin.IdMixin{},
		baseMixin.TimeMixin{},
		baseMixin.DeleteTimeMixin{},
		baseMixin.ByMixin{},
		baseMixin.RemarkMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the SysUser.
func (SysUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").StructTag(`json:"userName"`),
		field.String("nickname").StructTag(`json:"nickName"`).Optional().Comment("用户昵称"),
		field.String("mobile").StructTag(`json:"mobile"`).Optional(),
		field.String("birthday").StructTag(`json:"birthday"`).Optional(),
		field.String("password").StructTag(`json:"password"`).Optional(),
		field.String("salt").StructTag(`json:"salt"`).Optional(),
		field.String("email").StructTag(`json:"email"`).Optional(),
		field.Int8("sex").StructTag(`json:"sex"`).Optional(),
		field.String("avatar").StructTag(`json:"avatar"`).Optional().Comment("用户头像地址"),
		field.Int8("user_status").StructTag(`json:"userStatus"`).Optional().Comment("用户状态(0禁用,1正常,2未知)"),
		field.Int64("dept_id").StructTag(`json:"deptId,string"`).Optional().Comment("用户所属部门"),
		field.String("address").StructTag(`json:"address"`).Optional().Comment("用户联系地址"),
		field.String("desc").StructTag(`json:"desc"`).Optional(),
		field.String("last_login_ip").StructTag(`json:"lastLoginIp"`).Optional(),
		field.String("last_login_time").StructTag(`json:"lastLoginTime"`).Optional(),
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
