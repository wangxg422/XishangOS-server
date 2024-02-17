package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SysCasbinRule holds the schema definition for the SysCasbinRule entity.
type SysCasbinRule struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysCasbinRule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_casbin_rule"},
		entsql.WithComments(true),
		schema.Comment("casbin权限表"),
	}
}

// Fields of the SysCasbinRule.
func (SysCasbinRule) Fields() []ent.Field {
	return []ent.Field{
		field.String("ptype").Optional(),
		field.String("v0").Optional(),
		field.String("v1").Optional(),
		field.String("v2").Optional(),
		field.String("v3").Optional(),
		field.String("v4").Optional(),
		field.String("v5").Optional(),
	}
}

// Edges of the SysCasbinRule.
func (SysCasbinRule) Edges() []ent.Edge {
	return nil
}
