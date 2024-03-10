package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/common/constant"
	"go.uber.org/zap"
)

func Debug(tag string, fields ...zap.Field) {
	zapLog.Debug(tag, fields...)
}

func DebugF(ctx *gin.Context, tag string, fields ...zap.Field) {
	t, exist := ctx.Get(constant.TraceCtx)
	if !exist {
		Debug(tag, fields...)
	}
	trace := t.(*Trace)

	Debug(tag,
		append(fields, zap.String("traceId", trace.TraceId), zap.Int64("userId", trace.UserId))...,
	)
}

func Info(tag string, fields ...zap.Field) {
	zapLog.Info(tag, fields...)
}

func InfoF(ctx *gin.Context, tag string, fields ...zap.Field) {
	t, exist := ctx.Get(constant.TraceCtx)
	if !exist {
		Info(tag, fields...)
	}
	trace := t.(*Trace)

	Info(tag,
		append(fields, zap.String("traceId", trace.TraceId), zap.Int64("userId", trace.UserId))...,
	)
}

func Warn(tag string, fields ...zap.Field) {
	zapLog.Warn(tag, fields...)
}

func WarnF(ctx *gin.Context, tag string, fields ...zap.Field) {
	t, exist := ctx.Get(constant.TraceCtx)
	if !exist {
		Warn(tag, fields...)
	}
	trace := t.(*Trace)

	Warn(tag,
		append(fields, zap.String("traceId", trace.TraceId), zap.Int64("userId", trace.UserId))...,
	)
}

func Error(tag string, fields ...zap.Field) {
	zapLog.Error(tag, fields...)
}

func ErrorF(ctx *gin.Context, tag string, fields ...zap.Field) {
	t, exist := ctx.Get(constant.TraceCtx)
	if !exist {
		Error(tag, fields...)
	}
	trace := t.(*Trace)

	Error(tag,
		append(fields, zap.String("traceId", trace.TraceId), zap.Int64("userId", trace.UserId))...,
	)
}

func Fatal(tag string, fields ...zap.Field) {
	zapLog.Fatal(tag, fields...)
}

func FatalF(ctx *gin.Context, tag string, fields ...zap.Field) {
	t, exist := ctx.Get(constant.TraceCtx)
	if !exist {
		Fatal(tag, fields...)
	}
	trace := t.(*Trace)

	Fatal(tag,
		append(fields, zap.String("traceId", trace.TraceId), zap.Int64("userId", trace.UserId))...,
	)
}
