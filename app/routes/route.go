package routes

import (
	aHandler "product-manager-api/internal/auth/handler"
	"product-manager-api/internal/jwt/service"

	"github.com/gin-gonic/gin"
)

type Route struct {
	authHandler *aHandler.AuthHandler
	jwtService service.JWTService
}

func NewRoute(authHandler *aHandler.AuthHandler, jwtService service.JWTService) *Route {
	return &Route{
		authHandler: authHandler,
		jwtService: jwtService,
	}
}

func (r *Route) RegisterRoutes(ge *gin.Engine) {
	authGroup := ge.Group("/auth")
	{
		authGroup.POST("/register", r.authHandler.Register)
	}
}