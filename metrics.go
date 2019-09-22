package blockatlas

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	namespace = "client"
)

var (
	labels = []string{"status", "endpoint", "method"}

	reqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made.",
		}, labels,
	)
	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latencies in seconds.",
		}, labels,
	)
)

// init registers the prometheus metrics
func init() {
	prometheus.MustRegister(reqCount, reqDuration)
}

func getMetrics(status, url, method string, start time.Time) {
	lvs := []string{status, url, method}
	reqCount.WithLabelValues(lvs...).Inc()
	reqDuration.WithLabelValues(lvs...).Observe(time.Since(start).Seconds())
}
