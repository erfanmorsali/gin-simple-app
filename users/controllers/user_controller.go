package controllers

import (
	"github.com/erfanmorsali/gin-simple-app.git/users/dtos"
	"github.com/erfanmorsali/gin-simple-app.git/users/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userController struct {
	userService interfaces.UserService
}

func NewUserController(userService interfaces.UserService) interfaces.UserController {
	return &userController{userService: userService}
}

func (c userController) GetAllUsers(context *gin.Context) {
	context.JSON(http.StatusOK, c.userService.GetAll())
	return
}

func (c userController) CreateUser(context *gin.Context) {
	var userIn dtos.UserIn

	err := context.ShouldBindJSON(&userIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userOut, err := c.userService.Create(userIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, userOut)
	return
}

func (c userController) GetUserById(context *gin.Context) {
	pathId := context.Param("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input for id",
		})
		return
	}

	userOut, err := c.userService.GetById(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, userOut)
	return
}

func (c userController) DeleteUserById(context *gin.Context) {
	pathId := context.Param("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.userService.DeleteById(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, true)
}
