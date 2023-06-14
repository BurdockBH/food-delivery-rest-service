package viewmodels

import (
	"regexp"
)

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (*User) Validate() (bool, string) {
	if regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(User{}.Email) == false {
		return false, "Invalid email"
	} else if regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(User{}.Name) == false {
		return false, "Invalid name"
	} else if regexp.MustCompile(`^[0-9]{10}$`).MatchString(User{}.Phone) == false {
		return false, "Invalid phone"
	} else if regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(User{}.Password) == false {
		return false, "Invalid password"
	}
	return true, ""
}
