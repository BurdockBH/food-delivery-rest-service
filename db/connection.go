package db

import (
	"database/sql"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/config"
	_ "github.com/go-sql-driver/mysql"
)

// Connect to the database
func Connect(cfg *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Add validation for db fields
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
