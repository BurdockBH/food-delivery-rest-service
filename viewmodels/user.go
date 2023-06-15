package viewmodels

import (
	"regexp"
)

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

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
	} else if regexp.MustCompile(`^(?=.*[a-zA-Z0-9!@#$%^&*()-_=+{};:,<.>]).{8,20}$`).MatchString(u.Password) == false {
		return false, "Invalid password"
	}
	return true, ""
}

func (u *UserLogin) ValidateLogin() (bool, string) {
	if regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email) == false {
		return false, "Invalid email"
	} else if regexp.MustCompile(`^(?=.*[a-zA-Z0-9!@#$%^&*()-_=+{};:,<.>]).{8,20}$`).MatchString(u.Password) == false {
		return false, "Invalid password"
	}
	return true, ""
}
