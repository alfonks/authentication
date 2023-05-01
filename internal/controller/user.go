package controller

import (
	"context"
	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/internal/entity"
	"deall-alfon/pkg/api"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/errs"
	"deall-alfon/pkg/util/fn"
	"deall-alfon/pkg/util/logger"
	"github.com/labstack/echo/v4"
)

type User struct {
	cfg        config.ConfigStr
	userFacade internal.UserFacade
}

func NewUserController(cfg config.ConfigStr, userFacade internal.UserFacade) internal.UserController {
	return &User{
		cfg:        cfg,
		userFacade: userFacade,
	}
}

func (u *User) SignUpUser(c echo.Context) error {
	op := fn.Name()
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	var (
		req entity.User
		res = api.DefaultResponse(c)
	)

	err := api.Bind(c, &req)
	if err != nil {
		return res.ResponseBadRequestError(err.Error())
	}

	msg, err := u.userFacade.SignUpUser(ctx, req)
	if err != nil {
		logger.Printf("[%v] fail to insert user data with error: %v, user email: %v", op, err.Error(), req.Email)
		return res.ResponseInternalServerError(nil, errs.GetUserError(err))
	}

	return res.ReponseOK(nil, msg)
}

func (u *User) SignUpUserAdmin(c echo.Context) error {
	op := fn.Name()
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	var (
		req entity.User
		res = api.DefaultResponse(c)
	)

	err := api.Bind(c, &req)
	if err != nil {
		return res.ResponseBadRequestError(err.Error())
	}

	msg, err := u.userFacade.SignUpUserAdmin(ctx, req)
	if err != nil {
		logger.Printf("[%v] fail to insert user data with error: %v, user email: %v", op, err.Error(), req.Email)
		return res.ResponseInternalServerError(nil, errs.GetUserError(err))
	}

	return res.ReponseOK(nil, msg)
}

func (u *User) LoginUser(c echo.Context) error {
	op := fn.Name()
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	var (
		req entity.User
		res = api.DefaultResponse(c)
	)

	err := api.Bind(c, &req)
	if err != nil {
		return res.ResponseBadRequestError(err.Error())
	}

	token, err := u.userFacade.LoginUser(ctx, req)
	if err != nil {
		logger.Printf("[%v] error login with user email: %v, error: %v", op, req.Email, err.Error())
		return res.ResponseInternalServerError(nil, errs.GetUserError(err))
	}

	return res.ReponseOK(token, constant.LoginSuccess)
}
