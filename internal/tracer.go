package internal

import (
	"context"
	"go.opentelemetry.io/otel/sdk/trace"
	tr "go.opentelemetry.io/otel/trace"
	"time"
)

type TracerServiceImpl struct {
	TracerProvider *trace.TracerProvider
}

func (t *TracerServiceImpl) Tracer(tracerName string, ctx context.Context, spanName string, opt ...tr.SpanStartOption) (context.Context, tr.Span) {
	return t.TracerProvider.Tracer(tracerName).Start(ctx, spanName, opt...)
}

func (t *TracerServiceImpl) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return t.TracerProvider.Shutdown(ctx)
}
