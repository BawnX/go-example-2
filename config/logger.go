package config

import (
	"github.com/hashicorp/go-hclog"
	"log"
)

func NewLoggerStandard(name string, isDebug bool) *log.Logger {
	logger := hclog.New(
		&hclog.LoggerOptions{
			Name:       name,
			Level:      hclog.LevelFromString(*LogLevel),
			JSONFormat: isDebug,
		},
	)

	return logger.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})
}

func NewLogger(name string, isDebug bool) hclog.Logger {
	return hclog.New(
		&hclog.LoggerOptions{
			Name:       name,
			Level:      hclog.LevelFromString(*LogLevel),
			JSONFormat: isDebug,
		},
	)
}
