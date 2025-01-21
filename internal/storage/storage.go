package storage

import "github.com/AnkitBishen/courseHub/internal/stype"

type Storage interface {
	UserRegister(user stype.UserRegister) (int64, error)
	UserLoginValidation(user stype.UserRegister) (bool, error)
	UserValidateFromOauth(user stype.UserRegister) (bool, error)
}
