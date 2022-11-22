package metrics

import (
	"math/rand"
	"runtime"
	"sync"
)

type Metrics struct {
	Counters map[string]int64
	Gauges   map[string]float64
	Mutex    sync.Mutex
	rtm      runtime.MemStats
}

func (m *Metrics) Update() {
	runtime.ReadMemStats(&m.rtm)
	m.Mutex.Lock()
	m.Counters["PollCount"] += 1
	m.Gauges["RandomValue"] = rand.Float64()
	m.Gauges["Alloc"] = float64(m.rtm.Alloc)
	m.Gauges["BuckHashSys"] = float64(m.rtm.BuckHashSys)
	m.Gauges["Frees"] = float64(m.rtm.Frees)
	m.Gauges["GCCPUFraction"] = float64(m.rtm.GCCPUFraction)
	m.Gauges["GCSys"] = float64(m.rtm.GCSys)
	m.Gauges["HeapAlloc"] = float64(m.rtm.HeapAlloc)
	m.Gauges["HeapIdle"] = float64(m.rtm.HeapIdle)
	m.Gauges["HeapInuse"] = float64(m.rtm.HeapInuse)
	m.Gauges["HeapObjects"] = float64(m.rtm.HeapObjects)
	m.Gauges["HeapReleased"] = float64(m.rtm.HeapReleased)
	m.Gauges["HeapSys"] = float64(m.rtm.HeapSys)
	m.Gauges["LastGC"] = float64(m.rtm.LastGC)
	m.Gauges["Lookups"] = float64(m.rtm.Lookups)
	m.Gauges["MCacheInuse"] = float64(m.rtm.MCacheInuse)
	m.Gauges["MCacheSys"] = float64(m.rtm.MCacheSys)
	m.Gauges["MSpanSys"] = float64(m.rtm.MSpanSys)
	m.Gauges["Mallocs"] = float64(m.rtm.Mallocs)
	m.Gauges["NextGC"] = float64(m.rtm.NextGC)
	m.Gauges["NumForcedGC"] = float64(m.rtm.NumForcedGC)
	m.Gauges["NextGC"] = float64(m.rtm.NumGC)
	m.Gauges["OtherSys"] = float64(m.rtm.OtherSys)
	m.Gauges["PauseTotalNs"] = float64(m.rtm.PauseTotalNs)
	m.Gauges["StackInuse"] = float64(m.rtm.StackInuse)
	m.Gauges["StackSys"] = float64(m.rtm.StackSys)
	m.Gauges["Sys"] = float64(m.rtm.Sys)
	m.Gauges["TotalAlloc"] = float64(m.rtm.TotalAlloc)
	m.Mutex.Unlock()
}

func NewMetrics() *Metrics {
	var counters = map[string]int64{
		"PollCount": 0,
	}

	var gauges = map[string]float64{
		"Alloc":         0,
		"BuckHashSys":   0,
		"Frees":         0,
		"GCCPUFraction": 0,
		"GCSys":         0,
		"HeapAlloc":     0,
		"HeapIdle":      0,
		"HeapInuse":     0,
		"HeapObjects":   0,
		"HeapReleased":  0,
		"HeapSys":       0,
		"LastGC":        0,
		"Lookups":       0,
		"MCacheInuse":   0,
		"MCacheSys":     0,
		"MSpanInuse":    0,
		"MSpanSys":      0,
		"Mallocs":       0,
		"NextGC":        0,
		"NumForcedGC":   0,
		"NumGC":         0,
		"OtherSys":      0,
		"PauseTotalNs":  0,
		"StackInuse":    0,
		"StackSys":      0,
		"Sys":           0,
		"TotalAlloc":    0,
		"RandomValue":   0,
	}
	return &Metrics{Counters: counters, Gauges: gauges, Mutex: sync.Mutex{}, rtm: runtime.MemStats{}}
}
