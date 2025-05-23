package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "path to config directory")
}

// Option is a functional extension option
type Option func(v *viper.Viper, cfg *Config) error

func WithExtension[T any](key string, target *T) Option {
	return func(v *viper.Viper, cfg *Config) error {
		sub := v.Sub(key)
		if sub == nil {
			return nil // ignore missing sections
		}
		if err := sub.Unmarshal(target); err != nil {
			return fmt.Errorf("failed to unmarshal %s: %w", key, err)
		}
		if cfg.Extra == nil {
			cfg.Extra = map[string]any{}
		}
		cfg.Extra[key] = target
		return nil
	}
}

func Load(opts ...Option) (*Config, error) {
	_ = godotenv.Load()
	flag.Parse()

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	if configPath == "" {
		if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
			configPath = envPath
		} else {
			dir, err := callerDir()
			if err != nil {
				return nil, err
			}
			configPath = dir
		}
	}

	v := viper.New()
	v.SetConfigName(fmt.Sprintf("config.%s", env))
	v.SetConfigType("json")
	v.AddConfigPath(configPath)

	v.AutomaticEnv()
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var core CoreConfig
	if err := v.Unmarshal(&core); err != nil {
		return nil, fmt.Errorf("unmarshal core: %w", err)
	}

	full := &Config{Core: core}

	for _, opt := range opts {
		if err := opt(v, full); err != nil {
			return nil, err
		}
	}

	return full, nil
}

func callerDir() (string, error) {
	// Walk up the stack until we’re outside the core package
	for i := 1; i < 10; i++ {
		_, file, _, ok := runtime.Caller(i)
		if !ok {
			break
		}
		
		if !strings.Contains(file, "config") {
			return filepath.Dir(file), nil
		}
	}
	return "", fmt.Errorf("unable to detect caller directory")
}
