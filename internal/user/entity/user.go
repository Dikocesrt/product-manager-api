package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

func (User) TableName() string {
	return "users"
}

func ToUserModel(user User) User {
	return User{
		Model:    user.Model,
		Email:    user.Email,
		Password: user.Password,
	}
}