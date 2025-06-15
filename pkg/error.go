package pkg

import (
	"errors"
	"net/http"
)

var (
	ErrEmptyField       = errors.New("field cannot be empty")
	ErrInvalidToken     = errors.New("invalid token")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound	   = errors.New("user not found")
	ErrUnAuthorizedAccess = errors.New("unauthorized access")
	ErrProductNotFound  = errors.New("product not found")

	ErrInternalServerError = errors.New("internal server error")
)

func ConvertErrorCode(err error) int {
	switch err {
	case ErrEmptyField:
		return http.StatusBadRequest
	case ErrInvalidToken:
		return http.StatusUnauthorized
	case ErrEmailAlreadyExists:
		return http.StatusConflict
	case ErrInvalidCredentials:
		return http.StatusUnauthorized
	case ErrUserNotFound:
		return http.StatusNotFound
	case ErrUnAuthorizedAccess:
		return http.StatusForbidden
	case ErrProductNotFound:
		return http.StatusNotFound
	default:
		// Check if it's a validation error
		if _, ok := err.(ValidationErrors); ok {
			return http.StatusBadRequest
		}
		return http.StatusInternalServerError
	}
}