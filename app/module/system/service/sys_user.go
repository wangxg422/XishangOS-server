package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/base/model/response"
	"github.com/wangxg422/XishangOS-backend/app/base/util"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysuser"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/exception"
)

type SysUserService struct {
	BaseService
}

func (m *SysUserService) List(req *request.SysUserListReq, c *gin.Context) (*response.PaginationRes, error) {
	query := initial.SysDbClient.SysUser.Query()
	if req.Keyword != "" {
		query.Where(sysuser.Or(sysuser.UserNameContains(req.Keyword), sysuser.UserNicknameContains(req.Keyword)))
	}
	if req.Mobile != "" {
		query.Where(sysuser.MobileContains(req.Mobile))
	}
	if req.Email != "" {
		query.Where(sysuser.UserEmailContains(req.Email))
	}
	if req.Sex != 0 {
		query.Where(sysuser.Sex(req.Sex))
	}
	if req.Status != 0 {
		query.Where(sysuser.UserStatus(req.Status))
	}
	if req.DeptId != 0 {
		query.Where(sysuser.DeptID(req.DeptId))
	}
	if req.OrderBy != "" {
		query.Order(codegen.Asc(sysuser.FieldCreatedAt))
	}

	total, err := query.Count(c)
	if err != nil {
		return nil, err
	}

	limit, offset := util.PageLimitOffset(req.PageNo, req.PageSize)
	users, err := query.WithSysDept().WithSysRoles().WithSysPosts().Offset(offset).Limit(limit).All(c)
	if err != nil {
		return nil, err
	}

	res := new(response.PaginationRes)
	res.List = users
	res.PageNo = req.PageNo
	res.Total = total

	return res, nil
}

func (m *SysUserService) GetUserByUsername(username string, c *gin.Context) (*codegen.SysUser, error) {
	u, err := initial.SysDbClient.SysUser.Query().Where(sysuser.UserNameEQ(username)).First(c)
	if exception.NotNoRecordError(err) {
		return nil, err
	}
	return u, err
}

func (m *SysUserService) addUser(tx *codegen.Tx, c *gin.Context, req *request.SysUserCreateReq) error {
	// 检查用户名、手机号是否已经存在,保证用户名唯一、手机号唯一
	if err := m.UserNameOrMobileExist(c, req.UserName, req.Mobile, 0); err != nil {
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
	if err := m.UserNameOrMobileExist(c, req.UserName, req.Mobile, req.Id); err != nil {
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
	u, err := initial.SysDbClient.SysUser.Query().
		Where(sysuser.IDEQ(id)).
		WithSysPosts().WithSysRoles().WithSysDept().
		First(c)
	if exception.NotNoRecordError(err) {
		return nil, err
	}
	return u, nil
}

// UserNameOrMobileExist 按照username和手机号查询用户是否已经存在
func (m *SysUserService) UserNameOrMobileExist(c *gin.Context, username string, mobile string, id int64) error {
	user, err := initial.SysDbClient.SysUser.Query().
		Where(sysuser.Or(sysuser.UserNameEQ(username), sysuser.Mobile(mobile))).
		Where(sysuser.IDNEQ(id)).
		First(c)

	if exception.NotNoRecordError(err) {
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

func (m *SysUserService) Delete(id int64, c *gin.Context) (interface{}, error) {
	return initial.SysDbClient.SysUser.Delete().Where(sysuser.ID(id)).Exec(c)
}
