package service

import (
	"github.com/wangxg422/XishangOS-backend/app/module/system/dao"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
)

type SysUserService struct {
}

func (m *SysUserService) List() ([]*codegen.SysUser, error) {
	return dao.SysUserDao.List()
}

func (m *SysUserService) GetUserByUsername(username string) (*codegen.SysUser, error) {
	return dao.SysUserDao.GetUserByUsername(username)
}
