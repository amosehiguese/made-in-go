package store

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func RedisConn() (*redis.Client, error) {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPwd := os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}

	return redis.NewClient(&redis.Options{
		Addr: redisAddr,
		Password: redisPwd,
		DB: redisDB,
	}), nil
}
