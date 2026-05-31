package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

// Init 初始化日志
func Init(filename string, level string, maxSize, maxBackups, maxAge int, compress bool) {
	// 日志级别
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	// 日志编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 日志轮转
	writer := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(writer), zapcore.AddSync(os.Stdout)),
		zapLevel,
	)

	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// Debug ...
func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

// Info ...
func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

// Warn ...
func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

// Error ...
func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

// Fatal ...
func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}