package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/base/controller"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/service"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/param"
	"github.com/wangxg422/XishangOS-backend/common/result"
)

type SysPost struct {
	controller.BaseController
}

func (m *SysPost) List(c *gin.Context) {
	req := new(request.SysPostListReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	res, err := service.SysPostService.List(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.OkWithData(res, c)
}

func (m *SysPost) Add(c *gin.Context) {
	req := new(request.SysPostCreateReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	err = service.SysPostService.Add(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysPost) Update(c *gin.Context) {
	req := new(request.SysPostUpdateReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	err = service.SysPostService.Update(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysPost) Delete(c *gin.Context) {
	id, err := param.ParamInt64("id", c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	_, err = service.SysPostService.Delete(id, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}
