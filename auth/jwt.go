package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint
	Role   string
	jwt.RegisteredClaims
}

func GenerateJWT(user User) (string, error) {
	// JWT generation logic
	return "", nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	// JWT validation logic
	return nil, nil
}
