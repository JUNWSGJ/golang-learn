package service

import (
	"github.com/pkg/errors"
	"golang-demo/error_demo/dao"
)

func CreateUser(name string, age uint8) error {
	if err := dao.CreateUser(name, age); err != nil {
		return errors.Wrap(err, "create user fail")
	}
	return nil
}

func FindUserById(id uint) (*dao.User, error) {
	if user, err := dao.FindUserById(id); err != nil {
		return nil, errors.WithMessage(err, "find users by id error")
	} else {
		return user, nil
	}
}

func FindUsersByAge(age uint8) (*[]dao.User, error) {
	if users, err := dao.FindUsersByAge(age); err != nil {
		return nil, errors.WithMessage(err, "find users by id error")
	} else {
		return users, nil
	}
}
