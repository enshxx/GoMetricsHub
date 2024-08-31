package main

import (
	"github.com/enshxx/GoMetricsHub/internal/agent/metric"
	"time"
)

const (
	serverAddress  = "http://localhost:8080"
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func main() {
	var pollCount int64 = 0
	randomValue := 0.0
	go func() {
		for {
			metric.UpdateMetrics(&pollCount, &randomValue)
			time.Sleep(pollInterval)
		}
	}()

	for {
		metric.ReportMetrics(serverAddress, &pollCount, &randomValue)
		time.Sleep(reportInterval)
	}
}
