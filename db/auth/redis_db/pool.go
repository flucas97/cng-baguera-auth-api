package redis_db

import (
	"context"

	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	redis "github.com/go-redis/redis/v8"
)

var (
	ctx    = context.Background()
	Client = redis.NewClient(&redis.Options{
		Addr:       "cachedb:6379",
		Password:   "Root!2020",
		PoolSize:   500,
		MaxRetries: 2,
		DB:         0,
	})
)

func init() {
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		logger.Error("cannot ping redis", err)
		panic(err)
	}

	logger.Info("redis successfuly connected")
}
