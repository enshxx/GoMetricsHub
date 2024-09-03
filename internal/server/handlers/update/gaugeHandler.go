package update

import (
	"github.com/enshxx/GoMetricsHub/internal/storage/memstorage"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"strings"
)

func GaugeHandler(storage memstorage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		name := strings.ToLower(chi.URLParam(r, "name"))
		value := strings.ToLower(chi.URLParam(r, "val"))

		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(w, "invalid gauge value", http.StatusBadRequest)
			return
		}
		storage.AddGauge(name, v)

		w.WriteHeader(http.StatusOK)

	}
}
