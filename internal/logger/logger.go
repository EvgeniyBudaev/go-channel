package logger

import (
	"go.uber.org/zap"
)

type Logger = *zap.Logger

func New(level string) (Logger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = lvl
	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return zl, nil
}

func GetDefaultLevel() string {
	return "DEBUG"
}
