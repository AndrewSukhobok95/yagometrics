package reporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/datastorage"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/serialization"
)

func sendByJSON(client *http.Client, address string, metric serialization.Metrics) {
	metricMarshal, _ := json.Marshal(metric)
	request, err := http.NewRequest(http.MethodPost, address, bytes.NewBuffer(metricMarshal))
	if err != nil {
		log.Printf("Error in creating the request:\n")
		log.Printf(err.Error() + "\n\n")
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error in receiving the response:\n")
		log.Printf(err.Error() + "\n\n")
	} else {
		defer response.Body.Close()
	}
}

func Report(client *http.Client, storage datastorage.Storage, endpoint string, reportInterval time.Duration) {
	ticker := time.NewTicker(reportInterval)
	address := fmt.Sprintf("http://%s/update/", endpoint)
	for {
		<-ticker.C
		counterNames := storage.GetCounterMetricNames()
		for _, name := range counterNames {
			metricToReturn, err := datastorage.GetFilledMetricFromStorage(name, "counter", storage)
			if err != nil {
				log.Printf("Error in extracting the metric from storage:\n")
				log.Printf(err.Error() + "\n\n")
			} else {
				sendByJSON(client, address, metricToReturn)
			}
		}
		gaugeNames := storage.GetGaugeMetricNames()
		for _, name := range gaugeNames {
			metricToReturn, err := datastorage.GetFilledMetricFromStorage(name, "gauge", storage)
			if err != nil {
				log.Printf("Error in extracting the metric from storage:\n")
				log.Printf(err.Error() + "\n\n")
			} else {
				sendByJSON(client, address, metricToReturn)
			}
		}
	}
}
