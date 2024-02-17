package dao

import (
	"context"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
)

type SysUser struct {
}

func (m *SysUser) List() ([]*codegen.SysUser, error) {
	return initial.SysDbClient.SysUser.Query().All(context.Background())
}
