package casbin

import (
	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求path
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 角色应该从token解析出来，此处为了节约时间，写死了值
		sub := "admin"
		// 引入casbin
		e := AppCasbinService.GetCasbinEnforcer()

		//判断策略是否存在
		success, _ := e.Enforce(sub, obj, act)
		//如果环境变量是开发者模式或者casbin校验通过
		if success {
			c.Next()
		} else {
			c.Abort()
			return
		}
	}
}
