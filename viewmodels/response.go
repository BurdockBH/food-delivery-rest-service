package viewmodels

// BaseResponse is the response model
type BaseResponse struct {
	Status string `json:"status"`
}

type UserList struct {
	BaseResponse
	Users []User `json:"users"`
}

// LoginResponse is the response model for login
type LoginResponse struct {
	BaseResponse
	AccessToken string `json:"token"`
}
