package viewmodels

import (
	"regexp"
)

// User is the user model
type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// UserLogin is the user login model
type UserLogin struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate validates the user
func (u *User) Validate() (bool, string) {
	if regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email) == false {
		return false, "Invalid email"
	} else if regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(u.Name) == false {
		return false, "Invalid name"
	} else if regexp.MustCompile(`^[0-9]+$`).MatchString(u.Phone) == false {
		return false, "Invalid phone"
	} else if len(u.Password) > 20 || len(u.Password) < 8 {
		return false, "Invalid password"
	}
	return true, ""
}

// ValidateLogin validates the user login credentials
func (u *UserLogin) ValidateLogin() (bool, string) {
	if regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email) == false {
		return false, "Invalid email"
	} else if len(u.Password) > 20 || len(u.Password) < 8 {
		return false, "Invalid password"
	}
	return true, ""
}
