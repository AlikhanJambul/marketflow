package worker

import (
	"context"
	"log/slog"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
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
		case data, ok := <-input:
			if !ok {
				slog.Warn("Канал закрыт, записываю остатки и выхожу...")
				if len(batch) > 0 {
					if err := r.BatchInsert(context.Background(), batch); err != nil {
						slog.Error(err.Error())
					}
				}
				return
			}

			batch = append(batch, data)
			if len(batch) >= 3000 {
				slog.Info("BatchInsert: 3000 записей")
				if err := r.BatchInsert(context.Background(), batch); err != nil {
					slog.Error(err.Error())
				}
				batch = batch[:0]
			}

			//case <-ticker.C:
			//	if len(batch) > 0 {
			//		slog.Info("BatchInsert: по таймеру, кол-во:", slog.Int("count", len(batch)))
			//		if err := r.BatchInsert(context.Background(), batch); err != nil {
			//			slog.Error(err.Error())
			//		}
			//		batch = batch[:0]
			//	}
		}
	}
}
