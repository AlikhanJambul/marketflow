package worker

import (
	"marketflow/internal/core/ports"
	"marketflow/internal/domain/models"
)

type BatchInserter struct {
	ResultChan <-chan models.Prices
	Repo       ports.PostgresDB
}

func (w *BatchInserter) StartBatchInsert() {
	batch := make([]models.Prices, 0, 45)

	for v := range w.ResultChan {
		batch = append(batch, v)

		if len(batch) >= 500 {
			_ = w.Repo.BatchInsert(batch)
			batch = nil
		}
	}
	if len(batch) > 0 {
		_ = w.Repo.BatchInsert(batch)
	}

}
