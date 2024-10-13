package domains

import "github.com/dgrijalva/jwt-go"

// Define a struct to hold the JWT claims
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
