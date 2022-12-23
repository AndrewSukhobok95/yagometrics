package main

import (
	"net/http"
	"sync"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/configuration"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/poller"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/reporter"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
)

func main() {
	var wg sync.WaitGroup
	memStorage := storage.NewMemStorage()
	client := &http.Client{}

	config := configuration.GetConfig()

	wg.Add(2)
	go poller.Poll(memStorage, config.PollInterval)
	go reporter.Report(client, memStorage, config.Address, config.ReportInterval)
	wg.Wait()
}
