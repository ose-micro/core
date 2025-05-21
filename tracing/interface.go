package tracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)


type Tracer interface {
	Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span)
	Shutdown(ctx context.Context) error
}
