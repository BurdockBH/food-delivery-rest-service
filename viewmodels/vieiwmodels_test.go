package viewmodels

import (
	"errors"
	"log"
	"testing"
	"time"
)

type testCase struct {
	ExpectedResult error
	Name           string
	Data           interface{}
}

func TestUser_Validate(t *testing.T) {

	var testCases = []testCase{
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
	for _, u := range testCases {
		user := u.Data.(User)
		err := user.Validate()
		if err != nil && err.Error() != u.ExpectedResult.Error() {
			t.Errorf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			log.Printf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			failed = true
		} else if err == nil && u.ExpectedResult != nil {
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
	testCases := []testCase{
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
	}

	failed := false
	for _, u := range testCases {
		user := u.Data.(UserLoginRequest)
		err := user.ValidateLogin()
		if err != u.ExpectedResult {
			t.Errorf("%v unexpected error: %v", u.Name, err.Error())
			log.Printf("%v unexpected error: %v", u.Name, err.Error())
			failed = true
		}
	}

	correctUserLogin := UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password123",
	}

	err := correctUserLogin.ValidateLogin()
	if err != nil {
		t.Errorf("Error on valid data: %v", correctUserLogin.Email)
		log.Printf("Error on valid data: %v", correctUserLogin.Email)
		return
	}

	if failed {
		return
	}

	t.Log("User login validation is validating correctly")
}

func TestProduct_ValidateProduct(t *testing.T) {
	testCases := []Product{
		{
			Name:        "",
			Description: "A delicious burger",
			Price:       5.99,
		},
		{
			Name:        "Burger",
			Description: "",
			Price:       5.99,
		},
		{
			Name:        "Burger",
			Description: "A delicious burger",
			Price:       -5.99,
		},
	}

	failed := false
	for _, p := range testCases {
		err := p.ValidateProduct()
		if err == nil {
			t.Errorf("Product: %v should have failed validation", p.Name)
			failed = true
		}
	}

	correctProduct := Product{
		Name:        "Burger",
		Description: "A delicious burger",
		Price:       5.99,
	}

	err := correctProduct.ValidateProduct()
	if err != nil {
		t.Errorf("Error on valid data: %v", correctProduct.Name)
		return
	}

	if failed {
		return
	}

	t.Log("Product validation is validating correctly")
}

func TestFoodVenue_ValidateFoodVenue(t *testing.T) {
	testCases := []FoodVenue{
		{
			Name:    "",
			Address: "1234 Main St",
		},
		{
			Name:    "Burger",
			Address: "",
		},
	}

	failed := false
	for _, fv := range testCases {
		err := fv.ValidateFoodVenue()
		if err == nil {
			t.Errorf("FoodVenue: %v should have failed validation", fv.Name)
			failed = true
		}
	}

	correctFoodVenue := FoodVenue{
		Name:    "Burger",
		Address: "1234 Main St",
	}

	err := correctFoodVenue.ValidateFoodVenue()
	if err != nil {
		t.Errorf("Error on valid data: %v", correctFoodVenue.Name)
		return
	}

	if failed {
		return
	}

	t.Log("FoodVenue validation is validating correctly")
}
