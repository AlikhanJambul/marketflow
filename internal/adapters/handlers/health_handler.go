package handlers

import (
	"marketflow/internal/core/utils"
	"net/http"
)

func (h *Handler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	res := h.Service.CheckRedisDb(ctx)

	code := 200

	if res.Status == "down" {
		code = 500
	}

	utils.ResponseInJson(w, code, res)
}
