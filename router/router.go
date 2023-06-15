package router

import (
	"github.com/BurdockBH/food-delivery-rest-service/service/user"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func InitializeRouter() *http.ServeMux {
	router := http.NewServeMux()

	//TODO: Add more routes here
	router.HandleFunc("/user/register", Chain(Post)(user.RegisterUser))
	router.HandleFunc("/user/login", Chain(Post)(user.LoginUser))
	router.HandleFunc("/user/delete-user", Chain(Delete)(user.DeleteUser))

	return router
}

func Chain(middlewares ...Middleware) func(http.HandlerFunc) http.HandlerFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}
		return handler
	}
}

func Post(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && r.Method != http.MethodOptions {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}

func Delete(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete && r.Method != http.MethodOptions {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}

func Put(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && r.Method != http.MethodOptions {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}
