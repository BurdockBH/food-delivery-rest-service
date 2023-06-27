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

func DeleteProduct(id int) error {
	query := "CALL DeleteProduct(?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL DeleteProduct(%v): %v", id, err)
		return err
	}
	defer st.Close()

	var deleted int
	err = st.QueryRow(id).Scan(&deleted)
	if err != nil {
		log.Printf("Error executing query: CALL DeleteProduct(%v): %v", id, err)
		return err
	}

	if deleted == 0 {
		log.Printf("Product with id %v does not exist", id)
		return errors.New(fmt.Sprintf("Product with id %v does not exist", id))
	}

	return nil
}

func EditProduct(product *viewmodels.Product) error {
	query := "CALL EditProduct(?, ?, ?, ?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL EditProduct(%v, %v, %v, %v): %v", product.ID, product.Name, product.Description, product.Price, err)
		return err
	}
	defer st.Close()

	var edited int
	err = st.QueryRow(product.ID, product.Name, product.Description, product.Price).Scan(&edited)
	if err != nil {
		log.Printf("Error executing query: CALL EditProduct(%v, %v, %v, %v): %v", product.ID, product.Name, product.Description, product.Price, err)
		return err
	}

	if edited != 1 {
		log.Printf("Product with id %v does not exist", product.ID)
		return errors.New(fmt.Sprintf("Product with id %v does not exist", product.ID))
	}

	return nil
}

func GetProducts(venueId int) ([]viewmodels.Product, error) {
	query := "CALL GetProducts(?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL GetProducts(%v): %v", venueId, err)
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(venueId)
	if err != nil {
		log.Printf("Error executing query: CALL GetProducts(%v): %v", venueId, err)
		return nil, err
	}
	defer rows.Close()

	var products []viewmodels.Product
	for rows.Next() {
		var p viewmodels.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.FoodVenue.ID, &p.CreatedBy, &p.CreatedAt, &p.UpdatedAt, &p.FoodVenue.Name, &p.FoodVenue.Address)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			products = append(products, viewmodels.Product{})
		}
		products = append(products, p)
	}

	if len(products) == 0 {
		log.Println("No products found")
		return nil, errors.New("no products found")
	}

	return products, nil
}
