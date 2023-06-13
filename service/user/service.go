package user

import (
	"fmt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	// Handler logic goes here

	// RegisterUser()
	fmt.Fprintf(w, "User registered successfully!")
}
