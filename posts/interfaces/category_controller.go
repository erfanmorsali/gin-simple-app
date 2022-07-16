package interfaces

import "github.com/gin-gonic/gin"

type CategoryController interface {
	GetAllCategories(context *gin.Context)
	GetCategoryById(context *gin.Context)
	CreateCategory(context *gin.Context)
	DeleteCategoryById(context *gin.Context)
}
