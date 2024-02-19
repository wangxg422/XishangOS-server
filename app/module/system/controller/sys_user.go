package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/service"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/param"
	"github.com/wangxg422/XishangOS-backend/common/result"
)

type SysUserController struct {
}

func (m *SysUserController) List(c *gin.Context) {
	req := new(request.SysUserListReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	res, err := service.AppSysUserService.List(req, c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.OkWithData(res, c)
}

func (m *SysUserController) Add(c *gin.Context) {
	req := new(request.SysUserCreateReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	err = service.AppSysUserService.Add(c, req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysUserController) Update(c *gin.Context) {
	req := new(request.SysUserUpdateReq)
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	err = service.AppSysUserService.Update(c, req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.Ok(c)
}

func (m *SysUserController) Delete(c *gin.Context) {

}

func (m *SysUserController) GetUserInfo(c *gin.Context) {
	id, err := param.ParamInt64("id", c)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	user, err := service.AppSysUserService.GetUserInfo(c, id)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	result.OkWithData(user, c)
}
