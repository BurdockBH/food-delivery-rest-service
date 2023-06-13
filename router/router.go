package router

import (
	"github.com/BurdockBH/food-delivery-rest-service/service/user"
	"net/http"
)

func InitializeRouter() *http.ServeMux {
	router := http.NewServeMux()

	//TODO: Add more routes here
	router.HandleFunc("/api/register", user.RegisterUser)

	return router
}
