package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/base/controller"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/service"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/param"
	"github.com/wangxg422/XishangOS-backend/common/result"
)

type SysMenuController struct {
	controller.BaseController
}

func (m *SysMenuController) List(c *gin.Context) {
	req := new(request.SysMenuListReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	res, err := service.AppSysMenuService.List(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.OkWithData(res, c)
}

func (m *SysMenuController) Add(c *gin.Context) {
	req := new(request.SysMenuCreateReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	err = service.AppSysMenuService.Add(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysMenuController) Update(c *gin.Context) {
	req := new(request.SysMenuUpdateReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	err = service.AppSysMenuService.Update(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysMenuController) Delete(c *gin.Context) {
	id, err := param.ParamInt64("id", c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	_, err = service.AppSysMenuService.Delete(id, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}
