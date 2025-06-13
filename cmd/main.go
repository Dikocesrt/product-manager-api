package main

import (
	"fmt"
	"os"
	"product-manager-api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	r := gin.Default()

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(fmt.Sprintf(":%s", port))
}