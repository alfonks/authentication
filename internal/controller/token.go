package controller

import (
	"context"
	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/pkg/api"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/errs"
	"deall-alfon/pkg/util/converter"
	"deall-alfon/pkg/util/fn"
	"deall-alfon/pkg/util/logger"
	"github.com/labstack/echo/v4"
)

type Token struct {
	cfg    config.ConfigStr
	facade internal.TokenFacade
}

func NewTokenController(cfg config.ConfigStr, facade internal.TokenFacade) internal.TokenController {
	return &Token{
		cfg:    cfg,
		facade: facade,
	}
}

func (t *Token) GenerateNewPairToken(c echo.Context) error {
	op := fn.Name()
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	req := c.Request()

	res := api.DefaultResponse(c)

	refreshToken := req.Header[constant.HeaderJWTRefreshToken]
	if len(refreshToken) <= 0 {
		logger.Printf("[%v] refresh token not found", op)
		return res.ResponseBadRequestError("refresh token not found")
	}

	realRefreshToken := converter.ToString(refreshToken[0])

	newTokenPair, err := t.facade.GenerateNewPairToken(ctx, realRefreshToken)
	if err != nil {
		logger.Printf("[%v] fail generate new token pair", op)
		return res.ResponseInternalServerError(nil, errs.GetUserError(err))
	}

	return res.ReponseOK(newTokenPair, constant.GenerateNewTokenSuccess)
}
