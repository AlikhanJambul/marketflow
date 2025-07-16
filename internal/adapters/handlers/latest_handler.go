package handlers

import (
	"marketflow/internal/core/utils"
	"net/http"
)

func (h *Handler) LatestSymbol(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	symbol := r.PathValue("symbol")

	result, err := h.Service.GetLatestBySymbolService(symbol)
	if err != nil {
		utils.ErrResponseInJson(w, err)
		return
	}

	utils.ResponseInJson(w, 200, result)
}
