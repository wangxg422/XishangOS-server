package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/controller"
	"github.com/wangxg422/XishangOS-backend/middleware/casbin"
)

func AddSystemRouter(group *gin.RouterGroup) {
	router := group.Group("system/v1")
	router.Use(casbin.CasbinHandler())
	{
		router.POST("/login", controller.AppSysLoginController.Login)
		router.GET("/user/list", controller.AppSysUserController.List)
		router.POST("/user/add", controller.AppSysLoginController.Add)
	}
}
