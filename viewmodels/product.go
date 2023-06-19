package viewmodels

import "errors"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	FoodVenue   FoodVenue `json:"food_venue"`
	CreatedAt   int64     `json:"created_at"`
	UpdatedAt   int64     `json:"updated_at"`
}

type FoodVenue struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (p *Product) ValidateProduct() error {
	if len(p.Name) < 1 {
		return errors.New("price cannot be negative")
	} else if p.Price < 0 {
		return errors.New("name cannot be empty")
	} else if len(p.Description) < 1 {
		return errors.New("description cannot be empty")
	}
	return nil
}

func (fv *FoodVenue) ValidateFoodVenue() error {
	if fv.Address == "" {
		return errors.New("address cannot be empty")
	} else if fv.Name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}
