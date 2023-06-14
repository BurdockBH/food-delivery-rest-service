package user

import (
	"encoding/json"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"log"
	"net/http"
)

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
