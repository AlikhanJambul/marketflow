package handlers

import (
	"log/slog"
	"marketflow/internal/core/utils"
	"net/http"
)

func (h *Handler) GetAverage(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.Service.GetAvgService(symbol, exchange, duration)
	if err != nil {
		slog.Error(err.Error())
		utils.ErrResponseInJson(w, err)
		return
	}

	var sum float64

	for _, stat := range res {
		sum += stat.Average
	}

	utils.ResponseInJson(w, 200, map[string]interface{}{
		"pair_name":     symbol,
		"exchange":      exchange,
		"average_price": sum / float64(len(res)),
	})
}
