package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"marketflow/internal/core/ports"
	"marketflow/internal/domain/models"
	"time"
)

type RedisCache struct {
	rbd *redis.Client
}

func NewRedisCache(rdb *redis.Client) ports.Cache {
	return &RedisCache{rbd: rdb}
}

func (r *RedisCache) Set(key string, value models.Prices) error {
	ctx := context.Background()

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.rbd.Set(ctx, key, string(jsonValue), 60*time.Second).Err()
}

func (r *RedisCache) Get(s string) (models.Prices, error) {
	ctx := context.Background()

	value, err := r.rbd.Get(ctx, s).Result()
	if err != nil {
		return models.Prices{}, err
	}

	var result models.Prices

	if err = json.Unmarshal([]byte(value), &result); err != nil {
		return models.Prices{}, err
	}

	return result, nil
}
