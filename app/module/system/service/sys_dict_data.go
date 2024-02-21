package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/base/model/response"
	"github.com/wangxg422/XishangOS-backend/app/base/util"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdictdata"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/exception"
)

type SysDictData struct {
}

func (m *SysDictData) List(req *request.SysDictDataListReq, c *gin.Context) (*response.PaginationRes, error) {
	query := initial.SysDbClient.SysDictData.Query()

	if req.Status != 0 {
		query.Where(sysdictdata.Status(req.Status))
	}
	if req.Keyword != "" {
		query.Where(sysdictdata.Or(sysdictdata.DictLabel(req.Keyword)), sysdictdata.DictValue(req.Keyword))
	}

	query.Order(codegen.Asc(sysdictdata.FieldSort, sysdictdata.FieldCreatedAt))

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

func (m *SysDictData) Add(req *request.SysDictDataCreateReq, c *gin.Context) error {
	// 保证 dict_value 唯一
	err := m.SysDictDataExist(req.DictValue, 0, c)
	if err != nil {
		return err
	}

	return initial.SysDbClient.SysDictData.Create().
		SetDictLabel(req.DictLabel).SetDictValue(req.DictValue).SetDictTypeID(req.DictTypeId).
		SetCSSClass(req.CssClass).SetListClass(req.ListCss).SetIsDefault(req.IsDefault).
		SetStatus(req.Status).SetRemark(req.Remark).SetSort(req.Sort).
		Exec(c)
}

func (m *SysDictData) Update(req *request.SysDictDataUpdateReq, c *gin.Context) error {
	// 保证 dict_value 唯一
	err := m.SysDictDataExist(req.DictValue, 0, c)
	if err != nil {
		return err
	}

	return initial.SysDbClient.SysDictData.Update().
		Where(sysdictdata.ID(req.Id)).
		SetDictLabel(req.DictLabel).SetDictValue(req.DictValue).SetDictTypeID(req.DictTypeId).
		SetCSSClass(req.CssClass).SetListClass(req.ListCss).SetIsDefault(req.IsDefault).
		SetStatus(req.Status).SetRemark(req.Remark).SetSort(req.Sort).
		Exec(c)
}

func (m *SysDictData) Delete(id int64, c *gin.Context) (int, error) {
	return initial.SysDbClient.SysDictData.Delete().Where(sysdictdata.ID(id)).Exec(c)
}

func (m *SysDictData) SysDictDataExist(postCode string, id int64, c *gin.Context) error {
	_, err := initial.SysDbClient.SysDictData.Query().Where(sysdictdata.DictValue(postCode), sysdictdata.IDNEQ(id)).First(c)
	if exception.NotNoRecordError(err) {
		return err
	}
	return errors.New("字典数据已存在")
}
