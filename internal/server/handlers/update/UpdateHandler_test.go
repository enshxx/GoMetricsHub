package update

import (
	"github.com/enshxx/GoMetricsHub/internal/storage/memstorage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	type want struct {
		code         int
		response     string
		contentType  string
		storageKey   string
		storageValue float64
	}
	tests := []struct {
		name       string
		url        string
		method     string
		metricType string
		want       want
	}{
		{
			name:       "simple counter test POST",
			method:     http.MethodPost,
			url:        "/counter/someMetric/527",
			metricType: "counter",
			want: want{
				code:         http.StatusOK,
				response:     "",
				contentType:  "text/plain; charset=utf-8",
				storageKey:   "someMetric",
				storageValue: 527,
			},
		},
		{
			name:       "Get method",
			method:     http.MethodGet,
			url:        "/counter/someMetric/527",
			metricType: "counter",
			want: want{
				code:         http.StatusMethodNotAllowed,
				response:     "only POST requests are allowed\n",
				contentType:  "text/plain; charset=utf-8",
				storageKey:   "",
				storageValue: 0,
			},
		},
		{
			name:       "Get method with invalid params",
			method:     http.MethodGet,
			url:        "/countersomeMetric527",
			metricType: "counter",
			want: want{
				code:         http.StatusMethodNotAllowed,
				response:     "only POST requests are allowed\n",
				contentType:  "text/plain; charset=utf-8",
				storageKey:   "",
				storageValue: 0,
			},
		},
		{
			name:       "Counter with float value",
			method:     http.MethodPost,
			url:        "/counter/banana/527.0",
			metricType: "counter",
			want: want{
				code:         http.StatusBadRequest,
				response:     "invalid counter value\n",
				contentType:  "text/plain; charset=utf-8",
				storageKey:   "",
				storageValue: 0,
			},
		},
		{
			name:       "Gauge with float value",
			method:     http.MethodPost,
			url:        "/gauge/banana/527.0",
			metricType: "gauge",
			want: want{
				code:         http.StatusOK,
				response:     "",
				contentType:  "text/plain; charset=utf-8",
				storageKey:   "banana",
				storageValue: 527.0,
			},
		},
		{
			name:       "not found",
			method:     http.MethodPost,
			url:        "/gauge/",
			metricType: "gauge",
			want: want{
				code:         http.StatusNotFound,
				response:     "\n",
				contentType:  "text/plain; charset=utf-8",
				storageKey:   "\n",
				storageValue: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stg := memstorage.New()
			r := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()
			Handler(stg)(w, r)

			// Проверка ответа
			assert.Equal(t, tt.want.code, w.Code, "Invalid status code")
			assert.Equal(t, tt.want.contentType, w.Header().Get("Content-Type"), "Invalid header: Content-type")
			assert.Equal(t, tt.want.response, w.Body.String(), "Invalid body")

			// Проверка изменения данных
			if tt.want.storageKey == "" {
				return
			}
			switch tt.metricType {
			case "counter":
				val, _ := stg.GetCounter(tt.want.storageKey)
				assert.Equal(t, tt.want.storageValue, float64(val))
			case "gauge":
				val, _ := stg.GetGauge(tt.want.storageKey)
				assert.Equal(t, tt.want.storageValue, val)
			}
		})
	}
}
