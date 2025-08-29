package utils

import (
	"encoding/json"
	"log/slog"
	"marketflow/internal/core/apperrors"
	"net/http"
)

func ResponseInJson(w http.ResponseWriter, statusCode int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(object)
}

func ErrResponseInJson(w http.ResponseWriter, err error) {
	slog.Error(err.Error())

	statusCode := apperrors.CheckCode(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
