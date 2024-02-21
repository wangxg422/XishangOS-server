package request

import "github.com/wangxg422/XishangOS-backend/app/base/model/request"

type SysDictTypeListReq struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  int8   `json:"status" form:"status"`
	request.Pagination
}

type SysDictTypeCreateReq struct {
	DictName string `json:"dictName"`
	DictType string `json:"dictType"`
	Status   int8   `json:"status"`
	Remark   string `json:"remark"`
}

type SysDictTypeUpdateReq struct {
	Id       int64  `json:"id,string"`
	DictName string `json:"dictName"`
	DictType string `json:"dictType"`
	Status   int8   `json:"status"`
	Remark   string `json:"remark"`
}
