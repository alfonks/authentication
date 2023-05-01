package rediscache

import (
	"github.com/go-redis/redis/v8"
)

func connectRedis(cfg redisConfig) *redis.ClusterClient {
	redisClusterConfig := &redis.ClusterOptions{
		Addrs:       cfg.URL,
		MaxRetries:  cfg.MaxRetries,
		DialTimeout: cfg.DialTimeout,
		IdleTimeout: cfg.IdleTimeout,
	}

	client := redis.NewClusterClient(redisClusterConfig)
	return client
}
