package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	baseMixin "github.com/wangxg422/XishangOS-backend/app/module/base/schema/mixin"
)

// SysLoginLog holds the schema definition for the SysLoginLog entity.
type SysLoginLog struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysLoginLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_login_log"},
		entsql.WithComments(true),
		schema.Comment("系统用户登录日志表"),
	}
}

// Mixin 嵌入字段
func (SysLoginLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		baseMixin.IdMixin{},
		baseMixin.StatusMixin{},
	}
}

// Fields of the SysLoginLog.
func (SysLoginLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("login_name").Optional().Comment("登录账号"),
		field.String("ip_addr").Optional().Comment("登录IP地址"),
		field.String("login_location").Optional().Comment("登录地点"),
		field.String("browser").Optional().Comment("登录浏览器类型"),
		field.String("os").Optional().Comment("登录操作系统"),
		field.String("msg").Optional().Comment("提示消息"),
		field.Time("login_time").Optional().Comment("登录时间"),
		field.String("module").Optional().Comment("登录模块"),
	}
}

// Edges of the SysLoginLog.
func (SysLoginLog) Edges() []ent.Edge {
	return nil
}
