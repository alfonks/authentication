package entity

import (
	"time"
)

type User struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	Gender    string `json:"gender" bson:"gender"`

	BirthDate time.Time `json:"birth_date" bson:"birth_date"`
}

type AdminUser struct {
	ID         string    `bson:"_id,omitempty"`
	UserLevel  int64     `bson:"user_level"`
	CreateTime time.Time `bson:"create_time"`
	User       User      `bson:"user"`
}

func (a *AdminUser) IsEmpty() bool {
	return a.ID == ""
}
