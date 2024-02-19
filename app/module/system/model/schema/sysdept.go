package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	mixin2 "github.com/wangxg422/XishangOS-backend/app/base/model/schema/mixin"
)

// SysDept holds the schema definition for the SysDept entity.
type SysDept struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysDept) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dept"},
		entsql.WithComments(true),
		schema.Comment("系统部门表"),
	}
}

// Mixin 嵌入字段
func (SysDept) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin2.IdMixin{},
		mixin2.TimeMixin{},
		mixin2.ByMixin{},
		mixin2.RemarkMixin{},
		mixin2.StatusMixin{},
		mixin2.SortMixin{},
	}
}

// Fields of the SysDept.
func (SysDept) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("parent_id").Comment("父级部门id"),
		field.String("ancestors").Optional().Comment("祖先部门列表"),
		field.String("dept_name").Optional().Comment("部门名称"),
		field.String("dept_code").Optional().Comment("部门编码"),
		field.String("leader").Optional().Comment("负责人"),
		field.String("phone").Optional().Comment("部门联系电话"),
		field.String("email").Optional().Comment("部门电子邮箱"),
	}
}

// Edges of the SysDept.
func (SysDept) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sysUsers", SysUser.Type),
		edge.From("sysRoles", SysRole.Type).
			Ref("sysDepts"),
	}
}
