package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level          string   `json:"level" yaml:"level"`
	Encoding       string   `json:"encoding" yaml:"encoding"`
	OutputPaths    []string `json:"output_paths" yaml:"output_paths"`
	ErrOutputPaths string   `json:"err_output_paths" yaml:"err_output_paths"`
	InitialFields  string   `json:"initial_fields" yaml:"initial_fields"`
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
