package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	commonService "github.com/wangxg422/XishangOS-backend/app/module/common/service"
	"github.com/wangxg422/XishangOS-backend/app/module/system/enmu"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/response"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/common/constant"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/initial/logger"
	"github.com/wangxg422/XishangOS-backend/middleware/jwt"
	"github.com/wangxg422/XishangOS-backend/utils"
	"go.uber.org/zap"
	"time"
)

type SysLogin struct {
	BaseService
}

func (m *SysLogin) Login(req *request.SysLoginReq, c *gin.Context) (*response.LoginRes, error) {
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

	// 登录成功
	user.UserPassword = ""

	// 记录登录成功日志
	loginLog.Msg = "登录成功"
	loginLog.LoginSuccess = 1
	_ = SysLoginLogService.Add(loginLog, c)

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
		// 检查是否是超级管理员
		if r.RoleCode == constant.SuperAdminRoleCode {
			isSuperAdmin = true
			break
		}
		if len(r.Edges.SysMenus) > 0 {
			// TODO 需要将菜单去重
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
	menuList = SysMenuService.BuildMenuTree(SysMenuService.FormatMenu(menus), 0)

	token, err := m.GenToken(user)
	if err != nil {
		loginLog.LoginSuccess = 2
		loginLog.Msg = "token签发失败"
		_ = SysLoginLogService.Add(loginLog, c)
		logger.Error("token签发失败", zap.Error(err))
		return nil, errors.New("token签发失败")
	}

	res := &response.LoginRes{
		Token: token,
		UserInfo: &response.UserInfo{
			Username:    req.Username,
			UserId:      user.ID,
			DeptId:      user.DeptID,
			MenuList:    menuList,
			Permissions: permissions,
		},
	}

	return res, nil
}

// GenToken 签发jwt
func (m *SysLogin) GenToken(user *codegen.SysUser) (string, error) {
	j := &jwt.JWT{SigningKey: []byte(global.AppConfig.Jwt.SigningKey)}
	claims := j.CreateClaims(jwt.BaseClaims{
		UserId:   user.ID,
		NickName: user.UserNickname,
		Username: user.UserName,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", errors.New("token签发失败")
	}
	if !global.AppConfig.App.UseMultipoint {
		return token, nil
	}

	if jwtStr, err := jwt.GetRedisJWT(user.UserName); errors.Is(err, redis.Nil) {
		if err := jwt.SetRedisJWT(token, user.UserName); err != nil {
			return "", errors.New("设置登录状态失败")
		}
	} else if err != nil {
		return "", errors.New("设置登录状态失败")
	} else {
		if _, err := jwt.InBlackList(jwtStr); err != nil {
			return "", errors.New("jwt作废失败")
		}
		if err := jwt.SetRedisJWT(token, user.UserName); err != nil {
			return "", errors.New("设置登录状态失败")
		}
	}
	return token, nil
}
