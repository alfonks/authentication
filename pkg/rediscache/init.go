package rediscache

import (
	"sync"
	"time"

	"deall-alfon/pkg/config"
	"deall-alfon/pkg/util/converter"
	"github.com/go-redis/redis/v8"
)

var (
	redisClientOnce sync.Once
	redisClient     *redis.ClusterClient
)

func GetRedisCluster() *redis.ClusterClient {
	redisClientOnce.Do(func() {
		cfg := config.GetConfig()
		redisClient = connectRedis(redisConfig{
			URL:         []string{cfg.Redis.URL},
			MaxRetries:  converter.ToInt(cfg.Redis.MaxRetries),
			DialTimeout: time.Duration(cfg.Redis.DialTimeout) * time.Second,
			IdleTimeout: time.Duration(cfg.Redis.IdleTimeout) * time.Second,
		})
	})

	return redisClient
}
