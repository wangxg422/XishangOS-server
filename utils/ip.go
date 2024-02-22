package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetClientIp(c *gin.Context) string {
	if c.ClientIP() != "" {
		return c.ClientIP()
	}
	ip := ""
	realIps := c.GetHeader("X-Forwarded-For")
	if realIps != "" && len(realIps) != 0 && !strings.EqualFold("unknown", realIps) {
		ipArray := strings.Split(realIps, ",")
		ip = ipArray[0]
	}
	if ip == "" {
		ip = c.GetHeader("Proxy-Client-IP")
	}
	if ip == "" {
		ip = c.GetHeader("WL-Proxy-Client-IP")
	}
	if ip == "" {
		ip = c.GetHeader("HTTP_CLIENT_IP")
	}
	if ip == "" {
		ip = c.GetHeader("HTTP_X_FORWARDED_FOR")
	}
	if ip == "" {
		ip = c.GetHeader("X-Real-IP")
	}

	return ip
}
