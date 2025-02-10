package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken, _ := c.Get("user")
		token := userToken.(*jwt.Token)

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["role"] != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak"})
			c.Abort()
			return
		}

		c.Next()
	}
}
