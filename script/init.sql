CREATE TABLE aggregated_prices (
                                   id SERIAL PRIMARY KEY,
                                   pair_name VARCHAR(255),
                                   exchange VARCHAR(255),
                                   timestamp TIMESTAMPTZ,
                                   value FLOAT
);
