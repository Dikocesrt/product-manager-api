package repository

import (
	"context"
	"encoding/json"
	"product-manager-api/config"
	"product-manager-api/internal/user/entity"
	"time"

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
	// Use project prefix for Redis key to avoid collisions
	ctx := context.Background()
	cacheKey := "product-manager:user:email:" + email

	// Check if user exists in cache
	userJSON, err := config.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		// User found in cache
		var user entity.User
		if err := json.Unmarshal([]byte(userJSON), &user); err == nil {
			return &user, nil
		}
	}

	// User not in cache, query database
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// Store in cache for future requests (cache for 15 minutes)
	if userJSON, err := json.Marshal(user); err == nil {
		config.Redis.Set(ctx, cacheKey, userJSON, 15*time.Minute)
	}

	return &user, nil
}

func (r *UserRepositoryImpl) CreateUser(user *entity.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	// Use project prefix for Redis key to avoid collisions
	ctx := context.Background()
	cacheKey := "product-manager:user:email:" + user.Email

	if userJSON, err := json.Marshal(user); err == nil {
		config.Redis.Set(ctx, cacheKey, userJSON, 15*time.Minute)
	}

	return nil
}