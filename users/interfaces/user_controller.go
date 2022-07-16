package interfaces

import "github.com/gin-gonic/gin"

type UserController interface {
	GetAllUsers(context *gin.Context)
	CreateUser(context *gin.Context)
	GetUserById(context *gin.Context)
	DeleteUserById(context *gin.Context)
}
