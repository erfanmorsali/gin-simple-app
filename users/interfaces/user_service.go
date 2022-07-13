package interfaces

import "github.com/erfanmorsali/gin-simple-app.git/users/models"

type UserService interface {
	GetAllUsers() []models.UserOut
	CreateUser(userIn models.UserIn) *models.UserOut
	GetUserById(id uint) (*models.UserOut, error)
	DeleteUserById(id uint) error
}
