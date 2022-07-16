package main

import (
	"github.com/erfanmorsali/gin-simple-app.git/api"
	"github.com/erfanmorsali/gin-simple-app.git/database"
	"github.com/erfanmorsali/gin-simple-app.git/middlewares"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	db, err := database.ConnectToDb()
	if err != nil {
		return
	}

	engine := gin.Default()

	apiGroup := engine.Group("/api")
	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middlewares.JwtAuthenticationMiddleware())

	publicGroup := apiGroup.Group("/pub")

	api.RunAdminApis(adminGroup, db)
	api.RunPublicApis(publicGroup, db)

	log.Fatalln(engine.Run(":8080"))

}
