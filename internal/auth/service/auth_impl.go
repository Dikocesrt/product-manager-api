package service

import (
	uDomain "product-manager-api/internal/auth/domain"
	"product-manager-api/internal/jwt/service"
	"product-manager-api/internal/user/entity"
	uRepository "product-manager-api/internal/user/repository"
	"product-manager-api/pkg"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepository uRepository.UserRepository
}

func NewAuthService(userRepository uRepository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: userRepository,
	}
}

func (s *AuthServiceImpl) Register(request uDomain.RegisterRequest, jwtService service.JWTService) (uDomain.RegisterResponse, error) {
	if err := pkg.Validate(request); err != nil {
		return uDomain.RegisterResponse{}, err
	}

	existingUser, _ := s.userRepository.FindByEmail(request.Email)
	if existingUser != nil {
		return uDomain.RegisterResponse{}, pkg.ErrEmailAlreadyExists
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return uDomain.RegisterResponse{}, pkg.ErrInternalServerError
	}

	user := &entity.User{
		Email:    request.Email,
		Password: string(hashPassword),
	}

	if err := s.userRepository.CreateUser(user); err != nil {
		return uDomain.RegisterResponse{}, err
	}

	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		return uDomain.RegisterResponse{}, err
	}

	return uDomain.RegisterResponse{Token: token}, nil
}