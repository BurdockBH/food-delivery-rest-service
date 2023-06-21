package viewmodels

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestUser_Validate(t *testing.T) {

	var TestCases = []TestCase{
		{
			Name:           "request User failed because of invalid name",
			ExpectedResult: errors.New("invalid name"),
			Data: User{
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
			Data: User{
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
			Data: User{
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
			Data: User{
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
			Data: User{
				Name:      "Burdock",
				Email:     "edocicak@gmail.com",
				Password:  "password",
				Phone:     "1234567890",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}

	failed := false
	for _, u := range TestCases {
		user := u.Data.(User)
		err := user.Validate()
		if (err != nil && err.Error() != u.ExpectedResult.Error()) || (err == nil && u.ExpectedResult != nil) {
			t.Errorf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			log.Printf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			failed = true
		}
	}

	if failed {
		return
	}

	t.Log("User validation is validating correctly")
}

func TestUserLoginRequest_ValidateLogin(t *testing.T) {
	TestCases := []TestCase{
		{
			Name:           "request UserLogin failed because of invalid email",
			ExpectedResult: errors.New("invalid email"),
			Data: UserLoginRequest{
				Email:    "edocicak.com",
				Password: "password",
			},
		},
		{
			Name:           "request UserLogin failed because of invalid password",
			ExpectedResult: errors.New("invalid password"),
			Data: UserLoginRequest{
				Email:    "edocicak@gmail.com",
				Password: "passasdofjadsjfoiaj0erjt0j092j9i2",
			},
		},
		{
			Name:           "all data is valid, should return nil",
			ExpectedResult: nil,
			Data: UserLoginRequest{
				Email:    "edocicak@gmail.com",
				Password: "password",
			},
		},
	}

	failed := false
	for _, u := range TestCases {
		user := u.Data.(UserLoginRequest)
		err := user.ValidateLogin()
		if (err != nil && err.Error() != u.ExpectedResult.Error()) || (err == nil && u.ExpectedResult != nil) {
			t.Errorf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			log.Printf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			failed = true
		}
	}

	if failed {
		return
	}

	t.Log("User login validation is validating correctly")
}

func TestProduct_ValidateProduct(t *testing.T) {
	TestCases := []TestCase{
		{
			Name:           "request Product failed because of invalid name",
			ExpectedResult: errors.New("name cannot be empty"),
			Data: Product{
				Name:        "",
				Description: "A delicious burger",
				Price:       5.99,
			},
		},
		{
			Name:           "request Product failed because of invalid description",
			ExpectedResult: errors.New("description cannot be empty"),
			Data: Product{
				Name:        "Burger",
				Description: "",
				Price:       5.99,
			},
		},
		{
			Name:           "request Product failed because of invalid price",
			ExpectedResult: errors.New("price cannot be negative"),
			Data: Product{
				Name: "Burger",

				Description: "A delicious burger",
				Price:       -5.99,
			},
		},
	}

	failed := false
	for _, u := range TestCases {
		user := u.Data.(Product)
		err := user.ValidateProduct()
		if (err != nil && err.Error() != u.ExpectedResult.Error()) || (err == nil && u.ExpectedResult != nil) {
			t.Errorf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			log.Printf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			failed = true
		}
	}

	if failed {
		return
	}

	t.Log("Product validation is validating correctly")
}

func TestFoodVenue_ValidateFoodVenue(t *testing.T) {
	TestCases := []TestCase{
		{
			Name:           "request FoodVenue failed because of invalid name",
			ExpectedResult: errors.New("name cannot be empty"),
			Data: FoodVenue{
				Name:    "",
				Address: "1234 Main St",
			},
		},

		{
			Name:           "request FoodVenue failed because of invalid address",
			ExpectedResult: errors.New("address cannot be empty"),
			Data: FoodVenue{
				Name:    "Burger",
				Address: "",
			},
		},
	}

	failed := false
	for _, u := range TestCases {
		user := u.Data.(FoodVenue)
		err := user.ValidateFoodVenue()
		if (err != nil && err.Error() != u.ExpectedResult.Error()) || (err == nil && u.ExpectedResult != nil) {
			t.Errorf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			log.Printf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			failed = true
		}
	}

	if failed {
		return
	}

	t.Log("FoodVenue validation is validating correctly")
}
