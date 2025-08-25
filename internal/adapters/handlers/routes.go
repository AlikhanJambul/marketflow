package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"marketflow/internal/application/mode"
	"marketflow/internal/core/utils"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
)

func InitNewServer(h *Handler, out chan<- models.Prices) *http.ServeMux {
	mux := http.NewServeMux()

	// Market Data API
	mux.HandleFunc("/prices/latest/{symbol}", h.GetLatest)
	mux.HandleFunc("GET /prices/latest/{exchange}/{symbol}", h.GetLatest)

	mux.HandleFunc("GET /prices/highest/{symbol}", h.GetHighest)
	mux.HandleFunc("GET /prices/highest/{exchange}/{symbol}", h.GetHighest)

	mux.HandleFunc("GET /prices/lowest/{symbol}", h.GetLowest)
	mux.HandleFunc("GET /prices/lowest/{exchange}/{symbol}", h.GetLowest)

	mux.HandleFunc("GET /prices/average/{symbol}", h.GetAverage)
	mux.HandleFunc("GET /prices/average/{exchange}/{symbol}", h.GetAverage)

	// Data Mode API
	mux.HandleFunc("POST /mode/test", h.SetModeTest(out))
	mux.HandleFunc("POST /mode/live", h.SetModeLive(out))

	// System Health
	mux.HandleFunc("GET /health", h.CheckHealth)

	return mux
}

type Handler struct {
	Service ports.ServiceMethods
	Manager *mode.Manager
}

func InitHandlers(service ports.ServiceMethods, manager *mode.Manager) *Handler {
	return &Handler{
		Service: service,
		Manager: manager,
	}
}

func (h *Handler) SetModeTest(input chan<- models.Prices) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		if err := h.Manager.Start(ctx, "test", input); err != nil {
			slog.Error(err.Error())
			utils.ErrResponseInJson(w, err)
			return
		}

		utils.ResponseInJson(w, 200, "ok!")
	}
}

func (h *Handler) SetModeLive(input chan<- models.Prices) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		if err := h.Manager.Start(ctx, "live", input); err != nil {
			slog.Error(err.Error())
			utils.ErrResponseInJson(w, err)
			return
		}

		utils.ResponseInJson(w, 200, "ok!")
	}
}
