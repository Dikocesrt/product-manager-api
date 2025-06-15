package pkg

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func parseToken(token string) (*jwt.Token, error) {
	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, ErrInvalidToken
	}
	tokenString := parts[1]
	
	result, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	
	if err != nil {
		return nil, ErrInvalidToken
	}

	return result, nil
}

func GetIDFromToken(token string) (uint, error) {
	result, err := parseToken(token)
	if err != nil {
		return 0, err 
	}
	claims, ok := result.Claims.(jwt.MapClaims)
	if ok && result.Valid {
		return uint(claims["id"].(float64)), nil
	}

	return 0, ErrInvalidToken
}