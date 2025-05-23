package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zapLogger wraps a zap.Logger instance to implement
// the Logger interface with structured logging methods.
type zapLogger struct {
	logger *zap.Logger
}

// Zap implements Logger.
func (z *zapLogger) Zap() *zap.Logger {
	return z.logger
}

// Fatal implements Logger.
func (z *zapLogger) Fatal(msg string, keysAndValues ...any) {
	z.logger.Sugar().Fatalw(msg, keysAndValues...)
}

// Panic implements Logger.
func (z *zapLogger) Panic(msg string, keysAndValues ...any) {
	z.logger.Sugar().Panicw(msg, keysAndValues...)
}

// Debug logs a message at Debug level with optional key-value pairs.
// Implements the Debug method of the Logger interface.
func (z *zapLogger) Debug(msg string, keysAndValues ...any) {
	z.logger.Sugar().Debugw(msg, keysAndValues...)
}

// Error logs a message at Error level with optional key-value pairs.
func (z *zapLogger) Error(msg string, keysAndValues ...any) {
	z.logger.Sugar().Errorw(msg, keysAndValues...)
}

// Info logs a message at Info level with optional key-value pairs.
func (z *zapLogger) Info(msg string, keysAndValues ...any) {
	z.logger.Sugar().Infow(msg, keysAndValues...)
}

// Warn logs a message at Warn level with optional key-value pairs.
func (z *zapLogger) Warn(msg string, keysAndValues ...any) {
	z.logger.Sugar().Warnw(msg, keysAndValues...)
}

// loggerLevelMap maps string log levels to zapcore.Level values.
var loggerLevelMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"error": zap.ErrorLevel,
	"panic": zap.PanicLevel,
	"fatal": zap.FatalLevel,
}

const (
	// ENV_PRODUCTION indicates the production environment.
	ENV_PRODUCTION = "Production"
	// ENV_DEVELOPMENT indicates the development environment.
	ENV_DEVELOPMENT = "Development"
)

// NewZap creates and returns a new zapLogger instance configured
// according to the given Config.
// The logger level and environment (production/development) influence
// the configuration and log output format.
//
// Returns an error if the underlying zap logger cannot be built.
func NewZap(conf Config) (Logger, error) {
	// Default to DebugLevel if level is not recognized
	level := zap.DebugLevel

	if lvl, exists := loggerLevelMap[conf.Level]; exists {
		level = lvl
	}

	// Use development config by default
	zc := zap.NewDevelopmentConfig()

	// Use production config if environment is set to Production
	if conf.Environment == ENV_PRODUCTION {
		zc = zap.NewProductionConfig()
	}

	// Set the logging level
	zc.Level = zap.NewAtomicLevelAt(level)

	// Build the zap logger with caller info and caller skip for cleaner stack traces
	logger, err := zc.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	return &zapLogger{
		logger: logger,
	}, nil
}
