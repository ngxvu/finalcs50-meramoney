package server

import (
	"encoding/json"
	"meramoney/backend/infrastructure/domains"
	"net/http"
)

type ProfileResponse struct {
	Username string `json:"user_name"`
}

func (s *Server) GetProfile(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("user").(string)
	if !ok {
		http.Error(w, "Failed to retrieve user information", http.StatusInternalServerError)
		return
	}

	profile, err := s.Profile(username)
	if err != nil {
		http.Error(w, "Failed to retrieve user profile", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(profile)
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

	if err := s.DB.Model(&domains.User{}).Where("user_name = ?", username).Updates(updatedUser).Error; err != nil {
		http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
		return
	}

	profile, err := s.Profile(updatedUser.UserName)
	if err != nil {
		http.Error(w, "Failed to retrieve user profile", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(profile)
}

func (s *Server) Profile(username string) (*ProfileResponse, error) {
	var user domains.User

	if err := s.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	// Create a ProfileResponse without the password field
	userProfile := &ProfileResponse{
		Username: user.UserName,
	}

	return userProfile, nil

}
