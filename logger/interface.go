package logger

// Config holds configuration options for creating a new zapLogger.
type Config struct {
	// Environment specifies the deployment environment: e.g. "Production" or "Development".
	Environment string `mapstructure:"environment"`
	// Level is the logging level as a string: "debug", "info", "warn", "error", "panic", or "fatal".
	Level string `mapstructure:"level"`
}

// Logger defines the logging methods supported by the logger package.
type Logger interface {
	Info(msg string, keysAndValues ...any)
	Warn(msg string, keysAndValues ...any)
	Error(msg string, keysAndValues ...any)
	Debug(msg string, keysAndValues ...any)
	Fatal(msg string, keysAndValues ...any)
	Panic(msg string, keysAndValues ...any)
}
