package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/common/service"
	"github.com/wangxg422/XishangOS-backend/app/module/system/enmu"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	userService "github.com/wangxg422/XishangOS-backend/app/module/system/service"
	"github.com/wangxg422/XishangOS-backend/common/result"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/utils"
)

type SysLoginController struct {
}

func (m *SysLoginController) Login(c *gin.Context) {
	req := &request.SysLoginReq{}
	err := c.ShouldBind(req)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	// 如果开启了验证码，校验验证码是否正确
	if global.AppConfig.Captcha.Enabled {
		if service.AppCaptchaService.VerifyString(req.VerifyKey, req.VerifyCode) {
			result.FailWithMessage("验证码输入错误", c)
			return
		}
	}

	// 按用户名查询用户
	user := userService.AppSysUserService.GetUserByUsername(req.Username)
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}

	// 校验密码是否正确
	if utils.EncryptPassword(req.Password, user.UserSalt) != user.UserPassword {
		result.FailWithMessage("用户名密码错误", c)
		return
	}

	// 检查用户状态是否正常
	if enmu.UserStatusDisabled.Equals(user.UserStatus) {
		result.FailWithMessage("用户已禁用，禁止登录!", c)
		return
	}

	user.UserPassword = ""
	// 登录成功，签发token
	// TODO 签发jwt token

	// 查询用户菜单及权限
	userService.AppSysMenuService.GetUserMenuAndPermissions(user)
}
