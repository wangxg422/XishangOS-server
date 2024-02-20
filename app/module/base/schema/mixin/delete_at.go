package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type DeleteTimeMixin struct {
	mixin.Schema
}

func (DeleteTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_at").Optional(),
	}
}
