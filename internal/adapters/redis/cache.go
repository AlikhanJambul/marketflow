package redis

import (
	"context"
	"encoding/json"
	"fmt"
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

func (r *RedisCache) SetLatest(value models.Prices) error {
	ctx := context.Background()

	firstKey := fmt.Sprintf("latest/%s", value.Symbol)
	secondKey := fmt.Sprintf("latest/%s/%s", value.Exchange, value.Symbol)

	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.rbd.Set(ctx, firstKey, string(jsonValue), 10*time.Second).Err()
	if err != nil {
		return err
	}

	return r.rbd.Set(ctx, secondKey, string(jsonValue), 10*time.Second).Err()
}

func (r *RedisCache) GetLatest(key string) (models.Prices, error) {
	ctx := context.Background()

	value, err := r.rbd.Get(ctx, key).Result()
	if err != nil {
		return models.Prices{}, err
	}

	var result models.Prices

	if err = json.Unmarshal([]byte(value), &result); err != nil {
		return models.Prices{}, err
	}
	return result, nil
}

func (r *RedisCache) Check(ctx context.Context) error {
	return r.rbd.Ping(ctx).Err()
}
