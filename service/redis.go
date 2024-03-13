package service

import (
	"github.com/redis/go-redis/v9"
)

func ConnToRedis() (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		PoolSize: 20,
	})

	return
}
