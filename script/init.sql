CREATE TABLE birge_prices (
                                   id SERIAL PRIMARY KEY,
                                   symbol VARCHAR(255),
                                   price FLOAT,
                                   timestamp TIMESTAMPTZ,
                                   exchange VARCHAR(255)
);
