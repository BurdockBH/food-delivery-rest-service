package viewmodels

// BaseResponse is the response model
type BaseResponse struct {
	StatusCode int64  `json:"status code"`
	Message    string `json:"message"`
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
