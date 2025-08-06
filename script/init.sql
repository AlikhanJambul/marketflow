CREATE TABLE birge_prices (
                              pair_name TEXT NOT NULL,
                              exchange TEXT NOT NULL,
                              timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                              average_price FLOAT8 NOT NULL,
                              min_price FLOAT8 NOT NULL,
                              max_price FLOAT8 NOT NULL
);


CREATE INDEX idx_symbol ON birge_prices(pair_name);
CREATE INDEX idx_symbol_time ON birge_prices(pair_name, timestamp);
CREATE INDEX idx_symbol_exchange ON birge_prices(pair_name, exchange);