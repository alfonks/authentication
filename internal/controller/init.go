package controller

import (
	"sync"

	"deall-alfon/internal"
	"deall-alfon/internal/facade"
	"deall-alfon/pkg/config"
)

type Controllers struct {
	internal.HealthCheckController
	internal.UserController
	internal.TokenController
}

var (
	controllersServiceOnce sync.Once
	controllersService     Controllers

	healthCheckControllerOnce sync.Once
	healthCheckController     internal.HealthCheckController

	userControllerOnce sync.Once
	userController     internal.UserController

	tokenControllerOnce sync.Once
	tokenController     internal.TokenController
)

func GetControllers() Controllers {
	controllersServiceOnce.Do(func() {
		controllersService = Controllers{
			getHealthCheckController(),
			getUserController(),
			getTokenController(),
		}
	})

	return controllersService
}

func getHealthCheckController() internal.HealthCheckController {
	healthCheckControllerOnce.Do(func() {
		healthCheckController = NewHealthCheckController()
	})

	return healthCheckController
}

func getUserController() internal.UserController {
	userControllerOnce.Do(func() {
		userController = NewUserController(
			config.GetConfig(),
			facade.GetUserFacade(),
		)
	})

	return userController
}

func getTokenController() internal.TokenController {
	tokenControllerOnce.Do(func() {
		tokenController = NewTokenController(
			config.GetConfig(),
			facade.GetTokenFacade(),
		)
	})

	return tokenController
}
