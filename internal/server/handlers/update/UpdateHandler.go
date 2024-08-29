package update

import (
	"github.com/enshxx/GoMetricsHub/internal/storage/memstorage"
	"net/http"
	"strconv"
	"strings"
)

func compileURL(path string) (string, string, string, int) {
	params := strings.Split(path, "/")
	if len(params) != 3 {
		return "", "", "", http.StatusNotFound
	}
	return params[0], params[1], params[2], 0
}

func Handler(storage memstorage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}
		mType, name, value, errCode := compileURL(r.URL.String())
		if errCode != 0 {
			http.Error(w, "", errCode)
			return
		}
		switch mType {
		case "gauge":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				http.Error(w, "invalid gauge value", http.StatusBadRequest)
				return
			}
			storage.GaugeSet(name, v)
		case "counter":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				http.Error(w, "invalid counter value", http.StatusBadRequest)
				return
			}
			storage.CounterInc(name, v)
		default:
			http.Error(w, "invalid metric type", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
