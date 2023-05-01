package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/internal/entity"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/rediscache"
	"deall-alfon/pkg/util/timeutil"
	"github.com/go-redis/redis/v8"
)

type User struct {
	cfg    config.ConfigStr
	client *redis.ClusterClient
}

func NewUserCache(cfg config.ConfigStr, client *redis.ClusterClient) internal.UserCache {
	return &User{
		cfg:    cfg,
		client: client,
	}
}

func (u *User) GetUserDataByEmail(ctx context.Context, email string) (userData entity.AdminUser, err error) {
	key := fmt.Sprintf(constant.KeyUserDataByEmail, email)

	redisItf := u.client.Get(ctx, key)
	if err = redisItf.Err(); err != nil {
		if err == rediscache.RedisNotFound {
			return userData, nil
		}
		return userData, err
	}

	err = json.Unmarshal([]byte(redisItf.Val()), &userData)
	if err != nil {
		return userData, err
	}

	return userData, err
}

func (u *User) SetUserDataByEmail(ctx context.Context, userData entity.AdminUser) error {
	key := fmt.Sprintf(constant.KeyUserDataByEmail, userData.User.Email)
	data, err := json.Marshal(userData)
	if err != nil {
		return err
	}

	redisItf := u.client.SetEX(ctx, key, data, timeutil.TimeOneDay)
	if err = redisItf.Err(); err != nil {
		return err
	}

	return nil
}
