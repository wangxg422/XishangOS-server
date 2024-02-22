// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysConfigColumns holds the columns for the "sys_config" table.
	SysConfigColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_by", Type: field.TypeInt64, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "del_flag", Type: field.TypeInt8, Default: 1},
		{Name: "config_name", Type: field.TypeString, Nullable: true, Comment: "配置名称"},
		{Name: "config_key", Type: field.TypeString, Nullable: true, Comment: "配置项"},
		{Name: "config_value", Type: field.TypeString, Nullable: true, Comment: "配置值"},
		{Name: "config_type", Type: field.TypeInt8, Nullable: true, Comment: "系统内置配置(0是1否)"},
	}
	// SysConfigTable holds the schema information for the "sys_config" table.
	SysConfigTable = &schema.Table{
		Name:       "sys_config",
		Comment:    "系统配置表",
		Columns:    SysConfigColumns,
		PrimaryKey: []*schema.Column{SysConfigColumns[0]},
	}
	// SysDeptColumns holds the columns for the "sys_dept" table.
	SysDeptColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_by", Type: field.TypeInt64, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "sort", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "del_flag", Type: field.TypeInt8, Default: 1},
		{Name: "parent_id", Type: field.TypeInt64, Comment: "父级部门id"},
		{Name: "ancestors", Type: field.TypeString, Nullable: true, Comment: "祖先部门列表"},
		{Name: "dept_name", Type: field.TypeString, Nullable: true, Comment: "部门名称"},
		{Name: "dept_code", Type: field.TypeString, Nullable: true, Comment: "部门编码"},
		{Name: "leader", Type: field.TypeString, Nullable: true, Comment: "负责人"},
		{Name: "phone", Type: field.TypeString, Nullable: true, Comment: "部门联系电话"},
		{Name: "email", Type: field.TypeString, Nullable: true, Comment: "部门电子邮箱"},
		{Name: "address", Type: field.TypeString, Nullable: true, Comment: "部门地址"},
	}
	// SysDeptTable holds the schema information for the "sys_dept" table.
	SysDeptTable = &schema.Table{
		Name:       "sys_dept",
		Comment:    "系统部门表",
		Columns:    SysDeptColumns,
		PrimaryKey: []*schema.Column{SysDeptColumns[0]},
	}
	// SysDictDataColumns holds the columns for the "sys_dict_data" table.
	SysDictDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_by", Type: field.TypeInt64, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "sort", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "del_flag", Type: field.TypeInt8, Default: 1},
		{Name: "dict_label", Type: field.TypeString, Nullable: true, Comment: "字典标签"},
		{Name: "dict_value", Type: field.TypeString, Nullable: true, Comment: "字典值"},
		{Name: "css_class", Type: field.TypeString, Nullable: true, Comment: "样式属性"},
		{Name: "list_class", Type: field.TypeString, Nullable: true, Comment: "表格回显样式"},
		{Name: "is_default", Type: field.TypeInt8, Nullable: true, Comment: "是否默认(1是0否)"},
		{Name: "dict_type_id", Type: field.TypeInt64, Nullable: true, Comment: "字典类型ID"},
	}
	// SysDictDataTable holds the schema information for the "sys_dict_data" table.
	SysDictDataTable = &schema.Table{
		Name:       "sys_dict_data",
		Comment:    "系统字典数据表",
		Columns:    SysDictDataColumns,
		PrimaryKey: []*schema.Column{SysDictDataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_dict_data_sys_dict_type_sysDictDatas",
				Columns:    []*schema.Column{SysDictDataColumns[16]},
				RefColumns: []*schema.Column{SysDictTypeColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysDictTypeColumns holds the columns for the "sys_dict_type" table.
	SysDictTypeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_by", Type: field.TypeInt64, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "del_flag", Type: field.TypeInt8, Default: 1},
		{Name: "dict_name", Type: field.TypeString, Nullable: true, Comment: "字典类型名称"},
		{Name: "dict_type", Type: field.TypeString, Unique: true, Nullable: true, Comment: "字典类型"},
	}
	// SysDictTypeTable holds the schema information for the "sys_dict_type" table.
	SysDictTypeTable = &schema.Table{
		Name:       "sys_dict_type",
		Comment:    "系统字典类型表",
		Columns:    SysDictTypeColumns,
		PrimaryKey: []*schema.Column{SysDictTypeColumns[0]},
	}
	// SysLoginLogColumns holds the columns for the "sys_login_log" table.
	SysLoginLogColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "login_name", Type: field.TypeString, Nullable: true, Comment: "登录账号"},
		{Name: "ip_addr", Type: field.TypeString, Nullable: true, Comment: "登录IP地址"},
		{Name: "login_location", Type: field.TypeString, Nullable: true, Comment: "登录地点"},
		{Name: "browser", Type: field.TypeString, Nullable: true, Comment: "登录浏览器类型"},
		{Name: "os", Type: field.TypeString, Nullable: true, Comment: "登录操作系统"},
		{Name: "msg", Type: field.TypeString, Nullable: true, Comment: "提示消息"},
		{Name: "login_time", Type: field.TypeTime, Nullable: true, Comment: "登录时间"},
		{Name: "login_success", Type: field.TypeInt8, Nullable: true, Comment: "登录是否成功"},
		{Name: "module", Type: field.TypeString, Nullable: true, Comment: "登录模块"},
	}
	// SysLoginLogTable holds the schema information for the "sys_login_log" table.
	SysLoginLogTable = &schema.Table{
		Name:       "sys_login_log",
		Comment:    "系统用户登录日志表",
		Columns:    SysLoginLogColumns,
		PrimaryKey: []*schema.Column{SysLoginLogColumns[0]},
	}
	// SysMenuColumns holds the columns for the "sys_menu" table.
	SysMenuColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_by", Type: field.TypeInt64, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "sort", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "del_flag", Type: field.TypeInt8, Default: 1},
		{Name: "menu_name", Type: field.TypeString, Nullable: true, Comment: "菜单名称"},
		{Name: "menu_title", Type: field.TypeString, Nullable: true, Comment: "菜单标题"},
		{Name: "menu_icon", Type: field.TypeString, Nullable: true, Comment: "菜单图标"},
		{Name: "condition", Type: field.TypeString, Nullable: true, Comment: "条件"},
		{Name: "menu_type", Type: field.TypeInt8, Nullable: true, Comment: "菜单类型(0目录,1菜单,2按钮)"},
		{Name: "path", Type: field.TypeString, Nullable: true, Comment: "路由地址"},
		{Name: "component", Type: field.TypeString, Nullable: true, Comment: "组件路径"},
		{Name: "module_type", Type: field.TypeString, Nullable: true, Comment: "所属模块"},
		{Name: "model_id", Type: field.TypeInt64, Nullable: true, Comment: "模型id"},
		{Name: "is_hide", Type: field.TypeInt8, Nullable: true, Comment: "显示状态"},
		{Name: "is_iframe", Type: field.TypeInt8, Nullable: true, Comment: "是否是内嵌iframe(1是,0否)"},
		{Name: "is_link", Type: field.TypeInt8, Nullable: true, Comment: "是否是外部链接(1是,0否)"},
		{Name: "is_cached", Type: field.TypeInt8, Nullable: true, Comment: "是否缓存(1是,0否)"},
		{Name: "is_affix", Type: field.TypeInt8, Nullable: true, Comment: "是否固定(1是,0否)"},
		{Name: "redirect", Type: field.TypeString, Nullable: true, Comment: "路由重定向地址"},
		{Name: "link_url", Type: field.TypeString, Nullable: true, Comment: "链接地址"},
		{Name: "pid", Type: field.TypeInt64, Nullable: true, Comment: "父菜单ID"},
	}
	// SysMenuTable holds the schema information for the "sys_menu" table.
	SysMenuTable = &schema.Table{
		Name:       "sys_menu",
		Comment:    "系统菜单权限表",
		Columns:    SysMenuColumns,
		PrimaryKey: []*schema.Column{SysMenuColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menu_sys_menu_children",
				Columns:    []*schema.Column{SysMenuColumns[27]},
				RefColumns: []*schema.Column{SysMenuColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysOperLogColumns holds the columns for the "sys_oper_log" table.
	SysOperLogColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "title", Type: field.TypeString, Nullable: true, Comment: "模块标题"},
		{Name: "business_type", Type: field.TypeInt, Nullable: true, Comment: "业务类型(0其他,1新增,2修改,3删除)"},
		{Name: "method", Type: field.TypeString, Nullable: true, Comment: "方法名称"},
		{Name: "request_method", Type: field.TypeString, Nullable: true, Comment: "请求方式"},
		{Name: "oper_type", Type: field.TypeInt8, Nullable: true, Comment: "操作分类(0其他,1后台用户,2手机端用户)"},
		{Name: "oper_name", Type: field.TypeString, Nullable: true, Comment: "操作用户"},
		{Name: "dept_name", Type: field.TypeString, Nullable: true, Comment: "部门名称"},
		{Name: "oper_url", Type: field.TypeString, Nullable: true, Comment: "请求url"},
		{Name: "oper_ip", Type: field.TypeString, Nullable: true, Comment: "主机地址"},
		{Name: "oper_location", Type: field.TypeString, Nullable: true, Comment: "操作地点"},
		{Name: "oper_param", Type: field.TypeString, Nullable: true, Comment: "请求参数"},
		{Name: "error_msg", Type: field.TypeString, Nullable: true, Comment: "错误消息"},
		{Name: "oper_time", Type: field.TypeTime, Nullable: true, Comment: "操作时间"},
	}
	// SysOperLogTable holds the schema information for the "sys_oper_log" table.
	SysOperLogTable = &schema.Table{
		Name:       "sys_oper_log",
		Comment:    "系统操作日志表",
		Columns:    SysOperLogColumns,
		PrimaryKey: []*schema.Column{SysOperLogColumns[0]},
	}
	// SysPostColumns holds the columns for the "sys_post" table.
	SysPostColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_by", Type: field.TypeInt64, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "sort", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "del_flag", Type: field.TypeInt8, Default: 1},
		{Name: "post_code", Type: field.TypeString, Nullable: true, Comment: "岗位编码"},
		{Name: "post_name", Type: field.TypeString, Nullable: true, Comment: "岗位名称"},
	}
	// SysPostTable holds the schema information for the "sys_post" table.
	SysPostTable = &schema.Table{
		Name:       "sys_post",
		Comment:    "系统岗位表",
		Columns:    SysPostColumns,
		PrimaryKey: []*schema.Column{SysPostColumns[0]},
	}
	// SysRoleColumns holds the columns for the "sys_role" table.
	SysRoleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_by", Type: field.TypeInt64, Nullable: true},
		{Name: "status", Type: field.TypeInt8, Default: 1},
		{Name: "sort", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "del_flag", Type: field.TypeInt8, Default: 1},
		{Name: "role_name", Type: field.TypeString, Nullable: true, Comment: "角色名称"},
		{Name: "data_scope", Type: field.TypeInt8, Nullable: true, Comment: "数据权限范围(1全部数据权限 2自定数据权限 3本部门数据权限 4本部门及以下数据权限)"},
	}
	// SysRoleTable holds the schema information for the "sys_role" table.
	SysRoleTable = &schema.Table{
		Name:       "sys_role",
		Comment:    "系统角色表",
		Columns:    SysRoleColumns,
		PrimaryKey: []*schema.Column{SysRoleColumns[0]},
	}
	// SysUserColumns holds the columns for the "sys_user" table.
	SysUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "delete_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_by", Type: field.TypeInt64, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_by", Type: field.TypeInt64, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true},
		{Name: "del_flag", Type: field.TypeInt8, Default: 1},
		{Name: "user_name", Type: field.TypeString},
		{Name: "user_nickname", Type: field.TypeString, Nullable: true, Comment: "用户昵称"},
		{Name: "mobile", Type: field.TypeString, Nullable: true},
		{Name: "birthday", Type: field.TypeString, Nullable: true},
		{Name: "user_password", Type: field.TypeString, Nullable: true},
		{Name: "user_salt", Type: field.TypeString, Nullable: true},
		{Name: "user_email", Type: field.TypeString, Nullable: true},
		{Name: "sex", Type: field.TypeInt8, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Comment: "用户头像地址"},
		{Name: "user_status", Type: field.TypeInt8, Nullable: true, Comment: "用户状态(0禁用,1正常,2未知)"},
		{Name: "address", Type: field.TypeString, Nullable: true, Comment: "用户联系地址"},
		{Name: "describe", Type: field.TypeString, Nullable: true},
		{Name: "last_login_ip", Type: field.TypeString, Nullable: true},
		{Name: "last_login_time", Type: field.TypeString, Nullable: true},
		{Name: "dept_id", Type: field.TypeInt64, Nullable: true, Comment: "用户所属部门"},
	}
	// SysUserTable holds the schema information for the "sys_user" table.
	SysUserTable = &schema.Table{
		Name:       "sys_user",
		Comment:    "系统用户表",
		Columns:    SysUserColumns,
		PrimaryKey: []*schema.Column{SysUserColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_user_sys_dept_sysUsers",
				Columns:    []*schema.Column{SysUserColumns[23]},
				RefColumns: []*schema.Column{SysDeptColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysUserOnlineColumns holds the columns for the "sys_user_online" table.
	SysUserOnlineColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "uuid", Type: field.TypeInt64, Nullable: true, Comment: "用户标识"},
		{Name: "token", Type: field.TypeInt64, Nullable: true, Comment: "用户token"},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true, Comment: "登录时间"},
		{Name: "user_name", Type: field.TypeInt64, Nullable: true, Comment: "用户名"},
		{Name: "ip_addr", Type: field.TypeInt64, Nullable: true, Comment: "登录ip"},
		{Name: "browser", Type: field.TypeInt64, Nullable: true, Comment: "浏览器"},
		{Name: "os", Type: field.TypeInt64, Nullable: true, Comment: "操作系统"},
	}
	// SysUserOnlineTable holds the schema information for the "sys_user_online" table.
	SysUserOnlineTable = &schema.Table{
		Name:       "sys_user_online",
		Comment:    "在线用户表",
		Columns:    SysUserOnlineColumns,
		PrimaryKey: []*schema.Column{SysUserOnlineColumns[0]},
	}
	// SysRoleDeptColumns holds the columns for the "sys_role_dept" table.
	SysRoleDeptColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt64},
		{Name: "dept_id", Type: field.TypeInt64},
	}
	// SysRoleDeptTable holds the schema information for the "sys_role_dept" table.
	SysRoleDeptTable = &schema.Table{
		Name:       "sys_role_dept",
		Columns:    SysRoleDeptColumns,
		PrimaryKey: []*schema.Column{SysRoleDeptColumns[0], SysRoleDeptColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_role_dept_role_id",
				Columns:    []*schema.Column{SysRoleDeptColumns[0]},
				RefColumns: []*schema.Column{SysRoleColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "sys_role_dept_dept_id",
				Columns:    []*schema.Column{SysRoleDeptColumns[1]},
				RefColumns: []*schema.Column{SysDeptColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SysRoleMenuColumns holds the columns for the "sys_role_menu" table.
	SysRoleMenuColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt64},
		{Name: "menu_id", Type: field.TypeInt64},
	}
	// SysRoleMenuTable holds the schema information for the "sys_role_menu" table.
	SysRoleMenuTable = &schema.Table{
		Name:       "sys_role_menu",
		Columns:    SysRoleMenuColumns,
		PrimaryKey: []*schema.Column{SysRoleMenuColumns[0], SysRoleMenuColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_role_menu_role_id",
				Columns:    []*schema.Column{SysRoleMenuColumns[0]},
				RefColumns: []*schema.Column{SysRoleColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "sys_role_menu_menu_id",
				Columns:    []*schema.Column{SysRoleMenuColumns[1]},
				RefColumns: []*schema.Column{SysMenuColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SysUserPostColumns holds the columns for the "sys_user_post" table.
	SysUserPostColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "post_id", Type: field.TypeInt64},
	}
	// SysUserPostTable holds the schema information for the "sys_user_post" table.
	SysUserPostTable = &schema.Table{
		Name:       "sys_user_post",
		Columns:    SysUserPostColumns,
		PrimaryKey: []*schema.Column{SysUserPostColumns[0], SysUserPostColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_user_post_user_id",
				Columns:    []*schema.Column{SysUserPostColumns[0]},
				RefColumns: []*schema.Column{SysUserColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "sys_user_post_post_id",
				Columns:    []*schema.Column{SysUserPostColumns[1]},
				RefColumns: []*schema.Column{SysPostColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SysUserRoleColumns holds the columns for the "sys_user_role" table.
	SysUserRoleColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "role_id", Type: field.TypeInt64},
	}
	// SysUserRoleTable holds the schema information for the "sys_user_role" table.
	SysUserRoleTable = &schema.Table{
		Name:       "sys_user_role",
		Columns:    SysUserRoleColumns,
		PrimaryKey: []*schema.Column{SysUserRoleColumns[0], SysUserRoleColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_user_role_user_id",
				Columns:    []*schema.Column{SysUserRoleColumns[0]},
				RefColumns: []*schema.Column{SysUserColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "sys_user_role_role_id",
				Columns:    []*schema.Column{SysUserRoleColumns[1]},
				RefColumns: []*schema.Column{SysRoleColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysConfigTable,
		SysDeptTable,
		SysDictDataTable,
		SysDictTypeTable,
		SysLoginLogTable,
		SysMenuTable,
		SysOperLogTable,
		SysPostTable,
		SysRoleTable,
		SysUserTable,
		SysUserOnlineTable,
		SysRoleDeptTable,
		SysRoleMenuTable,
		SysUserPostTable,
		SysUserRoleTable,
	}
)

func init() {
	SysConfigTable.Annotation = &entsql.Annotation{
		Table: "sys_config",
	}
	SysDeptTable.Annotation = &entsql.Annotation{
		Table: "sys_dept",
	}
	SysDictDataTable.ForeignKeys[0].RefTable = SysDictTypeTable
	SysDictDataTable.Annotation = &entsql.Annotation{
		Table: "sys_dict_data",
	}
	SysDictTypeTable.Annotation = &entsql.Annotation{
		Table: "sys_dict_type",
	}
	SysLoginLogTable.Annotation = &entsql.Annotation{
		Table: "sys_login_log",
	}
	SysMenuTable.ForeignKeys[0].RefTable = SysMenuTable
	SysMenuTable.Annotation = &entsql.Annotation{
		Table: "sys_menu",
	}
	SysOperLogTable.Annotation = &entsql.Annotation{
		Table: "sys_oper_log",
	}
	SysPostTable.Annotation = &entsql.Annotation{
		Table: "sys_post",
	}
	SysRoleTable.Annotation = &entsql.Annotation{
		Table: "sys_role",
	}
	SysUserTable.ForeignKeys[0].RefTable = SysDeptTable
	SysUserTable.Annotation = &entsql.Annotation{
		Table: "sys_user",
	}
	SysUserOnlineTable.Annotation = &entsql.Annotation{
		Table: "sys_user_online",
	}
	SysRoleDeptTable.ForeignKeys[0].RefTable = SysRoleTable
	SysRoleDeptTable.ForeignKeys[1].RefTable = SysDeptTable
	SysRoleMenuTable.ForeignKeys[0].RefTable = SysRoleTable
	SysRoleMenuTable.ForeignKeys[1].RefTable = SysMenuTable
	SysUserPostTable.ForeignKeys[0].RefTable = SysUserTable
	SysUserPostTable.ForeignKeys[1].RefTable = SysPostTable
	SysUserRoleTable.ForeignKeys[0].RefTable = SysUserTable
	SysUserRoleTable.ForeignKeys[1].RefTable = SysRoleTable
}
