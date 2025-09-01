# MarketFlow: Система обработки рыночных данных в реальном времени

## Описание

MarketFlow - это система для обработки рыночных данных в реальном времени, разработанная на Go. Система получает данные о ценах криптовалютных пар с нескольких бирж, обрабатывает их с использованием паттернов конкурентности, кэширует в Redis и сохраняет агрегированные данные в PostgreSQL.

## Особенности

- **Два режима работы**: Live Mode (реальные данные с бирж) и Test Mode (синтетические данные)
- **Конкурентная обработка**: Использование паттернов Fan-in, Fan-out и Worker Pool
- **Кэширование**: Redis для быстрого доступа к последним ценам
- **Хранение данных**: PostgreSQL для долговременного хранения агрегированных данных
- **REST API**: Эндпоинты для получения информации о ценах и статистики
- **Graceful shutdown**: Корректное завершение работы при получении сигналов

## Архитектура

Проект следует гексагональной архитектуре:

- **Domain Layer**: Бизнес-логика и модели
- **Application Layer**: Use cases и оркестрация компонентов
- **Adapters Layer**: 
  - Web Adapter (HTTP handlers)
  - Storage Adapter (PostgreSQL)
  - Cache Adapter (Redis)
  - Exchange Adapter (источники данных)

## Запуск проекта

### Предварительные требования

- Docker и Docker Compose
- Go 1.21+ (для локальной разработки)

### Настройка окружения

1. Скопируйте файл `.env.example` в `.env`:
```bash
cp .env.example .env
```

2. Заполните переменные окружения в файле `.env`:

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

### Загрузка образов бирж

Запустите скрипт для загрузки образов бирж:

```bash
cd script
./load_images.sh
```

### Запуск с Docker Compose

```bash
docker-compose up -d
```

Приложение будет доступно на порту, указанном в `SERVER_PORT` (по умолчанию 8080).

### Локальная разработка

Для локальной разработки:

```bash
go build -o marketflow .
./marketflow
```

## Использование

### Командная строка

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

- `GET /prices/latest/{symbol}` - Последняя цена для символа
- `GET /prices/latest/{exchange}/{symbol}` - Последняя цена для символа с конкретной биржи
- `GET /prices/highest/{symbol}` - Наивысшая цена за период
- `GET /prices/highest/{exchange}/{symbol}` - Наивысшая цена за период с конкретной биржи
- `GET /prices/highest/{symbol}?period={duration}` - Наивысшая цена за указанный период
- `GET /prices/lowest/{symbol}` - Наинизшая цена за период
- `GET /prices/lowest/{exchange}/{symbol}` - Наинизшая цена за период с конкретной биржи
- `GET /prices/lowest/{symbol}?period={duration}` - Наинизшая цена за указанный период
- `GET /prices/average/{symbol}` - Средняя цена за период
- `GET /prices/average/{exchange}/{symbol}` - Средняя цена за период с конкретной биржи
- `GET /prices/average/{exchange}/{symbol}?period={duration}` - Средняя цена за указанный период с конкретной биржи

#### Data Mode API

- `POST /mode/test` - Переключиться в Test Mode
- `POST /mode/live` - Переключиться в Live Mode

#### System Health

- `GET /health` - Статус системы (соединения, доступность Redis)

## Структура проекта

```
marketflow/
├── cmd/                 # Точка входа приложения
├── internal/            # Внутренние пакеты
│   ├── adapters/        # Адаптеры
│   │   ├── exchange/    # Адаптеры бирж
│   │   ├── handlers/    # HTTP handlers
│   │   ├── postgres/    # PostgreSQL адаптер
│   │   └── redis/       # Redis адаптер
│   ├── application/     # Use cases и сервисы
│   │   ├── aggregator/  # Агрегация данных
│   │   ├── mode/        # Управление режимами
│   │   ├── usecase/     # Бизнес-логика
│   │   └── worker/      # Worker pool
│   ├── bootstrap/       # Инициализация приложения
│   ├── core/            # Основные утилиты
│   │   ├── apperrors/   # Ошибки приложения
│   │   ├── config/      # Конфигурация
│   │   └── utils/       # Вспомогательные утилиты
│   └── domain/          # Доменный слой
│       ├── models/      # Модели данных
│       └── ports/       # Интерфейсы
├── script/              # Скрипты
│   └── load_images.sh   # Скрипт загрузки образов бирж
└── docker-compose.yml   # Docker Compose конфигурация
```

## Обрабатываемые торговые пары

Система обрабатывает следующие криптовалютные пары:
- BTCUSDT
- DOGEUSDT
- TONUSDT
- SOLUSDT
- ETHUSDT

## База данных

В PostgreSQL создается таблица `birge_prices` со следующей структурой:

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

## Логирование

Используется пакет `log/slog` для логирования с различными уровнями:
- Info: Информационные сообщения
- Warning: Предупреждения
- Error: Ошибки

## Graceful Shutdown

Приложение корректно обрабатывает сигналы SIGINT и SIGTERM, обеспечивая clean shutdown всех компонентов.

## Мониторинг

Эндпоинт `/health` предоставляет информацию о состоянии системы, включая доступность Redis и соединений с биржами.