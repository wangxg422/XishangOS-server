package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/dao"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuser"
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

func (m *SysUserService) addUser(tx *codegen.Tx, c *gin.Context, req *request.SysUserCreateReq) error {
	// 检查用户名、手机号是否已经存在,保证用户名唯一、手机号唯一
	if err := dao.SysUserDao.UserNameOrMobileExist(c, req.UserName, req.Mobile, req.Id); err != nil {
		return err
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
func (m *SysUserService) Add(c *gin.Context, req *request.SysUserCreateReq) error {
	return m.WithTx(c, initial.SysDbClient, func(tx *codegen.Tx) error {
		return m.addUser(tx, c, req)
	})
}

func (m *SysUserService) Update(c *gin.Context, req *request.SysUserUpdateReq) error {
	return m.WithTx(c, initial.SysDbClient, func(tx *codegen.Tx) error {
		return m.updateUser(tx, c, req)
	})
}

func (m *SysUserService) updateUser(tx *codegen.Tx, c *gin.Context, req *request.SysUserUpdateReq) error {
	// 检查用户名、手机号是否已经存在,保证用户名唯一、手机号唯一
	if err := dao.SysUserDao.UserNameOrMobileExist(c, req.UserName, req.Mobile, req.Id); err != nil {
		return err
	}

	// 先删除sys_user和sys_post的关联关系
	if err := tx.SysUser.Update().ClearSysPosts().ClearSysRoles().Exec(c); err != nil {
		return err
	}

	// 创建用户
	return tx.SysUser.Update().
		Where(sysuser.ID(req.Id)).
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
}

// GetUserInfo 获取用户的详细信息，包含用户的角色，部门，职位
func (m *SysUserService) GetUserInfo(c *gin.Context, id int64) (*codegen.SysUser, error) {
	return dao.SysUserDao.GetUserInfo(c, id)
}
