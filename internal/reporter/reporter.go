package reporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/serialization"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
)

/*func send(client *http.Client, endpoint string) {
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

func Report(client *http.Client, storage storage.Storage, endpoint string, reportInterval time.Duration) {
	ticker := time.NewTicker(reportInterval)
	counters := make(map[string]int64)
	gauges := make(map[string]float64)
	for {
		<-ticker.C
		storage.FillCounterMetricMap(counters)
		storage.FillGaugeMetricMap(gauges)
		for k, v := range counters {
			send(client, fmt.Sprintf("http://%s/update/%s/%s/%f", endpoint, "gauge", k, v))
		}
		for k, v := range gauges {
			send(client, fmt.Sprintf("http://%s/update/%s/%s/%f", endpoint, "gauge", k, v))
		}
	}
}*/

func sendByJSON(client *http.Client, address string, metric serialization.Metrics) {
	metricMarshal, _ := json.Marshal(metric)
	request, err := http.NewRequest(http.MethodPost, address, bytes.NewBuffer(metricMarshal))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
}

func Report(client *http.Client, storage storage.Storage, endpoint string, reportInterval time.Duration) {
	ticker := time.NewTicker(reportInterval)
	address := fmt.Sprintf("http://%s/update/", endpoint)
	for {
		<-ticker.C
		counterNames := storage.GetCounterMetricNames()
		for _, name := range counterNames {
			metricToReturn, err := serialization.GetFilledMetricFromStorage(name, "counter", storage)
			if err != nil {
				log.Fatal(err)
			}
			sendByJSON(client, address, metricToReturn)
		}
		gaugeNames := storage.GetGaugeMetricNames()
		for _, name := range gaugeNames {
			metricToReturn, err := serialization.GetFilledMetricFromStorage(name, "gauge", storage)
			if err != nil {
				log.Fatal(err)
			}
			sendByJSON(client, address, metricToReturn)
		}
	}
}
