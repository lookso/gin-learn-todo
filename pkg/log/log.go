/*
@Time : 2020-02-01 11:13 
@Author : peanut
@File : log
@Software: GoLand
*/
package log

import (
	"go.uber.org/zap"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	cnf := zap.NewDevelopmentConfig()
	cnf.OutputPaths = []string{
		"stdout",
	}
	logger, _ = cnf.Build()
}

//func Init(cfg *Config) error {
//	var err error
//	if logger, err = cfg.Build(); err != nil {
//		return  err
//	}
//	return nil
//}
//
//func InitWithFile(path string) error {
//	cfg, err := NewConfigWithFile(path)
//	if err != nil {
//		return err
//	}
//	return Init(cfg)
//}
//
func Logger() *zap.Logger {
	return logger
}

func Sugar() *zap.SugaredLogger {
	if sugar == nil {
		sugar = logger.Sugar()
	}
	return sugar
}

func Sync() error{
	return logger.Sync()
}

