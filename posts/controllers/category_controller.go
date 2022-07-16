package controllers

import (
	"github.com/erfanmorsali/gin-simple-app.git/posts/dtos"
	"github.com/erfanmorsali/gin-simple-app.git/posts/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type categoryController struct {
	categoryService interfaces.CategoryService
}

func NewCategoryController(categoryService interfaces.CategoryService) interfaces.CategoryController {
	return &categoryController{categoryService: categoryService}
}

func (c categoryController) GetAllCategories(context *gin.Context) {
	categories, err := c.categoryService.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	context.JSON(http.StatusOK, categories)
	return
}

func (c categoryController) GetCategoryById(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category, err := c.categoryService.GetById(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, category)
	return
}

func (c categoryController) CreateCategory(context *gin.Context) {
	var categoryIn dtos.CategoryIn

	err := context.ShouldBindJSON(&categoryIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category, err := c.categoryService.Create(categoryIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, category)
	return
}

func (c categoryController) DeleteCategoryById(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.categoryService.DeleteById(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, true)
	return
}
