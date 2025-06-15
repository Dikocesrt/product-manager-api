package routes

import (
	"product-manager-api/app/middleware"
	aHandler "product-manager-api/internal/auth/handler"
	"product-manager-api/internal/jwt/service"
	pHandler "product-manager-api/internal/product/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	authHandler *aHandler.AuthHandler
	jwtService service.JWTService
	productHandler *pHandler.ProductHandler
}

func NewRoute(authHandler *aHandler.AuthHandler, jwtService service.JWTService, productHandler *pHandler.ProductHandler) *Route {
	return &Route{
		authHandler: authHandler,
		jwtService: jwtService,
		productHandler: productHandler,
	}
}

func (r *Route) RegisterRoutes(ge *gin.Engine) {
	authGroup := ge.Group("/auth")
	{
		authGroup.POST("/register", r.authHandler.Register)
		authGroup.POST("/login", r.authHandler.Login)
	}

	productGroup := ge.Group("/products", middleware.Authentication(r.jwtService))
    {
        productGroup.POST("", r.productHandler.CreateProduct)
        productGroup.GET("", r.productHandler.GetAllProducts)
        productGroup.GET("/:id", r.productHandler.GetProductByID)
        productGroup.PUT("/:id", r.productHandler.UpdateProduct)
        productGroup.DELETE("/:id", r.productHandler.DeleteProduct)
    }
}