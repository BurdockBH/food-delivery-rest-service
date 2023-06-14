package user

import (
	"errors"
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"time"
)

func RegisterUser(u viewmodels.User) error {
	existingUserQuery := "SELECT COUNT(*) FROM users WHERE email = ?"
	var count int
	err := db.DB.QueryRow(existingUserQuery, u.Email).Scan(&count)
	if err != nil {
		log.Println("Error checking existing user:", err)
		return err
	}

	if count > 0 {
		log.Println("User with that email already exists")
		return errors.New("user with that email already exists")
	}

	query := "INSERT INTO users (name, email, phone, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	database, err := db.DB.Exec(query, u.Name, u.Email, u.Phone, u.Password, time.Now().Unix(), time.Now().Unix())
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
		return errors.New("no rows affected")
	}

	return nil
}
