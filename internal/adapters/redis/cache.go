package redis

import (
	"context"
	"encoding/json"
	"marketflow/internal/domain/models"
	"marketflow/internal/domain/ports"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	rbd *redis.Client
}

func NewRedisCache(rdb *redis.Client) ports.Cache {
	return &RedisCache{rbd: rdb}
}

func (r *RedisCache) Set(firstKey, secondKey string, value models.PriceStats) error {
	ctx := context.Background()

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.rbd.Set(ctx, firstKey, string(jsonValue), 30*time.Second).Err()
	if err != nil {
		return err
	}

	return r.rbd.Set(ctx, secondKey, string(jsonValue), 30*time.Second).Err()
}

func (r *RedisCache) Get(s string) (models.PriceStats, error) {
	ctx := context.Background()

	value, err := r.rbd.Get(ctx, s).Result()
	if err != nil {
		return models.PriceStats{}, err
	}

	var result models.PriceStats

	if err = json.Unmarshal([]byte(value), &result); err != nil {
		return models.PriceStats{}, err
	}

	return result, nil
}

func (r *RedisCache) SetLatest(firstKey, secondKey string, latest models.LatestPrice) error {
	ctx := context.Background()

	jsonValue, err := json.Marshal(latest)
	if err != nil {
		return err
	}

	err = r.rbd.Set(ctx, firstKey, string(jsonValue), 30*time.Second).Err()
	if err != nil {
		return err
	}

	return r.rbd.Set(ctx, secondKey, string(jsonValue), 30*time.Second).Err()
}

func (r *RedisCache) GetLatest(key string) (models.LatestPrice, error) {
	ctx := context.Background()

	value, err := r.rbd.Get(ctx, key).Result()
	if err != nil {
		return models.LatestPrice{}, err
	}

	var result models.LatestPrice

	if err = json.Unmarshal([]byte(value), &result); err != nil {
		return models.LatestPrice{}, err
	}
	return result, nil
}

func (r *RedisCache) Check(ctx context.Context) error {
	return r.rbd.Ping(ctx).Err()
}
