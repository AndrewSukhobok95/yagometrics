package poller_test

import (
	"runtime"
	"testing"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/poller"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
)

func TestUpdateMetrics(t *testing.T) {
	tests := []struct {
		name             string
		numUpdates       int
		wantCounterValue int64
	}{
		{
			name:             "Positive test: Counter metric",
			numUpdates:       10,
			wantCounterValue: 10,
		},
	}
	var rtm runtime.MemStats
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Init Storage
			memStorage := storage.NewMemStorage()
			memStorage.InsertCounterMetric("PollCount", 0)
			// Update metric numUpdates times
			for i := 0; i < tt.numUpdates; i++ {
				poller.UpdateMetrics(memStorage, &rtm)
			}
			value, err := memStorage.GetCounterMetric("PollCount")
			if err != nil {
				t.Errorf("Something wrong with the PollCount key")
			}
			if value != tt.wantCounterValue {
				t.Errorf("Expected PollCount %d, got %d", tt.wantCounterValue, value)
			}
		})
	}
}

func TestPoll(t *testing.T) {
	tests := []struct {
		name         string
		pollInterval time.Duration
	}{
		{
			name:         "Positive test: Poll collects metrics",
			pollInterval: 10 * time.Millisecond,
		},
	}
	numUpdates := 3
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memStorage := storage.NewMemStorage()
			go poller.Poll(memStorage, tt.pollInterval)
			waitInterval := time.Duration(numUpdates) * tt.pollInterval
			time.Sleep(waitInterval)
			metricNames := memStorage.GetAllMetricNames()
			if len(metricNames) == 0 {
				t.Errorf("Poller didn't collect any metrics during %d updates", numUpdates)
			}
		})
	}
}
