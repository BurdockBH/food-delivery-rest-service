package viewmodels

type Response struct {
	Status string `json:"status"`
}

type LoginResponse struct {
	Response
	AccessToken string `json:"token"`
}
