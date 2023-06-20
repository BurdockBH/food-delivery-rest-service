package viewmodels

import "errors"

type FoodVenue struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (fv *FoodVenue) ValidateFoodVenue() error {
	if fv.Address == "" {
		return errors.New("address cannot be empty")
	} else if fv.Name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}
