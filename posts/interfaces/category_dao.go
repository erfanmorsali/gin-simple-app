package interfaces

import "github.com/erfanmorsali/gin-simple-app/database/models"

type CategoryDao interface {
	GetByIds(ids []uint) ([]models.Category, error)
	GetAll() ([]models.Category, error)
	GetById(id uint) (*models.Category, error)
	Create(category models.Category) (*models.Category, error)
	Delete(category models.Category) error
}
