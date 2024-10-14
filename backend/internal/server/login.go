package server

import (
	"encoding/json"
	"meramoney/backend/infrastructure/domains"
	auth "meramoney/backend/infrastructure/middlewares"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

// LoginHandler handles user login and returns JWT tokens
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database
	var user domains.User
	if err := s.DB.Where("user_name = ?", req.Username).First(&user).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate tokens
	accessToken, refreshToken, err := auth.GenerateTokens(user.ID, user.UserName)
	if err != nil {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}

	// Return tokens as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Logout logs out the user by invalidating their session or token
func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	// Assuming you are using a token-based authentication system
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token is required", http.StatusUnauthorized)
		return
	}

	// Invalidate the token (this could be done by removing it from a database or cache)
	if err := s.DB.Exec("DELETE FROM tokens WHERE token = ?", token).Error; err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}
