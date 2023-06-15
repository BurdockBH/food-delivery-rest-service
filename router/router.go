package router

import (
	"github.com/BurdockBH/food-delivery-rest-service/service/user"
	"net/http"
)

func InitializeRouter() *http.ServeMux {
	router := http.NewServeMux()

	//TODO: Add more routes here
	router.HandleFunc("/user/register", user.RegisterUser)
	router.HandleFunc("/user/login", user.LoginUser)
	router.HandleFunc("/user/delete-user", user.DeleteUser)
	router.HandleFunc("/api/delete-user", user.DeleteUser)

	return router
}
