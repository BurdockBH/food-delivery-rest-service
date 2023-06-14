package user

import (
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
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
		return err
	}

	rowsAffected, err := database.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}
