package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/base/schema/mixin"
)

// SysRoleDept holds the schema definition for the SysRoleDept entity.
type SysRoleDept struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysRoleDept) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_role_dept"},
		entsql.WithComments(true),
		schema.Comment("系统角色部门关联表"),
	}
}

// Mixin 嵌入字段
func (SysRoleDept) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
	}
}

// Fields of the SysRoleDept.
func (SysRoleDept) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("role_id").Comment("角色id"),
		field.Int64("dept_id").Comment("部门id"),
	}
}

// Edges of the SysRoleDept.
func (SysRoleDept) Edges() []ent.Edge {
	return nil
}
