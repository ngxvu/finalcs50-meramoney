package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm/schema"
)

func (s *Server) MigrationRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/migration", s.Migration)
	return r
}

func (s *Server) Migration(w http.ResponseWriter, r *http.Request) {

	models := []interface{}{}

	r.DB.Config.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "public.",
	}

	if err := r.DB.AutoMigrate(models...); err != nil {
		return err
	}
	return nil

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
