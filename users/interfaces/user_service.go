package interfaces

import "github.com/erfanmorsali/gin-simple-app.git/users/dtos"

type UserService interface {
	GetAllUsers() []dtos.UserOut
	CreateUser(userIn dtos.UserIn) *dtos.UserOut
	GetUserById(id uint) (*dtos.UserOut, error)
	DeleteUserById(id uint) error
}
