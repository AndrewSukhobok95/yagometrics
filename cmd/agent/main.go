package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/metrics"
)

const (
	pollInterval   = time.Duration(2 * time.Second)
	reportInterval = time.Duration(10 * time.Second)
	endpoint       = "127.0.0.1:8080"
)

func poll(m *metrics.Metrics) {
	ticker := time.NewTicker(pollInterval)
	for {
		<-ticker.C
		m.Update()
		fmt.Println("poll")
	}
}

func send(client *http.Client, endpoint string) {
	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBufferString(""))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Content-Type", "text/plain")
	fmt.Println("do-start")
	response, err := client.Do(request)
	fmt.Println("do-end")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
}

func report(m *metrics.Metrics) {
	ticker := time.NewTicker(reportInterval)
	client := &http.Client{}
	fmt.Println("report-init")
	for {
		<-ticker.C
		fmt.Println("report-tick")
		m.Mutex.Lock()
		counters := m.Counters
		gauges := m.Gauges
		m.Mutex.Unlock()
		fmt.Println("report-start")
		for k, v := range counters {
			send(client, fmt.Sprintf("http://%s/update/%s/%s/%d", endpoint, "counter", k, v))
		}
		for k, v := range gauges {
			send(client, fmt.Sprintf("http://%s/update/%s/%s/%f", endpoint, "gauge", k, v))
		}
		fmt.Println("report-end")
	}
}

func main() {
	var wg sync.WaitGroup
	m := metrics.NewMetrics()
	fmt.Println("start")
	wg.Add(2)
	go poll(m)
	go report(m)
	wg.Wait()
}
