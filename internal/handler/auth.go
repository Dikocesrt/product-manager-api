package handler

import (
	"net/http"
	"product-manager-api/internal/domain"
	"product-manager-api/internal/service"
	"product-manager-api/pkg"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
	jwtService service.JWTService
}

func NewAuthHandler(authService service.AuthService, jwtService service.JWTService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		jwtService: jwtService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request domain.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(pkg.ConvertErrorCode(err), pkg.NewBaseErrorResponse(err.Error()))
		return
	}

	response, err := h.authService.Register(request, h.jwtService)
	if err != nil {
		c.JSON(pkg.ConvertErrorCode(err), pkg.NewBaseErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, pkg.NewBaseSuccessResponse("success register", response))
}