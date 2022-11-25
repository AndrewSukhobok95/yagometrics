package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/handlers"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
)

const (
	endpoint = "127.0.0.1:8080"
)

func showMetrics(ms storage.Storage) {
	for {
		vg, err := ms.GetGaugeMetric("Alloc")
		if err != nil {
			fmt.Println("Alloc error")
		} else {
			fmt.Printf("Alloc %f \n", vg)
		}
		vc, err := ms.GetCounterMetric("PollCount")
		if err != nil {
			fmt.Println("PollCount error")
		} else {
			fmt.Printf("PollCount %d \n", vc)
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	memStorage := storage.NewMemStorage()
	handler := handlers.NewMetricHandler(memStorage)
	mux := http.NewServeMux()
	mux.Handle("/update/", handler)

	go showMetrics(memStorage)

	fmt.Println("Start server")
	server := &http.Server{
		Addr:    endpoint,
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
