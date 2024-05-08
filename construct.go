package zlog

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewWithBufferedWriteSyncer(writeSyncer *zapcore.BufferedWriteSyncer, level Level) *Logger {
	if writeSyncer == nil {
		writeSyncer = &zapcore.BufferedWriteSyncer{
			WS:   zapcore.AddSync(os.Stdout),
			Size: 4096,
		}
	}

	cores := []zapcore.Core{
		zapcore.NewCore(getDefaultEncoder(), zapcore.AddSync(writeSyncer), level),
	}

	tee := zapcore.NewTee(cores...)
	l := zap.New(tee).WithOptions(zap.AddCallerSkip(1), zap.AddStacktrace(ErrorLevel))

	lg := &Logger{
		zl:  l,
		szl: l.Sugar(),
	}

	return lg
}

func getDefaultEncoder() zapcore.Encoder {
	cfg := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		FunctionKey:   zapcore.OmitKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime:    timeEncoder,
		EncodeLevel:   levelEncoder,
		EncodeCaller:  callerEncoder,
	}

	return zapcore.NewConsoleEncoder(cfg)
}

func levelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(time.DateTime) + "]")
}

func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}
