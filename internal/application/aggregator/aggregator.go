package aggregator

import (
	"context"
	"fmt"
	"log/slog"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
	"strings"
	"time"
)

type Aggregator struct {
	Repo ports.PostgresDB
	out  <-chan models.Prices
}

func NewAggregator(repo ports.PostgresDB, cache ports.Cache, out <-chan models.Prices) *Aggregator {
	return &Aggregator{
		Repo: repo,
		out:  out,
	}
}

func (a *Aggregator) Start(ctx context.Context) {
	buffer := make(map[string][]float64)
	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-a.out:
			if !ok {
				slog.Info("aggregator stopped")
				return
			}

			key := fmt.Sprintf("%s/%s", v.Exchange, v.Symbol)

			buffer[key] = append(buffer[key], v.Price)

		case <-ticker.C:
			if len(buffer) > 0 {
				a.Aggregate(ctx, buffer)
				buffer = make(map[string][]float64)
			}
		}
	}
}

func (a *Aggregator) Aggregate(ctx context.Context, buffer map[string][]float64) {
	var result []models.PriceStats

	for key, prices := range buffer {
		fmt.Println(key, len(prices))

		part := strings.Split(key, "/")

		exc, sym := part[0], part[1]

		var sum, avg, max, min float64

		max, min = prices[0], prices[0]

		for _, val := range prices {
			if val > max {
				max = val
			}

			if val < min {
				min = val
			}

			sum += val
		}

		avg = sum / float64(len(prices))

		resultKey := models.PriceStats{
			Exchange:  exc,
			Pair:      sym,
			Timestamp: time.Now(),
			Average:   avg,
			Max:       max,
			Min:       min,
		}

		result = append(result, resultKey)
	}

	if len(result) > 0 {
		if err := a.Repo.NewBatchInsert(ctx, result); err != nil {
			slog.Error("aggregator err", err)
			return
		} else {
			slog.Info("aggregator success")
		}
	}
}
