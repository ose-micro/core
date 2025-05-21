package main

import (
	ose "github.com/ose-micro/core"
	"github.com/ose-micro/core/config"
	"github.com/ose-micro/core/logger"
	"github.com/ose-micro/core/tracing"
	"go.uber.org/fx"
)

type Redis struct {
	Address string `mapstructure:"address"`
}

func main() {
	app := ose.New(
		fx.Provide(
			loadConfig,
		),
		fx.Invoke(func (log logger.Logger)  error {
			log.Info("Load info")
			return nil
		}),
	)

	app.Run()

}

func loadConfig() (config.Service, logger.Config, tracing.Config, error) {
	conf, err := config.Load()
	if err != nil {
		return config.Service{}, logger.Config{}, tracing.Config{}, err
	}

	return conf.Core.Service, conf.Core.Service.Logger, conf.Core.Service.Tracer, nil
}
