package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/application/service"
	"github.com/wangxg422/XishangOS-backend/common/result"
)

type AppInstanceController struct {
}

func (m *AppInstanceController) List(c *gin.Context) {
	res, err := service.AppAppInstanceService.List(c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.OkWithData(res, c)
}
