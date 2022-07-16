package main

import (
	"github.com/erfanmorsali/gin-simple-app.git/api"
	"github.com/erfanmorsali/gin-simple-app.git/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	// Connect To Database
	db, err := database.ConnectToDb()
	if err != nil {
		return
	}

	// Create Gin Engine
	engine := gin.Default()

	// Run All Apis
	api.RunAllApis(engine, db)

	// Run Server
	log.Fatalln(engine.Run(":8080"))

}
