package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	CacheHits = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cache_hits",
		Help: "Number of cache hits",
	})
	CacheSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cache_size",
		Help: "Length of the cache",
	})
	SchedulerFailures = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "scheduler_failures",
		Help: "Number of scheduler failures",
	}, []string{"job_id"})
	SchedulerSuccesses = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "scheduler_successes",
		Help: "Number of scheduler successes",
	}, []string{"job_id"})
	Requests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "requests",
		Help: "Number of requests",
	}, []string{"type", "status"})
)

func InitMetrics() {
	prometheus.MustRegister(CacheHits, CacheSize, SchedulerFailures, SchedulerSuccesses, Requests)
}
