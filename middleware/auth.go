package middleware

import (
	"github.com/amrremam/EBE.git/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTMiddleware ensures only authenticated users access protected routes
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Parse and validate token
		claims, err := auth.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store user ID from JWT in context for use in handlers
		c.Set("userID", claims.Issuer)

		// Continue to the next middleware or handler
		c.Next()
	}
}
