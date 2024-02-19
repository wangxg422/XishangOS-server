package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/dao"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
)

type SysUserService struct {
	BaseService
}

func (m *SysUserService) List() ([]*codegen.SysUser, error) {
	return dao.SysUserDao.List()
}

func (m *SysUserService) GetUserByUsername(username string) *codegen.SysUser {
	return dao.SysUserDao.GetUserByUsername(username)
}

func (m *SysUserService) addUser(tx *codegen.Tx, c *gin.Context, req *request.SysUserCreateUpdateReq) error {
	// 检查用户是否已经存在,保证用户名唯一
	existUser := dao.SysUserDao.GetUserByUsername(req.UserName)
	if existUser != nil {
		return errors.New("用户已存在")
	}

	// 创建用户
	return tx.SysUser.Create().
		SetUserName(req.UserName).
		SetUserPassword(req.Password).
		SetDeptID(req.DeptId).
		SetUserEmail(req.Email).
		SetUserNickname(req.NickName).
		SetMobile(req.Mobile).
		SetRemark(req.Remark).
		SetSex(req.Sex).
		SetUserStatus(req.Status).
		SetAddress(req.Address).
		AddSysPostIDs(req.PostIds...).
		AddSysRoleIDs(req.RoleIds...).
		SetDeptID(req.DeptId).
		Exec(c)

	// 关联角色 TODO 向权限数据库中添加权限数据
	//enforcer := casbin.AppCasbinService.GetCasbinEnforcer()
	//for _, v := range req.RoleIds {
	//	_, err = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", "u_", user.ID), gconv.String(v))
	//}
}
func (m *SysUserService) Add(c *gin.Context, req *request.SysUserCreateUpdateReq) error {
	return m.WithTx(c, initial.SysDbClient, func(tx *codegen.Tx) error {
		return m.addUser(tx, c, req)
	})
}
