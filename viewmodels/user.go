package viewmodels

import (
	"regexp"
	"time"
)

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (*User) Validate(u *User) (bool, string) {
	if regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email) == false {
		return false, "Invalid email"
	} else if regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(u.Name) == false {
		return false, "Invalid name"
	} else if regexp.MustCompile(`^[0-9]{10}$`).MatchString(u.Phone) == false {
		return false, "Invalid phone"
	} else if regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(u.Password) == false {
		return false, "Invalid password"
	}
	return true, ""
}
