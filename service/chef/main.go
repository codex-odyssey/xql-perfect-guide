package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"app/internal/metrics"
	recipe "app/internal/recipe"
	tracing "app/internal/trace"
)

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
				otelgin.Middleware("chef-service")(c)
			}
		},
		metrics.PrometheusMiddleware(),
	)

	r.GET("/", Handler)

	// === Recipe
	r.GET("/karubikuppa", recipe.Karubikuppa)

	r.GET("/metrics", metrics.HandlerMetrics)
	r.Run(":8090")
	log.Printf("Start Server")
}

func Handler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
