package postgres

//
//func (r *Repository) GetLastestSym(ctx context.Context, symbol string) (models.Prices, error) {
//	var result models.PriceStats
//
//	err := r.db.QueryRowContext(ctx, `
//		SELECT symbol, price, timestamp
//		FROM birge_prices
//		WHERE symbol = $1
//		ORDER BY timestamp DESC
//		LIMIT 1`, symbol).Scan(&result.Symbol, &result.Price, &result.Timestamp)
//
//	if err != nil {
//		return models.Prices{}, err
//	}
//
//	return result, nil
//}
//
//func (r *Repository) GetLatestSymExc(ctx context.Context, symbol, exchange string) (models.Prices, error) {
//	var result models.Prices
//
//	err := r.db.QueryRowContext(ctx, `
//		SELECT symbol, price, timestamp, exchange
//		FROM birge_prices
//		WHERE symbol = $1 AND exchange = $2
//		ORDER BY timestamp DESC
//		LIMIT 1`, symbol, exchange).Scan(&result.Symbol, &result.Price, &result.Timestamp, &result.Exchange)
//	if err != nil {
//		return models.Prices{}, err
//	}
//
//	return result, nil
//}
