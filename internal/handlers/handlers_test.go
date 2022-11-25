package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/handlers"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
)

func TestMetricHandler(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code:        200,
				contentType: "text/plain",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/update/gauge/Name1/111", bytes.NewBufferString(""))
			w := httptest.NewRecorder()
			memStorage := storage.NewMemStorage()
			handler := handlers.NewMetricHandler(memStorage)
			h := http.Handler(handler)
			h.ServeHTTP(w, request)
			res := w.Result()
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}
			defer res.Body.Close()
			if res.Header.Get("Content-Type") != tt.want.contentType {
				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
			}
		})
	}
}
