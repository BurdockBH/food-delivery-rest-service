package user

import (
	"encoding/json"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/db/user"
	"github.com/BurdockBH/food-delivery-rest-service/viewmodels"
	"net/http"
	"regexp"
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

	Validate(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, &u.Email, &w)
	Validate(`^[a-zA-Z\s]+$`, &u.Name, &w)
	Validate(`^[0-9]{10}$`, &u.Phone, &w)
	Validate(`^[a-zA-Z0-9]+$`, &u.Password, &w)

	repository := user.UserRepository{}
	err = repository.RegisterUser(u)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User registered successfully!")
}

func Validate(s string, p *string, w *http.ResponseWriter) {
	if regexp.MustCompile(s).MatchString(*p) == false {
		http.Error(*w, "Invalid "+*p, http.StatusBadRequest)
		return
	}
}
