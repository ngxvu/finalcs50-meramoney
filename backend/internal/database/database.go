package database

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDatabase initializes and returns a GORM DB instance.
func ConnectDatabase() (*gorm.DB, error) {
	// Load environment variables
	database := os.Getenv("DB_DATABASE")
	password := os.Getenv("DB_PASSWORD")
	username := os.Getenv("DB_USERNAME")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	sslmode := "disable" // Set sslmode to disable

	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, username, password, database, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: %w", err)
	}

	return db, nil
}
