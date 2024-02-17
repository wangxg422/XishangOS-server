package initial

import (
	"github.com/wangxg422/XishangOS-backend/app/router"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/middleware/logger"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// debug or release
	gin.SetMode(global.AppConfig.App.Mode)
	r := gin.New()
	// 使用日志框架接管gin日志
	r.Use(logger.GinLogger(global.Log), logger.GinRecovery(global.Log, true))

	if addr := global.AppConfig.App.Addresses; addr == "" || addr == "0.0.0.0" {

	} else {
		_ = r.SetTrustedProxies(strings.Split(global.AppConfig.App.Addresses, ","))
	}

	router.AddRouter(r)

	_ = r.Run(":" + global.AppConfig.App.Port)
}
