package cmd

import (
	"marketflow/internal/config"
	"net/http"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()

	http.ListenAndServe(cfg.Port, mux)
}
