package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/common/controller"
	"github.com/wangxg422/XishangOS-backend/global"
)

func AddCommonPublicRouter(group *gin.RouterGroup) {
	publicRouter := group.Group("api/pub/v1")

	router := publicRouter
	if global.AppConfig.Module.Public.ApiVersionEnabled {
		if global.AppConfig.Module.Public.ApiVersion != "" {
			router = publicRouter.Group(global.AppConfig.Module.Public.ApiVersion)
		}
	}

	{
		router.GET("/captcha/get", controller.CaptchaController.GetCaptchaImgString)
	}

	group.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ok",
		})
	})
}

func AddCommonRouter(group *gin.RouterGroup) {
	commonRouter := group.Group("common")

	router := commonRouter
	if global.AppConfig.Module.Common.ApiVersionEnabled {
		if global.AppConfig.Module.Common.ApiVersion != "" {
			router = commonRouter.Group(global.AppConfig.Module.Common.ApiVersion)
		}
	}

	{
		router.GET("/test", func(context *gin.Context) {
		})
	}
}
