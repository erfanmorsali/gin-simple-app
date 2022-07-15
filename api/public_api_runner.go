package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunPublicApis(publicRouterGroup *gin.RouterGroup, db *gorm.DB) {
	RunAuthApi(publicRouterGroup, db)
}
