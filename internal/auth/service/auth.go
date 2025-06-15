package service

import (
	"product-manager-api/internal/auth/domain"
	"product-manager-api/internal/jwt/service"
)

type AuthService interface {
	Register(request domain.RegisterRequest, jwtService service.JWTService) (domain.RegisterResponse, error)
}