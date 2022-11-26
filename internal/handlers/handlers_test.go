package handlers_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/handlers"
	"github.com/AndrewSukhobok95/yagometrics.git/internal/storage"
	"github.com/go-chi/chi/v5"
)

func TestMetricHandlerUpdateMetric(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name        string
		metricType  string
		metricName  string
		metricValue string
		want        want
	}{
		{
			name:        "Positive test: Gauge metric",
			metricType:  "gauge",
			metricName:  "name",
			metricValue: "111.2",
			want: want{
				code:        200,
				contentType: "text/plain",
			},
		},
		{
			name:        "Positive test: Counter metric",
			metricType:  "counter",
			metricName:  "name",
			metricValue: "111",
			want: want{
				code:        200,
				contentType: "text/plain",
			},
		},
		{
			name:        "Negative test: Unkonwn metric",
			metricType:  "unknown",
			metricName:  "name",
			metricValue: "111",
			want: want{
				code:        501,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:        "Negative test: Wrong value",
			metricType:  "counter",
			metricName:  "name",
			metricValue: "nan",
			want: want{
				code:        400,
				contentType: "text/plain; charset=utf-8",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/update/{metricType}/{metricName}/{metricValue}", bytes.NewBufferString(""))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("metricType", tt.metricType)
			rctx.URLParams.Add("metricName", tt.metricName)
			rctx.URLParams.Add("metricValue", tt.metricValue)
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))
			w := httptest.NewRecorder()
			memStorage := storage.NewMemStorage()
			handler := handlers.NewMetricHandler(memStorage)
			h := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				handler.UpdateMetric(rw, r)
			})
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

func TestMetricHandlerGetMetric(t *testing.T) {
	type want struct {
		code        int
		contentType string
		metricValue string
	}
	tests := []struct {
		name       string
		metricType string
		metricName string
		want       want
	}{
		{
			name:       "Positive test: Gauge metric",
			metricType: "gauge",
			metricName: "G1",
			want: want{
				code:        200,
				contentType: "text/plain",
				metricValue: "111.2",
			},
		},
		{
			name:       "Positive test: Counter metric",
			metricType: "counter",
			metricName: "C1",
			want: want{
				code:        200,
				contentType: "text/plain",
				metricValue: "111",
			},
		},
	}
	memStorage := storage.NewMemStorage()
	memStorage.InsertCounterMetric("C1", 111)
	memStorage.InsertGaugeMetric("G1", 111.2)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/value/{metricType}/{metricName}", bytes.NewBufferString(""))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("metricType", tt.metricType)
			rctx.URLParams.Add("metricName", tt.metricName)
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, rctx))
			w := httptest.NewRecorder()
			handler := handlers.NewMetricHandler(memStorage)
			h := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				handler.GetMetric(rw, r)
			})
			h.ServeHTTP(w, request)
			res := w.Result()
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(resBody) != tt.want.metricValue {
				t.Errorf("Expected body %s, got %s", tt.want.metricValue, w.Body.String())
			}
			if res.Header.Get("Content-Type") != tt.want.contentType {
				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
			}
		})
	}
}
