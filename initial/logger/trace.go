package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/common/constant"
)

// Trace 定义trace结构体
type Trace struct {
	TraceId   string  `json:"traceId"`
	SpanId    string  `json:"spanId"`
	Caller    string  `json:"caller"`
	SrcMethod *string `json:"srcMethod,omitempty"`
	UserId    int64   `json:"userId"`
}

// GetTraceCtx 根据gin的context获取context，使log trace更加通用
func GetTraceCtx(c *gin.Context) *Trace {
	return c.MustGet(constant.TraceCtx).(*Trace)
}
