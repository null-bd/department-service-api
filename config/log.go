package config

import (
	"github.com/null-bd/logger"
)

func GetLogger(cfg *Config) (logger.Logger, error) {
	defaultLogger, err := logger.New(&logger.Config{
		ServiceName: cfg.App.Name,
		Environment: cfg.App.Env,
		LogLevel:    logger.InfoLevel,
		Format:      "json",
		OutputPaths: []string{"stdout"},
	})
	if err != nil {
		return nil, err
	}

	return defaultLogger, nil
}
