package server

import (
	"encoding/json"
	"net/http"

	"meramoney/backend/infrastructure/domains"
	auth "meramoney/backend/infrastructure/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create the user
	user := domains.User{
		UserName: req.Username,
		Password: string(hashedPassword),
	}

	// Save the user to the database
	if err := s.DB.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	if err := s.DB.Where("user_name = ?", req.Username).First(&user).Error; err != nil {
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
