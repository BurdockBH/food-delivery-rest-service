package router

import (
	"github.com/BurdockBH/food-delivery-rest-service/router/middlewares"
	"github.com/BurdockBH/food-delivery-rest-service/service/user"
	"net/http"
)

// InitializeRouter initializes the router
func InitializeRouter() *http.ServeMux {
	router := http.NewServeMux()

	// User routes
	router.HandleFunc("/users/register", middlewares.Chain(middlewares.Post)(user.RegisterUser))
	router.HandleFunc("/users/login", middlewares.Chain(middlewares.Post)(user.LoginUser))
	router.HandleFunc("/users/delete", middlewares.Chain(middlewares.Delete)(user.DeleteUser))
	router.HandleFunc("/users/edit", middlewares.Chain(middlewares.Put)(user.EditUser))
	router.HandleFunc("/users/get", middlewares.Chain(middlewares.Get)(user.GetUsers))

	return router
}
