package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/poller"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/reporter"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
)

const (
	pollInterval   = time.Duration(2 * time.Second)
	reportInterval = time.Duration(10 * time.Second)
	endpoint       = "127.0.0.1:8080"
)

func main() {
	var wg sync.WaitGroup
	memStorage := storage.NewMemStorage()
	client := &http.Client{}
	wg.Add(2)
	go poller.Poll(memStorage, pollInterval)
	go reporter.Report(client, memStorage, endpoint, reportInterval)
	wg.Wait()
}
