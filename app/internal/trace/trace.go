package tracing

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	logging "app/internal/log"
)

const (
	TraceBackendAddressEnv = "TRACE_BACKEND_ADDRESS"
)

func InitTracer() (*sdktrace.TracerProvider, error) {
	ctx := context.Background()

	otelCollectorAddress := os.Getenv(TraceBackendAddressEnv)

	fmt.Print("Tempo Connection... ", otelCollectorAddress)
	conn, err := grpc.DialContext(ctx,
		otelCollectorAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}
	fmt.Printf("Tempo Connected.")

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// LogQLのcontainerタグにひも付けるため、コンテナ名と同じにする必要がある
			semconv.ServiceNameKey.String("waiter-service"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return tp, nil
}

var tracer = otel.GetTracerProvider().Tracer("")

func CreateTrace(ctx context.Context, sleep int, step string, logger *zap.SugaredLogger) {
	ctx, span := tracer.Start(ctx, step)
	logger = logging.WithTrace(ctx, logger)
	logger.Infoln(step)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	defer span.End()
}

func CreateTraceWithAttribute(ctx context.Context, sleep int, step string, logger *zap.SugaredLogger, attrs map[string]interface{}) {
	ctx, span := tracer.Start(ctx, step)
	for key, value := range attrs {
		span.SetAttributes(attribute.String(key, value.(string)))
	}
	logger = logging.WithTrace(ctx, logger)
	logger.Info(step)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	defer span.End()
}
