package main

import (
	"log"
	"net/http"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/handlers"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
	"github.com/go-chi/chi/v5"
)

const (
	endpoint = "127.0.0.1:8080"
)

func main() {
	memStorage := storage.NewMemStorage()
	handler := handlers.NewMetricHandler(memStorage)

	r := chi.NewRouter()

	//r.Use(middleware.RequestID)
	//r.Use(middleware.RealIP)
	//r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
			handler.GetMetricList(rw, r)
		})
		r.Post("/update/{metricType}/{metricName}/{metricValue}", func(rw http.ResponseWriter, r *http.Request) {
			handler.UpdateMetric(rw, r)
		})
		r.Get("/value/{metricType}/{metricName}", func(rw http.ResponseWriter, r *http.Request) {
			handler.GetMetric(rw, r)
		})
		r.Post("/update/", func(rw http.ResponseWriter, r *http.Request) {
			handler.UpdateMetricFromJSON(rw, r)
		})
		r.Post("/value/", func(rw http.ResponseWriter, r *http.Request) {
			handler.GetMetricJSON(rw, r)
		})
	})

	log.Fatal(http.ListenAndServe(endpoint, r))
}
