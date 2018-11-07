package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

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

func ObserveSystemLatency(sysURL string, serviceId string, observed time.Duration) {
	systemLatency.WithLabelValues(sysURL, serviceId).Observe(observed.Seconds())
}

func ObserveBackendLatency(backendURL string, serviceId string, observed time.Duration) {
	backendLatency.WithLabelValues(backendURL, serviceId).Observe(observed.Seconds())
}

func IncrementTotalRequests() {
	totalRequests.Inc()
}

func IncrementCacheHits() {
	cacheHits.Inc()
}
