package cache

import (
	"sync"

	"deall-alfon/internal"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/rediscache"
)

var (
	userCacheOnce sync.Once
	userCache     internal.UserCache

	tokenCacheOnce sync.Once
	tokenCache     internal.TokenCache
)

func GetUserCache() internal.UserCache {
	userCacheOnce.Do(func() {
		userCache = NewUserCache(
			config.GetConfig(),
			rediscache.GetRedisCluster(),
		)
	})

	return userCache
}

func GetTokenCache() internal.TokenCache {
	tokenCacheOnce.Do(func() {
		tokenCache = NewTokenCache(
			config.GetConfig(),
			rediscache.GetRedisCluster(),
		)
	})

	return tokenCache
}
