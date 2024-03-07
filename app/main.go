package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	logging "app/internal/log"
	"app/internal/metrics"
	tracing "app/internal/trace"
)

var tracer = otel.GetTracerProvider().Tracer("")

func main() {
	// Start Tracing
	tp, err := tracing.InitTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	r := gin.New()
	r.Use(
		otelgin.Middleware("xql-sample-app"),
		metrics.PrometheusMiddleware(),
	)

	r.GET("/", Handler)
	r.GET("/log", HandlerLog)
	r.GET("/karubikuppa", HandlerKarubikuppa)
	r.GET("/metrics", metrics.PrometheusHandler())
	r.Run(":8080")
	log.Printf("Start Server")
}

func Handler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func HandlerLog(c *gin.Context) {
	ctx := c.Request.Context()
	logger := logging.GetLoggerFromCtx(ctx)

	logger.Infoln("foo")
	logger.Infoln("var")
	logger.Infoln("foo")
	logger.Infoln("var")
	logger.Infoln("var")
	logger.Infoln("foo")
	c.String(http.StatusOK, "ok")
}

func HandlerKarubikuppa(c *gin.Context) {
	ctx := c.Request.Context()
	logger := logging.GetLoggerWithTraceID(ctx)

	func() {
		ctx, span := tracer.Start(ctx, "食材を準備する", trace.WithAttributes(
			attribute.String("カルビ", "たくさん"),
			attribute.String("コチュジャン", "いっぱい"),
			attribute.String("ごま油", "できるだけ多く"),
		))
		logger := logging.WithTrace(ctx, logger)
		logger.Infoln("食材を準備する")
		defer span.End()
	}()

	func() {
		ctx, span := tracer.Start(ctx, "カルビを炒める")
		logger := logging.WithTrace(ctx, logger)
		logger.Infoln("カルビを炒める")
		time.Sleep(1 * time.Second)
		defer span.End()
	}()

	func() {
		ctx, span := tracer.Start(ctx, "カルビクッパを煮込む")
		logger := logging.WithTrace(ctx, logger)
		logger.Infoln("カルビクッパを煮込む")
		time.Sleep(3 * time.Second)
		defer span.End()
	}()

	func() {
		ctx, span := tracer.Start(ctx, "ごま油を入れる")
		logger := logging.WithTrace(ctx, logger)
		logger.Infoln("ごま油を入れる")
		defer span.End()
	}()

	c.String(http.StatusOK, "dekita!!")
}
