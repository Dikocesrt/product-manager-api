package repository

import "product-manager-api/internal/entity"

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
}