package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	commonService "github.com/wangxg422/XishangOS-backend/app/module/common/service"
	"github.com/wangxg422/XishangOS-backend/app/module/system/enmu"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/utils"
)

type SysLogin struct {
	BaseService
}

func (l SysLogin) Login(req *request.SysLoginReq, c *gin.Context) (map[string]any, error) {
	// 如果开启了验证码，校验验证码是否正确
	if global.AppConfig.Captcha.Enabled {
		if commonService.CaptchaService.VerifyString(req.VerifyKey, req.VerifyCode) {
			return nil, errors.New("验证码错误")
		}
	}

	// 按用户名查询用户
	user, err := SysUserService.GetUserByUsername(req.Username, c)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 校验密码是否正确
	enPassword, err := utils.EncryptPasswordWithSalt(req.Password, user.UserSalt)
	if err != nil {
		return nil, err
	}
	if enPassword != user.UserPassword {
		// TODO 记录用户登录失败日志
		return nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态是否正常
	if enmu.UserStatusDisabled.Equals(user.UserStatus) {
		return nil, errors.New("用户已禁用,禁止登录")
	}

	user.UserPassword = ""
	// 登录成功，签发token
	// TODO 签发jwt token
	// TODO 记录登录成功日志

	// 查询用户菜单及权限
	var permissions []string
	var menuList []*codegen.SysMenu
	// 如果是超级管理员，拥有所有权限及菜单
	if user.IsAdmin == 1 {
		permissions = []string{"*/*/*"}
		menuList, err = SysMenuService.GetAllLayoutMenu(c)
		if err != nil {
			return nil, err
		}
	}

	resMap := make(map[string]any)
	resMap["token"] = ""
	resMap["userInfo"] = ""
	resMap["permissions"] = permissions
	resMap["menuList"] = menuList

	return resMap, nil
}
