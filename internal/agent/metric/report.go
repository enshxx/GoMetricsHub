package metric

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
)

func UpdateMetrics(pollCount *int64, randomValue *float64) {
	*pollCount++
	*randomValue = rand.Float64()
}

func ReportMetrics(serverAddress string, pollCount *int64, randomValue *float64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	reportGauge(serverAddress, "Alloc", float64(m.Alloc))
	reportGauge(serverAddress, "BuckHashSys", float64(m.BuckHashSys))
	reportGauge(serverAddress, "Frees", float64(m.Frees))
	reportGauge(serverAddress, "GCCPUFraction", m.GCCPUFraction)
	reportGauge(serverAddress, "GCSys", float64(m.GCSys))
	reportGauge(serverAddress, "HeapAlloc", float64(m.HeapAlloc))
	reportGauge(serverAddress, "HeapIdle", float64(m.HeapIdle))
	reportGauge(serverAddress, "HeapInuse", float64(m.HeapInuse))
	reportGauge(serverAddress, "HeapObjects", float64(m.HeapObjects))
	reportGauge(serverAddress, "HeapReleased", float64(m.HeapReleased))
	reportGauge(serverAddress, "HeapSys", float64(m.HeapSys))
	reportGauge(serverAddress, "LastGC", float64(m.LastGC))
	reportGauge(serverAddress, "Lookups", float64(m.Lookups))
	reportGauge(serverAddress, "MCacheInuse", float64(m.MCacheInuse))
	reportGauge(serverAddress, "MCacheSys", float64(m.MCacheSys))
	reportGauge(serverAddress, "MSpanInuse", float64(m.MSpanInuse))
	reportGauge(serverAddress, "MSpanSys", float64(m.MSpanSys))
	reportGauge(serverAddress, "Mallocs", float64(m.Mallocs))
	reportGauge(serverAddress, "NextGC", float64(m.NextGC))
	reportGauge(serverAddress, "NumForcedGC", float64(m.NumForcedGC))
	reportGauge(serverAddress, "NumGC", float64(m.NumGC))
	reportGauge(serverAddress, "OtherSys", float64(m.OtherSys))
	reportGauge(serverAddress, "PauseTotalNs", float64(m.PauseTotalNs))
	reportGauge(serverAddress, "StackInuse", float64(m.StackInuse))
	reportGauge(serverAddress, "StackSys", float64(m.StackSys))
	reportGauge(serverAddress, "Sys", float64(m.Sys))
	reportGauge(serverAddress, "TotalAlloc", float64(m.TotalAlloc))

	reportGauge(serverAddress, "RandomValue", *randomValue)
	reportCounter(serverAddress, "PollCount", *pollCount)
}

func reportMetric(serverAddress string, metricType, name string, value interface{}) {
	url := fmt.Sprintf("%s/update/%s/%s/%v", serverAddress, metricType, name, value)
	resp, err := http.Post(url, "text/plain", nil)

	if err != nil {
		fmt.Println("failed to sent POST request with metric: %w", err)
		return
	}
	err = resp.Body.Close()
	if err != nil {
		fmt.Println("failed to close response body: %w", err)
		return
	}

}

func reportGauge(serverAddress string, name string, value float64) {
	reportMetric(serverAddress, "gauge", name, value)
}

func reportCounter(serverAddress string, name string, value int64) {
	reportMetric(serverAddress, "counter", name, value)
}
