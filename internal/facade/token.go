package facade

import (
	"context"
	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/internal/entity"
	"deall-alfon/internal/usecase"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/errs"
	"deall-alfon/pkg/util/converter"
	"deall-alfon/pkg/util/fn"
	"deall-alfon/pkg/util/logger"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	cfg      config.ConfigStr
	usecases usecase.UseCases
}

func NewTokenFacade(cfg config.ConfigStr, usecases usecase.UseCases) internal.TokenFacade {
	return &Token{
		cfg:      cfg,
		usecases: usecases,
	}
}

func (t *Token) GenerateNewPairToken(ctx context.Context, refreshToken string) (token entity.Tokens, err error) {
	op := fn.Name()

	jwttoken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.SetUserError(fmt.Errorf("wrong jwt algorithm"), constant.GenerateNewtokenPairRefreshTokenInvalid)
		}

		return []byte(t.cfg.JWT.SecretKey), nil
	})

	jwtclaimmap, ok := jwttoken.Claims.(jwt.MapClaims)
	if !ok {
		return token, errs.SetUserError(fmt.Errorf("invalid jwt claim"), constant.GenerateNewtokenPairRefreshTokenInvalid)
	}

	email := converter.ToString(jwtclaimmap[constant.KeyTokenEmail])

	if err != nil {
		logger.Printf("[%v] fail parse refresh token for email: %v, error: %v", op, email, err)
		return token, errs.SetUserError(err, constant.GenerateNewtokenPairRefreshTokenInvalid)
	}

	if !jwttoken.Valid {
		return token, errs.SetUserError(fmt.Errorf("invalid token"), constant.GenerateNewTokenPairRefreshTokenExpired)
	}

	userData, err := t.usecases.UserUC.GetUserDataByEmail(ctx, email)
	if err != nil {
		logger.Printf("[%v] fail to get user data for email: %v, error: %v", op, email, err)
		return token, errs.SetUserError(err, constant.GenerateNewtokenPairRefreshTokenInvalid)
	}

	err = t.usecases.TokenUC.DeleteCacheTokenPair(ctx, email)
	if err != nil {
		logger.Printf("[%v] fail to delete token pair from cache for email: %v, error: %v", op, email, err)
		return token, errs.SetUserError(err, constant.GenerateTokenFail)
	}

	token, err = t.usecases.TokenUC.GenerateTokenPair(ctx, userData)
	if err != nil {
		return token, errs.SetUserError(err, constant.GenerateTokenFail)
	}

	return token, err
}
