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

	rows, err := st.Exec(fv.Name, fv.Address, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		log.Printf("Error executing query: CALL CreateFoodVenue(%v, %v): %v", fv.Name, fv.Address, err)
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("errror with rows affected: %v", err)
		return errors.New(fmt.Sprintf("error with rows affected: %v", err))
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

func GetVenues(fv *viewmodels.FoodVenue) ([]viewmodels.FoodVenue, error) {
	query := "CALL GetVenues(?, ?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: CALL GetVenues(%v, %v): %v", fv.Name, fv.Address, err)
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(fv.Name, fv.Address)
	if err != nil {
		log.Printf("Error executing query: CALL GetVenues(%v, %v): %v", fv.Name, fv.Address, err)
		return nil, err
	}
	defer rows.Close()

	var venues []viewmodels.FoodVenue
	for rows.Next() {
		var venue viewmodels.FoodVenue
		err = rows.Scan(&venue.ID, &venue.Name, &venue.Address, &venue.CreatedAt, &venue.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		venues = append(venues, venue)
	}

	if venues == nil {
		log.Printf("No food venues found")
		return nil, errors.New("no food venues found")
	}

	return venues, nil
}
