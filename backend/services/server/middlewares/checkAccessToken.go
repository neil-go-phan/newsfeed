package middlewares

import (
	"errors"
	"net/http"
	"server/entities"
	"server/helpers"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authTokenHeader := c.Request.Header.Get("Authorization")
		if authTokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized access"})
			c.Abort()
			return
		}
		tokenString := strings.Split(authTokenHeader, " ")[1]
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

func validateToken(tokenString string) (*entities.JWTClaim, error) {
	claims := &entities.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return helpers.TOKEN_SERECT_KEY, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token invalid")
	}
	return claims, nil
}
