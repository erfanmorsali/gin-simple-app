package dtos

import "github.com/dgrijalva/jwt-go"

type CustomClaim struct {
	Email string `json:"email"`
	User  bool   `json:"user"`
	jwt.StandardClaims
}
