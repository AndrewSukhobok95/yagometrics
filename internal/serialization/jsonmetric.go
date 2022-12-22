package serialization

import (
	"fmt"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
)

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

func (m *Metrics) ToString() string {
	metric := ""
	metric += "Type: " + m.MType
	metric += "Name: " + m.ID
	switch m.MType {
	case "gauge":
		metric += "Value: " + fmt.Sprintf("%f", *m.Value)
	case "counter":
		metric += "Delta: " + fmt.Sprintf("%d", *m.Delta)
	}
	return metric
}

func GetFilledMetricFromStorage(mName, mType string, storage storage.Storage) (Metrics, error) {
	var metric Metrics
	var value float64
	var delta int64
	var err error
	metric.ID = mName
	metric.MType = mType
	switch {
	case mType == "gauge":
		value, err = storage.GetGaugeMetric(mName)
		metric.Value = &value
	case mType == "counter":
		delta, err = storage.GetCounterMetric(mName)
		metric.Delta = &delta
	default:
		err = fmt.Errorf("the given metric type %s doesn't exist", mType)
	}
	return metric, err
}
