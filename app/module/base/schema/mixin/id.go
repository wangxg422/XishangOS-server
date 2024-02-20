package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/wangxg422/XishangOS-backend/app/tools/sql"
)

type IdMixin struct {
	mixin.Schema
}

func (IdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			DefaultFunc(sql.IdFunc()),
	}
}
