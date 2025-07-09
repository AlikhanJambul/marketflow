package redis

import (
	"github.com/redis/go-redis/v9"
	"marketflow/internal/domain/models"
)

func ConnRedis(redisData models.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisData.Host + ":" + redisData.Port,
		Password: redisData.Password,
		DB:       redisData.DB,
	})

	//err := rdb.Set(ctx, "mykey", "Hello, Redis!", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := rdb.Get(ctx, "mykey").Result()
	//if err != nil {
	//	panic(err)
	//}

	return rdb
}
