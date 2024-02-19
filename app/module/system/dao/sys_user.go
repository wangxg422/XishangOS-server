package dao

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuser"
)

type SysUser struct {
}

func (m *SysUser) List() ([]*codegen.SysUser, error) {
	return initial.SysDbClient.SysUser.Query().All(context.Background())
}

func (m *SysUser) GetUserByUsername(username string) *codegen.SysUser {
	u, err := initial.SysDbClient.SysUser.Query().Where(sysuser.UserNameEQ(username)).First(context.Background())
	if codegen.IsNotFound(err) {
		return nil
	}
	return u
}

// UserNameOrMobileExist 按照username和手机号查询用户是否已经存在
func (m *SysUser) UserNameOrMobileExist(c *gin.Context, username string, mobile string, id int64) error {
	user, err := initial.SysDbClient.SysUser.Query().
		Where(sysuser.Or(sysuser.UserNameEQ(username), sysuser.Mobile(mobile))).
		Where(sysuser.IDNEQ(id)).
		First(c)

	if err != nil && codegen.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}

	if user != nil && user.UserName == username {
		return errors.New("用户名已存在")
	}
	if user != nil && user.Mobile == mobile {
		return errors.New("手机号已存在")
	}
	return nil
}

func (m *SysUser) GetUserInfo(c *gin.Context, id int64) (*codegen.SysUser, error) {
	return initial.SysDbClient.SysUser.Query().Where(sysuser.IDEQ(id)).WithSysPosts().WithSysRoles().WithSysDept().First(c)
}
