package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/common/controller"
)

func AddCommonPublicRouter(group *gin.RouterGroup) {
	router := group.Group("api/pub/v1")
	{
		router.GET("/captcha/get", controller.AppCaptchaController.GetCaptchaImgString)
	}

	group.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ok",
		})
	})
}

func AddCommonRouter(group *gin.RouterGroup) {
	router := group.Group("common")
	{
		router.GET("/test", func(context *gin.Context) {
		})
	}
}
