package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wangxg422/XishangOS-backend/common/result"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/middleware/logger"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetToken(c)
		if token == "" {
			result.FailWithMessage("未登录或非法访问", c)
			c.Abort()
			return
		}
		//if jwtService.IsBlacklist(token) {
		//	response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		//	utils.ClearToken(c)
		//	c.Abort()
		//	return
		//}
		j := NewJWT()

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			result.FailWithMessage(err.Error(), c)
			ClearToken(c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := ParseDuration(global.AppConfig.Jwt.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)

			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			SetToken(c, newToken, int(dr.Seconds()))
			if !global.AppConfig.App.UseMultipoint {
				redisJwtToken, err := GetRedisJWT(newClaims.Username)
				if err != nil {
					logger.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_, _ = ToBlackList(redisJwtToken)
				}
				// 无论如何都要记录当前的活跃状态
				_ = SetRedisJWT(redisJwtToken, newClaims.Username)
			}
		}
	}
}
