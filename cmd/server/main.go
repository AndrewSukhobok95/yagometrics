package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/configuration"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/datastorage"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config := configuration.GetServerConfig()

	var wg sync.WaitGroup
	memStorage := datastorage.NewMemStorage()
	handler := handlers.NewMetricHandler(memStorage)

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(handlers.GzipHandle)

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

	datastorage.BackUpToFile(memStorage, config.StoreFile, config.StoreInterval, config.Restore, &wg)
	log.Fatal(http.ListenAndServe(config.Address, r))
	wg.Wait()
}
