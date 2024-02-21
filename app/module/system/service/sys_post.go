package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/base/model/response"
	"github.com/wangxg422/XishangOS-backend/app/base/util"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/syspost"
	"github.com/wangxg422/XishangOS-backend/app/module/system/util/exception"
)

type SysPost struct {
}

func (m *SysPost) List(req *request.SysPostListReq, c *gin.Context) (*response.PaginationRes, error) {
	query := initial.SysDbClient.SysPost.Query()

	if req.Status != 0 {
		query.Where(syspost.Status(req.Status))
	}
	if req.Keyword != "" {
		query.Where(syspost.Or(syspost.PostCodeContains(req.Keyword)), syspost.PostName(req.Keyword))
	}

	query.Order(codegen.Asc(syspost.FieldSort, syspost.FieldPostCode, syspost.FieldCreatedAt))

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

func (m *SysPost) Add(req *request.SysPostCreateReq, c *gin.Context) error {
	// 保证岗位编码唯一
	err := m.SysPostExist(req.PostCode, 0, c)
	if err != nil {
		return err
	}

	return initial.SysDbClient.SysPost.Create().
		SetPostName(req.PostName).
		SetPostCode(req.PostCode).
		SetStatus(req.Status).
		SetRemark(req.Remark).
		SetSort(req.Sort).Exec(c)
}

func (m *SysPost) Update(req *request.SysPostUpdateReq, c *gin.Context) error {
	// 保证岗位编码唯一
	err := m.SysPostExist(req.PostCode, 0, c)
	if err != nil {
		return err
	}

	return initial.SysDbClient.SysPost.Update().
		Where(syspost.ID(req.Id)).
		SetPostCode(req.PostCode).
		SetPostName(req.PostName).
		SetStatus(req.Status).
		SetRemark(req.Remark).
		SetSort(req.Sort).Exec(c)
}

func (m *SysPost) Delete(id int64, c *gin.Context) (int, error) {
	// 校验是否能够删除
	// 删除关联关系

	return initial.SysDbClient.SysPost.Delete().Where(syspost.ID(id)).Exec(c)
}

func (m *SysPost) SysPostExist(postCode string, id int64, c *gin.Context) error {
	_, err := initial.SysDbClient.SysPost.Query().Where(syspost.PostCode(postCode), syspost.IDNEQ(id)).First(c)
	if exception.NotNoRecordError(err) {
		return err
	}
	return errors.New("岗位编码已存在")
}
