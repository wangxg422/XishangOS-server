package service

import (
	"context"
	"github.com/wangxg422/XishangOS-backend/app/module/system/dao"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
)

type SysMenuService struct {
}

func (m *SysMenuService) List() ([]*codegen.SysUser, error) {
	return dao.SysUserDao.List()
}

func (m *SysMenuService) GetUserMenuAndPermissions(user *codegen.SysUser) (menuList []*codegen.SysMenu, permissions []string, err error) {
	// 超级管理员获取全部菜单，并设置顶级权限[*/*/*]
	if user.IsAdmin == 1 {
		permissions = []string{"*/*/*"}
		//menuList, err := m.GetAllMenu()
		if err != nil {
			return nil, nil, err
		}
	}

	return nil, nil, nil
}

func (m *SysMenuService) GetAllMenu() ([]*codegen.SysMenu, error) {
	return initial.SysDbClient.SysMenu.Query().All(context.Background())
}
