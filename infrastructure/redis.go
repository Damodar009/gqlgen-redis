package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// Database modal
type RedisClient struct {
	RDB *redis.Client
}

// NewDatabase creates a new database instance
func NewRedisClient() RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	var ctx = context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis:", pong)
	return RedisClient{
		RDB: rdb,
	}
}
