package router

import (
	"net/http"

	"github.com/BurdockBH/food-delivery-rest-service/service"
)

func InitializeRouter() *http.ServeMux {
	router := http.NewServeMux()

	newService := service.NewService()

	//TODO: Add more routes here
	router.HandleFunc("/api/register", newService.RegisterUser)

	return router
}
