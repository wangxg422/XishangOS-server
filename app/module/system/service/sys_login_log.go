package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/base/model/response"
	"github.com/wangxg422/XishangOS-backend/app/base/util"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysloginlog"
)

type SysLoginLog struct {
	BaseService
}

func (m *SysLoginLog) List(req *request.SysLoginLogListReq, c *gin.Context) (*response.PaginationRes, error) {
	query := initial.SysDbClient.SysLoginLog.Query()

	if req.Keyword != "" {
		query.Where(sysloginlog.Or(sysloginlog.Msg(req.Keyword)), sysloginlog.Module(req.Keyword), sysloginlog.LoginName(req.Keyword))
	}

	query.Order(codegen.Desc(sysloginlog.FieldLoginTime))

	total, err := query.Count(c)
	if err != nil {
		return nil, err
	}

	limit, offset := util.PageLimitOffset(req.PageNo, req.PageSize)
	list, err := query.Offset(offset).Limit(limit).All(c)
	if err != nil {
		return nil, err
	}

	res := new(response.PaginationRes)
	res.List = list
	res.PageNo = req.PageNo
	res.Total = total

	return res, nil
}

func (m *SysLoginLog) Add(req *request.SysLoginLogCreateReq, c *gin.Context) error {
	return initial.SysDbClient.SysLoginLog.Create().
		SetLoginName(req.LoginName).SetIPAddr(req.IpAddr).SetLoginLocation(req.LoginLocation).
		SetBrowser(req.Browser).SetOs(req.Os).SetMsg(req.Msg).SetLoginTime(req.LoginTime).SetModule(req.Module).
		Exec(c)
}
