package interfaces

import (
	"github.com/erfanmorsali/gin-simple-app/database/models"
)

type UserDao interface {
	GetAll() ([]models.User, error)
	Create(user *models.User) (*models.User, error)
	GetById(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	Delete(user *models.User) error
}
