package db

import (
	"github.com/BurdockBH/food-delivery-rest-service/config"
	"testing"
)

func TestValidateDbConfig(t *testing.T) {
	const (
		username = "root"
		password = "root"
		host     = "localhost"
		port     = "3306"
		dbname   = "food_delivery"
	)

	cfg := &config.DatabaseConfig{
		DBUsername: username,
		DBPassword: password,
		DBHost:     host,
		DBPort:     port,
		DBName:     dbname,
	}

	err := ValidateDbConfig(cfg)
	if err != nil {
		t.Errorf("got error: %v", err)
		return
	}
	t.Logf("config is valid")
}
