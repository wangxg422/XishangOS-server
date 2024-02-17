package service

import (
	"github.com/wangxg422/XishangOS-backend/app/module/system/dao"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
)

type SysMenuService struct {
}

func (m *SysMenuService) List() ([]*codegen.SysUser, error) {
	return dao.SysUserDao.List()
}

func (m *SysMenuService) GetUserMenuAndPermissions(user *codegen.SysUser) {
	
}
