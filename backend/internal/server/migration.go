package server

import (
	"encoding/json"
	"meramoney/backend/infrastructure/domains"
	"net/http"

	"gorm.io/gorm/schema"
)

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
