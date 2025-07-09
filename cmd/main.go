package main

import (
	"fmt"
	"log/slog"
	"marketflow/internal/adapters/exchange"
	"marketflow/internal/adapters/postgres"
	"marketflow/internal/adapters/redis"
	"marketflow/internal/application/worker"
	"marketflow/internal/core/config"
	"marketflow/internal/domain/models"
	"net/http"
	"os"
)

func main() {
	cfg := config.Load()

	fmt.Println(cfg)

	db, err := postgres.ConnDb(cfg.DB)
	if err != nil {
		slog.Error("postgres didn't connect")
		os.Exit(1)
	}

	rdb := redis.ConnRedis(cfg.Redis)

	repo := postgres.NewRepository(db)
	_ = redis.NewRedisCache(rdb)

	sourseArr := []models.Sourse{}

	for _, v := range cfg.Exchanges {
		s := make(chan models.Prices, 30)
		sourse := models.Sourse{SourseChan: s, Addr: v}
		sourseArr = append(sourseArr, sourse)
	}

	resultChan := make(chan models.Prices, 90)

	exchange.GetDataBirge(sourseArr)

	worker.StartFanInWorkers(sourseArr, resultChan)

	inserter := worker.BatchInserter{
		ResultChan: resultChan,
		Repo:       repo,
	}

	go inserter.StartBatchInsert()

	mux := http.NewServeMux()

	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		slog.Error("server failed", slog.String("error", err.Error()))
	}
}
