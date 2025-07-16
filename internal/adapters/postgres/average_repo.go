package postgres

import (
	"context"
	"marketflow/internal/domain/models"
)

func (r *Repository) GetAvgSym(ctx context.Context, symbol string) (models.Prices, error) {
	var result models.Prices

	err := r.db.QueryRowContext(ctx, `
		SELECT symbol, AVG(price)
		FROM birge_prices
		WHERE symbol = $1
		GROUP BY symbol
	`, symbol).Scan(&result.Symbol, &result.Price)

	if err != nil {
		return models.Prices{}, err
	}

	return result, nil
}

func (r *Repository) GetAvgSymExc(ctx context.Context, symbol, exchange string) (models.Prices, error) {
	var result models.Prices

	err := r.db.QueryRowContext(ctx, `
		SELECT symbol, AVG(price), exchange
		FROM birge_prices
		WHERE symbol = $1 AND exchange = $2
		GROUP BY symbol, exchange
	`, symbol, exchange).Scan(&result.Symbol, &result.Price, &result.Exchange)

	if err != nil {
		return models.Prices{}, err
	}

	return result, nil
}
