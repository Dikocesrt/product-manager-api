package service

import (
	"product-manager-api/internal/auth/domain"
	"product-manager-api/internal/jwt/service"
)

type AuthService interface {
	Register(request domain.Request, jwtService service.JWTService) (domain.Response, error)
	Login(request domain.Request, jwtService service.JWTService) (domain.Response, error)
}