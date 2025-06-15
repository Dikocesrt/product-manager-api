package service

import (
	"os"
	"product-manager-api/pkg"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTServiceImpl struct {

}

func NewJWTService() *JWTServiceImpl {
	return &JWTServiceImpl{}
}

func (s JWTServiceImpl) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1 * 7).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s JWTServiceImpl) GetIDFromToken(token string) (uint, error) {
	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, pkg.ErrInvalidToken
	}
	tokenString := parts[1]
	
	result, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	
	if err != nil {
		return 0, pkg.ErrInvalidToken
	}

	claims, ok := result.Claims.(jwt.MapClaims)
	if ok && result.Valid {
		return uint(claims["id"].(float64)), nil
	}

	return 0, pkg.ErrInvalidToken
}