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

// Database queries and logic for user

// RegisterUser registers a new user
func RegisterUser(u viewmodels.User) error {

	hashedPassword, err := helper.HashPassword(u.Password)
	if err != nil {
		println("Error hashing password:", err)
		return err
	}

	query := "CALL RegisterUser(?, ?, ?, ?, ?, ?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(`Error preparing query "CALL RegisterUser(?, ?, ?, ?, ?, ?)"`, err)
		return err
	}
	defer st.Close()

	var created int
	err = st.QueryRow(u.Name, u.Email, hashedPassword, u.Phone, time.Now().Unix(), time.Now().Unix()).Scan(&created)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	if created == 0 {
		log.Printf("User with that email already exists")
		return errors.New("user with that email already exists")
	}

	return nil
}

// LoginUser logs in a user, it checks if the user exists and if the password matches
func LoginUser(u viewmodels.UserLoginRequest) error {
	//query := "CALL LoginUser(?, ?)"
	//st, err := db.DB.Prepare(query)
	//if err != nil {
	//	log.Println(`Error preparing query "CALL LoginUser(?, ?)"`, err)
	//	return err
	//}
	//defer st.Close()
	//
	//hashedPassword, err := helper.HashPassword(u.Password)
	//
	//result, err := db.DB.Exec(query, u.Email, hashedPassword)
	//if err != nil {
	//	log.Println(`Error executing query "CALL LoginUser(?, ?)"`, err)
	//	return err
	//}
	//
	//rowsAffected, err := result.RowsAffected()
	//if err != nil {
	//	println("Error getting rows affected:", err)
	//	return err
	//}
	//
	//if rowsAffected == 0 {
	//	log.Println("Entered wrong email or password")
	//	return errors.New("entered wrong email or password")
	//}

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

	return nil
}

// DeleteUser deletes a user from the database
func DeleteUser(u viewmodels.UserLoginRequest) error {
	query := "SELECT id, password FROM users WHERE email = ?"
	var id int
	var password string
	err := db.DB.QueryRow(query, u.Email).Scan(&id, &password)
	if err != nil {
		log.Println("User does not exist:", err)
		return errors.New("user does not exist")
	}

	err = helper.CompareHashedPassword(password, u.Password)
	if err != nil {
		log.Println("Error comparing password:", err)
		return errors.New("error comparing password")
	}

	query = "DELETE FROM users WHERE email = ?"
	database, err := db.DB.Exec(query, u.Email)
	if err != nil {
		log.Printf("Failed to delete user with email %v. error is %v \n ", u.Email, err)
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

// EditUser edits a user's information
func EditUser(tokenString string, u viewmodels.User) error {
	// Retrieve user information from the token
	claims, err := helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		return errors.New("token validation failed")
	}

	hashedPassword, err := helper.HashPassword(u.Password)
	if err != nil {
		println("Error hashing password:", err)
		return err
	}

	idQuery := "SELECT id FROM users WHERE email = ?"
	var id int
	err = db.DB.QueryRow(idQuery, u.Email).Scan(&id)
	if err != nil {
		log.Println("Error getting user id:", err)
		return errors.New("error getting user id")
	}

	query := "UPDATE users SET name = ?, email = ?, password = ?, phone = ?, updated_at = ? WHERE id = ?"
	database, err := db.DB.Exec(query, u.Name, u.Email, hashedPassword, u.Phone, time.Now().Unix(), id)
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
