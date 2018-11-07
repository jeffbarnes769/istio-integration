package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// DefaultMetricsPort - Default port that metrics endpoint will be served on
const DefaultMetricsPort = 8080

var systemLatency = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "threescale_system_latency",
		Help:    "Request latency for requests to 3scale system URL",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"systemURL", "serviceID"},
)

var backendLatency = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "threescale_backend_latency",
		Help:    "Request latency for requests to 3scale backend",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"backendURL", "serviceID"},
)

var totalRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "handle_authorization_requests",
		Help: "Total number of requests to adapter",
	},
)

var cacheHits = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "system_cache_hits",
		Help: "Total number of requests to 3scale system fetched from cache",
	},
)

func init() {
	prometheus.MustRegister(systemLatency, backendLatency)
}

// ObserveSystemLatency reports a metric to system latency histogram
func ObserveSystemLatency(sysURL string, serviceID string, observed time.Duration) {
	systemLatency.WithLabelValues(sysURL, serviceID).Observe(observed.Seconds())
}

// ObserveBackendLatency reports a metric to backend latency histogram
func ObserveBackendLatency(backendURL string, serviceID string, observed time.Duration) {
	backendLatency.WithLabelValues(backendURL, serviceID).Observe(observed.Seconds())
}

// IncrementTotalRequests increments the request count for authorization handler
func IncrementTotalRequests() {
	totalRequests.Inc()
}

// IncrementCacheHits increments proxy configurations that have bee read from the cach
func IncrementCacheHits() {
	cacheHits.Inc()
}
