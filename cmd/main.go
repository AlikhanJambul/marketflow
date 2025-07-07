package cmd

import (
	"log/slog"
	"marketflow/internal/adapters/database"
	"marketflow/internal/adapters/exchange"
	"marketflow/internal/adapters/redis"
	"marketflow/internal/config"
	"marketflow/internal/core/models"
	"net/http"
)

func main() {
	cfg := config.Load()

	_, err := database.ConnDb(cfg.DB)
	if err != nil {
		slog.Error("database didn't connect")
		return
	}

	_ = redis.ConnRedis(cfg.Redis)

	sourseOne := make(chan models.Exchange, 333)
	//sourseTwo := make(chan models.Exchange, 333)
	//sourseThree := make(chan models.Exchange, 333)

	sourse := make([]chan models.Exchange, 333)

	exchange.GetDataBirge(cfg.Exchanges, sourseOne)

	mux := http.NewServeMux()

	http.ListenAndServe(cfg.Port, mux)
}
