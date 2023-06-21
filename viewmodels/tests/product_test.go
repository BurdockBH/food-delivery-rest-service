package tests

import (
	"errors"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"testing"
)

func TestProduct_ValidateProduct(t *testing.T) {
	TestCases := []viewmodels.TestCase{
		{
			Name:           "request Product failed because of invalid name",
			ExpectedResult: errors.New("name cannot be empty"),
			Data: viewmodels.Product{
				Name:        "",
				Description: "A delicious burger",
				Price:       5.99,
			},
		},
		{
			Name:           "request Product failed because of invalid description",
			ExpectedResult: errors.New("description cannot be empty"),
			Data: viewmodels.Product{
				Name:        "Burger",
				Description: "",
				Price:       5.99,
			},
		},
		{
			Name:           "request Product failed because of invalid price",
			ExpectedResult: errors.New("price cannot be negative"),
			Data: viewmodels.Product{
				Name: "Burger",

				Description: "A delicious burger",
				Price:       -5.99,
			},
		},
	}

	for _, u := range TestCases {
		user := u.Data.(viewmodels.Product)
		err := user.ValidateProduct()
		if (err != nil && err.Error() != u.ExpectedResult.Error()) || (err == nil && u.ExpectedResult != nil) {
			t.Errorf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			log.Printf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			return
		}
	}

	t.Log("Product validation is validating correctly")
}

func TestFoodVenue_ValidateFoodVenue(t *testing.T) {
	TestCases := []viewmodels.TestCase{
		{
			Name:           "request FoodVenue failed because of invalid name",
			ExpectedResult: errors.New("name cannot be empty"),
			Data: viewmodels.FoodVenue{
				Name:    "",
				Address: "1234 Main St",
			},
		},

		{
			Name:           "request FoodVenue failed because of invalid address",
			ExpectedResult: errors.New("address cannot be empty"),
			Data: viewmodels.FoodVenue{
				Name:    "Burger",
				Address: "",
			},
		},
	}

	for _, u := range TestCases {
		user := u.Data.(viewmodels.FoodVenue)
		err := user.ValidateFoodVenue()
		if (err != nil && err.Error() != u.ExpectedResult.Error()) || (err == nil && u.ExpectedResult != nil) {
			t.Errorf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			log.Printf("Test for %v\nShould get error: %v but got: %v", u.Name, err, u.ExpectedResult)
			return
		}
	}

	t.Log("FoodVenue validation is validating correctly")
}
