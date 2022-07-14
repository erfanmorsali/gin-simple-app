package interfaces

import "github.com/erfanmorsali/gin-simple-app.git/posts/dtos"

type PostService interface {
	GetPostById(id uint) (*dtos.PostOut, error)
	GetAllPosts() ([]dtos.PostOut, error)
	CreatePost(postIn dtos.PostIn) (*dtos.PostOut, error)
	DeletePostById(id uint) error
}
