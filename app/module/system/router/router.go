package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/controller"
)

func AddSystemRouter(group *gin.RouterGroup) {
	router := group.Group("system/v1")
	{
		router.POST("/login", controller.AppSysLoginController.Login)
		router.GET("/user/list", controller.AppSysUserController.List)
		router.POST("/user/add", controller.AppSysLoginController.Add)
	}
}
