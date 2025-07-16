package application

import (
	"log/slog"
	"marketflow/internal/adapters/exchange"
	"marketflow/internal/application/worker"
	"marketflow/internal/bootstrap"
	"marketflow/internal/domain/models"
	"net/http"
)

func RunServer() {
	cfg := bootstrap.Cfg
	repo := bootstrap.Repo
	mux := bootstrap.Mux

	var sourceArr []models.Sourse
	for _, addr := range cfg.Exchanges {
		ch := make(chan models.Prices, 1000)
		sourceArr = append(sourceArr, models.Sourse{
			SourseChan: ch,
			Addr:       addr,
		})
	}

	var chans []<-chan models.Prices
	for _, s := range sourceArr {
		chans = append(chans, s.SourseChan)
	}

	go exchange.GetDataBirge(sourceArr)
	resultChan := worker.FanIn(chans...)

	go worker.StartBatchInsertLoop(resultChan, repo)

	addr := ":" + cfg.Port
	slog.Info("Сервер запущен", slog.String("port", cfg.Port))
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("Сервер упал", slog.String("error", err.Error()))
	}
}
