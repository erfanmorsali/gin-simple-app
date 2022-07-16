package interfaces

import "github.com/erfanmorsali/gin-simple-app.git/users/dtos"

type UserService interface {
	GetAll() ([]dtos.UserOut, error)
	Create(userIn dtos.UserIn) (*dtos.UserOut, error)
	GetById(id uint) (*dtos.UserOut, error)
	DeleteById(id uint) error
}
