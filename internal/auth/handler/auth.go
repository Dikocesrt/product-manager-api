package handler

import (
	"net/http"

	aDomain "product-manager-api/internal/auth/domain"
	aService "product-manager-api/internal/auth/service"
	jService "product-manager-api/internal/jwt/service"
	"product-manager-api/pkg"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService aService.AuthService
	jwtService jService.JWTService
}

func NewAuthHandler(authService aService.AuthService, jwtService jService.JWTService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		jwtService: jwtService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request aDomain.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewBaseErrorResponse("Invalid request format"))
		return
	}

	response, err := h.authService.Register(request, h.jwtService)
	if err != nil {
		if validationErrors, ok := err.(pkg.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, pkg.NewValidationErrorResponse("Validation failed", validationErrors))
			return
		}
		
		c.JSON(pkg.ConvertErrorCode(err), pkg.NewBaseErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, pkg.NewBaseSuccessResponse("success register", response))
}