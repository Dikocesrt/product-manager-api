package middleware

import (
	"net/http"
	"product-manager-api/internal/jwt/service"
	"product-manager-api/pkg"

	"github.com/gin-gonic/gin"
)

func Authentication(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, pkg.NewBaseErrorResponse("Unauthorized"))
			c.Abort()
			return
		}

		// Extract user ID from token
		userID, err := jwtService.GetIDFromToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, pkg.NewBaseErrorResponse(err.Error()))
			c.Abort()
			return
		}

		// Set userID in context for use in subsequent handlers
		c.Set("userID", userID)
		c.Next()
	}
}
