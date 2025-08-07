package worker

import (
	"marketflow/internal/domain/models"
)

type Worker struct {
	InputCh  <-chan models.Prices
	OutputCh chan<- models.Prices
}

func (w *Worker) FanIn() {
	go func() {
		for price := range w.InputCh {
			w.OutputCh <- price
		}
	}()
}
