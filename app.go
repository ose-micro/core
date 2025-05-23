package ose

import (
	"context"

	"github.com/ose-micro/core/logger"
	"github.com/ose-micro/core/timestamp"
	"github.com/ose-micro/core/tracing"
	"go.uber.org/fx"
)

type App interface {
	Run()
}

type app struct {
	app *fx.App
}

// Run implements App.
func (a *app) Run() {
	a.app.Run()
}

func New(additional ...fx.Option) App {
	base := fx.Options(
		fx.Provide(
			logger.NewZap,
			tracing.NewOtel,
			timestamp.NewTimestamp,
		),
		fx.Invoke(registerLifecycleHooks),
	)
	return &app{
		app: fx.New(
			fx.Options(
				base,
				fx.Options(additional...),
			),
		),
	}
}

func registerLifecycleHooks(lc fx.Lifecycle, tracer tracing.Tracer, logger logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down tracer")
			if err := tracer.Shutdown(ctx); err != nil {
				logger.Error("Failed to shutdown tracer", "error", err)
				return err
			}

			logger.Info("Shutdown complete")
			return nil
		},
	})
}
