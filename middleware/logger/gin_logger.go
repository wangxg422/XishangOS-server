package logger

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wangxg422/XishangOS-backend/common/constant"
	logUtil "github.com/wangxg422/XishangOS-backend/initial/logger"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		uuidStr := strings.ReplaceAll(uuid.New().String(), "-", "")
		userId := c.GetInt64("userId")

		c.Set(constant.TraceCtx, &logUtil.Trace{TraceId: uuidStr, Caller: path, UserId: userId})
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("Status", c.Writer.Status()),
			zap.String("Method", c.Request.Method),
			zap.String("Path", path),
			zap.String("Query", query),
			zap.String("IP", c.ClientIP()),
			zap.String("UserAgent", c.Request.UserAgent()),
			zap.String("Error", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("Cost", cost),
			zap.String("TraceId", uuidStr),
			zap.Int64("UserId", userId),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne.Err, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
