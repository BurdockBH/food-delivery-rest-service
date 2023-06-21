package viewmodels

import (
	"testing"
	"time"
)

func TestUser_Validate(t *testing.T) {

	var testCases = []User{
		{
			Name:      "Edo123",
			Email:     "edocicak@gmail.com",
			Password:  "password",
			Phone:     "1234567890",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
		{
			Name:      "Burdock",
			Email:     "edodafkoa.com",
			Password:  "password",
			Phone:     "1234567890",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
		{
			Name:      "Burdock",
			Email:     "edocicak@gmail.com",
			Password:  "pass",
			Phone:     "1234567890",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
		{
			Name:      "Burdock",
			Email:     "edocicak@gmail.com",
			Password:  "password",
			Phone:     "1234567abcd",
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}

	failed := false
	for _, u := range testCases {
		err := u.Validate()
		if err == nil {
			t.Errorf("User: %v should have failed validation", u.Name)
			failed = true
		}
	}

	correctUser := User{
		Name:      "Burdock",
		Email:     "edocicak@gmail.com",
		Password:  "password",
		Phone:     "1234567890",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	err := correctUser.Validate()
	if err != nil {
		t.Errorf("Error on valid data: %v", correctUser.Name)
		return
	}

	if failed {
		return
	}

	t.Log("User validation is validating correctly")
}

func TestUserLoginRequest_ValidateLogin(t *testing.T) {
	testCases := []UserLoginRequest{
		{
			Email:    "edocicak.com",
			Password: "password",
		},
		{
			Email:    "edocicak@gmail.com",
			Password: "passasdofjadsjfoiaj0erjt0j092j9i2",
		},
	}

	failed := false
	for _, u := range testCases {
		err := u.ValidateLogin()
		if err == nil {
			t.Errorf("User: %v should have failed validation", u.Email)
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
