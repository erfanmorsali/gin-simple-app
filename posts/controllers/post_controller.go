package controllers

import (
	"github.com/erfanmorsali/gin-simple-app.git/posts/dtos"
	"github.com/erfanmorsali/gin-simple-app.git/posts/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostController struct {
	postService interfaces.PostService
}

func NewPostController(service interfaces.PostService) *PostController {
	return &PostController{postService: service}
}

func (c PostController) GetAllPosts(context *gin.Context) {
	posts, err := c.postService.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	context.JSON(http.StatusOK, posts)
	return
}

func (c PostController) GetPostById(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postOut, err := c.postService.GetById(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, postOut)
	return
}

func (c PostController) CreatePost(context *gin.Context) {
	var postIn dtos.PostIn
	err := context.ShouldBindJSON(&postIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post, err := c.postService.Create(postIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, post)
	return
}

func (c PostController) DeletePostById(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.postService.DeleteById(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, true)
	return
}
