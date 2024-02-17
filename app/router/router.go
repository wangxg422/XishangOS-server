package router

import "github.com/gin-gonic/gin"
import commonRouter "github.com/wangxg422/XishangOS-backend/app/module/common/router"
import sysRouter "github.com/wangxg422/XishangOS-backend/app/module/system/router"
import appRouter "github.com/wangxg422/XishangOS-backend/app/module/application/router"

func AddRouter(r *gin.Engine) {
	// 开放接口,健康检查、登录、注册、忘记密码、接口文档等
	publicGroup := r.Group("")
	{
		commonRouter.AddCommonPublicRouter(publicGroup)
	}

	// 模块路由
	moduleGroup := r.Group("/api")
	{
		// 注册公共模块路由
		commonRouter.AddCommonRouter(moduleGroup)
		// 注册系统模块路由
		sysRouter.AddSystemRouter(moduleGroup)
		// 注册应用模块路由
		appRouter.AddAppRouter(moduleGroup)
	}
}
