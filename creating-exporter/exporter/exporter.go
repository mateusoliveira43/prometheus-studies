package main

import (
	"log"
	"net/http"
	"github.com/pbnjay/memory"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getFreeMemory() float64 {
    // `:=` operator is a short assignment statement
    // `a := b` it is the same as `var a = b`
    // cannot be used for constants
	freeMemory := memory.FreeMemory()
	return float64(freeMemory)
}

func getTotalMemory() float64 {
	totalMemory := memory.TotalMemory()
	return float64(totalMemory)
}

var (
    // https://prometheus.io/docs/practices/naming/
	freeMemoryBytesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "exporter_free_memory_bytes",
		Help: "Amount of free memory in bytes",
	})

	totalMemoryBytesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "exporter_memory_bytes_total",
		Help: "Total amount of memory in bytes",
	})
)

func init() {
	prometheus.MustRegister(freeMemoryBytesGauge)
	prometheus.MustRegister(totalMemoryBytesGauge)
}

func main() {
	freeMemoryBytesGauge.Set(getFreeMemory())
	totalMemoryBytesGauge.Set(getTotalMemory())

	http.Handle("/metrics", promhttp.Handler())
    // TODO pass PORT through environment variable
	log.Fatal(http.ListenAndServe(":8081", nil))
}
