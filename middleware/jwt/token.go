package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/initial/logger"
	"net"
	"strconv"
	"strings"
	"time"
)

func ClearToken(c *gin.Context) {
	// 增加cookie token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	c.SetCookie(global.AppConfig.Jwt.TokenName, "", -1, "/", host, true, false)
}

func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	c.SetCookie(global.AppConfig.Jwt.TokenName, token, maxAge, "/", host, true, false)
}

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie(global.AppConfig.Jwt.TokenName)
	if token == "" {
		token = c.Request.Header.Get(global.AppConfig.Jwt.TokenName)
	}
	return token
}

func GetClaims(c *gin.Context) (*CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		logger.Error("token解析失败")
	}
	return claims, err
}

// GetUserId 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserId(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.UserId
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.BaseClaims.UserId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户信息
func GetUserInfo(c *gin.Context) *CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.Username
	}
}

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
