package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	logging "app/internal/log"
	"app/internal/metrics"
	"app/internal/questions"
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
	excludePaths := map[string]bool{
		"/metrics": true,
	}
	r.Use(
		func(c *gin.Context) {
			if !excludePaths[c.Request.URL.Path] {
				otelgin.Middleware("xql-sample-app")(c)
			}
		},
		metrics.PrometheusMiddleware(),
	)

	r.GET("/", Handler)
	r.GET("/log", questions.HandlerLogQ)
	r.GET("/karubikuppa", HandlerKarubikuppa)
	r.GET("/metrics", metrics.HandlerMetrics)
	r.Run(":8080")
	log.Printf("Start Server")
}

func Handler(c *gin.Context) {
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

	// temporary
	serverURL := "http://chef-service:8090/chef"
	params := url.Values{}
	params.Add("dish_name", "karubikuppa")
	req, err := http.NewRequestWithContext(ctx, "GET", serverURL+"?"+params.Encode(), nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ReadAll: ", err)
		return
	}

	c.String(http.StatusOK, string(body)+"秒で dekita!!")
}
