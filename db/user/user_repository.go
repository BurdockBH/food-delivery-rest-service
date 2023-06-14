package user

import (
	"database/sql"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) RegisterUser(u viewmodels.User) error {
	existingUserQuery := "SELECT COUNT(*) FROM users WHERE email = ? OR position()"
	var count int
	err := ur.db.QueryRow(existingUserQuery, u.Email).Scan(&count)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (name, email, phone, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = ur.db.Exec(query, u.Name, u.Email, u.Phone, u.Password, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}
