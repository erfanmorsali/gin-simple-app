package controllers

import (
	"github.com/erfanmorsali/gin-simple-app.git/users/dtos"
	"github.com/erfanmorsali/gin-simple-app.git/users/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (c UserController) CreateUser(context *gin.Context) {
	var userIn dtos.UserIn

	err := context.ShouldBindJSON(&userIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userOut, err := c.UserService.CreateUser(userIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, userOut)
	return
}

func (c UserController) GetUserById(context *gin.Context) {
	pathId := context.Param("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input for id",
		})
		return
	}

	userOut, err := c.UserService.GetUserById(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, userOut)
	return
}

func (c UserController) DeleteUserById(context *gin.Context) {
	pathId := context.Param("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.UserService.DeleteUserById(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, true)
}
