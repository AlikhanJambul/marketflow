package worker

import (
	"log/slog"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
)

type Worker struct {
	InputCh  <-chan models.Prices
	OutputCh chan<- models.Prices
	cache    ports.Cache
}

func (w *Worker) FanIn() {
	go func() {
		for price := range w.InputCh {
			if err := w.cache.SetLatest(price); err != nil {
				slog.Error(err.Error())
			}

			w.OutputCh <- price
		}
	}()
}
