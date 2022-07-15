package interfaces

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
)

type UserDao interface {
	GetAll() []models.User
	Create(user *models.User) (*models.User, error)
	GetById(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	Delete(user *models.User) error
}
