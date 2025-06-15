package routes

import (
	"product-manager-api/internal/handler"
	"product-manager-api/internal/service"

	"github.com/gin-gonic/gin"
)

type Route struct {
	authHandler *handler.AuthHandler
	jwtService service.JWTService
}

func NewRoute(authHandler *handler.AuthHandler, jwtService service.JWTService) *Route {
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