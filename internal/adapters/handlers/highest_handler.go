package handlers

import (
	"log/slog"
	"marketflow/internal/core/utils"
	"net/http"
)

func (h *Handler) GetHighest(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.Service.GetStatService(symbol, exchange, "max", duration)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrResponseInJson(w, err)
		return
	}

	utils.ResponseInJson(w, 200, map[string]interface{}{
		"pair_name":     res.Pair,
		"exchange":      res.Exchange,
		"highest_price": res.Max,
		"timestamp":     res.Timestamp,
	})
}
