package request

import "github.com/wangxg422/XishangOS-backend/app/base/model/request"

type SysMenuListReq struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  int8   `json:"status" form:"status"`
	request.Pagination
}

type SysMenuCreateReq struct {
	Pid        int64  `json:"pId,string"`
	MenuName   string `json:"menuName"`
	MenuTitle  string `json:"menuTitle"`
	MenuIcon   string `json:"menuIcon"`
	MenuType   int8   `json:"menuType"`
	Condition  string `json:"condition"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	ModuleType string `json:"moduleType"`
	ModuleId   int64  `json:"moduleId"`
	IsHide     int8   `json:"isHide"`
	IsIframe   int8   `json:"isIframe"`
	IsCached   int8   `json:"isCached"`
	IsAffix    int8   `json:"isAffix"`
	IsLink     int8   `json:"isLink"`
	LinkUrl    string `json:"linkUrl"`
	Redirect   string `json:"redirect"`
	Sort       int    `json:"sort"`
	Status     int8   `json:"status"`
	Remark     string `json:"remark"`
}

type SysMenuUpdateReq struct {
	Id         int64  `json:"id,string"`
	Pid        int64  `json:"pId,string"`
	MenuName   string `json:"menuName"`
	MenuTitle  string `json:"menuTitle"`
	MenuIcon   string `json:"menuIcon"`
	MenuType   int8   `json:"menuType"`
	Condition  string `json:"condition"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	ModuleType string `json:"moduleType"`
	ModuleId   int64  `json:"moduleId"`
	IsHide     int8   `json:"isHide"`
	IsIframe   int8   `json:"isIframe"`
	IsCached   int8   `json:"isCached"`
	IsAffix    int8   `json:"isAffix"`
	IsLink     int8   `json:"isLink"`
	LinkUrl    string `json:"linkUrl"`
	Redirect   string `json:"redirect"`
	Sort       int    `json:"sort"`
	Status     int8   `json:"status"`
	Remark     string `json:"remark"`
}
