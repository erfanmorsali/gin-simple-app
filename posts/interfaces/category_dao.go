package interfaces

import "github.com/erfanmorsali/gin-simple-app.git/database/models"

type CategoryDao interface {
	GetCategoriesByIds(ids []uint) ([]models.Category, error)
}
