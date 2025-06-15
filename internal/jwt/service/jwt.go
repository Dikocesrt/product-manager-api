package service

type JWTService interface {
	GenerateToken(userID uint) (string, error)
	GetIDFromToken(token string) (uint, error)
}