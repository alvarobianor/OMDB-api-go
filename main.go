package main

import (
	"OMDB-api/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to run", "error", err)
		os.Exit(1)
	}
}

func run() error {
	handler := api.Handler()

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
