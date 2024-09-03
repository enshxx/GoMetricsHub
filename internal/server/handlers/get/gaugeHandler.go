package get

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

		val, ok := storage.GetGauge(name)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, err := w.Write([]byte(strconv.FormatFloat(val, 'f', -1, 64)))
		if err != nil {
			http.Error(w, "write error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
