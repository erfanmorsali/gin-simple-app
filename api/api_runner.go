package api

import (
	"github.com/erfanmorsali/gin-simple-app.git/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunAllApis(engine *gin.Engine, db *gorm.DB) {

	// Create Api Group ... All Apis Starts With /api
	apiGroup := engine.Group("/api")

	// Public Group ... Apis That Don't Need Authentication
	publicGroup := apiGroup.Group("/pub")
	RunAuthApi(publicGroup, db)

	// Admin Group ... Apis That Need Authentication
	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middlewares.JwtAuthenticationMiddleware())
	RunUsersApi(adminGroup, db)
	RunPostsApi(adminGroup, db)
	RunCategoriesApi(adminGroup, db)

}
