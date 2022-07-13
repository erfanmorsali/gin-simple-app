package main

import (
	"github.com/erfanmorsali/gin-simple-app.git/api"
	"github.com/erfanmorsali/gin-simple-app.git/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	db, err := database.ConnectToDb()
	if err != nil {
		return
	}

	api.RunAllApis(engine, db)

	log.Fatalln(engine.Run(":8080"))

}
