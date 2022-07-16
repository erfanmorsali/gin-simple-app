package api

import (
	"github.com/erfanmorsali/gin-simple-app/posts/controllers"
	"github.com/erfanmorsali/gin-simple-app/posts/repositories"
	"github.com/erfanmorsali/gin-simple-app/posts/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunCategoriesApi(routerGroup *gin.RouterGroup, db *gorm.DB) {
	dao := repositories.NewCategoryDao(db)
	service := services.NewCategoryService(dao)
	controller := controllers.NewCategoryController(service)

	categoryGroup := routerGroup.Group("/categories")

	categoryGroup.GET("/", controller.GetAllCategories)
	categoryGroup.GET("/:id", controller.GetCategoryById)
	categoryGroup.POST("/", controller.CreateCategory)
	categoryGroup.DELETE("/:id", controller.DeleteCategoryById)
}
