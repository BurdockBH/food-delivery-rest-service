package db

import (
	"github.com/BurdockBH/food-delivery-rest-service/config"
	"testing"
)

func TestValidateDbConfig(t *testing.T) {

	testCases := []config.DatabaseConfig{
		{
			DBUsername: "",
			DBPassword: "password",
			DBHost:     "localhost",
			DBPort:     "port",
			DBName:     "dbname",
		},
		{
			DBUsername: "user",
			DBPassword: "",
			DBHost:     "localhost",
			DBPort:     "port",
			DBName:     "dbname",
		},
		{
			DBUsername: "user",
			DBPassword: "password",
			DBHost:     "",
			DBPort:     "port",
			DBName:     "dbname",
		},
		{
			DBUsername: "user",
			DBPassword: "password",
			DBHost:     "localhost",
			DBPort:     "",
			DBName:     "dbname",
		},
		{
			DBUsername: "user",
			DBPassword: "password",
			DBHost:     "localhost",
			DBPort:     "port",
			DBName:     "",
		},
	}

	correctConfig := config.DatabaseConfig{
		DBUsername: "user",
		DBPassword: "password",
		DBHost:     "localhost",
		DBPort:     "port",
		DBName:     "dbname",
	}

	failed := false
	for _, test := range testCases {
		err := ValidateDbConfig(&test)
		if err == nil {
			t.Errorf("Config %v should have failed validation", test.DBName)
			failed = true
		}
	}

	err := ValidateDbConfig(&correctConfig)
	if err != nil {
		t.Errorf("Error on valid data: %v", correctConfig.DBName)
	}

	if failed {
		return
	}

	t.Logf("Validation function is validating correctly")
}
