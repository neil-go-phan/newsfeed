package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func ExpiredAccessTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("X-Refresh-Token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "No refresh token provided"})
			c.Abort()
			return
		}

		claims, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized access"})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}