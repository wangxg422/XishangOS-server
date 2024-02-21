package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	baseMixin "github.com/wangxg422/XishangOS-backend/app/module/base/schema/mixin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/mixin"
)

// SysMenu holds the schema definition for the SysMenu entity.
type SysMenu struct {
	ent.Schema
}

// Annotations 修改表名称
func (SysMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_menu"},
		entsql.WithComments(true),
		schema.Comment("系统菜单权限表"),
	}
}

// Mixin 嵌入字段
func (SysMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		baseMixin.IdMixin{},
		baseMixin.TimeMixin{},
		baseMixin.DeleteTimeMixin{},
		baseMixin.ByMixin{},
		baseMixin.StatusMixin{},
		baseMixin.RemarkMixin{},
		baseMixin.SortMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the SysMenu.
func (SysMenu) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("pid").Comment("父菜单ID"),
		field.String("menu_name").Optional().Comment("菜单名称"),
		field.String("menu_title").Optional().Comment("菜单标题"),
		field.String("menu_icon").Optional().Comment("菜单图标"),
		field.String("condition").Optional().Comment("条件"),
		field.Int8("menu_type").Optional().Comment("菜单类型(0目录,1菜单,2按钮)"),
		field.String("path").Optional().Comment("路由地址"),
		field.String("component").Optional().Comment("组件路径"),
		field.String("module_type").Optional().Comment("所属模块"),
		field.Int64("model_id").Optional().Comment("模型id"),
		field.Int8("is_hide").Optional().Comment("显示状态"),
		field.Int8("is_iframe").Optional().Comment("是否是内嵌iframe(1是,0否)"),
		field.Int8("is_link").Optional().Comment("是否是外部链接(1是,0否)"),
		field.Int8("is_cached").Optional().Comment("是否缓存(1是,0否)"),
		field.Int8("is_affix").Optional().Comment("是否固定(1是,0否)"),
		field.String("redirect").Optional().Comment("路由重定向地址"),
		field.String("link_url").Optional().Comment("链接地址"),
	}
}

// Edges of the SysMenu.
func (SysMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sysRoles", SysRole.Type).
			Ref("sysMenus"),
	}
}
