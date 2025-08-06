package postgres

import (
	"context"
	"marketflow/internal/domain/models"
)

func (r *Repository) GetAvgSym(ctx context.Context, symbol string) (models.PriceStats, error) {
	var result models.PriceStats

	err := r.db.QueryRowContext(ctx, `
		SELECT pair_name, AVG(average_price), exchange
		FROM birge_prices
		WHERE pair_name = $1
		GROUP BY pair_name, exchange
	`, symbol).Scan(&result.Pair, &result.Average, &result.Exchange)

	if err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}

func (r *Repository) GetAvgSymExc(ctx context.Context, symbol, exchange string) (models.PriceStats, error) {
	var result models.PriceStats

	err := r.db.QueryRowContext(ctx, `
		SELECT pair_name, AVG(average_price), exchange 
		FROM birge_prices
		WHERE pair_name = $1 AND exchange = $2
		GROUP BY pair_name, exchange
	`, symbol, exchange).Scan(&result.Pair, &result.Average, &result.Exchange)

	if err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}
