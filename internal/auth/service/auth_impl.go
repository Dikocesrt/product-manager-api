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

func (s *AuthServiceImpl) Register(request uDomain.Request, jwtService service.JWTService) (uDomain.Response, error) {
	if err := pkg.Validate(request); err != nil {
		return uDomain.Response{}, err
	}

	existingUser, _ := s.userRepository.FindByEmail(request.Email)
	if existingUser != nil {
		return uDomain.Response{}, pkg.ErrEmailAlreadyExists
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return uDomain.Response{}, pkg.ErrInternalServerError
	}

	user := &entity.User{
		Email:    request.Email,
		Password: string(hashPassword),
	}

	if err := s.userRepository.CreateUser(user); err != nil {
		return uDomain.Response{}, err
	}

	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		return uDomain.Response{}, err
	}

	return uDomain.Response{Token: token}, nil
}

func (s *AuthServiceImpl) Login(request uDomain.Request, jwtService service.JWTService) (uDomain.Response, error) {
	if err := pkg.Validate(request); err != nil {
		return uDomain.Response{}, err
	}

	user, err := s.userRepository.FindByEmail(request.Email)
	if err != nil {
		return uDomain.Response{}, pkg.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return uDomain.Response{}, pkg.ErrInvalidCredentials
	}

	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		return uDomain.Response{}, err
	}

	return uDomain.Response{Token: token}, nil
}