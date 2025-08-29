package postgres

import (
	"context"
	"fmt"
	"marketflow/internal/domain/models"
)

func (r *Repository) GetAverage(ctx context.Context, symbol, exchange string, duration string) ([]models.PriceStats, error) {
	query := fmt.Sprintf(`
		SELECT pair_name, exchange, average_price, timestamp
		FROM birge_prices
		WHERE pair_name = $1
		  AND timestamp >= NOW() - $2::interval
	`)

	args := []interface{}{symbol, duration}

	if exchange != "" {
		query += " AND exchange = $3"
		args = append(args, exchange)
	}

	var result []models.PriceStats
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var row models.PriceStats

		if err := rows.Scan(&row.Pair, &row.Exchange, &row.Average, &row.Timestamp); err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	return result, nil
}
