---

# MarketFlow: Real-Time Market Data Processing System

## Description

**MarketFlow** is a real-time market data processing system built with Go.
It collects cryptocurrency pair prices from multiple exchanges, processes them using concurrency patterns, caches the results in Redis, and stores aggregated data in PostgreSQL.

## Features

* **Two operation modes**:

  * **Live Mode** — real data from exchanges
  * **Test Mode** — synthetic data
* **Concurrent processing** using Fan-in, Fan-out, and Worker Pool patterns
* **Caching** with Redis for fast access to the latest prices
* **Data storage** in PostgreSQL for long-term aggregated data
* **REST API** providing endpoints for market data and statistics
* **Graceful shutdown** — clean shutdown on termination signals

## Architecture

The project follows **Hexagonal Architecture (Ports & Adapters)**:

* **Domain Layer** — business logic and data models
* **Application Layer** — use cases and component orchestration
* **Adapters Layer**:

  * **Web Adapter** — HTTP handlers
  * **Storage Adapter** — PostgreSQL
  * **Cache Adapter** — Redis
  * **Exchange Adapter** — data sources

## Project Setup

### Prerequisites

* Docker and Docker Compose
* Go 1.21+ (for local development)

### Environment Configuration

1. Copy the `.env.example` file to `.env`:

```bash
cp .env.example .env
```

2. Fill in the environment variables in `.env`:

```env
SERVER_PORT=8080

POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=marketflow
POSTGRES_PASSWORD=password
POSTGRES_DB=marketflow
POSTGRES_SSLMODE=disable

REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=redis_password
REDIS_DB=0

EXCHANGE1=exchange1:40101
EXCHANGE2=exchange2:40102
EXCHANGE3=exchange3:40103
```

### Loading Exchange Images

Run the script to load exchange images:

```bash
cd script
./load_images.sh
```

### Running with Docker Compose

```bash
docker-compose up -d
```

The application will be available on the port specified in `SERVER_PORT` (default: **8080**).

### Local Development

For local runs:

```bash
go build -o marketflow .
./marketflow
```

## Usage

### Command Line

```bash
./marketflow --help

Usage:
  marketflow [--port <N>]
  marketflow --help

Options:
  --port N     Port number
```

### API Endpoints

#### Market Data API

* `GET /prices/latest/{symbol}` — Get the latest price for a symbol
* `GET /prices/latest/{exchange}/{symbol}` — Get the latest price from a specific exchange
* `GET /prices/highest/{symbol}` — Get the highest price over a period
* `GET /prices/highest/{exchange}/{symbol}` — Get the highest price from a specific exchange
* `GET /prices/highest/{symbol}?period={duration}` — Get the highest price for a given time period
* `GET /prices/lowest/{symbol}` — Get the lowest price over a period
* `GET /prices/lowest/{exchange}/{symbol}` — Get the lowest price from a specific exchange
* `GET /prices/lowest/{symbol}?period={duration}` — Get the lowest price for a given time period
* `GET /prices/average/{symbol}` — Get the average price over a period
* `GET /prices/average/{exchange}/{symbol}` — Get the average price from a specific exchange
* `GET /prices/average/{exchange}/{symbol}?period={duration}` — Get the average price for a given period from a specific exchange

#### Data Mode API

* `POST /mode/test` — Switch to **Test Mode**
* `POST /mode/live` — Switch to **Live Mode**

#### System Health

* `GET /health` — Check system health (connections, Redis availability, etc.)

## Project Structure

```
marketflow/
├── cmd/                 # Application entry point
├── internal/            # Internal packages
│   ├── adapters/        # Adapters
│   │   ├── exchange/    # Exchange adapters
│   │   ├── handlers/    # HTTP handlers
│   │   ├── postgres/    # PostgreSQL adapter
│   │   └── redis/       # Redis adapter
│   ├── application/     # Use cases and services
│   │   ├── aggregator/  # Data aggregation
│   │   ├── mode/        # Mode management
│   │   ├── usecase/     # Business logic
│   │   └── worker/      # Worker pool
│   ├── bootstrap/       # Application initialization
│   ├── core/            # Core utilities
│   │   ├── apperrors/   # Application errors
│   │   ├── config/      # Configuration
│   │   └── utils/       # Helper utilities
│   └── domain/          # Domain layer
│       ├── models/      # Data models
│       └── ports/       # Interfaces
├── script/              # Scripts
│   └── load_images.sh   # Exchange image loading script
└── docker-compose.yml   # Docker Compose configuration
```

## Supported Trading Pairs

The system processes the following cryptocurrency pairs:

* BTCUSDT
* DOGEUSDT
* TONUSDT
* SOLUSDT
* ETHUSDT

## Database

A `birge_prices` table is created in PostgreSQL with the following schema:

```sql
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
```

## Logging

The system uses Go’s `log/slog` package with different logging levels:

* **Info** — informational messages
* **Warning** — warnings
* **Error** — errors

## Graceful Shutdown

The application gracefully handles **SIGINT** and **SIGTERM** signals, ensuring clean shutdown of all components.

## Monitoring

The `/health` endpoint provides system health information, including Redis and exchange connection status.

---
