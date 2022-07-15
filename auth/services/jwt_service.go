package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/erfanmorsali/gin-simple-app.git/auth/dtos"
	"time"
)

type JwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: getSecretKey(),
		issuer:    "test_issuer",
	}
}

func getSecretKey() string {
	secret := "erfan123456789" // TODO: Read From Env
	return secret
}

func (service *JwtService) GenerateToken(email string, isUser bool) string {
	claims := &dtos.CustomClaim{
		Email: email,
		User:  isUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
}