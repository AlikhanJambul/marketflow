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
		s := make(chan models.Prices, 1000)
		sourse := models.Sourse{SourseChan: s, Addr: v}
		sourseArr = append(sourseArr, sourse)
	}

	var chans []<-chan models.Prices
	for _, s := range sourseArr {
		chans = append(chans, s.SourseChan)
	}

	go exchange.GetDataBirge(sourseArr)
	resultChan := worker.FanIn(chans...)

	fmt.Println(len(resultChan))

	worker.StartBatchInsertLoop(resultChan, repo)

	mux := http.NewServeMux()
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		slog.Error("server failed", slog.String("error", err.Error()))
	}
}
