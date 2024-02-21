package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysmenu"
	"github.com/wangxg422/XishangOS-backend/common/enmu"
)

type SysMenu struct {
}

func (m *SysMenu) List(req *request.SysMenuListReq, c *gin.Context) ([]*codegen.SysMenu, error) {
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

// GetAllMenu 获取所有菜单，含目录、菜单、按钮
func (m *SysMenu) GetAllMenu(c *gin.Context) ([]*codegen.SysMenu, error) {
	return initial.SysDbClient.SysMenu.Query().All(c)
}

// GetAllLayoutMenu GetAllMenu 获取所有菜单，含目录、菜单，不含按钮
func (m *SysMenu) GetAllLayoutMenu(c *gin.Context) ([]*codegen.SysMenu, error) {
	return initial.SysDbClient.SysMenu.Query().
		Where(sysmenu.Or(sysmenu.MenuType(enmu.MenuTypeMenu.Value()), sysmenu.MenuType(enmu.MenuTypeDir.Value()))).
		All(c)
}

func (m *SysMenu) Add(req *request.SysMenuCreateReq, c *gin.Context) error {
	// 先不检查唯一性
	return initial.SysDbClient.SysMenu.Create().
		SetPid(req.Pid).SetMenuName(req.MenuName).SetMenuTitle(req.MenuTitle).
		SetMenuIcon(req.MenuIcon).SetMenuType(req.MenuType).SetCondition(req.Condition).SetPath(req.Path).
		SetComponent(req.Component).SetModuleType(req.ModuleType).SetModelID(req.ModuleId).SetIsHide(req.IsHide).
		SetIsIframe(req.IsIframe).SetIsCached(req.IsCached).SetIsAffix(req.IsAffix).SetIsLink(req.IsLink).
		SetLinkURL(req.LinkUrl).SetRedirect(req.Redirect).SetSort(req.Sort).SetStatus(req.Status).SetRemark(req.Remark).
		Exec(c)
}

func (m *SysMenu) Update(req *request.SysMenuUpdateReq, c *gin.Context) error {
	return initial.SysDbClient.SysMenu.Update().Where(sysmenu.ID(req.Id)).
		SetPid(req.Pid).SetMenuName(req.MenuName).SetMenuTitle(req.MenuTitle).
		SetMenuIcon(req.MenuIcon).SetMenuType(req.MenuType).SetCondition(req.Condition).SetPath(req.Path).
		SetComponent(req.Component).SetModuleType(req.ModuleType).SetModelID(req.ModuleId).SetIsHide(req.IsHide).
		SetIsIframe(req.IsIframe).SetIsCached(req.IsCached).SetIsAffix(req.IsAffix).SetIsLink(req.IsLink).
		SetLinkURL(req.LinkUrl).SetRedirect(req.Redirect).SetSort(req.Sort).SetStatus(req.Status).SetRemark(req.Remark).
		Exec(c)
}

func (m *SysMenu) Delete(id int64, c *gin.Context) (int, error) {
	// 先不检查是否能删除
	return initial.SysDbClient.SysMenu.Delete().Where(sysmenu.ID(id)).Exec(c)
}
