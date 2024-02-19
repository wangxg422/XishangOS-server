package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
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

func (m *SysUserController) Delete(context *gin.Context) {

}

func (m *SysUserController) GetUserInfo(context *gin.Context) {

}
