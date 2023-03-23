package gosdk

import (
	"context"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/tracer"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	tr "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

var tracerService tracer.Service

// NewGRPUnaryClientInterceptor returns unary client interceptor. It is used
// with `grpc.WithUnaryInterceptor` method.
func NewGRPUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return otelgrpc.UnaryClientInterceptor()
}

// NewGRPUnaryServerInterceptor returns unary server interceptor. It is used
// with `grpc.UnaryInterceptor` method.
func NewGRPUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return otelgrpc.UnaryServerInterceptor()
}

// NewGRPCStreamClientInterceptor returns stream client interceptor. It is used
// with `grpc.WithStreamInterceptor` method.
func NewGRPCStreamClientInterceptor() grpc.StreamClientInterceptor {
	return otelgrpc.StreamClientInterceptor()
}

// NewGRPCStreamServerInterceptor returns stream server interceptor. It is used
// with `grpc.StreamInterceptor` method.
func NewGRPCStreamServerInterceptor() grpc.StreamServerInterceptor {
	return otelgrpc.StreamServerInterceptor()
}

type JaegerConfig struct {
	Host        string `mapstructure:"host"`
	Environment string `mapstructure:"env"`
	ServiceName string `mapstructure:"service-name"`
}

// SetUpTracer set up the jaeger url and resource setting
func SetUpTracer(conf *JaegerConfig) error {
	logger, _ := NewLogger()
	service, err := tracer.NewService(conf.Host, conf.Environment, conf.ServiceName)
	if err != nil {
		return err
	}

	tracerService = service

	logger.Info(
		"successfully setup jaeger tracing",
		zap.String("service", "jaeger"),
		zap.String("host", conf.Host),
		zap.String("environment", conf.Environment),
		zap.String("service_name", conf.ServiceName),
	)

	return nil
}

// StartTracer create new trace span
func StartTracer(tracerName string, ctx context.Context, name string, opt ...tr.SpanStartOption) (context.Context, tr.Span) {
	if tracerService == nil {
		return nil, nil
	}

	return tracerService.Tracer(tracerName, ctx, name, opt...)
}

// CloseTracer close connection from tracer provider
func CloseTracer() error {
	if tracerService == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return tracerService.Shutdown(ctx)
}
