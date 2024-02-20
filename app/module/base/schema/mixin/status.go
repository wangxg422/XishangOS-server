package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type StatusMixin struct {
	mixin.Schema
}

func (StatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("status").
			Immutable().
			Default(1).
			Optional(),
	}
}
