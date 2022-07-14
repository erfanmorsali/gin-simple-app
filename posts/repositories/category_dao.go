package repositories

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"gorm.io/gorm"
)

type CategoryDao struct {
	Db *gorm.DB
}

func NewCategoryDao(db *gorm.DB) *CategoryDao {
	return &CategoryDao{Db: db}
}

func (d CategoryDao) GetCategoriesByIds(ids []uint) ([]models.Category, error) {
	var categories []models.Category

	if err := d.Db.Find(&categories, "id in (?)", ids).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
