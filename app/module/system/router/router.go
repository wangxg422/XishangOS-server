package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/controller"
)

func AddSystemRouter(group *gin.RouterGroup) {
	router := group.Group("system/v1")
	//router.Use(casbin.CasbinHandler())
	{
		router.POST("/login", controller.AppSysLoginController.Login)
		router.GET("/user/list", controller.AppSysUserController.List)
		router.POST("/user/add", controller.AppSysUserController.Add)
		router.PUT("/user/update", controller.AppSysUserController.Update)
		router.DELETE("/user/delete/id", controller.AppSysUserController.Delete)
		router.GET("/user/details/:id", controller.AppSysUserController.GetUserInfo)
	}
	{
		router.GET("/dept/list", controller.AppSysDeptController.List)
		router.POST("/dept/add", controller.AppSysDeptController.Add)
		router.PUT("/dept/update", controller.AppSysDeptController.Update)
		router.DELETE("/dept/delete/id", controller.AppSysDeptController.Delete)
	}
}
