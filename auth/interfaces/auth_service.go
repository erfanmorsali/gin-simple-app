package interfaces

import "github.com/erfanmorsali/gin-simple-app.git/auth/dtos"

type AuthService interface {
	Login(email string, password string) bool
	Register(registerInput dtos.RegisterIn) (*dtos.RegisterOut, error)
}
