package main

import (
	"net/http"
	"sync"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/configuration"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/datastorage"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/poller"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/reporter"
)

func main() {
	var wg sync.WaitGroup
	memStorage := datastorage.NewMemStorage()
	client := &http.Client{}

	config := configuration.GetAgentConfig()

	wg.Add(2)
	go poller.Poll(memStorage, config.PollInterval)
	go reporter.Report(client, memStorage, config.Address, config.ReportInterval)
	wg.Wait()
}
