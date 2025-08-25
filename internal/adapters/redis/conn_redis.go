package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
	"marketflow/internal/domain/models"
)

func ConnRedis(redisData models.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         redisData.Host + ":" + redisData.Port,
		Password:     redisData.Password,
		DB:           redisData.DB,
		DialTimeout:  100 * time.Millisecond,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,
	})

	return rdb
}
