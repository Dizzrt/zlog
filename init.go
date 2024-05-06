package zlog

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

type ZlogConfig struct {
	LogPath     string
	LogFileName string
	LogLevel    Level
	MaxCount    int
	MaxAge      time.Duration
	RotateTime  time.Duration
}

var once sync.Once

func InitZlog(cfg ZlogConfig) {
	once.Do(func() {
		initZlog(cfg)
	})
}

func initZlog(cfg ZlogConfig) {
	initLogPath(cfg.LogPath)

	rotateOptions := getRotateOptions(cfg)
	rotateLogName := fmt.Sprintf("%s_%%Y%%m%%d%%H", cfg.LogFileName)
	fmt.Println(rotateLogName)

	rotater, err := rotatelogs.New(
		filepath.Join(cfg.LogPath, rotateLogName),
		rotateOptions...,
	)

	if err != nil {
		panic(err)
	}

	writeSyncer := &zapcore.BufferedWriteSyncer{
		WS:   zapcore.AddSync(rotater),
		Size: 4096,
	}

	std = NewWithBufferedWriteSyncer(writeSyncer, cfg.LogLevel)
}

func getRotateOptions(cfg ZlogConfig) []rotatelogs.Option {
	options := make([]rotatelogs.Option, 0)
	options = append(options, rotatelogs.WithLinkName(cfg.LogPath+cfg.LogFileName))

	if cfg.RotateTime > 0 {
		options = append(options, rotatelogs.WithRotationTime(cfg.RotateTime))
	}

	if cfg.MaxAge > 0 {
		options = append(options, rotatelogs.WithMaxAge(cfg.MaxAge))
	}

	if cfg.MaxCount > 0 {
		options = append(options, rotatelogs.WithRotationCount(cfg.MaxCount))
	}

	return options
}

func initLogPath(path string) {
	// check if log path is exists
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// creat log path
			if err = os.Mkdir(path, 0744); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
}
