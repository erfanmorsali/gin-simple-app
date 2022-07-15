package controllers

import (
	"github.com/erfanmorsali/gin-simple-app.git/posts/dtos"
	"github.com/erfanmorsali/gin-simple-app.git/posts/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryController struct {
	CategoryService interfaces.CategoryService
}

func NewCategoryController(categoryService interfaces.CategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

func (c CategoryController) GetAllCategories(context *gin.Context) {
	categories, err := c.CategoryService.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	context.JSON(http.StatusOK, categories)
	return
}

func (c CategoryController) GetCategoryById(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category, err := c.CategoryService.GetById(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, category)
	return
}

func (c CategoryController) CreateCategory(context *gin.Context) {
	var categoryIn dtos.CategoryIn

	err := context.ShouldBindJSON(&categoryIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category, err := c.CategoryService.Create(categoryIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, category)
	return
}

func (c CategoryController) DeleteCategoryById(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.CategoryService.DeleteById(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, true)
	return
}