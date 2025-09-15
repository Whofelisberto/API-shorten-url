// Package api fornece handlers e utilitários para a API HTTP do projeto.
package api

import (
	"encoding/json"
	"errors"
	"http-golang/store"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

func NewHandler(store store.Store) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", handlePost(store))
	r.Get("/{code}", handleGet(store))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("failed to encode json data", "error", err)
		http.Error(w, `{"error":"something went wrong"}`, http.StatusInternalServerError)
	}
}

func handlePost(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "invalid request body"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.ParseRequestURI(body.URL); err != nil {
			sendJSON(w, Response{Error: "invalid URL passed"}, http.StatusBadRequest)
			return
		}

		code, err := store.SaveShortenedURL(r.Context(), body.URL)
		if err != nil {
			sendJSON(w, Response{Error: "failed to save url"}, http.StatusInternalServerError)
			return
		}
		sendJSON(w, Response{Data: code}, http.StatusCreated)
	}
}

func handleGet(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		fullURL, err := store.GetFullURL(r.Context(), code)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				http.Error(w, "url não encontrada", http.StatusNotFound)
				return
			}
			sendJSON(w, Response{Error: "failed to retrieve url"}, http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fullURL, http.StatusPermanentRedirect)
	}
}
