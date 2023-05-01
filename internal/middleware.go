package internal

import "github.com/labstack/echo/v4"

type MiddlewareItf interface {
	VerifyJWTAccess() echo.MiddlewareFunc
}
