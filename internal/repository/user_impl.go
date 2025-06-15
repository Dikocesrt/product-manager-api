package repository

import (
	"product-manager-api/internal/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&entity.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) CreateUser(user *entity.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}