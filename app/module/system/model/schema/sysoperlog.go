package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/base/schema/mixin"
)

// SysOperLog holds the schema definition for the SysOperLog entity.
type SysOperLog struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysOperLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_oper_log"},
		entsql.WithComments(true),
		schema.Comment("系统操作日志表"),
	}
}

// Mixin 嵌入字段
func (SysOperLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
	}
}

// Fields of the SysOperLog.
func (SysOperLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").StructTag(`json:"title"`).Optional().Comment("模块标题"),
		field.Int("business_type").StructTag(`json:"businessType"`).Optional().Comment("业务类型(0其他,1新增,2修改,3删除)"),
		field.String("method").StructTag(`json:"method"`).Optional().Comment("方法名称"),
		field.String("request_method").StructTag(`json:"requestMethod"`).Optional().Comment("请求方式"),
		field.Int8("oper_type").StructTag(`json:"operType"`).Optional().Comment("操作分类(0其他,1后台用户,2手机端用户)"),
		field.String("oper_name").StructTag(`json:"operName"`).Optional().Comment("操作用户"),
		field.String("dept_name").StructTag(`json:"deptName"`).Optional().Comment("部门名称"),
		field.String("oper_url").StructTag(`json:"operUrl"`).Optional().Comment("请求url"),
		field.String("oper_ip").StructTag(`json:"ip"`).Optional().Comment("主机地址"),
		field.String("oper_location").StructTag(`json:"location"`).Optional().Comment("操作地点"),
		field.String("oper_param").StructTag(`json:"operParam"`).Optional().Comment("请求参数"),
		field.String("error_msg").StructTag(`json:"errorMsg"`).Optional().Comment("错误消息"),
		field.Time("oper_time").StructTag(`json:"operTime"`).Optional().Comment("操作时间"),
	}
}

// Edges of the SysOperLog.
func (SysOperLog) Edges() []ent.Edge {
	return nil
}
