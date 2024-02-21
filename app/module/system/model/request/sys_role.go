package request

import "github.com/wangxg422/XishangOS-backend/app/base/model/request"

type SysRoleListReq struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  int8   `json:"status" form:"status"`
	request.Pagination
}

type SysRoleCreateReq struct {
	RoleName  string `json:"roleName"`
	DataScope int8   `json:"dataScope"`
	Sort      int    `json:"sort"`
	Status    int8   `json:"status"`
	Remark    string `json:"remark"`
}

type SysRoleUpdateReq struct {
	Id        int64  `json:"id,string"`
	RoleName  string `json:"roleName"`
	DataScope int8   `json:"dataScope"`
	Sort      int    `json:"sort"`
	Status    int8   `json:"status"`
	Remark    string `json:"remark"`
}
