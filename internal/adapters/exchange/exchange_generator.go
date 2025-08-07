package exchange

import (
	"context"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
	"math/rand"
	"time"
)

type TestClient struct {
	out      chan<- models.Prices
	exchange string
	stopCh   chan struct{}
}

func NewTestClient(out chan<- models.Prices, exchange string) ports.Client {
	return &TestClient{
		out:      out,
		exchange: exchange,
	}
}

func (c *TestClient) Start(ctx context.Context) error {
	pairs := []string{"BTCUSDT", "ETHUSDT", "DOGEUSDT", "TONUSDT", "SOLUSDT"}
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			for _, pair := range pairs {
				price := models.Prices{
					Symbol:    pair,
					Price:     randomPrice(pair),
					Timestamp: time.Now(),
					Exchange:  c.exchange,
				}

				c.out <- price
			}
		case <-c.stopCh:
			return nil
		}
	}

}

func (c *TestClient) Stop() {
	close(c.stopCh)
	return
}

func randomPrice(pair string) float64 {
	base := map[string]float64{
		"BTCUSDT":  60000,
		"ETHUSDT":  3000,
		"DOGEUSDT": 0.12,
		"TONUSDT":  5.5,
		"SOLUSDT":  160,
	}[pair]
	return base + rand.Float64()*base*0.02
}
