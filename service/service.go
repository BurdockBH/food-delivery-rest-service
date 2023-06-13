package service

import (
	"fmt"
	"net/http"
)

type User struct {
	//TODO: Add user fields here
}

func NewService() *User {
	return &User{}
}

func (s *User) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "User registered successfully!")

	} else {
		http.NotFound(w, r)
	}
}
