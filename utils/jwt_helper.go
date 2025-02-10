package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// var secretKey = []byte(os.Getenv("JWT_SECRET")) // ✅ Ensure consistency

func GenerateToken(userID int, role string) (string, string, error) {
	godotenv.Load()
	var secretKey = []byte(os.Getenv("JWT_SECRET"))
	fmt.Println(secretKey)
	// Access Token (2 hours)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	})

	// Refresh Token (7 days)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	accessString, err := accessToken.SignedString(secretKey) // ✅ Use global secretKey
	if err != nil {
		return "", "", err
	}

	refreshString, err := refreshToken.SignedString(secretKey) // ✅ Use global secretKey
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}
