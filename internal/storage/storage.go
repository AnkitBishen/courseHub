package storage

import "github.com/AnkitBishen/courseHub/internal/stype"

type Storage interface {
	UserRegister(user stype.UserRegister) error
}
