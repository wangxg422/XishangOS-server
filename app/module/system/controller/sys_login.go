package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/service"
	"github.com/wangxg422/XishangOS-backend/common/result"
)

type SysLogin struct {
}

func (m *SysLogin) Login(c *gin.Context) {
	req := new(request.SysLoginReq)
	err := c.ShouldBind(&req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	res, err := service.SysLoginService.Login(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.OkWithData(res, c)
}
