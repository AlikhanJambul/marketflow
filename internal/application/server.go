package application

import (
	"context"
	"log/slog"
	"marketflow/internal/adapters/exchange"
	"marketflow/internal/application/aggregator"
	"marketflow/internal/application/worker"
	"marketflow/internal/bootstrap"
	"marketflow/internal/domain/models"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func RunServer() {
	cfg := bootstrap.Cfg
	repo := bootstrap.Repo
	cache := bootstrap.Cache
	mux := bootstrap.Mux

	var sourceArr []models.Sourse
	for _, addr := range cfg.Exchanges {
		ch := make(chan models.Prices, 15)
		sourceArr = append(sourceArr, models.Sourse{
			SourseChan: ch,
			Addr:       addr,
		})
	}

	ctx, close := context.WithCancel(context.Background())
	defer close()

	chans := []<-chan models.Prices{}

	for _, source := range sourceArr {
		chans = append(chans, source.SourseChan)
	}

	for _, s := range sourceArr {
		parts := strings.Split(s.Addr, ":")
		exchangeCount := parts[0]

		client := exchange.NewBirgeClient(exchangeCount, s.Addr, s.SourseChan)
		go client.Start(ctx)
	}

	resultChan := worker.FanIn(chans...)

	agr := aggregator.NewAggregator(repo, cache, resultChan)

	go agr.Start(ctx)

	svr := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	if err := svr.ListenAndServe(); err != nil {
		slog.Error("Сервер упал", slog.String("error", err.Error()))
	}
	slog.Info("Сервер запущен", slog.String("port", cfg.Port))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	slog.Info("shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := svr.Shutdown(shutdownCtx); err != nil {
		slog.Error("API shutdown error", "error", err)
	}
	slog.Info("shutdown complete")

}
