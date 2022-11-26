package storage

import (
	"fmt"
	"sync"
)

type Storage interface {
	InsertGaugeMetric(name string, value float64)
	InsertCounterMetric(name string, value int64)
	AddCounterMetric(name string, value int64)
	GetGaugeMetric(name string) (float64, error)
	GetCounterMetric(name string) (int64, error)
	GetAllMetricNames() []string
}

type MemStorage struct {
	counters map[string]int64
	gauges   map[string]float64
	mutex    sync.Mutex
}

func NewMemStorage() *MemStorage {
	var ms MemStorage
	ms.counters = make(map[string]int64)
	ms.gauges = make(map[string]float64)
	ms.mutex = sync.Mutex{}
	return &ms
}

func (ms *MemStorage) InsertGaugeMetric(name string, value float64) {
	ms.mutex.Lock()
	ms.gauges[name] = value
	ms.mutex.Unlock()
}

func (ms *MemStorage) InsertCounterMetric(name string, value int64) {
	ms.mutex.Lock()
	ms.counters[name] = value
	ms.mutex.Unlock()
}

func (ms *MemStorage) AddCounterMetric(name string, value int64) {
	ms.mutex.Lock()
	ms.counters[name] += value
	ms.mutex.Unlock()
}

func (ms *MemStorage) GetGaugeMetric(name string) (float64, error) {
	ms.mutex.Lock()
	value, ok := ms.gauges[name]
	ms.mutex.Unlock()
	if ok {
		return value, nil
	} else {
		e := fmt.Errorf("The given metric name %s doesn't exist", name)
		return 0, e
	}
}

func (ms *MemStorage) GetCounterMetric(name string) (int64, error) {
	ms.mutex.Lock()
	value, ok := ms.counters[name]
	ms.mutex.Unlock()
	if ok {
		return value, nil
	} else {
		e := fmt.Errorf("The given metric name %s doesn't exist", name)
		return 0, e
	}
}

func (ms *MemStorage) GetAllMetricNames() []string {
	names := make([]string, len(ms.counters)+len(ms.gauges))
	for k := range ms.counters {
		names = append(names, k)
	}
	for k := range ms.gauges {
		names = append(names, k)
	}
	return names
}
