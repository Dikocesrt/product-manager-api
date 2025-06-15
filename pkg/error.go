package pkg

import (
	"errors"
	"net/http"
)

var (
	ErrEmptyField       = errors.New("field cannot be empty")
	ErrInvalidToken     = errors.New("invalid token")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

func ConvertErrorCode(err error) int {
	switch err {
	case ErrEmptyField:
		return http.StatusBadRequest
	case ErrInvalidToken:
		return http.StatusUnauthorized
	case ErrEmailAlreadyExists:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}