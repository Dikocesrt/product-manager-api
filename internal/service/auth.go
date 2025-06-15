package service

import (
	"product-manager-api/internal/domain"
)

type AuthService interface {
	Register(request domain.RegisterRequest, jwtService JWTService) (domain.RegisterResponse, error)
}