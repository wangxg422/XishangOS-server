package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/service"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/param"
	"github.com/wangxg422/XishangOS-backend/common/result"
)

type SysDept struct {
}

func (m *SysDept) List(c *gin.Context) {
	req := new(request.SysDeptListReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	list, err := service.SysDeptService.List(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.OkWithData(list, c)
}

func (m *SysDept) Add(c *gin.Context) {
	req := new(request.SysDeptCreateReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	err = service.SysDeptService.Add(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysDept) Update(c *gin.Context) {
	req := new(request.SysDeptUpdateReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	err = service.SysDeptService.Update(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysDept) Delete(c *gin.Context) {
	id, err := param.ParamInt64("id", c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	_, err = service.SysDeptService.Delete(id, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}
