package viewmodels

import (
	"testing"
	"time"
)

func TestUser_Validate(t *testing.T) {
	var u = User{
		Name:      "Burdock",
		Email:     "edocicak@gmail.com",
		Password:  "password",
		Phone:     "1234567890",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	err := u.Validate()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("User: %v validated succesfully", u.Name)
}

func TestUserLoginRequest_ValidateLogin(t *testing.T) {
	var u = UserLoginRequest{
		Email:    "edocicak@gmail.com",
		Password: "password"}

	err := u.ValidateLogin()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("User: %v validated succesfully", u.Email)
}

func TestProduct_ValidateProduct(t *testing.T) {
	var p = Product{
		Name:        "Burger",
		Description: "Burger with fries",
		Price:       10,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	err := p.ValidateProduct()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("Product: %v validated succesfully", p.Name)
}

func TestFoodVenue_ValidateFoodVenue(t *testing.T) {
	var fv = FoodVenue{
		Name:      "Burger King",
		Address:   "123 Main St",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	err := fv.ValidateFoodVenue()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("FoodVenue: %v validated succesfully", fv.Name)
}
