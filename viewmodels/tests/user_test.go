package tests

import (
	"errors"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"testing"
	"time"
)

func TestUser_Validate(t *testing.T) {

	var TestCases = []viewmodels.TestCase{
		{
			Name:           "request User failed because of invalid name",
			ExpectedResult: errors.New("invalid name"),
			Data: viewmodels.User{
				Name:      "Edo123",
				Email:     "edocicak@gmail.com",
				Password:  "password",
				Phone:     "1234567890",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
		},
		{
			Name:           "request User failed because of invalid email",
			ExpectedResult: errors.New("invalid email"),
			Data: viewmodels.User{
				Name:      "Burdock",
				Email:     "edocicakgmail.com",
				Password:  "password",
				Phone:     "1234567890",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
		},
		{
			Name:           "request User failed because of invalid password",
			ExpectedResult: errors.New("invalid password"),
			Data: viewmodels.User{
				Name:      "Burdock",
				Email:     "edocicak@gmail.com",
				Password:  "pass",
				Phone:     "1234567890",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
		},
		{
			Name:           "request User failed because of invalid phone",
			ExpectedResult: errors.New("invalid phone"),
			Data: viewmodels.User{
				Name:      "Burdock",
				Email:     "edocicak@gmail.com",
				Password:  "password",
				Phone:     "1234567abcd",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
		},
		{
			Name:           "all data is valid, should return nil",
			ExpectedResult: nil,
			Data: viewmodels.User{
				Name:      "Burdock",
				Email:     "edocicak@gmail.com",
				Password:  "password",
				Phone:     "1234567890",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}

	for _, u := range TestCases {
		user := u.Data.(viewmodels.User)
		err := user.Validate()
		if (err != nil && err.Error() != u.ExpectedResult.Error()) || (err == nil && u.ExpectedResult != nil) {
			t.Errorf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			log.Printf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			return
		}
	}

	t.Log("User validation is validating correctly")

}

func TestUserLoginRequest_ValidateLogin(t *testing.T) {
	TestCases := []viewmodels.TestCase{
		{
			Name:           "request UserLogin failed because of invalid email",
			ExpectedResult: errors.New("invalid email"),
			Data: viewmodels.UserLoginRequest{
				Email:    "edocicak.com",
				Password: "password",
			},
		},
		{
			Name:           "request UserLogin failed because of invalid password",
			ExpectedResult: errors.New("invalid password"),
			Data: viewmodels.UserLoginRequest{
				Email:    "edocicak@gmail.com",
				Password: "passasdofjadsjfoiaj0erjt0j092j9i2",
			},
		},
		{
			Name:           "all data is valid, should return nil",
			ExpectedResult: nil,
			Data: viewmodels.UserLoginRequest{
				Email:    "edocicak@gmail.com",
				Password: "password",
			},
		},
	}

	for _, u := range TestCases {
		user := u.Data.(viewmodels.UserLoginRequest)
		err := user.ValidateLogin()
		if (err != nil && err.Error() != u.ExpectedResult.Error()) || (err == nil && u.ExpectedResult != nil) {
			t.Errorf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			log.Printf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			return
		}
	}

	t.Log("User login validation is validating correctly")
}
