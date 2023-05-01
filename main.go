package main

import (
	_ "deall-alfon/docs"
	"deall-alfon/internal/constant"
	"deall-alfon/internal/controller"
	internalmiddleware "deall-alfon/internal/middleware"
	"deall-alfon/internal/usecase"
	"deall-alfon/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	swagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	// init some dependencies
	cfg := config.GetConfig()
	ctrl := controller.GetControllers()
	mlwr := internalmiddleware.NewMiddleware(cfg, usecase.GetUseCases())

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())

	// documentation
	{
		d := e.Group("/doc")
		if cfg.Environment != constant.EnvironmentProduction {
			d.GET("/swagger/*", swagger.WrapHandler)
		}
	}
	// health check
	{
		s := e.Group("/status")
		s.GET("", ctrl.StatusHealthCheck)
	}

	// user
	{
		u := e.Group("/user")
		u.POST("/sign-up/admin", ctrl.SignUpUserAdmin, mlwr.VerifyJWTAccess())
		u.POST("/sign-up", ctrl.SignUpUser)
		u.POST("/log-in", ctrl.LoginUser)
	}

	// token
	{
		t := e.Group("/token")
		t.GET("/refresh-token", ctrl.GenerateNewPairToken)
	}

	e.Logger.Fatal(e.Start(":9090"))
}
