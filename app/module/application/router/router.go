package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/application/controller"
	"github.com/wangxg422/XishangOS-backend/global"
)

func AddAppRouter(group *gin.RouterGroup) {
	appRouter := group.Group("/app/v1")

	router := appRouter
	if global.AppConfig.Module.Application.ApiVersionEnabled {
		if global.AppConfig.Module.Application.ApiVersion != "" {
			router = appRouter.Group(global.AppConfig.Module.Application.ApiVersion)
		}
	}

	{
		router.GET("/instance/list", controller.AppInstanceController.List)
	}
}
