package facade

import (
	"context"
	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/internal/entity"
	"deall-alfon/internal/usecase"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/errs"
)

type User struct {
	cfg      config.ConfigStr
	useCases usecase.UseCases
}

func NewUserFacade(cfg config.ConfigStr, useCases usecase.UseCases) internal.UserFacade {
	return &User{
		cfg:      cfg,
		useCases: useCases,
	}
}

func (u *User) SignUpUser(ctx context.Context, user entity.User) (msg string, err error) {
	adminUser := entity.AdminUser{
		UserLevel: constant.UserLevelNormal,
		User:      user,
	}
	msg, err = u.useCases.UserUC.SignUpUser(ctx, adminUser)

	return msg, err
}

func (u *User) SignUpUserAdmin(ctx context.Context, user entity.User) (msg string, err error) {
	adminUser := entity.AdminUser{
		UserLevel: constant.UserLevelAdmin,
		User:      user,
	}
	msg, err = u.useCases.UserUC.SignUpUser(ctx, adminUser)

	return msg, err
}

func (u *User) LoginUser(ctx context.Context, user entity.User) (token entity.Tokens, err error) {
	userRequest := entity.AdminUser{User: user}
	userData, err := u.useCases.UserUC.GetUserDataByEmail(ctx, user.Email)
	if err != nil {
		return token, err
	}

	valid, err := u.useCases.UserUC.CheckUserCredential(ctx, userRequest, userData)
	if !valid {
		return token, err
	}

	token, err = u.useCases.TokenUC.GenerateTokenPair(ctx, userData)
	if err != nil {
		return token, errs.SetUserError(err, constant.GenerateTokenFail)
	}

	return token, err
}
