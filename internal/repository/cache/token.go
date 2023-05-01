package cache

import (
	"context"
	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/internal/entity"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/rediscache"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Token struct {
	cfg    config.ConfigStr
	client *redis.ClusterClient
}

func NewTokenCache(cfg config.ConfigStr, client *redis.ClusterClient) internal.TokenCache {
	return &Token{
		cfg:    cfg,
		client: client,
	}
}

func (t *Token) GetGeneratedTokenPair(ctx context.Context, email string) (token entity.Tokens, err error) {
	key := fmt.Sprintf(constant.KeyTokenKeyPair, email)

	redisItf := t.client.Get(ctx, key)
	if err = redisItf.Err(); err != nil {
		if err == rediscache.RedisNotFound {
			return token, nil
		}
		return token, err
	}

	err = json.Unmarshal([]byte(redisItf.Val()), &token)
	if err != nil {
		return token, err
	}

	return token, err
}

func (t *Token) SetGeneratedTokenPair(ctx context.Context, email string, token entity.Tokens) error {
	key := fmt.Sprintf(constant.KeyTokenKeyPair, email)
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}

	redisItf := t.client.SetEX(ctx, key, data, time.Minute*time.Duration(t.cfg.JWT.ExpiryTime-1))
	if err = redisItf.Err(); err != nil {
		return err
	}

	return nil
}

func (t *Token) DeleteCacheTokenPair(ctx context.Context, email string) error {
	key := fmt.Sprintf(constant.KeyTokenKeyPair, email)

	redisItf := t.client.Del(ctx, key)
	if err := redisItf.Err(); err != nil {
		return err
	}

	return nil
}
