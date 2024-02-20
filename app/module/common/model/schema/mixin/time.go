package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Optional(),
		field.Time("delete_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Optional(),
	}
}
