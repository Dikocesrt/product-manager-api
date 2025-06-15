package service

import (
	"product-manager-api/internal/domain"
	"product-manager-api/internal/entity"
	"product-manager-api/internal/repository"
	"product-manager-api/pkg"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: userRepository,
	}
}

func (s *AuthServiceImpl) Register(request domain.RegisterRequest, jwtService JWTService) (domain.RegisterResponse, error) {
	if err := pkg.Validate(request); err != nil {
		return domain.RegisterResponse{}, err
	}

	existingUser, _ := s.userRepository.FindByEmail(request.Email)
	if existingUser != nil {
		return domain.RegisterResponse{}, pkg.ErrEmailAlreadyExists
	}

	hashPassword,_ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	user := &entity.User{
		Email:    request.Email,
		Password: string(hashPassword),
	}

	if err := s.userRepository.CreateUser(user); err != nil {
		return domain.RegisterResponse{}, err
	}

	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		return domain.RegisterResponse{}, err
	}

	return domain.RegisterResponse{Token: token}, nil
}