package internal

import "github.com/labstack/echo/v4"

//go:generate mockgen -destination=../mocks/mock_internal/mock_health_check.go -source=healthcheck.go
type HealthCheckController interface {
	StatusHealthCheck(c echo.Context) error
}
