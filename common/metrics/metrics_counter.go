package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var TotalRequests = *promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of requests.",
	},
	[]string{"code", "endpoint", "method"},
)

var Duration = *promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "A histogram of latencies for requests.",
		Buckets: []float64{.25, .5, 1, 2.5, 5, 10},
	},
	[]string{"endpoint", "method"},
)
