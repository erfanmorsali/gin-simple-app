package interfaces

import "github.com/gin-gonic/gin"

type PostController interface {
	GetAllPosts(context *gin.Context)
	GetPostById(context *gin.Context)
	CreatePost(context *gin.Context)
	DeletePostById(context *gin.Context)
}
