package api

import (
	"github.com/erfanmorsali/gin-simple-app.git/auth/controllers"
	"github.com/erfanmorsali/gin-simple-app.git/auth/services"
	"github.com/erfanmorsali/gin-simple-app.git/users/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunAuthApi(routerGroup *gin.RouterGroup, db *gorm.DB) {
	userDao := repositories.NewUserDao(db)
	authService := services.NewAuthService(userDao)
	jwtService := services.NewJwtService()
	controller := controllers.NewAuthController(authService, jwtService)

	postGroup := routerGroup.Group("")

	postGroup.POST("/login", controller.Login)
	postGroup.POST("/register", controller.Register)
}
