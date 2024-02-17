package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/dao"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
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
	// 检查用户是否已经存在
	existUser, err := dao.SysUserDao.GetUserByUsername(req.UserName)
	if existUser.ID != 0 {
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
	user, err := tx.SysUser.Create().
		SetUserName(req.UserName).
		SetUserPassword(req.Password).
		SetDeptID(req.DeptId).
		SetUserEmail(req.Email).
		SetUserNickname(req.NickName).
		SetMobile(req.Mobile)).
		SetRemark(req.Remark).
		SetSex(req.Sex).
		SetUserStatus(req.Status).
		SetAddress(req.Sex).Save(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("failed creating the group: %w", err))
	}
	// 创建 admin 组
	dan, err := tx.User.
		Create().
		SetAge(29).
		SetName("Dan").
		AddManage(hub).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	// 创建 "Ariel" 用户。
	a8m, err := tx.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		AddGroups(hub).
		AddFriends(dan).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	fmt.Println(a8m)
	// 输出:
	// User(id=2, age=30, name=Ariel)

	// 提交事务
	return tx.Commit()
}

// 如果发生错误则调用 tx.Rollback 回滚并返回嵌套的错误信息
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
	dao.SysUserDao.Add(req)
}
