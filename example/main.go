package main

import (
	ose "github.com/ose-micro/core"
	"github.com/ose-micro/core/config"
	"github.com/ose-micro/core/logger"
	"github.com/ose-micro/core/timestamp"
	"github.com/ose-micro/core/tracing"
	"go.uber.org/fx"
)

func main() {
	app := ose.New(
		fx.Provide(
			loadConfig,
		),
		fx.Invoke(func(log logger.Logger, timestamp timestamp.Timestamp) error {
			log.Info("Load info")
			log.Info(timestamp.Format(timestamp.Now()))
			return nil
		}),
	)

	app.Run()

}

func loadConfig() (config.Service, logger.Config, tracing.Config, timestamp.Config, error) {
	conf, err := config.Load()
	if err != nil {
		return config.Service{}, logger.Config{}, tracing.Config{}, timestamp.Config{}, err
	}

	return conf.Core.Service, conf.Core.Service.Logger, conf.Core.Service.Tracer, conf.Core.Service.Timestamp, nil
}
