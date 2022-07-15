package repositories

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"gorm.io/gorm"
)

type CategoryDao struct {
	db *gorm.DB
}

func NewCategoryDao(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db: db}
}

func (d CategoryDao) GetByIds(ids []uint) ([]models.Category, error) {
	var categories []models.Category

	if err := d.db.Find(&categories, "id in (?)", ids).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (d CategoryDao) GetAll() ([]models.Category, error) {
	var categories []models.Category

	if err := d.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (d CategoryDao) GetById(id uint) (*models.Category, error) {
	var category models.Category

	if err := d.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (d CategoryDao) Create(category models.Category) (*models.Category, error) {
	if err := d.db.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (d CategoryDao) Delete(category models.Category) error {
	if err := d.db.Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
