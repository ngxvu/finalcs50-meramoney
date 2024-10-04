package server

import (
	"encoding/json"
	"meramoney/backend/infrastructure/domains"
	"net/http"
)

// GetProfile retrieves the user's profile information
func (s *Server) GetProfile(w http.ResponseWriter, r *http.Request) {
	// Assuming user information is stored in the context after authentication
	user, ok := r.Context().Value("user").(domains.User)
	if !ok {
		http.Error(w, "Failed to retrieve user information", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// UpdateProfile updates the user's profile information
func (s *Server) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// Assuming user information is stored in the context after authentication
	user, ok := r.Context().Value("user").(domains.User)
	if !ok {
		http.Error(w, "Failed to retrieve user information", http.StatusInternalServerError)
		return
	}

	var updatedUser domains.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update user information in the database
	if err := s.DB.Model(&user).Updates(updatedUser).Error; err != nil {
		http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}
