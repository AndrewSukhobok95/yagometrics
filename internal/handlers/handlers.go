package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
	"github.com/go-chi/chi/v5"
)

type MetricHandler struct {
	storage storage.Storage
}

func NewMetricHandler(storage storage.Storage) MetricHandler {
	return MetricHandler{storage: storage}
}

func (mh *MetricHandler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
		return
	}
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	metricValueString := chi.URLParam(r, "metricValue")
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

func (mh *MetricHandler) GetMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed.", http.StatusMethodNotAllowed)
		return
	}
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	var metricValueString string
	switch {
	case metricType == "gauge":
		metricValue, err := mh.storage.GetGaugeMetric(metricName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Metric %s doesn't exist.", metricName), http.StatusNotFound)
			return
		}
		metricValueString = strconv.FormatFloat(metricValue, 'f', -1, 64)
	case metricType == "counter":
		metricValue, err := mh.storage.GetCounterMetric(metricName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Metric %s doesn't exist.", metricName), http.StatusNotFound)
			return
		}
		metricValueString = fmt.Sprintf("%d", metricValue)
	default:
		http.Error(w, fmt.Sprintf("%s metric type is not implemented.", metricName), http.StatusNotImplemented)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(metricValueString))
}

type MainPageContent struct {
	Metrics string
}

func (mh *MetricHandler) GetMetricList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed.", http.StatusMethodNotAllowed)
		return
	}
	metricsSlice := mh.storage.GetAllMetricNames()
	var metricsNames string
	if len(metricsSlice) != 0 {
		metricsNames = strings.Join(metricsSlice, "; ")
	} else {
		metricsNames = ""
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	// Reading html
	content := MainPageContent{Metrics: metricsNames}
	parsedTemplate, _ := template.ParseFiles("./web/main.html")
	err := parsedTemplate.Execute(w, content)
	if err != nil {
		http.Error(w, "Page not found", http.StatusInternalServerError)
	}
}
