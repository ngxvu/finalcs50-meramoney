package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"meramoney/backend/internal/database"
	"meramoney/backend/internal/server"

	"github.com/gorilla/mux"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the database connection
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ensure the database connection is closed when the application exits
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from GORM DB: %v", err)
	}
	defer sqlDB.Close()

	// Initialize the server with the database connection
	srv := &server.Server{DB: db}

	// Set up the router
	r := mux.NewRouter()
	srv.Routes(r)

	// Initialize and start the HTTP server
	httpServer := server.NewServer(r)

	err = httpServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
