package helper

import (
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// ValidateToken validates the token
func ValidateToken(tokenString string) (jwt.MapClaims, error) {

	jwtSecret, err := config.LoadJWTConfig()
	if err != nil {
		println("Failed to load JWT config:", err)
		return nil, err
	}

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

	jwtConfig, err := config.LoadJWTConfig()
	if err != nil {
		println("Failed to load JWT config:", err)
		return "", err
	}

	jwtSecret := []byte(jwtConfig.JWTSecret)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		println("Failed to sign token:", err)
		return "", err
	}

	return tokenString, nil
}

// HashPassword hashes the password
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHashedPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
