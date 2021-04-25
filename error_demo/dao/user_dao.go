package dao

import (
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model `json:"gorm_model,omitempty" :"gorm_model" :"gorm_model"`
	Name       string `json:"name,omitempty,unique" :"name" :"name"`
	Age        uint8  `:"age"`
}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}

func CreateUser(name string, age uint8) error {
	if err := db.Create(&User{
		Name: name,
		Age:  age,
	}).Error; err != nil {
		return errors.Wrap(err, "create user error")
	}
	return nil
}

func FindUserById(id uint) (*User, error) {
	user := &User{}
	err := db.Model(&User{}).Where("id = ?", id).First(user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

func FindUsersByAge(age uint8) (*[]User, error) {
	users := make([]User, 0)
	err := db.Model(&User{}).Where("age = ?", age).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return &users, nil
	}
	return &users, err
}
