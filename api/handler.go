package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Handler() http.Handler {
	handler := chi.NewMux()

	handler.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	return handler
}
