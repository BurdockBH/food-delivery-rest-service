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
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: "Failed to decode request body"})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = u.Validate()
	if err != nil {
		log.Println("Failed to validate login:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to validate user: %v", err)})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = user.RegisterUser(u)
	if err != nil {
		log.Println("Failed to register user:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to register user: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("User %v registered successfully!", u.Name)})
	helper.BaseResponse(w, response, http.StatusOK)
}

// LoginUser logs in a user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLogin viewmodels.UserLoginRequest

	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: "Failed to decode request body"})
		helper.BaseResponse(w, response, http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = userLogin.ValidateLogin()
	if err != nil {
		log.Println("Failed to validate login:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to validate login: %v", err)})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = user.LoginUser(userLogin)
	if err != nil {
		log.Println("Failed to login user: ", userLogin.Email, err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to login user: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	token, err := helper.GenerateToken(userLogin.Email)
	if err != nil {
		log.Println("Failed to generate token:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to generate token: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(viewmodels.LoginResponse{AccessToken: token, BaseResponse: viewmodels.BaseResponse{Status: fmt.Sprintf("User %v logged in successfully!", userLogin.Email)}})
	if err != nil {
		log.Println("Failed to marshal json:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to marshal json: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println("User logged in successfully!")
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println("Failed to write response:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to write response: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}
}

// DeleteUser deletes a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Token not found")
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: "Token not found"})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	claims, err := helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Token validation failed: %v", err)})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	var userLogin viewmodels.UserLoginRequest

	err = json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: "Failed to decode request body"})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = userLogin.ValidateLogin()
	if err != nil {
		log.Println("Failed to validate login:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to validate login: %v", err)})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	if _, ok := claims["email"]; !ok {
		log.Println("Invalid claim")
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: "Invalid claim"})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: "Failed to decode request body"})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = user.DeleteUser(userLogin)
	if err != nil {
		log.Println("Failed to delete user:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to delete user: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	log.Println("User deleted successfully!")
	response, _ := json.Marshal(viewmodels.BaseResponse{Status: "User deleted successfully!"})
	helper.BaseResponse(w, response, http.StatusOK)
}

// EditUser edits a user in the database
func EditUser(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Token not found")
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: "Token not found"})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	claims, err := helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Token validation failed: %v", err)})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	var u viewmodels.User

	if _, ok := claims["email"]; !ok {
		log.Println("Invalid claim")
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: "Invalid claim"})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: "Failed to decode request body"})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = u.Validate()
	if err != nil {
		log.Println("Failed to validate login:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to validate login: %v", err)})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = user.EditUser(tokenString, u)
	if err != nil {
		log.Println("Failed to update user:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to update user: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	token, err := helper.GenerateToken(u.Email)
	if err != nil {
		log.Println("Failed to generate token:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to generate token: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(viewmodels.LoginResponse{AccessToken: token, BaseResponse: viewmodels.BaseResponse{Status: fmt.Sprintf("User %v updated successfully!", u.Email)}})
	if err != nil {
		log.Println("Failed to marshal json:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to marshal json: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	log.Println("User edited successfully!")
	helper.BaseResponse(w, jsonResponse, http.StatusOK)
}

// GetUsers gets a user from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var u viewmodels.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: "Failed to decode request body"})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	users, err := user.GetUsersByDetails(u)
	if err != nil {
		log.Println("Failed to get user:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to get user: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(viewmodels.UserList{Users: users, BaseResponse: viewmodels.BaseResponse{Status: "Users retrieved successfully!"}})
	if err != nil {
		log.Println("Failed to marshal json:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{Status: fmt.Sprintf("Failed to marshal json: %v", err)})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	log.Println("User retrieved successfully!")
	helper.BaseResponse(w, jsonResponse, http.StatusOK)
}
