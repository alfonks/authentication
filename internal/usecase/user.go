package usecase

import (
	"context"
	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/internal/entity"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/encryption"
	"deall-alfon/pkg/errs"
	"deall-alfon/pkg/util/fn"
	"deall-alfon/pkg/util/logger"
	"time"
)

type User struct {
	cfg       config.ConfigStr
	userCache internal.UserCache
	userDB    internal.UserDB
	enc       encryption.EncItf
}

func NewUserUseCase(
	cfg config.ConfigStr,
	userCache internal.UserCache,
	userDB internal.UserDB,
	enc encryption.EncItf,
) internal.UserUseCase {
	return &User{
		cfg:       cfg,
		userDB:    userDB,
		userCache: userCache,
		enc:       enc,
	}
}

func (u *User) SignUpUser(ctx context.Context, user entity.AdminUser) (msg string, err error) {
	user.User.Password, err = u.enc.HashSalt([]byte(user.User.Password))
	if err != nil {
		return "", errs.SetUserError(err, constant.SignUpServerError)
	}
	user.CreateTime = time.Now()

	err = u.userDB.SignUpUser(ctx, user)
	if err != nil {
		return "", err
	}

	msg = constant.SignUpSuccess

	return msg, err
}

func (u *User) GetUserDataByEmail(ctx context.Context, email string) (userData entity.AdminUser, err error) {
	op := fn.Name()

	userData, err = u.userCache.GetUserDataByEmail(ctx, email)
	if err != nil {
		logger.Printf("[%v] fail to get user data from cache for email: %v, error: %v", op, email, err)
	}

	if !userData.IsEmpty() {
		return userData, err
	}

	userData, err = u.userDB.GetUserByEmail(ctx, email)
	if err != nil {
		return userData, err
	}

	err = u.userCache.SetUserDataByEmail(ctx, userData)
	if err != nil {
		logger.Printf("[%v] fail to set user data to cache for email: %v, error: %v", op, email, err)
	}

	return userData, err
}

func (u *User) CheckUserCredential(ctx context.Context, userRequest entity.AdminUser, userData entity.AdminUser) (valid bool, err error) {
	valid, err = u.enc.ValidateHashData([]byte(userRequest.User.Password), []byte(userData.User.Password))
	if err != nil {
		err = errs.SetUserError(err, constant.LoginInternalServerError)
	}
	if !valid {
		err = errs.SetUserError(err, constant.LoginUserDataInvalidOrNotFound)
	}

	return valid, err
}
