package user

import (
	"errors"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// Database queries and logic for user

// RegisterUser registers a new user
func RegisterUser(u *viewmodels.User) error {

	hashedPassword, err := helper.HashPassword(u.Password)
	if err != nil {
		println("Error hashing password:", err)
		return err
	}

	query := "CALL RegisterUser(?, ?, ?, ?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf(`Error preparing query "CALL RegisterUser(%v, %v, %v, %v): %v"`, u.Name, u.Email, hashedPassword, u.Phone, err)
		return err
	}
	defer st.Close()

	var created int
	err = st.QueryRow(u.Name, u.Email, hashedPassword, u.Phone).Scan(&created)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	if created != 1 {
		log.Printf("User with email: %v or phone number: %v already exists", u.Email, u.Phone)
		return errors.New(fmt.Sprintf("user with email %v or phone number %v already exists", u.Email, u.Phone))
	}

	return nil
}

// LoginUser logs in a user, it checks if the user exists and if the password matches
func LoginUser(u *viewmodels.UserLoginRequest) error {

	query := "CALL LoginUser(?)"
	var password string

	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf(`Error preparing query "CALL LoginUser(%v)": %v`, u.Email, err)
		return err
	}

	err = st.QueryRow(u.Email).Scan(&password)
	if err != nil && err.Error() == "sql: no rows in result set" {
		log.Printf("User with email %v does not exist", u.Email)
		return errors.New(fmt.Sprintf("user with email %v does not exist", u.Email))
	} else if err != nil { // other error
		log.Println("Error executing query:", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
	if err != nil {
		log.Println("Error comparing password:", err)
		return errors.New("error comparing password")
	}
	return nil
}

// DeleteUser deletes a user from the database
func DeleteUser(u *viewmodels.UserLoginRequest) error {
	passwordQuery := "CALL LoginUser(?)"
	var password string
	err := db.DB.QueryRow(passwordQuery, u.Email).Scan(&password)
	if err != nil {
		log.Println("User does not exist:", err)
		return errors.New(fmt.Sprintf("user %v does not exist", u.Email))
	}

	err = helper.CompareHashedPassword(password, u.Password)
	if err != nil {
		log.Println("Error comparing password:", err)
		return errors.New("error comparing password")
	}

	query := "CALL DeleteUser(?)"

	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf(`Error preparing query "CALL DeleteUser(%v)": %v`, u.Email, err)
		return err
	}

	var deleted int
	err = st.QueryRow(u.Email).Scan(&deleted)
	if err != nil {
		log.Printf("Failed to delete user with email %v. error is %v \n ", u.Email, err)
		return err
	}

	if deleted != 1 {
		log.Printf("Couldn't delete %v. No rows affected\n", u.Email)
		return errors.New(fmt.Sprintf("couldn't delete user %v. No rows affected", u.Email))
	}

	return nil
}

// EditUser edits a user's information
func EditUser(u *viewmodels.User) error {
	query := "CALL EditUser(?, ?, ?, ?, ?)"

	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf(`Error preparing query "CALL EditUser(%v, %v, %v, %v, %v": %v`, u.Name, u.Email, u.Password, u.Phone, time.Now().Unix(), err)
		return err
	}

	hashedPassword, err := helper.HashPassword(u.Password)
	if err != nil {
		log.Println("Failed to hash password", err)
		return err
	}

	var updated int
	err = st.QueryRow(u.Name, u.Email, hashedPassword, u.Phone, time.Now().Unix()).Scan(&updated)
	if err != nil {
		log.Printf("Failed to update user with email %v. Error: %v\n\n\n", u.Email, err)
		return err
	}

	if updated == -1 {
		log.Printf("User with email %v does not exist", u.Email)
		return errors.New(fmt.Sprintf("user with email %v does not exist", u.Email))
	} else if updated == -2 {
		log.Printf("User with phone number: %v already exists", u.Phone)
		return errors.New(fmt.Sprintf("user with phone number: %v already exists", u.Phone))
	}

	return nil
}

func GetUsers(u *viewmodels.User) ([]viewmodels.User, error) {
	query := "CALL GetUsersByDetails(?, ?, ?)"
	st, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf(`Error preparing query "CALL GetUsersByDetails(%v, %v, %v)": %v`, u.Name, u.Email, u.Phone, err)
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(u.Name, u.Email, u.Phone)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var users []viewmodels.User
	for rows.Next() {
		var user viewmodels.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			users = append(users, viewmodels.User{})
		}
		users = append(users, user)
	}

	if users == nil {
		log.Println("No users found")
		return nil, errors.New("no users found")
	}

	return users, nil
}
