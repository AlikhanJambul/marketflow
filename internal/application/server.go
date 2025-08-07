package application

import (
	"context"
	"log/slog"
	"marketflow/internal/application/aggregator"
	"marketflow/internal/application/worker"
	"marketflow/internal/bootstrap"
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
	mux := bootstrap.Mux
	manager := bootstrap.Manager

	ctx, close := context.WithCancel(context.Background())
	defer close()

	chans, err := manager.Start(ctx, "live")
	if err != nil {
		slog.Error(err.Error())
		return
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
