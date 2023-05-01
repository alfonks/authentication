package internal

import (
	"context"
	"deall-alfon/internal/entity"
	"github.com/labstack/echo/v4"
)

//go:generate mockgen -destination=../mocks/mock_internal/mock_token.go -source=token.go
type TokenController interface {
	GenerateNewPairToken(c echo.Context) error
}

type TokenFacade interface {
	GenerateNewPairToken(ctx context.Context, refreshToken string) (token entity.Tokens, err error)
}

type TokenUseCase interface {
	GenerateTokenPair(ctx context.Context, user entity.AdminUser) (token entity.Tokens, err error)
	DeleteCacheTokenPair(ctx context.Context, email string) error
}

type TokenCache interface {
	GetGeneratedTokenPair(ctx context.Context, email string) (token entity.Tokens, err error)
	SetGeneratedTokenPair(ctx context.Context, email string, token entity.Tokens) error
	DeleteCacheTokenPair(ctx context.Context, email string) error
}
