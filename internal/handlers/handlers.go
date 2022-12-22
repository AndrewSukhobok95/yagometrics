package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/serialization"
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

func (mh *MetricHandler) UpdateMetricFromJSON(w http.ResponseWriter, r *http.Request) {
	log.Printf("Attempt to update metric by json\n")
	var metric serialization.Metrics
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &metric)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Received metric: " + metric.ToString() + "\n")
	/*if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}*/
	switch {
	case metric.MType == "gauge":
		mh.storage.InsertGaugeMetric(metric.ID, *metric.Value)
	case metric.MType == "counter":
		mh.storage.AddCounterMetric(metric.ID, *metric.Delta)
	default:
		http.Error(w, fmt.Sprintf("%s metric type is not implemented.", metric.MType), http.StatusNotImplemented)
		return
	}
	metricToReturn, err := serialization.GetFilledMetricFromStorage(metric.ID, metric.MType, mh.storage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}
	log.Printf("Returned metric: " + metricToReturn.ToString() + "\n")
	metricToReturnMarshaled, _ := json.Marshal(metricToReturn)
	log.Printf("Sending response to agent\n\n")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(metricToReturnMarshaled)
	//json.NewEncoder(w).Encode(metricToReturn)
}

func (mh *MetricHandler) GetMetricJSON(w http.ResponseWriter, r *http.Request) {
	log.Printf("Attempt to get metric by json\n")
	var metric serialization.Metrics
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &metric)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Received metric: " + metric.ToString() + "\n")
	/*if err := json.NewDecoder(r.Body).Decode(&metric); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}*/
	metricToReturn, err := serialization.GetFilledMetricFromStorage(metric.ID, metric.MType, mh.storage)
	log.Printf("Returned metric: " + metricToReturn.ToString() + "\n")
	if err != nil {
		log.Printf(err.Error() + "\n\n")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	log.Printf("Marshling metric to return\n")
	metricToReturnMarshaled, _ := json.Marshal(metricToReturn)
	log.Printf("Sending response to agent\n\n")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(metricToReturnMarshaled)
	/*if err := json.NewEncoder(w).Encode(metricToReturn); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}*/
}

type MainPageContent struct {
	Metrics string
}

func (mh *MetricHandler) GetMetricList(w http.ResponseWriter, r *http.Request) {
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
