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
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
}

func Report(m *metrics.Metrics, endpoint string, reportInterval time.Duration) {
	ticker := time.NewTicker(reportInterval)
	client := &http.Client{}
	for {
		<-ticker.C
		m.Mutex.Lock()
		counters := m.Counters
		gauges := m.Gauges
		m.Mutex.Unlock()
		for k, v := range counters {
			send(client, fmt.Sprintf("http://%s/update/%s/%s/%d", endpoint, "counter", k, v))
		}
		for k, v := range gauges {
			send(client, fmt.Sprintf("http://%s/update/%s/%s/%f", endpoint, "gauge", k, v))
		}
	}
}
