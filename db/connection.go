package db

import (
	"database/sql"
	"fmt"
	"github.com/BurdockBH/food-delivery-rest-service/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

var DB *sql.DB

// Connect to the database
func Connect(cfg *config.DatabaseConfig) (*sql.DB, error) {
	err := ValidateDbConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid db config: %w", err)
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		defer db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	DB = db
	return db, nil
}

func RunMigrations() error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}
	if err := goose.Up(DB, "migrations"); err != nil {
		return err
	}
	return nil
}

func ValidateDbConfig(cfg *config.DatabaseConfig) error {
	if cfg.DBUsername == "" {
		return fmt.Errorf("db username is required")
	}
	if cfg.DBPassword == "" {
		return fmt.Errorf("db password is required")
	}
	if cfg.DBHost == "" {
		return fmt.Errorf("db host is required")
	}
	if cfg.DBPort == "" {
		return fmt.Errorf("db port is required")
	}
	if cfg.DBName == "" {
		return fmt.Errorf("db name is required")
	}
	return nil
}
