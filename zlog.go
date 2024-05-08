package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	DebugLevel  = zapcore.DebugLevel
	InfoLevel   = zapcore.InfoLevel
	WarnLevel   = zapcore.WarnLevel
	ErrorLevel  = zapcore.ErrorLevel
	DPanicLevel = zapcore.DPanicLevel
	PanicLevel  = zapcore.PanicLevel
	FatalLevel  = zapcore.FatalLevel
)

type Logger struct {
	zl  *zap.Logger
	szl *zap.SugaredLogger
}

var std *Logger

func Sync() error {
	return std.zl.Sync()
}

func L() *zap.Logger {
	return std.zl.WithOptions(zap.AddCallerSkip(-1))
}

func SL() *zap.SugaredLogger {
	return std.szl.WithOptions(zap.AddCallerSkip(-1))
}

// -------------------- slogger functions mapping --------------------
func Debug(args ...interface{}) {
	std.szl.Debug(args...)
}

func Info(args ...interface{}) {
	std.szl.Info(args...)
}

func Warn(args ...interface{}) {
	std.szl.Warn(args...)
}

func Error(args ...interface{}) {
	std.szl.Error(args...)
}

func DPanic(args ...interface{}) {
	std.szl.DPanic(args...)
}

func Panic(args ...interface{}) {
	std.szl.Panic(args...)
}

func Fatal(args ...interface{}) {
	std.szl.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	std.szl.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	std.szl.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	std.szl.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	std.szl.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	std.szl.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	std.szl.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	std.szl.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	std.szl.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	std.szl.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	std.szl.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	std.szl.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	std.szl.DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	std.szl.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	std.szl.Fatalw(msg, keysAndValues...)
}

func Debugln(args ...interface{}) {
	std.szl.Debugln(args...)
}

func Infoln(args ...interface{}) {
	std.szl.Infoln(args...)
}

func Warnln(args ...interface{}) {
	std.szl.Warnln(args...)
}

func Errorln(args ...interface{}) {
	std.szl.Errorln(args...)
}

func DPanicln(args ...interface{}) {
	std.szl.DPanicln(args...)
}

func Panicln(args ...interface{}) {
	std.szl.Panicln(args...)
}

func Fatalln(args ...interface{}) {
	std.szl.Fatalln(args...)
}

// -------------------- logger functions mapping --------------------
func FDebug(msg string, fields ...zapcore.Field) {
	std.zl.Debug(msg, fields...)
}

func FInfo(msg string, fields ...zapcore.Field) {
	std.zl.Info(msg, fields...)
}

func FWarn(msg string, fields ...zapcore.Field) {
	std.zl.Warn(msg, fields...)
}

func FError(msg string, fields ...zapcore.Field) {
	std.zl.Error(msg, fields...)
}

func FDPanic(msg string, fields ...zapcore.Field) {
	std.zl.DPanic(msg, fields...)
}

func FPanic(msg string, fields ...zapcore.Field) {
	std.zl.Panic(msg, fields...)
}

func FFatal(msg string, fields ...zapcore.Field) {
	std.zl.Fatal(msg, fields...)
}
