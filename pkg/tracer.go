package pkg

import (
	"context"
	"github.com/thinc-org/newbie-gosdk/internal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	jg "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	tr "go.opentelemetry.io/otel/trace"
)

func initJaegerTracerProvider(host string, environment string, serviceName string) (*tracesdk.TracerProvider, error) {
	exp, err := jg.New(jg.WithCollectorEndpoint(jg.WithEndpoint(host + "/api/traces")))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			attribute.String("environment", environment),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func NewTracerService(host string, environment string, serviceName string) (Service, error) {
	tracerProvider, err := initJaegerTracerProvider(host, environment, serviceName)
	if err != nil {
		return nil, err
	}

	return &internal.TracerServiceImpl{
		TracerProvider: tracerProvider,
	}, nil
}

type Service interface {
	Tracer(tracerName string, ctx context.Context, spanName string, opt ...tr.SpanStartOption) (context.Context, tr.Span)
	Shutdown(ctx context.Context) error
}
