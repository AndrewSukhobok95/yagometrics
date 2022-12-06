package poller

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
)

func UpdateMetrics(storage storage.Storage, rtm *runtime.MemStats) {
	runtime.ReadMemStats(rtm)
	storage.AddCounterMetric("PollCount", 1)
	storage.InsertGaugeMetric("RandomValue", rand.Float64())
	storage.InsertGaugeMetric("Alloc", float64(rtm.Alloc))
	storage.InsertGaugeMetric("BuckHashSys", float64(rtm.BuckHashSys))
	storage.InsertGaugeMetric("Frees", float64(rtm.Frees))
	storage.InsertGaugeMetric("GCCPUFraction", float64(rtm.GCCPUFraction))
	storage.InsertGaugeMetric("GCSys", float64(rtm.GCSys))
	storage.InsertGaugeMetric("HeapAlloc", float64(rtm.HeapAlloc))
	storage.InsertGaugeMetric("HeapIdle", float64(rtm.HeapIdle))
	storage.InsertGaugeMetric("HeapInuse", float64(rtm.HeapInuse))
	storage.InsertGaugeMetric("HeapObjects", float64(rtm.HeapObjects))
	storage.InsertGaugeMetric("HeapReleased", float64(rtm.HeapReleased))
	storage.InsertGaugeMetric("HeapSys", float64(rtm.HeapSys))
	storage.InsertGaugeMetric("LastGC", float64(rtm.LastGC))
	storage.InsertGaugeMetric("Lookups", float64(rtm.Lookups))
	storage.InsertGaugeMetric("MCacheInuse", float64(rtm.MCacheInuse))
	storage.InsertGaugeMetric("MCacheSys", float64(rtm.MCacheSys))
	storage.InsertGaugeMetric("MSpanSys", float64(rtm.MSpanSys))
	storage.InsertGaugeMetric("Mallocs", float64(rtm.Mallocs))
	storage.InsertGaugeMetric("NextGC", float64(rtm.NextGC))
	storage.InsertGaugeMetric("NumForcedGC", float64(rtm.NumForcedGC))
	storage.InsertGaugeMetric("NextGC", float64(rtm.NextGC))
	storage.InsertGaugeMetric("OtherSys", float64(rtm.OtherSys))
	storage.InsertGaugeMetric("PauseTotalNs", float64(rtm.PauseTotalNs))
	storage.InsertGaugeMetric("StackInuse", float64(rtm.StackInuse))
	storage.InsertGaugeMetric("StackSys", float64(rtm.StackSys))
	storage.InsertGaugeMetric("Sys", float64(rtm.Sys))
	storage.InsertGaugeMetric("TotalAlloc", float64(rtm.TotalAlloc))
}

func Poll(storage storage.Storage, pollInterval time.Duration) {
	ticker := time.NewTicker(pollInterval)
	var rtm runtime.MemStats
	storage.InsertCounterMetric("PollCount", 0)
	for {
		<-ticker.C
		UpdateMetrics(storage, &rtm)
	}
}
