package update

import (
	"github.com/enshxx/GoMetricsHub/internal/storage/memstorage"
	"net/http"
	"strconv"
	"strings"
)

func compileURL(path string) (string, string, string, int) {
	params := strings.Split(path, "/")
	if len(params) != 4 {
		return "", "", "", http.StatusNotFound
	}
	return params[1], params[2], params[3], 0
}

func Handler(storage memstorage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		if r.Method != http.MethodPost {
			http.Error(w, "only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}
		mType, name, value, errCode := compileURL(r.URL.String())
		if errCode != 0 {
			http.Error(w, "", http.StatusNotFound)
			return
		}
		switch mType {
		case "gauge":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				http.Error(w, "invalid gauge value", http.StatusBadRequest)
				return
			}
			storage.AddGauge(name, v)
		case "counter":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				http.Error(w, "invalid counter value", http.StatusBadRequest)
				return
			}
			storage.AddCounter(name, v)
		default:
			http.Error(w, "invalid metric type", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)

	}
}
