package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysmenu"
)

type SysMenuService struct {
}

func (m *SysMenuService) List(req *request.SysMenuListReq, c *gin.Context) ([]*codegen.SysMenu, error) {
	query := initial.SysDbClient.SysMenu.Query()
	if req.Status != 0 {
		query.Where(sysmenu.Status(req.Status))
	}
	if req.Keyword != "" {
		query.Where(sysmenu.Or(sysmenu.MenuNameContains(req.Keyword), sysmenu.MenuTitle(req.Keyword),
			sysmenu.PathContains(req.Keyword), sysmenu.ComponentContains(req.Keyword)))
	}
	query.Order(codegen.Asc(sysmenu.FieldPid, sysmenu.FieldSort, sysmenu.FieldCreatedAt))
	return query.All(c)
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

func (m *SysMenuService) Add(req *request.SysMenuCreateReq, c *gin.Context) error {
	// 先不检查唯一性
	return initial.SysDbClient.SysMenu.Create().
		SetPid(req.Pid).SetMenuName(req.MenuName).SetMenuTitle(req.MenuTitle).
		SetMenuIcon(req.MenuIcon).SetMenuType(req.MenuType).SetCondition(req.Condition).SetPath(req.Path).
		SetComponent(req.Component).SetModuleType(req.ModuleType).SetModelID(req.ModuleId).SetIsHide(req.IsHide).
		SetIsIframe(req.IsIframe).SetIsCached(req.IsCached).SetIsAffix(req.IsAffix).SetIsLink(req.IsLink).
		SetLinkURL(req.LinkUrl).SetRedirect(req.Redirect).SetSort(req.Sort).SetStatus(req.Status).SetRemark(req.Remark).
		Exec(c)
}

func (m *SysMenuService) Update(req *request.SysMenuUpdateReq, c *gin.Context) error {
	return initial.SysDbClient.SysMenu.Update().Where(sysmenu.ID(req.Id)).
		SetPid(req.Pid).SetMenuName(req.MenuName).SetMenuTitle(req.MenuTitle).
		SetMenuIcon(req.MenuIcon).SetMenuType(req.MenuType).SetCondition(req.Condition).SetPath(req.Path).
		SetComponent(req.Component).SetModuleType(req.ModuleType).SetModelID(req.ModuleId).SetIsHide(req.IsHide).
		SetIsIframe(req.IsIframe).SetIsCached(req.IsCached).SetIsAffix(req.IsAffix).SetIsLink(req.IsLink).
		SetLinkURL(req.LinkUrl).SetRedirect(req.Redirect).SetSort(req.Sort).SetStatus(req.Status).SetRemark(req.Remark).
		Exec(c)
}

func (m *SysMenuService) Delete(id int64, c *gin.Context) (int, error) {
	// 先不检查是否能删除
	return initial.SysDbClient.SysMenu.Delete().Where(sysmenu.ID(id)).Exec(c)
}
