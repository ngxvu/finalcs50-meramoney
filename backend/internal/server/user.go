package server

import (
	"encoding/json"
	"meramoney/backend/infrastructure/domains"
	"net/http"
)

func (s *Server) GetProfile(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("user").(string)
	if !ok {
		http.Error(w, "Failed to retrieve user information", http.StatusInternalServerError)
		return
	}

	var user domains.User
	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("user").(string)
	if !ok {
		http.Error(w, "Failed to retrieve user information", http.StatusInternalServerError)
		return
	}

	var updatedUser domains.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := s.DB.Model(&domains.User{}).Where("username = ?", username).Updates(updatedUser).Error; err != nil {
		http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}
