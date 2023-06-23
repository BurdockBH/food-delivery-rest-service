package user

import (
	"encoding/json"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/router/helper"
	"github.com/BurdockBH/food-delivery-rest-service/statusCodes"
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
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = u.Validate()
	if err != nil {
		log.Println("Failed to validate login:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateLogin,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateLogin],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = user.RegisterUser(&u)
	if err != nil {
		log.Println("Failed to register user:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToCreateUser,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToCreateUser] + ": " + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyCreatedUser,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyCreatedUser] + ": " + u.Email,
	})
	helper.BaseResponse(w, response, http.StatusOK)
}

// LoginUser logs in a user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var userLogin viewmodels.UserLoginRequest

	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
	}
	defer r.Body.Close()

	err = userLogin.ValidateLogin()
	if err != nil {
		log.Println("Failed to validate login:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateLogin,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateLogin] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = user.LoginUser(&userLogin)
	if err != nil {
		log.Println("Failed to login user: ", userLogin.Email, err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToLoginUser,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToLoginUser] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	token, err := helper.GenerateToken(userLogin.Email)
	if err != nil {
		log.Println("Failed to generate token:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToGenerateToken,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToGenerateToken] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(viewmodels.LoginResponse{AccessToken: token, BaseResponse: viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyLoggedInUser,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyLoggedInUser] + ":" + userLogin.Email,
	}})
	if err != nil {
		log.Println("Failed to marshal json:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToMarshalJSON,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToMarshalJSON] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	helper.BaseResponse(w, jsonResponse, http.StatusOK)
}

// DeleteUser deletes a user from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Token not found")
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenNotFound,
			Message:    statusCodes.StatusCodes[statusCodes.TokenNotFound],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	claims, err := helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenValidationFailed,
			Message:    statusCodes.StatusCodes[statusCodes.TokenValidationFailed],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	var userLogin viewmodels.UserLoginRequest

	if _, ok := claims["email"]; !ok {
		log.Println("Invalid claim")
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.InvalidClaims,
			Message:    statusCodes.StatusCodes[statusCodes.InvalidClaims],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = userLogin.ValidateLogin()
	if err != nil {
		log.Println("Failed to validate login:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateLogin,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateLogin] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = user.DeleteUser(&userLogin)
	if err != nil {
		log.Println("Failed to delete user:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDeleteUser,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDeleteUser] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	log.Printf("User %v deleted successfully!\n", userLogin.Email)
	response, _ := json.Marshal(viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyDeletedUser,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyDeletedUser] + ":" + userLogin.Email,
	})
	helper.BaseResponse(w, response, http.StatusOK)
}

// EditUser edits a user in the database
func EditUser(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Println("Token not found")
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenNotFound,
			Message:    statusCodes.StatusCodes[statusCodes.TokenNotFound],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	claims, err := helper.ValidateToken(tokenString)
	if err != nil {
		log.Println("Token validation failed:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.TokenValidationFailed,
			Message:    statusCodes.StatusCodes[statusCodes.TokenValidationFailed] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	var u viewmodels.User

	if _, ok := claims["email"]; !ok {
		log.Println("Invalid claim")
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.InvalidClaims,
			Message:    statusCodes.StatusCodes[statusCodes.InvalidClaims],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = u.Validate()
	if err != nil {
		log.Println("Failed to validate login:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToValidateLogin,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToValidateLogin] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}

	err = user.EditUser(&u)
	if err != nil {
		log.Println("Failed to update user:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToUpdateUser,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToUpdateUser] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	token, err := helper.GenerateToken(u.Email)
	if err != nil {
		log.Println("Failed to generate token:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToGenerateToken,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToGenerateToken] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(viewmodels.LoginResponse{AccessToken: token, BaseResponse: viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyUpdatedUser,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyUpdatedUser] + ":" + u.Email,
	}})
	if err != nil {
		log.Println("Failed to marshal json:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToMarshalJSON,
		})
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
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToDecodeRequestBody,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToDecodeRequestBody],
		})
		helper.BaseResponse(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	users, err := user.GetUsers(&u)
	if err != nil {
		log.Println("Failed to get user:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToFetchUsers,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToFetchUsers] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(viewmodels.UserList{Users: users, BaseResponse: viewmodels.BaseResponse{
		StatusCode: statusCodes.SuccesfullyFetchedUsers,
		Message:    statusCodes.StatusCodes[statusCodes.SuccesfullyFetchedUsers],
	}})
	if err != nil {
		log.Println("Failed to marshal json:", err)
		response, _ := json.Marshal(viewmodels.BaseResponse{
			StatusCode: statusCodes.FailedToMarshalJSON,
			Message:    statusCodes.StatusCodes[statusCodes.FailedToMarshalJSON] + ":" + err.Error(),
		})
		helper.BaseResponse(w, response, http.StatusInternalServerError)
		return
	}

	log.Println("User retrieved successfully!")
	helper.BaseResponse(w, jsonResponse, http.StatusOK)
}
