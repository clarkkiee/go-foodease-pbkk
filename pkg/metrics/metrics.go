package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP Requests",
		},
		[]string{"method", "endpoint"},
	)

	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_requests_duration_seconds",
			Help: "Histogram of response time for handler in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	HTTPResponseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_status_total",
			Help: "Total number of HTTP responses, classified by status code",
		},
		[]string{"status_code"},
	)

)

func Initialize() {
	prometheus.MustRegister(HTTPRequestsTotal)
	prometheus.MustRegister(HTTPRequestDuration)
	prometheus.MustRegister(HTTPResponseStatus)
}