package serialization

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
	Hash  string   `json:"hash,omitempty"`  // значение хеш-функции
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

func (m *Metrics) GetHash(key string) string {
	var data string
	switch m.MType {
	case "gauge":
		data = fmt.Sprintf("%s:gauge:%f", m.ID, *m.Value)
	case "counter":
		data = fmt.Sprintf("%s:counter:%d", m.ID, *m.Delta)
	}
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (m *Metrics) ToJSON() []byte {
	metric := make(map[string]interface{})
	metric["id"] = m.ID
	metric["type"] = m.MType
	switch m.MType {
	case "gauge":
		metric["value"] = *m.Value
	case "counter":
		metric["delta"] = *m.Delta
	}
	if m.Hash != "" {
		metric["hash"] = m.Hash
	}
	metricMarshal, _ := json.Marshal(metric)
	return metricMarshal
}

func (m *Metrics) ToJSONString() string {
	metricMarshal := m.ToJSON()
	return string(metricMarshal)
}
