package router

import (
	"github.com/BurdockBH/food-delivery-rest-service/router/middlewares"
	"github.com/BurdockBH/food-delivery-rest-service/service/user"
	"net/http"
)

func InitializeRouter() *http.ServeMux {
	router := http.NewServeMux()

	//TODO: Add more routes here
	router.HandleFunc("/user/register", middlewares.Chain(middlewares.Post)(user.RegisterUser))
	router.HandleFunc("/user/login", middlewares.Chain(middlewares.Post)(user.LoginUser))
	router.HandleFunc("/user/delete-user", middlewares.Chain(middlewares.Delete)(user.DeleteUser))
	router.HandleFunc("/user/edit-user", middlewares.Chain(middlewares.Put)(user.EditUser))

	return router
}
