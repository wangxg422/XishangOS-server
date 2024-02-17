package router

import (
	"github.com/gin-gonic/gin"
)

func AddAppRouter(group *gin.RouterGroup) {
	router := group.Group("app")
	{
		router.GET("/instance/list", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "ok",
			})
		})
	}
}
