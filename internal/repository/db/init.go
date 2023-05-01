package db

import (
	"sync"

	"deall-alfon/internal"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/mongodb"
)

var (
	userDBOnce sync.Once
	userDB     internal.UserDB
)

func GetUserDB() internal.UserDB {
	userDBOnce.Do(func() {
		userDB = NewUserDB(
			config.GetConfig(),
			mongodb.GetMongoDBinstance(),
		)
	})

	return userDB
}
