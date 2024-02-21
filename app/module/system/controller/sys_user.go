package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/service"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/param"
	"github.com/wangxg422/XishangOS-backend/common/result"
)

type SysUser struct {
}

func (m *SysUser) List(c *gin.Context) {
	req := new(request.SysUserListReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	res, err := service.SysUserService.List(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.OkWithData(res, c)
}

func (m *SysUser) Add(c *gin.Context) {
	req := new(request.SysUserCreateReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	err = service.SysUserService.Add(c, req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysUser) Update(c *gin.Context) {
	req := new(request.SysUserUpdateReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	err = service.SysUserService.Update(c, req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysUser) Delete(c *gin.Context) {
	id, err := param.ParamInt64("id", c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	_, err = service.SysUserService.Delete(id, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysUser) GetUserInfo(c *gin.Context) {
	id, err := param.ParamInt64("id", c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	user, err := service.SysUserService.GetUserInfo(c, id)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.OkWithData(user, c)
}
