package viewmodels

import "errors"

type Order struct {
	ID         int64 `json:"id"`
	ProductID  int64 `json:"product_id"`
	UserId     int64 `json:"user_id"`
	Quantity   int64 `json:"quantity"`
	Price      int64 `json:"price"`
	TotalPrice int64 `json:"total_price"`
}

func (o *Order) ValidateOrder() error {

	if o.ProductID == 0 {
		return errors.New("product id cannot be empty")
	}
	if o.Quantity == 0 {
		return errors.New("quantity cannot be empty")
	}

	return nil
}
