package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
	"time"
)

// ValidateToken validates the token
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	// Retrieve the JWT secret key from the environment variable
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// GenerateToken generates a JWT token
func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	err := godotenv.Load()
	if err != nil {
		return "Failed to load .env", err
	}
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		println("Failed to sign token:", err)
		return "Failed to sign token:", err
	}

	return tokenString, nil
}
