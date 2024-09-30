package server

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Server struct with a reference to the database connection
type Server struct {
	DB *gorm.DB
}

func (s *Server) Routes(r *mux.Router) {
	r.HandleFunc("/migrate", s.MigrationHandler).Methods("POST")
	r.HandleFunc("/sign-up", s.SignUpHandler).Methods("POST")
	r.HandleFunc("/login", s.LoginHandler).Methods("POST")
}
