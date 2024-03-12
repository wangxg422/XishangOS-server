package response

import "github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"

type SysMenuTree struct {
	*codegen.SysMenu
	Children []*SysMenuTree `json:"children"`
}

type UserMenu struct {
	Id        int64  `json:"id,string"`
	Pid       int64  `json:"pid,string"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	*MenuMeta `json:"meta"`
	Children  []*UserMenu `json:"children"`
}

type MenuMeta struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	IsLink      string `json:"isLink"`
	IsHide      bool   `json:"isHide"`
	IsFull      bool   `json:"isFull"`
	IsKeepAlive bool   `json:"isKeepAlive"`
	IsAffix     bool   `json:"isAffix"`
	IsIframe    bool   `json:"isIframe"`
	ActiveMenu  string `json:"activeMenu"`
}
