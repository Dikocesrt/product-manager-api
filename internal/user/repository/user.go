package repository

import "product-manager-api/internal/user/entity"


type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
}