// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"time"

	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysconfig"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdept"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdictdata"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdicttype"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysloginlog"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysmenu"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysoperlog"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/syspost"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysrole"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuser"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuseronline"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sysconfigMixin := schema.SysConfig{}.Mixin()
	sysconfigMixinFields0 := sysconfigMixin[0].Fields()
	_ = sysconfigMixinFields0
	sysconfigMixinFields1 := sysconfigMixin[1].Fields()
	_ = sysconfigMixinFields1
	sysconfigFields := schema.SysConfig{}.Fields()
	_ = sysconfigFields
	// sysconfigDescCreatedAt is the schema descriptor for created_at field.
	sysconfigDescCreatedAt := sysconfigMixinFields1[0].Descriptor()
	// sysconfig.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysconfig.DefaultCreatedAt = sysconfigDescCreatedAt.Default.(func() time.Time)
	// sysconfigDescUpdatedAt is the schema descriptor for updated_at field.
	sysconfigDescUpdatedAt := sysconfigMixinFields1[1].Descriptor()
	// sysconfig.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysconfig.DefaultUpdatedAt = sysconfigDescUpdatedAt.Default.(func() time.Time)
	// sysconfig.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysconfig.UpdateDefaultUpdatedAt = sysconfigDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysconfigDescDeleteAt is the schema descriptor for delete_at field.
	sysconfigDescDeleteAt := sysconfigMixinFields1[2].Descriptor()
	// sysconfig.DefaultDeleteAt holds the default value on creation for the delete_at field.
	sysconfig.DefaultDeleteAt = sysconfigDescDeleteAt.Default.(func() time.Time)
	// sysconfig.UpdateDefaultDeleteAt holds the default value on update for the delete_at field.
	sysconfig.UpdateDefaultDeleteAt = sysconfigDescDeleteAt.UpdateDefault.(func() time.Time)
	// sysconfigDescID is the schema descriptor for id field.
	sysconfigDescID := sysconfigMixinFields0[0].Descriptor()
	// sysconfig.DefaultID holds the default value on creation for the id field.
	sysconfig.DefaultID = sysconfigDescID.Default.(func() int64)
	sysdeptMixin := schema.SysDept{}.Mixin()
	sysdeptMixinFields0 := sysdeptMixin[0].Fields()
	_ = sysdeptMixinFields0
	sysdeptMixinFields1 := sysdeptMixin[1].Fields()
	_ = sysdeptMixinFields1
	sysdeptMixinFields4 := sysdeptMixin[4].Fields()
	_ = sysdeptMixinFields4
	sysdeptFields := schema.SysDept{}.Fields()
	_ = sysdeptFields
	// sysdeptDescCreatedAt is the schema descriptor for created_at field.
	sysdeptDescCreatedAt := sysdeptMixinFields1[0].Descriptor()
	// sysdept.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysdept.DefaultCreatedAt = sysdeptDescCreatedAt.Default.(func() time.Time)
	// sysdeptDescUpdatedAt is the schema descriptor for updated_at field.
	sysdeptDescUpdatedAt := sysdeptMixinFields1[1].Descriptor()
	// sysdept.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysdept.DefaultUpdatedAt = sysdeptDescUpdatedAt.Default.(func() time.Time)
	// sysdept.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysdept.UpdateDefaultUpdatedAt = sysdeptDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysdeptDescDeleteAt is the schema descriptor for delete_at field.
	sysdeptDescDeleteAt := sysdeptMixinFields1[2].Descriptor()
	// sysdept.DefaultDeleteAt holds the default value on creation for the delete_at field.
	sysdept.DefaultDeleteAt = sysdeptDescDeleteAt.Default.(func() time.Time)
	// sysdept.UpdateDefaultDeleteAt holds the default value on update for the delete_at field.
	sysdept.UpdateDefaultDeleteAt = sysdeptDescDeleteAt.UpdateDefault.(func() time.Time)
	// sysdeptDescStatus is the schema descriptor for status field.
	sysdeptDescStatus := sysdeptMixinFields4[0].Descriptor()
	// sysdept.DefaultStatus holds the default value on creation for the status field.
	sysdept.DefaultStatus = sysdeptDescStatus.Default.(int8)
	// sysdeptDescID is the schema descriptor for id field.
	sysdeptDescID := sysdeptMixinFields0[0].Descriptor()
	// sysdept.DefaultID holds the default value on creation for the id field.
	sysdept.DefaultID = sysdeptDescID.Default.(func() int64)
	sysdictdataMixin := schema.SysDictData{}.Mixin()
	sysdictdataMixinFields0 := sysdictdataMixin[0].Fields()
	_ = sysdictdataMixinFields0
	sysdictdataMixinFields1 := sysdictdataMixin[1].Fields()
	_ = sysdictdataMixinFields1
	sysdictdataMixinFields4 := sysdictdataMixin[4].Fields()
	_ = sysdictdataMixinFields4
	sysdictdataFields := schema.SysDictData{}.Fields()
	_ = sysdictdataFields
	// sysdictdataDescCreatedAt is the schema descriptor for created_at field.
	sysdictdataDescCreatedAt := sysdictdataMixinFields1[0].Descriptor()
	// sysdictdata.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysdictdata.DefaultCreatedAt = sysdictdataDescCreatedAt.Default.(func() time.Time)
	// sysdictdataDescUpdatedAt is the schema descriptor for updated_at field.
	sysdictdataDescUpdatedAt := sysdictdataMixinFields1[1].Descriptor()
	// sysdictdata.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysdictdata.DefaultUpdatedAt = sysdictdataDescUpdatedAt.Default.(func() time.Time)
	// sysdictdata.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysdictdata.UpdateDefaultUpdatedAt = sysdictdataDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysdictdataDescDeleteAt is the schema descriptor for delete_at field.
	sysdictdataDescDeleteAt := sysdictdataMixinFields1[2].Descriptor()
	// sysdictdata.DefaultDeleteAt holds the default value on creation for the delete_at field.
	sysdictdata.DefaultDeleteAt = sysdictdataDescDeleteAt.Default.(func() time.Time)
	// sysdictdata.UpdateDefaultDeleteAt holds the default value on update for the delete_at field.
	sysdictdata.UpdateDefaultDeleteAt = sysdictdataDescDeleteAt.UpdateDefault.(func() time.Time)
	// sysdictdataDescStatus is the schema descriptor for status field.
	sysdictdataDescStatus := sysdictdataMixinFields4[0].Descriptor()
	// sysdictdata.DefaultStatus holds the default value on creation for the status field.
	sysdictdata.DefaultStatus = sysdictdataDescStatus.Default.(int8)
	// sysdictdataDescID is the schema descriptor for id field.
	sysdictdataDescID := sysdictdataMixinFields0[0].Descriptor()
	// sysdictdata.DefaultID holds the default value on creation for the id field.
	sysdictdata.DefaultID = sysdictdataDescID.Default.(func() int64)
	sysdicttypeMixin := schema.SysDictType{}.Mixin()
	sysdicttypeMixinFields0 := sysdicttypeMixin[0].Fields()
	_ = sysdicttypeMixinFields0
	sysdicttypeMixinFields1 := sysdicttypeMixin[1].Fields()
	_ = sysdicttypeMixinFields1
	sysdicttypeMixinFields4 := sysdicttypeMixin[4].Fields()
	_ = sysdicttypeMixinFields4
	sysdicttypeFields := schema.SysDictType{}.Fields()
	_ = sysdicttypeFields
	// sysdicttypeDescCreatedAt is the schema descriptor for created_at field.
	sysdicttypeDescCreatedAt := sysdicttypeMixinFields1[0].Descriptor()
	// sysdicttype.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysdicttype.DefaultCreatedAt = sysdicttypeDescCreatedAt.Default.(func() time.Time)
	// sysdicttypeDescUpdatedAt is the schema descriptor for updated_at field.
	sysdicttypeDescUpdatedAt := sysdicttypeMixinFields1[1].Descriptor()
	// sysdicttype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysdicttype.DefaultUpdatedAt = sysdicttypeDescUpdatedAt.Default.(func() time.Time)
	// sysdicttype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysdicttype.UpdateDefaultUpdatedAt = sysdicttypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysdicttypeDescDeleteAt is the schema descriptor for delete_at field.
	sysdicttypeDescDeleteAt := sysdicttypeMixinFields1[2].Descriptor()
	// sysdicttype.DefaultDeleteAt holds the default value on creation for the delete_at field.
	sysdicttype.DefaultDeleteAt = sysdicttypeDescDeleteAt.Default.(func() time.Time)
	// sysdicttype.UpdateDefaultDeleteAt holds the default value on update for the delete_at field.
	sysdicttype.UpdateDefaultDeleteAt = sysdicttypeDescDeleteAt.UpdateDefault.(func() time.Time)
	// sysdicttypeDescStatus is the schema descriptor for status field.
	sysdicttypeDescStatus := sysdicttypeMixinFields4[0].Descriptor()
	// sysdicttype.DefaultStatus holds the default value on creation for the status field.
	sysdicttype.DefaultStatus = sysdicttypeDescStatus.Default.(int8)
	// sysdicttypeDescID is the schema descriptor for id field.
	sysdicttypeDescID := sysdicttypeMixinFields0[0].Descriptor()
	// sysdicttype.DefaultID holds the default value on creation for the id field.
	sysdicttype.DefaultID = sysdicttypeDescID.Default.(func() int64)
	sysloginlogMixin := schema.SysLoginLog{}.Mixin()
	sysloginlogMixinFields0 := sysloginlogMixin[0].Fields()
	_ = sysloginlogMixinFields0
	sysloginlogMixinFields1 := sysloginlogMixin[1].Fields()
	_ = sysloginlogMixinFields1
	sysloginlogFields := schema.SysLoginLog{}.Fields()
	_ = sysloginlogFields
	// sysloginlogDescStatus is the schema descriptor for status field.
	sysloginlogDescStatus := sysloginlogMixinFields1[0].Descriptor()
	// sysloginlog.DefaultStatus holds the default value on creation for the status field.
	sysloginlog.DefaultStatus = sysloginlogDescStatus.Default.(int8)
	// sysloginlogDescID is the schema descriptor for id field.
	sysloginlogDescID := sysloginlogMixinFields0[0].Descriptor()
	// sysloginlog.DefaultID holds the default value on creation for the id field.
	sysloginlog.DefaultID = sysloginlogDescID.Default.(func() int64)
	sysmenuMixin := schema.SysMenu{}.Mixin()
	sysmenuMixinFields0 := sysmenuMixin[0].Fields()
	_ = sysmenuMixinFields0
	sysmenuMixinFields1 := sysmenuMixin[1].Fields()
	_ = sysmenuMixinFields1
	sysmenuFields := schema.SysMenu{}.Fields()
	_ = sysmenuFields
	// sysmenuDescCreatedAt is the schema descriptor for created_at field.
	sysmenuDescCreatedAt := sysmenuMixinFields1[0].Descriptor()
	// sysmenu.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysmenu.DefaultCreatedAt = sysmenuDescCreatedAt.Default.(func() time.Time)
	// sysmenuDescUpdatedAt is the schema descriptor for updated_at field.
	sysmenuDescUpdatedAt := sysmenuMixinFields1[1].Descriptor()
	// sysmenu.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysmenu.DefaultUpdatedAt = sysmenuDescUpdatedAt.Default.(func() time.Time)
	// sysmenu.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysmenu.UpdateDefaultUpdatedAt = sysmenuDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysmenuDescDeleteAt is the schema descriptor for delete_at field.
	sysmenuDescDeleteAt := sysmenuMixinFields1[2].Descriptor()
	// sysmenu.DefaultDeleteAt holds the default value on creation for the delete_at field.
	sysmenu.DefaultDeleteAt = sysmenuDescDeleteAt.Default.(func() time.Time)
	// sysmenu.UpdateDefaultDeleteAt holds the default value on update for the delete_at field.
	sysmenu.UpdateDefaultDeleteAt = sysmenuDescDeleteAt.UpdateDefault.(func() time.Time)
	// sysmenuDescID is the schema descriptor for id field.
	sysmenuDescID := sysmenuMixinFields0[0].Descriptor()
	// sysmenu.DefaultID holds the default value on creation for the id field.
	sysmenu.DefaultID = sysmenuDescID.Default.(func() int64)
	sysoperlogMixin := schema.SysOperLog{}.Mixin()
	sysoperlogMixinFields0 := sysoperlogMixin[0].Fields()
	_ = sysoperlogMixinFields0
	sysoperlogFields := schema.SysOperLog{}.Fields()
	_ = sysoperlogFields
	// sysoperlogDescID is the schema descriptor for id field.
	sysoperlogDescID := sysoperlogMixinFields0[0].Descriptor()
	// sysoperlog.DefaultID holds the default value on creation for the id field.
	sysoperlog.DefaultID = sysoperlogDescID.Default.(func() int64)
	syspostMixin := schema.SysPost{}.Mixin()
	syspostMixinFields0 := syspostMixin[0].Fields()
	_ = syspostMixinFields0
	syspostMixinFields1 := syspostMixin[1].Fields()
	_ = syspostMixinFields1
	syspostMixinFields4 := syspostMixin[4].Fields()
	_ = syspostMixinFields4
	syspostFields := schema.SysPost{}.Fields()
	_ = syspostFields
	// syspostDescCreatedAt is the schema descriptor for created_at field.
	syspostDescCreatedAt := syspostMixinFields1[0].Descriptor()
	// syspost.DefaultCreatedAt holds the default value on creation for the created_at field.
	syspost.DefaultCreatedAt = syspostDescCreatedAt.Default.(func() time.Time)
	// syspostDescUpdatedAt is the schema descriptor for updated_at field.
	syspostDescUpdatedAt := syspostMixinFields1[1].Descriptor()
	// syspost.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	syspost.DefaultUpdatedAt = syspostDescUpdatedAt.Default.(func() time.Time)
	// syspost.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	syspost.UpdateDefaultUpdatedAt = syspostDescUpdatedAt.UpdateDefault.(func() time.Time)
	// syspostDescDeleteAt is the schema descriptor for delete_at field.
	syspostDescDeleteAt := syspostMixinFields1[2].Descriptor()
	// syspost.DefaultDeleteAt holds the default value on creation for the delete_at field.
	syspost.DefaultDeleteAt = syspostDescDeleteAt.Default.(func() time.Time)
	// syspost.UpdateDefaultDeleteAt holds the default value on update for the delete_at field.
	syspost.UpdateDefaultDeleteAt = syspostDescDeleteAt.UpdateDefault.(func() time.Time)
	// syspostDescStatus is the schema descriptor for status field.
	syspostDescStatus := syspostMixinFields4[0].Descriptor()
	// syspost.DefaultStatus holds the default value on creation for the status field.
	syspost.DefaultStatus = syspostDescStatus.Default.(int8)
	// syspostDescID is the schema descriptor for id field.
	syspostDescID := syspostMixinFields0[0].Descriptor()
	// syspost.DefaultID holds the default value on creation for the id field.
	syspost.DefaultID = syspostDescID.Default.(func() int64)
	sysroleMixin := schema.SysRole{}.Mixin()
	sysroleMixinFields0 := sysroleMixin[0].Fields()
	_ = sysroleMixinFields0
	sysroleMixinFields1 := sysroleMixin[1].Fields()
	_ = sysroleMixinFields1
	sysroleMixinFields2 := sysroleMixin[2].Fields()
	_ = sysroleMixinFields2
	sysroleFields := schema.SysRole{}.Fields()
	_ = sysroleFields
	// sysroleDescCreatedAt is the schema descriptor for created_at field.
	sysroleDescCreatedAt := sysroleMixinFields1[0].Descriptor()
	// sysrole.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysrole.DefaultCreatedAt = sysroleDescCreatedAt.Default.(func() time.Time)
	// sysroleDescUpdatedAt is the schema descriptor for updated_at field.
	sysroleDescUpdatedAt := sysroleMixinFields1[1].Descriptor()
	// sysrole.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysrole.DefaultUpdatedAt = sysroleDescUpdatedAt.Default.(func() time.Time)
	// sysrole.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysrole.UpdateDefaultUpdatedAt = sysroleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysroleDescDeleteAt is the schema descriptor for delete_at field.
	sysroleDescDeleteAt := sysroleMixinFields1[2].Descriptor()
	// sysrole.DefaultDeleteAt holds the default value on creation for the delete_at field.
	sysrole.DefaultDeleteAt = sysroleDescDeleteAt.Default.(func() time.Time)
	// sysrole.UpdateDefaultDeleteAt holds the default value on update for the delete_at field.
	sysrole.UpdateDefaultDeleteAt = sysroleDescDeleteAt.UpdateDefault.(func() time.Time)
	// sysroleDescStatus is the schema descriptor for status field.
	sysroleDescStatus := sysroleMixinFields2[0].Descriptor()
	// sysrole.DefaultStatus holds the default value on creation for the status field.
	sysrole.DefaultStatus = sysroleDescStatus.Default.(int8)
	// sysroleDescID is the schema descriptor for id field.
	sysroleDescID := sysroleMixinFields0[0].Descriptor()
	// sysrole.DefaultID holds the default value on creation for the id field.
	sysrole.DefaultID = sysroleDescID.Default.(func() int64)
	sysuserMixin := schema.SysUser{}.Mixin()
	sysuserMixinFields0 := sysuserMixin[0].Fields()
	_ = sysuserMixinFields0
	sysuserMixinFields1 := sysuserMixin[1].Fields()
	_ = sysuserMixinFields1
	sysuserFields := schema.SysUser{}.Fields()
	_ = sysuserFields
	// sysuserDescCreatedAt is the schema descriptor for created_at field.
	sysuserDescCreatedAt := sysuserMixinFields1[0].Descriptor()
	// sysuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	sysuser.DefaultCreatedAt = sysuserDescCreatedAt.Default.(func() time.Time)
	// sysuserDescUpdatedAt is the schema descriptor for updated_at field.
	sysuserDescUpdatedAt := sysuserMixinFields1[1].Descriptor()
	// sysuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sysuser.DefaultUpdatedAt = sysuserDescUpdatedAt.Default.(func() time.Time)
	// sysuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	sysuser.UpdateDefaultUpdatedAt = sysuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sysuserDescDeleteAt is the schema descriptor for delete_at field.
	sysuserDescDeleteAt := sysuserMixinFields1[2].Descriptor()
	// sysuser.DefaultDeleteAt holds the default value on creation for the delete_at field.
	sysuser.DefaultDeleteAt = sysuserDescDeleteAt.Default.(func() time.Time)
	// sysuser.UpdateDefaultDeleteAt holds the default value on update for the delete_at field.
	sysuser.UpdateDefaultDeleteAt = sysuserDescDeleteAt.UpdateDefault.(func() time.Time)
	// sysuserDescID is the schema descriptor for id field.
	sysuserDescID := sysuserMixinFields0[0].Descriptor()
	// sysuser.DefaultID holds the default value on creation for the id field.
	sysuser.DefaultID = sysuserDescID.Default.(func() int64)
	sysuseronlineMixin := schema.SysUserOnline{}.Mixin()
	sysuseronlineMixinFields0 := sysuseronlineMixin[0].Fields()
	_ = sysuseronlineMixinFields0
	sysuseronlineFields := schema.SysUserOnline{}.Fields()
	_ = sysuseronlineFields
	// sysuseronlineDescID is the schema descriptor for id field.
	sysuseronlineDescID := sysuseronlineMixinFields0[0].Descriptor()
	// sysuseronline.DefaultID holds the default value on creation for the id field.
	sysuseronline.DefaultID = sysuseronlineDescID.Default.(func() int64)
}
