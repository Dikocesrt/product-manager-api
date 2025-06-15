package middleware

import (
	"product-manager-api/internal/jwt/service"
	"product-manager-api/pkg"

	"github.com/gin-gonic/gin"
)

func Authentication(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(pkg.ConvertErrorCode(pkg.ErrUnAuthorizedAccess), pkg.NewBaseErrorResponse(pkg.ErrUnAuthorizedAccess.Error()))
			c.Abort()
			return
		}

		// Extract user ID from token
		userID, err := jwtService.GetIDFromToken(authHeader)
		if err != nil {
			c.JSON(pkg.ConvertErrorCode(err), pkg.NewBaseErrorResponse(err.Error()))
			c.Abort()
			return
		}

		// Set userID in context for use in subsequent handlers
		c.Set("userID", userID)
		c.Next()
	}
}
