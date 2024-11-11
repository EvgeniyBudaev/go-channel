package config

import (
	"github.com/EvgeniyBudaev/go-channel/internal/logger"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Port        string `envconfig:"PORT"`
	LoggerLevel string `envconfig:"LOGGER_LEVEL"`
}

func Load(l logger.Logger) (*Config, error) {
	var cfg Config
	err := envconfig.Process("APP", &cfg)
	if err != nil {
		l.Debug("error func Load, method Process by path internal/config/config.go", zap.Error(err))
		return nil, err
	}
	return &cfg, nil
}
