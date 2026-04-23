package visibility

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// ================= METRICS =================

// Request count (Counter)
var HttpRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "route", "status"},
)

// Request latency (Histogram)
var HttpDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "HTTP request latency",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"method", "route"},
)

// In-flight requests (Gauge)
var InFlight = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "http_in_flight_requests",
		Help: "Current number of in-flight requests",
	},
)

func init() {
	prometheus.MustRegister(HttpRequests, HttpDuration, InFlight)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		InFlight.Inc()
		defer InFlight.Dec()
		c.Next()
		duration := time.Since(start).Seconds()
		route := c.FullPath()

		if route == "" {
			route = "unknown"
		}

		status := c.Writer.Status()

		HttpRequests.WithLabelValues(
			c.Request.Method,
			route,
			strconv.Itoa(status),
		).Inc()

		HttpDuration.WithLabelValues(
			c.Request.Method,
			route).Observe(duration)
	}

}
