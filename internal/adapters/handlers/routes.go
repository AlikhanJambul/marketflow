package handlers

import (
	"marketflow/internal/domain/ports"
	"net/http"
)

func InitNewServer(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	// Market Data API
	mux.HandleFunc("/prices/latest/{symbol}", h.GetLatestSym)
	mux.HandleFunc("GET /prices/latest/{exchange}/{symbol}", h.GetLatestSymExc)

	mux.HandleFunc("GET /prices/highest/{symbol}", h.GetHighestSym)
	mux.HandleFunc("GET /prices/highest/{exchange}/{symbol}", h.GetHighestSymExc)

	mux.HandleFunc("GET /prices/lowest/{symbol}", h.GetLowestSym)
	mux.HandleFunc("GET /prices/lowest/{exchange}/{symbol}", h.GetLowestSymExc)

	mux.HandleFunc("GET /prices/average/{symbol}", h.GetAvgSym)
	mux.HandleFunc("GET /prices/average/{exchange}/{symbol}", h.GetAvgSymExc)

	// Data Mode API
	mux.HandleFunc("POST /mode/test", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("POST /mode/live", func(w http.ResponseWriter, r *http.Request) {})

	// System Health
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {})

	return mux
}

type Handler struct {
	Service ports.ServiceMethods
}

func InitHandlers(service ports.ServiceMethods) *Handler {
	return &Handler{
		Service: service,
	}
}
