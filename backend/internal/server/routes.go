package server

import (
	auth "meramoney/backend/infrastructure/middlewares"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Server struct with a reference to the database connection
type Server struct {
	DB *gorm.DB
}

func (s *Server) Routes(r *mux.Router) {
	r.HandleFunc("/upload", s.UploadHandler).Methods("POST")
	r.HandleFunc("/migrate", s.Migration).Methods("POST")
	r.HandleFunc("/sign-up", s.SignUp).Methods("POST")
	r.HandleFunc("/login", s.Login).Methods("POST")

	// Create a subrouter for protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(auth.VerifyUserMiddleware)

	// Protected routes
	protected.HandleFunc("/logout", s.Logout).Methods("POST")

	// User Profile Management
	protected.HandleFunc("/profile", s.GetProfile).Methods("GET")
	protected.HandleFunc("/profile", s.UpdateProfile).Methods("PUT")

	// CRUD for category
	protected.HandleFunc("/category", s.CreateCategory).Methods("POST")
	protected.HandleFunc("/category", s.GetAllCategories).Methods("GET")
	protected.HandleFunc("/category/{id}", s.GetCategory).Methods("GET")
	protected.HandleFunc("/category/{id}", s.UpdateCategory).Methods("PUT")
	protected.HandleFunc("/category/{id}", s.DeleteCategory).Methods("DELETE")

	// CRUD for transaction income and expense
	protected.HandleFunc("/transaction", s.CreateTransaction).Methods("POST")
	protected.HandleFunc("/transaction", s.GetAllTransactions).Methods("GET")
	protected.HandleFunc("/transaction/{id}", s.GetTransaction).Methods("GET")
	protected.HandleFunc("/transaction/{id}", s.UpdateTransaction).Methods("PUT")
	protected.HandleFunc("/transaction/{id}", s.DeleteTransaction).Methods("DELETE")

	// Balance and total APIs
	protected.HandleFunc("/balance", s.GetBalance).Methods("GET")
	protected.HandleFunc("/total-income", s.GetTotalIncome).Methods("GET")
	protected.HandleFunc("/total-expense", s.GetTotalExpense).Methods("GET")

}
