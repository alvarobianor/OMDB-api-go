package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

const apiUrl = "https://www.omdbapi.com/?s=%s&apikey=%s"

func Handler(apiKey string) http.Handler {
	router := chi.NewMux()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(corsMiddleware.Handler)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("X-CUSTOM-HEADER-API-KEY", apiKey)
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.Get("/", getMovie())

	return router
}

func getMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKeyHeader := r.Header.Get("X-CUSTOM-HEADER-API-KEY")
		if apiKeyHeader == "" {
			errMessage := "X-CUSTOM-HEADER-API-KEY is not set"
			slog.Error("Unauthorized", "error", errMessage)
			SendJson(w, Response{Error: errMessage}, http.StatusUnauthorized)
			return
		}

		movieName := r.URL.Query().Get("movie")

		if movieName == "" {
			errMessage := "Movie name is not set"
			slog.Error("Bad Request", "error", errMessage)
			SendJson(w, Response{Error: errMessage}, http.StatusBadRequest)
			return
		}

		concatUrlApi := fmt.Sprintf(apiUrl, movieName, apiKeyHeader)

		resp, err := http.Get(concatUrlApi)
		if err != nil {
			errMessage := "Failed to get movie"
			slog.Error("Failed to get movie", "error", err)
			SendJson(w, Response{Error: errMessage}, http.StatusInternalServerError)
			return
		}

		var data any

		if errDecode := json.NewDecoder(resp.Body).Decode(&data); errDecode != nil {
			errMessage := "Failed to decode movie"
			slog.Error("Failed to decode movie", "error", errDecode)
			SendJson(w, Response{Error: errMessage}, http.StatusInternalServerError)
			return
		}

		SendJson(w, Response{Data: data}, http.StatusOK)
	}
}

func SendJson(w http.ResponseWriter, r Response, code int) {
	data, err := json.Marshal(r)

	if err != nil {
		slog.Error(err.Error(), "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	_, err = w.Write(data)

	if err != nil {
		slog.Error(err.Error(), "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
