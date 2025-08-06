package worker

//
//import (
//	"context"
//	"fmt"
//	"log/slog"
//	"marketflow/internal/domain/models"
//	"marketflow/internal/domain/ports"
//	"time"
//)
//
////type BatchInserter struct {
////	ResultChan <-chan models.Prices
////	Repo       ports.PostgresDB
////}
////
////func (w *BatchInserter) StartBatchInsert() {
////	batch := make([]models.Prices, 0, 500)
////
////	for v := range w.ResultChan {
////		batch = append(batch, v)
////
////		if len(batch) >= 500 {
////			_ = w.Repo.BatchInsert(batch)
////			batch = nil
////		}
////	}
////	if len(batch) > 0 {
////		_ = w.Repo.BatchInsert(batch)
////	}
////
////}
//
//func StartBatchInsertLoop(input <-chan models.Prices, r ports.PostgresDB) {
//	ticker := time.NewTicker(1 * time.Minute)
//	defer ticker.Stop()
//
//	batch := make([]models.Prices, 0, 3000)
//
//	for {
//		select {
//		case data, ok := <-input:
//			if !ok {
//				slog.Warn("Канал закрыт, записываю остатки и выхожу...")
//				if len(batch) > 0 {
//					if err := r.BatchInsert(context.Background(), batch); err != nil {
//						slog.Error(err.Error())
//					}
//				}
//				return
//			}
//
//			batch = append(batch, data)
//			if len(batch) >= 3000 {
//				fmt.Println(len(batch))
//				slog.Info("BatchInsert: 3000 записей")
//
//				if err := r.BatchInsert(context.Background(), batch); err != nil {
//					slog.Error(err.Error())
//				}
//
//				batch = batch[:0]
//			}
//
//			//case <-ticker.C:
//			//	if len(batch) == 2000 || len(batch) == 1000 {
//			//		slog.Info("BatchInsert: по таймеру, кол-во:", slog.Int("count", len(batch)))
//			//
//			//		if err := r.BatchInsert(context.Background(), batch); err != nil {
//			//			slog.Error(err.Error())
//			//		}
//			//
//			//		batch = batch[:0]
//			//	} else if len(batch) != 3000 {
//			//		time.Sleep(2 * time.Second)
//			//		slog.Info("BatchInsert: по таймеру, кол-во:", slog.Int("count", len(batch)))
//			//
//			//		if err := r.BatchInsert(context.Background(), batch); err != nil {
//			//			slog.Error(err.Error())
//			//		}
//			//
//			//		batch = batch[:0]
//			//	}
//		}
//	}
//}
