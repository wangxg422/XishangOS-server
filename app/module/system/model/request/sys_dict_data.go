package request

import "github.com/wangxg422/XishangOS-backend/app/base/model/request"

type SysDictDataListReq struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  int8   `json:"status" form:"status"`
	request.Pagination
}

type SysDictDataCreateReq struct {
	DictLabel  string `json:"dictLabel"`
	DictValue  string `json:"dictValue"`
	DictTypeId int64  `json:"dictTypeId"`
	CssClass   string `json:"cssClass"`
	ListCss    string `json:"listCss"`
	IsDefault  int8   `json:"IsDefault"`
	Sort       int    `json:"sort"`
	Status     int8   `json:"status"`
	Remark     string `json:"remark"`
}

type SysDictDataUpdateReq struct {
	Id         int64  `json:"id,string"`
	DictLabel  string `json:"dictLabel"`
	DictValue  string `json:"dictValue"`
	DictTypeId int64  `json:"dictTypeId"`
	CssClass   string `json:"cssClass"`
	ListCss    string `json:"listCss"`
	IsDefault  int8   `json:"IsDefault"`
	Sort       int    `json:"sort"`
	Status     int8   `json:"status"`
	Remark     string `json:"remark"`
}
