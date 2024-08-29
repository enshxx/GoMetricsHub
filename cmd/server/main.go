package main

import (
	"github.com/enshxx/GoMetricsHub/internal/server/handlers/update"
	"github.com/enshxx/GoMetricsHub/internal/storage/memstorage"
	"net/http"
)

func main() {
	storage := memstorage.New()
	mux := http.NewServeMux()
	mux.Handle("/update/", http.StripPrefix(`/update/`, update.Handler(storage)))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
