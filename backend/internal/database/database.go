package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDatabase initializes and returns a GORM DB instance.
func ConnectDatabase() (*gorm.DB, error) {
	// Load environment variables
	database := os.Getenv("DB_DATABASE")
	password := os.Getenv("DB_PASSWORD")
	username := os.Getenv("DB_USERNAME")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	schema := os.Getenv("DB_SCHEMA")

	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s", host, username, password, database, port, schema)

	// Open the database connection using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
