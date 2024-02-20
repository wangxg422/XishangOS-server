package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/base/model/response"
	"github.com/wangxg422/XishangOS-backend/app/base/util"
	"github.com/wangxg422/XishangOS-backend/app/module/system/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/request"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysconfig"
)

type SysConfigService struct {
	BaseService
}

func (m SysConfigService) Add(req *request.SysConfigCreateReq, c *gin.Context) error {
	// 保证config_key唯一
	err := m.SysConfigKeyExist(req.ConfigKey, 0, c)
	if err != nil {
		return err
	}

	return initial.SysDbClient.SysConfig.Create().
		SetConfigName(req.ConfigName).
		SetConfigKey(req.ConfigKey).
		SetConfigValue(req.ConfigValue).
		SetConfigType(req.ConfigType).
		SetRemark(req.Remark).
		Exec(c)
}

func (m SysConfigService) SysConfigKeyExist(key string, id int64, c *gin.Context) error {
	count, err := initial.SysDbClient.SysConfig.Query().Where(sysconfig.ConfigKey(key), sysconfig.IDNEQ(id)).Count(c)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("系统配置已存在")
	}
	return nil
}

func (m SysConfigService) Update(req *request.SysConfigUpdateReq, c *gin.Context) error {
	// 保证config_key唯一
	err := m.SysConfigKeyExist(req.ConfigKey, req.Id, c)
	if err != nil {
		return err
	}

	return initial.SysDbClient.SysConfig.Update().
		Where(sysconfig.ID(req.Id)).
		SetConfigName(req.ConfigName).
		SetConfigKey(req.ConfigKey).
		SetConfigValue(req.ConfigValue).
		SetConfigType(req.ConfigType).
		SetRemark(req.Remark).
		Exec(c)
}

func (m SysConfigService) List(req *request.SysConfigListReq, c *gin.Context) (*response.PaginationRes, error) {
	query := initial.SysDbClient.SysConfig.Query()
	if req.ConfigType != 0 {
		query.Where(sysconfig.ConfigType(req.ConfigType))
	}
	if req.Keyword != "" {
		query.Where(sysconfig.Or(sysconfig.ConfigNameContains(req.Keyword), sysconfig.ConfigKeyContains(req.Keyword), sysconfig.ConfigValueContains(req.Keyword)))
	}
	query.Order(codegen.Asc(sysconfig.FieldCreatedAt, sysconfig.FieldConfigKey))

	total, err := query.Count(c)
	if err != nil {
		return nil, err
	}
	limit, offset := util.PageLimitOffset(req.PageNo, req.PageSize)
	all, err := query.Offset(offset).Limit(limit).All(c)

	res := new(response.PaginationRes)
	res.List = all
	res.Total = total
	res.PageNo = req.PageNo

	return res, nil
}

func (m SysConfigService) Delete(id int64, c *gin.Context) (int, error) {
	return initial.SysDbClient.SysConfig.Delete().Where(sysconfig.ID(id)).Exec(c)
}
