package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/base/schema/mixin"
)

// SysRole holds the schema definition for the SysRole entity.
type SysRole struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_role"},
		entsql.WithComments(true),
		schema.Comment("系统角色表"),
	}
}

// Mixin 嵌入字段
func (SysRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
		mixin.StatusMixin{},
		mixin.RemarkMixin{},
	}
}

// Fields of the SysRole.
func (SysRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("list_order").Optional().Comment("排序"),
		field.String("name").Optional().Comment("角色名称"),
		field.Int8("data_scope").Optional().Comment("数据权限范围(1全部数据权限 2自定数据权限 3本部门数据权限 4本部门及以下数据权限)"),
	}
}

// Edges of the SysRole.
func (SysRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("depts", SysDept.Type),
	}
}
