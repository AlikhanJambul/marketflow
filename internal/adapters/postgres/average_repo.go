package postgres

import (
	"context"
	"marketflow/internal/domain/models"
	"time"
)

func (r *Repository) GetAvgSym(ctx context.Context, symbol string, duration time.Duration) (models.PriceStats, error) {
	var result models.PriceStats

	err := r.db.QueryRowContext(ctx, `
		SELECT  pair_name, average_price, exchange, timestamp
		FROM birge_prices
		WHERE pair_name = $1
		AND timestamp >= NOW() - $2::interval
		ORDER BY timestamp DESC
		LIMIT 1;
	`, symbol, duration).Scan(&result.Pair, &result.Average, &result.Exchange, &result.Timestamp)

	if err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}

func (r *Repository) GetAvgSymExc(ctx context.Context, symbol, exchange string, duration time.Duration) (models.PriceStats, error) {
	var result models.PriceStats

	err := r.db.QueryRowContext(ctx, `
		SELECT  pair_name, average_price, exchange, timestamp
		FROM birge_prices
		WHERE pair_name = $1
		AND exchange = $2
		AND timestamp >= NOW() - $3::interval
		ORDER BY timestamp DESC
		LIMIT 1;
	`, symbol, exchange, duration).Scan(&result.Pair, &result.Average, &result.Exchange, &result.Timestamp)

	if err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}
