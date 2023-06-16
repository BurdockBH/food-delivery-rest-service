package user

import (
	"encoding/json"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"net/http"
	"strings"
)

// RegisterUser registers a new user in the database
func RegisterUser(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, fmt.Sprintf(`{"status" : "%v"}`, errString), http.StatusBadRequest)
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
	var userLogin viewmodels.UserLogin

	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, `{"status" : "Failed to decode request body"}`, http.StatusBadRequest)
	}
	defer r.Body.Close()

	validRequest, errString := userLogin.ValidateLogin()
	if !validRequest {
		log.Println(errString)
		http.Error(w, fmt.Sprintf(`{"status" : "%v"}`, errString), http.StatusBadRequest)
		return
	}

	err = user.LoginUser(userLogin)
	if err != nil {
		log.Println("Failed to login user: ", userLogin.Email, err)
		http.Error(w, fmt.Sprintf(`{"status" : "Failed to login user: %v %v"}`, userLogin.Email, err), http.StatusInternalServerError)
		return
	}

	token, err := helper.GenerateToken(userLogin.Email)
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

// DeleteUser deletes a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Token not found")
		http.Error(w, `{"status" : "Token not found"}`, http.StatusBadRequest)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	claims, err := helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		http.Error(w, fmt.Sprintf(`{"status" : "Token validation failed: %v"}`, err), http.StatusInternalServerError)
		return
	}

	var userLogin viewmodels.UserLogin

	validRequest, errString := userLogin.ValidateLogin()
	if !validRequest {
		log.Println(errString)
		http.Error(w, fmt.Sprintf(`{"status" : "%v"}`, errString), http.StatusBadRequest)
		return
	}

	if _, ok := claims["email"]; ok {

		err := json.NewDecoder(r.Body).Decode(&userLogin)
		if err != nil {
			log.Println("Failed to decode request body:", err)
			http.Error(w, `{"status" : "Failed to decode request body"}`, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = user.DeleteUser(userLogin)
		if err != nil {
			log.Println("Failed to delete user:", err)
			http.Error(w, fmt.Sprintf(`{"status" : "Failed to delete user: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		log.Println("User deleted successfully!")
		fmt.Fprintf(w, `{"status" : "User deleted successfully!"}`)
	} else {
		log.Println("Invalid token")
		http.Error(w, fmt.Sprintf(`{"status" : "Invalid token for user %v"}`, userLogin.Email), http.StatusBadRequest)
		return
	}
}

// EditUser edits a user in the database
func EditUser(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Token not found")
		http.Error(w, `{"status" : "Token not found"}`, http.StatusBadRequest)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	claims, err := helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		http.Error(w, fmt.Sprintf(`{"status" : "Token validation failed: %v"}`, err), http.StatusInternalServerError)
		return
	}

	var u viewmodels.User

	validRequest, errString := u.Validate()
	if !validRequest {
		log.Println(errString)
		http.Error(w, fmt.Sprintf(`{"status" : "%v"}`, errString), http.StatusBadRequest)
		return
	}

	if _, ok := claims["email"]; ok {
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println("Failed to decode request body:", err)
			http.Error(w, `{"status" : "Failed to decode request body"}`, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = user.EditUser(tokenString, u)
		if err != nil {
			log.Println("Failed to update user:", err)
			http.Error(w, fmt.Sprintf(`{"status" : "Failed to update user: %v"}`, err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		log.Println("User updated successfully!")
		fmt.Fprintf(w, `{"status" : "User updated successfully!"}`)
	} else {
		log.Println("Invalid token")
		http.Error(w, fmt.Sprintf(`{"status" : "Invalid token for user %v"}`, u.Email), http.StatusBadRequest)
		return
	}
}
