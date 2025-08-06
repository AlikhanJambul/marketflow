package postgres

import (
	"context"
	"marketflow/internal/domain/models"
)

func (r *Repository) GetLowestSym(ctx context.Context, symbol string) (models.PriceStats, error) {
	var result models.PriceStats

	err := r.db.QueryRowContext(ctx, `
		SELECT pair_name, min_price, timestamp
		FROM birge_prices
		WHERE pair_name = $1
		ORDER BY min_price ASC
		LIMIT 1
`, symbol).Scan(&result.Pair, &result.Min, &result.Timestamp)

	if err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}

func (r *Repository) GetLowestSymExc(ctx context.Context, symbol string, exchange string) (models.PriceStats, error) {
	var result models.PriceStats

	err := r.db.QueryRowContext(ctx, `
		SELECT pair_name, min_price, timestamp, exchange
		FROM birge_prices
		WHERE pair_name = $1 AND exchange = $2
		ORDER BY min_price ASC
		LIMIT 1
`, symbol, exchange).Scan(&result.Pair, &result.Min, &result.Timestamp, &result.Exchange)

	if err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}
