package user_dao

import  "gorm.io/gorm"

type User struct {
	gorm.Model `json:"gorm_model,omitempty"`
	Name       string `json:"name,omitempty"`
}



