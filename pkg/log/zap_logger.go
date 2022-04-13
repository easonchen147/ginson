package log

import (
	"context"
	"fmt"
	"os"

	"ginson/cfg"
	"ginson/pkg/constant"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	AccessLogger *zap.Logger
	Logger       *zap.Logger
	SqlLogger    *zap.Logger
)

// Init 配置日志模块
func Init(cfg *cfg.AppConfig) {
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
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var core, accessCore, sqlCore zapcore.Core
	switch cfg.LogMode {
	case "console":
		core = zapcore.NewTee(zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(os.Stdout), level))

		accessCore = zapcore.NewTee(zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(os.Stdout), level))
	case "file":
		core = newLoggerCore(cfg.LogFile, core, encoderConfig, level)
		accessCore = newLoggerCore(cfg.AccessLogFile, core, encoderConfig, level)
		sqlCore = newLoggerCore(cfg.SqlLogFile, core, encoderConfig, level)
	}

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	AccessLogger = zap.New(accessCore, zap.AddCaller(), zap.AddCallerSkip(1))
	SqlLogger = zap.New(sqlCore, zap.AddCaller(), zap.AddCallerSkip(1))
}

func newLoggerCore(logFilePath string, core zapcore.Core, encoderConfig zapcore.EncoderConfig, level zapcore.Level) zapcore.Core {
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    500, // megabytes
		MaxBackups: 0,
		MaxAge:     30, // days
		LocalTime:  true,
		Compress:   true,
	})
	return zapcore.NewTee(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(writer), level))
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

func Access(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zapDefaultFields(ctx)...)
	AccessLogger.Info(msg, fields...)
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
