package service

import (
	"github.com/wangxg422/XishangOS-backend/app/module/system/dao"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
)

type SysUser struct {
}

func (m *SysUser) List() ([]*codegen.SysUser, error) {
	return dao.SysUserDao.List()
}
