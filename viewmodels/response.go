package viewmodels

type LoginResponse struct {
	Status      string `json:"status"`
	AccessToken string `json:"access_token"`
}
