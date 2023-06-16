package viewmodels

// Response is the response model
type Response struct {
	Status string `json:"status"`
}

// LoginResponse is the response model for login
type LoginResponse struct {
	Response
	AccessToken string `json:"token"`
}
