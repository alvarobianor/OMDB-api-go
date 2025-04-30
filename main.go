package main

import (
	"OMDB-api/api"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to run", "error", err)
		os.Exit(1)
	}
}

func run() error {
	errDotEnv := godotenv.Load()

	if errDotEnv != nil {
		slog.Error("failed to load .env file", "error", errDotEnv)
		os.Exit(1)
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		slog.Error("API_KEY is not set")
		os.Exit(1)
	}

	handler := api.Handler(apiKey)

	server := http.Server{
		Addr:    ":8081",
		Handler: handler,

		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       time.Minute,
		MaxHeaderBytes:    1024 * 10,
	}

	slog.Info("Starting server ->", "addr", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
