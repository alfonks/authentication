package facade

import (
	"deall-alfon/internal/usecase"
	"sync"

	"deall-alfon/internal"
	"deall-alfon/pkg/config"
)

var (
	userFacadeOnce sync.Once
	userFacade     internal.UserFacade

	tokenFacadeOnce sync.Once
	tokenFacade     internal.TokenFacade
)

func GetUserFacade() internal.UserFacade {
	userFacadeOnce.Do(func() {
		userFacade = NewUserFacade(
			config.GetConfig(),
			usecase.GetUseCases(),
		)
	})

	return userFacade
}

func GetTokenFacade() internal.TokenFacade {
	tokenFacadeOnce.Do(func() {
		tokenFacade = NewTokenFacade(
			config.GetConfig(),
			usecase.GetUseCases(),
		)
	})

	return tokenFacade
}
