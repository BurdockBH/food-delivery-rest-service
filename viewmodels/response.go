package viewmodels

// BaseResponse is the response model
type BaseResponse struct {
	Status string `json:"status"`
}

// LoginResponse is the response model for login
type LoginResponse struct {
	BaseResponse
	AccessToken string `json:"token"`
}
