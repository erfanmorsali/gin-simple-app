package api

import (
	"github.com/erfanmorsali/gin-simple-app/posts/controllers"
	"github.com/erfanmorsali/gin-simple-app/posts/repositories"
	"github.com/erfanmorsali/gin-simple-app/posts/services"
	userRepositories "github.com/erfanmorsali/gin-simple-app/users/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunPostsApi(routerGroup *gin.RouterGroup, db *gorm.DB) {
	dao := repositories.NewPostDao(db)
	userDao := userRepositories.NewUserDao(db)
	categoryDao := repositories.NewCategoryDao(db)
	service := services.NewPostService(dao, userDao, categoryDao)
	controller := controllers.NewPostController(service)

	postGroup := routerGroup.Group("/posts")

	postGroup.GET("/", controller.GetAllPosts)
	postGroup.GET("/:id", controller.GetPostById)
	postGroup.POST("/", controller.CreatePost)
	postGroup.DELETE("/:id", controller.DeletePostById)
}
