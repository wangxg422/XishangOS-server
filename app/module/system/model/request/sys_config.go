package request

import "github.com/wangxg422/XishangOS-backend/app/base/model/request"

type SysConfigListReq struct {
	Keyword    string `json:"keyword" form:"keyword"`
	ConfigType int8   `json:"configType" form:"configType"`
	request.Pagination
}

type SysConfigCreateReq struct {
	ConfigName  string `json:"configName"`
	ConfigKey   string `json:"configKey"`
	ConfigValue string `json:"configValue"`
	ConfigType  int8   `json:"configType"`
	Remark      string `json:"remark"`
}

type SysConfigUpdateReq struct {
	Id          int64  `json:"id,string"`
	ConfigName  string `json:"configName"`
	ConfigKey   string `json:"configKey"`
	ConfigValue string `json:"configValue"`
	ConfigType  int8   `json:"configType"`
	Remark      string `json:"remark"`
}
