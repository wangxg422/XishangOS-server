package request

import "github.com/wangxg422/XishangOS-backend/app/base/model/request"

type SysPostListReq struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  int8   `json:"status" form:"status"`
	request.Pagination
}

type SysPostCreateReq struct {
	PostName string `json:"postName"`
	PostCode string `json:"postCode"`
	Sort     int    `json:"sort"`
	Status   int8   `json:"status"`
	Remark   string `json:"remark"`
}

type SysPostUpdateReq struct {
	Id       int64  `json:"id,string"`
	PostName string `json:"postName"`
	PostCode string `json:"postCode"`
	Status   int8   `json:"status"`
	Remark   string `json:"remark"`
	Sort     int    `json:"sort"`
}
