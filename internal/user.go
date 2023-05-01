package internal

import (
	"context"

	"deall-alfon/internal/entity"
	"github.com/labstack/echo/v4"
)

//go:generate mockgen -destination=../mocks/mock_internal/mock_user.go -source=user.go

type UserController interface {
	SignUpUser(c echo.Context) error
	SignUpUserAdmin(c echo.Context) error
	LoginUser(c echo.Context) error
}

type UserFacade interface {
	SignUpUser(ctx context.Context, user entity.User) (msg string, err error)
	SignUpUserAdmin(ctx context.Context, user entity.User) (msg string, err error)
	LoginUser(ctx context.Context, user entity.User) (token entity.Tokens, err error)
}

type UserUseCase interface {
	SignUpUser(ctx context.Context, user entity.AdminUser) (msg string, err error)
	GetUserDataByEmail(ctx context.Context, email string) (userData entity.AdminUser, err error)
	CheckUserCredential(ctx context.Context, userRequest entity.AdminUser, userData entity.AdminUser) (valid bool, err error)
}

type UserCache interface {
	GetUserDataByEmail(ctx context.Context, email string) (userData entity.AdminUser, err error)
	SetUserDataByEmail(ctx context.Context, userData entity.AdminUser) error
}

type UserDB interface {
	GetUserByEmail(ctx context.Context, email string) (entity.AdminUser, error)
	SignUpUser(ctx context.Context, user entity.AdminUser) error
}
