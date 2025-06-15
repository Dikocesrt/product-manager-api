package main

import (
	"fmt"
	"os"
	"product-manager-api/app/routes"
	"product-manager-api/config"
	aHandler "product-manager-api/internal/auth/handler"
	aService "product-manager-api/internal/auth/service"
	jService "product-manager-api/internal/jwt/service"
	pHandler "product-manager-api/internal/product/handler"
	pRepository "product-manager-api/internal/product/repository"
	pService "product-manager-api/internal/product/service"
	uRepository "product-manager-api/internal/user/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	r := gin.New()
	
	// Add middleware
	// r.Use(gin.Recovery())
	// r.Use(middleware.Logger())
	// r.Use(middleware.CORS())

	userRepo := uRepository.NewUserRepository(config.DB)

	jwtService := jService.NewJWTService()

	authService := aService.NewAuthService(userRepo)
	authHandler := aHandler.NewAuthHandler(authService, jwtService)

	productRepo := pRepository.NewProductRepository(config.DB)
	productService := pService.NewProductService(productRepo)
	productHandler := pHandler.NewProductHandler(productService)

	routes := routes.NewRoute(authHandler, jwtService, productHandler)

	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(fmt.Sprintf(":%s", port))
}