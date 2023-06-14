package user

import (
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"time"
)

func RegisterUser(u viewmodels.User) error {
	existingUserQuery := "SELECT COUNT(*) FROM users WHERE email = ? OR position()"
	var count int
	err := db.DB.QueryRow(existingUserQuery, u.Email).Scan(&count)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (name, email, phone, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	database, err := db.DB.Exec(query, u.Name, u.Email, u.Phone, u.Password, time.Now(), time.Now())
	if err != nil {
		log.Printf("Failed to create user with name %v. error is %v \n ", u.Name, err)
		return err
	}

	rowsAffected, err := database.RowsAffected()
	if err != nil {
		log.Printf("Error with rows affected %v \n", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("No rows affected")
		return err
	}

	return nil
}
