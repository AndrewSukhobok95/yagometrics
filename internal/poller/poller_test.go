package poller_test

import (
	"testing"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/poller"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
)

func TestPoll(t *testing.T) {
	tests := []struct {
		name             string
		pollInterval     time.Duration
		waitInterval     time.Duration
		wantCounterValue int64
	}{
		{
			name:             "Positive test: Counter metric",
			pollInterval:     10 * time.Millisecond,
			waitInterval:     105 * time.Millisecond,
			wantCounterValue: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memStorage := storage.NewMemStorage()
			go poller.Poll(memStorage, tt.pollInterval)
			time.Sleep(tt.waitInterval)
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
