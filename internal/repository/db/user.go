package db

import (
	"context"

	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/internal/entity"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/errs"
	"deall-alfon/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	cfg      config.ConfigStr
	instance mongodb.MongoDBInstance
}

func NewUserDB(cfg config.ConfigStr, instance mongodb.MongoDBInstance) internal.UserDB {
	return &User{
		cfg:      cfg,
		instance: instance,
	}
}

func (u *User) GetUserByEmail(ctx context.Context, email string) (entity.AdminUser, error) {
	filter := bson.D{
		{
			"user.email", email,
		},
	}

	var res entity.AdminUser

	err := u.
		instance.
		Slave.
		Collection(mongodb.DeallUser).
		FindOne(ctx, filter).
		Decode(&res)

	if err != nil {
		return res, errs.SetUserError(err, constant.LoginUserDataInvalidOrNotFound)
	}

	return res, nil
}

func (u *User) SignUpUser(ctx context.Context, user entity.AdminUser) error {
	_, err := u.
		instance.
		Master.
		Collection(mongodb.DeallUser).
		InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errs.SetUserError(err, constant.SignUpDuplicateEmailError)
		}
		return err
	}

	return nil
}
