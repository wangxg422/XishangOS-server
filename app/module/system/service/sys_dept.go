package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdept"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/exception"
)

type SysDeptService struct {
}

func (m *SysDeptService) DeptExist(deptCode string, id int64, c *gin.Context) error {
	count, err := initial.SysDbClient.SysDept.Query().Where(sysdept.DeptCode(deptCode), sysdept.IDNEQ(id)).Count(c)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("部门编码已存在")
	}
	return nil
}

func (m *SysDeptService) GetDeptByCode(deptCode string, c *gin.Context) (*codegen.SysDept, error) {
	d, err := initial.SysDbClient.SysDept.Query().Where(sysdept.DeptCode(deptCode)).First(c)
	if exception.NotNoRecordError(err) {
		return nil, err
	}
	return d, nil
}

func (m *SysDeptService) Add(req *request.SysDeptCreateReq, c *gin.Context) error {
	// 校验dept_code唯一
	err := m.DeptExist(req.DeptCode, 0, c)
	if err != nil {
		return err
	}

	return initial.SysDbClient.SysDept.Create().
		SetParentID(req.ParentId).
		SetDeptName(req.DeptName).
		SetDeptCode(req.DeptCode).
		SetEmail(req.Email).
		SetPhone(req.Phone).
		SetAddress(req.Address).
		SetRemark(req.Remark).
		SetLeader(req.Leader).Exec(c)
}

func (m *SysDeptService) Update(req *request.SysDeptUpdateReq, c *gin.Context) error {
	err := m.DeptExist(req.DeptCode, req.Id, c)
	if err != nil {
		return err
	}

	return initial.SysDbClient.SysDept.Update().Where(sysdept.ID(req.Id)).
		SetParentID(req.ParentId).
		SetDeptName(req.DeptName).
		SetDeptCode(req.DeptCode).
		SetEmail(req.Email).
		SetPhone(req.Phone).
		SetAddress(req.Address).
		SetRemark(req.Remark).
		SetLeader(req.Leader).Exec(c)
}

func (m *SysDeptService) List(req *request.SysDeptListReq, c *gin.Context) ([]*codegen.SysDept, error) {
	query := initial.SysDbClient.SysDept.Query()

	if req.Status != 0 {
		query.Where(sysdept.Status(req.Status))
	}
	if req.Keyword != "" {
		query.Where(sysdept.Or(sysdept.DeptCodeContains(req.Keyword)), sysdept.DeptName(req.Keyword), sysdept.LeaderContains(req.Keyword))
	}
	return query.Order(codegen.Asc(sysdept.FieldParentID, sysdept.FieldID, sysdept.FieldCreatedAt)).All(c)
}
