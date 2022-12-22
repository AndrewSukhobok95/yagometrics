package serialization

import (
	"fmt"
	"log"

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
	metric += "Type: " + m.MType + " "
	metric += "Name: " + m.ID + " "
	switch {
	case m.Value == nil && m.Delta == nil:
		metric += "Value: nil Delta: nil"
	case m.MType == "gauge":
		metric += "Value: " + fmt.Sprintf("%f", *m.Value)
	case m.MType == "counter":
		metric += "Delta: " + fmt.Sprintf("%d", *m.Delta)
	default:
		metric += "Unknown Type"
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
	log.Printf("Extracting data from storage\n")
	switch {
	case mType == "gauge":
		log.Printf("Extracting Gauge metric\n")
		value, err = storage.GetGaugeMetric(mName)
		log.Printf("Gauge metric extracted\n")
		metric.Value = &value
	case mType == "counter":
		log.Printf("Extracting Counter metric\n")
		delta, err = storage.GetCounterMetric(mName)
		log.Printf("Counter metric extracted\n")
		metric.Delta = &delta
	default:
		err = fmt.Errorf("the given metric type %s doesn't exist", mType)
	}
	return metric, err
}
