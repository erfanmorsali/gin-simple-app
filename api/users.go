package api

import (
	"github.com/erfanmorsali/gin-simple-app.git/users/controllers"
	"github.com/erfanmorsali/gin-simple-app.git/users/repositories"
	"github.com/erfanmorsali/gin-simple-app.git/users/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunUsersApi(routerGroup *gin.RouterGroup, db *gorm.DB) {
	dao := repositories.NewUserDao(db)
	service := services.NewUserService(dao)
	controller := controllers.NewUserController(service)

	usersGroup := routerGroup.Group("/users")

	usersGroup.GET("/", controller.GetAllUsers)
	usersGroup.GET("/:id", controller.GetUserById)
	usersGroup.POST("/", controller.CreateUser)
	usersGroup.DELETE("/:id", controller.DeleteUserById)
}
