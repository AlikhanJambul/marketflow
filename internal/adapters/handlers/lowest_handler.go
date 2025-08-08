package handlers

import (
	"marketflow/internal/core/utils"
	"net/http"
)

func (h *Handler) GetLowestSym(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	symbol := r.PathValue("symbol")
	duration := r.URL.Query().Get("period")

	if duration == "" {
		duration = "1m"
	}

	res, err := h.Service.GetLowestSymService(symbol, duration)
	if err != nil {
		utils.ErrResponseInJson(w, err)
		return
	}

	utils.ResponseInJson(w, 200, res)
}

func (h *Handler) GetLowestSymExc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	symbol := r.PathValue("symbol")
	exchange := r.PathValue("exchange")
	duration := r.URL.Query().Get("period")

	if duration == "" {
		duration = "1m"
	}

	res, err := h.Service.GetLowestSymExcService(symbol, exchange, duration)
	if err != nil {
		utils.ErrResponseInJson(w, err)
		return
	}

	utils.ResponseInJson(w, 200, res)
}
