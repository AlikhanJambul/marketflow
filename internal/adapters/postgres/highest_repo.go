package postgres

import (
	"context"
	"marketflow/internal/domain/models"
	"time"
)

func (r *Repository) GetHighestSym(ctx context.Context, symbol string, duration time.Duration) (models.PriceStats, error) {
	var result models.PriceStats

	err := r.db.QueryRowContext(ctx, `
		SELECT pair_name, max_price, timestamp, exchange
		FROM birge_prices
		WHERE pair_name = $1
		AND timestamp >= NOW() - $2::interval
		ORDER BY max_price DESC
		LIMIT 1
`, symbol, duration).Scan(&result.Pair, &result.Max, &result.Timestamp, &result.Exchange)

	if err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}

func (r *Repository) GetHighestSymExc(ctx context.Context, symbol string, exchange string, duration time.Duration) (models.PriceStats, error) {
	var result models.PriceStats

	err := r.db.QueryRowContext(ctx, `
		SELECT pair_name, max_price, timestamp, exchange
		FROM birge_prices
		WHERE pair_name = $1 
		AND exchange = $2
		AND timestamp >= NOW() - $3::interval
		ORDER BY max_price DESC
		LIMIT 1
`, symbol, exchange, duration).Scan(&result.Pair, &result.Max, &result.Timestamp, &result.Exchange)

	if err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}
