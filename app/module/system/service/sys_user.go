package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/wangxg422/XishangOS-backend/app/module/system/dao"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/syspost"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/exception"
	"github.com/wangxg422/XishangOS-backend/middleware/casbin"
)

type SysUserService struct {
}

func (m *SysUserService) List() ([]*codegen.SysUser, error) {
	return dao.SysUserDao.List()
}

func (m *SysUserService) GetUserByUsername(username string) (*codegen.SysUser, error) {
	return dao.SysUserDao.GetUserByUsername(username)
}

func (m *SysUserService) Add(c *gin.Context, req *request.SysUserCreateUpdateReq) error {
	// 检查用户是否已经存在,保证用户名唯一
	existUser, err := dao.SysUserDao.GetUserByUsername(req.UserName)
	if existUser != nil {
		return errors.New("用户已存在")
	}
	if err != nil {
		return err
	}

	// 使用事务添加用户
	tx, err := initial.SysDbClient.Tx(c)
	if err != nil {
		return err
	}

	// 创建用户
	user, err := tx.SysUser.Create().
		SetUserName(req.UserName).
		SetUserPassword(req.Password).
		SetDeptID(req.DeptId).
		SetUserEmail(req.Email).
		SetUserNickname(req.NickName).
		SetMobile(req.Mobile).
		SetRemark(req.Remark).
		SetSex(req.Sex).
		SetUserStatus(req.Status).
		SetAddress(req.Address).Save(c)
	if err != nil {
		return exception.Rollback(tx, fmt.Errorf("failed creating the group: %w", err))
	}

	// 关联角色
	enforcer := casbin.AppCasbinService.GetCasbinEnforcer()
	for _, v := range req.RoleIds {
		_, err = enforcer.AddGroupingPolicy(fmt.Sprintf("%s%d", "u_", user.ID), gconv.String(v))
	}

	// 添加用户岗位
	if len(req.PostIds) != 0 {
		// 先删除原先岗位
		_, err := tx.SysPost.Delete().Where(syspost.IDIn(req.PostIds...)).Exec(c)
		if err != nil {
			return exception.Rollback(tx, fmt.Errorf("failed creating the group: %w", err))
		}
		// 添加新的岗位

	}

	return tx.Commit()
}
