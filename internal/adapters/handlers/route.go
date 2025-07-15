package handlers

import (
	"net/http"
)

func InitNewServer() *http.ServeMux {
	mux := http.NewServeMux()

	// Market Data API
	mux.HandleFunc("GET /prices/latest/{symbol}", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /prices/latest/{exchange}/{symbol}", func(w http.ResponseWriter, r *http.Request) {})

	mux.HandleFunc("GET /prices/highest/{symbol}", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /prices/highest/{exchange}/{symbol}", func(w http.ResponseWriter, r *http.Request) {})

	mux.HandleFunc("GET /prices/lowest/{symbol}", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /prices/lowest/{exchange}/{symbol}", func(w http.ResponseWriter, r *http.Request) {})

	mux.HandleFunc("GET /prices/average/{symbol}", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /prices/average/{exchange}/{symbol}", func(w http.ResponseWriter, r *http.Request) {})

	// Data Mode API
	mux.HandleFunc("POST /mode/test", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("POST /mode/live", func(w http.ResponseWriter, r *http.Request) {})

	// System Health
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {})

	return mux
}
