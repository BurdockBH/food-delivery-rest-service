package food_venue

import (
	"errors"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"time"
)

func CreateFoodVenue(fv *viewmodels.FoodVenue) error {
	query := "CALL CreateFoodVenue(?, ?, ?, ?)"

	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL CreateFoodVenue(%v, %v): %v", fv.Name, fv.Address, err)
		return err
	}
	defer st.Close()

	var created int
	err = st.QueryRow(fv.Name, fv.Address, time.Now().Unix(), time.Now().Unix()).Scan(&created)
	if err != nil {
		log.Printf("Error executing query: CALL CreateFoodVenue(%v, %v): %v", fv.Name, fv.Address, err)
		return err
	}

	if created == 0 {
		log.Printf("Food venue with name %v and address %v already exists", fv.Name, fv.Address)
		return errors.New(fmt.Sprintf("Food venue with name: %v and address %v already exists", fv.Name, fv.Address))
	}

	return nil
}

func DeleteFoodVenue(fv *viewmodels.FoodVenue) error {
	var deleted int
	query := "CALL DeleteFoodVenue(?, ?, ?)"

	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL DeleteFoodVenue(%v, %v): %v", fv.Name, fv.Address, err)
		return err
	}
	defer st.Close()

	err = st.QueryRow(fv.ID, fv.Name, fv.Address).Scan(&deleted)
	if err != nil {
		log.Printf("Error executing query: CALL DeleteFoodVenue(%v, %v): %v", fv.Name, fv.Address, err)
		return err
	}

	if deleted == 0 {
		log.Printf("Food venue with name %v and address %v or id %v does not exist", fv.Name, fv.Address, fv.ID)
		return errors.New(fmt.Sprintf("Food venue with name %v and address %v or id %v does not exist", fv.Name, fv.Address, fv.ID))
	}

	return nil
}
