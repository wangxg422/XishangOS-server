package dao

import (
	"context"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuser"
)

type SysUser struct {
}

func (m *SysUser) List() ([]*codegen.SysUser, error) {
	return initial.SysDbClient.SysUser.Query().All(context.Background())
}

func (m *SysUser) GetUserByUsername(username string) *codegen.SysUser {
	u, err := initial.SysDbClient.SysUser.Query().Where(sysuser.UserNameEQ(username)).First(context.Background())
	if codegen.IsNotFound(err) {
		return nil
	}
	return u
}
