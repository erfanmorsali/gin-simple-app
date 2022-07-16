package interfaces

import "github.com/erfanmorsali/gin-simple-app/auth/dtos"

type AuthService interface {
	Login(email string, password string) bool
	Register(registerInput dtos.RegisterIn) (*dtos.RegisterOut, error)
}
