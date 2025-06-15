package main

import (
	"fmt"
	"os"
	"product-manager-api/config"
	"product-manager-api/internal/handler"
	"product-manager-api/internal/repository"
	"product-manager-api/internal/routes"
	"product-manager-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	r := gin.New()
	
	// Add middleware
	// r.Use(gin.Recovery())
	// r.Use(middleware.Logger())
	// r.Use(middleware.CORS())

	userRepo := repository.NewUserRepository(config.DB)

	jwtService := service.NewJWTService()

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService, jwtService)

	routes := routes.NewRoute(authHandler, jwtService)

	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(fmt.Sprintf(":%s", port))
}