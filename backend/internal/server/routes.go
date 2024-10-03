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
	r.HandleFunc("/migrate", s.Migration).Methods("POST")
	r.HandleFunc("/sign-up", s.SignUp).Methods("POST")
	r.HandleFunc("/login", s.Login).Methods("POST")

	// CRUD for category

	// create category
	r.HandleFunc("/category", s.CreateCategory).Methods("POST")

	// get all categories
	r.HandleFunc("/category", s.GetAllCategories).Methods("GET")

	// get category by id
	r.HandleFunc("/category/{id}", s.GetCategory).Methods("GET")
	// update category
	r.HandleFunc("/category/{id}", s.UpdateCategory).Methods("PUT")

	// delete category
	r.HandleFunc("/category/{id}", s.DeleteCategory).Methods("DELETE")
}
