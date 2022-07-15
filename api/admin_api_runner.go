package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunAdminApis(adminRouterGroup *gin.RouterGroup, db *gorm.DB) {
	RunUsersApi(adminRouterGroup, db)
	RunPostsApi(adminRouterGroup, db)
	RunCategoriesApi(adminRouterGroup, db)
}
