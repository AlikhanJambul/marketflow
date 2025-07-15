package worker

import (
	"context"
	"log/slog"
	"marketflow/internal/core/ports"
	"marketflow/internal/domain/models"
	"time"
)

//type BatchInserter struct {
//	ResultChan <-chan models.Prices
//	Repo       ports.PostgresDB
//}
//
//func (w *BatchInserter) StartBatchInsert() {
//	batch := make([]models.Prices, 0, 500)
//
//	for v := range w.ResultChan {
//		batch = append(batch, v)
//
//		if len(batch) >= 500 {
//			_ = w.Repo.BatchInsert(batch)
//			batch = nil
//		}
//	}
//	if len(batch) > 0 {
//		_ = w.Repo.BatchInsert(batch)
//	}
//
//}

func StartBatchInsertLoop(input <-chan models.Prices, r ports.PostgresDB) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	batch := make([]models.Prices, 0, 3000)

	for {
		select {
		case data := <-input:
			batch = append(batch, data)

			if len(batch) >= 3000 {
				slog.Info("Запрос отправлен")
				r.BatchInsert(context.Background(), batch)
				batch = batch[:0]
			}

		case <-ticker.C:
			if len(batch) > 0 {
				r.BatchInsert(context.Background(), batch)
				batch = batch[:0]
			}
		}
	}
}
