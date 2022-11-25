package main

import (
	"sync"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/metrics"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/poller"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/reporter"
)

const (
	pollInterval   = time.Duration(2 * time.Second)
	reportInterval = time.Duration(10 * time.Second)
	endpoint       = "127.0.0.1:8080"
)

func main() {
	var wg sync.WaitGroup
	m := metrics.NewMetrics()
	wg.Add(2)
	go poller.Poll(m, pollInterval)
	go reporter.Report(m, endpoint, reportInterval)
	wg.Wait()
}
