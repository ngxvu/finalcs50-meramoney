package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

// NewServer initializes and returns a new HTTP server.
func NewServer(handler http.Handler) *http.Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080 // Default port if not specified
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
