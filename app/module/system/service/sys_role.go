package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/base/model/response"
	"github.com/wangxg422/XishangOS-backend/app/base/util"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysrole"
)

type SysRoleService struct {
}

func (m *SysRoleService) List(req *request.SysRoleListReq, c *gin.Context) (*response.PaginationRes, error) {
	query := initial.SysDbClient.SysRole.Query()

	if req.Status != 0 {
		query.Where(sysrole.Status(req.Status))
	}
	if req.Keyword != "" {
		query.Where(sysrole.Or(sysrole.RoleName(req.Keyword)), sysrole.Remark(req.Keyword))
	}

	query.Order(codegen.Asc(sysrole.FieldSort, sysrole.FieldCreatedAt))

	total, err := query.Count(c)
	if err != nil {
		return nil, err
	}

	limit, offset := util.PageLimitOffset(req.PageNo, req.PageSize)
	list, err := query.Offset(offset).Limit(limit).All(c)
	if err != nil {
		return nil, err
	}

	res := new(response.PaginationRes)
	res.List = list
	res.PageNo = req.PageNo
	res.Total = total

	return res, nil
}

func (m *SysRoleService) Add(req *request.SysRoleCreateReq, c *gin.Context) error {
	return initial.SysDbClient.SysRole.Create().
		SetRoleName(req.RoleName).
		SetDataScope(req.DataScope).
		SetStatus(req.Status).
		SetRemark(req.Remark).
		SetSort(req.Sort).Exec(c)
}

func (m *SysRoleService) Update(req *request.SysRoleUpdateReq, c *gin.Context) error {
	return initial.SysDbClient.SysRole.Update().
		Where(sysrole.ID(req.Id)).
		SetRoleName(req.RoleName).
		SetDataScope(req.DataScope).
		SetStatus(req.Status).
		SetRemark(req.Remark).
		SetSort(req.Sort).Exec(c)
}

func (m *SysRoleService) Delete(id int64, c *gin.Context) (int, error) {
	// 校验是否能够删除
	// 删除关联关系

	return initial.SysDbClient.SysRole.Delete().Where(sysrole.ID(id)).Exec(c)
}

//func (m *SysRoleService) SysRoleExist(postCode string, id int64, c *gin.Context) error {
//	_, err := initial.SysDbClient.SysRole.Query().Where(sysrole.PostCode(postCode), sysrole.IDNEQ(id)).First(c)
//	if exception.NotNoRecordError(err) {
//		return err
//	}
//	return errors.New("岗位编码已存在")
//}
