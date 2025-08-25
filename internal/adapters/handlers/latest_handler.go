package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"marketflow/internal/core/utils"
)

func (h *Handler) GetLatest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	symbol := r.PathValue("symbol")
	exchange := r.PathValue("exchange")

	fmt.Println(symbol, exchange)

	result, err := h.Service.GetLatestService(symbol, exchange)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrResponseInJson(w, err)
		return
	}

	utils.ResponseInJson(w, 200, result)
}
