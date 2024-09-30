package server

import (
	"encoding/json"
	"meramoney/backend/infrastructure/domains"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Server struct with a reference to the database connection
type Server struct {
	DB *gorm.DB
}

func (s *Server) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/migrate", s.Migration).Methods("POST")
}

func (s *Server) Migration(w http.ResponseWriter, r *http.Request) {
	models := []interface{}{
		domains.User{},
		domains.Category{},
		domains.Transaction{},
	}

	s.DB.Config.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "public.",
	}

	if err := s.DB.AutoMigrate(models...); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Migration successful"})
}
