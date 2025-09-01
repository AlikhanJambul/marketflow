package application

import (
	"context"
	"log/slog"
	"marketflow/internal/adapters/handlers"
	"marketflow/internal/application/aggregator"
	"marketflow/internal/application/worker"
	"marketflow/internal/bootstrap"
	"marketflow/internal/domain/models"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	cfg := bootstrap.Cfg
	repo := bootstrap.Repo
	cache := bootstrap.Cache
	manager := bootstrap.Manager
	handler := bootstrap.Handlers

	ctx, close := context.WithCancel(context.Background())
	defer close()

	input := make(chan models.Prices, 50)
	output := make(chan models.Prices, 50)

	agr := aggregator.NewAggregator(repo, cache, output)

	go agr.Start(ctx)

	for i := 0; i <= 5; i++ {
		pool := worker.Worker{
			InputCh:  input,
			OutputCh: output,
			Cache:    cache,
		}
		go pool.FanIn()
	}

	err := manager.Start(ctx, "live", input)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	svr := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handlers.InitNewServer(handler, input),
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
