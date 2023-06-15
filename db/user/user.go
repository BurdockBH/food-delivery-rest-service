package user

import (
	"errors"
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// RegisterUser registers a new user
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

	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		println("Error hashing password:", err)
		return err
	}

	query := "INSERT INTO users (name, email, phone, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	database, err := db.DB.Exec(query, u.Name, u.Email, u.Phone, hashedPassword, time.Now().Unix(), time.Now().Unix())
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

func LoginUser(u viewmodels.UserLogin) error {
	query := "SELECT id, password FROM users WHERE email = ?"
	var id int
	var password string
	err := db.DB.QueryRow(query, u.Email).Scan(&id, &password)
	if err != nil {
		log.Println("User does not exist:", err)
		return errors.New("user does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
	if err != nil {
		log.Println("Error comparing password:", err)
		return errors.New("error comparing password")
	}
	return nil
}

func DeleteUser(u viewmodels.UserLogin) error {
	query := "SELECT id, password FROM users WHERE email = ?"
	var id int
	var password string
	err := db.DB.QueryRow(query, u.Email).Scan(&id, &password)
	if err != nil {
		log.Println("User does not exist:", err)
		return errors.New("user does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
	if err != nil {
		log.Println("Error comparing password:", err)
		return errors.New("error comparing password")
	}

	query = "DELETE FROM users WHERE email = ?"
	database, err := db.DB.Exec(query, u.Email)
	if err != nil {
		log.Printf("Failed to create user with email %v. error is %v \n ", u.Email, err)
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

func EditUser(tokenString string, u viewmodels.User) error {
	// Retrieve user information from the token
	claims, err := helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		return errors.New("token validation failed")
	}

	userEMAIL, ok := claims["email"].(string)
	if !ok {
		log.Println("User ID is missing or has an invalid type in the token claims")
		return errors.New("invalid user ID in token claims")
	}

	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		println("Error hashing password:", err)
		return err
	}

	query := "UPDATE users SET name = ?, email = ?, password = ?, phone = ?, updated_at = ? WHERE email = ?"
	database, err := db.DB.Exec(query, u.Name, u.Email, hashedPassword, u.Phone, time.Now().Unix(), userEMAIL)
	if err != nil {
		log.Printf("Failed to update user with email %v. Error: %v\n", claims["email"].(string), err)
		return err
	}

	_, err = database.RowsAffected()
	if err != nil {
		log.Printf("Error with rows affected: %v\n", err)
		return err
	}

	return nil
}

// Function for password hashing using bcrypt
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
