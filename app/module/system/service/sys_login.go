package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	commonService "github.com/wangxg422/XishangOS-backend/app/module/common/service"
	"github.com/wangxg422/XishangOS-backend/app/module/system/enmu"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/response"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/initial/logger"
	"github.com/wangxg422/XishangOS-backend/utils"
	"go.uber.org/zap"
	"time"
)

type SysLogin struct {
	BaseService
}

func (m *SysLogin) Login(req *request.SysLoginReq, c *gin.Context) (map[string]any, error) {
	logger.WarnF(c, "log", zap.String("test", "666"))

	// 如果开启了验证码，校验验证码是否正确
	if global.AppConfig.Captcha.Enabled {
		if commonService.CaptchaService.VerifyString(req.VerifyKey, req.VerifyCode) {
			return nil, errors.New("验证码错误")
		}
	}

	loginLog := &request.SysLoginLogCreateReq{
		LoginName: req.Username,
		IpAddr:    utils.GetClientIp(c),
		Browser:   c.GetHeader("User-Agent"),
		LoginTime: time.Now(),
		Module:    "system",
	}

	// 按用户名查询用户
	user, err := SysUserService.GetUserByUsername(req.Username, c)
	if err != nil {
		loginLog.Msg = err.Error()
		_ = SysLoginLogService.Add(loginLog, c)
		return nil, err
	}

	if user == nil {
		loginLog.LoginSuccess = 2
		loginLog.Msg = "用户名密码错误"
		_ = SysLoginLogService.Add(loginLog, c)
		return nil, errors.New("用户名或密码错误")
	}

	// 校验密码是否正确
	enPassword, err := utils.EncryptPasswordWithSalt(req.Password, user.UserSalt)
	if err != nil {
		return nil, err
	}
	if enPassword != user.UserPassword {
		loginLog.LoginSuccess = 2
		loginLog.Msg = "用户名密码错误"
		_ = SysLoginLogService.Add(loginLog, c)
		return nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态是否正常
	if enmu.UserStatusDisabled.Equals(user.UserStatus) {
		loginLog.LoginSuccess = 2
		loginLog.Msg = "用户已被禁止登录"
		_ = SysLoginLogService.Add(loginLog, c)
		return nil, errors.New("用户已禁用,禁止登录")
	}

	user.UserPassword = ""
	// 登录成功，签发token
	// TODO 签发jwt token

	// 记录登录成功日志
	loginLog.Msg = "登录成功"
	loginLog.LoginSuccess = 1
	_ = SysLoginLogService.Add(loginLog, c)

	resMap := make(map[string]any)

	// 查询用户角色，角色关联的菜单、权限
	var permissions []string
	var menuList []*response.UserMenu

	userInfo, err := SysUserService.GetUserRoleAndMenu(user.ID, c)
	if err != nil {
		return nil, err
	}

	isSuperAdmin := false
	roles := userInfo.Edges.SysRoles
	menus := make([]*codegen.SysMenu, 0)
	for _, r := range roles {
		// 检查是否是超级管理员 TODO 更改判定条件
		if r.ID == 11 {
			isSuperAdmin = true
			break
		}
		if len(r.Edges.SysMenus) > 0 {
			menus = append(menus, r.Edges.SysMenus...)
		}
	}

	if isSuperAdmin {
		// 如果是超级管理员，拥有所有权限及菜单
		permissions = []string{"*/*/*"}
		menus, err = SysMenuService.GetAllLayoutMenu(c)
		if err != nil {
			return nil, err
		}
	}

	// 格式化菜单数据为前端需要的格式并构建菜单树
	menuList = SysMenuService.FormatMenu(menus)
	resMap["menuList"] = SysMenuService.BuildMenuTree(menuList, 0)

	resMap["token"] = ""
	resMap["userInfo"] = userInfo
	resMap["permissions"] = permissions

	return resMap, nil
}
