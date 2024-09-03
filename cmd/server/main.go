package main

import (
	"github.com/enshxx/GoMetricsHub/internal/server/handlers/get"
	"github.com/enshxx/GoMetricsHub/internal/server/handlers/update"
	"github.com/enshxx/GoMetricsHub/internal/storage/memstorage"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	storage := memstorage.New()
	r := chi.NewRouter()

	r.Route("/update", func(r chi.Router) {
		r.Post("/counter/{name}/{val}", update.CounterHandler(storage))
		r.Post("/gauge/{name}/{val}", update.GaugeHandler(storage))
		r.Post("/{type}/{name}/{val}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
		})
	})
	r.Route("/value", func(r chi.Router) {
		r.Get("/counter/{name}", get.CounterHandler(storage))
		r.Get("/gauge/{name}", get.GaugeHandler(storage))
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
