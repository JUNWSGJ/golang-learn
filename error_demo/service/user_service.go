package service

import (
	"github.com/pkg/errors"
	"golang-demo/error_demo/dao"
	"golang-demo/error_demo/domain"
)

func CreateUser(name string, age uint8) error {
	if err := dao.CreateUser(name, age); err != nil {
		return errors.WithMessage(err, "create user fail")
	}
	return nil
}

func FindUserById(id uint) (domain.NullUser, error) {
	if user, err := dao.FindUserById(id); err != nil {
		return user, errors.WithMessage(err, "find user by id error")
	} else {
		return user, err
	}

}

func FindUsersByAge(age uint8) (*[]domain.User, error) {
	if users, err := dao.FindUsersByAge(age); err != nil {
		return nil, errors.WithMessage(err, "find users by id error")
	} else {
		return users, nil
	}
}
