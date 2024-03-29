package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/wangxg422/XishangOS-backend/global"
	"github.com/wangxg422/XishangOS-backend/utils"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger

func GetLogger() *zap.Logger {
	return zapLog
}

func InitLog() {
	// 如果不存在日志目录，则新建
	if ok, _ := utils.PathExists(global.AppConfig.Zap.Path); !ok {
		fmt.Printf("create %v directory\n", global.AppConfig.Zap.Path)
		_ = os.Mkdir(global.AppConfig.Zap.Path, os.ModePerm)
	}

	cores := GetZapCores()
	zapLog = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1))
}

func GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := transportLevel(global.AppConfig.Zap.Level); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, getEncoderCore(level, getLevelPriority(level)))
	}
	return cores
}

func getEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	// 使用file-rotatelogs进行日志分割
	writer, err := GetWriteSyncer(l.String())
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err)
		return nil
	}

	return zapcore.NewCore(getEncoder(), writer, level)
}

func getEncoder() zapcore.Encoder {
	if global.AppConfig.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	} else if global.AppConfig.Zap.Format == "console" {
		return zapcore.NewConsoleEncoder(getEncoderConfig())
	} else {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.AppConfig.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    getEncodeLevel(global.AppConfig.Zap.EncodeLevel),
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.AppConfig.Zap.Prefix + t.Format("2006/01/02 15:04:05.000"))
}

func getLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool {
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool {
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool {
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool {
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool {
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool {
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	}
}

func getEncodeLevel(encodeLevel string) zapcore.LevelEncoder {
	switch {
	case encodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case encodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case encodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case encodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func transportLevel(level string) zapcore.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	logFile := path.Join(global.AppConfig.Zap.Path, global.AppConfig.App.Name+"-%Y-%m-%d"+".log")
	if global.AppConfig.Zap.FileShunt {
		logFile = path.Join(global.AppConfig.Zap.Path, global.AppConfig.App.Name+"-%Y-%m-%d-"+level+".log")
	}

	fileWriter, err := rotatelogs.New(
		logFile,
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.AppConfig.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.AppConfig.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
