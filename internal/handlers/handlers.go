package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
)

type MetricHandler struct {
	storage storage.Storage
}

func NewMetricHandler(storage storage.Storage) MetricHandler {
	return MetricHandler{storage: storage}
}

func (mh MetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	}
	fields := ParseURL(r.URL.Path)
	if len(fields) != 4 {
		http.Error(w, "Broken address.", http.StatusNotFound)
		return
	}
	metricType := fields[1]
	metricName := fields[2]
	metricValueString := fields[3]
	switch {
	case metricType == "gauge":
		metricValue, err := strconv.ParseFloat(metricValueString, 64)
		if err != nil {
			http.Error(w, "Broken address.", http.StatusBadRequest)
			return
		}
		mh.storage.InsertGaugeMetric(metricName, metricValue)
	case metricType == "counter":
		metricValue, err := strconv.ParseInt(metricValueString, 10, 64)
		if err != nil {
			http.Error(w, "Broken address.", http.StatusBadRequest)
			return
		}
		mh.storage.AddCounterMetric(metricName, metricValue)
	default:
		http.Error(w, fmt.Sprintf("%s metric type is not implemented.", metricName), http.StatusNotImplemented)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

func ParseURL(url string) []string {
	url = strings.TrimPrefix(url, "/")
	url = strings.TrimSuffix(url, "/")
	fields := strings.Split(url, "/")
	return fields
}
