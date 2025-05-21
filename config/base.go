package config

import (
	"github.com/ose-micro/core/logger"
	"github.com/ose-micro/core/tracing"
)

type Service struct {
	Name   string         `mapstructure:"name"`
	Logger logger.Config  `mapstructure:"logger"`
	Tracer tracing.Config `mapstructure:"tracer"`
}

type CoreConfig struct {
	Service Service `mapstructure:"service"`
}

type Config struct {
	Core  CoreConfig
	Extra map[string]any
}
