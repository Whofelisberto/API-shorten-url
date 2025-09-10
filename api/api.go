// Package api fornece handlers e utilitários para a API HTTP do projeto.
package api

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func NewHandler(db map[string]string) http.Handler {
  r := chi.NewMux()

 r.Use(middleware.Recoverer)
 r.Use(middleware.RequestID)
 r.Use(middleware.Logger)

 r.Post("/api/shorten", handlePost(db))
 r.Get("/{code}", handleGet(db))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data any `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("failed to encode json data", "error", err)
		http.Error(w, `{"error":"something went wrong"}`, http.StatusInternalServerError)
	}
}

// func de post
func handlePost(db map[string]string) http.HandlerFunc { return func (w http.ResponseWriter, r *http.Request) {
	var body PostBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sendJSON(w, Response{Error: "invalid request body"}, http.StatusUnprocessableEntity)
		return
	}
if  _, err:= url.Parse(body.URL); err != nil {
	sendJSON (w, Response{Error: "invalid URL passed"}, http.StatusBadRequest)
}
 code := genCode()
 db[code] = body.URL
 sendJSON(w, Response{Data: code}, http.StatusCreated)
}}

const characters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func genCode() string {
	const n = 8
	byts := make([]byte, n)
	for i := range n {
		byts[i] = characters[rand.Intn(len(characters))]

	}
	return string(byts)
}

// func de get
func handleGet(db map[string]string) http.HandlerFunc { return func (w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	data, ok := db[code]
	if !ok {
		http.Error(w, "url não encontrada", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, data, http.StatusPermanentRedirect)
}}