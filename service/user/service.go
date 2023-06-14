package user

import (
	"encoding/json"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
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
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	v, errString := u.Validate()
	if !v {
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	err = user.RegisterUser(u)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User registered successfully!")
}
