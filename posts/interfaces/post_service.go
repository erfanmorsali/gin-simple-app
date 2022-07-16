package interfaces

import "github.com/erfanmorsali/gin-simple-app/posts/dtos"

type PostService interface {
	GetById(id uint) (*dtos.PostOut, error)
	GetAll() ([]dtos.PostOut, error)
	Create(postIn dtos.PostIn) (*dtos.PostOut, error)
	DeleteById(id uint) error
}
