package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunAllApis(engine *gin.Engine, db *gorm.DB) {
	RunUsersApi(engine, db)
}
