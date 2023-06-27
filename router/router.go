package router

import (
	"github.com/BurdockBH/food-delivery-rest-service/router/middlewares"
	"github.com/BurdockBH/food-delivery-rest-service/service/food_venue"
	"github.com/BurdockBH/food-delivery-rest-service/service/product"
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

	// Food venue routes
	router.HandleFunc("/food-venues/create", middlewares.Chain(middlewares.Post)(food_venue.CreateFoodVenue))
	router.HandleFunc("/food-venues/delete", middlewares.Chain(middlewares.Delete)(food_venue.DeleteFoodVenue))
	router.HandleFunc("/food-venues/get", middlewares.Chain(middlewares.Get)(food_venue.GetFoodVenues))

	// Product routes
	router.HandleFunc("/product/create", middlewares.Chain(middlewares.Post)(product.CreateProduct))
	router.HandleFunc("/product/delete", middlewares.Chain(middlewares.Delete)(product.DeleteProduct))
	router.HandleFunc("/product/edit", middlewares.Chain(middlewares.Put)(product.EditProduct))
	router.HandleFunc("/products/get", middlewares.Chain(middlewares.Get)(product.GetProducts))

	return router
}
