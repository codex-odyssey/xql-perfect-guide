package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"app/internal/metrics"
	"app/internal/questions"
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
				otelgin.Middleware("xql-sample-app")(c)
			}
		},
		metrics.PrometheusMiddleware(),
	)

	r.GET("/", Handler)
	r.GET("/log", questions.HandlerLogQ)

	// === Recipe
	r.GET("/karubikuppa", recipe.Karubikuppa)
	r.GET("/curry", recipe.Curry)
	r.GET("/spaghetti", recipe.Spaghetti)
	r.GET("/meuniere", recipe.Meuniere)
	r.GET("/sandwich", recipe.Sandwich)
	r.GET("/salad", recipe.Salad)
	r.GET("/smoothie", recipe.Smoothie)

	r.GET("/metrics", metrics.HandlerMetrics)
	r.Run(":8080")
	log.Printf("Start Server")
}

func Handler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
