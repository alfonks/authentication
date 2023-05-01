package usecase

import (
	"context"
	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/internal/entity"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/util/fn"
	"deall-alfon/pkg/util/logger"
	"deall-alfon/pkg/util/timeutil"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Token struct {
	cfg   config.ConfigStr
	cache internal.TokenCache
}

func NewTokenUseCase(cfg config.ConfigStr, cache internal.TokenCache) internal.TokenUseCase {
	return &Token{
		cfg:   cfg,
		cache: cache,
	}
}

func (t *Token) GenerateTokenPair(ctx context.Context, userData entity.AdminUser) (token entity.Tokens, err error) {
	op := fn.Name()

	token, err = t.cache.GetGeneratedTokenPair(ctx, userData.User.Email)
	if err != nil {
		logger.Printf("[%v] fail to get token pair for email: %v, error: %v", op, userData.User.Email, err)
	}

	jwtToken := jwt.New(jwt.SigningMethodHS512)
	jwtTokenExpiry := time.Now().Add(time.Minute * time.Duration(t.cfg.JWT.ExpiryTime))

	claims := jwtToken.Claims.(jwt.MapClaims)
	claims[constant.KeyTokenUserLevel] = userData.UserLevel
	claims[constant.KeyTokenEmail] = userData.User.Email
	claims[constant.KeytokenExpiryTime] = jwtTokenExpiry.Unix()

	accessToken, err := jwtToken.SignedString([]byte(t.cfg.JWT.SecretKey))
	if err != nil {
		return token, err
	}

	refreshJWT := jwt.New(jwt.SigningMethodHS512)
	rtClaims := refreshJWT.Claims.(jwt.MapClaims)
	rtClaims[constant.KeyTokenEmail] = userData.User.Email
	rtClaims[constant.KeytokenExpiryTime] = time.Now().Add(timeutil.TimeOneDay * time.Duration(t.cfg.JWT.RefreshTokenExpiryTime)).Unix()

	refreshToken, err := refreshJWT.SignedString([]byte(t.cfg.JWT.SecretKey))
	if err != nil {
		return token, err
	}

	token = entity.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	err = t.cache.SetGeneratedTokenPair(ctx, userData.User.Email, token)
	if err != nil {
		logger.Printf("[%v] fail to set cache token for email: %v, error: %v", op, userData.User.Email, err)
	}

	return token, err
}

func (t *Token) DeleteCacheTokenPair(ctx context.Context, email string) error {
	return t.cache.DeleteCacheTokenPair(ctx, email)
}
