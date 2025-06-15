package main

import (
	"fmt"
	"os"
	"product-manager-api/app/routes"
	"product-manager-api/config"
	aHandler "product-manager-api/internal/auth/handler"
	aService "product-manager-api/internal/auth/service"
	jService "product-manager-api/internal/jwt/service"
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

	routes := routes.NewRoute(authHandler, jwtService)

	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(fmt.Sprintf(":%s", port))
}