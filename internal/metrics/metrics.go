package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define metrics for queue length, retries, and processed requests
var (
	QueueLengthGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "metadata_processing_queue_length",
		Help: "Current length of the metadata processing queue.",
	})

	ProcessedRequestsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "metadata_processing_processed_requests_total",
			Help: "Total number of processed metadata requests.",
		},
		[]string{"status"}, // status can be "success" or "failure"
	)
)

// Initialize registers the metrics with Prometheus
func Initialize() {
	// Register metrics with Prometheus
	prometheus.MustRegister(QueueLengthGauge)
	prometheus.MustRegister(ProcessedRequestsCounter)
}

// ExposeMetrics exposes the metrics on an HTTP endpoint for Prometheus to scrape
func ExposeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Starting Prometheus metrics server on :2112")
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()
}

// IncrementProcessedRequests increments the success or failure count for processed requests
func IncrementProcessedRequests(status string) {
	ProcessedRequestsCounter.WithLabelValues(status).Inc()
}

func IncrementQueueLength() {
	QueueLengthGauge.Inc()
}

func DecrementQueueLength() {
	QueueLengthGauge.Dec()
}
