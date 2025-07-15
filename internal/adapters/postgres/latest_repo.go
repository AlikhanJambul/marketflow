package postgres

import (
	"context"
	"marketflow/internal/domain/models"
)

func (r *Repository) GetLastestBySymbol(ctx context.Context) (models.Prices, error) {
	var result models.Prices

	err := r.db.QueryRowContext(ctx, `
		SELECT symbol, price, timestamp 
		FROM prices 
		ORDER BY timestamp DESC 
		LIMIT 1`).Scan(&result.Symbol, &result.Price, &result.Timestamp)

	if err != nil {
		return models.Prices{}, err
	}

	return result, nil
}
