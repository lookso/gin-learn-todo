/*
@Time : 2019-11-24 16:10 
@Author : Tenlu
@File : config
@Software: GoLand
*/
package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level          string   `json:"level" toml:"level"`
	Encoding       string   `json:"encoding" toml:"encoding"`
	OutputPaths    []string `json:"output_paths" toml:"output_paths"`
	ErrOutputPaths string   `json:"err_output_paths" toml:"err_output_paths"`
	InitialFields  string   `json:"initial_fields" toml:"initial_fields"`
}

func (cfg *Config) BuildZapConfig() zap.Config {
	cnf := zap.NewProductionConfig()
	cnf.Development = true
	cnf.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	if cfg.Encoding != "" {
		cnf.Encoding = cfg.Encoding
	}

	if len(cfg.OutputPaths) > 0 {
		cnf.OutputPaths = cfg.OutputPaths
	}

	return cnf
}

