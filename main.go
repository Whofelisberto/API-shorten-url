package main

import (
	"http-golang/api"
	"http-golang/store"
	"log/slog"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all systems online")
}

func run() error {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	store := store.NewStore(rdb)
	handler := api.NewHandler(store)

	s := &http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}