package user

import (
	"encoding/json"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

var jwtSecret = []byte("secret-key")

// RegisterUser registers a new user in the database
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	// Handler logic goes here
	var u viewmodels.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	validRequest, errString := u.Validate()
	if !validRequest {
		log.Println(errString)
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	err = user.RegisterUser(u)
	if err != nil {
		log.Println("Failed to register user:", err)
		http.Error(w, fmt.Sprintf("Failed to register user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("User registered successfully!")
	fmt.Fprintf(w, "User registered successfully!")
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	var userLogin viewmodels.UserLogin

	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = user.LoginUser(userLogin)
	if err != nil {
		log.Println("Failed to login user:", err)
		http.Error(w, fmt.Sprintf("Failed to login user: %v", err), http.StatusInternalServerError)
		return
	}

	token, err := GenerateToken(userLogin.Email)
	if err != nil {
		log.Println("Failed to generate token:", err)
		http.Error(w, fmt.Sprintf("Failed to generate token: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("User logged in successfully!")
	fmt.Fprintf(w, "User logged in successfully! Token: %v", token)
}

// GenerateToken Token generation when user logs in
func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "No token", err
	}

	return tokenString, nil
}
