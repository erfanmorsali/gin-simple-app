package interfaces

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
)

type UserDao interface {
	GetAll() []models.User
	CreateUser(user *models.User) *models.User
	GetById(id uint) (*models.User, error)
	DeleteUser(user *models.User) error
}
