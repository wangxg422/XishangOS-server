package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/system/controller"
)

func AddSystemRouter(group *gin.RouterGroup) {
	router := group.Group("system/v1")
	//router.Use(casbin.CasbinHandler())
	{
		router.POST("/login", controller.SysLoginController.Login)
		router.GET("/user/list", controller.SysUserController.List)
		router.POST("/user/add", controller.SysUserController.Add)
		router.PUT("/user/update", controller.SysUserController.Update)
		router.DELETE("/user/delete/id", controller.SysUserController.Delete)
		router.GET("/user/details/:id", controller.SysUserController.GetUserInfo)
	}
	{
		router.GET("/dept/list", controller.SysDeptController.List)
		router.POST("/dept/add", controller.SysDeptController.Add)
		router.PUT("/dept/update", controller.SysDeptController.Update)
		router.DELETE("/dept/delete/id", controller.SysDeptController.Delete)
	}
	{
		router.GET("/post/list", controller.SysPostController.List)
		router.POST("/post/add", controller.SysPostController.Add)
		router.PUT("/post/update", controller.SysPostController.Update)
		router.DELETE("/post/delete/id", controller.SysPostController.Delete)
	}
	{
		router.GET("/config/list", controller.SysConfigController.List)
		router.POST("/config/add", controller.SysConfigController.Add)
		router.PUT("/config/update", controller.SysConfigController.Update)
		router.DELETE("/config/delete/:id", controller.SysConfigController.Delete)
	}
}
