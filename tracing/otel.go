// Package tracing provides OpenTelemetry tracing setup using OTLP exporters.
//
// Copyright (c) 2025 Moriba. All rights reserved.

package tracing

import (
	"context"
	"fmt"

	"github.com/ose-micro/core/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

// otelTracing implements the Tracer interface using OpenTelemetry.
type otelTracing struct {
	tp *sdktrace.TracerProvider
}

// Start begins a new span with the provided name and options using the default tracer.
func (o otelTracing) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return otel.Tracer("ose").Start(ctx, spanName, opts...)
}

// Shutdown flushes and shuts down the tracer provider.
func (o otelTracing) Shutdown(ctx context.Context) error {
	return o.tp.Shutdown(ctx)
}

// NewOtel sets up and returns an OpenTelemetry tracer with OTLP exporter using gRPC.
func NewOtel(conf Config, log logger.Logger) (Tracer, error) {
	ctx := context.Background()

	// Create OTLP trace exporter
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(), // For insecure connection; use WithTLS for production
		otlptracegrpc.WithEndpoint(conf.Endpoint),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)
	if err != nil {
		log.Error("failed to create exporter: %w", err)
		return nil, fmt.Errorf("failed to create exporter: %w", err)
	}

	// Create tracer provider with batch span processor and resource attributes
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(conf.ServiceName),
		)),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(conf.SampleRatio)),
	)

	// Register the global tracer provider
	otel.SetTracerProvider(tp)

	return &otelTracing{tp: tp}, nil
}
