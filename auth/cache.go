package auth

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

func InitRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	if redisHost == "" {
		redisHost = "localhost"
	}
	if redisPort == "" {
		redisPort = "6379"
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
	})
}

// CacheToken сохраняет токен в Redis с TTL 1 час
func CacheToken(token string) error {
	if redisClient == nil {
		InitRedis()
	}
	return redisClient.Set(ctx, token, "1", time.Hour).Err()
}

// IsTokenCached проверяет, есть ли токен в Redis
func IsTokenCached(token string) bool {
	if redisClient == nil {
		InitRedis()
	}
	res, err := redisClient.Get(ctx, token).Result()
	return err == nil && res == "1"
}
