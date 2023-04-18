package gosdk

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)
	config.EncoderConfig.TimeKey = "time"

	return config.Build()
}

func LogWithTraceID(ctx context.Context) logging.Fields {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return logging.Fields{"traceID", span.TraceID().String()}
	}
	return nil
}
