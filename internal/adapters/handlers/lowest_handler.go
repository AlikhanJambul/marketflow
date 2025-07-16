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

	res, err := h.Service.GetLatestSymService(symbol)
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

	res, err := h.Service.GetLowestSymExcService(symbol, exchange)
	if err != nil {
		utils.ErrResponseInJson(w, err)
		return
	}

	utils.ResponseInJson(w, 200, res)
}
