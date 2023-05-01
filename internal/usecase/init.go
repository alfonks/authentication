package usecase

import (
	"deall-alfon/pkg/encryption"
	"sync"

	"deall-alfon/internal"
	"deall-alfon/internal/repository/cache"
	"deall-alfon/internal/repository/db"
	"deall-alfon/pkg/config"
)

type UseCases struct {
	UserUC  internal.UserUseCase
	TokenUC internal.TokenUseCase
}

var (
	allUseCasesOnce sync.Once
	allUseCases     UseCases

	userUseCaseOnce sync.Once
	userUseCase     internal.UserUseCase

	tokenUseCaseOnce sync.Once
	tokenUseCase     internal.TokenUseCase
)

func GetUseCases() UseCases {
	allUseCasesOnce.Do(func() {
		allUseCases = UseCases{
			UserUC:  getUserUseCase(),
			TokenUC: getTokenUseCase(),
		}
	})
	return allUseCases
}

func getUserUseCase() internal.UserUseCase {
	userUseCaseOnce.Do(func() {
		userUseCase = NewUserUseCase(
			config.GetConfig(),
			cache.GetUserCache(),
			db.GetUserDB(),
			encryption.GetEncryption(),
		)
	})

	return userUseCase
}

func getTokenUseCase() internal.TokenUseCase {
	tokenUseCaseOnce.Do(func() {
		tokenUseCase = NewTokenUseCase(
			config.GetConfig(),
			cache.GetTokenCache(),
		)
	})

	return tokenUseCase
}
