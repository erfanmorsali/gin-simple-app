package interfaces

import (
	"github.com/erfanmorsali/gin-simple-app.git/posts/dtos"
)

type CategoryService interface {
	GetAll() ([]dtos.CategoryOut, error)
	GetById(id uint) (*dtos.CategoryOut, error)
	Create(categoryIn dtos.CategoryIn) (*dtos.CategoryOut, error)
	DeleteById(id uint) error
}
