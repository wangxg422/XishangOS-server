package request

import (
	"github.com/wangxg422/XishangOS-backend/app/base/model/request"
	"time"
)

type SysLoginLogListReq struct {
	Keyword    string `json:"keyword" form:"keyword"`
	ConfigType int8   `json:"configType" form:"configType"`
	request.Pagination
}

type SysLoginLogCreateReq struct {
	LoginName     string    `json:"loginName"`
	IpAddr        string    `json:"ipAddr"`
	LoginLocation string    `json:"loginLocation"`
	Browser       string    `json:"browser"`
	Os            string    `json:"os"`
	Msg           string    `json:"msg"`
	LoginTime     time.Time `json:"loginTime"`
	Module        string    `json:"module"`
}
