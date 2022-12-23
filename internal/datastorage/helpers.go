package datastorage

import (
	"fmt"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/serialization"
)

func GetFilledMetricFromStorage(mName, mType string, storage Storage) (serialization.Metrics, error) {
	var metric serialization.Metrics
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
