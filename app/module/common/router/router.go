package router

import "github.com/gin-gonic/gin"

func AddCommonPublicRouter(group *gin.RouterGroup) {
	router := group.Group("public")
	{
		router.GET("", func(context *gin.Context) {

		})
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
