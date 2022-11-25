package poller

import (
	"fmt"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/metrics"
)

func Poll(m *metrics.Metrics, pollInterval time.Duration) {
	ticker := time.NewTicker(pollInterval)
	for {
		<-ticker.C
		m.Update()
		fmt.Println("poll")
	}
}
