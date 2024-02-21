package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/response"
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
	return initial.SysDbClient.SysMenu.Query().
		Order(codegen.Asc(sysmenu.FieldPid, sysmenu.FieldSort, sysmenu.FieldCreatedAt)).
		All(c)
}

// GetAllLayoutMenu GetAllMenu 获取所有状态正常的菜单，含目录、菜单，不含按钮
func (m *SysMenu) GetAllLayoutMenu(c *gin.Context) ([]*codegen.SysMenu, error) {
	return initial.SysDbClient.SysMenu.Query().
		Where(sysmenu.Or(sysmenu.MenuType(enmu.MenuTypeMenu.Value()), sysmenu.MenuType(enmu.MenuTypeDir.Value()))).
		Where(sysmenu.Status(enmu.StatusNormal.Value())).
		Order(codegen.Asc(sysmenu.FieldPid, sysmenu.FieldSort, sysmenu.FieldCreatedAt)).
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

// FormatMenu 格式化菜单列表数据为前端需要的格式
func (m *SysMenu) FormatMenu(menus []*codegen.SysMenu) []*response.UserMenu {
	if menus == nil || len(menus) == 0 {
		return []*response.UserMenu{}
	}
	newMenus := make([]*response.UserMenu, len(menus))
	for k, v := range menus {
		newMenus[k] = m.SetMenuData(v)
	}
	return newMenus
}

func (m *SysMenu) SetMenuData(menu *codegen.SysMenu) *response.UserMenu {
	newMenu := &response.UserMenu{
		Id:        menu.ID,
		Pid:       menu.Pid,
		Name:      menu.MenuName,
		Component: menu.Component,
		Path:      menu.Path,
		MenuMeta: &response.MenuMeta{
			Icon:        menu.MenuIcon,
			Title:       menu.MenuTitle,
			IsLink:      "",
			IsHide:      menu.IsHide == 1,
			IsKeepAlive: menu.IsCached == 1,
			IsAffix:     menu.IsAffix == 1,
			IsIframe:    menu.IsIframe == 1,
		},
	}
	if newMenu.MenuMeta.IsIframe || menu.IsLink == 1 {
		newMenu.MenuMeta.IsLink = menu.LinkURL
	}

	return newMenu
}

// BuildMenuTree 由菜单列表构建菜单树
func (m *SysMenu) BuildMenuTree(menus []*response.UserMenu, root int64) []*response.UserMenu {
	menuList := make([]*response.UserMenu, 0)
	for _, v := range menus {
		if v.Pid == root {
			v.Children = m.BuildMenuTree(menus, v.Id)
			menuList = append(menuList, v)
		}
	}
	return menuList
}
