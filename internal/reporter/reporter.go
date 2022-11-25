package reporter

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/metrics"
)

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

func Report(m *metrics.Metrics, endpoint string, reportInterval time.Duration) {
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
