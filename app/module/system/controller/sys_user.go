package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/service"
	"github.com/wangxg422/XishangOS-backend/common/result"
)

type SysUserController struct {
}

func (m *SysUserController) List(c *gin.Context) {
	list, err := service.AppSysUserService.List()
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.OkWithData(list, c)
}
