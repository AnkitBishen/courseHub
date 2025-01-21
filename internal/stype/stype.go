package stype

import "time"

type UserRegister struct {
	Id           string    `json:"id"`
	FirstName    string    `json:"name" validate:"required"`
	LastName     string    `json:"last_name" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password" validate:"required,min=6"`
	Picture      string    `json:"picture"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime   time.Time `json:"update_time"`
}
