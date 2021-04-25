package dao

import (
	"github.com/pkg/errors"
	"golang-demo/error_demo/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.User{})
}

func CreateUser(name string, age uint8) error {
	if err := db.Create(&domain.User{
		Name: name,
		Age:  age,
	}).Error; err != nil {
		return errors.Wrap(err, "insert user error")
	}
	return nil
}

func FindUserById(id uint) (domain.NullUser, error) {
	user := &domain.User{}
	err := db.Model(&domain.User{}).Where("id = ?", id).First(user).Error
	if err == gorm.ErrRecordNotFound {
		return domain.NullUser{Valid: false, User: user}, nil
	}
	if err != nil {
		return domain.NullUser{}, err
	}
	return domain.NullUser{Valid: true, User: user}, err
}

func FindUsersByAge(age uint8) (*[]domain.User, error) {
	users := make([]domain.User, 0)
	err := db.Model(&domain.User{}).Where("age = ?", age).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return &users, nil
	} else if err != nil {
		return &users, errors.Wrap(err, "find users error")
	}
	return &users, err
}
