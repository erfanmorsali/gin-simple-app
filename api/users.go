package api

import (
	"github.com/erfanmorsali/gin-simple-app.git/users/controllers"
	"github.com/erfanmorsali/gin-simple-app.git/users/repositories"
	"github.com/erfanmorsali/gin-simple-app.git/users/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunUsersApi(engine *gin.Engine, db *gorm.DB) {
	dao := repositories.UserDao{Db: db}
	service := services.UserService{UserDao: dao}
	controller := controllers.UserController{UserService: service}

	usersGroup := engine.Group("/users")

	usersGroup.GET("/", controller.GetAllUsers)
}
