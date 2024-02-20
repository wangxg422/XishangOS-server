package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/application/controller"
)

func AddAppRouter(group *gin.RouterGroup) {
	router := group.Group("/app/v1")
	{
		router.GET("/instance/list", controller.AppAppInstanceController.List)
	}
}
