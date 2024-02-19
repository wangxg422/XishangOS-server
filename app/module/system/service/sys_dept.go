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

func (m *SysDeptService) GetDeptByCode(deptCode string, c *gin.Context) (*codegen.SysDept, error) {
	d, err := initial.SysDbClient.SysDept.Query().Where(sysdept.DeptCode(deptCode)).First(c)
	if exception.NotNoRecordError(err) {
		return nil, err
	}
	return d, nil
}

func (m *SysDeptService) Add(req *request.SysDeptCreateReq, c *gin.Context) error {
	// 校验dept_code唯一
	d, err := m.GetDeptByCode(req.DeptCode, c)
	if err != nil {
		return err
	}
	if d != nil {
		return errors.New("部门编码已存在")
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
