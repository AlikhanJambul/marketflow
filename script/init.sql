CREATE TABLE birge_prices (
                                   id SERIAL PRIMARY KEY,
                                   symbol VARCHAR(255),
                                   price FLOAT,
                                   timestamp TIMESTAMPTZ,
                                   exchange VARCHAR(255)
);

CREATE INDEX idx_symbol ON birge_prices(symbol);
CREATE INDEX idx_symbol_time ON birge_prices(symbol, timestamp);
CREATE INDEX idx_symbol_exchange ON birge_prices(symbol, exchange);