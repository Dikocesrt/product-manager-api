package routes

import (
	"product-manager-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	authHandler *handler.AuthHandler
}

func NewRoute(authHandler *handler.AuthHandler) *Route {
	return &Route{
		authHandler: authHandler,
	}
}

func (r *Route) RegisterRoutes(ge *gin.Engine) {
	authGroup := ge.Group("/auth")
	{
		authGroup.POST("/register", r.authHandler.Register)
	}
}