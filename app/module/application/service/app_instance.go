package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/application/initial"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen"
)

type AppInstanceService struct {
}

func (m *AppInstanceService) List(c *gin.Context) ([]*codegen.AppInstance, error) {
	return initial.AppDbClient.AppInstance.Query().All(c)
}
