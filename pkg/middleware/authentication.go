package middleware

import (
	"clinic-management/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		token, isValid := utils.ValidateToken(tokenStr)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("userID", claims["id"])
			c.Set("role", claims["role"])
		}
		c.Next()
	}
}
