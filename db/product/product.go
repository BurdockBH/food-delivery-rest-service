package product

import (
	"errors"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
)

func CreateProduct(product *viewmodels.Product, email string) error {
	query := "CALL CreateProduct(?, ?, ?, ?, ? ,?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL CreateProduct(%v, %v, %v, %v, %v, %v, %v): %v", product.Name, product.Description, product.Price, product.FoodVenue.Name, product.FoodVenue.Address, email, err)
		return err
	}
	defer st.Close()

	var created int
	err = st.QueryRow(product.Name, product.Description, product.Price, product.FoodVenue.Name, product.FoodVenue.Address, email).Scan(&created)
	if err != nil {
		log.Printf("Error executing query: CALL CreateProduct(%v, %v, %v, %v, %v, %v): %v", product.Name, product.Description, product.Price, product.FoodVenue.Name, product.FoodVenue.Address, email, err)
		return err
	}

	if created == 0 {
		log.Printf("Product with name %v already exists in venue %v", product.Name, product.FoodVenue.Name)
		return errors.New(fmt.Sprintf("product with name %v already exists in venue %v", product.Name, product.FoodVenue.Name))
	}

	if created == -1 {
		log.Printf("There is no venue with name %v and address %v", product.FoodVenue.Name, product.FoodVenue.Address)
		return errors.New(fmt.Sprintf("there is no venue with name %v and address %v", product.FoodVenue.Name, product.FoodVenue.Address))
	}

	return nil
}
