package interfaces

import "github.com/erfanmorsali/gin-simple-app.git/database/models"

type PostDao interface {
	GetAll() ([]*models.Post, error)
	GetById(id uint) (*models.Post, error)
	Create(post models.Post) (*models.Post, error)
	Delete(post models.Post) error
}
