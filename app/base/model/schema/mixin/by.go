package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type ByMixin struct {
	mixin.Schema
}

func (ByMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_by").Optional(),
		field.Int64("updated_by").Optional(),
		field.Int64("delete_by").Optional(),
	}
}
