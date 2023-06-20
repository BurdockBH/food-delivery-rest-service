package main

import (
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/config"
	"github.com/BurdockBH/food-delivery-rest-service/db"
	"github.com/BurdockBH/food-delivery-rest-service/router"
	"log"
	"net/http"
)

// main is the entry point of the application
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbConnected, err := db.Connect(cfg.DatabaseConfig)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer dbConnected.Close()

	err = db.RunMigrations()
	if err != nil {
		log.Fatal(err)
	}

	r := router.InitializeRouter()
	fmt.Print("Server is running on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
