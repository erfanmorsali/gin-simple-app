package interfaces

import "github.com/dgrijalva/jwt-go"

type JwtService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(encodedToken string) (*jwt.Token, error)
}
