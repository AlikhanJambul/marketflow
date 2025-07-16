package handlers

import (
	"fmt"
	"marketflow/internal/core/utils"
	"net/http"
)

func (h *Handler) GetLatestSym(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	symbol := r.PathValue("symbol")

	result, err := h.Service.GetLatestSymService(symbol)
	if err != nil {
		utils.ErrResponseInJson(w, err)
		return
	}

	utils.ResponseInJson(w, 200, result)
}

func (h *Handler) GetLatestSymExc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	symbol := r.PathValue("symbol")
	exchange := r.PathValue("exchange")

	fmt.Println(symbol, exchange)

	result, err := h.Service.GetLatestSymExcService(symbol, exchange)
	if err != nil {
		utils.ErrResponseInJson(w, err)
		return
	}

	utils.ResponseInJson(w, 200, result)
}
