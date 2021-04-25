package domain

import (
	"errors"
	"gorm.io/gorm"
)

var UserNotFound = errors.New("UserNotFound")

type User struct {
	gorm.Model `json:"gorm_model,omitempty"`
	Name       string `json:"name,omitempty", gorm:"unique_index"`
	Age        uint8
}

type NullUser struct {
	User  *User
	Valid bool
}
