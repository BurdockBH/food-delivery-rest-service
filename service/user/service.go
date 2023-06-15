package user

import (
	"encoding/json"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

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
		http.Error(w, `{"status": "Failed to decode request body"}`, http.StatusBadRequest)
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
		http.Error(w, fmt.Sprintf(`{"status" : "Failed to register user: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("User registered successfully!")
	fmt.Fprintf(w, `{"status" : "User registered successfully!"}`)
}

// LoginUser logs in a user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	var userLogin viewmodels.UserLogin

	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, `{"status" : "Failed to decode request body"}`, http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = user.LoginUser(userLogin)
	if err != nil {
		log.Println("Failed to login user: ", userLogin.Email, err)
		http.Error(w, fmt.Sprintf(`{"status" : "Failed to login user: %v %v"}`, userLogin.Email, err), http.StatusInternalServerError)
		return
	}

	token, err := GenerateToken(userLogin.Email)
	if err != nil {
		log.Println("Failed to generate token:", err)
		http.Error(w, fmt.Sprintf(`"status" : "Failed to generate token: %v"}`, err), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(viewmodels.LoginResponse{AccessToken: token, Response: viewmodels.Response{Status: fmt.Sprintf("User %v logged in successfully!", userLogin.Email)}})
	if err != nil {
		log.Println("Failed to marshal json:", err)
		http.Error(w, fmt.Sprintf(`{"stauts" : "Failed to marshal json: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("User logged in successfully!")
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println("Failed to write response:", err)
		http.Error(w, fmt.Sprintf(`{"status" : "Failed to write response: %v"`, err), http.StatusInternalServerError)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.NotFound(w, r)
		return
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Token not found")
		http.Error(w, "Token not found", http.StatusBadRequest)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file:", err)
		http.Error(w, "Failed to load .env file", http.StatusInternalServerError)
		return
	}

	// Retrieve the JWT secret key from the environment variable
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		log.Println("Failed to parse token:", err)
		http.Error(w, fmt.Sprintf(`{"status": "Failed to parse token: %v"}`, err), http.StatusInternalServerError)
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var u viewmodels.UserLogin

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println("Failed to decode request body:", err)
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = user.DeleteUser(u)
		if err != nil {
			log.Println("Failed to delete user:", err)
			http.Error(w, fmt.Sprintf("Failed to delete user: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		log.Println("User deleted successfully!")
		fmt.Fprintf(w, "User deleted successfully!")
	} else {
		log.Println("Invalid token")
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}
}

// GenerateToken Token generation when user logs in
func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	err := godotenv.Load()
	if err != nil {
		return "Failed to load .env", err
	}
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		println("Failed to sign token:", err)
		return "Failed to sign token:", err
	}

	return tokenString, nil
}
