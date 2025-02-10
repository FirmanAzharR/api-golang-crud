package middleware

import (
	"net/http"
	"os"
	"strings"

	"api-golang-crud/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// AuthMiddleware untuk memvalidasi token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load environment variables
		godotenv.Load()
		secretKey := []byte(os.Getenv("JWT_SECRET"))

		// Ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
			c.Abort()
			return
		}

		// Format: "Bearer {token}"
		tokenPart := strings.Split(authHeader, " ")
		if len(tokenPart) != 2 || tokenPart[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token salah"})
			c.Abort()
			return
		}

		tokenString := tokenPart[1]

		// Cek apakah token ada di blacklist
		if utils.IsTokenBlacklisted(tokenString) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token sudah logout"})
			c.Abort()
			return
		}

		// Parsing dan verifikasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Pastikan metode signing sesuai
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		c.Set("user", token)

		c.Next()
	}
}
