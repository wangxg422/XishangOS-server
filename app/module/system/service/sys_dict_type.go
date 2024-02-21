package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/base/model/response"
	"github.com/wangxg422/XishangOS-backend/app/base/util"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdicttype"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/exception"
)

type SysDictType struct {
}

func (m *SysDictType) List(req *request.SysDictTypeListReq, c *gin.Context) (*response.PaginationRes, error) {
	query := initial.SysDbClient.SysDictType.Query()

	if req.Status != 0 {
		query.Where(sysdicttype.Status(req.Status))
	}
	if req.Keyword != "" {
		query.Where(sysdicttype.Or(sysdicttype.DictName(req.Keyword)), sysdicttype.DictType(req.Keyword))
	}

	query.Order(codegen.Asc(sysdicttype.FieldDictType, sysdicttype.FieldCreatedAt))

	total, err := query.Count(c)
	if err != nil {
		return nil, err
	}

	limit, offset := util.PageLimitOffset(req.PageNo, req.PageSize)
	users, err := query.Offset(offset).Limit(limit).All(c)
	if err != nil {
		return nil, err
	}

	res := new(response.PaginationRes)
	res.List = users
	res.PageNo = req.PageNo
	res.Total = total

	return res, nil
}

func (m *SysDictType) Add(req *request.SysDictTypeCreateReq, c *gin.Context) error {
	// 保证 dict_type 唯一
	err := m.SysDictTypeExist(req.DictType, 0, c)
	if err != nil {
		return err
	}

	return initial.SysDbClient.SysDictType.Create().
		SetDictName(req.DictName).
		SetDictType(req.DictType).
		SetStatus(req.Status).
		SetRemark(req.Remark).
		Exec(c)
}

func (m *SysDictType) Update(req *request.SysDictTypeUpdateReq, c *gin.Context) error {
	// 保证 dict_type 唯一
	err := m.SysDictTypeExist(req.DictType, req.Id, c)
	if err != nil {
		return err
	}

	return initial.SysDbClient.SysDictType.Update().
		Where(sysdicttype.ID(req.Id)).
		SetDictName(req.DictName).
		SetDictType(req.DictType).
		SetStatus(req.Status).
		SetRemark(req.Remark).
		Exec(c)
}

func (m *SysDictType) Delete(id int64, c *gin.Context) (int, error) {
	// 校验是否能够删除
	// 删除关联关系

	return initial.SysDbClient.SysDictType.Delete().Where(sysdicttype.ID(id)).Exec(c)
}

func (m *SysDictType) SysDictTypeExist(postCode string, id int64, c *gin.Context) error {
	_, err := initial.SysDbClient.SysDictType.Query().Where(sysdicttype.DictType(postCode), sysdicttype.IDNEQ(id)).First(c)
	if exception.NotNoRecordError(err) {
		return err
	}
	return errors.New("字典已存在")
}
