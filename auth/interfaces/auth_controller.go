package interfaces

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}
