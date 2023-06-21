package db

import (
	"errors"
	"github.com/BurdockBH/food-delivery-rest-service/config"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"testing"
)

func TestValidateDbConfig(t *testing.T) {

	TestCases := []viewmodels.TestCase{
		{
			Name:           "request DatabaseConfig failed because of invalid DBUsername",
			ExpectedResult: errors.New("db username is required"),
			Data: config.DatabaseConfig{
				DBUsername: "",
				DBPassword: "password",
				DBHost:     "localhost",
				DBPort:     "port",
				DBName:     "dbname"},
		},
		{
			Name:           "request DatabaseConfig failed because of invalid DBPassword",
			ExpectedResult: errors.New("db password is required"),
			Data: config.DatabaseConfig{
				DBUsername: "user",
				DBPassword: "",
				DBHost:     "localhost",
				DBPort:     "port",
				DBName:     "dbname",
			},
		},
		{
			Name:           "request DatabaseConfig failed because of invalid DBHost",
			ExpectedResult: errors.New("db host is required"),
			Data: config.DatabaseConfig{
				DBUsername: "user",
				DBPassword: "password",
				DBHost:     "",
				DBPort:     "port",
				DBName:     "dbname",
			},
		},
		{
			Name:           "request DatabaseConfig failed because of invalid DBPort",
			ExpectedResult: errors.New("db port is required"),
			Data: config.DatabaseConfig{
				DBUsername: "user",
				DBPassword: "password",
				DBHost:     "localhost",
				DBPort:     "",
				DBName:     "dbname",
			},
		},
		{
			Name:           "request DatabaseConfig failed because of invalid DBName",
			ExpectedResult: errors.New("db name is required"),
			Data: config.DatabaseConfig{
				DBUsername: "user",
				DBPassword: "password",
				DBHost:     "localhost",
				DBPort:     "port",
				DBName:     "",
			},
		},
		{
			Name:           "all fields are valid should return nil",
			ExpectedResult: nil,
			Data: config.DatabaseConfig{
				DBUsername: "user",
				DBPassword: "password",
				DBHost:     "localhost",
				DBPort:     "port",
				DBName:     "dbname",
			},
		},
	}

	failed := false
	for _, u := range TestCases {
		user := u.Data.(config.DatabaseConfig)
		err := ValidateDbConfig(&user)
		if (err != nil && err.Error() != u.ExpectedResult.Error()) || (err == nil && u.ExpectedResult != nil) {
			t.Errorf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			log.Printf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			failed = true
		}
	}

	if failed {
		return
	}

	t.Logf("Validation function is validating correctly")
}
