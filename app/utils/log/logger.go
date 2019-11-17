package log

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// init 默认初始化一份production的config
func init() {
	l, _ := zap.NewDevelopment()
	logger = l.Sugar()
}

// Init 根据config初始化logger
func New(cfg *Config) (*zap.SugaredLogger, error) {

	var l *zap.Logger
	var err error
	if l, err = cfg.BuildZapConfig().Build(); err != nil {
		return nil, err
	}
	logger = l.Sugar()
	return logger,nil
}

// Info
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Debug
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Sync flushes any buffered log entries.
func Sync() error {
	return logger.Sync()
}
