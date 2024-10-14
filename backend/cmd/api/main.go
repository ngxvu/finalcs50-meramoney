package main

import (
	"fmt"
	"log"
	"meramoney/backend/internal/database"
	"meramoney/backend/internal/server"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gorilla/handlers"
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

	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
	// Define CORS options
	corsOptions := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Initialize and start the HTTP server with CORS support
	httpServer := server.NewServer(handlers.CORS(corsOptions, corsHeaders, corsMethods)(r))

	uploadsDir := "uploads"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadsDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create uploads directory: %v", err)
		}
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
