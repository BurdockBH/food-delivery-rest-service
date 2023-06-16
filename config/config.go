package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
//type Config struct {
//	DBUsername string
//	DBPassword string
//	DBHost     string
//	DBPort     string
//	DBName     string
//}

type DatabaseConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

type JWTConfig struct {
	JWTSecret string
}

func LoadConfig() (*DatabaseConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	config := &DatabaseConfig{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}

	return config, nil
}

func LoadJWTConfig() (*JWTConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	config := &JWTConfig{
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	return config, nil
}
