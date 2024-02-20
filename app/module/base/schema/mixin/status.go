package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/wangxg422/XishangOS-backend/common/enmu"
)

type StatusMixin struct {
	mixin.Schema
}

func (StatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("status").
			Default(enmu.StatusNormal.Value()),
	}
}
