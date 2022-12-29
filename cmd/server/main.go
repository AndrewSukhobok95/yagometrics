package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/configuration"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/datastorage"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	var wg sync.WaitGroup
	memStorage := datastorage.NewMemStorage()
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

	config := configuration.GetServerConfig()
	datastorage.BackUpToFile(memStorage, config.StoreFile, config.StoreInterval, config.Restore, &wg)
	log.Fatal(http.ListenAndServe(config.Address, r))
	wg.Wait()
}
