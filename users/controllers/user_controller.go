package controllers

import (
	"github.com/erfanmorsali/gin-simple-app.git/users/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService interfaces.UserService
}

func NewUserController(userService interfaces.UserService) *UserController {
	return &UserController{userService}
}

func (c UserController) GetAllUsers(context *gin.Context) {
	context.JSON(http.StatusOK, c.UserService.GetAllUsers())
	return
}
