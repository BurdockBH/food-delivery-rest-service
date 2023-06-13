package main

import (
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/config"
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/router"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	_, err = db.Connect(cfg)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	r := router.InitializeRouter()
	fmt.Print("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
