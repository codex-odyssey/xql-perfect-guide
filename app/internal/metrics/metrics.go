package metrics

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"

	"app/internal/constant"
)

const (
	exponentialBucketsStart  = 128
	exponentialBucketsFactor = 2
	exponentialBucketsCount  = 10
)

// メトリクス定義
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests processed",
		},
		[]string{"method", "path"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	httpRequestSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_size_bytes",
			Help: "Size of HTTP requests",
			Buckets: prometheus.ExponentialBuckets(
				exponentialBucketsStart, exponentialBucketsFactor, exponentialBucketsCount,
			),
		},
		[]string{"method", "path"},
	)

	httpResponseSize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_response_size_bytes",
			Help: "Size of HTTP responses",
			Buckets: prometheus.ExponentialBuckets(
				exponentialBucketsStart, exponentialBucketsFactor, exponentialBucketsCount,
			),
		},
		[]string{"method", "path"},
	)

	httpResponseTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_response_time_seconds",
			Help: "Time of the last HTTP response",
		},
		[]string{"method", "path"},
	)
)

func init() {
	registerMetrics()
}

func registerMetrics() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(httpRequestSize)
	prometheus.MustRegister(httpResponseSize)
	prometheus.MustRegister(httpResponseTime)
}

func updateMetrics(ctx context.Context, method, path string, requestSize, responseSize int, duration time.Duration) {
	traceID := trace.SpanContextFromContext(ctx).TraceID()

	rt := httpRequestsTotal.WithLabelValues(method, path)
	if exemplarAdder, ok := rt.(prometheus.ExemplarAdder); ok && traceID.IsValid() {
		exemplarAdder.AddWithExemplar(1, prometheus.Labels{constant.TrackID: traceID.String()})
	} else {
		rt.Inc()
	}

	rd := httpRequestDuration.WithLabelValues(method, path)
	if exemplarObserver, ok := rd.(prometheus.ExemplarObserver); ok && traceID.IsValid() {
		exemplarObserver.ObserveWithExemplar(duration.Seconds(), prometheus.Labels{constant.TrackID: traceID.String()})
	} else {
		rd.Observe(duration.Seconds())
	}

	httpRequestSize.WithLabelValues(method, path).Observe(float64(requestSize))
	httpResponseSize.WithLabelValues(method, path).Observe(float64(responseSize))
	httpResponseTime.WithLabelValues(method, path).Set(float64(time.Now().Unix()))

}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)

		requestSize := c.Request.ContentLength
		responseSize := c.Writer.Size()

		updateMetrics(c.Request.Context(), c.Request.Method, c.Request.URL.Path, int(requestSize), responseSize, duration)
	}
}

func PrometheusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := promhttp.InstrumentMetricHandler(
			prometheus.DefaultRegisterer,
			promhttp.HandlerFor(
				prometheus.DefaultGatherer,
				promhttp.HandlerOpts{EnableOpenMetrics: true},
			),
		)
		h.ServeHTTP(c.Writer, c.Request)
	}
}
