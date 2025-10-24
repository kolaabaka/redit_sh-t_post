package monitoring

import (
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var totalHttpReqCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "total_http_request_count",
		Help: "Every http request increment this counter",
	},
)

var httpReqCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "",
	},
	[]string{"url"},
)

var memoryUsage = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "go_memory_usage_bytes",
		Help: "Current memory usage statistics",
	},
	[]string{"type"},
)

var memoryStats = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "go_memory_stats_bytes",
		Help: "Detailed memory statistics from runtime",
	},
	[]string{"stat"},
)

var gcStats = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "go_gc_stats",
		Help: "Garbage collector statistics",
	},
	[]string{"stat"},
)

func MustInitPrometheusStat() {
	prometheus.MustRegister(totalHttpReqCounter)
	prometheus.MustRegister(httpReqCounter)
	//Memory statistic
	prometheus.MustRegister(memoryUsage)
	prometheus.MustRegister(memoryStats)
	prometheus.MustRegister(gcStats)
	go trackingMemmoryStatGauge()
}

func IncrementEndpointHttpCounter(urlPath string) {
	httpReqCounter.WithLabelValues(urlPath).Inc()
}

func IncrementTotalhttpCounter() {
	totalHttpReqCounter.Inc()
}

func trackingMemmoryStatGauge() {
	var memStats = runtime.MemStats{}
	var ticker = time.NewTicker(time.Second)
	for {
		<-ticker.C

		time.Sleep(time.Second * 1)
		runtime.ReadMemStats(&memStats)

		updateMemoryMetrics(memStats)
	}
}

func updateMemoryMetrics(memStats runtime.MemStats) {
	memoryUsage.WithLabelValues("heap").Set(float64(memStats.HeapAlloc))
	memoryUsage.WithLabelValues("stack").Set(float64(memStats.StackInuse))
	memoryUsage.WithLabelValues("system").Set(float64(memStats.Sys))
	memoryUsage.WithLabelValues("alloc").Set(float64(memStats.Alloc))

	memoryStats.WithLabelValues("heap_sys").Set(float64(memStats.HeapSys))
	memoryStats.WithLabelValues("heap_idle").Set(float64(memStats.HeapIdle))
	memoryStats.WithLabelValues("heap_inuse").Set(float64(memStats.HeapInuse))
	memoryStats.WithLabelValues("heap_released").Set(float64(memStats.HeapReleased))
	memoryStats.WithLabelValues("heap_objects").Set(float64(memStats.HeapObjects))

	gcStats.WithLabelValues("num_gc").Set(float64(memStats.NumGC))
	gcStats.WithLabelValues("pause_total_ns").Set(float64(memStats.PauseTotalNs))
	gcStats.WithLabelValues("next_gc_bytes").Set(float64(memStats.NextGC))
}
