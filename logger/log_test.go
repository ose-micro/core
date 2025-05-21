// Package logger_test contains tests for the logger package.
//
// Copyright (c) 2025 Moriba. All rights reserved.

package logger_test

import (
	"testing"

	"github.com/ose-micro/core/logger"
)

func TestNewZapLogger(t *testing.T) {
	tests := []struct {
		name        string
		conf        logger.Config
		expectError bool
	}{
		{
			name: "Development debug level",
			conf: logger.Config{
				Environment: logger.ENV_DEVELOPMENT,
				Level:       "debug",
			},
			expectError: false,
		},
		{
			name: "Production info level",
			conf: logger.Config{
				Environment: logger.ENV_PRODUCTION,
				Level:       "info",
			},
			expectError: false,
		},
		{
			name: "Unknown level falls back to debug",
			conf: logger.Config{
				Environment: logger.ENV_DEVELOPMENT,
				Level:       "unknown",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logr, err := logger.NewZap(tt.conf)
			if (err != nil) != tt.expectError {
				t.Fatalf("NewZapLogger() error = %v, expectError %v", err, tt.expectError)
			}
			if logr == nil {
				t.Fatalf("NewZapLogger() returned nil logger")
			}

			// Basic smoke test for logging methods (should not panic)
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Logging method panicked: %v", r)
					}
				}()
				logr.Debug("test debug", "key1", "value1")
				logr.Info("test info", "key2", "value2")
				logr.Warn("test warn", "key3", "value3")
				logr.Error("test error", "key4", "value4")
			}()
		})
	}
}

func TestLoggingOutput(t *testing.T) {
	cfg := logger.Config{
		Environment: logger.ENV_DEVELOPMENT,
		Level:       "debug",
	}

	logr, err := logger.NewZap(cfg)
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}

	// Call logging methods and ensure no panic occurs
	logr.Debug("buffer test debug", "foo", "bar")
	logr.Info("buffer test info", "foo", "bar")

	// Note: Capturing and asserting log output requires exposing zap.Logger
	// or injecting a custom zapcore.Core, which is not implemented yet.
}
