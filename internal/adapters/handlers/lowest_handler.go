package handlers

import (
	"log/slog"
	"net/http"

	"marketflow/internal/core/utils"
)

func (h *Handler) GetLowest(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.Service.GetStatService(symbol, exchange, "min", duration)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrResponseInJson(w, err)
		return
	}

	utils.ResponseInJson(w, 200, map[string]interface{}{
		"pair_name":    res.Pair,
		"exchange":     res.Exchange,
		"lowest_price": res.Min,
		"timestamp":    res.Timestamp,
	})
}
