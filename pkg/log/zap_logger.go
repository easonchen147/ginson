package log

import (
	"context"
	"fmt"
	"ginson/pkg/constant"
	"os"

	"ginson/pkg/conf"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Init 配置日志模块
func Init(cfg *conf.AppConfig) {
	var level zapcore.Level
	if level.UnmarshalText([]byte(cfg.LogLevel)) != nil {
		level = zapcore.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		NameKey:        "name",
		TimeKey:        "time",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		CallerKey:      "location",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var core zapcore.Core
	switch cfg.LogMode {
	case "console":
		core = zapcore.NewTee(zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), level))
	case "file":
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.LogFile,
			MaxSize:    500, // megabytes
			MaxBackups: 0,
			MaxAge:     28, // days
			LocalTime:  true,
		})
		core = zapcore.NewTee(zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(writer), level))
	}
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Debug(ctx context.Context, msg string, val ...interface{}) {
	Logger.Debug(fmt.Sprintf(msg, val...), zapDefaultFields(ctx)...)
}

func Info(ctx context.Context, msg string, val ...interface{}) {
	Logger.Info(fmt.Sprintf(msg, val...), zapDefaultFields(ctx)...)
}

func Warn(ctx context.Context, msg string, val ...interface{}) {
	Logger.Warn(fmt.Sprintf(msg, val...), zapDefaultFields(ctx)...)
}

func Error(ctx context.Context, msg string, val ...interface{}) {
	Logger.Error(fmt.Sprintf(msg, val...), zapDefaultFields(ctx)...)
}

func Panic(ctx context.Context, msg string, val ...interface{}) {
	Logger.Panic(fmt.Sprintf(msg, val...), zapDefaultFields(ctx)...)
}

func zapDefaultFields(ctx context.Context) []zap.Field {
	fields := make([]zap.Field, 0)
	fields = append(fields, zap.String("traceId", getTraceId(ctx)))
	return fields
}

func getTraceId(ctx context.Context) string {
	obj := ctx.Value(constant.TraceIdKey)
	if obj == nil {
		return ""
	}
	traceId, ok := obj.(string)
	if !ok {
		return ""
	}
	return traceId
}
