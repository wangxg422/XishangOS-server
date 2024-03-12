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
		field.Int64("pid").StructTag(`json:"pId,string"`).Optional().Comment("父菜单ID"),
		field.String("path").StructTag(`json:"path"`).Optional().Comment("路由地址"),
		field.String("name").StructTag(`json:"name"`).Optional().Comment("菜单名称"),
		field.String("component").StructTag(`json:"component"`).Optional().Comment("组件路径"),
		field.String("redirect").StructTag(`json:"redirect"`).Optional().Comment("路由重定向地址"),
		field.String("title").StructTag(`json:"title"`).Optional().Comment("菜单标题"),
		field.String("icon").StructTag(`json:"icon"`).Optional().Comment("菜单图标"),
		field.Int8("is_link").StructTag(`json:"isLink"`).Optional().Comment("是否是外部链接(1是,2否)"),
		field.Int8("is_hide").StructTag(`json:"isHide"`).Optional().Comment("显示状态"),
		field.Int8("is_affix").StructTag(`json:"isAffix"`).Optional().Comment("是否固定(1是,2否)"),
		field.Int8("is_keep_alive").StructTag(`json:"isKeepAlive"`).Optional().Comment("是否缓存(1是,2否)"),
		field.Int8("is_full").StructTag(`json:"isFull"`).Optional().Comment("是否是全屏页面(1是,2否)"),
		field.Int8("is_iframe").StructTag(`json:"isIframe"`).Optional().Comment("是否是内嵌iframe(1是,0否)"),
		field.Int8("type").StructTag(`json:"type"`).Optional().Comment("菜单类型(1目录,2菜单,3按钮)"),
		field.String("module_type").StructTag(`json:"moduleType"`).Optional().Comment("所属模块"),
		field.String("link_url").StructTag(`json:"linkUrl"`).Optional().Comment("链接地址"),
	}
}

// Edges of the SysMenu.
func (SysMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sysRoles", SysRole.Type).
			Ref("sysMenus"),
		// 添加一对多的（O2M ）自引用，即一个节点有且仅有一个父节点，有多个子节点
		edge.To("children", SysMenu.Type).
			From("parent").
			Unique().
			Field("pid"),
	}
}
