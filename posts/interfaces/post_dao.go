package interfaces

import "github.com/erfanmorsali/gin-simple-app.git/database/models"

type PostDao interface {
	GetAll() ([]*models.Post, error)
	GetById(id uint) (*models.Post, error)
	CreatePost(post models.Post) (*models.Post, error)
	DeletePost(post models.Post) error
}
